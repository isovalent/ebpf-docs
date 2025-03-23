---
title: "Libbpf userspace function 'bpf_token_create'"
description: "This page documents the 'bpf_token_create' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_token_create`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.4.0](https://github.com/libbpf/libbpf/releases/tag/v1.4.0)
<!-- [/LIBBPF_TAG] -->

Creates a new instance of BPF token derived from specified BPF file system mount point. A wrapper around the [`BPF_TOKEN_CREATE`](../../../linux/syscall/BPF_TOKEN_CREATE.md) syscall command.

## Definition

`#!c int bpf_token_create(int bpffs_fd, struct bpf_token_create_opts *opts);`

**Parameters**

- `bpffs_fd`: FD for BPF FS instance from which to derive a BPF token instance.
- `opts`: optional BPF token creation options, can be `NULL`

**Return**

BPF token FD > `0`, on success; negative error code, otherwise ([`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) is also set to the error code)

### `struct bpf_token_create_opts`

```c
struct bpf_token_create_opts {
	size_t sz; /* size of this struct for forward/backward compatibility */
	__u32 flags;
	size_t :0;
};
```

## Usage

BPF token created with this API can be passed to BPF syscall for commands like [`BPF_PROG_LOAD`](../../../linux/syscall/BPF_PROG_LOAD.md), [`BPF_MAP_CREATE`](../../../linux/syscall/BPF_MAP_CREATE.md), etc.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
