---
title: "KFunc 'bpf_get_dentry_xattr'"
description: "This page documents the 'bpf_get_dentry_xattr' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_get_dentry_xattr`

<!-- [FEATURE_TAG](bpf_get_dentry_xattr) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/ac13a4261afe81ca423eddd8e6571078fe2a7cea)
<!-- [/FEATURE_TAG] -->

This function gets xattr of a dentry.

## Definition

Get xattr `name__str` of `dentry` and store the output in `value_ptr`.

For security reasons, only `name__str` with prefix "user." is allowed.

**Parameters**

`dentry`: dentry to get xattr from

`name__str`: name of the xattr

`value_p`: output buffer of the xattr value

**Returns**

0 on success, a negative value on error.


**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_get_dentry_xattr(struct dentry *dentry, const char *name__str, struct bpf_dynptr *value_p)`

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).
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

