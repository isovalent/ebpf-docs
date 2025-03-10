---
title: "Libbpf userspace function 'bpf_object__set_kversion'"
description: "This page documents the 'bpf_object__set_kversion' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_object__set_kversion`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.4.0](https://github.com/libbpf/libbpf/releases/tag/v0.4.0)
<!-- [/LIBBPF_TAG] -->

Set the kernel version associated with the BPF object.

## Definition

`#!c int bpf_object__set_kversion(struct bpf_object *obj, __u32 kern_version);`

**Parameters**

- `obj`: The BPF object to set the kernel version for.
- `kern_version`: The kernel version to set, encoded as an integer. See [`KERNEL_VERSION`](../ebpf/KERNEL_VERSION.md) details.

**Returns**

- 0 on success, or a negative error code on failure.

## Usage

This function allows you to override the kernel version associated with the BPF object. This version number is used when loading the object into the kernel, see [`kern_version`](../../../linux/syscall/BPF_PROG_LOAD.md#kern_version).

!!! note
    This value was once used for cross-kernel compatibility, but is now unused. [CO-RE](../../../concepts/core.md) has replaced this functionality.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
