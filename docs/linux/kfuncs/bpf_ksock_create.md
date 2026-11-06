---
title: "KFunc 'bpf_ksock_create'"
description: "This page documents the 'bpf_ksock_create' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_ksock_create`

<!-- [FEATURE_TAG](bpf_ksock_create) -->
[:octicons-tag-24: v7.3](https://github.com/torvalds/linux/commit/7ae4eb14c5f9d9bf0e0feabeab206151b1280512)
<!-- [/FEATURE_TAG] -->

Create a BPF kernel socket.

## Definition

Allocates and creates a kernel socket, wrapped in a reference counted `struct bpf_ksock` context. The context returned by this function must either be stored in a map as a kptr, or freed with [`bpf_ksock_release`](bpf_ksock_release.md).

The socket is created via the same code path as a [`socket(2)`](https://man7.org/linux/man-pages/man2/socket.2.html) call made by the current task. So the socket is created in the network namespace of the task invoking the BPF program, it is created with the credentials and cGroup of that task, and is subject to the same LSM and cGroup policies as a socket created by that task from userspace.

`opts`: Pointer to a `struct bpf_ksock_create_opts` containing the socket parameters.

`opts__sz`: Size of the `opts` struct, must be `sizeof(struct bpf_ksock_create_opts)`.

`err__uninit`: Pointer to an integer in which an error code is stored when `NULL` is returned, or `0` on success. Does not need to be initialized, see [`__uninit` annotation](../concepts/kfuncs.md#__uninit-annotation).

```c
struct bpf_ksock_create_opts {
	__u8 family;
	__u8 type;
	__u8 protocol;
	__u8 reserved;
};
```

`family`: The address family, `AF_INET` or `AF_INET6`.

`type`: The socket type, only `SOCK_DGRAM` is supported for now.

`protocol`: The protocol number, only `IPPROTO_UDP` or `0` (the default protocol for the given type) are supported for now.

`reserved`: Must be zero, reserved for future use.

**Returns**

Returns a pointer to a new BPF kernel socket context on success, or `NULL` on failure. On failure, `err__uninit` is set to one of the following errors:

- `-EOPNOTSUPP`: Not called from the context of a user task, for example from a [BPF work-queue](bpf_wq_start.md) callback or a kernel thread.
- `-EINVAL`: `opts` is `NULL`, `opts__sz` is incorrect, or `reserved` is not zero.
- `-EAFNOSUPPORT`: `family` is not `AF_INET` or `AF_INET6`.
- `-EPROTONOSUPPORT`: `type` is not `SOCK_DGRAM`, or `protocol` is not `IPPROTO_UDP` or `0`.
- `-ENOMEM`: Not enough memory to allocate the context.
- Any other error returned by socket creation, for example when an LSM denies the creation of the socket.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct bpf_ksock *bpf_ksock_create(const struct bpf_ksock_create_opts *opts, u32 opts__sz, int *err__uninit)`

!!! note
	This kfunc returns a pointer to a refcounted object. The verifier will then ensure that the pointer to the object 
	is eventually released using a release kfunc, or transferred to a map using a referenced kptr 
	(by invoking [`bpf_kptr_xchg`](../helper-function/bpf_kptr_xchg.md)). If not, the verifier fails the 
	loading of the BPF program until no lingering references remain in all possible explored states of the program.

!!! note
	The pointer returned by the kfunc may be NULL. Hence, it forces the user to do a NULL check on the pointer returned 
	from the kfunc before making use of it (dereferencing or passing to another helper).

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).
<!-- [/KFUNC_DEF] -->

## Usage

The `bpf_ksock_*` kfuncs allow BPF programs to create UDP sockets and send data over them. The main use case is to allow BPF programs, for example LSM programs, to emit telemetry or security alerts over the network, independently of a userspace agent. Data can still be emitted after the userspace agent that loaded the programs has stopped.

This kfunc is used to create such a socket. Since socket creation may sleep, this must be done in a [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md) program, which is typically executed by userspace via the [`BPF_PROG_TEST_RUN`](../syscall/BPF_PROG_TEST_RUN.md) syscall command. The socket will then be created in the network namespace, and with the credentials, of the userspace process running the program.

The typical lifecycle looks like this:

1. A syscall program creates a socket with this kfunc, connects it to a remote address with [`bpf_ksock_connect`](bpf_ksock_connect.md) and stores it in a map as a kptr using [`bpf_kptr_xchg`](../helper-function/bpf_kptr_xchg.md).
2. Other programs, such as LSM programs, look up the kptr in the map, acquire a reference to it with [`bpf_ksock_acquire`](bpf_ksock_acquire.md), send data with [`bpf_ksock_send`](bpf_ksock_send.md) and release their reference with [`bpf_ksock_release`](bpf_ksock_release.md).
3. When the kptr is removed from the map, or the map itself is freed, the reference held by the map is released. When the last reference is released, the socket is closed.

The receive buffer of the created socket is set to the minimum size and locked, since the socket is only meant to be used for sending data.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

This example creates and connects a UDP socket in a syscall program and stores it in a map. Every time the process with PID `target_pid` binds a socket, the sleepable LSM program sends a message over the socket.

Example based on [`tools/testing/selftests/bpf/progs/ksock_common.h`](https://github.com/torvalds/linux/blob/c03c1a14507dfaf8b856c2608f9ac216021fc222/tools/testing/selftests/bpf/progs/ksock_common.h) and [`tools/testing/selftests/bpf/progs/ksock_lsm.c`](https://github.com/torvalds/linux/blob/c03c1a14507dfaf8b856c2608f9ac216021fc222/tools/testing/selftests/bpf/progs/ksock_lsm.c)

!!! example
	```c
	// SPDX-License-Identifier: GPL-2.0
	/* Copyright (c) 2026 Isovalent */

	#include "vmlinux.h"
	#include <bpf/bpf_helpers.h>
	#include <bpf/bpf_tracing.h>
	#include <bpf/bpf_endian.h>

	#define AF_INET     2
	#define SOCK_DGRAM  2
	#define IPPROTO_UDP 17
	#define ENOENT      2
	#define EEXIST      17

	struct bpf_ksock *bpf_ksock_create(const struct bpf_ksock_create_opts *opts,
					   u32 opts__sz, int *err__uninit) __ksym;
	int bpf_ksock_connect(struct bpf_ksock *ks, const union bpf_ksock_addr *addr,
			      u32 addr__sz) __ksym;
	struct bpf_ksock *bpf_ksock_acquire(struct bpf_ksock *ks) __ksym;
	void bpf_ksock_release(struct bpf_ksock *ks) __ksym;
	int bpf_ksock_send(struct bpf_ksock *ks, const void *data, u32 data__sz) __ksym;
	void bpf_rcu_read_lock(void) __ksym;
	void bpf_rcu_read_unlock(void) __ksym;

	struct __ksock_ctx_value {
		struct bpf_ksock __kptr * ctx;
	};

	struct {
		__uint(type, BPF_MAP_TYPE_ARRAY);
		__type(key, int);
		__type(value, struct __ksock_ctx_value);
		__uint(max_entries, 1);
	} __ksock_ctx_map SEC(".maps");

	static inline struct __ksock_ctx_value *ksock_ctx_value_lookup(void)
	{
		u32 key = 0;

		return bpf_map_lookup_elem(&__ksock_ctx_map, &key);
	}

	static inline struct bpf_ksock *ksock_ctx_get(void)
	{
		struct __ksock_ctx_value *v;
		struct bpf_ksock *ks = NULL, *tmp;

		v = ksock_ctx_value_lookup();
		if (!v)
			return NULL;

		bpf_rcu_read_lock();
		tmp = v->ctx;
		if (tmp)
			ks = bpf_ksock_acquire(tmp);
		bpf_rcu_read_unlock();

		return ks;
	}

	static inline int ksock_ctx_insert(struct bpf_ksock *ctx)
	{
		struct __ksock_ctx_value *v;
		struct bpf_ksock *old;

		v = ksock_ctx_value_lookup();
		if (!v) {
			bpf_ksock_release(ctx);
			return -ENOENT;
		}

		old = bpf_kptr_xchg(&v->ctx, ctx);
		if (old) {
			bpf_ksock_release(old);
			return -EEXIST;
		}

		return 0;
	}

	char send_data[32] = "hello from bpf ksock";

	__be32 ipv4_remote;
	__u16 remote_port;
	int target_pid;
	int send_ret = -1;

	SEC("syscall")
	int ksock_setup(void *ctx)
	{
		struct bpf_ksock_create_opts create_opts = {};
		union bpf_ksock_addr addr = {};
		struct bpf_ksock *ks;
		int err = 0;

		create_opts.family = AF_INET;
		create_opts.type = SOCK_DGRAM;
		create_opts.protocol = IPPROTO_UDP;

		ks = bpf_ksock_create(&create_opts, sizeof(create_opts), &err);
		if (!ks)
			return err;

		addr.sin.sin_family = AF_INET;
		addr.sin.sin_port = bpf_htons(remote_port);
		addr.sin.sin_addr.s_addr = ipv4_remote;

		err = bpf_ksock_connect(ks, &addr, sizeof(addr));
		if (err) {
			bpf_ksock_release(ks);
			return err;
		}

		err = ksock_ctx_insert(ks);
		if (err && err != -EEXIST)
			return err;
		return 0;
	}

	SEC("lsm.s/socket_bind")
	int BPF_PROG(ksock_socket_bind, struct socket *sock, struct sockaddr *address,
		     int addrlen, int ret)
	{
		struct bpf_ksock *ks;
		u32 pid = bpf_get_current_pid_tgid() >> 32;

		if (ret || pid != target_pid)
			return ret;

		ks = ksock_ctx_get();
		if (!ks) {
			send_ret = -ENOENT;
			return ret;
		}

		send_ret = bpf_ksock_send(ks, send_data, sizeof(send_data));
		bpf_ksock_release(ks);

		return ret;
	}

	char __license[] SEC("license") = "GPL";
	```
