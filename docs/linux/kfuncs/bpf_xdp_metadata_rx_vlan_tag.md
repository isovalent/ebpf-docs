---
title: "KFunc 'bpf_xdp_metadata_rx_vlan_tag' - eBPF Docs"
description: "This page documents the 'bpf_xdp_metadata_rx_vlan_tag' eBPF kfunc, including its defintion, usage, program types that can use it, and examples."
---
# KFunc `bpf_xdp_metadata_rx_vlan_tag`

<!-- [FEATURE_TAG](bpf_xdp_metadata_rx_vlan_tag) -->
[:octicons-tag-24: v6.8](https://github.com/torvalds/linux/commit/e6795330f88b4f643c649a02662d47b779340535)
<!-- [/FEATURE_TAG] -->

Get XDP packet outermost VLAN tag

## Definition

In case of success, ``vlan_proto`` contains *Tag protocol identifier (TPID)*,
usually ``ETH_P_8021Q`` or ``ETH_P_8021AD``, but some networks can use
custom TPIDs. ``vlan_proto`` is stored in **network byte order (BE)**
and should be used as follows:
``if (vlan_proto == bpf_htons(ETH_P_8021Q)) do_something();``

``vlan_tci`` contains the remaining 16 bits of a VLAN tag.
Driver is expected to provide those in **host byte order (usually LE)**,
so the bpf program should not perform byte conversion.
According to 802.1Q standard, *VLAN TCI (Tag control information)*
is a bit field that contains:
*VLAN identifier (VID)* that can be read with ``vlan_tci & 0xfff``,
*Drop eligible indicator (DEI)* - 1 bit,
*Priority code point (PCP)* - 3 bits.
For detailed meaning of DEI and PCP, please refer to other sources.

`ctx`: XDP context pointer.
`vlan_proto`: Destination pointer for VLAN Tag protocol identifier (TPID).
`vlan_tci`: Destination pointer for VLAN TCI (VID + DEI + PCP)

**Return**
 * Returns 0 on success or ``-errno`` on error.
 * ``-EOPNOTSUPP`` : device driver doesn't implement kfunc
 * ``-ENODATA``    : VLAN tag was not stripped or is not available

<!-- [KFUNC_DEF] -->
`#!c int bpf_xdp_metadata_rx_vlan_tag(const struct xdp_md *ctx, __be16 *vlan_proto, u16 *vlan_tci)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

