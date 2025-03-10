---
title: "Libbpf userspace function 'bpf_program__attach_trace_opts'"
description: "This page documents the 'bpf_program__attach_trace_opts' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_trace_opts`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)
<!-- [/LIBBPF_TAG] -->

Attach a [`BPF_PROG_TYPE_TRACING`](../../../linux/program-type/BPF_PROG_TYPE_TRACING.md) program. Like [`bpf_program__attach_trace`](bpf_program__attach_trace.md), but with additional options.

## Definition

`#!c struct bpf_link * bpf_program__attach_trace_opts(const struct bpf_program *prog, const struct bpf_trace_opts *opts);`

**Parameters**

- `prog`: BPF program to attach
- `opts`: Trace options

**Return**

Reference to the newly created BPF link; or `NULL` is returned on error, error code is stored in [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html)

### `struct bpf_trace_opts`

```c
struct bpf_trace_opts {
	/* size of this struct, for forward/backward compatibility */
	size_t sz;
	/* custom user-provided value fetchable through bpf_get_attach_cookie() */
	__u64 cookie;
};
```

#### `cookie`

[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/commit/aec48fffeeb00f93d5ed7f87c5e81412647c363c)

Custom user-provided value fetchable through [`bpf_get_attach_cookie`](../../../linux/helper-function/bpf_get_attach_cookie.md). This allows you to write one program, load it once, and then attach it to multiple perf events with different `bpf_cookie` values, allowing the program to detect which event it is attached to.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
