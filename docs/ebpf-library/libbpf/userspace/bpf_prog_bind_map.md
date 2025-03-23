---
title: "Libbpf userspace function 'bpf_prog_bind_map'"
description: "This page documents the 'bpf_prog_bind_map' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_prog_bind_map`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_PROG_BIND_MAP`](../../../linux/syscall/BPF_PROG_BIND_MAP.md) syscall command.

## Definition

`#!c int bpf_prog_bind_map(int prog_fd, int map_fd, const struct bpf_prog_bind_opts *opts);`

**Parameters**

- `prog_fd`: BPF program file descriptor
- `map_fd`: BPF map file descriptor
- `opts`: options for configuring the binding

**Return**

`0`, on success; negative error code, otherwise

### `struct bpf_prog_bind_opts`

```c
struct bpf_prog_bind_opts {
	size_t sz; /* size of this struct for forward/backward compatibility */
	__u32 flags;
};
```

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
