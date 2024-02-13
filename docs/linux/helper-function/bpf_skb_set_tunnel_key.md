# Helper function `bpf_skb_set_tunnel_key`

<!-- [FEATURE_TAG](bpf_skb_set_tunnel_key) -->
[:octicons-tag-24: v4.3](https://github.com/torvalds/linux/commit/d3aa45ce6b94c65b83971257317867db13e5f492)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Populate tunnel metadata for packet associated to _skb._ The tunnel metadata is set to the contents of _key_, of _size_. The _flags_ can be set to a combination of the following values:

**BPF_F_TUNINFO_IPV6**

&nbsp;&nbsp;&nbsp;&nbsp;Indicate that the tunnel is based on IPv6 protocol instead of IPv4.

**BPF_F_ZERO_CSUM_TX**

&nbsp;&nbsp;&nbsp;&nbsp;For IPv4 packets, add a flag to tunnel metadata indicating that checksum computation should be skipped and checksum set to zeroes.

**BPF_F_DONT_FRAGMENT**

&nbsp;&nbsp;&nbsp;&nbsp;Add a flag to tunnel metadata indicating that the packet should not be fragmented.

**BPF_F_SEQ_NUMBER**

&nbsp;&nbsp;&nbsp;&nbsp;Add a flag to tunnel metadata indicating that a sequence number should be added to tunnel header before sending the packet. This flag was added for GRE encapsulation, but might be used with other protocols as well in the future.

**BPF_F_NO_TUNNEL_KEY**

&nbsp;&nbsp;&nbsp;&nbsp;Add a flag to tunnel metadata indicating that no tunnel key should be set in the resulting tunnel header.

Here is a typical usage on the transmit path:

```
struct bpf_tunnel_key key;      populate key ... bpf_skb_set_tunnel_key(skb, &key, sizeof(key), 0); bpf_clone_redirect(skb, vxlan_dev_ifindex, 0);
```

See also the description of the **bpf_skb_get_tunnel_key**() helper for additional information.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (*bpf_skb_set_tunnel_key)(struct __sk_buff *skb, struct bpf_tunnel_key *key, __u32 size, __u64 flags) = (void *) 21;`
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
