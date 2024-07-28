---
title: "KFunc 'bpf_crypto_ctx_acquire'"
description: "This page documents the 'bpf_crypto_ctx_acquire' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_crypto_ctx_acquire`

<!-- [FEATURE_TAG](bpf_crypto_ctx_acquire) -->
[:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/3e1c6f35409f9e447bf37f64840f5b65576bfb78)
<!-- [/FEATURE_TAG] -->

Acquire a reference to a BPF crypto context.

## Definition

Acquires a reference to a BPF crypto context. The context returned by this function must either be embedded in a map as a kptr, or freed with [`bpf_crypto_ctx_release`](bpf_crypto_ctx_release.md).

`ctx`: The BPF crypto context being acquired. The ctx must be a trusted pointer.

**Returns**

Returns `ctx` on success, or `NULL` if a reference could not be acquired.

<!-- [KFUNC_DEF] -->
`#!c struct bpf_crypto_ctx *bpf_crypto_ctx_acquire(struct bpf_crypto_ctx *ctx)`

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

This kfunc can be used to acquire a reference to a BPF crypto context that was previously created using [`bpf_crypto_ctx_create`](bpf_crypto_ctx_create.md). This allows you to add the same context to multiple values in the same map or to multiple maps.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

see [bpf_crypto_ctx_create](bpf_crypto_ctx_create.md#example) for an example

