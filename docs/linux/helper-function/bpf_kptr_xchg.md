---
title: "Helper Function 'bpf_kptr_xchg' - eBPF Docs"
description: "This page documents the 'bpf_kptr_xchg' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_kptr_xchg`

<!-- [FEATURE_TAG](bpf_kptr_xchg) -->
[:octicons-tag-24: v5.19](https://github.com/torvalds/linux/commit/c0a5a21c25f37c9fd7b36072f9968cdff1e4aa13)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Exchange kptr at pointer _map_value_ with _ptr_, and return the old value. _ptr_ can be NULL, otherwise it must be a referenced pointer which will be released when this helper is called.

### Returns

The old value of kptr (which can be NULL). The returned pointer if not NULL, is a reference which must be released using its corresponding release function, or moved into a BPF map before program exit.

`#!c static void *(*bpf_kptr_xchg)(void *map_value, void *ptr) = (void *) 194;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_SYSCTL](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
