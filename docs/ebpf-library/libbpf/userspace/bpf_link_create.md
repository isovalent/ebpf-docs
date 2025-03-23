---
title: "Libbpf userspace function 'bpf_link_create'"
description: "This page documents the 'bpf_link_create' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_link_create`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.8](https://github.com/libbpf/libbpf/releases/tag/v0.0.8)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_LINK_CREATE`](../../../linux/syscall/BPF_LINK_CREATE.md) syscall command.

## Definition

`#!c int bpf_link_create(int prog_fd, int target_fd, enum bpf_attach_type attach_type, const struct bpf_link_create_opts *opts);`

**Parameters**

- `prog_fd`: file descriptor of the program to attach
- `target_fd`: file descriptor of the target object
- `attach_type`: type of the attachment
- `opts`: options for configuring the attachment

**Return**

`>0`, file descriptor of the created link; negative error code, otherwise

### `struct bpf_link_create_opts`

```c
struct bpf_link_create_opts {
	size_t sz; /* size of this struct for forward/backward compatibility */
	__u32 flags;
	union bpf_iter_link_info *iter_info;
	__u32 iter_info_len;
	__u32 target_btf_id;
	union {
		struct {
			__u64 bpf_cookie;
		} perf_event;
		struct {
			__u32 flags;
			__u32 cnt;
			const char **syms;
			const unsigned long *addrs;
			const __u64 *cookies;
		} kprobe_multi;
		struct {
			__u32 flags;
			__u32 cnt;
			const char *path;
			const unsigned long *offsets;
			const unsigned long *ref_ctr_offsets;
			const __u64 *cookies;
			__u32 pid;
		} uprobe_multi;
		struct {
			__u64 cookie;
		} tracing;
		struct {
			__u32 pf;
			__u32 hooknum;
			__s32 priority;
			__u32 flags;
		} netfilter;
		struct {
			__u32 relative_fd;
			__u32 relative_id;
			__u64 expected_revision;
		} tcx;
		struct {
			__u32 relative_fd;
			__u32 relative_id;
			__u64 expected_revision;
		} netkit;
	};
	size_t :0;
};
```

## Usage

This function should only be used for specific program types that need to be attached via the `BPF_LINK_CREATE` syscall command and you need specific control over this process. In most cases, the [`bpf_program__attach`](bpf_program__attach.md) or specific `bpf_program__attach_*` functions should be used instead.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
