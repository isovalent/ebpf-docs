---
title: "KFunc 'bpf_ksock_acquire'"
description: "This page documents the 'bpf_ksock_acquire' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_ksock_acquire`

<!-- [FEATURE_TAG](bpf_ksock_acquire) -->
[:octicons-tag-24: v7.3](https://github.com/torvalds/linux/commit/7ae4eb14c5f9d9bf0e0feabeab206151b1280512)
<!-- [/FEATURE_TAG] -->

Acquire a reference to a BPF kernel socket.

## Definition

Acquires a reference to a BPF kernel socket context. The context returned by this function must either be stored in a map as a kptr, or freed with [`bpf_ksock_release`](bpf_ksock_release.md).

`ks`: The BPF kernel socket context to acquire. Must be a trusted or RCU protected pointer, for example a kptr loaded from a map value while holding the RCU read lock (see [`KF_RCU`](../concepts/kfuncs.md#kf_rcu)).

**Returns**

Returns `ks` on success, or `NULL` if a reference could not be acquired because the last reference to the context has already been released.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct bpf_ksock *bpf_ksock_acquire(struct bpf_ksock *ks)`

!!! note
	This kfunc returns a pointer to a refcounted object. The verifier will then ensure that the pointer to the object 
	is eventually released using a release kfunc, or transferred to a map using a referenced kptr 
	(by invoking [`bpf_kptr_xchg`](../helper-function/bpf_kptr_xchg.md)). If not, the verifier fails the 
	loading of the BPF program until no lingering references remain in all possible explored states of the program.

!!! note
	The pointer returned by the kfunc may be NULL. Hence, it forces the user to do a NULL check on the pointer returned 
	from the kfunc before making use of it (dereferencing or passing to another helper).
<!-- [/KFUNC_DEF] -->

## Usage

This kfunc is used to acquire a reference to a BPF kernel socket context that was previously created with [`bpf_ksock_create`](bpf_ksock_create.md) and stored in a map as a kptr. This allows programs to use the socket stored in the map, for example to send data with [`bpf_ksock_send`](bpf_ksock_send.md), without removing it from the map. It also allows you to add the same context to multiple values in the same map or to multiple maps.

A typical pattern is to take the RCU read lock with [`bpf_rcu_read_lock`](bpf_rcu_read_lock.md), read the kptr from the map value, acquire a reference with this kfunc and then release the RCU read lock with [`bpf_rcu_read_unlock`](bpf_rcu_read_unlock.md). The acquired reference remains valid after the RCU read lock is released.

!!! warning
	This kfunc can't be used by LSM programs attached to the `socket_sendmsg` hook. Since [`bpf_ksock_send`](bpf_ksock_send.md) invokes this hook, such programs could cause infinite recursion. The verifier rejects such programs with the message `calling kernel function bpf_ksock_acquire is not allowed`.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

See [`bpf_ksock_create`](bpf_ksock_create.md#example) for an example.
