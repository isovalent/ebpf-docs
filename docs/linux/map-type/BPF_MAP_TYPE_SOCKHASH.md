# Map type `BPF_MAP_TYPE_SOCKHASH`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_SOCKHASH) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/81110384441a59cff47430f20f049e69b98c17f4)
<!-- [/FEATURE_TAG] -->

The socket map is a specialized map type which hold network sockets as value.

## Usage

This map type can be use too lookup a pointer to a socket with the [bpf_map_lookup_elem](../helper-function/bpf_map_lookup_elem.md) helper function, which then can be passed to helpers such as [`bpf_sk_assign`](../helper-function/bpf_sk_assign.md) or the a map reference can be used directly in a range of helpers such as [bpf_sk_redirect_map](../helper-function/bpf_sk_redirect_map.md), [bpf_msg_redirect_map](../helper-function/bpf_msg_redirect_map.md) and [bpf_sk_select_reuseport](../helper-function/bpf_sk_select_reuseport.md). All of the above cases redirect a packet or connection in some way, the details differ depending on the program type and the helper function, so please visit the specific pages for details.

This map can also be manipulated from kernel space, the main use-case for doing so seems to be to manage the contents of the map automatically from program types that trigger on socket events. This would allow 1 program to manage the contents of the map, and another to do the actual redirecting on packet events.

## Attributes

The [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) must always be `4` and the [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) must always be `8`. 

<!-- TODO link to generic page for attributes which are the same for every map type -->

## Syscall commands

The following syscall commands work with this map type:

* [`BPF_MAP_LOOKUP_ELEM`](../syscall/BPF_MAP_LOOKUP_ELEM.md)
* [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md)
* [`BPF_MAP_DELETE_ELEM`](../syscall/BPF_MAP_DELETE_ELEM.md)

## Helper functions

<!-- DO NOT EDIT MANUALLY -->
<!-- [MAP_HELPER_FUNC_REF] -->
 * [bpf_sk_redirect_hash](../helper-function/bpf_sk_redirect_hash.md)
 * [bpf_sock_hash_update](../helper-function/bpf_sock_hash_update.md)
 * [bpf_map_delete_elem](../helper-function/bpf_map_delete_elem.md)
 * [bpf_msg_redirect_hash](../helper-function/bpf_msg_redirect_hash.md)
 * [bpf_sk_select_reuseport](../helper-function/bpf_sk_select_reuseport.md)
 * [bpf_map_lookup_elem](../helper-function/bpf_map_lookup_elem.md)
 * [bpf_map_update_elem](../helper-function/bpf_map_update_elem.md)
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

### `BPF_F_RDONLY_PROG`

[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/591fe9888d7809d9ee5c828020b6c6ae27c37229)

Setting this flag will make it so the map can only be read via [helper functions](../helper-function/index.md), but not written to.

For details please check the [generic description](../syscall/BPF_MAP_CREATE.md#bpf_f_rdonly_prog).

### `BPF_F_WRONLY_PROG`

[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/591fe9888d7809d9ee5c828020b6c6ae27c37229)

Setting this flag will make it so the map can only be written to via [helper functions](../helper-function/index.md), but not read from.

For details please check the [generic description](../syscall/BPF_MAP_CREATE.md#bpf_f_wronly_prog).
