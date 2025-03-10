---
title: "Libbpf userspace function 'bpf_map__ifindex'"
description: "This page documents the 'bpf_map__ifindex' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__ifindex`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.1.0](https://github.com/libbpf/libbpf/releases/tag/v0.1.0)
<!-- [/LIBBPF_TAG] -->

Get the interface index associated with the BPF map.

## Definition

`#!c __u32 bpf_map__ifindex(const struct bpf_map *map);`

**Parameters**

- `map`: Pointer to the BPF map.

**Return**

Interface index associated with the map. Or `0` if no interface index is available.

## Usage

When offloading XDP programs to the hardware, any BPF map that is used by the XDP program is also offloaded to the hardware. Since the driver might impose restrictions on maps, the interface index on which the program/map will be offloaded is associated with it, so the driver for that interface can be looped in during map loading.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
