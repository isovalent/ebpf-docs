---
title: "KFunc 'bpf_crypto_ctx_release'"
description: "This page documents the 'bpf_crypto_ctx_release' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_crypto_ctx_release`

<!-- [FEATURE_TAG](bpf_crypto_ctx_release) -->
[:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/3e1c6f35409f9e447bf37f64840f5b65576bfb78)
<!-- [/FEATURE_TAG] -->

Release a previously acquired BPF crypto context.

## Definition

Releases a previously acquired reference to a BPF crypto context. When the final reference of the BPF crypto context has been released, its memory will be released.

`ctx`: The crypto context being released.

<!-- [KFUNC_DEF] -->
`#!c void bpf_crypto_ctx_release(struct bpf_crypto_ctx *ctx)`

!!! note
	This kfunc releases the pointer passed in to it. There can be only one referenced pointer that can be passed in. 
	All copies of the pointer being released are invalidated as a result of invoking this kfunc.
<!-- [/KFUNC_DEF] -->

## Usage

This kfunc is used to release a reference held on a BPF crypto context previously acquired using [`bpf_crypto_ctx_acquire`](bpf_crypto_ctx_acquire.md) or [`bpf_crypto_ctx_create`](bpf_crypto_ctx_create.md).

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

see [bpf_crypto_ctx_create](bpf_crypto_ctx_create.md#example) for an example

