---
title: "KFunc 'bpf_crypto_ctx_create'"
description: "This page documents the 'bpf_crypto_ctx_create' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_crypto_ctx_create`

<!-- [FEATURE_TAG](bpf_crypto_ctx_create) -->
[:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/3e1c6f35409f9e447bf37f64840f5b65576bfb78)
<!-- [/FEATURE_TAG] -->

Create a mutable BPF crypto context.

## Definition

Allocates a crypto context that can be used, acquired, and released by a BPF program. The crypto context returned by this function must either be embedded in a map as a kptr, or freed with [`bpf_crypto_ctx_release`](bpf_crypto_ctx_release.md). As crypto API functions use GFP_KERNEL allocations, this function can only be used in sleepable BPF programs.

`params`: pointer to struct bpf_crypto_params which contains all the details needed to initialise crypto context.

`params__sz`: size of steuct bpf_crypto_params usef by bpf program

`err`: integer to store error code when NULL is returned.

**Returns**

Returns an allocated crypto context on success, may return NULL if no memory is available.

<!-- [KFUNC_DEF] -->
`#!c struct bpf_crypto_ctx *bpf_crypto_ctx_create(const struct bpf_crypto_params *params, u32 params__sz, int *err)`

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

This kfunc is used to allocate a new BPF crypto context which can then be used in `bpf_crypto_encrypt` and `bpf_crypto_decrypt` to encrypt or decrypt network packets. The creation allocates memory and thus may sleep, so this must be done outside of packet processing context in a syscall program.

The created context can be stored and shared with network programs via a map containing a kernel pointer.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_SYSCALL](../program-type/BPF_PROG_TYPE_SYSCALL.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example
	```c
	/* Copyright (c) 2024 Meta Platforms, Inc. and affiliates. */

	#include "vmlinux.h"
	#include "bpf_tracing_net.h"
	#include <bpf/bpf_helpers.h>
	#include <bpf/bpf_endian.h>
	#include <bpf/bpf_tracing.h>
	#include "bpf_misc.h"
	#include "bpf_kfuncs.h"

	struct bpf_crypto_ctx *bpf_crypto_ctx_create(const struct bpf_crypto_params *params,
							u32 params__sz, int *err) __ksym;
	struct bpf_crypto_ctx *bpf_crypto_ctx_acquire(struct bpf_crypto_ctx *ctx) __ksym;
	void bpf_crypto_ctx_release(struct bpf_crypto_ctx *ctx) __ksym;

	struct __crypto_ctx_value {
		struct bpf_crypto_ctx __kptr * ctx;
	};

	struct array_map {
		__uint(type, BPF_MAP_TYPE_ARRAY);
		__type(key, int);
		__type(value, struct __crypto_ctx_value);
		__uint(max_entries, 1);
	} __crypto_ctx_map SEC(".maps");

	static inline int crypto_ctx_insert(struct bpf_crypto_ctx *ctx)
	{
		struct __crypto_ctx_value local, *v;
		struct bpf_crypto_ctx *old;
		u32 key = 0;
		int err;

		local.ctx = NULL;
		err = bpf_map_update_elem(&__crypto_ctx_map, &key, &local, 0);
		if (err) {
			bpf_crypto_ctx_release(ctx);
			return err;
		}

		v = bpf_map_lookup_elem(&__crypto_ctx_map, &key);
		if (!v) {
			bpf_crypto_ctx_release(ctx);
			return -ENOENT;
		}

		old = bpf_kptr_xchg(&v->ctx, ctx);
		if (old) {
			bpf_crypto_ctx_release(old);
			return -EEXIST;
		}

		return 0;
	}

	char cipher[128] = {};
	u32 key_len, authsize;
	u8 key[256] = {};
	int status;

	SEC("syscall")
	int crypto_setup(void *args)
	{
		struct bpf_crypto_ctx *cctx;
		struct bpf_crypto_params params = {
			.type = "skcipher",
			.key_len = key_len,
			.authsize = authsize,
		};
		int err = 0;

		status = 0;

		if (!cipher[0] || !key_len || key_len > 256) {
			status = -EINVAL;
			return 0;
		}

		__builtin_memcpy(&params.algo, cipher, sizeof(cipher));
		__builtin_memcpy(&params.key, key, sizeof(key));
		cctx = bpf_crypto_ctx_create(&params, sizeof(params), &err);

		if (!cctx) {
			status = err;
			return 0;
		}

		err = crypto_ctx_insert(cctx);
		if (err && err != -EEXIST)
			status = err;

		return 0;
	}
	```
