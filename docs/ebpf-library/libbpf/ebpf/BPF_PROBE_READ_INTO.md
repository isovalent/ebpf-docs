---
title: "Libbpf eBPF macro 'BPF_PROBE_READ_INTO'"
description: "This page documents the 'BPF_PROBE_READ_INTO' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `BPF_PROBE_READ_INTO`

[:octicons-tag-24: v0.4.0](https://github.com/libbpf/libbpf/releases/tag/v0.4.0)

The `BPF_PROBE_READ_INTO` macro is a more performance-conscious variant of [`BPF_PROBE_READ`](BPF_PROBE_READ.md), in which final field is read into user-provided storage.

## Definition

```c
#define BPF_PROBE_READ_INTO(dst, src, a, ...) ({			    \
	___core_read(bpf_probe_read_kernel, bpf_probe_read_kernel,	    \
		     dst, (src), a, ##__VA_ARGS__)			    \
})
```

## Usage

`BPF_PROBE_READ_INTO` is very similar to [`BPF_PROBE_READ`](BPF_PROBE_READ.md), but instead of returning the value, it writes the value into the provided destination.

This macro does not emit CO-RE relocations. Its value is mostly in the pointer chasing use cases where this macro can convert one accessor into multiple [`bpf_probe_read_kernel`](../../../linux/helper-function/bpf_probe_read_kernel.md) calls.

The following two code snippets are equivalent:

```c
int x = BPF_PROBE_READ_INTO(s, a.b.c, d.e, f, g);
```

```c
int x;
BPF_PROBE_READ(&x, s, a.b.c, d.e, f, g);
```

Please refer to the [`BPF_PROBE_READ`](BPF_PROBE_READ.md) documentation for more details on usage of it and its variants like this macros.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
