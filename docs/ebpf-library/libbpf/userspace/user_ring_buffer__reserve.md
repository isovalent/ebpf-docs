---
title: "Libbpf userspace function 'user_ring_buffer__reserve'"
description: "This page documents the 'user_ring_buffer__reserve' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `user_ring_buffer__reserve`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.1.0](https://github.com/libbpf/libbpf/releases/tag/v1.1.0)
<!-- [/LIBBPF_TAG] -->

Reserves a pointer to a sample in the user ring buffer.

## Definition

`#!c void *user_ring_buffer__reserve(struct user_ring_buffer *rb, __u32 size);`

**Parameters**

- `rb`: A pointer to a user ring buffer.
- `size`: The size of the sample, in bytes.

**Return**

A pointer to an 8-byte aligned reserved region of the user ring
buffer; `NULL`, and [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) being set if a sample could not be reserved.

## Usage

This function is `not` thread safe, and callers must synchronize accessing this function if there are multiple producers. If a size is requested that is larger than the size of the entire ring buffer, [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) will be set to `E2BIG` and `NULL` is returned. If the ring buffer could accommodate the size, but currently does not have enough space, [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) is set to `ENOSPC` and `NULL` is returned.

After initializing the sample, callers must invoke [`user_ring_buffer__submit`](user_ring_buffer__submit.md) to post the sample to the kernel. Otherwise, the sample must be freed with [`user_ring_buffer__discard`](user_ring_buffer__discard.md).

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
