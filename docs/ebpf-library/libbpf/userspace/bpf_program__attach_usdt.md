---
title: "Libbpf userspace function 'bpf_program__attach_usdt'"
description: "This page documents the 'bpf_program__attach_usdt' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_usdt`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)
<!-- [/LIBBPF_TAG] -->

Is just like [`bpf_program__attach_uprobe_opts`](bpf_program__attach_uprobe_opts.md) except it covers USDT (User-space Statically Defined Tracepoint) attachment, instead of attaching to user-space function entry or exit.

## Definition

`#!c struct bpf_link * bpf_program__attach_usdt(const struct bpf_program *prog, pid_t pid, const char *binary_path, const char *usdt_provider, const char *usdt_name, const struct bpf_usdt_opts *opts);`

**Parameters**

- `prog`: BPF program to attach
- `pid`: Process ID to attach the uprobe to, `0` for self (own process), `-1` for all processes
- `binary_path`: Path to binary that contains provided USDT probe
- `usdt_provider`: USDT provider name
- `usdt_name`: USDT probe name
- `opts`: Options for altering program attachment

**Return**

Reference to the newly created BPF link; or `NULL` is returned on error, error code is stored in [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html)

### `struct bpf_usdt_opts`

```c
struct bpf_usdt_opts {
	/* size of this struct, for forward/backward compatibility */
	size_t sz;
	__u64 usdt_cookie;
	size_t :0;
};
```

#### `usdt_cookie`

[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/commit/1b4b798916b3eb0da6a149343fa6b5deabf74517)

Custom user-provided value fetchable through [`bpf_get_attach_cookie`](../../../linux/helper-function/bpf_get_attach_cookie.md). This allows you to write one program, load it once, and then attach it to multiple perf events with different `bpf_cookie` values, allowing the program to detect which event it is attached to.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
