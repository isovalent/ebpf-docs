---
title: "SCX eBPF macro 'READ_ONCE'"
description: "This page documents the 'READ_ONCE' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `READ_ONCE`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/2a52ca7c98960aafb0eca9ef96b2d0c932171357)

The `READ_ONCE` macro prevents the compiler caching, redoing or reordering reads.

## Definition

```c
typedef __u8  __attribute__((__may_alias__))  __u8_alias_t;
typedef __u16 __attribute__((__may_alias__)) __u16_alias_t;
typedef __u32 __attribute__((__may_alias__)) __u32_alias_t;
typedef __u64 __attribute__((__may_alias__)) __u64_alias_t;

static __always_inline void __read_once_size(const volatile void *p, void *res, int size)
{
	switch (size) {
	case 1: *(__u8_alias_t  *) res = *(volatile __u8_alias_t  *) p; break;
	case 2: *(__u16_alias_t *) res = *(volatile __u16_alias_t *) p; break;
	case 4: *(__u32_alias_t *) res = *(volatile __u32_alias_t *) p; break;
	case 8: *(__u64_alias_t *) res = *(volatile __u64_alias_t *) p; break;
	default:
		[barrier](../libbpf/ebpf/barrier.md)();
		__builtin_memcpy((void *)res, (const void *)p, size);
		[barrier](../libbpf/ebpf/barrier.md)();
	}
}

#define READ_ONCE(x)                                \
({                                                  \
	union { typeof(x) __val; char __c[1]; } __u =   \
		{ .__c = { 0 } };                           \
	__read_once_size(&(x), __u.__c, sizeof(x));     \
	__u.__val;                                      \
})
```

## Usage

Compilers will try to optimize code in any way possible within the constraints they are aware of. They may for example:

* Keep the result of a memory read in a register and re-use it instead of re-reading the memory location. (caching)
* Reorder reads to improve cache locality (reordering)
* Re-read from memory instead of keeping the value in a register (redoing)

The `READ_ONCE` macro makes sure the compiler can't do any of these optimizations. This is useful when you want to make sure you read a value from memory exactly once, and at the exact point in the code where you expect it to be read. For example when reading a value from a shared data structure that may be modified by another thread.

### Example

In this example, we want to force the compiler to read the value of `shared_data->flag` from memory every time we check it in the loop. Without `READ_ONCE`, the compiler might pull the read out of the loop and cache the value (because by default the compiler will not be aware the memory can be changed by something else).

```c
for (int i = 0; i < 10; i++) {
    if (READ_ONCE(shared_data->flag) == 1) {
        // Do something
    }
}
```
