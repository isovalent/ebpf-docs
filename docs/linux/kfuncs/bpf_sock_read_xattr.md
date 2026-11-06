---
title: "KFunc 'bpf_sock_read_xattr'"
description: "This page documents the 'bpf_sock_read_xattr' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_sock_read_xattr`

<!-- [FEATURE_TAG](bpf_sock_read_xattr) -->
[:octicons-tag-24: v7.3](https://github.com/torvalds/linux/commit/b4e124d16855213409f5dfa6aa18b81cd00fdbba)
<!-- [/FEATURE_TAG] -->

Read an extended attribute of a socket.

## Definition

Get the extended attribute (xattr) `name__str` of the inode of socket `sock` in the socket file system and store its value in `value_p`. For security reasons, only `name__str` with prefix `user.` is allowed.

The read never sleeps and does not take any locks, the value is looked up in the in-memory xattr store of the inode under RCU protection.

**Parameters**

`sock`: The socket to get the xattr from. Must be a trusted or RCU protected pointer.

`name__str`: The name of the xattr, must be a constant string starting with `user.`.

`value_p`: The output buffer for the xattr value, a [dynamic pointer](../concepts/dynptrs.md).

**Returns**

The length of the xattr value on success, or a negative error code on failure:

- `-EPERM`: `name__str` does not start with `user.`.
- `-EINVAL`: `value_p` is not a valid, writable dynamic pointer.
- `-EOPNOTSUPP`: `sock` is not a socket in the socket file system.
- `-ENODATA`: The socket does not have an xattr named `name__str`.
- `-ERANGE`: The xattr value does not fit in `value_p`.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_sock_read_xattr(struct socket *sock, const char *name__str, struct bpf_dynptr *value_p)`
<!-- [/KFUNC_DEF] -->

## Usage

This kfunc allows a BPF LSM program to read labels which userspace attached to a socket. Processes can set `user.*` extended attributes on sockets using [`fsetxattr(2)`](https://man7.org/linux/man-pages/man2/fsetxattr.2.html) on the socket file descriptor. A BPF LSM program can then read these labels in socket LSM hooks, which only provide a `struct socket`, without having to acquire a reference to a file and perform a sleeping operation. For example, `systemd` uses `user.*` extended attributes to tag sockets, which a BPF program can read back to implement policies.

This kfunc mirrors [`bpf_cgroup_read_xattr`](bpf_cgroup_read_xattr.md), including the restriction to the `user.*` namespace. It only works for sockets in the socket file system, sockets that are file system objects, such as bound Unix sockets, should be read with [`bpf_get_file_xattr`](bpf_get_file_xattr.md) or [`bpf_get_dentry_xattr`](bpf_get_dentry_xattr.md) instead.

Since the read never sleeps, this kfunc can be used from both sleepable and non-sleepable LSM programs.

!!! note
	A `struct socket` is only guaranteed to live in the socket file system when it is passed to an LSM socket hook. So while the kfunc is part of a kfunc set registered for multiple program types, the verifier only allows calling it from [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md) programs (and `binfmt_misc` struct ops programs).

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

This example reads the `user.bpf_test` xattr of a socket when it is being connected by the process with PID `monitored_pid`.

Example based on [`tools/testing/selftests/bpf/progs/sock_read_xattr.c`](https://github.com/torvalds/linux/blob/c03c1a14507dfaf8b856c2608f9ac216021fc222/tools/testing/selftests/bpf/progs/sock_read_xattr.c)

!!! example
	```c
	// SPDX-License-Identifier: GPL-2.0
	/* Copyright (c) 2026 Christian Brauner */

	#include <vmlinux.h>
	#include <bpf/bpf_tracing.h>
	#include <bpf/bpf_helpers.h>

	extern int bpf_sock_read_xattr(struct socket *sock, const char *name__str,
				       struct bpf_dynptr *value_p) __weak __ksym;

	char _license[] SEC("license") = "GPL";

	char value[16];
	int read_ret = -1;
	__u32 monitored_pid = 0;

	SEC("lsm/socket_connect")
	int BPF_PROG(read_sock_xattr, struct socket *sock)
	{
		struct bpf_dynptr value_ptr;
		__u32 pid = bpf_get_current_pid_tgid() >> 32;

		if (pid != monitored_pid)
			return 0;

		bpf_dynptr_from_mem(value, sizeof(value), 0, &value_ptr);
		read_ret = bpf_sock_read_xattr(sock, "user.bpf_test", &value_ptr);
		return 0;
	}
	```
