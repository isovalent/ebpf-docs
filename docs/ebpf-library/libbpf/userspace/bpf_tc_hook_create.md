---
title: "Libbpf userspace function 'bpf_tc_hook_create'"
description: "This page documents the 'bpf_tc_hook_create' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_tc_hook_create`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.4.0](https://github.com/libbpf/libbpf/releases/tag/v0.4.0)
<!-- [/LIBBPF_TAG] -->

Create a TC qdisc to which a [`BPF_PROG_TYPE_SCHED_CLS`](../../../linux/program-type/BPF_PROG_TYPE_SCHED_CLS.md) program can be attached via [`bpf_tc_attach`](bpf_tc_attach.md).

## Definition

`#!c int bpf_tc_hook_create(struct bpf_tc_hook *hook);`

**Parameters**

- `hook`: A pointer to a `struct bpf_tc_hook` that will be filled with the hook information.

**Return**

`0` on success. A negative error code on failure.

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

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
