# Map type `BPF_MAP_TYPE_REUSEPORT_SOCKARRAY`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_REUSEPORT_SOCKARRAY) -->
[:octicons-tag-24: v4.19](https://github.com/torvalds/linux/commit/5dc4c4b7d4e8115e7cde96a030f98cb3ab2e458c)
<!-- [/FEATURE_TAG] -->

This is a special purpose map type which holds references to sockets with the `SO_REUSEPORT` option set.

## Usage

This map is used together with the [bpf_sk_select_reuseport](../helper-function/bpf_sk_select_reuseport.md) helper function to select a socket from a reuse port socket group to handle traffic.

## Attributes

The [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) can be `4` or `8` and the [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) must always be `4`. 

<!-- TODO link to generic page for attributes which are the same for every map type -->

## Syscall commands

The following syscall commands work with this map type:

* [`BPF_MAP_LOOKUP_ELEM`](../syscall/BPF_MAP_LOOKUP_ELEM.md)
* [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md)
* [`BPF_MAP_DELETE_ELEM`](../syscall/BPF_MAP_DELETE_ELEM.md)

## Helper functions

<!-- DO NOT EDIT MANUALLY -->
<!-- [MAP_HELPER_FUNC_REF] -->
 * [bpf_sk_select_reuseport](../helper-function/bpf_sk_select_reuseport.md)
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
