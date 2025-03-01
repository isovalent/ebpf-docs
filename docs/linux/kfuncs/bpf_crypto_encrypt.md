---
title: "KFunc 'bpf_crypto_encrypt'"
description: "This page documents the 'bpf_crypto_encrypt' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_crypto_encrypt`

<!-- [FEATURE_TAG](bpf_crypto_encrypt) -->
[:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/3e1c6f35409f9e447bf37f64840f5b65576bfb78)
<!-- [/FEATURE_TAG] -->

Encrypt buffer using configured context and IV provided.

## Definition

Encrypts provided buffer using IV data and the crypto context. Crypto context must be configured.

`ctx`: The crypto context being used. The ctx must be a trusted pointer.

`src`: bpf_dynptr to the plain data. Must be a trusted pointer.

`dst`: bpf_dynptr to buffer where to store the result. Must be a trusted pointer.

`siv`: bpf_dynptr to IV data and state data to be used by decryptor.

**Returns**

Return 0 on success, or a negative error code on failure.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_crypto_encrypt(struct bpf_crypto_ctx *ctx, const struct bpf_dynptr *src, const struct bpf_dynptr *dst, const struct bpf_dynptr *siv__nullable)`
<!-- [/KFUNC_DEF] -->

## Usage

This kfunc allows network programs to encrypt network packets using the kernels cryptographic functions. This requires a cryptographic context which can be created using [`bpf_crypto_ctx_create`](bpf_crypto_ctx_create.md).

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
- [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
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

    int bpf_crypto_encrypt(struct bpf_crypto_ctx *ctx, const struct bpf_dynptr *src,
                const struct bpf_dynptr *dst, const struct bpf_dynptr *iv) __ksym;

    struct __crypto_ctx_value {
        struct bpf_crypto_ctx __kptr * ctx;
    };

    struct array_map {
        __uint(type, BPF_MAP_TYPE_ARRAY);
        __type(key, int);
        __type(value, struct __crypto_ctx_value);
        __uint(max_entries, 1);
    } __crypto_ctx_map SEC(".maps");

    static inline struct __crypto_ctx_value *crypto_ctx_value_lookup(void)
    {
        u32 key = 0;

        return bpf_map_lookup_elem(&__crypto_ctx_map, &key);
    }

    const volatile unsigned int len = 16;
    char dst[256] = {};

    SEC("tc")
    int crypto_encrypt(struct __sk_buff *skb)
    {
        struct __crypto_ctx_value *v;
        struct bpf_crypto_ctx *ctx;
        struct bpf_dynptr psrc, pdst, iv;

        v = crypto_ctx_value_lookup();
        if (!v) {
            status = -ENOENT;
            return 0;
        }

        ctx = v->ctx;
        if (!ctx) {
            status = -ENOENT;
            return 0;
        }

        bpf_dynptr_from_skb(skb, 0, &psrc);
        bpf_dynptr_from_mem(dst, len, 0, &pdst);
        bpf_dynptr_from_mem(dst, 0, 0, &iv);

        status = bpf_crypto_encrypt(ctx, &psrc, &pdst, &iv);
        __sync_add_and_fetch(&hits, 1);

        return 0;
    }

    char __license[] SEC("license") = "GPL";
    ```
