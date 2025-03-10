---
title: "Libbpf userspace function 'bpf_program__attach_netfilter'"
description: "This page documents the 'bpf_program__attach_netfilter' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_netfilter`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/releases/tag/v1.3.0)
<!-- [/LIBBPF_TAG] -->

Attach a [`BPF_PROG_TYPE_NETFILTER`](../../../linux/program-type/BPF_PROG_TYPE_NETFILTER.md) program to a netfilter hook.

## Definition

`#!c struct bpf_link * bpf_program__attach_netfilter(const struct bpf_program *prog, const struct bpf_netfilter_opts *opts);`

**Parameters**

- `prog`: BPF program to attach
- `opts`: netfilter options

**Return**

Reference to the newly created BPF link; or `NULL` is returned on error, error code is stored in [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html)

### `struct bpf_netfilter_opts`

```c
struct bpf_netfilter_opts {
	/* size of this struct, for forward/backward compatibility */
	size_t sz;

	__u32 pf;
	__u32 hooknum;
	__s32 priority;
	__u32 flags;
};
```

#### `pf`

The protocol family of the hook.

#### `hooknum`

The hook number.

#### `priority`

The priority of the hook.

#### `flags`

The flags of the hook.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
