---
title: "Libbpf userspace function 'bpf_btf_get_fd_by_id_opts'"
description: "This page documents the 'bpf_btf_get_fd_by_id_opts' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_btf_get_fd_by_id_opts`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.1.0](https://github.com/libbpf/libbpf/releases/tag/v1.1.0)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_BTF_GET_FD_BY_ID`](../../../linux/syscall/BPF_BTF_GET_FD_BY_ID.md) syscall command.

## Definition

`#!c int bpf_btf_get_fd_by_id_opts(__u32 id, const struct bpf_get_fd_by_id_opts *opts);`

**Parameters**

- `id`: BTF object ID
- `opts`: options for configuring the file descriptor retrieval

**Return**

`>0`, a file descriptor for the BTF object; negative error code, otherwise

### `struct bpf_get_fd_by_id_opts`

```c
struct bpf_get_fd_by_id_opts {
	size_t sz; /* size of this struct for forward/backward compatibility */
	__u32 open_flags; /* permissions requested for the operation on fd */
	size_t :0;
};
```

#### `open_flags`

[:octicons-tag-24: 1.1.0](https://github.com/libbpf/libbpf/commit/a719cae6aaa3bd40b553329a936f8783510f9d71)

## Usage

This function returns the file descriptor of the BTF object with the given `id`. This causes the current process to hold a reference to the BTF object, preventing the kernel from unloading it. The file descriptor should be closed with [`close`](https://man7.org/linux/man-pages/man2/close.2.html) when it is no longer needed.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
