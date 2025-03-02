---
title: "SCX eBPF macro 'BPF_FOR_EACH_ITER'"
description: "This page documents the 'BPF_FOR_EACH_ITER' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `BPF_FOR_EACH_ITER`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/4c30f5ce4f7af4f639af99e0bdeada8b268b7361)

The `BPF_FOR_EACH_ITER` macro is used as value for `it__iter` when calling [`scx_bpf_dsq_move`](../../linux/kfuncs/scx_bpf_dsq_move.md) or [`scx_bpf_dsq_move_vtime`](../../linux/kfuncs/scx_bpf_dsq_move_vtime.md) from within [`bpf_for_each`](../libbpf/ebpf/bpf_for_each.md) loops.

## Definition

```c
#define BPF_FOR_EACH_ITER	(&___it)
```

## Usage

Some kfuncs require a pointer to an open coded iterator. When using the `bpf_for_each` macro, this iterator is implicitly defined. The `BPF_FOR_EACH_ITER` macro takes the address of this implicitly defined iterator so it can be passed.

### Example

Loop over all tasks in the shared DSQ and move them to bpf scheduler defined `SOME_DSQ`.

```c hl_lines="3"
struct task_struct *p;
[bpf_for_each](../libbpf/ebpf/bpf_for_each.md)(scx_dsq, p, SHARED_DSQ, 0) {
    [scx_bpf_dsq_move_vtime](../../linux/kfuncs/scx_bpf_dsq_move_vtime.md)(BPF_FOR_EACH_ITER, p, SOME_DSQ, 0);
}
```
