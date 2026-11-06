---
title: "KFunc 'bpf_ksock_send'"
description: "This page documents the 'bpf_ksock_send' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_ksock_send`

<!-- [FEATURE_TAG](bpf_ksock_send) -->
[:octicons-tag-24: v7.3](https://github.com/torvalds/linux/commit/7ae4eb14c5f9d9bf0e0feabeab206151b1280512)
<!-- [/FEATURE_TAG] -->

Send data through a BPF kernel socket.

## Definition

Sends data on a connected BPF kernel socket. Sending is best-effort and non-blocking (`MSG_DONTWAIT`), if the data can't be sent immediately an error is returned instead of waiting.

The data is sent via the same code path as a [`sendmsg(2)`](https://man7.org/linux/man-pages/man2/sendmsg.2.html) call made by the current task, so the `socket_sendmsg` LSM hook is invoked like it would be for a userspace socket.

`ks`: The BPF kernel socket context. Must be an acquired reference, see [`bpf_ksock_acquire`](bpf_ksock_acquire.md).

`data`: Pointer to the data to send.

`data__sz`: Size of the data to send.

**Returns**

Returns the number of bytes sent on success, or a negative error code on failure:

- `-EOPNOTSUPP`: Not called from the context of a user task, for example from a [BPF work-queue](bpf_wq_start.md) callback or a kernel thread.
- `-EAGAIN`: The data can't be sent without blocking, for example because the send buffer is full.
- `-EDESTADDRREQ`: The socket has not been connected with [`bpf_ksock_connect`](bpf_ksock_connect.md).
- Any other error returned by the send operation, for example when an LSM denies the send or the data is too large to fit in a single datagram (`-EMSGSIZE`).

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_ksock_send(struct bpf_ksock *ks, const void *data, u32 data__sz)`

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).
<!-- [/KFUNC_DEF] -->

## Usage

This kfunc is used to send data, such as telemetry or security alerts, over a BPF kernel socket previously created with [`bpf_ksock_create`](bpf_ksock_create.md) and connected with [`bpf_ksock_connect`](bpf_ksock_connect.md). The main use case is to allow LSM programs to emit events over the network without depending on a userspace agent to read them from a ring buffer and forward them.

Typically, the socket is stored in a map as a kptr by a syscall program. The LSM program then acquires a reference with [`bpf_ksock_acquire`](bpf_ksock_acquire.md), sends the data and releases the reference with [`bpf_ksock_release`](bpf_ksock_release.md).

Since this kfunc may sleep, LSM programs must be loaded as sleepable (`SEC("lsm.s/...")`) to use it.

!!! warning
	This kfunc can't be used by LSM programs attached to the `socket_sendmsg` hook, since this kfunc invokes that hook itself, which could cause infinite recursion. The verifier rejects such programs with the message `calling kernel function bpf_ksock_send is not allowed`. This restriction also applies to [`bpf_ksock_acquire`](bpf_ksock_acquire.md) and [`bpf_ksock_release`](bpf_ksock_release.md).

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

See [`bpf_ksock_create`](bpf_ksock_create.md#example) for an example.
