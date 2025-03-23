---
title: "Libbpf userspace function 'bpf_map_get_next_id'"
description: "This page documents the 'bpf_map_get_next_id' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map_get_next_id`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_MAP_GET_NEXT_ID`](../../../linux/syscall/BPF_MAP_GET_NEXT_ID.md) syscall command.

## Definition

`#!c int bpf_map_get_next_id(__u32 start_id, __u32 *next_id);`

**Parameters**

- `start_id`: map ID to start from
- `next_id`: next map ID

**Return**

`0`, on success; negative error code, otherwise

## Usage

This function allows you to iterate over loaded BPF maps. It returns the next map ID after the `start_id` provided. If `start_id` is `0`, the first map ID is returned.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
