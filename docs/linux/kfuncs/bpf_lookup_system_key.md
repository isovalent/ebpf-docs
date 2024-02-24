---
title: "KFunc 'bpf_lookup_system_key'"
description: "This page documents the 'bpf_lookup_system_key' eBPF kfunc, including its defintion, usage, program types that can use it, and examples."
---
# KFunc `bpf_lookup_system_key`

<!-- [FEATURE_TAG](bpf_lookup_system_key) -->
[:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/f3cf4134c5c6c47b9b5c7aa3cb2d67e107887a7b)
<!-- [/FEATURE_TAG] -->

Lookup a key by a system-defined ID

## Definition

Obtain a bpf_key structure with a key pointer set to the passed key ID.
The key pointer is marked as invalid, to prevent [`bpf_key_put()`](bpf_key_put.md) from
attempting to decrement the key reference count on that pointer. The key
pointer set in such way is currently understood only by
[`verify_pkcs7_signature()`](verify_pkcs7_signature.md).

Set `id` to one of the values defined in `include/linux/verification.h`:

- `0` for the primary keyring (immutable keyring of system keys)
- `VERIFY_USE_SECONDARY_KEYRING` for both the primary and secondary keyring
(where keys can be added only if they are vouched for by existing keys
in those keyrings)
- `VERIFY_USE_PLATFORM_KEYRING` for the platform
keyring (primarily used by the integrity subsystem to verify a kexec'ed
kerned image and, possibly, the initramfs signature).

**Return**

a bpf_key pointer with an invalid key pointer set from the pre-determined ID on success, a NULL pointer otherwise

<!-- [KFUNC_DEF] -->
`#!c struct bpf_key *bpf_lookup_system_key(u64 id)`

!!! note
	This kfunc returns a pointer to a refcounted object. The verifier will then ensure that the pointer to the object 
	is eventually released using a release kfunc, or transferred to a map using a referenced kptr 
	(by invoking [`bpf_kptr_xchg`](../../helper-function/bpf_kptr_xchg.md)). If not, the verifier fails the 
	loading of the BPF program until no lingering references remain in all possible explored states of the program.

!!! note
	The pointer returned by the kfunc may be NULL. Hence, it forces the user to do a NULL check on the pointer returned 
	from the kfunc before making use of it (dereferencing or passing to another helper).
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

