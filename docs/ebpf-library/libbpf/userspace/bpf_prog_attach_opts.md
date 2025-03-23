---
title: "Libbpf userspace function 'bpf_prog_attach_opts'"
description: "This page documents the 'bpf_prog_attach_opts' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_prog_attach_opts`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.8](https://github.com/libbpf/libbpf/releases/tag/v0.0.8)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_PROG_ATTACH`](../../../linux/syscall/BPF_PROG_ATTACH.md) syscall command.

## Definition

`#!c int bpf_prog_attach_opts(int prog_fd, int target, enum bpf_attach_type type, const struct bpf_prog_attach_opts *opts);`

**Parameters**

- `prog_fd`: BPF program file descriptor
- `target`: attach location file descriptor or ifindex
- `type`: attach type for the BPF program
- `opts`: options for configuring the attachment

### `struct bpf_prog_attach_opts`

```c
struct bpf_prog_attach_opts {
	size_t sz; /* size of this struct for forward/backward compatibility */
	__u32 flags;
	union {
		int replace_prog_fd;
		int replace_fd;
	};
	int relative_fd;
	__u32 relative_id;
	__u64 expected_revision;
	size_t :0;
};
```

#### `flags`

[:octicons-tag-24: 0.0.8](https://github.com/libbpf/libbpf/commit/8b20ffa4b913c13e2a0712b454af7bce65664003)

#### `replace_prog_fd`

[:octicons-tag-24: 0.0.8](https://github.com/libbpf/libbpf/commit/8b20ffa4b913c13e2a0712b454af7bce65664003)

#### `replace_fd`

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

Attaches the BPF program corresponding to `prog_fd` to a `target` which can represent a file descriptor or netdevice ifindex.

This function should only be used for specific program types that need to be attached via the `BPF_PROG_ATTACH` syscall command and you need specific control over this process. In most cases, the [`bpf_program__attach`](bpf_program__attach.md) or specific `bpf_program__attach_*` functions should be used instead.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
