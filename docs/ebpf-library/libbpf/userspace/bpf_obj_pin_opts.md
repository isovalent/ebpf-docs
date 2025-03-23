---
title: "Libbpf userspace function 'bpf_obj_pin_opts'"
description: "This page documents the 'bpf_obj_pin_opts' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_obj_pin_opts`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/releases/tag/v1.3.0)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_OBJ_PIN`](../../../linux/syscall/BPF_OBJ_PIN.md) syscall command.

## Definition

`#!c int bpf_obj_pin_opts(int fd, const char *pathname, const struct bpf_obj_pin_opts *opts);`

**Parameters**

- `fd`: file descriptor of the object to pin
- `pathname`: path to the directory where the object will be pinned
- `opts`: pointer to a `bpf_obj_pin_opts` structure

**Return**

`0`, on success; negative error code, otherwise

### `struct bpf_obj_pin_opts`

```c
struct bpf_obj_pin_opts {
	size_t sz; /* size of this struct for forward/backward compatibility */

	__u32 file_flags;
	int path_fd;

	size_t :0;
};
```

#### `file_flags`

[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/commit/a50544ef45c0ba1d02bcf283e86c2d90abaa29a1)


#### `path_fd`

[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/commit/a50544ef45c0ba1d02bcf283e86c2d90abaa29a1)

## Usage

This function should only be used if you need precise control over the object pinning process. In most cases the [`bpf_object__pin`](bpf_object__pin.md) function should be used instead.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
