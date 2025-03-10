---
title: "Libbpf userspace function 'bpf_program__attach_netkit'"
description: "This page documents the 'bpf_program__attach_netkit' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_netkit`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/releases/tag/v1.3.0)
<!-- [/LIBBPF_TAG] -->

Attach a netkit program to a netkit network interface.

## Definition

`#!c struct bpf_link * bpf_program__attach_netkit(const struct bpf_program *prog, int ifindex, const struct bpf_netkit_opts *opts);`

**Parameters**

- `prog`: BPF program to attach
- `ifindex`: index of the network interface to attach the program to
- `opts`: netkit options

**Return**

Reference to the newly created BPF link; or `NULL` is returned on error, error code is stored in [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html)

### `struct bpf_netkit_opts`

```c
struct bpf_netkit_opts {
	/* size of this struct, for forward/backward compatibility */
	size_t sz;
	__u32 flags;
	__u32 relative_fd;
	__u32 relative_id;
	__u64 expected_revision;
	size_t :0;
};
```


#### `flags`

Flags to modify attachment behavior.

`BPF_F_ID` - If set, identify a relative object based on the ID in `relative_id`. Otherwise use the FD in `relative_fd`.
`BPF_F_LINK` - If set, the `relative_fd`/`relative_id` fields refer to links, not programs.
`BPF_F_REPLACE` - If set, replace the existing program/link with the relative one.
`BPF_F_BEFORE` - If set, add the new program/link before the relative one.
`BPF_F_AFTER` - If set, add the new program/link after the relative one.

#### `relative_fd`

File descriptor of the relative object.

#### `relative_id`

ID of the relative object.

#### `expected_revision`

The current <nospell>mprog</nospell> revision. A revision number shared by all programs attached to the same hook point. It can be queried via [`BPF_PROG_QUERY`](../../../linux/syscall/BPF_PROG_QUERY.md) / [`bpf_prog_query`](bpf_prog_query.md) / [`bpf_prog_query_opts`](bpf_prog_query_opts.md).

## Usage

Netkit programs essentially TCX programs with more restrictions (namely that netkit programs can only attach to netkit links). Their attachment is very similar.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
