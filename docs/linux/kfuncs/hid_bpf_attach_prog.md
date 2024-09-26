---
title: "KFunc 'hid_bpf_attach_prog'"
description: "This page documents the 'hid_bpf_attach_prog' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `hid_bpf_attach_prog`

[:octicons-tag-24: v6.3](https://github.com/torvalds/linux/commit/f5c27da4e3c8a2e42fb4f41a0c685debcb9af294) - [:octicons-tag-24: v6.11](https://github.com/torvalds/linux/commit/4a86220e046da009bef0948e9f51d1d26d68f93c)

Attach the given `prog_fd` to the given HID device

!!! warning
	This kfunc has been removed in [:octicons-tag-24: v6.11](https://github.com/torvalds/linux/commit/4a86220e046da009bef0948e9f51d1d26d68f93c) HID eBPF programs are now attached as [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md).

## Definition

**Parameters**

`hid_id`: the system unique identifier of the HID device

`prog_fd`: an fd in the user process representing the program to attach

`flags`: any logical OR combination of &enum hid_bpf_attach_flags

**Returns**

A file descriptor of a `bpf_link` object on success (> %0), an error code otherwise. Closing this file descriptor will detach the program from the HID device (unless the `bpf_link` is pinned to the BPF file system).

<!-- [KFUNC_DEF] -->
`#!c int hid_bpf_attach_prog(unsigned int hid_id, int prog_fd, __u32 flags)`
<!-- [/KFUNC_DEF] -->

## Usage

This kfunc is used to attach a BPF program to the jump-table in the BPF-HID subsystem. 

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

```c
// SPDX-License-Identifier: GPL-2.0
/* Copyright (c) 2022 Red hat */
#include "vmlinux.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>

char _license[] SEC("license") = "GPL";

extern __u8 *hid_bpf_get_data(struct hid_bpf_ctx *ctx,
			      unsigned int offset,
			      const size_t __sz) __ksym;
extern int hid_bpf_attach_prog(unsigned int hid_id, int prog_fd, u32 flags) __ksym;

struct attach_prog_args {
	int prog_fd;
	unsigned int hid;
	int retval;
};

__u64 callback_check = 52;
__u64 callback2_check = 52;

SEC("?fmod_ret/hid_bpf_device_event")
int BPF_PROG(hid_first_event, struct hid_bpf_ctx *hid_ctx)
{
	__u8 *rw_data = hid_bpf_get_data(hid_ctx, 0 /* offset */, 3 /* size */);

	if (!rw_data)
		return 0; /* EPERM check */

	callback_check = rw_data[1];

	rw_data[2] = rw_data[1] + 5;

	return 0;
}

SEC("syscall")
int attach_prog(struct attach_prog_args *ctx)
{
	ctx->retval = hid_bpf_attach_prog(ctx->hid,
					  ctx->prog_fd,
					  0);
	return 0;
}
```
