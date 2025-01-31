---
title: "KFunc 'bpf_put_file'"
description: "This page documents the 'bpf_put_file' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_put_file`

<!-- [FEATURE_TAG](bpf_put_file) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/d08e2045ebf0f5f2a97ad22cc7dae398b35354ba)
<!-- [/FEATURE_TAG] -->

This function puts a reference on the supplied file.

## Definition

Put a reference on the supplied `file`. Only referenced file pointers may be passed to this BPF kfunc. Attempting to pass an unreferenced file pointer, or any other arbitrary pointer for that matter, will result in the BPF program being rejected by the BPF verifier.

This BPF kfunc may only be called from BPF LSM programs.

**Parameters**

`file`: file to put a reference on

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void bpf_put_file(struct file *file)`

!!! note
	This kfunc releases the pointer passed in to it. There can be only one referenced pointer that can be passed in. 
	All copies of the pointer being released are invalidated as a result of invoking this kfunc.
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

