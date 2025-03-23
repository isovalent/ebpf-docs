---
title: "Libbpf userspace function 'bpf_btf_get_next_id'"
description: "This page documents the 'bpf_btf_get_next_id' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_btf_get_next_id`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.5](https://github.com/libbpf/libbpf/releases/tag/v0.0.5)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_BTF_GET_NEXT_ID`](../../../linux/syscall/BPF_BTF_GET_NEXT_ID.md) syscall command.

## Definition

`#!c int bpf_btf_get_next_id(__u32 start_id, __u32 *next_id);`

**Parameters**

- `start_id`: BTF object ID to start from
- `next_id`: next BTF object ID

**Return**

`0`, on success; negative error code, otherwise

## Usage

This function allows you to iterate over loaded BTF objects. It returns the next BTF object ID after the `start_id` provided. If `start_id` is `0`, the first BTF object ID is returned.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
