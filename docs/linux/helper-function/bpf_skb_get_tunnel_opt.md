# Helper function `bpf_skb_get_tunnel_opt`

<!-- [FEATURE_TAG](bpf_skb_get_tunnel_opt) -->
[:octicons-tag-24: v4.6](https://github.com/torvalds/linux/commit/14ca0751c96f8d3d0f52e8ed3b3236f8b34d3460)
<!-- [/FEATURE_TAG] -->

## Definition

<!-- [HELPER_FUNC_DEF] -->
Retrieve tunnel options metadata for the packet associated to _skb_, and store the raw tunnel option data to the buffer _opt_ of _size_.

This helper can be used with encapsulation devices that can operate in "collect metadata" mode (please refer to the related note in the description of **bpf_skb_get_tunnel_key**() for more details). A particular example where this can be used is in combination with the Geneve encapsulation protocol, where it allows for pushing (with **bpf_skb_get_tunnel_opt**() helper) and retrieving arbitrary TLVs (Type-Length-Value headers) from the eBPF program. This allows for full customization of these headers.

### Returns

The size of the option data retrieved.

`#!c static long (*bpf_skb_get_tunnel_opt)(struct __sk_buff *skb, void *opt, __u32 size) = (void *) 29;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
