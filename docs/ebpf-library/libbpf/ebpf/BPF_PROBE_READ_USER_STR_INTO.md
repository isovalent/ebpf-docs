---
title: "Libbpf eBPF macro 'BPF_PROBE_READ_USER_STR_INTO'"
description: "This page documents the 'BPF_PROBE_READ_USER_STR_INTO' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `BPF_PROBE_READ_USER_STR_INTO`

[:octicons-tag-24: v0.4.0](https://github.com/libbpf/libbpf/releases/tag/v0.4.0)

The `BPF_PROBE_READ_USER_STR_INTO` macro is a variant of the [`BPF_PROBE_READ_USER_INTO`](BPF_PROBE_READ_USER_INTO.md) macro, which is used to do a string read into user-provided storage.

## Definition

```c
#define BPF_PROBE_READ_USER_STR_INTO(dst, src, a, ...) ({		    \
	___core_read(bpf_probe_read_user_str, bpf_probe_read_user,	    \
		     dst, (src), a, ##__VA_ARGS__)			    \
})
```

## Usage

`BPF_PROBE_READ_USER_STR_INTO` is very similar to [`BPF_PROBE_READ_USER_INTO`](BPF_PROBE_READ_USER_INTO.md), but it uses the [`bpf_probe_read_user_str`](../../../linux/helper-function/bpf_probe_read_kernel_str.md) helper function instead of the [`bpf_probe_read_kernel_str`](../../../linux/helper-function/bpf_probe_read_kernel_str.md) helper function. This makes it better suites for reading NUL-terminated strings from userspace memory.

This macro does not emit CO-RE relocations. Its value is mostly in the pointer chasing use cases where this macro can convert one accessor into multiple [`bpf_probe_read_user_str`](../../../linux/helper-function/bpf_probe_read_user_str.md) calls. As no CO-RE relocations are emitted, source types can be arbitrary and are not restricted to kernel types only.

Please refer to the [`BPF_PROBE_READ`](BPF_PROBE_READ.md) documentation for more details on usage of it and its variants like this macros.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
