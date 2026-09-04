---
title: "KFunc 'bpf_set_dentry_xattr'"
description: "This page documents the 'bpf_set_dentry_xattr' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_set_dentry_xattr`

<!-- [FEATURE_TAG](bpf_set_dentry_xattr) -->
[:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/56467292794b800164df20c076c409ac548e56ec)
<!-- [/FEATURE_TAG] -->

Set a extended attribute(<nospell>xattr</nospell>) of a directory entry(<nospell>dentry</nospell>).

## Definition

Set <nospell>xattr</nospell> `name__str` of `dentry` to the value in `value_ptr`. For security reasons, only `name__str` with prefix `security.bpf.` is allowed. The caller has not locked `dentry->d_inode`.

**Parameters**

`dentry`: dentry to get <nospell>xattr</nospell> from

`name__str`: name of the <nospell>xattr</nospell>

`value_p`: <nospell>xattr</nospell> value

`flags`: flags to pass into filesystem operations


**Returns**

`0` on success, a negative value on error.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_set_dentry_xattr(struct dentry *dentry, const char *name__str, const struct bpf_dynptr *value_p, int flags)`

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
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

