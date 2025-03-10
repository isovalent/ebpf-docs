---
title: "Libbpf userspace function 'bpf_program__attach_iter'"
description: "This page documents the 'bpf_program__attach_iter' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_iter`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.9](https://github.com/libbpf/libbpf/releases/tag/v0.0.9)
<!-- [/LIBBPF_TAG] -->

Attach a [iterator](../../../linux/program-type/BPF_PROG_TYPE_TRACING.md/#iterator) program.

## Definition

`#!c struct bpf_link * bpf_program__attach_iter(const struct bpf_program *prog, const struct bpf_iter_attach_opts *opts);`

**Parameters**

- `prog`: BPF program to attach
- `opts`: options for attaching the iterator program

**Return**

Reference to the newly created BPF link; or `NULL` is returned on error, error code is stored in [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html)

### `struct bpf_iter_attach_opts`

```c
struct bpf_iter_attach_opts {
	size_t sz; /* size of this struct for forward/backward compatibility */
	union bpf_iter_link_info *link_info;
	__u32 link_info_len;
};
```

#### `link_info`

```c
union bpf_iter_link_info {
	struct {
		__u32	map_fd;
	} map;
};
```

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
