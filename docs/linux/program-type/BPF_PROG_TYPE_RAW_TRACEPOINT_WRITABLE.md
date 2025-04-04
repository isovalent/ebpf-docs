---
title: "Program Type 'BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE'"
description: "This page documents the 'BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE' eBPF program type, including its definition, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE) -->
[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/9df1c28bb75217b244257152ab7d788bb2a386d0)
<!-- [/FEATURE_TAG] -->

Raw tracepoint writable programs are similar to [raw tracepoint](BPF_PROG_TYPE_RAW_TRACEPOINT.md) programs, but they allow you to write to the given context.

## Usage

This program type can be attached to tracepoints that were placed at specific locations in the kernel by the kernel developers. Unlike non-writable tracepoints, these ones can write to the whole context or parts of the context. This essentially allows you to modify the kernel's behavior at runtime in a very specific way.

Writable raw tracepoint programs can only be attached to tracepoints which have been created with the [`DEFINE_EVENT_WRITABLE`](https://elixir.bootlin.com/linux/v6.13.7/source/include/trace/bpf_probe.h#L108) or [`DECLARE_TRACE_WRITABLE`](https://elixir.bootlin.com/linux/v6.13.7/source/include/trace/bpf_probe.h#L126) macros.

In practice there are very limited of such tracepoints, on only one as of kernel v6.14 is [`nbd_send_request`](https://elixir.bootlin.com/linux/v6.13.7/source/include/trace/events/nbd.h#L94)

## Context

The context to this program type is an array of `u64` values. Each element representing an argument of the tracepoint. The program has to cast the elements to their proper type, libbpf provides the [`BPF_PROG`](../../ebpf-library/libbpf/ebpf/BPF_PROG.md) macro to help with this.

The first element of the context is referred to as the "writable buffer", it will be a pointer to a values which is allowed to be modified. The verifier will check that you do not attempt to modify any other parts or modify outside of the bounds of the writable buffer.

## Attachment

Raw tracepoints can be attached in two ways, first is with a dedicated syscall, the second method is with the more generic [BPF link](../syscall/BPF_LINK_CREATE.md) syscall.

### Syscall

The dedicated syscall `BPF_RAW_TRACEPOINT_OPEN` can be used to attach the raw tracepoint. This requires the `name` field to be set to a string containing the name of the tracepoint to which the user whishes to attach to. The `prog_fd` attribute field should be set to the file descriptor of the BPF program to attach.

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
    <!-- An libbpf example would be nice -->

## Example

```c
// SPDX-License-Identifier: GPL-2.0
/* Copyright (c) 2020 Facebook */

[SEC](../../ebpf-library/libbpf/ebpf/SEC.md)("raw_tp.w/bpf_testmod_test_writable_bare")
int [BPF_PROG](../../ebpf-library/libbpf/ebpf/BPF_PROG.md)(handle_raw_tp_writable_bare,
	     struct bpf_testmod_test_writable_ctx *writable)
{
	raw_tp_writable_bare_in_val = writable->val;
	writable->early_ret = raw_tp_writable_bare_early_ret;
	writable->val = raw_tp_writable_bare_out_val;
	return 0;
}
```

## Helper functions

Not all helper functions are available in all program types. These are the helper calls available for raw tracepoint writable programs:

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [`bpf_cgrp_storage_delete`](../helper-function/bpf_cgrp_storage_delete.md)
    * [`bpf_cgrp_storage_get`](../helper-function/bpf_cgrp_storage_get.md)
    * [`bpf_copy_from_user`](../helper-function/bpf_copy_from_user.md)
    * [`bpf_copy_from_user_task`](../helper-function/bpf_copy_from_user_task.md)
    * [`bpf_current_task_under_cgroup`](../helper-function/bpf_current_task_under_cgroup.md)
    * [`bpf_dynptr_data`](../helper-function/bpf_dynptr_data.md)
    * [`bpf_dynptr_from_mem`](../helper-function/bpf_dynptr_from_mem.md)
    * [`bpf_dynptr_read`](../helper-function/bpf_dynptr_read.md)
    * [`bpf_dynptr_write`](../helper-function/bpf_dynptr_write.md)
    * [`bpf_find_vma`](../helper-function/bpf_find_vma.md)
    * [`bpf_for_each_map_elem`](../helper-function/bpf_for_each_map_elem.md)
    * [`bpf_get_branch_snapshot`](../helper-function/bpf_get_branch_snapshot.md)
    * [`bpf_get_current_ancestor_cgroup_id`](../helper-function/bpf_get_current_ancestor_cgroup_id.md)
    * [`bpf_get_current_cgroup_id`](../helper-function/bpf_get_current_cgroup_id.md)
    * [`bpf_get_current_comm`](../helper-function/bpf_get_current_comm.md)
    * [`bpf_get_current_pid_tgid`](../helper-function/bpf_get_current_pid_tgid.md)
    * [`bpf_get_current_task`](../helper-function/bpf_get_current_task.md)
    * [`bpf_get_current_task_btf`](../helper-function/bpf_get_current_task_btf.md)
    * [`bpf_get_current_uid_gid`](../helper-function/bpf_get_current_uid_gid.md)
    * [`bpf_get_func_ip`](../helper-function/bpf_get_func_ip.md)
    * [`bpf_get_ns_current_pid_tgid`](../helper-function/bpf_get_ns_current_pid_tgid.md)
    * [`bpf_get_numa_node_id`](../helper-function/bpf_get_numa_node_id.md)
    * [`bpf_get_prandom_u32`](../helper-function/bpf_get_prandom_u32.md)
    * [`bpf_get_smp_processor_id`](../helper-function/bpf_get_smp_processor_id.md)
    * [`bpf_get_stack`](../helper-function/bpf_get_stack.md)
    * [`bpf_get_stackid`](../helper-function/bpf_get_stackid.md)
    * [`bpf_get_task_stack`](../helper-function/bpf_get_task_stack.md)
    * [`bpf_jiffies64`](../helper-function/bpf_jiffies64.md)
    * [`bpf_kptr_xchg`](../helper-function/bpf_kptr_xchg.md)
    * [`bpf_ktime_get_boot_ns`](../helper-function/bpf_ktime_get_boot_ns.md)
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
    * [`bpf_perf_event_read`](../helper-function/bpf_perf_event_read.md)
    * [`bpf_perf_event_read_value`](../helper-function/bpf_perf_event_read_value.md)
    * [`bpf_probe_read`](../helper-function/bpf_probe_read.md)
    * [`bpf_probe_read_kernel`](../helper-function/bpf_probe_read_kernel.md)
    * [`bpf_probe_read_kernel_str`](../helper-function/bpf_probe_read_kernel_str.md)
    * [`bpf_probe_read_str`](../helper-function/bpf_probe_read_str.md)
    * [`bpf_probe_read_user`](../helper-function/bpf_probe_read_user.md)
    * [`bpf_probe_read_user_str`](../helper-function/bpf_probe_read_user_str.md)
    * [`bpf_probe_write_user`](../helper-function/bpf_probe_write_user.md)
    * [`bpf_ringbuf_discard`](../helper-function/bpf_ringbuf_discard.md)
    * [`bpf_ringbuf_discard_dynptr`](../helper-function/bpf_ringbuf_discard_dynptr.md)
    * [`bpf_ringbuf_output`](../helper-function/bpf_ringbuf_output.md)
    * [`bpf_ringbuf_query`](../helper-function/bpf_ringbuf_query.md)
    * [`bpf_ringbuf_reserve`](../helper-function/bpf_ringbuf_reserve.md)
    * [`bpf_ringbuf_reserve_dynptr`](../helper-function/bpf_ringbuf_reserve_dynptr.md)
    * [`bpf_ringbuf_submit`](../helper-function/bpf_ringbuf_submit.md)
    * [`bpf_ringbuf_submit_dynptr`](../helper-function/bpf_ringbuf_submit_dynptr.md)
    * [`bpf_send_signal`](../helper-function/bpf_send_signal.md)
    * [`bpf_send_signal_thread`](../helper-function/bpf_send_signal_thread.md)
    * [`bpf_snprintf`](../helper-function/bpf_snprintf.md)
    * [`bpf_snprintf_btf`](../helper-function/bpf_snprintf_btf.md)
    * [`bpf_spin_lock`](../helper-function/bpf_spin_lock.md)
    * [`bpf_spin_unlock`](../helper-function/bpf_spin_unlock.md)
    * [`bpf_strncmp`](../helper-function/bpf_strncmp.md)
    * [`bpf_tail_call`](../helper-function/bpf_tail_call.md)
    * [`bpf_task_pt_regs`](../helper-function/bpf_task_pt_regs.md)
    * [`bpf_task_storage_delete`](../helper-function/bpf_task_storage_delete.md)
    * [`bpf_task_storage_get`](../helper-function/bpf_task_storage_get.md)
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
There are currently no kfuncs supported for this program type
<!-- [/PROG_KFUNC_REF] -->
