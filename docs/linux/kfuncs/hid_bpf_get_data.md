# KFunc `hid_bpf_get_data`

<!-- [FEATURE_TAG](hid_bpf_get_data) -->
[:octicons-tag-24: v6.3](https://github.com/torvalds/linux/commit/f5c27da4e3c8a2e42fb4f41a0c685debcb9af294)
<!-- [/FEATURE_TAG] -->

Get the kernel memory pointer associated with the context @ctx

## Definition

**Parameters**

`ctx`: The HID-BPF context

`offset`: The offset within the memory

`rdwr_buf_size`: the const size of the buffer

**Returns**

`NULL` on error, an `__u8` memory pointer on success

<!-- [KFUNC_DEF] -->
`#!c __u8 *hid_bpf_get_data(struct hid_bpf_ctx *ctx, unsigned int offset, const size_t rdwr_buf_size)`

!!! note
	The pointer returned by the kfunc may be NULL. Hence, it forces the user to do a NULL check on the pointer returned 
	from the kfunc before making use of it (dereferencing or passing to another helper).
<!-- [/KFUNC_DEF] -->

## Usage

The goal of HID-BPF is to partially replace drivers, so this situation can be problematic because we might have programs which will step on each other toes.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_LSM](../../program-types/BPF_PROG_TYPE_LSM.md)
- [BPF_PROG_TYPE_TRACING](../../program-types/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

```c
// SPDX-License-Identifier: GPL-2.0

#include "vmlinux.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>

/* following are kfuncs exported by HID for HID-BPF */
extern int hid_bpf_attach_prog(unsigned int hid_id, int prog_fd, u32 flags) __ksym;
extern __u8 *hid_bpf_get_data(struct hid_bpf_ctx *ctx,
			      unsigned int offset,
			      const size_t __sz) __ksym;
extern void hid_bpf_data_release(__u8 *data) __ksym;
extern int hid_bpf_hw_request(struct hid_bpf_ctx *ctx) __ksym;

struct attach_prog_args {
	int prog_fd;
	unsigned int hid;
	int retval;
};

SEC("syscall")
int attach_prog(struct attach_prog_args *ctx)
{
	ctx->retval = hid_bpf_attach_prog(ctx->hid,
					  ctx->prog_fd,
					  0);
	return 0;
}

SEC("fmod_ret/hid_bpf_device_event")
int BPF_PROG(hid_y_event, struct hid_bpf_ctx *hctx)
{
	s16 y;
	__u8 *data = hid_bpf_get_data(hctx, 0 /* offset */, 9 /* size */);

	if (!data)
		return 0; /* EPERM check */

	bpf_printk("event: size: %d", hctx->size);
	bpf_printk("incoming event: %02x %02x %02x",
		   data[0],
		   data[1],
		   data[2]);
	bpf_printk("                %02x %02x %02x",
		   data[3],
		   data[4],
		   data[5]);
	bpf_printk("                %02x %02x %02x",
		   data[6],
		   data[7],
		   data[8]);

	y = data[3] | (data[4] << 8);

	y = -y;

	data[3] = y & 0xFF;
	data[4] = (y >> 8) & 0xFF;

	bpf_printk("modified event: %02x %02x %02x",
		   data[0],
		   data[1],
		   data[2]);
	bpf_printk("                %02x %02x %02x",
		   data[3],
		   data[4],
		   data[5]);
	bpf_printk("                %02x %02x %02x",
		   data[6],
		   data[7],
		   data[8]);

	return 0;
}

SEC("fmod_ret/hid_bpf_device_event")
int BPF_PROG(hid_x_event, struct hid_bpf_ctx *hctx)
{
	s16 x;
	__u8 *data = hid_bpf_get_data(hctx, 0 /* offset */, 9 /* size */);

	if (!data)
		return 0; /* EPERM check */

	x = data[1] | (data[2] << 8);

	x = -x;

	data[1] = x & 0xFF;
	data[2] = (x >> 8) & 0xFF;
	return 0;
}

SEC("fmod_ret/hid_bpf_rdesc_fixup")
int BPF_PROG(hid_rdesc_fixup, struct hid_bpf_ctx *hctx)
{
	__u8 *data = hid_bpf_get_data(hctx, 0 /* offset */, 4096 /* size */);

	if (!data)
		return 0; /* EPERM check */

	bpf_printk("rdesc: %02x %02x %02x",
		   data[0],
		   data[1],
		   data[2]);
	bpf_printk("       %02x %02x %02x",
		   data[3],
		   data[4],
		   data[5]);
	bpf_printk("       %02x %02x %02x ...",
		   data[6],
		   data[7],
		   data[8]);

	/*
	 * The original report descriptor contains:
	 *
	 * 0x05, 0x01,                    //   Usage Page (Generic Desktop)      30
	 * 0x16, 0x01, 0x80,              //   Logical Minimum (-32767)          32
	 * 0x26, 0xff, 0x7f,              //   Logical Maximum (32767)           35
	 * 0x09, 0x30,                    //   Usage (X)                         38
	 * 0x09, 0x31,                    //   Usage (Y)                         40
	 *
	 * So byte 39 contains Usage X and byte 41 Usage Y.
	 *
	 * We simply swap the axes here.
	 */
	data[39] = 0x31;
	data[41] = 0x30;

	return 0;
}

char _license[] SEC("license") = "GPL";
```
