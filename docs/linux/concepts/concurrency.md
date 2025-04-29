---
title: "Concurrency"
description: "This page explains how to deal with concurrency in eBPF programs. It explains multiple methods to deal with concurrency, their pros and cons, and when to use them."
---
# Concurrency

Concurrency in the BPF world is something to be aware of when writing BPF programs. A BPF program can be seen as a function called by the kernel, thus the same program can in theory be invoked concurrently by every kernel thread. The only guarantee given by the kernel is that the same program invocation always runs on the same logical CPU.

This is particularly important when accessing memory that is shared between multiple programs or invocations of the same program such as non-per-CPU maps and kernel memory. Accesses and modifications to such kinds of memory are subject to [race conditions](https://en.wikipedia.org/wiki/Race_condition). Same goes for programs and userspace accessing the same map value at the same time.

There are a few methods to avoid race conditions.

## Atomic operations

Atomic operations refers to atomic CPU instructions. A normal `i += 1` operation will at some level break down into:

1. Read `i` into some CPU register
2. Increment the CPU register with `1`
3. Write the register value back to `i`

Since this happens in multiple steps, even such a simple operation is subject to a [race condition](https://en.wikipedia.org/wiki/Race_condition).

There is a class of CPU instructions that can perform specific tasks in a single CPU instruction which is serialized at the hardware level. These are also available in BPF. When compiling with Clang/LLVM these special instructions can be accessed via a list of special builtin functions:

* `__sync_fetch_and_add(*a, b)` - Read value at `a`, add `b` and write it back, return the original value of `a`
* `__sync_fetch_and_sub(*a, b)` - Read value at `a`, subtract a number and write it back, return the original value of `a`
* `__sync_fetch_and_or(*a, b)` - Read value at `a`, binary OR a number and write it back, return the original value of `a` :octicons-tag-24: [v5.12](https://lwn.net/ml/linux-kernel/20210114181751.768687-1-jackmanb@google.com/)
* `__sync_fetch_and_xor(*a, b)` - Read value at `a`, binary XOR a number and write it back, return the original value of `a` :octicons-tag-24: [v5.12](https://lwn.net/ml/linux-kernel/20210114181751.768687-1-jackmanb@google.com/)
* `__sync_val_compare_and_swap(*a, b, c)` - Read value at `a`, check if it is equal to `b`, if true write `c` to `a` and return the original value of `a`. On fail leave `a` be and return `c`. :octicons-tag-24: [v5.12](https://lwn.net/ml/linux-kernel/20210114181751.768687-1-jackmanb@google.com/)
* `__sync_lock_test_and_set(*a, b)` - Read value at `a`, write `b` to `a`, return original value of `a` :octicons-tag-24: [v5.12](https://lwn.net/ml/linux-kernel/20210114181751.768687-1-jackmanb@google.com/)

If you want to perform one of the above sequences on a variable you can do so with the atomic builtin functions. A common example is to increment a shared counter with `__sync_fetch_and_add`.

Atomic instructions work on variable of 1, 2, 4, or 8 bytes. Any variables larger than that such as multiple struct fields require multiple atomic instructions or other synchronization mechanisms.

Here is a simple example using atomic instructions to count the number of times the `sys_enter` tracepoint is called.

```c
int counter = 0;

SEC("tp_btf/sys_enter")
int sys_enter_count(void *ctx) {
	__sync_fetch_and_add(&counter, 1);
	return 0;
}
```

!!! note
	Atomic instructions still synchronize at the hardware level, so using atomic instructions will still decrease performance compared to its non-atomic variant.

## Spin locks

A common technique in the kernel for synchronization is a [spinlock](https://en.wikipedia.org/wiki/Spinlock). eBPF also provides spinlock capabilities for map values. The main advantage of spinlocks over atomic instructions is that it guarantees multiple fields are updated together.

To use spin locks, you first have to include a `struct bpf_spin_lock` at the top of your map value.

```c
struct concurrent_element {
	struct bpf_spin_lock semaphore;
	int count;
}

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__type(key, int);
	__type(value, struct concurrent_element);
	__uint(max_entries, 100);
} concurrent_map SEC(".maps");
```

Then in your code, you can take the lock with [`bpf_spin_lock`](../helper-function/bpf_spin_lock.md), do whatever you need to do, and release the lock with [`bpf_spin_unlock`](../helper-function/bpf_spin_unlock.md). In this example, we simply increment the number of times the `sys_enter` tracepoint is called.

```c
SEC("tp_btf/sys_enter")
int sys_enter_count(void *ctx) {
	int key = 0;
	struct concurrent_element init_value = {};
	struct concurrent_element *read_value;
	bpf_map_update_elem(&concurrent_map, &key, &init_value, BPF_NOEXIST);

	read_value = bpf_map_lookup_elem(&concurrent_map, &key);
	if(!read_value)
	{
		return 0;
	}

	bpf_spin_lock(&read_value->semaphore);
	read_value->count += 1;
	bpf_spin_unlock(&read_value->semaphore);
	return 0;
}
```

!!! warning
	The verifier will fail if there exists a code path where you take a lock and never release it. You are also not to take more than one lock at a time since that can cause a [deadlock](https://en.wikipedia.org/wiki/Deadlock) scenario.

!!! warning
	Not all BPF program types support `bpf_spin_lock` so be sure to check the [supported program types list](../helper-function/bpf_spin_lock.md#program-types).

On the userspace side we can also request that the spinlock in a value is taken when performing a lookup or update with the [`BPF_F_LOCK`](../syscall/BPF_MAP_LOOKUP_ELEM.md#bpf_f_lock) flag.

## Per CPU maps

Per-CPU maps are map types which have a copy of the map for each logical CPU. By giving each CPU its own memory we side step the issue of synchronizing memory access since there is no shared access. This is the most CPU efficient way to deal with race-conditions for write-heavy tasks. It does however, come at the cost of memory since you need significantly more memory depending on the logical CPU count.

This scheme also increases the complexity on the userspace side since more data needs to be read and the values of the individual CPUs combined.

## Map RCU

In niche use-cases it might be possible to get away with the helper functions built-in RCU logic. This method work by never modifying the map value directly via the pointer you get via the `bpf_map_lookup_elem` helper. But instead copying the map value to the BPF stack, modifying its value there, then calling `bpf_map_update_elem` on the modified copy. The helper functions will guarantee that we transition cleanly from the initial state to the updated state. This property might be important if there exists a relation between fields in the map value. This technique may result in missing updates if multiple updates happen at the same time, but values will never be "mixed".

Performance wise there is a trade off. This technique does perform additional memory copies, but is also does not block or synchronize. So this may or may not be faster than spin-locking depending on the size of the values.

It should be noted that updates via userspace always follow this principle, it is only for BPF programs where this distinction matters.

## Map-in-Map swapping

In most situations it is not possible for userspace to read the contents of a map all at once. Userspace needs to iterate over all keys and perform lookups. This means that during the time it takes to iterate and read, the values in the map can change. This can be problematic for use-case which desire a snapshot of the map at a given time, for statistics for example where the relation between values and time need to be very accurate.

Map-in-maps can be used to get this snapshot behavior. The BPF program first performs a lookup in the outer map which gives the pointer to a inner map. Userspace can swap out the inner map when it wants to collect a snapshot. This is in principle like [multiple buffering](https://en.wikipedia.org/wiki/Multiple_buffering) seen in graphics.
