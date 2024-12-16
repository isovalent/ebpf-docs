---
title: "Libbpf eBPF macro 'BPF_CORE_READ_BITFIELD_PROBED'"
description: "This page documents the 'BPF_CORE_READ_BITFIELD_PROBED' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `BPF_CORE_READ_BITFIELD_PROBED`

[:octicons-tag-24: v0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)

The `BPF_CORE_READ_BITFIELD_PROBED` macro extracts a bitfield from a given structure in a CO-RE relocatable way.

## Definition

```c
#define BPF_CORE_READ_BITFIELD_PROBED(s, field) ({			      \
	unsigned long long val = 0;					      \
									      \
	__CORE_BITFIELD_PROBE_READ(&val, s, field);			      \
	val <<= __CORE_RELO(s, field, LSHIFT_U64);			      \
	if (__CORE_RELO(s, field, SIGNED))				      \
		val = ((long long)val) >> __CORE_RELO(s, field, RSHIFT_U64);  \
	else								      \
		val = val >> __CORE_RELO(s, field, RSHIFT_U64);		      \
	val;								      \
})
```

## Usage

`BPF_CORE_READ_BITFIELD` extract bitfield, identified by `s->field`, and return its value as u64. All this is done in relocatable manner, so bitfield changes such as signedness, bit size, offset changes, this will be handled automatically. This version of macro is using [`bpf_probe_read_kernel`](../../../linux/helper-function/bpf_probe_read_kernel.md) to read underlying integer storage. Macro functions as an expression and its return type is [`bpf_probe_read_kernel`](../../../linux/helper-function/bpf_probe_read_kernel.md) return value: 0, on success, <0 on error.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
