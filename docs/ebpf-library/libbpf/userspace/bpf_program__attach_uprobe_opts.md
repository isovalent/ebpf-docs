---
title: "Libbpf userspace function 'bpf_program__attach_uprobe_opts'"
description: "This page documents the 'bpf_program__attach_uprobe_opts' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_uprobe_opts`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.5.0](https://github.com/libbpf/libbpf/releases/tag/v0.5.0)
<!-- [/LIBBPF_TAG] -->

Is just like [`bpf_program__attach_uprobe`](bpf_program__attach_uprobe.md) except with a options struct for various configurations.

## Definition

`#!c struct bpf_link * bpf_program__attach_uprobe_opts(const struct bpf_program *prog, pid_t pid, const char *binary_path, size_t func_offset, const struct bpf_uprobe_opts *opts);`

**Parameters**

- `prog`: BPF program to attach
- `pid`: Process ID to attach the uprobe to, 0 for self (own process), `-1` for all processes
- `binary_path`: Path to binary that contains the function symbol
- `func_offset`: Offset within the binary of the function symbol
- `opts`: Options for altering program attachment

**Return**

Reference to the newly created BPF link; or `NULL` is returned on error, error code is stored in [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html)

### `struct bpf_uprobe_opts`

```c
struct bpf_uprobe_opts {
	/* size of this struct, for forward/backward compatibility */
	size_t sz;
	size_t ref_ctr_offset;
	__u64 bpf_cookie;
	bool retprobe;
	const char *func_name;
	enum probe_attach_mode attach_mode;
	size_t :0;
};
```

#### `ref_ctr_offset`

[:octicons-tag-24: 0.5.0](https://github.com/libbpf/libbpf/commit/6d67d5314399822be224da81858c786526435c63)

Offset of kernel reference counted USDT semaphore, added in a6ca88b241d5 ("trace_uprobe: support reference counter in <nospell>fd-based</nospell> uprobe")

#### `bpf_cookie`

[:octicons-tag-24: 0.5.0](https://github.com/libbpf/libbpf/commit/91259bc676ae64bb376cff666055d09640773737)

Custom user-provided value fetchable through [`bpf_get_attach_cookie`](../../../linux/helper-function/bpf_get_attach_cookie.md). This allows you to write one program, load it once, and then attach it to multiple perf events with different `bpf_cookie` values, allowing the program to detect which event it is attached to.

#### `retprobe`

[:octicons-tag-24: 0.5.0](https://github.com/libbpf/libbpf/commit/91259bc676ae64bb376cff666055d09640773737)

uprobe is return probe, invoked at function return time.

#### `func_name`

[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/commit/d112c9ce249bd2999235be92fd3d58c94778499d)

Function name to attach to. Could be an unqualified <nospell>("abc")</nospell> or library-qualified <nospell>"abc@LIBXYZ"</nospell> name.  To specify function entry, `func_name` should be set while `func_offset` argument to `bpf_prog__attach_uprobe_opts` should be `0`. To trace an offset within a function, specify `func_name` and use `func_offset` argument to specify offset within the function. Shared library functions must specify the shared library `binary_path`.

#### `attach_mode`

[:octicons-tag-24: 1.2.0](https://github.com/libbpf/libbpf/commit/cc7177624f04ef6b6f877ec7ed87299c2cad72f7)

The mode to attach uprobe force libbpf to attach uprobe in specific mode, `-ENOTSUP` will be returned if it is not supported by the kernel.

```c
enum probe_attach_mode {
	/* attach probe in latest supported mode by kernel */
	PROBE_ATTACH_MODE_DEFAULT = 0,
	/* attach probe in legacy mode, using debugfs/tracefs */
	PROBE_ATTACH_MODE_LEGACY,
	/* create perf event with perf_event_open() syscall */
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
