---
title: "KFunc 'bpf_session_cookie'"
description: "This page documents the 'bpf_session_cookie' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_session_cookie`

<!-- [FEATURE_TAG](bpf_session_cookie) -->
[:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/5c919acef85147886eb2abf86fb147f94680a8b0)
<!-- [/FEATURE_TAG] -->

Get a pointer to a 64-bit session cookie.

## Definition

**Returns**

Returns pointer to the cookie value. The bpf program can use the pointer to store (on entry) and load (on return) the value.

<!-- [KFUNC_DEF] -->
`#!c __u64 *bpf_session_cookie()`
<!-- [/KFUNC_DEF] -->

## Usage

The session cookie is u64 value, implemented via fprobe feature that allows to share values between entry and return ftrace fprobe callbacks.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_KPROBE](../program-type/BPF_PROG_TYPE_KPROBE.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

