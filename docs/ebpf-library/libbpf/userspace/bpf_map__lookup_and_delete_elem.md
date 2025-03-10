---
title: "Libbpf userspace function 'bpf_map__lookup_and_delete_elem'"
description: "This page documents the 'bpf_map__lookup_and_delete_elem' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__lookup_and_delete_elem`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)
<!-- [/LIBBPF_TAG] -->

Allows to lookup BPF map value corresponding to provided key and atomically delete it afterwards.

## Definition

`#!c int bpf_map__lookup_and_delete_elem(const struct bpf_map *map, const void *key, size_t key_sz, void *value, size_t value_sz, __u64 flags);`

**Parameters**

- `map`: BPF map to lookup element in
- `key`: pointer to memory containing bytes of the key used for lookup
- `key_sz`: size in bytes of key data, needs to match BPF map definition's `key_size`
- `value`: pointer to memory in which looked up value will be stored
- `value_sz`: size in byte of value data memory; it has to match BPF map definition's `value_size`. For per-CPU BPF maps value size has to be a product of BPF map value size and number of possible CPUs in the system (could be fetched with [`libbpf_num_possible_cpus()`](libbpf_num_possible_cpus.md)). Note also that for per-CPU values value size has to be aligned up to closest 8 bytes for alignment reasons, so expected size is: `round_up(value_size, 8) * libbpf_num_possible_cpus()`.
- `flags`: flags passed to kernel for this operation

**Return**

`0`, on success; negative error, otherwise

## Usage

[`bpf_map__lookup_and_delete_elem()`](bpf_map__lookup_and_delete_elem.md) is high-level equivalent of [`bpf_map_lookup_and_delete_elem()`](bpf_map_lookup_and_delete_elem.md) API with added check for key and value size.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
