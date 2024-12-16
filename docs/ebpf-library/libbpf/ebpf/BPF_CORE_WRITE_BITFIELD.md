---
title: "Libbpf eBPF macro 'BPF_CORE_WRITE_BITFIELD'"
description: "This page documents the 'BPF_CORE_WRITE_BITFIELD' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `BPF_CORE_WRITE_BITFIELD`

[:octicons-tag-24: v0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)

The `BPF_CORE_WRITE_BITFIELD` macro writes a bitfield to a given structure in a CO-RE relocatable way.

## Definition

```c
#define BPF_CORE_WRITE_BITFIELD(s, field, new_val) ({			\
	void *p = (void *)s + __CORE_RELO(s, field, BYTE_OFFSET);	\
	unsigned int byte_size = __CORE_RELO(s, field, BYTE_SIZE);	\
	unsigned int lshift = __CORE_RELO(s, field, LSHIFT_U64);	\
	unsigned int rshift = __CORE_RELO(s, field, RSHIFT_U64);	\
	unsigned long long mask, val, nval = new_val;			\
	unsigned int rpad = rshift - lshift;				\
									\
	asm volatile("" : "+r"(p));					\
									\
	switch (byte_size) {						\
	case 1: val = *(unsigned char *)p; break;			\
	case 2: val = *(unsigned short *)p; break;			\
	case 4: val = *(unsigned int *)p; break;			\
	case 8: val = *(unsigned long long *)p; break;			\
	}								\
									\
	mask = (~0ULL << rshift) >> lshift;				\
	val = (val & ~mask) | ((nval << rpad) & mask);			\
									\
	switch (byte_size) {						\
	case 1: *(unsigned char *)p      = val; break;			\
	case 2: *(unsigned short *)p     = val; break;			\
	case 4: *(unsigned int *)p       = val; break;			\
	case 8: *(unsigned long long *)p = val; break;			\
	}								\
})
```

## Usage

`BPF_CORE_WRITE_BITFIELD` writes to a bitfield, identified by `s->field`, the inverse of [`BPF_CORE_READ_BITFIELD`](BPF_CORE_READ_BITFIELD.md) This macro is using direct memory reads and should be used from BPF program types that support such functionality (e.g., typed [raw tracepoints](../../../linux/program-type/BPF_PROG_TYPE_TRACING.md#raw-tracepoint)).

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
