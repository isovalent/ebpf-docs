---
title: "KFunc 'hid_bpf_hw_request'"
description: "This page documents the 'hid_bpf_hw_request' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `hid_bpf_hw_request`

<!-- [FEATURE_TAG](hid_bpf_hw_request) -->
[:octicons-tag-24: v6.3](https://github.com/torvalds/linux/commit/91a7f802d1852f60139712bdcfa98db547ce0531)
<!-- [/FEATURE_TAG] -->

Communicate with a HID device

## Definition

**Parameters**

`ctx`: the HID-BPF context previously allocated in [`hid_bpf_allocate_context`](hid_bpf_allocate_context.md)

`buf`: a `PTR_TO_MEM` buffer

`buf__sz`: the size of the data to transfer

`rtype`: the type of the report (`HID_INPUT_REPORT`, `HID_FEATURE_REPORT`, `HID_OUTPUT_REPORT`)

`reqtype`: the type of the request (`HID_REQ_GET_REPORT`, `HID_REQ_SET_REPORT`, ...)

**Returns**

`0` on success, a negative error code otherwise.

`#!c int hid_bpf_hw_request(struct hid_bpf_ctx *ctx, __u8 *buf, size_t buf__sz, hid_report_type rtype, hid_class_request reqtype)`

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).
    This is only true when not used from [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md) programs.

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
