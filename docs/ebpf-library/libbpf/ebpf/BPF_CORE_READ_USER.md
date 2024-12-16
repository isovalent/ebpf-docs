---
title: "Libbpf eBPF macro 'BPF_CORE_READ_USER'"
description: "This page documents the 'BPF_CORE_READ_USER' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `BPF_CORE_READ_USER`

[:octicons-tag-24: v0.4.0](https://github.com/libbpf/libbpf/releases/tag/v0.4.0)

The `BPF_CORE_READ_USER` macro is the userspace variant of the [`BPF_CORE_READ`](BPF_CORE_READ.md) macro.

## Definition

```c
#define BPF_CORE_READ_USER(src, a, ...) ({				    \
	___type((src), a, ##__VA_ARGS__) __r;				    \
	BPF_CORE_READ_USER_INTO(&__r, (src), a, ##__VA_ARGS__);		    \
	__r;								    \
})
```

## Usage

The `BPF_CORE_READ_USER` macro is the userspace variant of the [`BPF_CORE_READ`](BPF_CORE_READ.md) macro. The difference being that the [`bpf_probe_read_user`](../../../linux/helper-function/bpf_probe_read_user.md) helper function is used instead of the [`bpf_probe_read_kernel`](../../../linux/helper-function/bpf_probe_read_kernel.md) helper function. This makes it able to read from userspace memory.

All the source types involved are still *kernel types* and need to exist in kernel (or kernel module) BTF, otherwise CO-RE relocation will fail. Custom user types are not relocatable with CO-RE. The typical situation in which `BPF_CORE_READ_USER` might be used is to read kernel UAPI types from the user-space memory passed in as a syscall input argument.

Please refer to the [`BPF_CORE_READ`](BPF_CORE_READ.md) documentation for more details on usage of it and its variants like this macros.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
