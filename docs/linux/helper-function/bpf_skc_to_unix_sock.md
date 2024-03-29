---
title: "Helper Function 'bpf_skc_to_unix_sock'"
description: "This page documents the 'bpf_skc_to_unix_sock' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_skc_to_unix_sock`

<!-- [FEATURE_TAG](bpf_skc_to_unix_sock) -->
[:octicons-tag-24: v5.16](https://github.com/torvalds/linux/commit/9eeb3aa33ae005526f672b394c1791578463513f)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Dynamically cast a _sk_ pointer to a _unix_sock_ pointer.

### Returns

_sk_ if casting is valid, or **NULL** otherwise.

`#!c static struct unix_sock *(* const bpf_skc_to_unix_sock)(void *sk) = (void *) 178;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_FLOW_DISSECTOR](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md)
 * [BPF_PROG_TYPE_SK_LOOKUP](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md)
 * [BPF_PROG_TYPE_SK_MSG](../program-type/BPF_PROG_TYPE_SK_MSG.md)
 * [BPF_PROG_TYPE_SK_SKB](../program-type/BPF_PROG_TYPE_SK_SKB.md)
 * [BPF_PROG_TYPE_SOCK_OPS](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
