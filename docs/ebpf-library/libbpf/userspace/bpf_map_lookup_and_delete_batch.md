---
title: "Libbpf userspace function 'bpf_map_lookup_and_delete_batch'"
description: "This page documents the 'bpf_map_lookup_and_delete_batch' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map_lookup_and_delete_batch`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.7](https://github.com/libbpf/libbpf/releases/tag/v0.0.7)
<!-- [/LIBBPF_TAG] -->

Allows for batch lookup and deletion of BPF map elements where each element is deleted after being retrieved.

## Definition

`#!c int bpf_map_lookup_and_delete_batch(int fd, void *in_batch, void *out_batch, void *keys, void *values, __u32 *count, const struct bpf_map_batch_opts *opts);`

**Parameters**

- `fd`: BPF map file descriptor
- `in_batch`: address of the first element in batch to read, can pass `NULL` to get address of the first element in `out_batch`. If not `NULL`, must be large enough to hold a key. For **BPF_MAP_TYPE_{HASH, PERCPU_HASH, LRU_HASH, LRU_PERCPU_HASH}**, the memory size must be at least 4 bytes wide regardless of key size.
- `out_batch`: output parameter that should be passed to next call as `in_batch`
- `keys`: pointer to an array of `count` keys
- `values`: pointer to an array large enough for `count` values
- `count`: input and output parameter; on input it's the number of elements in the map to read and delete in batch; on output it represents the number of elements that were successfully read and deleted If a non-`EFAULT` error code is returned and if the output `count` value is not equal to the input `count` value, up to `count` elements may have been deleted. if `EFAULT` is returned up to `count` elements may have been deleted without being returned via the `keys` and `values` output parameters.
- `opts`: options for configuring the way the batch lookup and delete works

**Return**

`0`, on success; negative error code, otherwise ([`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) is also set to the error code)

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
