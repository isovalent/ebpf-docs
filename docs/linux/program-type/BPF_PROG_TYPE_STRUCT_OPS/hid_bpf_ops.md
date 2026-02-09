---
title: "Struct ops 'hid_bpf_ops'"
description: "This page documents the 'hid_bpf_ops' struct ops, its semantics, capabilities, and limitations."
---
# Struct ops `hid_bpf_ops`

[:octicons-tag-24: v6.11](https://github.com/torvalds/linux/commit/ebc0d8093e8c97de459615438edefad1a4ac352c)

HID BPF ops is a type of struct_ops which allows for the manipulation of HID device events in a driver like fashion.

## Usage

It is not uncommon for HID devices to have deficiencies in their hardware or firmware that manufactures work around using drivers. Such drivers are often not provided for Linux devices. The HID BPF ops allows for the creation and distribution of such drivers using BPF outside of the kernels main release cycle.

For example enabling the rebinding of buttons on HID peripherals.

## Fields and ops

```c
struct hid_bpf_ops {
	int   [hid_id](#hid_id);
	u32   [flags](#flags);
	int (*[hid_device_event](#hid_device_event))(
		[struct hid_bpf_ctx](#struct-hid_bpf_ctx)  *ctx, 
		[enum hid_report_type](#enum-hid_report_type) report_type, 
		u64                  source
	);
	int (*[hid_rdesc_fixup](#hid_rdesc_fixup))([struct hid_bpf_ctx](#struct-hid_bpf_ctx) *ctx);
	int (*[hid_hw_request](#hid_hw_request))(
        [struct hid_bpf_ctx](#struct-hid_bpf_ctx)     *ctx, 
        unsigned char           reportnum, 
        [enum hid_report_type](#enum-hid_report_type)    rtype,
        [enum hid_class_request](#enum-hid_class_request)  reqtype,
        u64                     source
    );
	int (*[hid_hw_output_report](#hid_hw_output_report))([struct hid_bpf_ctx](#struct-hid_bpf_ctx) *ctx, u64 source);
};
```

### `hid_id`

[:octicons-tag-24: v6.11](https://github.com/torvalds/linux/commit/ebc0d8093e8c97de459615438edefad1a4ac352c)

`#!c int hid_id`

This field contains the HID unique ID of the device for which the struct_ops is attached.

### `flags`

[:octicons-tag-24: v6.11](https://github.com/torvalds/linux/commit/ebc0d8093e8c97de459615438edefad1a4ac352c)

`#!c u32 flags`
s
Flags used while attaching the struct_ops to the device. Currently only available value is `0` or ``BPF_F_BEFORE``(4).

If `0`, the current struct_ops will be registered at the end of the list of struct_ops to be called for events. If `BPF_F_BEFORE`, the current struct_ops will be registered at the beginning of the list of struct_ops to be called for events.

### `hid_device_event`

[:octicons-tag-24: v6.11](https://github.com/torvalds/linux/commit/ebc0d8093e8c97de459615438edefad1a4ac352c)

`#!c int (*hid_device_event)(struct hid_bpf_ctx *ctx, enum hid_report_type report_type, u64 source)`

This function/program is called whenever an event is coming in from the device. The callback is executed in the context of an interrupt and may therefore not be [sleepable](../../syscall/BPF_PROG_LOAD.md#bpf_f_sleepable).

**Parameters**

`ctx`: The HID-BPF context as [`struct hid_bpf_ctx`](#struct-hid_bpf_ctx)

`report_type`: See [`enum hid_report_type`](#enum-hid_report_type)

`source`: A u64 referring to a uniq but identifiable source. If `0`, the kernel itself emitted that call. For hidraw, `source` is set to the associated `struct file *`.

**Returns**

`0` on success and keep processing; a positive value to change the incoming size buffer; a negative error code to interrupt the processing of this event.

### `hid_rdesc_fixup`

[:octicons-tag-24: v6.11](https://github.com/torvalds/linux/commit/ebc0d8093e8c97de459615438edefad1a4ac352c)

`#!c int (*hid_rdesc_fixup)(struct hid_bpf_ctx *ctx)`

This function/program is called when the probe function parses the report descriptor of the HID device.

**Parameters**

`ctx`: The HID-BPF context as [`struct hid_bpf_ctx`](#struct-hid_bpf_ctx)

**Returns**

`0` on success and keep processing; a positive value to change the incoming size buffer; a negative error code to interrupt the processing of this device

### `hid_hw_request`

[:octicons-tag-24: v6.11](https://github.com/torvalds/linux/commit/8bd0488b5ea58655ad6fdcbe0408ef49b16882b1)

`#!c int (*hid_hw_request)(struct hid_bpf_ctx *ctx, unsigned char reportnum, enum hid_report_type rtype, enum hid_class_request reqtype, u64 source)`

This function/program is called whenever a HID report is requested from the HID device.

**Parameters**

`ctx`: The HID-BPF context as [`struct hid_bpf_ctx`](#struct-hid_bpf_ctx)

`reportnum`: The report number, as in `hid_hw_raw_request()`

`rtype`: The report type, see [`enum hid_class_request`](#enum-hid_class_request)

`source`: A u64 referring to a uniq but identifiable source. If `0`, the kernel itself emitted that call. For hidraw, `source` is set to the associated `struct file *`.

**Returns**

`0` to keep processing the request by hid-core; any other value  stops hid-core from processing that event. A positive value should be  returned with the number of bytes returned in the incoming buffer; a  negative error code interrupts the processing of this call.

### `hid_hw_output_report`

[:octicons-tag-24: v6.11](https://github.com/torvalds/linux/commit/9286675a2aed40a517be8cc4e283a04f473275b5)

`#!c int (*hid_hw_output_report)(struct hid_bpf_ctx *ctx, u64 source)`

This function/program is called whenever a output report is emitted by the device.

**Parameters**

`ctx`: The HID-BPF context as [`struct hid_bpf_ctx`](#struct-hid_bpf_ctx)

`source`: A u64 referring to a uniq but identifiable source. If `0`, the kernel itself emitted that call. For hidraw, `source` is set to the associated `struct file *`.

**Returns**

`0` to keep processing the request by hid-core; any other value stops hid-core from processing that event. A positive value should be returned with the number of bytes written to the device; a negative error code interrupts the processing of this call.

## Types

### `struct hid_bpf_ctx`

User accessible data for all HID programs.

`data` is not directly accessible from the context. We need to issue a call to [`hid_bpf_get_data`](../../kfuncs/hid_bpf_get_data.md) in order to get a pointer to that field.

```c
struct hid_bpf_ctx {
	struct hid_device *hid;
	__u32 allocated_size;
	union {
		__s32 retval;
		__s32 size;
	};
};
```

#### `hid`

The [`struct hid_device`](https://elixir.bootlin.com/linux/v6.13.1/source/include/linux/hid.h#L604) representing the device itself.

This struct is read-only, except for `name`, `uniq`, and `phys` which can be modified from the [`hid_rdesc_fixup`](#hid_rdesc_fixup) function/program.

#### `allocated_size`

Allocated size of data. This is how much memory is available and can be requested by the HID program. Note that for [`hid_rdesc_fixup`](#hid_rdesc_fixup), that memory is set to `4096` (4 KB)

#### `size`

Valid data in the data field. Programs can get the available valid size in data by fetching this field. Programs can also change this value by returning a positive number in the program. To discard the event, return a negative error code. `size` must always be less or equal than `allocated_size` (it is enforced once all BPF programs have been run).

#### `retval`

Return value of the previous program. `hid` and `allocated_size` are read-only, `size` and `retval` are read-write.

### `enum hid_report_type`

```c
enum hid_report_type {
	HID_INPUT_REPORT    = 0,
	HID_OUTPUT_REPORT   = 1,
	HID_FEATURE_REPORT  = 2,
};
```

### `enum hid_class_request`

```c
enum hid_class_request {
	HID_REQ_GET_REPORT      = 0x01,
	HID_REQ_GET_IDLE        = 0x02,
	HID_REQ_GET_PROTOCOL    = 0x03,
	HID_REQ_SET_REPORT      = 0x09,
	HID_REQ_SET_IDLE        = 0x0A,
	HID_REQ_SET_PROTOCOL    = 0x0B,
};
```

## Example

```c
// SPDX-License-Identifier: GPL-2.0-only
/* Copyright (c) 2023 Benjamin Tissoires
 */

#include "vmlinux.h"
#include "hid_bpf.h"
#include "hid_bpf_helpers.h"
#include <bpf/bpf_tracing.h>

#define VID_HP 0x03F0
#define PID_ELITE_PRESENTER 0x464A

HID_BPF_CONFIG(
	HID_DEVICE(BUS_BLUETOOTH, HID_GROUP_GENERIC, VID_HP, PID_ELITE_PRESENTER)
);

/*
 * Already fixed as of commit 0db117359e47 ("HID: add quirk for 03f0:464a
 * HP Elite Presenter Mouse") in the kernel, but this is a slightly better
 * fix.
 *
 * The HP Elite Presenter Mouse HID Record Descriptor shows
 * two mice (Report ID 0x1 and 0x2), one keypad (Report ID 0x5),
 * two Consumer Controls (Report IDs 0x6 and 0x3).
 * Prior to these fixes it registers one mouse, one keypad
 * and one Consumer Control, and it was usable only as a
 * digital laser pointer (one of the two mouses).
 * We replace the second mouse collection with a pointer collection,
 * allowing to use the device both as a mouse and a digital laser
 * pointer.
 */

SEC(HID_BPF_RDESC_FIXUP)
int BPF_PROG(hid_fix_rdesc, struct hid_bpf_ctx *hctx)
{
	__u8 *data = hid_bpf_get_data(hctx, 0 /* offset */, 4096 /* size */);

	if (!data)
		return 0; /* EPERM check */

	/* replace application mouse by application pointer on the second collection */
	if (data[79] == 0x02)
		data[79] = 0x01;

	return 0;
}

HID_BPF_OPS(hp_elite_presenter) = {
	.hid_rdesc_fixup = (void *)hid_fix_rdesc,
};

SEC("syscall")
int probe(struct hid_bpf_probe_args *ctx)
{
	ctx->retval = ctx->rdesc_size != 264;
	if (ctx->retval)
		ctx->retval = -EINVAL;

	return 0;
}

char _license[] SEC("license") = "GPL";
```
