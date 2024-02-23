---
title: "KFunc 'bpf_get_file_xattr' - eBPF Docs"
description: "This page documents the 'bpf_get_file_xattr' eBPF kfunc, including its defintion, usage, program types that can use it, and examples."
---
# KFunc `bpf_get_file_xattr`

<!-- [FEATURE_TAG](bpf_get_file_xattr) -->
[:octicons-tag-24: v6.8](https://github.com/torvalds/linux/commit/ac9c05e0e453cfcab2866f6d28f257590e4f66e5)
<!-- [/FEATURE_TAG] -->

Get xattr of a file

## Definition

Get xattr `name__str` of `file` and store the output in `value_ptr`.

For security reasons, only `name__str` with prefix "user." is allowed.

**Return**

0 on success, a negative value on error.

<!-- [KFUNC_DEF] -->
`#!c int bpf_get_file_xattr(struct file *file, const char *name__str, struct bpf_dynptr_kern *value_ptr)`

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../../syscall/BPF_PROG_LOAD/#bpf_f_sleepable).
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
- [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

