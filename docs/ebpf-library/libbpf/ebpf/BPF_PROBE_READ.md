---
title: "Libbpf eBPF macro 'BPF_PROBE_READ'"
description: "This page documents the 'BPF_PROBE_READ' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `BPF_PROBE_READ`

[:octicons-tag-24: v0.4.0](https://github.com/libbpf/libbpf/releases/tag/v0.4.0)

The `BPF_PROBE_READ` macro is the non CO-RE variant of [`BPF_CORE_READ`](BPF_CORE_READ.md).

## Definition

```c
#define BPF_PROBE_READ(src, a, ...) ({					    \
	___type((src), a, ##__VA_ARGS__) __r;				    \
	BPF_PROBE_READ_INTO(&__r, (src), a, ##__VA_ARGS__);		    \
	__r;								    \
})
```

## Usage

The `BPF_PROBE_READ` macro is the non CO-RE variant of [`BPF_CORE_READ`](BPF_CORE_READ.md). So it does not emit CO-RE relocations. Its value is mostly in the pointer chasing use cases where this macro can convert one accessor into multiple [`bpf_probe_read_kernel`](../../../linux/helper-function/bpf_probe_read_kernel.md) calls.

!!! note
    Only up to 9 "field accessors" are supported, which should be more than enough for any practical purpose.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
