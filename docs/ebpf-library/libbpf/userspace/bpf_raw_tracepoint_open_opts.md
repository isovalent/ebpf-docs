---
title: "Libbpf userspace function 'bpf_raw_tracepoint_open_opts'"
description: "This page documents the 'bpf_raw_tracepoint_open_opts' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_raw_tracepoint_open_opts`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.4.0](https://github.com/libbpf/libbpf/releases/tag/v1.4.0)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_RAW_TRACEPOINT_OPEN`](../../../linux/syscall/BPF_RAW_TRACEPOINT_OPEN.md) syscall command.

## Definition

`#!c int bpf_raw_tracepoint_open_opts(int prog_fd, struct bpf_raw_tp_opts *opts);`

**Parameters**

- `prog_fd`: BPF program file descriptor
- `opts`: options for configuring the raw tracepoint

**Return**

`>0`, file descriptor of the raw tracepoint; negative error code, otherwise

### `struct bpf_raw_tp_opts`

```c
struct bpf_raw_tp_opts {
	size_t sz; /* size of this struct for forward/backward compatibility */
	const char *tp_name;
	__u64 cookie;
	size_t :0;
};
```

## Usage

This function should only be used if you need precise control over the raw tracepoint opening process. In most cases the [`bpf_program__attach_raw_tracepoint`](bpf_program__attach_raw_tracepoint.md) or [`bpf_program__attach_raw_tracepoint_opts`](bpf_program__attach_raw_tracepoint_opts.md) function should be used instead.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
