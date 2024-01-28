# Helper function `bpf_skc_to_tcp_timewait_sock`

<!-- [FEATURE_TAG](bpf_skc_to_tcp_timewait_sock) -->
[:octicons-tag-24: v5.9](https://github.com/torvalds/linux/commit/478cfbdf5f13dfe09cfd0b1cbac821f5e27f6108)
<!-- [/FEATURE_TAG] -->

## Defintion

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Dynamically cast a _sk_ pointer to a _tcp_timewait_sock_ pointer.

### Returns

_sk_ if casting is valid, or **NULL** otherwise.

`#!c static struct tcp_timewait_sock *(*bpf_skc_to_tcp_timewait_sock)(void *sk) = (void *) 138;`
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
