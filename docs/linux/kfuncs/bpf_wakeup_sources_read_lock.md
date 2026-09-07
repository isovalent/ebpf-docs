---
title: "KFunc 'bpf_wakeup_sources_read_lock'"
description: "This page documents the 'bpf_wakeup_sources_read_lock' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_wakeup_sources_read_lock`

<!-- [FEATURE_TAG](bpf_wakeup_sources_read_lock) -->
[:octicons-tag-24: v7.2](https://github.com/torvalds/linux/commit/5ff44955447eb04f77161736ff5729c8c0994f7f)
<!-- [/FEATURE_TAG] -->

Acquire the sleepable RCU lock for wake-up sources.

## Definition

The underlying sleepable RCU lock returns an integer index. However, the BPF verifier requires a pointer ([`PTR_TO_BTF_ID`](https://elixir.bootlin.com/linux/v7.2.2/source/include/linux/bpf.h#L1060)) to strictly track the state of acquired resources using [`KF_ACQUIRE`](../concepts/kfuncs.md#kf_acquire) and [`KF_RELEASE`](../concepts/kfuncs.md#kf_release) semantics. We use an opaque structure pointer (`struct bpf_ws_lock *`) to satisfy the verifier while safely encoding the integer index within the pointer address itself.

**Returns**

An opaque pointer encoding the sleepable RCU lock `index + 1` (to avoid `NULL`).

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct bpf_ws_lock *bpf_wakeup_sources_read_lock()`

!!! note
	This kfunc returns a pointer to a refcounted object. The verifier will then ensure that the pointer to the object 
	is eventually released using a release kfunc, or transferred to a map using a referenced kptr 
	(by invoking [`bpf_kptr_xchg`](../helper-function/bpf_kptr_xchg.md)). If not, the verifier fails the 
	loading of the BPF program until no lingering references remain in all possible explored states of the program.
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

