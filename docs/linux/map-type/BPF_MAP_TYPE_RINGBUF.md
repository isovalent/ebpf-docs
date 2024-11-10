---
title: "Map Type 'BPF_MAP_TYPE_RINGBUF'"
description: "This page documents the 'BPF_MAP_TYPE_RINGBUF' eBPF map type, including its definition, usage, program types that can use it, and examples."
---
# Map type `BPF_MAP_TYPE_RINGBUF`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_RINGBUF) -->
[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/457f44363a8894135c85b7a9afd2bd8196db24ab)
<!-- [/FEATURE_TAG] -->

The ring-buffer map can be used to efficiently send large amounts of data from eBPF programs to userspace. Data is sent in a queue / first-in-first-out (FIFO) manner.

This map consists of a singular ring as opposed to the per-CPU design of the `BPF_MAP_TYPE_PERF_ARRAY` map type. This means that the order of events is preserved across all CPUs.

## Attributes

Since this map type does not have key-value pairs, and the communicated samples can be of any size, the [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) and [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) attributes have to both be set to `0`.

The [`max_entries`](../syscall/BPF_MAP_CREATE.md#max_entries) attribute is used to specify the size of the ring-buffer in bytes. It must be a power of 2 and a multiple of the page size (typically `4096`), so `4096`, `8192`, `16384`, `32768`, ect.

## Userspace map reading

!!! note
    eBPF loader libraries such as libbpf abstract the following steps away from the user. Unless you interested in the actual internals, you can skip this section.

Unlike most maps, the map lookup, update, and delete BPF syscall commands can not be used with this map type. To read from the map, a userspace program must memory map part of the ring-buffer into its address space. The `mmap` syscall is used to achieve this.

First the userspace program must map the "consumer" page

```c
consumer = mmap(NULL, rb->page_size, PROT_READ | PROT_WRITE, MAP_SHARED, map_fd, 0);
```

The consumer page is writable. While a whole page was mapped, only the first 8 bytes are used to represent a 64-bit unsigned integer, the index of the consumer.

The second area to map is the "producer".

```c
mmap_sz = rb->page_size + 2 * (__u64)info.max_entries;
producer = mmap(NULL, (size_t)mmap_sz, PROT_READ, MAP_SHARED, map_fd, rb->page_size);
```

The producer memory area is read-only. We map an area twice the size of the actual ring-buffer size plus 1 page, this is an optimization. The single, additional page is used to map the producer index, the same way we did for the consumer. The pages after that are the actual data. Since the buffer is a circular buffer, data can be split between the end of the buffer and the beginning. By mapping the physical memory twice into virtual memory we can read any overflowing data as if it was contiguous.

<div class="mono">

```
Single mapped
0          4096                           8192
+----------+---------------------------------+
+ Prod idx |...one sample|      | This is ...|
+----------+---------------------------------+


Double mapped
0          4096                                                     12288
+----------+------------------------------------------------------------+
+ Prod idx |...one sample|      | This is one sample|      | This is ...|
+----------+------------------------------------------------------------+
```

</div>

Both the consumer and producer are indexes into the buffer. When eBPF programs write to the buffer, they increment the producer index. The data between the consumer and producer indexes is the data that has not been read yet. 

!!! warning
    Reads from and writes to both indexes should use atomic operations to avoid race conditions, reading the actual data can be done without atomic operations.

Every sample has 8 byte header, containing 2 32-bit integers, the length of the sample and the page offset. The page offset is used by the kernel but not relevant.

<div class="mono">

```
+-------------+-----------------+
| len | pgoff | sample contents |
+-------------+-----------------+
```

</div>

The upper two bits of the length are used to store flags. The lower 30 bits are the actual length of the sample.

```
BPF_RINGBUF_BUSY_BIT		= (1U << 31),
BPF_RINGBUF_DISCARD_BIT		= (1U << 30),
```

So userspace first read the header. Check the busy bit, if it is set, room for the sample has been allocated by a program is still writing to it, userspace should wait until the busy bit is cleared. If the discard bit is set, the space was initially reserved but the program decided not to write to it. In this case, userspace can skip the sample.

The samples are always 8-byte aligned. So if a sample is 5 bytes long, 3 bytes of padding will be added to the end of the sample. For example:

<div class="mono">

```
<-----------Sample 1-------------><----------Sample 2--------->

0         8                 13    16        24               32
+---------+-----------------+-----+---------+-----------------+
| 5 | ... | sample contents | pad | 8 | ... | sample contents |
+---------+-----------------+-----+---+-----+-----------------+
```

</div>

So the consumer index should always be incremented by the length of the sample plus the padding (rounded up to multiple of 8).

### epoll

A userspace program can use the `epoll_*` syscalls to wait for new data to be available in the ring-buffer. This is often preferred since no CPU is consumed while waiting for new data, the OS scheduler will wake up the program when new data is available.

The eBPF program has some control over when the userspace program is woken up. When using the [`bpf_ringbuf_output`](../helper-function/bpf_ringbuf_output.md), [`bpf_ringbuf_submit`](../helper-function/bpf_ringbuf_submit.md), or [`bpf_ringbuf_submit_dynptr`](../helper-function/bpf_ringbuf_submit_dynptr.md) helpers the `BPF_RB_NO_WAKEUP` or `BPF_RB_FORCE_WAKEUP` flag can be set. When the `BPF_RB_NO_WAKEUP` flag is set, no notification of new data availability is sent. When the `BPF_RB_FORCE_WAKEUP` flag is set a notification of new data availability is sent unconditionally. If no flags are set the notification is sent "adaptively". This means a notification is sent whenever the userspace process has caught up and consumed all available payloads. If the userspace process is still processing a previous payload, then no notification is needed as it will process the newly added payload automatically.

## eBPF map writing

The ring-buffer map has its own unique collection of helper functions to write to it. There are effectively three modes of writing to the ring-buffer.

### Ring-buffer output

The first is to use [`bpf_ringbuf_output`](../helper-function/bpf_ringbuf_output.md). This helper function copies data from the eBPF program into the ring-buffer in one go, so you must already have the full sample data ready somewhere in memory.

### reserve, submit, and discard

The second is to use [`bpf_ringbuf_reserve`](../helper-function/bpf_ringbuf_reserve.md) to reserve space in the ring-buffer. You give the helper a size in bytes and it returns a pointer of that size in the ring-buffer (or a NULL pointer if allocation was not possible).
You can then write to this memory location from the eBPF program.

This has two advantages:
    * No additional memory copy is needed, the data is written directly to the ring-buffer, thus faster.
    * The reserved memory does not count towards the program stack, so you can write large samples without worrying about stack space.

The verifier has to assert that you do no overrun the reserved memory. If you do, the program will be rejected. Side effect of this is that the `size` provided to `bpf_ringbuf_reserve` must be known at compile time. (see the last method for a way around this).

Once you are done writing, you use [`bpf_ringbuf_submit`](../helper-function/bpf_ringbuf_submit.md) to make the data available to userspace. If you decide after reserving that you do not want to submit the sample you can use [`bpf_ringbuf_discard`](../helper-function/bpf_ringbuf_discard.md) to discard the reservation. 

!!! example
    ```c
    // Reserve space in the ring buffer
    struct ringbuf_data *rb_data = bpf_ringbuf_reserve(&my_ringbuf, sizeof(struct ringbuf_data), 0);
    if(!rb_data) {
        // if bpf_ringbuf_reserve fails, print an error message and return
        bpf_printk("bpf_ringbuf_reserve failed\n");
        return 1;
    }

    if(unhappy_flow) {
        // Discard the reserved data
        bpf_ringbuf_discard(rb_data, 0);
        return 1;
    }

    // Submit the reserved data
    bpf_ringbuf_submit(rb_data, 0);
    ```

### dynamic pointers

In this final method we utilize the [`bpf_ringbuf_reserve_dynptr`](../helper-function/bpf_ringbuf_reserve_dynptr.md) helper function to reserve space in the ring-buffer instead of `bpf_ringbuf_reserve`. This variant also takes a sample size but return a "dynamic pointer", a pointer with additional metadata, in this case about the size of the sample. This allows verification of the sample size at runtime and thus allows for dynamic sample sizes to be provided to helper.

You cannot read or write to this pointer directly, you must use the [`bpf_dynptr_read`](../helper-function/bpf_dynptr_read.md), [`bpf_dynptr_write`](../helper-function/bpf_dynptr_write.md), [`bpf_dynptr_data`](../helper-function/bpf_dynptr_data.md), or [`bpf_dynptr_from_mem`](../helper-function/bpf_dynptr_from_mem.md) helper functions to interact with the data which can fail at runtime if they cause accesses outside the bounds of the sample.

!!! note
    There are also a growing number of kfuncs which can be used to interact with dynamic pointers

Once you are done writing, you use [`bpf_ringbuf_submit_dynptr`](../helper-function/bpf_ringbuf_submit_dynptr.md) to make the data available to userspace. If you decide after reserving that you do not want to submit the sample you can use [`bpf_ringbuf_discard_dynptr`](../helper-function/bpf_ringbuf_discard_dynptr.md) to discard the reservation.

!!! example
    ```c
    // SPDX-License-Identifier: GPL-2.0
    /* Copyright (c) 2022 Facebook */
    
    SEC("?tp/syscalls/sys_enter_nanosleep")
    int test_read_write(void *ctx)
    {
        char write_data[64] = "hello there, world!!";
        char read_data[64] = {};
        struct bpf_dynptr ptr;
        int i;

        if (bpf_get_current_pid_tgid() >> 32 != pid)
            return 0;

        bpf_ringbuf_reserve_dynptr(&ringbuf, sizeof(write_data), 0, &ptr);

        /* Write data into the dynptr */
        err = bpf_dynptr_write(&ptr, 0, write_data, sizeof(write_data), 0);

        /* Read the data that was written into the dynptr */
        err = err ?: bpf_dynptr_read(read_data, sizeof(read_data), &ptr, 0, 0);

        /* Ensure the data we read matches the data we wrote */
        for (i = 0; i < sizeof(read_data); i++) {
            if (read_data[i] != write_data[i]) {
                err = 1;
                break;
            }
        }

        bpf_ringbuf_submit_dynptr(&ptr, 0);
        return 0;
    }
    ```

## Helper functions

The following helper functions work with this map type:

<!-- DO NOT EDIT MANUALLY -->
<!-- [MAP_HELPER_FUNC_REF] -->
 * [`bpf_ringbuf_discard`](../helper-function/bpf_ringbuf_discard.md)
 * [`bpf_ringbuf_discard_dynptr`](../helper-function/bpf_ringbuf_discard_dynptr.md)
 * [`bpf_ringbuf_output`](../helper-function/bpf_ringbuf_output.md)
 * [`bpf_ringbuf_query`](../helper-function/bpf_ringbuf_query.md)
 * [`bpf_ringbuf_reserve`](../helper-function/bpf_ringbuf_reserve.md)
 * [`bpf_ringbuf_reserve_dynptr`](../helper-function/bpf_ringbuf_reserve_dynptr.md)
 * [`bpf_ringbuf_submit`](../helper-function/bpf_ringbuf_submit.md)
 * [`bpf_ringbuf_submit_dynptr`](../helper-function/bpf_ringbuf_submit_dynptr.md)
<!-- [/MAP_HELPER_FUNC_REF] -->

## Flags

### `BPF_F_NUMA_NODE`

[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/96eabe7a40aa17e613cf3db2c742ee8b1fc764d0)

When set, the [`numa_node`](../syscall/BPF_MAP_CREATE.md#numa_node) attribute is respected during map creation.
