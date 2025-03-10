---
title: "Libbpf userspace function 'bpf_program__attach_raw_tracepoint_opts'"
description: "This page documents the 'bpf_program__attach_raw_tracepoint_opts' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_raw_tracepoint_opts`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.4.0](https://github.com/libbpf/libbpf/releases/tag/v1.4.0)
<!-- [/LIBBPF_TAG] -->


Attach a [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../../../linux/program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md) program. Like [`bpf_program__attach_raw_tracepoint`](bpf_program__attach_raw_tracepoint.md), but with additional options.

## Definition

`#!c struct bpf_link * bpf_program__attach_raw_tracepoint_opts(const struct bpf_program *prog, const char *tp_name, struct bpf_raw_tracepoint_opts *opts);`

**Parameters**

- `prog`: BPF program to attach
- `tp_name`: Tracepoint name
- `opts`: Tracepoint options

**Return**

Reference to the newly created BPF link; or `NULL` is returned on error, error code is stored in [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html)

### `struct bpf_raw_tracepoint_opts`

```c
struct bpf_raw_tracepoint_opts {
	size_t sz; /* size of this struct for forward/backward compatibility */
	__u64 cookie;
	size_t :0;
};
```

#### `cookie`

[:octicons-tag-24: 1.4.0](https://github.com/libbpf/libbpf/commit/f5828cc3520f12ceed531b11a551bafb65b36379)

Custom user-provided value fetchable through [`bpf_get_attach_cookie`](../../../linux/helper-function/bpf_get_attach_cookie.md). This allows you to write one program, load it once, and then attach it to multiple perf events with different `bpf_cookie` values, allowing the program to detect which event it is attached to.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
