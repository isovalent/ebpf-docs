---
title: "Libbpf userspace function 'bpf_object__kversion'"
description: "This page documents the 'bpf_object__kversion' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_object__kversion`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Get the kernel version associated with the BPF object.

## Definition

`#!c unsigned int bpf_object__kversion(const struct bpf_object *obj);`

**Returns**

The kernel version, encoded as an integer. See [`KERNEL_VERSION`](../ebpf/KERNEL_VERSION.md) details.

## Usage

This version number is gotten from the `version` ELF section and indicates the kernel version the object was built against. If the ELF file does not contain a version section, the kernel version is probed at load time.

This version number is passed to the kernel when loading the program, see [`kern_version`](../../../linux/syscall/BPF_PROG_LOAD.md#kern_version).

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
