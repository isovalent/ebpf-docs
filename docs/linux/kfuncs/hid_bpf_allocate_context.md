---
title: "KFunc 'hid_bpf_allocate_context'"
description: "This page documents the 'hid_bpf_allocate_context' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `hid_bpf_allocate_context`

<!-- [FEATURE_TAG](hid_bpf_allocate_context) -->
[:octicons-tag-24: v6.3](https://github.com/torvalds/linux/commit/91a7f802d1852f60139712bdcfa98db547ce0531)
<!-- [/FEATURE_TAG] -->

Allocate a context to the given HID device

## Definition

**Parameters**

`hid_id`: the system unique identifier of the HID device

**Returns**

A pointer to &struct hid_bpf_ctx on success, `NULL` on error.

`#!c struct hid_bpf_ctx *hid_bpf_allocate_context(unsigned int hid_id)`

!!! note
	This kfunc returns a pointer to a refcounted object. The verifier will then ensure that the pointer to the object 
	is eventually released using a release kfunc, or transferred to a map using a referenced kptr 
	(by invoking [`bpf_kptr_xchg`](../helper-function/bpf_kptr_xchg.md)). If not, the verifier fails the 
	loading of the BPF program until no lingering references remain in all possible explored states of the program.

!!! note
	The pointer returned by the kfunc may be NULL. Hence, it forces the user to do a NULL check on the pointer returned 
	from the kfunc before making use of it (dereferencing or passing to another helper).

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

```c
// SPDX-License-Identifier: GPL-2.0-only
/* Copyright (c) 2022 Benjamin Tissoires
 */

#include "vmlinux.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>
#include "hid_bpf_helpers.h"

#define HID_UP_BUTTON		0x0009
#define HID_GD_WHEEL		0x0038

SEC("fmod_ret/hid_bpf_device_event")
int BPF_PROG(hid_event, struct hid_bpf_ctx *hctx)
{
	__u8 *data = hid_bpf_get_data(hctx, 0 /* offset */, 9 /* size */);

	if (!data)
		return 0; /* EPERM check */

	/* Touch */
	data[1] &= 0xfd;

	/* X */
	data[4] = 0;
	data[5] = 0;

	/* Y */
	data[6] = 0;
	data[7] = 0;

	return 0;
}

/* 72 == 360 / 5 -> 1 report every 5 degrees */
int resolution = 72;
int physical = 5;

struct haptic_syscall_args {
	unsigned int hid;
	int retval;
};

static __u8 haptic_data[8];

SEC("syscall")
int set_haptic(struct haptic_syscall_args *args)
{
	struct hid_bpf_ctx *ctx;
	const size_t size = sizeof(haptic_data);
	u16 *res;
	int ret;

	if (size > sizeof(haptic_data))
		return -7; /* -E2BIG */

	ctx = hid_bpf_allocate_context(args->hid);
	if (!ctx)
		return -1; /* EPERM check */

	haptic_data[0] = 1;  /* report ID */

	ret = hid_bpf_hw_request(ctx, haptic_data, size, HID_FEATURE_REPORT, HID_REQ_GET_REPORT);

	bpf_printk("probed/remove event ret value: %d", ret);
	bpf_printk("buf: %02x %02x %02x",
		   haptic_data[0],
		   haptic_data[1],
		   haptic_data[2]);
	bpf_printk("     %02x %02x %02x",
		   haptic_data[3],
		   haptic_data[4],
		   haptic_data[5]);
	bpf_printk("     %02x %02x",
		   haptic_data[6],
		   haptic_data[7]);

	/* whenever resolution multiplier is not 3600, we have the fixed report descriptor */
	res = (u16 *)&haptic_data[1];
	if (*res != 3600) {
//		haptic_data[1] = 72; /* resolution multiplier */
//		haptic_data[2] = 0;  /* resolution multiplier */
//		haptic_data[3] = 0;  /* Repeat Count */
		haptic_data[4] = 3;  /* haptic Auto Trigger */
//		haptic_data[5] = 5;  /* Waveform Cutoff Time */
//		haptic_data[6] = 80; /* Retrigger Period */
//		haptic_data[7] = 0;  /* Retrigger Period */
	} else {
		haptic_data[4] = 0;
	}

	ret = hid_bpf_hw_request(ctx, haptic_data, size, HID_FEATURE_REPORT, HID_REQ_SET_REPORT);

	bpf_printk("set haptic ret value: %d -> %d", ret, haptic_data[4]);

	args->retval = ret;

	hid_bpf_release_context(ctx);

	return 0;
}

/* Convert REL_DIAL into REL_WHEEL */
SEC("fmod_ret/hid_bpf_rdesc_fixup")
int BPF_PROG(hid_rdesc_fixup, struct hid_bpf_ctx *hctx)
{
	__u8 *data = hid_bpf_get_data(hctx, 0 /* offset */, 4096 /* size */);
	__u16 *res, *phys;

	if (!data)
		return 0; /* EPERM check */

	/* Convert TOUCH into a button */
	data[31] = HID_UP_BUTTON;
	data[33] = 2;

	/* Convert REL_DIAL into REL_WHEEL */
	data[45] = HID_GD_WHEEL;

	/* Change Resolution Multiplier */
	phys = (__u16 *)&data[61];
	*phys = physical;
	res = (__u16 *)&data[66];
	*res = resolution;

	/* Convert X,Y from Abs to Rel */
	data[88] = 0x06;
	data[98] = 0x06;

	return 0;
}

char _license[] SEC("license") = "GPL";
u32 _version SEC("version") = 1;
```
