# Helper function `bpf_msg_redirect_hash`

<!-- [FEATURE_TAG](bpf_msg_redirect_hash) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/81110384441a59cff47430f20f049e69b98c17f4)
<!-- [/FEATURE_TAG] -->

The message redirect hash helper is used to redirect a message to a socket referenced by a hash map.

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


**Returns**
`SK_PASS` on success, or `SK_DROP` on error.

`#!c static long (*bpf_msg_redirect_hash)(struct sk_msg_md *msg, void *map, void *key, __u64 flags) = (void *) 71;`

## Usage

This helper is used in programs implementing policies at the socket level. If the message `msg` is allowed to pass (i.e. if the verdict eBPF program returns `SK_PASS`), redirect it to the socket referenced by `map` (of type [`BPF_MAP_TYPE_SOCKHASH`](../map-type/BPF_MAP_TYPE_SOCKHASH.md)) using hash `key`. Both ingress and egress interfaces can be used for redirection. The `BPF_F_INGRESS` value in `flags` is used to make the distinction (ingress path is selected if the flag is present, egress path otherwise). This is the only flag supported for now.

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_SK_MSG](../program-type/BPF_PROG_TYPE_SK_MSG.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Map types

This helper call can be used with the following map types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_MAP_REF] -->
 * [BPF_MAP_TYPE_SOCKHASH](../map-type/BPF_MAP_TYPE_SOCKHASH.md)
<!-- [/HELPER_FUNC_MAP_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
