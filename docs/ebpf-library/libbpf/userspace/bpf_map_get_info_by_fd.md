---
title: "Libbpf userspace function 'bpf_map_get_info_by_fd'"
description: "This page documents the 'bpf_map_get_info_by_fd' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map_get_info_by_fd`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.2.0](https://github.com/libbpf/libbpf/releases/tag/v1.2.0)
<!-- [/LIBBPF_TAG] -->

Obtains information about the BPF map corresponding to `map_fd`.

## Definition

`#!c int bpf_map_get_info_by_fd(int map_fd, struct bpf_map_info *info, __u32 *info_len);`

**Parameters**

- `map_fd`: BPF map file descriptor
- `info`: pointer to `struct bpf_map_info` that will be populated with BPF map information
- `info_len`: pointer to the size of `info`; on success updated with the number of bytes written to `info`

**Return**

`0`, on success; negative error code, otherwise ([`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) is also set to the error code)

## Usage

Populates up to `info_len` bytes of `info` and updates `info_len` with the actual number of bytes written to `info`. 

!!! note 
    `info` should be zero-initialized or initialized as expected by the requested `info` type. Failing to (zero-)initialize `info` under certain circumstances can result in this helper returning an error.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
