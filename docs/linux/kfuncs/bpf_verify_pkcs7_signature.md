---
title: "KFunc 'bpf_verify_pkcs7_signature'"
description: "This page documents the 'bpf_verify_pkcs7_signature' eBPF kfunc, including its defintion, usage, program types that can use it, and examples."
---
# KFunc `bpf_verify_pkcs7_signature`

<!-- [FEATURE_TAG](bpf_verify_pkcs7_signature) -->
[:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/865b0566d8f1a0c3937e5eb4bd6ba4ef03e7e98c)
<!-- [/FEATURE_TAG] -->

Verify a PKCS#7 signature

## Definition

Verify the PKCS#7 signature `sig_ptr` against the supplied `data_ptr` with keys in a keyring referenced by `trusted_keyring`.

<!-- [KFUNC_DEF] -->
`#!c int bpf_verify_pkcs7_signature(struct bpf_dynptr_kern *data_ptr, struct bpf_dynptr_kern *sig_ptr, struct bpf_key *trusted_keyring)`

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../../syscall/BPF_PROG_LOAD/#bpf_f_sleepable).
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
- [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

