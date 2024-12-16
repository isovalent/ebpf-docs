---
title: "Libbpf eBPF macro 'BPF_CORE_READ_STR_INTO'"
description: "This page documents the 'BPF_CORE_READ_STR_INTO' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `BPF_CORE_READ_STR_INTO`

[:octicons-tag-24: v0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)

The `BPF_CORE_READ_STR_INTO` macro is a variant of the [`BPF_CORE_READ_INTO`](BPF_CORE_READ_INTO.md) macro, which is used to do a BPF CO-RE relocatable string read into user-provided storage.

## Definition

```c
#define BPF_CORE_READ_STR_INTO(dst, src, a, ...) ({			    \
	___core_read(bpf_core_read_str, bpf_core_read,			    \
		     dst, (src), a, ##__VA_ARGS__)			    \
})
```

## Usage

`BPF_CORE_READ_STR_INTO` is very similar to [`BPF_CORE_READ_INTO`](BPF_CORE_READ_INTO.md), but it uses the [`bpf_probe_read_kernel_str`](../../../linux/helper-function/bpf_probe_read_kernel_str.md) helper function instead of the [`bpf_probe_read_kernel`](../../../linux/helper-function/bpf_probe_read_kernel.md) helper function. This makes it better suites for reading NUL-terminated strings from kernel memory.

Please refer to the [`BPF_CORE_READ`](BPF_CORE_READ.md) documentation for more details on usage of it and its variants like this macros.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
