---
title: "Libbpf userspace function 'bpf_map_update_batch'"
description: "This page documents the 'bpf_map_update_batch' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map_update_batch`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.7](https://github.com/libbpf/libbpf/releases/tag/v0.0.7)
<!-- [/LIBBPF_TAG] -->

Updates multiple elements in a map by specifying keys and their corresponding values.

## Definition

`#!c int bpf_map_update_batch(int fd, const void *keys, const void *values, __u32 *count, const struct bpf_map_batch_opts *opts);`

**Parameters**

- `fd`: BPF map file descriptor
- `keys`: pointer to an array of `count` keys
- `values`: pointer to an array of `count` values
- `count`: input and output parameter; on input it's the number of elements in the map to update in batch; on output if a non`-EFAULT` error is returned,
`count` represents the number of updated elements if the output `count` value is not equal to the input `count` value. If `EFAULT` is returned, `count` should not be trusted to be correct.
- `opts`: options for configuring the way the batch update works

**Return**

`0`, on success; negative error code, otherwise ([`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) is also set to the error code)

## Usage

The `keys` and `values` parameters must point to memory large enough to hold `count` items based on the key and value size of the map.

The `opts` parameter can be used to control how [`bpf_map_update_batch`](bpf_map_update_batch.md) should handle keys that either do or do not already exist in the map. In particular the `flags` parameter of `bpf_map_batch_opts` can be one of the following:

* `BPF_ANY` - Create new elements or update existing.
* `BPF_NOEXIST` -  Create new elements only if they do not exist.
* `BPF_EXIST` - Update existing elements.
* `BPF_F_LOCK` - Update spin_lock-ed map elements. This must be specified if the map value contains a spinlock.

!!! note
   `count` is an input and output parameter, where on output it represents how many elements were successfully updated. Also note that if `EFAULT` then `count` should not be trusted to be correct.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
