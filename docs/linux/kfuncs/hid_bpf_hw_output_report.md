---
title: "KFunc 'hid_bpf_hw_output_report'"
description: "This page documents the 'hid_bpf_hw_output_report' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `hid_bpf_hw_output_report`

<!-- [FEATURE_TAG](hid_bpf_hw_output_report) -->
[:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/5599f80196612efde96dbe6ef18f6ecc0cb4ba19)
<!-- [/FEATURE_TAG] -->

Send an output report to a HID device

## Definition

**Parameters**

`ctx`: the HID-BPF context previously allocated in `hid_bpf_allocate_context()`

`buf`: a `PTR_TO_MEM` buffer

`buf__sz`: the size of the data to transfer

**Returns**

Returns the number of bytes transferred on success, a negative error code otherwise.

<!-- [KFUNC_DEF] -->
`#!c int hid_bpf_hw_output_report(struct hid_bpf_ctx *ctx, __u8 *buf, size_t buf__sz)`
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

