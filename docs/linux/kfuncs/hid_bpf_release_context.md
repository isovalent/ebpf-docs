---
title: "KFunc 'hid_bpf_release_context'"
description: "This page documents the 'hid_bpf_release_context' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `hid_bpf_release_context`

<!-- [FEATURE_TAG](hid_bpf_release_context) -->
[:octicons-tag-24: v6.3](https://github.com/torvalds/linux/commit/91a7f802d1852f60139712bdcfa98db547ce0531)
<!-- [/FEATURE_TAG] -->

Release the previously allocated context @ctx

## Definition

`ctx`: the HID-BPF context to release

`#!c void hid_bpf_release_context(struct hid_bpf_ctx *ctx)`

!!! note
	This kfunc releases the pointer passed in to it. There can be only one referenced pointer that can be passed in. 
	All copies of the pointer being released are invalidated as a result of invoking this kfunc.

!!! note
	This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).
	This is only true when not used from [`BPF_PROG_SYSCALL`](../program-type/BPF_PROG_SYSCALL.md) programs.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md) Until [:octicons-tag-24: v6.11](https://github.com/torvalds/linux/commit/ebc0d8093e8c97de459615438edefad1a4ac352c)
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md) Until [:octicons-tag-24: v6.11](https://github.com/torvalds/linux/commit/ebc0d8093e8c97de459615438edefad1a4ac352c)
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md) Since [:octicons-tag-24: v6.11](https://github.com/torvalds/linux/commit/ebc0d8093e8c97de459615438edefad1a4ac352c)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)

### Example

See [hid_bpf_allocate_context](hid_bpf_allocate_context.md#example) for an example of how to use this kfunc.

