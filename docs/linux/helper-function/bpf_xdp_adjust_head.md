---
title: "Helper Function 'bpf_xdp_adjust_head'"
description: "This page documents the 'bpf_xdp_adjust_head' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_xdp_adjust_head`

<!-- [FEATURE_TAG](bpf_xdp_adjust_head) -->
[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/17bedab2723145d17b14084430743549e6943d03)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Adjust (move) _xdp_md_**->data** by _delta_ bytes. Note that it is possible to use a negative value for _delta_. This helper can be used to prepare the packet for pushing or popping headers.

A call to this helper is susceptible to change the underlying packet buffer. Therefore, at load time, all checks on pointers previously done by the verifier are invalidated and must be performed again, if the helper is used in combination with direct packet access.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_xdp_adjust_head)(struct xdp_md *xdp_md, int delta) = (void *) 44;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

This example, inspired by the Linux kernel's [`tools/testing/selftests/bpf/progs/test_xdp_vlan.c`][test_exdp_vlan.c]
demonstrates how to remove an outer `802.1q` VLAN tag (_4 bytes_) from Ethernet packets.
The program shifts the Ethernet header to overwrite the VLAN tag and adjusts the packet size using `bpf_xdp_adjust_head`.

[test_exdp_vlan.c]: https://elixir.bootlin.com/linux/v6.15/source/tools/testing/selftests/bpf/progs/test_xdp_vlan.c#L182

```c
#include <linux/bpf.h>
#include <linux/if_ether.h>
#include <bpf/bpf_helpers.h>

#define VLAN_HDR_SZ (4)

SEC("xdp")
int xdp_remove_outer_vlan_tag(struct xdp_md *ctx)
{
    void *data = (void *)(long)ctx->data;
    void *data_end = (void *)(long)ctx->data_end;
    struct ethhdr *eth = data;

    /* Ensure there's enough room for (ethernet + vlan) headers */
    int hdrsize = (sizeof(struct ethhdr) + VLAN_HDR_SZ);
    if ((data + hdrsize > data_end) || (eth->h_proto != bpf_htons(ETH_P_8021Q))) {
        return XDP_PASS;
    }

     /*
     * To remove the VLAN header, shift the source and
     * destination MAC addresses (12 bytes) right by VLAN_HDR_SZ (4 bytes),
     * overwriting the VLAN's TPID and TCI fields.
     * The encapsulated protocol remains unchanged.
     *
     * Since *dest* and *data* overlap, *__builtin_memmove*
     * ensures safe copying. Alternatively, you could save the MAC
     * addresses and restore them after shrinking the packet.
     */
    char *dest = (data + VLAN_HDR_SZ);
    __builtin_memmove(dest, data, ETH_ALEN * 2);

    /* Shrink packet by VLAN header size */
    bpf_xdp_adjust_head(ctx, VLAN_HDR_SZ);

    /*
     * Note: After bpf_xdp_adjust_head, the eBPF verifier invalidates all
     * previous pointer checks. Any subsequent pointer accesses (e.g., to data
     * or eth) must be revalidated to ensure safety.
     */

    return XDP_PASS;
}
```
