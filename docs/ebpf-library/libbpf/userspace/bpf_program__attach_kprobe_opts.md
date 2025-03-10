---
title: "Libbpf userspace function 'bpf_program__attach_kprobe_opts'"
description: "This page documents the 'bpf_program__attach_kprobe_opts' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_kprobe_opts`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.5.0](https://github.com/libbpf/libbpf/releases/tag/v0.5.0)
<!-- [/LIBBPF_TAG] -->

Attach a [`BPF_PROG_TYPE_KPROBE`](../../../linux/program-type/BPF_PROG_TYPE_KPROBE.md) program.

## Definition

`#!c struct bpf_link * bpf_program__attach_kprobe_opts(const struct bpf_program *prog, const char *func_name, const struct bpf_kprobe_opts *opts);`

**Parameters**

- `prog`: pointer to the `bpf_program` object.
- `func_name`: name of the kernel function to attach the probe to.
- `opts`: additional options for the kprobe.

### `struct bpf_kprobe_opts`

```c
struct bpf_kprobe_opts {
	/* size of this struct, for forward/backward compatibility */
	size_t sz;
	__u64 bpf_cookie;
	size_t offset;
	bool retprobe;
	/* kprobe attach mode */
	enum probe_attach_mode attach_mode;
	size_t :0;
};
```

#### `bpf_cookie`

[:octicons-tag-24: 0.5.0](https://github.com/libbpf/libbpf/releases/tag/v0.5.0)

Custom user-provided value fetchable through [`bpf_get_attach_cookie`](../../../linux/helper-function/bpf_get_attach_cookie.md). This allows you to write one program, load it once, and then attach it to multiple perf events with different `bpf_cookie` values, allowing the program to detect which event it is attached to.

#### `offset`

[:octicons-tag-24: 0.5.0](https://github.com/libbpf/libbpf/releases/tag/v0.5.0)

Function's offset to install kprobe to. By default, the probe is installed at the function's entry. By you can install it at any CPU instruction in the function by specifying the offset.

#### `retprobe`

[:octicons-tag-24: 0.5.0](https://github.com/libbpf/libbpf/releases/tag/v0.5.0)

Kprobe is return probe.

#### `attach_mode`

[:octicons-tag-24: 1.2.0](https://github.com/libbpf/libbpf/releases/tag/v1.2.0)

The mode to attach kprobe/uprobe. Values are:

```c
enum probe_attach_mode {
	/* attach probe in latest supported mode by kernel */
	PROBE_ATTACH_MODE_DEFAULT = 0,
	/* attach probe in legacy mode, using debugfs/tracefs */
	PROBE_ATTACH_MODE_LEGACY,
	/* create perf event with [perf_event_open](https://man7.org/linux/man-pages/man2/perf_event_open.2.html) syscall */
	PROBE_ATTACH_MODE_PERF,
	/* attach probe with BPF link */
	PROBE_ATTACH_MODE_LINK,
};
```

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
