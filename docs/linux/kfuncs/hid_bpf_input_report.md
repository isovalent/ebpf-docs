---
title: "KFunc 'hid_bpf_input_report'"
description: "This page documents the 'hid_bpf_input_report' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `hid_bpf_input_report`

<!-- [FEATURE_TAG](hid_bpf_input_report) -->
[:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/9be50ac30a83896a753ab9f64e941763bb7900be)
<!-- [/FEATURE_TAG] -->

Inject a HID report in the kernel from a HID device

## Definition

**Parameters**

`ctx`: the HID-BPF context previously allocated in `hid_bpf_allocate_context()`

`type`: the type of the report (`HID_INPUT_REPORT`, `HID_FEATURE_REPORT`, `HID_OUTPUT_REPORT`)

`buf`: a `PTR_TO_MEM` buffer

`buf__sz`: the size of the data to transfer


**Returns**

Returns `0` on success, a negative error code otherwise. This function will wait for the device to be available before injecting the event, thus needs to be called in sleepable context.

<!-- [KFUNC_DEF] -->
`#!c int hid_bpf_input_report(struct hid_bpf_ctx *ctx, hid_report_type type, u8 *buf, const size_t buf__sz)`

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).
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

