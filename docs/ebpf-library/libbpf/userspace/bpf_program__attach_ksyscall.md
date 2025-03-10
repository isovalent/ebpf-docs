---
title: "Libbpf userspace function 'bpf_program__attach_ksyscall'"
description: "This page documents the 'bpf_program__attach_ksyscall' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_ksyscall`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.0.0](https://github.com/libbpf/libbpf/releases/tag/v1.0.0)
<!-- [/LIBBPF_TAG] -->

Attaches a BPF program to kernel syscall handler of a specified syscall. Optionally it's possible to request to install retprobe that will be triggered at syscall exit. It's also possible to associate BPF cookie (though options).

## Definition

`#!c struct bpf_link * bpf_program__attach_ksyscall(const struct bpf_program *prog, const char *syscall_name, const struct bpf_ksyscall_opts *opts);`

**Parameters**

- `prog`: BPF program to attach
- `syscall_name`: Symbolic name of the syscall (e.g., "bpf")
- `opts`: Additional options (see `struct bpf_ksyscall_opts`)

**Return**

Reference to the newly created BPF link; or `NULL` is returned on error, error code is stored in [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html)

### `struct bpf_ksyscall_opts`

```c
struct bpf_ksyscall_opts {
	/* size of this struct, for forward/backward compatibility */
	size_t sz;
	__u64 bpf_cookie;
	bool retprobe;
	size_t :0;
};
```

#### `bpf_cookie`

[:octicons-tag-24: 1.0.0](https://github.com/libbpf/libbpf/commit/0862e4e54d1139e2f0ba39091244c2c7e6e24c8a)

Custom user-provided value fetchable through [`bpf_get_attach_cookie`](../../../linux/helper-function/bpf_get_attach_cookie.md). This allows you to write one program, load it once, and then attach it to multiple perf events with different `bpf_cookie` values, allowing the program to detect which event it is attached to.

#### `retprobe`

[:octicons-tag-24: 1.0.0](https://github.com/libbpf/libbpf/commit/0862e4e54d1139e2f0ba39091244c2c7e6e24c8a)

Attach as return probe. If set to `true`, the BPF program will be triggered at syscall exit.

## Usage

Libbpf automatically will determine correct full kernel function name, which depending on system architecture and kernel version/configuration could be of the form `__<arch>_sys_<syscall>` or `__se_sys_<syscall>`, and will attach specified program using kprobe/kretprobe mechanism.

[`bpf_program__attach_ksyscall`](bpf_program__attach_ksyscall.md) is an API counterpart of declarative
`SEC("ksyscall/<syscall>")` annotation of BPF programs.

At the moment `SEC("ksyscall")` and `bpf_program__attach_ksyscall()` do not handle all the calling convention quirks for `mmap()`, `clone()` and compat syscalls. It also only attaches to "native" syscall interfaces. If host system supports compat syscalls or defines 32-bit syscalls in 64-bit kernel, such syscall interfaces won't be attached to by libbpf.

These limitations may or may not change in the future. Therefore it is recommended to use `SEC("kprobe")` for these syscalls or if working with compat and 32-bit interfaces is required.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
