---
title: "Program Type 'BPF_PROG_TYPE_CGROUP_SYSCTL'"
description: "This page documents the 'BPF_PROG_TYPE_CGROUP_SYSCTL' eBPF program type, including its definition, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_CGROUP_SYSCTL`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_CGROUP_SYSCTL) -->
[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/7b146cebe30cb481b0f70d85779da938da818637)
<!-- [/FEATURE_TAG] -->

cGroup sysctl programs are called when a process in the cGroup to which the program is attached attempts to read or write a sysctl option in the `proc` file system.

## Usage

cGroup sysctl programs are typically located in the `cgroup/sysctl` ELF section. These programs can be used to inspect and filter sysctl usage.

These programs must return one of the following
return codes:

* `0` means "reject access to sysctl"
* `1` means "proceed with access"

If program returns `0` user space will get `-1` from [`read(2)`](https://linux.die.net/man/2/read) or [`write(2)`](https://linux.die.net/man/2/write) and `errno` will be set to `EPERM`.

!!! note
    ``BPF_PROG_TYPE_CGROUP_SYSCTL`` is intended to be used in **trusted** root
    environment, for example to monitor sysctl usage or catch unreasonable values
    an application, running as root in a separate cgroup, is trying to set.

    Since `task_dfl_cgroup(current)` is called at `sys_read` / `sys_write` time it
    may return results different from that at `sys_open` time, i.e. process that
    opened sysctl file in proc filesystem may differ from process that is trying
    to read from / write to it and two such processes may run in different
    cGroups, what means ``BPF_PROG_TYPE_CGROUP_SYSCTL`` should not be used as a
    security mechanism to limit sysctl usage.

    As with any cGroup-bpf program additional care should be taken if an
    application running as root in a cGroup should not be allowed to
    detach/replace BPF program attached by administrator.

### Special helpers

Since sysctl knob is represented by a name and a value, sysctl specific BPF helpers focus on providing access to these properties:

* [`bpf_sysctl_get_name`](../helper-function/bpf_sysctl_get_name.md) to get sysctl name as it is visible in `/proc/sys` into provided by BPF program buffer;

* [`bpf_sysctl_get_current_value`](../helper-function/bpf_sysctl_get_current_value.md) to get string value currently held by sysctl into provided by BPF program buffer. This helper is available on both [`read(2)`](https://linux.die.net/man/2/read) from and [`write(2)`](https://linux.die.net/man/2/write) to sysctl;

* [`bpf_sysctl_get_new_value`](../helper-function/bpf_sysctl_get_new_value.md) to get new string value currently being written to sysctl before actual write happens. This helper can be used only on `ctx->write == 1`;

* [`bpf_sysctl_set_new_value`](../helper-function/bpf_sysctl_set_new_value.md) to override new string value currently being written to sysctl before actual write happens. Sysctl value will be overridden starting from the current `ctx->file_pos`. If the whole value has to be overridden BPF program can set `file_pos` to zero before calling to the helper. This helper can be used only on `ctx->write == 1`. New string value set by the helper is treated and verified by kernel same way as an equivalent string passed by user space.

BPF program sees sysctl value same way as user space does in `proc` file system, i.e. as a string. Since many sysctl values represent an integer or a vector of integers, the following helpers can be used to get numeric value from the string:

* `bpf_strtol()` to convert initial part of the string to long integer similar to user space [`strtol(3)`](https://linux.die.net/man/3/strtol)
* `bpf_strtoul()` to convert initial part of the string to unsigned long integer similar to user space [`strtoul(3)`](https://linux.die.net/man/3/strtoul)

## Context

```c
struct bpf_sysctl {
    __u32 write;
    __u32 file_pos;
};
```

### `write`

This field indicates whether sysctl value is being read (`0`) or written (`1`). This field is read-only.

### `file_pos`

This field indicates file position sysctl is being accessed at, read or written. This field is read-write. Writing to the field sets the starting position in sysctl `proc` file [`read(2)`](https://linux.die.net/man/2/read) will be reading from or [`write(2)`](https://linux.die.net/man/2/write) will be writing to. Writing zero to the field can be used e.g. to override whole sysctl value by [`bpf_sysctl_set_new_value`](../helper-function/bpf_sysctl_get_new_value.md) on [`write(2)`](https://linux.die.net/man/2/write) even when it's called by user space on `file_pos > 0`. Writing non-zero value to the field can be used to access part of sysctl value starting from specified `file_pos`. Not all sysctl support access with `file_pos != 0`, e.g. writes to numeric sysctl entries must always be at file position `0`. See also `kernel.sysctl_writes_strict` sysctl.

## Attachment

cGroup socket buffer programs are attached to cGroups via the [`BPF_PROG_ATTACH`](../syscall/BPF_PROG_ATTACH.md) syscall or via [BPF link](../syscall/BPF_LINK_CREATE.md).

## Example

```c
// SPDX-License-Identifier: GPL-2.0
// Copyright (c) 2019 Facebook

#include <stdint.h>
#include <string.h>

#include <linux/stddef.h>
#include <linux/bpf.h>

#include <bpf/bpf_helpers.h>

/* Max supported length of a string with unsigned long in base 10 (pow2 - 1). */
#define MAX_ULONG_STR_LEN 0xF

/* Max supported length of sysctl value string (pow2). */
#define MAX_VALUE_STR_LEN 0x40

#ifndef ARRAY_SIZE
#define ARRAY_SIZE(x) (sizeof(x) / sizeof((x)[0]))
#endif

const char tcp_mem_name[] = "net/ipv4/tcp_mem";
static __always_inline int is_tcp_mem(struct bpf_sysctl *ctx)
{
	unsigned char i;
	char name[sizeof(tcp_mem_name)];
	int ret;

	memset(name, 0, sizeof(name));
	ret = bpf_sysctl_get_name(ctx, name, sizeof(name), 0);
	if (ret < 0 || ret != sizeof(tcp_mem_name) - 1)
		return 0;

#pragma clang loop unroll(full)
	for (i = 0; i < sizeof(tcp_mem_name); ++i)
		if (name[i] != tcp_mem_name[i])
			return 0;

	return 1;
}

SEC("cgroup/sysctl")
int sysctl_tcp_mem(struct bpf_sysctl *ctx)
{
	unsigned long tcp_mem[3] = {0, 0, 0};
	char value[MAX_VALUE_STR_LEN];
	unsigned char i, off = 0;
	volatile int ret;

	if (ctx->write)
		return 0;

	if (!is_tcp_mem(ctx))
		return 0;

	ret = bpf_sysctl_get_current_value(ctx, value, MAX_VALUE_STR_LEN);
	if (ret < 0 || ret >= MAX_VALUE_STR_LEN)
		return 0;

#pragma clang loop unroll(full)
	for (i = 0; i < ARRAY_SIZE(tcp_mem); ++i) {
		ret = bpf_strtoul(value + off, MAX_ULONG_STR_LEN, 0,
				  tcp_mem + i);
		if (ret <= 0 || ret > MAX_ULONG_STR_LEN)
			return 0;
		off += ret & MAX_ULONG_STR_LEN;
	}


	return tcp_mem[0] < tcp_mem[1] && tcp_mem[1] < tcp_mem[2];
}

char _license[] SEC("license") = "GPL";
```

## Helper functions

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [`bpf_cgrp_storage_delete`](../helper-function/bpf_cgrp_storage_delete.md)
    * [`bpf_cgrp_storage_get`](../helper-function/bpf_cgrp_storage_get.md)
    * [`bpf_dynptr_data`](../helper-function/bpf_dynptr_data.md)
    * [`bpf_dynptr_from_mem`](../helper-function/bpf_dynptr_from_mem.md)
    * [`bpf_dynptr_read`](../helper-function/bpf_dynptr_read.md)
    * [`bpf_dynptr_write`](../helper-function/bpf_dynptr_write.md)
    * [`bpf_for_each_map_elem`](../helper-function/bpf_for_each_map_elem.md)
    * [`bpf_get_current_cgroup_id`](../helper-function/bpf_get_current_cgroup_id.md)
    * [`bpf_get_current_pid_tgid`](../helper-function/bpf_get_current_pid_tgid.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
    * [`bpf_get_current_task`](../helper-function/bpf_get_current_task.md)
    * [`bpf_get_current_task_btf`](../helper-function/bpf_get_current_task_btf.md)
    * [`bpf_get_current_uid_gid`](../helper-function/bpf_get_current_uid_gid.md)
    * [`bpf_get_local_storage`](../helper-function/bpf_get_local_storage.md)
    * [`bpf_get_ns_current_pid_tgid`](../helper-function/bpf_get_ns_current_pid_tgid.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
    * [`bpf_get_numa_node_id`](../helper-function/bpf_get_numa_node_id.md)
    * [`bpf_get_prandom_u32`](../helper-function/bpf_get_prandom_u32.md)
    * [`bpf_get_retval`](../helper-function/bpf_get_retval.md)
    * [`bpf_get_smp_processor_id`](../helper-function/bpf_get_smp_processor_id.md)
    * [`bpf_jiffies64`](../helper-function/bpf_jiffies64.md)
    * [`bpf_kptr_xchg`](../helper-function/bpf_kptr_xchg.md)
    * [`bpf_ktime_get_boot_ns`](../helper-function/bpf_ktime_get_boot_ns.md)
    * [`bpf_ktime_get_coarse_ns`](../helper-function/bpf_ktime_get_coarse_ns.md)
    * [`bpf_ktime_get_ns`](../helper-function/bpf_ktime_get_ns.md)
    * [`bpf_ktime_get_tai_ns`](../helper-function/bpf_ktime_get_tai_ns.md)
    * [`bpf_loop`](../helper-function/bpf_loop.md)
    * [`bpf_map_delete_elem`](../helper-function/bpf_map_delete_elem.md)
    * [`bpf_map_lookup_elem`](../helper-function/bpf_map_lookup_elem.md)
    * [`bpf_map_lookup_percpu_elem`](../helper-function/bpf_map_lookup_percpu_elem.md)
    * [`bpf_map_peek_elem`](../helper-function/bpf_map_peek_elem.md)
    * [`bpf_map_pop_elem`](../helper-function/bpf_map_pop_elem.md)
    * [`bpf_map_push_elem`](../helper-function/bpf_map_push_elem.md)
    * [`bpf_map_update_elem`](../helper-function/bpf_map_update_elem.md)
    * [`bpf_per_cpu_ptr`](../helper-function/bpf_per_cpu_ptr.md)
    * [`bpf_perf_event_output`](../helper-function/bpf_perf_event_output.md)
    * [`bpf_probe_read_kernel`](../helper-function/bpf_probe_read_kernel.md)
    * [`bpf_probe_read_kernel_str`](../helper-function/bpf_probe_read_kernel_str.md)
    * [`bpf_probe_read_user`](../helper-function/bpf_probe_read_user.md)
    * [`bpf_probe_read_user_str`](../helper-function/bpf_probe_read_user_str.md)
    * [`bpf_ringbuf_discard`](../helper-function/bpf_ringbuf_discard.md)
    * [`bpf_ringbuf_discard_dynptr`](../helper-function/bpf_ringbuf_discard_dynptr.md)
    * [`bpf_ringbuf_output`](../helper-function/bpf_ringbuf_output.md)
    * [`bpf_ringbuf_query`](../helper-function/bpf_ringbuf_query.md)
    * [`bpf_ringbuf_reserve`](../helper-function/bpf_ringbuf_reserve.md)
    * [`bpf_ringbuf_reserve_dynptr`](../helper-function/bpf_ringbuf_reserve_dynptr.md)
    * [`bpf_ringbuf_submit`](../helper-function/bpf_ringbuf_submit.md)
    * [`bpf_ringbuf_submit_dynptr`](../helper-function/bpf_ringbuf_submit_dynptr.md)
    * [`bpf_set_retval`](../helper-function/bpf_set_retval.md)
    * [`bpf_snprintf`](../helper-function/bpf_snprintf.md)
    * [`bpf_snprintf_btf`](../helper-function/bpf_snprintf_btf.md)
    * [`bpf_spin_lock`](../helper-function/bpf_spin_lock.md)
    * [`bpf_spin_unlock`](../helper-function/bpf_spin_unlock.md)
    * [`bpf_strncmp`](../helper-function/bpf_strncmp.md)
    * [`bpf_strtol`](../helper-function/bpf_strtol.md)
    * [`bpf_strtoul`](../helper-function/bpf_strtoul.md)
    * [`bpf_sysctl_get_current_value`](../helper-function/bpf_sysctl_get_current_value.md)
    * [`bpf_sysctl_get_name`](../helper-function/bpf_sysctl_get_name.md)
    * [`bpf_sysctl_get_new_value`](../helper-function/bpf_sysctl_get_new_value.md)
    * [`bpf_sysctl_set_new_value`](../helper-function/bpf_sysctl_set_new_value.md)
    * [`bpf_tail_call`](../helper-function/bpf_tail_call.md)
    * [`bpf_task_pt_regs`](../helper-function/bpf_task_pt_regs.md)
    * [`bpf_this_cpu_ptr`](../helper-function/bpf_this_cpu_ptr.md)
    * [`bpf_timer_cancel`](../helper-function/bpf_timer_cancel.md)
    * [`bpf_timer_init`](../helper-function/bpf_timer_init.md)
    * [`bpf_timer_set_callback`](../helper-function/bpf_timer_set_callback.md)
    * [`bpf_timer_start`](../helper-function/bpf_timer_start.md)
    * [`bpf_trace_printk`](../helper-function/bpf_trace_printk.md)
    * [`bpf_trace_vprintk`](../helper-function/bpf_trace_vprintk.md)
    * [`bpf_user_ringbuf_drain`](../helper-function/bpf_user_ringbuf_drain.md)
<!-- [/PROG_HELPER_FUNC_REF] -->

## KFuncs

<!-- [PROG_KFUNC_REF] -->
??? abstract "Supported kfuncs"
    - [`bpf_arena_alloc_pages`](../kfuncs/bpf_arena_alloc_pages.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_arena_free_pages`](../kfuncs/bpf_arena_free_pages.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_cast_to_kern_ctx`](../kfuncs/bpf_cast_to_kern_ctx.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_copy_from_user_str`](../kfuncs/bpf_copy_from_user_str.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_dynptr_adjust`](../kfuncs/bpf_dynptr_adjust.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_dynptr_clone`](../kfuncs/bpf_dynptr_clone.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_dynptr_from_skb`](../kfuncs/bpf_dynptr_from_skb.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_dynptr_is_null`](../kfuncs/bpf_dynptr_is_null.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_dynptr_is_rdonly`](../kfuncs/bpf_dynptr_is_rdonly.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_dynptr_size`](../kfuncs/bpf_dynptr_size.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_dynptr_slice`](../kfuncs/bpf_dynptr_slice.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_dynptr_slice_rdwr`](../kfuncs/bpf_dynptr_slice_rdwr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_get_kmem_cache`](../kfuncs/bpf_get_kmem_cache.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_bits_destroy`](../kfuncs/bpf_iter_bits_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_bits_new`](../kfuncs/bpf_iter_bits_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_bits_next`](../kfuncs/bpf_iter_bits_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_css_destroy`](../kfuncs/bpf_iter_css_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_css_new`](../kfuncs/bpf_iter_css_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_css_next`](../kfuncs/bpf_iter_css_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_css_task_destroy`](../kfuncs/bpf_iter_css_task_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_css_task_new`](../kfuncs/bpf_iter_css_task_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_css_task_next`](../kfuncs/bpf_iter_css_task_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_kmem_cache_destroy`](../kfuncs/bpf_iter_kmem_cache_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_kmem_cache_new`](../kfuncs/bpf_iter_kmem_cache_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_kmem_cache_next`](../kfuncs/bpf_iter_kmem_cache_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_num_destroy`](../kfuncs/bpf_iter_num_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_num_new`](../kfuncs/bpf_iter_num_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_num_next`](../kfuncs/bpf_iter_num_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_task_destroy`](../kfuncs/bpf_iter_task_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_task_new`](../kfuncs/bpf_iter_task_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_task_next`](../kfuncs/bpf_iter_task_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_task_vma_destroy`](../kfuncs/bpf_iter_task_vma_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_task_vma_new`](../kfuncs/bpf_iter_task_vma_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_task_vma_next`](../kfuncs/bpf_iter_task_vma_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_local_irq_restore`](../kfuncs/bpf_local_irq_restore.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_local_irq_save`](../kfuncs/bpf_local_irq_save.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_map_sum_elem_count`](../kfuncs/bpf_map_sum_elem_count.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_preempt_disable`](../kfuncs/bpf_preempt_disable.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_preempt_enable`](../kfuncs/bpf_preempt_enable.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_rcu_read_lock`](../kfuncs/bpf_rcu_read_lock.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_rcu_read_unlock`](../kfuncs/bpf_rcu_read_unlock.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_rdonly_cast`](../kfuncs/bpf_rdonly_cast.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_sock_addr_set_sun_path`](../kfuncs/bpf_sock_addr_set_sun_path.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_wq_init`](../kfuncs/bpf_wq_init.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_wq_set_callback_impl`](../kfuncs/bpf_wq_set_callback_impl.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_wq_start`](../kfuncs/bpf_wq_start.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
<!-- [/PROG_KFUNC_REF] -->
