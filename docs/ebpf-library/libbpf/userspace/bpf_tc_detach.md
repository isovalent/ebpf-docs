---
title: "Libbpf userspace function 'bpf_tc_detach'"
description: "This page documents the 'bpf_tc_detach' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_tc_detach`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.4.0](https://github.com/libbpf/libbpf/releases/tag/v0.4.0)
<!-- [/LIBBPF_TAG] -->

Detach a [`BPF_PROG_TYPE_SCHED_CLS`](../../../linux/program-type/BPF_PROG_TYPE_SCHED_CLS.md) program from a TC qdisc created with [`bpf_tc_hook_create`](bpf_tc_hook_create.md). 

## Definition

`#!c int bpf_tc_detach(const struct bpf_tc_hook *hook, const struct bpf_tc_opts *opts);`

**Parameters**

- `hook`: A pointer to a `struct bpf_tc_hook` that contains the hook information.
- `opts`: A pointer to a `struct bpf_tc_opts` that contains the options for the TC qdisc.

### `struct bpf_tc_hook`

```c
struct bpf_tc_hook {
	size_t sz;
	int ifindex;
	enum bpf_tc_attach_point attach_point;
	__u32 parent;
	size_t :0;
};
```

#### `ifindex`

The interface index of the device to attach the TC qdisc to.

#### `attach_point`

The TC attach point. This can be one of the following values:

```c
enum bpf_tc_attach_point {
	BPF_TC_INGRESS = 1 << 0,
	BPF_TC_EGRESS  = 1 << 1,
	BPF_TC_CUSTOM  = 1 << 2,
};
```

#### `parent`

The ID of the parent qdisc.

### `struct bpf_tc_opts`

```c
struct bpf_tc_opts {
	size_t sz;
	int prog_fd;
	__u32 flags;
	__u32 prog_id;
	__u32 handle;
	__u32 priority;
	size_t :0;
};
```

#### `prog_fd`

The file descriptor of the BPF program to attach.

#### `flags`

The flags for the TC qdisc. This can be one of the following values:

```c
enum bpf_tc_flags {
	BPF_TC_F_REPLACE = 1 << 0, // Replace existing program, fail if no program is attached
};
```

#### `prog_id`

This field is unused for this operation. Must be set to `0`.

#### `handle`

The handle of the TC classifier.

#### `priority`

The priority of the TC classifier. Must be a value between `0` and `65535`.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
