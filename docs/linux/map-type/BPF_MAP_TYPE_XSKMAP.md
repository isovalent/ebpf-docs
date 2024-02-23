---
title: "Map Type 'BPF_MAP_TYPE_XSKMAP' - eBPF Docs"
description: "This page documents the 'BPF_MAP_TYPE_XSKMAP' eBPF map type, including its defintion, usage, program types that can use it, and examples."
---
# Map type `BPF_MAP_TYPE_XSKMAP`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_XSKMAP) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/fbfc504a24f53f7ebe128ab55cb5dba634f4ece8)
<!-- [/FEATURE_TAG] -->

This XDP Socket map is a specialized map which references XDP Sockets.

## Usage

This map type is used in combination with the [bpf_redirect_map](../helper-function/bpf_redirect_map.md) helper to redirect traffic to userspace, bypassing the kernel network stack. It is an array style map, where the indices go from `0` to `max_entries-1`. The values of this map are the file descriptor of specially prepared network sockets.

For details on the usage and preparation of AF_XDP sockets checkout out the [concept page](../concepts/af_xdp.md) and/or [kernel docs](https://www.kernel.org/doc/html/latest/networking/af_xdp.html).

### Example

The following examples shows how to use the map and helper from the BPF/kernel side. Important to note is that we always use the [`rx_queue_index`](../program-type/BPF_PROG_TYPE_XDP.md#rx_queue_index) as key, not doing so will cause the packet to be dropped.

```c
struct {
    __uint(type, BPF_MAP_TYPE_XSKMAP);
    __type(key, __u32);
    __type(value, __u32);
    __uint(max_entries, 64);
} xsks_map SEC(".maps");


SEC("xdp")
int xsk_redir_prog(struct xdp_md *ctx)
{
    __u32 index = ctx->rx_queue_index;

    if (bpf_map_lookup_elem(&xsks_map, &index))
            return bpf_redirect_map(&xsks_map, index, 0);
    return XDP_PASS;
}
```

## Attributes

The [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) must always be `4`. The [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) must always be `4`.

<!-- TODO link to generic page for attributes which are the same for every map type -->

## Syscall commands

The following syscall commands work with this map type:

* [`BPF_MAP_LOOKUP_ELEM`](../syscall/BPF_MAP_LOOKUP_ELEM.md)
* [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md)
* [`BPF_MAP_DELETE_ELEM`](../syscall/BPF_MAP_DELETE_ELEM.md)

## Helper functions

<!-- DO NOT EDIT MANUALLY -->
<!-- [MAP_HELPER_FUNC_REF] -->
 * [bpf_map_lookup_elem](../helper-function/bpf_map_lookup_elem.md)
 * [bpf_redirect_map](../helper-function/bpf_redirect_map.md)
<!-- [/MAP_HELPER_FUNC_REF] -->

## Flags

### `BPF_F_NUMA_NODE`

[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/96eabe7a40aa17e613cf3db2c742ee8b1fc764d0)

When set, the [`numa_node`](../syscall/BPF_MAP_CREATE.md#numa_node) attribute is respected during map creation.

### `BPF_F_RDONLY`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/6e71b04a82248ccf13a94b85cbc674a9fefe53f5)

Setting this flag will make it so the map can only be read via the [syscall](../syscall/index.md) interface, but not written to.

For details please check the [generic description](../syscall/BPF_MAP_CREATE.md#bpf_f_rdonly).

### `BPF_F_WRONLY`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/6e71b04a82248ccf13a94b85cbc674a9fefe53f5)

Setting this flag will make it so the map can only be written to via the [syscall](../syscall/index.md) interface, but not read from.
