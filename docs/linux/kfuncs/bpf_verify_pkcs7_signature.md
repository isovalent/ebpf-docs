---
title: "KFunc 'bpf_verify_pkcs7_signature'"
description: "This page documents the 'bpf_verify_pkcs7_signature' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_verify_pkcs7_signature`

<!-- [FEATURE_TAG](bpf_verify_pkcs7_signature) -->
[:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/865b0566d8f1a0c3937e5eb4bd6ba4ef03e7e98c)
<!-- [/FEATURE_TAG] -->

Verify a <nospell>PKCS#7</nospell> signature

## Definition

Verify the <nospell>PKCS#7</nospell> signature `sig_ptr` against the supplied `data_ptr` with keys in a keyring referenced by `trusted_keyring`.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_verify_pkcs7_signature(struct bpf_dynptr *data_p, struct bpf_dynptr *sig_p, struct bpf_key *trusted_keyring)`

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

