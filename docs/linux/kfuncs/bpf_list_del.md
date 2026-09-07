---
title: "KFunc 'bpf_list_del'"
description: "This page documents the 'bpf_list_del' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_list_del`

<!-- [FEATURE_TAG](bpf_list_del) -->
[:octicons-tag-24: v7.2](https://github.com/torvalds/linux/commit/187baa10963ac9f1db5123fa2ab761ab34ea06b9)
<!-- [/FEATURE_TAG] -->

Allow users to remove any node from a linked list.

## Definition

**Parameters**

`head`: The list head from which `node__nonown_allowed` should be removed.

`node__nonown_allowed`: The node to be removed from the list.

**Returns**

Pointer to bpf_list_node of that was removed, or `NULL` if given node was not in the list   .

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct bpf_list_node *bpf_list_del(struct bpf_list_head *head, struct bpf_list_node *node__nonown_allowed)`

!!! note
	This kfunc returns a pointer to a refcounted object. The verifier will then ensure that the pointer to the object 
	is eventually released using a release kfunc, or transferred to a map using a referenced kptr 
	(by invoking [`bpf_kptr_xchg`](../helper-function/bpf_kptr_xchg.md)). If not, the verifier fails the 
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
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
- [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

