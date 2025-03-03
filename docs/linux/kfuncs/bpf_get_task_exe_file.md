---
title: "KFunc 'bpf_get_task_exe_file'"
description: "This page documents the 'bpf_get_task_exe_file' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_get_task_exe_file`

<!-- [FEATURE_TAG](bpf_get_task_exe_file) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/d08e2045ebf0f5f2a97ad22cc7dae398b35354ba)
<!-- [/FEATURE_TAG] -->

This function gets a reference on the `exe_file` struct file member of the `mm_struct` that is nested within the supplied `task_struct`.

## Definition

Get a reference on the `exe_file` struct file member field of the `mm_struct` nested within the supplied `task`. The referenced file pointer acquired by this BPF kfunc must be released using [`bpf_put_file`](bpf_put_file.md). Failing to call [`bpf_put_file`](bpf_put_file.md) on the returned referenced struct file pointer that has been acquired by this BPF kfunc will result in the BPF program being rejected by the BPF verifier.

This BPF kfunc may only be called from BPF LSM programs.

Internally, this BPF kfunc leans on `get_task_exe_file`, such that calling `bpf_get_task_exe_file` would be analogous to calling `get_task_exe_file` directly in kernel context.

**Parameters**

`task`: `task_struct` of which the nested `mm_struct` `exe_file` member to get a reference on

**Returns** 

A referenced struct file pointer to the `exe_file` member of the `mm_struct` that is nested within the supplied `task`. On error, `NULL` is returned.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c struct file *bpf_get_task_exe_file(struct task_struct *task)`

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
- [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

