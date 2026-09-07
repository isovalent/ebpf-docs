---
title: "KFunc 'bpf_wakeup_sources_read_unlock'"
description: "This page documents the 'bpf_wakeup_sources_read_unlock' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_wakeup_sources_read_unlock`

<!-- [FEATURE_TAG](bpf_wakeup_sources_read_unlock) -->
[:octicons-tag-24: v7.2](https://github.com/torvalds/linux/commit/5ff44955447eb04f77161736ff5729c8c0994f7f)
<!-- [/FEATURE_TAG] -->

Release the sleepable RCU lock for wake-up sources.

## Definition

The BPF verifier guarantees that `lock` is a valid, unreleased pointer from the acquire function. We decode the pointer back into the integer sleepable RCU index by subtracting 1 and release the lock.

**Parameters**

`lock`: The opaque pointer returned by [`bpf_wakeup_sources_read_lock`](bpf_wakeup_sources_read_lock.md)

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void bpf_wakeup_sources_read_unlock(struct bpf_ws_lock *lock)`

!!! note
	This kfunc releases the pointer passed in to it. There can be only one referenced pointer that can be passed in. 
	All copies of the pointer being released are invalidated as a result of invoking this kfunc.
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

