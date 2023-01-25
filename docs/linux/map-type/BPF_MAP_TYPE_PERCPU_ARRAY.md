# Map type `BPF_MAP_TYPE_PERCPU_ARRAY`

[:octicons-tag-24: v4.6](https://github.com/torvalds/linux/commit/a10423b87a7eae75da79ce80a8d9475047a674ee)

This is the per-CPU variant of the [`BPF_MAP_TYPE_ARRAY`](BPF_MAP_TYPE_ARRAY.md) map type. 

This map type is a generic map type with no restrictions on the structure of the value. However the key is a numeric index between 0 and [`max_entries`](../syscall/BPF_MAP_CREATE.md#max_entries)

This per-CPU version has a separate array for each logical CPU. When accessing the map using most [helper function](../helper-function/index.md), the array assigned to the CPU the eBPF program is currently running on is accessed implicitly. 

Since preemption is disabled during program execution, no other programs will be able to concurrently access the same memory. This guarantees there will never be any race conditions and improves the performance due to the lack of congestion and synchronization logic, at the cost of having a large memory footprint.

<!-- TODO: On newer kernels CPU migration is disabled, not preemption, check the implications of that against the above statements -->
<!-- TODO: "preemption" need a link -->
## Attributes

While the [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) is essentially unrestricted, the [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) must always be `4` indicating the key is a 32-bit unsigned integer.

<!-- TODO link to generic page for attributes which are the same for every map type -->

## Syscall commands

The following syscall commands work with this map type:

* [`BPF_MAP_LOOKUP_ELEM`](../syscall/BPF_MAP_LOOKUP_ELEM.md)
* [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md)
* [`BPF_MAP_GET_NEXT_KEY`](../syscall/BPF_MAP_GET_NEXT_KEY.md)
* [`BPF_MAP_LOOKUP_BATCH`](../syscall/BPF_MAP_LOOKUP_BATCH.md)

## Helper functions

The following helper functions work with this map type:

<!-- DO NOT EDIT MANUALLY -->
<!-- [MAP_HELPER_FUNC_REF] -->
 * [bpf_map_lookup_elem](../helper-function/bpf_map_lookup_elem.md)
 * [bpf_map_update_elem](../helper-function/bpf_map_update_elem.md)
 * [bpf_map_delete_elem](../helper-function/bpf_map_delete_elem.md)
 * [bpf_map_for_each_callback](../helper-function/bpf_map_for_each_callback.md)
 * [bpf_map_lookup_percpu_elem](../helper-function/bpf_map_lookup_percpu_elem.md)
<!-- [/MAP_HELPER_FUNC_REF] -->

## Flags

The following flags are supported by this map type.

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

<!-- ## Usage -->
<!-- ### Global data -->
<!-- TODO make an example -->

<!-- ### Memory mapping -->
<!-- TODO make an example -->
