---
title: "Libbpf userspace function 'user_ring_buffer__reserve_blocking'"
description: "This page documents the 'user_ring_buffer__reserve_blocking' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `user_ring_buffer__reserve_blocking`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.1.0](https://github.com/libbpf/libbpf/releases/tag/v1.1.0)
<!-- [/LIBBPF_TAG] -->

Reserves a record in the ring buffer, possibly blocking for up to `timeout_ms` until a sample becomes available.

## Definition

`#!c void *user_ring_buffer__reserve_blocking(struct user_ring_buffer *rb, __u32 size, int timeout_ms);`

**Parameters**

- `rb`: The user ring buffer.
- `size`: The size of the sample, in bytes.
- `timeout_ms`: The amount of time, in milliseconds, for which the caller should block when waiting for a sample. `-1` causes the caller to block indefinitely.

**Return**

A pointer to an 8-byte aligned reserved region of the user ring buffer; `NULL`, and [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) being set if a sample could not be reserved.

## Usage

This function is `not` thread safe, and callers must synchronize accessing this function if there are multiple producers.

If `timeout_ms` is `-1`, the function will block indefinitely until a sample becomes available. Otherwise, `timeout_ms` must be non-negative, or [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) is set to EINVAL, and `NULL` is returned. If `timeout_ms` is 0, no blocking will occur and the function will return immediately after attempting to reserve a sample.

If `size` is larger than the size of the entire ring buffer, [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) is set to `E2BIG` and `NULL` is returned. If the ring buffer could accommodate `size`, but currently does not have enough space, the caller will block until at most `timeout_ms` has elapsed. If insufficient space is available at that time, [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) is set to `ENOSPC`, and `NULL` is returned.

The kernel guarantees that it will wake up this thread to check if sufficient space is available in the ring buffer at least once per invocation of the [`bpf_user_ringbuf_drain`](../../../linux/helper-function/bpf_user_ringbuf_drain.md) helper function, provided that at least one sample is consumed, and the BPF program did not invoke the function with `BPF_RB_NO_WAKEUP`. A wakeup may occur sooner than that, but the kernel does not guarantee this. If the helper function is invoked with `BPF_RB_FORCE_WAKEUP`, a wakeup event will be sent even if no sample is consumed.

When a sample of size `size` is found within `timeout_ms`, a pointer to the sample is returned. After initializing the sample, callers must invoke [`user_ring_buffer__submit`](user_ring_buffer__submit.md) to post the sample to the ring buffer. Otherwise, the sample must be freed with [`user_ring_buffer__discard`](user_ring_buffer__discard.md).

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
