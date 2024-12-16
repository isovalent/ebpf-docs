---
title: "Libbpf eBPF macro 'BPF_PROBE_READ_USER'"
description: "This page documents the 'BPF_PROBE_READ_USER' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `BPF_PROBE_READ_USER`

[:octicons-tag-24: v0.4.0](https://github.com/libbpf/libbpf/releases/tag/v0.4.0)

The `BPF_PROBE_READ_USER` macro is the userspace variant of the [`BPF_PROBE_READ`](BPF_PROBE_READ.md) macro.

## Definition

```c
#define BPF_PROBE_READ_USER(src, a, ...) ({				    \
	___type((src), a, ##__VA_ARGS__) __r;				    \
	BPF_PROBE_READ_USER_INTO(&__r, (src), a, ##__VA_ARGS__);	    \
	__r;								    \
})
```

## Usage

The `BPF_PROBE_READ_USER` macro is the userspace variant of the [`BPF_PROBE_READ`](BPF_PROBE_READ.md) macro. The difference being that the [`bpf_probe_read_user`](../../../linux/helper-function/bpf_probe_read_user.md) helper function is used instead of the [`bpf_probe_read_kernel`](../../../linux/helper-function/bpf_probe_read_kernel.md) helper function. This makes it able to read from userspace memory.

This macro does not emit CO-RE relocations. Its value is mostly in the pointer chasing use cases where this macro can convert one accessor into multiple [`bpf_probe_read_user`](../../../linux/helper-function/bpf_probe_read_kernel.md) calls. As no CO-RE relocations are emitted, source types can be arbitrary and are not restricted to kernel types only.

Please refer to the [`BPF_PROBE_READ`](BPF_PROBE_READ.md) documentation for more details on usage of it and its variants like this macros.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
