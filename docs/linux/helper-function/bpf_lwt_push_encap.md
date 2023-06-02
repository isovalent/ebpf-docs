# Helper function `bpf_lwt_push_encap`

<!-- [FEATURE_TAG](bpf_lwt_push_encap) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/fe94cc290f535709d3c5ebd1e472dfd0aec7ee79)
<!-- [/FEATURE_TAG] -->

## Definition

<!-- [HELPER_FUNC_DEF] -->
Encapsulate the packet associated to _skb_ within a Layer 3 protocol header. This header is provided in the buffer at address _hdr_, with _len_ its size in bytes. _type_ indicates the protocol of the header and can be one of:

**BPF_LWT_ENCAP_SEG6**

&nbsp;&nbsp;&nbsp;&nbsp;IPv6 encapsulation with Segment Routing Header (**struct ipv6_sr_hdr**). _hdr_ only contains the SRH, the IPv6 header is computed by the kernel.

**BPF_LWT_ENCAP_SEG6_INLINE**

&nbsp;&nbsp;&nbsp;&nbsp;Only works if _skb_ contains an IPv6 packet. Insert a Segment Routing Header (**struct ipv6_sr_hdr**) inside the IPv6 header.

**BPF_LWT_ENCAP_IP**

&nbsp;&nbsp;&nbsp;&nbsp;IP encapsulation (GRE/GUE/IPIP/etc). The outer header must be IPv4 or IPv6, followed by zero or more additional headers, up to **LWT_BPF_MAX_HEADROOM** total bytes in all prepended headers. Please note that if **skb_is_gso**(_skb_) is true, no more than two headers can be prepended, and the inner header, if present, should be either GRE or UDP/GUE.

**BPF_LWT_ENCAP_SEG6**\_ types can be called by BPF programs of type **BPF_PROG_TYPE_LWT_IN**; **BPF_LWT_ENCAP_IP** type can be called by bpf programs of types **BPF_PROG_TYPE_LWT_IN** and **BPF_PROG_TYPE_LWT_XMIT**.

A call to this helper is susceptible to change the underlying packet buffer. Therefore, at load time, all checks on pointers previously done by the verifier are invalidated and must be performed again, if the helper is used in combination with direct packet access.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (*bpf_lwt_push_encap)(struct __sk_buff *skb, __u32 type, void *hdr, __u32 len) = (void *) 73;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
