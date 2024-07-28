---
title: "Program Type 'BPF_PROG_TYPE_CGROUP_DEVICE'"
description: "This page documents the 'BPF_PROG_TYPE_CGROUP_DEVICE' eBPF program type, including its definition, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_CGROUP_DEVICE`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_CGROUP_DEVICE) -->
[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/ebc614f687369f9df99828572b1d85a7c2de3d92)
<!-- [/FEATURE_TAG] -->

cGroup device programs are executed when a process in the cGroup to which the program is attached wishes to utilize a device. The program can then decide whether or not to allow the process to allow that operation.

## Usage

This program type is the cGroup v2 variant of the [device whitelist controller](https://www.kernel.org/doc/Documentation/admin-guide/cgroup-v1/devices.rst).  This program type is typically located in a `cgroup/dev` ELF section. It is called with a context describing the access attempt, if the program returns `0`, the attempt fails with `-EPERM`, otherwise it succeeds.

## Context

The program takes a pointer to the `bpf_cgroup_dev_ctx` structure, which describes the device access attempt: access type (mknod/read/write) and device (type, major and minor numbers). If the program returns 0, the attempt fails with `-EPERM`, otherwise it succeeds.

```c
struct bpf_cgroup_dev_ctx {
	/* access_type encoded as (BPF_DEVCG_ACC_* << 16) | BPF_DEVCG_DEV_* */
	__u32 access_type;
	__u32 major;
	__u32 minor;
};

enum {
	BPF_DEVCG_ACC_MKNOD	= (1ULL << 0),
	BPF_DEVCG_ACC_READ	= (1ULL << 1),
	BPF_DEVCG_ACC_WRITE	= (1ULL << 2),
};

enum {
	BPF_DEVCG_DEV_BLOCK	= (1ULL << 0),
	BPF_DEVCG_DEV_CHAR	= (1ULL << 1),
};
```

## Attachment

cGroup socket buffer programs are attached to cGroups via the [`BPF_PROG_ATTACH`](../syscall/BPF_PROG_ATTACH.md) syscall or via [BPF link](../syscall/BPF_LINK_CREATE.md).

## Example

Example BPF program:

```c
/* Copyright (c) 2017 Facebook
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of version 2 of the GNU General Public
 * License as published by the Free Software Foundation.
 */

#include <linux/bpf.h>
#include <linux/version.h>
#include <bpf/bpf_helpers.h>

SEC("cgroup/dev")
int bpf_prog1(struct bpf_cgroup_dev_ctx *ctx)
{
	short type = ctx->access_type & 0xFFFF;
#ifdef DEBUG
	short access = ctx->access_type >> 16;
	char fmt[] = "  %d:%d    \n";

	switch (type) {
	case BPF_DEVCG_DEV_BLOCK:
		fmt[0] = 'b';
		break;
	case BPF_DEVCG_DEV_CHAR:
		fmt[0] = 'c';
		break;
	default:
		fmt[0] = '?';
		break;
	}

	if (access & BPF_DEVCG_ACC_READ)
		fmt[8] = 'r';

	if (access & BPF_DEVCG_ACC_WRITE)
		fmt[9] = 'w';

	if (access & BPF_DEVCG_ACC_MKNOD)
		fmt[10] = 'm';

	bpf_trace_printk(fmt, sizeof(fmt), ctx->major, ctx->minor);
#endif

	/* Allow access to /dev/zero and /dev/random.
	 * Forbid everything else.
	 */
	if (ctx->major != 1 || type != BPF_DEVCG_DEV_CHAR)
		return 0;

	switch (ctx->minor) {
	case 5: /* 1:5 /dev/zero */
	case 9: /* 1:9 /dev/urandom */
		return 1;
	}

	return 0;
}

char _license[] SEC("license") = "GPL";
```

## Helper functions

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [`bpf_get_current_uid_gid`](../helper-function/bpf_get_current_uid_gid.md)
    * [`bpf_get_local_storage`](../helper-function/bpf_get_local_storage.md)
    * [`bpf_get_current_cgroup_id`](../helper-function/bpf_get_current_cgroup_id.md)
    * [`bpf_perf_event_output`](../helper-function/bpf_perf_event_output.md)
    * [`bpf_get_retval`](../helper-function/bpf_get_retval.md)
    * [`bpf_set_retval`](../helper-function/bpf_set_retval.md)
    * [`bpf_map_lookup_elem`](../helper-function/bpf_map_lookup_elem.md)
    * [`bpf_map_update_elem`](../helper-function/bpf_map_update_elem.md)
    * [`bpf_map_delete_elem`](../helper-function/bpf_map_delete_elem.md)
    * [`bpf_map_push_elem`](../helper-function/bpf_map_push_elem.md)
    * [`bpf_map_pop_elem`](../helper-function/bpf_map_pop_elem.md)
    * [`bpf_map_peek_elem`](../helper-function/bpf_map_peek_elem.md)
    * [`bpf_map_lookup_percpu_elem`](../helper-function/bpf_map_lookup_percpu_elem.md)
    * [`bpf_get_prandom_u32`](../helper-function/bpf_get_prandom_u32.md)
    * [`bpf_get_smp_processor_id`](../helper-function/bpf_get_smp_processor_id.md)
    * [`bpf_get_numa_node_id`](../helper-function/bpf_get_numa_node_id.md)
    * [`bpf_tail_call`](../helper-function/bpf_tail_call.md)
    * [`bpf_ktime_get_ns`](../helper-function/bpf_ktime_get_ns.md)
    * [`bpf_ktime_get_boot_ns`](../helper-function/bpf_ktime_get_boot_ns.md)
    * [`bpf_ringbuf_output`](../helper-function/bpf_ringbuf_output.md)
    * [`bpf_ringbuf_reserve`](../helper-function/bpf_ringbuf_reserve.md)
    * [`bpf_ringbuf_submit`](../helper-function/bpf_ringbuf_submit.md)
    * [`bpf_ringbuf_discard`](../helper-function/bpf_ringbuf_discard.md)
    * [`bpf_ringbuf_query`](../helper-function/bpf_ringbuf_query.md)
    * [`bpf_for_each_map_elem`](../helper-function/bpf_for_each_map_elem.md)
    * [`bpf_loop`](../helper-function/bpf_loop.md)
    * [`bpf_strncmp`](../helper-function/bpf_strncmp.md)
    * [`bpf_spin_lock`](../helper-function/bpf_spin_lock.md)
    * [`bpf_spin_unlock`](../helper-function/bpf_spin_unlock.md)
    * [`bpf_jiffies64`](../helper-function/bpf_jiffies64.md)
    * [`bpf_per_cpu_ptr`](../helper-function/bpf_per_cpu_ptr.md)
    * [`bpf_this_cpu_ptr`](../helper-function/bpf_this_cpu_ptr.md)
    * [`bpf_timer_init`](../helper-function/bpf_timer_init.md)
    * [`bpf_timer_set_callback`](../helper-function/bpf_timer_set_callback.md)
    * [`bpf_timer_start`](../helper-function/bpf_timer_start.md)
    * [`bpf_timer_cancel`](../helper-function/bpf_timer_cancel.md)
    * [`bpf_trace_printk`](../helper-function/bpf_trace_printk.md)
    * [`bpf_get_current_task`](../helper-function/bpf_get_current_task.md)
    * [`bpf_get_current_task_btf`](../helper-function/bpf_get_current_task_btf.md)
    * [`bpf_probe_read_user`](../helper-function/bpf_probe_read_user.md)
    * [`bpf_probe_read_kernel`](../helper-function/bpf_probe_read_kernel.md)
    * [`bpf_probe_read_user_str`](../helper-function/bpf_probe_read_user_str.md)
    * [`bpf_probe_read_kernel_str`](../helper-function/bpf_probe_read_kernel_str.md)
    * [`bpf_snprintf_btf`](../helper-function/bpf_snprintf_btf.md)
    * [`bpf_snprintf`](../helper-function/bpf_snprintf.md)
    * [`bpf_task_pt_regs`](../helper-function/bpf_task_pt_regs.md)
    * [`bpf_trace_vprintk`](../helper-function/bpf_trace_vprintk.md)
    * [`bpf_cgrp_storage_get`](../helper-function/bpf_cgrp_storage_get.md)
    * [`bpf_cgrp_storage_delete`](../helper-function/bpf_cgrp_storage_delete.md)
    * [`bpf_dynptr_data`](../helper-function/bpf_dynptr_data.md)
    * [`bpf_dynptr_from_mem`](../helper-function/bpf_dynptr_from_mem.md)
    * [`bpf_dynptr_read`](../helper-function/bpf_dynptr_read.md)
    * [`bpf_dynptr_write`](../helper-function/bpf_dynptr_write.md)
    * [`bpf_kptr_xchg`](../helper-function/bpf_kptr_xchg.md)
    * [`bpf_ktime_get_tai_ns`](../helper-function/bpf_ktime_get_tai_ns.md)
    * [`bpf_ringbuf_discard_dynptr`](../helper-function/bpf_ringbuf_discard_dynptr.md)
    * [`bpf_ringbuf_reserve_dynptr`](../helper-function/bpf_ringbuf_reserve_dynptr.md)
    * [`bpf_ringbuf_submit_dynptr`](../helper-function/bpf_ringbuf_submit_dynptr.md)
    * [`bpf_user_ringbuf_drain`](../helper-function/bpf_user_ringbuf_drain.md)
<!-- [/PROG_HELPER_FUNC_REF] -->

## KFuncs

<!-- [PROG_KFUNC_REF] -->
There are currently no kfuncs supported for this program type
<!-- [/PROG_KFUNC_REF] -->
