---
title: "KFunc 'bpf_path_d_path'"
description: "This page documents the 'bpf_path_d_path' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_path_d_path`

<!-- [FEATURE_TAG](bpf_path_d_path) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/d08e2045ebf0f5f2a97ad22cc7dae398b35354ba)
<!-- [/FEATURE_TAG] -->

This function resolve the pathname for the supplied path.

## Definition

Resolve the pathname for the supplied `path` and store it in `buf`. This BPF kfunc is the safer variant of the legacy [`bpf_d_path`](../helper-function/bpf_d_path.md) helper and should be used in place of [`bpf_d_path`](../helper-function/bpf_d_path.md) whenever possible. It enforces `KF_TRUSTED_ARGS` semantics, meaning that the supplied `path` must itself hold a valid reference, or else the BPF program will be outright rejected by the BPF verifier.

This BPF kfunc may only be called from BPF LSM programs.

**Parameters**

`path`: path to resolve the pathname for

`buf`: buffer to return the resolved pathname in

`buf__sz`: length of the supplied buffer

**Returns**

A positive integer corresponding to the length of the resolved pathname in *buf*, including the NUL termination character. On error, a negative integer is returned.


**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_path_d_path(struct path *path, char *buf, size_t buf__sz)`
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

