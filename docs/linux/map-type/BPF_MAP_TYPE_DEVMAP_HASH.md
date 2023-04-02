# Map type `BPF_MAP_TYPE_DEVMAP_HASH`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_DEVMAP_HASH) -->
[:octicons-tag-24: v5.4](https://github.com/torvalds/linux/commit/6f9d451ab1a33728adb72d7ff66a7b374d665176)
<!-- [/FEATURE_TAG] -->

The device hash map is a specialized map type which holds references to network devices.

## Usage

This map type is used in combination with the [bpf_redirect_map](../helper-function/bpf_redirect_map.md) helper to redirect traffic to egress out of a different device.

Initially the value of this map was just the network interface index as `__u32`. But after [:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/fbee97feed9b3e4acdf9590e1f6b4a2eefecfffe) the value has been optionally extended to add a file descriptor to a secondary XDP program.

The C structure of the values look as follows:

```c
struct bpf_devmap_val {
	__u32 ifindex;   /* device index */
	union {
		int   fd;  /* prog fd on map write */
		__u32 id;  /* prog id on map read */
	} bpf_prog;
};
```

The `fd`/`id` refers to an XDP program optionally set by userspace. If set, the referred XDP program will execute on the packet, in the context of the new network device after the packet has been redirected but before it egresses the network interface.

!!! note
    Programs attached to a devmap must be loaded with the `BPF_XDP_DEVMAP` expected attach type.

## Attributes

The [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) can be `4` or `8` depending on kernel version and optional secondary program support. The [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) can be freely chosen.

<!-- TODO link to generic page for attributes which are the same for every map type -->

## Syscall commands

The following syscall commands work with this map type:

* [`BPF_MAP_LOOKUP_ELEM`](../syscall/BPF_MAP_LOOKUP_ELEM.md)
* [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md)
* [`BPF_MAP_DELETE_ELEM`](../syscall/BPF_MAP_DELETE_ELEM.md)

## Helper functions

<!-- DO NOT EDIT MANUALLY -->
<!-- [MAP_HELPER_FUNC_REF] -->
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
