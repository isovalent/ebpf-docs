---
title: "Map Type 'BPF_MAP_TYPE_LPM_TRIE'"
description: "This page documents the 'BPF_MAP_TYPE_LPM_TRIE' eBPF map type, including its definition, usage, program types that can use it, and examples."
---
# Map type `BPF_MAP_TYPE_LPM_TRIE`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_LPM_TRIE) -->
[:octicons-tag-24: v4.11](https://github.com/torvalds/linux/commit/b95a5c4db09bc7c253636cb84dc9b12c577fd5a0)
<!-- [/FEATURE_TAG] -->

The LPM (Longest Prefix Match) map is a generic map type which does prefix matching on the key upon lookup.

## Usage

One of the main use cases for this map type is to implement routing tables or policies for IP ranges. Take the following key-value pairs:

* `10.0.0.0/8`     -> 1
* `10.0.10.0/24`   -> 2
* `10.0.10.123/32` -> 3

A lookup for `10.0.10.123` will return value 3 because we have a specific entry for it in the map. A lookup for `10.0.10.200` will return value 2, because the `/24` key is more specific than the `/8` key. A lookup for `10.12.0.1` would return 1. And a lookup for `12.0.0.1` will not return any entry.

## Attributes

The [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) is unrestricted. The [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) is variable, but must adhere to a specific structure. The first 4 bytes of the key indicate the prefix length of the key (the number after the `/` in a CIDR), the rest of the bytes are the value.

The `bpf.h` header file provides the following struct as example:

```c
struct bpf_lpm_trie_key {
	__u32	prefixlen;	/* up to 32 for AF_INET, 128 for AF_INET6 */
	__u8	data[0];	/* Arbitrary size */
};
```

A common trick is to embed this struct in a custom key type:

```c
struct lpm_key {
	struct bpf_lpm_trie_key trie_key;
	__u32 data;
};
```

But this is not required, as long as the first field is a `__u32`:

```c
struct lpm_key {
	__u32 prefixlen;
	__be32 addr;
};
```

LPM tries may be created with a maximum prefix length that is a multiple of 8, in the range from 8 to 2048. So the `data` part of the key can't be larger than 256 bytes (4 or 16 is normal for IPv4 and IPv6 addresses).

> Data stored in @data of struct `bpf_lpm_key` and struct `lpm_trie_node` is interpreted as big endian, so data[0] stores the most significant byte. [...] one single element that matches 192.168.0.0/16. The data array would hence contain `[0xc0, 0xa8, 0x00, 0x00]` in big-endian notation.

!!! note
    The `BPF_F_NO_PREALLOC` flag must always be set when creating this map type since the implementation cannot pre-allocate the map.

## Syscall commands

The following syscall commands work with this map type:

* [`BPF_MAP_LOOKUP_ELEM`](../syscall/BPF_MAP_LOOKUP_ELEM.md)
* [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md)
* [`BPF_MAP_DELETE_ELEM`](../syscall/BPF_MAP_DELETE_ELEM.md)
* [`BPF_MAP_GET_NEXT_KEY`](../syscall/BPF_MAP_GET_NEXT_KEY.md)

## Helper functions

The following helper functions work with this map type:

<!-- DO NOT EDIT MANUALLY -->
<!-- [MAP_HELPER_FUNC_REF] -->
 * [`bpf_map_delete_elem`](../helper-function/bpf_map_delete_elem.md)
 * [`bpf_map_lookup_elem`](../helper-function/bpf_map_lookup_elem.md)
 * [`bpf_map_update_elem`](../helper-function/bpf_map_update_elem.md)
<!-- [/MAP_HELPER_FUNC_REF] -->

## Flags

The following flags are supported by this map type.

### `BPF_F_NO_PREALLOC`
[:octicons-tag-24: v4.6](https://github.com/torvalds/linux/commit/6c90598174322b8888029e40dd84a4eb01f56afe)

Hash maps are pre-allocated by default, this means that even a completely empty hash map will use the same amount of
kernel memory as a full map. 

If this flag is set, pre-allocation is disabled. Users might consider this for large maps since allocating large amounts of memory takes a lot of time during creation and might be undesirable.

!!! note
    Setting this flag during map creation is mandatory for the LPM map type.

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

For details please check the [generic description](../syscall/BPF_MAP_CREATE.md#bpf_f_wronly).

### `BPF_F_RDONLY_PROG`

[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/591fe9888d7809d9ee5c828020b6c6ae27c37229)

Setting this flag will make it so the map can only be read via [helper functions](../helper-function/index.md), but not written to.

For details please check the [generic description](../syscall/BPF_MAP_CREATE.md#bpf_f_rdonly_prog).

<!-- TODO:  -->

### `BPF_F_WRONLY_PROG`

[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/591fe9888d7809d9ee5c828020b6c6ae27c37229)

Setting this flag will make it so the map can only be written to via [helper functions](../helper-function/index.md), but not read from.

For details please check the [generic description](../syscall/BPF_MAP_CREATE.md#bpf_f_wronly_prog).

## Example

```c
#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>

struct ipv4_lpm_key {
        __u32 prefixlen;
        __u32 data;
};

struct {
        __uint(type, BPF_MAP_TYPE_LPM_TRIE);
        __type(key, struct ipv4_lpm_key);
        __type(value, __u32);
        __uint(map_flags, BPF_F_NO_PREALLOC);
        __uint(max_entries, 255);
} ipv4_lpm_map SEC(".maps");

void *lookup(__u32 ipaddr)
{
        struct ipv4_lpm_key key = {
                .prefixlen = 32,
                .data = ipaddr
        };

        return bpf_map_lookup_elem(&ipv4_lpm_map, &key);
}

```
