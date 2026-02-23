---
title: "Syscall command 'BPF_PROG_DETACH'"
description: "This page documents the 'BPF_PROG_DETACH' eBPF syscall command, including its definition, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_PROG_DETACH` command

<!-- [FEATURE_TAG](BPF_PROG_DETACH) -->
[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/f4324551489e8781d838f941b7aee4208e52e8bf)
<!-- [/FEATURE_TAG] -->

This syscall detaches programs that were previously attached with [`BPF_PROG_ATTACH`](BPF_PROG_ATTACH.md)

## Return value

This command will return a zero on success or an error number (negative integer) if something went wrong.

## Attributes

### `target_fd`

The file descriptor of the resource to detach the program from. The type of file descriptor changes per program type.

### `target_ifindex`

The network interface index of the network device to detach the program from.

### `attach_bpf_fd`

The file descriptor of the BPF program to detach from the `target_fd`/`target_ifindex`.

### `attach_flags`

Any flags relevant to attaching.

### `expected_revision`

The expected revision of the collection of programs, in the case of attach points that support multiple programs being attached at the same time.


