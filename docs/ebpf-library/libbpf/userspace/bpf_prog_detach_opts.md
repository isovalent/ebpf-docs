---
title: "Libbpf userspace function 'bpf_prog_detach_opts'"
description: "This page documents the 'bpf_prog_detach_opts' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_prog_detach_opts`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/releases/tag/v1.3.0)
<!-- [/LIBBPF_TAG] -->

Detaches the BPF program corresponding to `prog_fd` from a `target` which can represent a file descriptor or netdevice ifindex.

## Definition

`#!c int bpf_prog_detach_opts(int prog_fd, int target, enum bpf_attach_type type, const struct bpf_prog_detach_opts *opts);`

**Parameters**

- `prog_fd`: BPF program file descriptor
- `target`: detach location file descriptor or ifindex
- `type`: detach type for the BPF program
- `opts`: options for configuring the detachment

### `struct bpf_prog_detach_opts`

```c
struct bpf_prog_detach_opts {
	size_t sz; /* size of this struct for forward/backward compatibility */
	__u32 flags;
	int relative_fd;
	__u32 relative_id;
	__u64 expected_revision;
	size_t :0;
};
```

#### `flags`

[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/commit/bb5d7c1be8355c95b98223d6fe8d2c20c4bfcda9)

#### `relative_fd`

[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/commit/bb5d7c1be8355c95b98223d6fe8d2c20c4bfcda9)

#### `relative_id`

[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/commit/bb5d7c1be8355c95b98223d6fe8d2c20c4bfcda9)

#### `expected_revision`

[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/commit/bb5d7c1be8355c95b98223d6fe8d2c20c4bfcda9)

**Return**

`0`, on success; negative error code, otherwise ([`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) is also set to the error code)

## Usage

This function should only be used for specific program types that need to be detached via the `BPF_PROG_DETACH` syscall command and you need specific control over this process. In most cases, the [`bpf_link__detach`](bpf_link__detach.md) or [`bpf_link__destroy`](bpf_link__destroy.md) function should be used.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
