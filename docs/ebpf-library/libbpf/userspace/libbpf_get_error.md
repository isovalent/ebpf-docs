---
title: "Libbpf userspace function 'libbpf_get_error'"
description: "This page documents the 'libbpf_get_error' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `libbpf_get_error`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Extracts the error code from the passed pointer

## Definition

`#!c long libbpf_get_error(const void *ptr);`

**Parameters**

- `ptr`: pointer returned from libbpf API function

**Return**

error code; or 0 if no error occurred

## Usage

As of libbpf 1.0 this function is not necessary and not recommended to be used. Libbpf doesn't return error code embedded into the pointer itself. Instead, `NULL` is returned on error and error code is passed through thread-local [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) variable. [`libbpf_get_error()`](libbpf_get_error.md) is just returning -[`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) value if it receives `NULL`, which is correct only if [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) hasn't been modified between libbpf API call and corresponding [`libbpf_get_error()`](libbpf_get_error.md) call. Prefer to check return for `NULL` and use [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) directly.

This API is left in libbpf 1.0 to allow applications that were 1.0-ready before final libbpf 1.0 without needing to change them.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
