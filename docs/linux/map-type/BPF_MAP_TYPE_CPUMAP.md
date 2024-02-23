---
title: "Map Type 'BPF_MAP_TYPE_CPUMAP' - eBPF Docs"
description: "This page documents the 'BPF_MAP_TYPE_CPUMAP' eBPF map type, including its defintion, usage, program types that can use it, and examples."
---
# Map type `BPF_MAP_TYPE_CPUMAP`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_CPUMAP) -->
[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/6710e1126934d8b4372b4d2f9ae1646cd3f151bf)
<!-- [/FEATURE_TAG] -->

The CPU map is a specialized map type that holds references to logical CPUs.

## Usage

This map type can be used to redirect incoming packets to a different logical CPU. To do so, this map is used in combination with the [bpf_redirect_map](../helper-function/bpf_redirect_map.md) helper function.

This feature can for example be used to implement a form of [Receive Side Scaling (RSS)](https://www.kernel.org/doc/Documentation/networking/scaling.txt). This might be especially useful when dealing with network protocols for which Network Interface Cards, network drivers and/or the kernel have sub-optimal support.

Another theoretical example might be a multi tenancy situation where a set of logical CPUs is dedicated to a tenant. By redirecting traffic to its CPUs at the XDP level, the system can ensure heavy network load for one tenant does not impact others on the same system.

When packets are redirected, they are placed on a queue associated with each logical CPU. Initially the value of this map was just a single `__u32` representing the size of this queue. Writing a value of `0` disables that kv pair. The maximum queue size is `16384`.

After [:octicons-tag-24: v5.9](https://github.com/torvalds/linux/commit/9216477449f33cdbc9c9a99d49f500b7fbb81702), it becomes possible to add a secondary XDP program to the map entry which will be executed on the redirected packet on the new CPU. The new C structure of the map value look like this:

```c
struct bpf_cpumap_val {
	__u32 qsize;	/* queue size to remote target CPU */
	union {
		int   fd;	/* prog fd on map write */
		__u32 id;	/* prog id on map read */
	} bpf_prog;
};
```

!!! note
    Only programs which have been loaded with the `BPF_XDP_CPUMAP` expected attach type can be added to the `fd`/`id`.

## Attributes

The [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) can be `4` or `8` depending on kernel version and optional secondary program support. The [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) must always be `4`. The [`max_entries`](../syscall/BPF_MAP_CREATE.md#max_entries) must be smaller or equal to the amount of logical CPUs on the host.

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
