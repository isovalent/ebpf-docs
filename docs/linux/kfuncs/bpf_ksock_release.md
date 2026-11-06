---
title: "KFunc 'bpf_ksock_release'"
description: "This page documents the 'bpf_ksock_release' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_ksock_release`

<!-- [FEATURE_TAG](bpf_ksock_release) -->
[:octicons-tag-24: v7.3](https://github.com/torvalds/linux/commit/7ae4eb14c5f9d9bf0e0feabeab206151b1280512)
<!-- [/FEATURE_TAG] -->

Release a BPF kernel socket.

## Definition

Releases a reference to a BPF kernel socket context. When the final reference has been released, the socket is closed and the context is freed. Since closing a socket may sleep, this cleanup is deferred to a work-queue after an RCU grace period, so this kfunc itself does not sleep.

`ks`: The BPF kernel socket context to release.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void bpf_ksock_release(struct bpf_ksock *ks)`

!!! note
	This kfunc releases the pointer passed in to it. There can be only one referenced pointer that can be passed in. 
	All copies of the pointer being released are invalidated as a result of invoking this kfunc.
<!-- [/KFUNC_DEF] -->

## Usage

This kfunc is used to release a reference held on a BPF kernel socket context, previously acquired using [`bpf_ksock_create`](bpf_ksock_create.md) or [`bpf_ksock_acquire`](bpf_ksock_acquire.md), or taken out of a map with [`bpf_kptr_xchg`](../helper-function/bpf_kptr_xchg.md).

When a map value containing a BPF kernel socket kptr is deleted, or the map itself is freed, the kernel automatically calls this kfunc to release the reference held by the map.

!!! warning
	This kfunc can't be used by LSM programs attached to the `socket_sendmsg` hook. Since [`bpf_ksock_send`](bpf_ksock_send.md) invokes this hook, such programs could cause infinite recursion. The verifier rejects such programs with the message `calling kernel function bpf_ksock_release is not allowed`.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

See [`bpf_ksock_create`](bpf_ksock_create.md#example) for an example.
