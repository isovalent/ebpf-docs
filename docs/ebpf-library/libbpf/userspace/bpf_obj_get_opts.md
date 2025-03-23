---
title: "Libbpf userspace function 'bpf_obj_get_opts'"
description: "This page documents the 'bpf_obj_get_opts' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_obj_get_opts`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.0.0](https://github.com/libbpf/libbpf/releases/tag/v1.0.0)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_OBJ_GET`](../../../linux/syscall/BPF_OBJ_GET.md) syscall command.

## Definition

`#!c int bpf_obj_get_opts(const char *pathname, const struct bpf_obj_get_opts *opts);`

**Parameters**

- `pathname`: path to the object to retrieve
- `opts`: pointer to a `bpf_obj_get_opts` structure

**Return**

`>0`, file descriptor of the object; negative error code, otherwise

### `struct bpf_obj_get_opts`

```c
struct bpf_obj_get_opts {
    size_t sz; /* size of this struct for forward/backward compatibility */

    __u32 file_flags;
    int path_fd;

    size_t :0;
};
```

#### `file_flags`

[:octicons-tag-24: 1.0.0](https://github.com/libbpf/libbpf/commit/d8e2c9d9650ee15784f0413ae689a535c02f5e17) 

#### `path_fd`

[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/commit/a50544ef45c0ba1d02bcf283e86c2d90abaa29a1)


## Usage

This function should only be used if you need precise control over the object retrieval process. In most cases the [`bpf_object__open`](bpf_object__open.md) or similar high level API functions should be used instead.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
