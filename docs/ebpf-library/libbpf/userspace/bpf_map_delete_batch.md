---
title: "Libbpf userspace function 'bpf_map_delete_batch'"
description: "This page documents the 'bpf_map_delete_batch' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map_delete_batch`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.7](https://github.com/libbpf/libbpf/releases/tag/v0.0.7)
<!-- [/LIBBPF_TAG] -->

Allows for batch deletion of multiple elements in a BPF map.

## Definition

`#!c int bpf_map_delete_batch(int fd, const void *keys, __u32 *count, const struct bpf_map_batch_opts *opts);`

**Parameters**

- `fd`: BPF map file descriptor
- `keys`: pointer to an array of `count` keys
- `count`: input and output parameter; on input `count` represents the number of  elements in the map to delete in batch; on output if a non`-EFAULT` error is returned, `count` represents the number of deleted elements if the output `count` value is not equal to the input `count` value If `-EFAULT` is returned, `count` should not be trusted to be correct.
- `opts`: options for configuring the way the batch deletion works

**Return**

`0`, on success; negative error code, otherwise ([`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) is also set to the error code)

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
