---
title: "KFunc 'bpf_ksock_connect'"
description: "This page documents the 'bpf_ksock_connect' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_ksock_connect`

<!-- [FEATURE_TAG](bpf_ksock_connect) -->
[:octicons-tag-24: v7.3](https://github.com/torvalds/linux/commit/7ae4eb14c5f9d9bf0e0feabeab206151b1280512)
<!-- [/FEATURE_TAG] -->

Connect a BPF kernel socket to a remote address.

## Definition

Connects the socket of a BPF kernel socket context to the specified remote address and port.

The connect goes through the same code path as a [`connect(2)`](https://man7.org/linux/man-pages/man2/connect.2.html) call made by the current task, so the `socket_connect` LSM hook and [cGroup socket address](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md) connect programs are invoked like they would be for a userspace socket.

`ks`: The BPF kernel socket context, as returned by [`bpf_ksock_create`](bpf_ksock_create.md).

`addr`: Pointer to a `union bpf_ksock_addr` containing an IPv4 or IPv6 socket address.

`addr__sz`: Size of the `addr` union, must be `sizeof(union bpf_ksock_addr)`.

```c
union bpf_ksock_addr {
	struct sockaddr_in sin;
	struct sockaddr_in6 sin6;
};
```

`sin`: IPv4 socket address. Used when `sin.sin_family` is `AF_INET`.

`sin6`: IPv6 socket address. Used when `sin6.sin6_family` is `AF_INET6`.

The address and port are expected to be in network byte order, just like when using `struct sockaddr_in` or `struct sockaddr_in6` from userspace.

**Returns**

Returns `0` on success, or a negative error code on failure:

- `-EOPNOTSUPP`: Not called from the context of a user task, for example from a [BPF work-queue](bpf_wq_start.md) callback or a kernel thread.
- `-EINVAL`: `addr` is `NULL` or `addr__sz` is incorrect.
- `-EAFNOSUPPORT`: The address family is not `AF_INET` or `AF_INET6`.
- Any other error returned by the connect operation, for example when an LSM denies the connect.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_ksock_connect(struct bpf_ksock *ks, const union bpf_ksock_addr *addr, u32 addr__sz)`

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).
<!-- [/KFUNC_DEF] -->

## Usage

This kfunc is used to set the destination of a BPF kernel socket created with [`bpf_ksock_create`](bpf_ksock_create.md), before storing it in a map for use by other programs. Since only UDP sockets are supported for now, connecting does not involve any network traffic, it sets the default destination for data sent with [`bpf_ksock_send`](bpf_ksock_send.md).

A socket must be connected before data can be sent with [`bpf_ksock_send`](bpf_ksock_send.md), since that kfunc does not take a destination address.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

See [`bpf_ksock_create`](bpf_ksock_create.md#example) for an example.
