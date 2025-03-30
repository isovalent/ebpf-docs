---
title: "KFunc 'bpf_xdp_metadata_rx_vlan_tag'"
description: "This page documents the 'bpf_xdp_metadata_rx_vlan_tag' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
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

**Returns**
 * Returns 0 on success or ``-errno`` on error.
 * ``-EOPNOTSUPP`` : device driver doesn't implement kfunc
 * ``-ENODATA``    : VLAN tag was not stripped or is not available

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_xdp_metadata_rx_vlan_tag(const struct xdp_md *ctx, __be16 *vlan_proto, u16 *vlan_tci)`
<!-- [/KFUNC_DEF] -->

## Usage

This kernel function helps to obtain VLAN header information in XDP program context, if VLAN hardware offloads are enabled.
In this case VLAN header is already stripped by the driver code and cannot be seen in XDP linear data.

Care should be taken if the same `xdp_md` buffer is reused to build a "reply" packet. This is often the case when we need for example to construct a TCP SYN-ACK from the intercepted SYN.
When the `ndo_xdp_xmit` driver routine is called to send back the SYN-ACK, even if `tx-vlan-offload` is enabled, the driver will not insert a VLAN header into it.
Thus, if the outgoing packet is built in the XDP context we need to add the VLAN by ourselves. [`bpf_xdp_adjust_tail`](../helper-function/bpf_xdp_adjust_tail.md) call could be used in such a scenario to allocate some extra place for the VLAN header.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

```c
/* XDP program example to show a typical usage of bpf_xdp_metadata_rx_vlan_tag.
 * It could be loaded and attached via following commands:
 * bpftool prog load test_xdp.bpf.o /sys/fs/bpf/prog/xdp_test_vlan type xdp xdpmeta_dev eth1
 * bpftool net attach xdp name xdp_test_vlan dev eth1
 */

#include "vmlinux.h"

#include <bpf/bpf_helpers.h>
#include <bpf/bpf_endian.h>
#include <linux/errno.h>

#include "bpf_tracing_net.h"

#define NO_VLAN_DETECTED 0xffffffff

struct meta_info {
	u16 vlan_proto;
	u16 vlan_tci;
	int rx_vlan_tag_status;
};

extern int bpf_xdp_metadata_rx_vlan_tag(const struct xdp_md *ctx,
					__be16 *vlan_proto,
					__u16 *vlan_tci) __ksym;

SEC("xdp")
int xdp_test_vlan(struct xdp_md *xdp) {

	void *data, *data_end, *data_meta;
	struct ethhdr *eth;
	struct vlan_ethhdr *eth_vlan;
	struct meta_info *meta;
	u32 eth_proto;

	if (bpf_xdp_adjust_meta(xdp, -(int)sizeof(*meta)) < 0) {
		bpf_printk("Cannot allocate %ld bytes for metadata\n",
			   sizeof(*meta));
		return XDP_ABORTED;
	}

	data = (void *)(unsigned long)xdp->data;
	data_meta = (void *)(unsigned long)xdp->data_meta;
	data_end = (void *)(unsigned long)xdp->data_end;
	meta = data_meta;

	/* Check that data_meta have room for meta_info struct */
	if (meta + 1 > data) {
		return XDP_ABORTED;
	}
	meta->rx_vlan_tag_status = NO_VLAN_DETECTED;

	eth = (struct ethhdr *)data;
	if (eth + 1 > data_end)
		return XDP_DROP;

	eth_proto = bpf_ntohs(eth->h_proto);

	if (eth_proto == ETH_P_8021Q) {
		eth_vlan = (struct vlan_ethhdr *)data;
		if (eth_vlan + 1 > data_end)
			return XDP_DROP;

		meta->vlan_proto = eth_proto;
		meta->vlan_tci = bpf_ntohs(eth_vlan->h_vlan_TCI);
		meta->rx_vlan_tag_status = 0;
		bpf_printk("VID=%d from XDP linear data\n",
		           meta->vlan_tci & 0x0fff);
		eth_proto = bpf_ntohs(eth_vlan->h_vlan_encapsulated_proto);

	} else {
		/* check if VLAN was stripped by driver */
		meta->rx_vlan_tag_status = bpf_xdp_metadata_rx_vlan_tag(xdp, &meta->vlan_proto,
									&meta->vlan_tci);
		if (meta->rx_vlan_tag_status == 0)
			bpf_printk("VID=%d reported by driver\n", meta->vlan_tci & 0x0fff);
		if (meta->rx_vlan_tag_status == -ENODATA) {
			bpf_printk("no VLAN reported by driver\n");
			meta->rx_vlan_tag_status = NO_VLAN_DETECTED;
		}

	}

	return XDP_PASS;
}

char _license[] SEC("license") = "GPL";

/* A separate tc filter program, that works in collaboration with the XDP code
 * above and uses VLAN information stored by XDP in the skb->data_meta area.
 * bpftool prog loadall test_tc.bpf.o /sys/fs/bpf/prog
 * tc filter add dev eth1 ingress bpf object-pinned /sys/fs/bpf/prog/tc_ingress_test_vlan
 */

#include "vmlinux.h"

#include <bpf/bpf_helpers.h>
#include <bpf/bpf_endian.h>
#include "bpf_tracing_net.h"

#define NO_VLAN_DETECTED 0xffffffff

struct meta_info {
	u16 vlan_proto;
	u16 vlan_tci;
	int rx_vlan_tag_status;
};

SEC("tc")
int tc_ingress_test_vlan(struct __sk_buff *skb)
{

	void *data = (void *)(unsigned long)skb->data;
	void *data_meta = (void *)(unsigned long)skb->data_meta;
	struct meta_info *meta = data_meta;

	if (meta + 1 > data) {
		bpf_printk("No metadata associate with skb.\n");
		return TC_ACT_OK;
	}

	if (meta->rx_vlan_tag_status == 0)
		bpf_printk("VID=%d\n", meta->vlan_tci & 0x0fff);
	else if (meta->rx_vlan_tag_status == 0xffffffff)
		bpf_printk("No VLAN tag detected\n");
	else
		bpf_printk("VLAN was stripped by driver, "
		           "but bpf_xdp_metadata_rx_vlan_tag returned %d\n",
		           meta->rx_vlan_tag_status);

	return TC_ACT_OK;
}

char _license[] SEC("license") = "GPL";

```

