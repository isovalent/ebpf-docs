---
title: "KFunc 'bpf_key_put'"
description: "This page documents the 'bpf_key_put' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_key_put`

<!-- [FEATURE_TAG](bpf_key_put) -->
[:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/f3cf4134c5c6c47b9b5c7aa3cb2d67e107887a7b)
<!-- [/FEATURE_TAG] -->

Decrement key reference count if key is valid and free bpf_key

## Definition

Decrement the reference count of the key inside `bkey`, if the pointer is valid, and free `bkey`.

<!-- [KFUNC_DEF] -->
`#!c void bpf_key_put(struct bpf_key *bkey)`

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
- [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

