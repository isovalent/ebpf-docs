---
title: "SCX eBPF macro 'bpf_refcount_acquire'"
description: "This page documents the 'bpf_refcount_acquire' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `bpf_refcount_acquire`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/2a52ca7c98960aafb0eca9ef96b2d0c932171357)

The `bpf_refcount_acquire` macro wraps [`bpf_refcount_acquire_impl`](../../linux/kfuncs/bpf_refcount_acquire_impl.md) to provide a more ergonomic interface.

## Definition

```c
#define bpf_refcount_acquire(kptr) bpf_refcount_acquire_impl(kptr, NULL)
```

## Usage

The [`bpf_refcount_acquire_impl`](../../linux/kfuncs/bpf_refcount_acquire_impl.md) kfunc has a quirk where the second argument is always `NULL`, this wrapper abstracts that quirk away.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
