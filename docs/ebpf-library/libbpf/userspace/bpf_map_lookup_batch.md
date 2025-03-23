---
title: "Libbpf userspace function 'bpf_map_lookup_batch'"
description: "This page documents the 'bpf_map_lookup_batch' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map_lookup_batch`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.7](https://github.com/libbpf/libbpf/releases/tag/v0.0.7)
<!-- [/LIBBPF_TAG] -->

Allows for batch lookup of BPF map elements.

## Definition

`#!c int bpf_map_lookup_batch(int fd, void *in_batch, void *out_batch, void *keys, void *values, __u32 *count, const struct bpf_map_batch_opts *opts);`

**Parameters**

- `fd`: BPF map file descriptor
- `in_batch`: address of the first element in batch to read, can pass `NULL` to
indicate that the batched lookup starts from the beginning of the map.
- `out_batch`: output parameter that should be passed to next call as `in_batch`
- `keys`: pointer to an array large enough for `count` keys
- `values`: pointer to an array large enough for `count` values
- `count`: input and output parameter; on input it's the number of elements in the map to read in batch; on output it's the number of elements that were successfully read. If a non`-EFAULT` error is returned, count will be set as the number of elements that were read before the error occurred. If `-EFAULT` is returned, `count` should not be trusted to be correct.
- `opts`: options for configuring the way the batch lookup works

**Return**

`0`, on success; negative error code, otherwise ([`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) is also set to the error code)

## Usage

The parameter `in_batch` is the address of the first element in the batch to read. `out_batch` is an output parameter that should be passed as `in_batch` to subsequent calls to [`bpf_map_lookup_batch`](bpf_map_lookup_batch.md). `NULL` can be passed for `in_batch` to indicate that the batched lookup starts from the beginning of the map. Both `in_batch` and `out_batch` must point to memory large enough to hold a single key, except for maps of type `BPF_MAP_TYPE_{HASH, PERCPU_HASH, LRU_HASH, LRU_PERCPU_HASH}`, for which the memory size must be at least 4 bytes wide regardless of key size.

The `keys` and `values` are output parameters which must point to memory large enough to hold `count` items based on the key and value size of the map `map_fd`. The `keys` buffer must be of `key_size` * `count`. The `values` buffer must be of `value_size` * `count`.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
