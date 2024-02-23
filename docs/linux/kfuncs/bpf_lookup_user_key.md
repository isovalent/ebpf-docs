---
title: "KFunc 'bpf_lookup_user_key' - eBPF Docs"
description: "This page documents the 'bpf_lookup_user_key' eBPF kfunc, including its defintion, usage, program types that can use it, and examples."
---
# KFunc `bpf_lookup_user_key`

<!-- [FEATURE_TAG](bpf_lookup_user_key) -->
[:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/f3cf4134c5c6c47b9b5c7aa3cb2d67e107887a7b)
<!-- [/FEATURE_TAG] -->

Lookup a key by its serial

## Definition

Search a key with a given `serial` and the provided `flags`. If found, increment the reference count of the key by one, and return it in the `bpf_key` structure.

The `bpf_key` structure must be passed to [`bpf_key_put()`](bpf_key_put.md) when done with it, so that the key reference count is decremented and the `bpf_key` structure is freed.

Permission checks are deferred to the time the key is used by one of the available key-specific kfuncs.

Set `flags` with `KEY_LOOKUP_CREATE`, to attempt creating a requested special keyring (e.g. session keyring), if it doesn't yet exist.

Set `flags` with `KEY_LOOKUP_PARTIAL`, to lookup a key without waiting for the key construction, and to retrieve uninstantiated keys (keys without data attached to them).

**Return**

a bpf_key pointer with a valid key pointer if the key is found, a NULL pointer otherwise.

<!-- [KFUNC_DEF] -->
`#!c struct bpf_key *bpf_lookup_user_key(u32 serial, u64 flags)`

!!! note
	This kfunc returns a pointer to a refcounted object. The verifier will then ensure that the pointer to the object 
	is eventually released using a release kfunc, or transferred to a map using a referenced kptr 
	(by invoking [`bpf_kptr_xchg`](../../helper-function/bpf_kptr_xchg.md)). If not, the verifier fails the 
	loading of the BPF program until no lingering references remain in all possible explored states of the program.

!!! note
	The pointer returned by the kfunc may be NULL. Hence, it forces the user to do a NULL check on the pointer returned 
	from the kfunc before making use of it (dereferencing or passing to another helper).

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

