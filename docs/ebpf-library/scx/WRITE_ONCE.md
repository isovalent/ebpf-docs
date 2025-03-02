---
title: "SCX eBPF macro 'WRITE_ONCE'"
description: "This page documents the 'WRITE_ONCE' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `WRITE_ONCE`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/2a52ca7c98960aafb0eca9ef96b2d0c932171357)

The `WRITE_ONCE` macro prevents the compiler caching, redoing or reordering writes.

## Definition

```c
typedef __u8  __attribute__((__may_alias__))  __u8_alias_t;
typedef __u16 __attribute__((__may_alias__)) __u16_alias_t;
typedef __u32 __attribute__((__may_alias__)) __u32_alias_t;
typedef __u64 __attribute__((__may_alias__)) __u64_alias_t;

static __always_inline void __write_once_size(volatile void *p, void *res, int size)
{
	switch (size) {
	case 1: *(volatile  __u8_alias_t *) p = *(__u8_alias_t  *) res; break;
	case 2: *(volatile __u16_alias_t *) p = *(__u16_alias_t *) res; break;
	case 4: *(volatile __u32_alias_t *) p = *(__u32_alias_t *) res; break;
	case 8: *(volatile __u64_alias_t *) p = *(__u64_alias_t *) res; break;
	default:
		[barrier](../libbpf/ebpf/barrier.md)();
		__builtin_memcpy((void *)p, (const void *)res, size);
		[barrier](../libbpf/ebpf/barrier.md)();
	}
}

#define WRITE_ONCE(x, val)                          \
({                                                  \
	union { typeof(x) __val; char __c[1]; } __u =   \
		{ .__val = (val) };                         \
	__write_once_size(&(x), __u.__c, sizeof(x));    \
	__u.__val;                                      \
})
```

## Usage

Compilers will try to optimize code in any way possible within the constraints they are aware of. They may for example:

* Remove duplicate writes to the same memory location.
* Reorder writes to improve cache locality.

The `WRITE_ONCE` macro makes sure the compiler can't do any of these optimizations. This is useful when you want to make sure you write a value from memory exactly once, and at the exact point in the code where you expect it to be read. For example when writing a value from a shared data structure that may be modified by another thread.

### Example

Lets say that its crucial that we write to memory in this exact order. Without `WRITE_ONCE`, the compiler might reorder the writes, so the two writes to `shared_data` happen before the write to `some_other_var`, to improve cache locality.

```c
WRITE_ONCE(shared_data->field1, 1);
WRITE_ONCE(some_other_var, true);
WRITE_ONCE(shared_data->field2, 2);
```
