---
title: "Libbpf userspace function 'bpf_prog_query'"
description: "This page documents the 'bpf_prog_query' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_prog_query`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Queries the BPF programs and BPF links which are attached to `target` which can represent a file descriptor or netdevice ifindex.

## Definition

`#!c int bpf_prog_query(int target_fd, enum bpf_attach_type type, __u32 query_flags, __u32 *attach_flags, __u32 *prog_ids, __u32 *prog_cnt);`

**Parameters**

- `target`: query location file descriptor or ifindex
- `type`: attach type for the BPF program

**Return**

`0`, on success; negative error code, otherwise ([`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) is also set to the error code)

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
