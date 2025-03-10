---
title: "Libbpf userspace function 'bpf_map__set_ifindex'"
description: "This page documents the 'bpf_map__set_ifindex' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__set_ifindex`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Set the interface index associated with the BPF map.

## Definition

`#!c int bpf_map__set_ifindex(struct bpf_map *map, __u32 ifindex);`

**Parameters**

- `map`: Pointer to the BPF map.
- `ifindex`: Interface index associated with the BPF map.

**Return**

`0` on success, a negative error in case of failure.

## Usage

When offloading XDP programs to the hardware, any BPF map that is used by the XDP program is also offloaded to the hardware. Since the driver might impose restrictions on maps, the interface index on which the program/map will be offloaded is associated with it, so the driver for that interface can be looped in during map loading.

The interface index can only be set before the map is loaded.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
