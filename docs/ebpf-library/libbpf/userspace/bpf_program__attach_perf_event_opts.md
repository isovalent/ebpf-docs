---
title: "Libbpf userspace function 'bpf_program__attach_perf_event_opts'"
description: "This page documents the 'bpf_program__attach_perf_event_opts' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_perf_event_opts`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.5.0](https://github.com/libbpf/libbpf/releases/tag/v0.5.0)
<!-- [/LIBBPF_TAG] -->

Attach a [`BPF_PROG_TYPE_PERF_EVENT`](../../../linux/program-type/BPF_PROG_TYPE_PERF_EVENT.md) program to a perf event.

## Definition

`#!c struct bpf_link * bpf_program__attach_perf_event_opts(const struct bpf_program *prog, int pfd, const struct bpf_perf_event_opts *opts);`

**Parameters**

- `prog`: BPF program to attach
- `pfd`: File descriptor of the perf event to attach to
- `opts`: Additional options

### `struct bpf_perf_event_opts`

```c
struct bpf_perf_event_opts {
	/* size of this struct, for forward/backward compatibility */
	size_t sz;
	__u64 bpf_cookie;
	bool force_ioctl_attach;
	size_t :0;
};
```

#### `bpf_cookie`

[:octicons-tag-24: 0.5.0](https://github.com/libbpf/libbpf/releases/tag/v0.5.0)

Custom user-provided value fetchable through [`bpf_get_attach_cookie`](../../../linux/helper-function/bpf_get_attach_cookie.md). This allows you to write one program, load it once, and then attach it to multiple perf events with different `bpf_cookie` values, allowing the program to detect which event it is attached to.

#### `force_ioctl_attach`

[:octicons-tag-24: 1.2.0](https://github.com/libbpf/libbpf/releases/tag/v1.2.0)

Don't use BPF link when attach BPF program. BPF links are the newer way to attach BPF programs to various kernel hooks. By default, libbpf will attempt to use the newer links and fallback to the older `ioctl` method if the kernel doesn't support links. This option allows you to force the use of the older `ioctl` method.

## Usage

The `pfd` is obtained by first creating a perf event using the [`perf_event_open`](https://man7.org/linux/man-pages/man2/perf_event_open.2.html) syscall.

This function calls [`bpf_program__attach_perf_event_opts`](bpf_program__attach_perf_event_opts.md) with default values.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
