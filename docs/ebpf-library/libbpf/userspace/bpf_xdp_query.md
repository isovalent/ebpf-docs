---
title: "Libbpf userspace function 'bpf_xdp_query'"
description: "This page documents the 'bpf_xdp_query' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_xdp_query`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.7.0](https://github.com/libbpf/libbpf/releases/tag/v0.7.0)
<!-- [/LIBBPF_TAG] -->

Query a network interface to get information about the XDP program attached to it.

## Definition

`#!c int bpf_xdp_query(int ifindex, int flags, struct bpf_xdp_query_opts *opts);`

**Parameters**

- `ifindex`: The index of the network interface to query.
- `flags`: Flags to control the query.
- `opts`: A pointer to a `struct bpf_xdp_query_opts` structure that will be filled with the query results.

**Flags**

* `XDP_FLAGS_SKB_MODE` = `(1U << 1)` - If set, query for programs in SKB (generic) mode.
* `XDP_FLAGS_DRV_MODE` = `(1U << 2)` - If set, query for programs in DRV (driver / native) mode.
* `XDP_FLAGS_HW_MODE` = `(1U << 3)` - If set, query for programs in hardware offload mode.

**Return**

`0` on success. A negative error code on failure.

### `struct bpf_xdp_query_opts`

```c
struct bpf_xdp_query_opts {
	size_t sz;
	__u32 prog_id;		/* output */
	__u32 drv_prog_id;	/* output */
	__u32 hw_prog_id;	/* output */
	__u32 skb_prog_id;	/* output */
	__u8 attach_mode;	/* output */
	__u64 feature_flags;	/* output */
	__u32 xdp_zc_max_segs;	/* output */
	size_t :0;
};
```

#### `prog_id`

[:octicons-tag-24: 0.7.0](https://github.com/libbpf/libbpf/commit/8fbe7eec3aacca51d81785f95da295d40e1cb965)

The ID of the program that is attached (if attached via netlink, programs attached via BPF link will not be returned). Regardless of the mode the program is attached in.

#### `drv_prog_id`

[:octicons-tag-24: 0.7.0](https://github.com/libbpf/libbpf/commit/8fbe7eec3aacca51d81785f95da295d40e1cb965)

The ID of the program if attached in driver mode.

#### `hw_prog_id`

[:octicons-tag-24: 0.7.0](https://github.com/libbpf/libbpf/commit/8fbe7eec3aacca51d81785f95da295d40e1cb965)

The ID of the program if attached in hardware offload mode.

#### `skb_prog_id`

[:octicons-tag-24: 0.7.0](https://github.com/libbpf/libbpf/commit/8fbe7eec3aacca51d81785f95da295d40e1cb965)

The ID of the program if attached in SKB mode.

#### `attach_mode`

[:octicons-tag-24: 0.7.0](https://github.com/libbpf/libbpf/commit/8fbe7eec3aacca51d81785f95da295d40e1cb965)

The mode the program is attached in. One of:

```c
enum {
	XDP_ATTACHED_NONE = 0,
	XDP_ATTACHED_DRV,
	XDP_ATTACHED_SKB,
	XDP_ATTACHED_HW,
	XDP_ATTACHED_MULTI,
};
```

#### `feature_flags`

[:octicons-tag-24: 1.2.0](https://github.com/libbpf/libbpf/commit/547881e04e771d050f7b450bcbaa4c19e64e9654)

XDP features the network interface is capable of. A combination of:

```c
/**
 * enum netdev_xdp_act
 * @NETDEV_XDP_ACT_BASIC: XDP features set supported by all drivers
 *   (XDP_ABORTED, XDP_DROP, XDP_PASS, XDP_TX)
 * @NETDEV_XDP_ACT_REDIRECT: The netdev supports XDP_REDIRECT
 * @NETDEV_XDP_ACT_NDO_XMIT: This feature informs if netdev implements
 *   ndo_xdp_xmit callback.
 * @NETDEV_XDP_ACT_XSK_ZEROCOPY: This feature informs if netdev supports AF_XDP
 *   in zero copy mode.
 * @NETDEV_XDP_ACT_HW_OFFLOAD: This feature informs if netdev supports XDP hw
 *   offloading.
 * @NETDEV_XDP_ACT_RX_SG: This feature informs if netdev implements non-linear
 *   XDP buffer support in the driver napi callback.
 * @NETDEV_XDP_ACT_NDO_XMIT_SG: This feature informs if netdev implements
 *   non-linear XDP buffer support in ndo_xdp_xmit callback.
 */
enum netdev_xdp_act {
	NETDEV_XDP_ACT_BASIC = 1,
	NETDEV_XDP_ACT_REDIRECT = 2,
	NETDEV_XDP_ACT_NDO_XMIT = 4,
	NETDEV_XDP_ACT_XSK_ZEROCOPY = 8,
	NETDEV_XDP_ACT_HW_OFFLOAD = 16,
	NETDEV_XDP_ACT_RX_SG = 32,
	NETDEV_XDP_ACT_NDO_XMIT_SG = 64,
};
```

#### `xdp_zc_max_segs`

[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/commit/8ae70bcbdf996369502e5613510c17687ce3d7ad)

An unsigned integer stating the max number of [frags](../../../linux/program-type/BPF_PROG_TYPE_XDP.md#xdp-fragments) that are supported by this device in [XSK](../../../linux/concepts/af_xdp.md) [zero-copy mode](../../../linux/concepts/af_xdp.md#zero-copy-mode).

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
