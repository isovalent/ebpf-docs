---
title: "Program Type 'BPF_PROG_TYPE_LSM'"
description: "This page documents the 'BPF_PROG_TYPE_LSM' eBPF program type, including its defintion, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_LSM`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_LSM) -->
[:octicons-tag-24: v5.7](https://github.com/torvalds/linux/commit/fc611f47f2188ade2b48ff6902d5cce8baac0c58)
<!-- [/FEATURE_TAG] -->

`BPF_PROG_TYPE_LSM` are eBPF programs that can attach to LSM (Linux Security Module) hooks. These are the same hooks as used by programs such as SELinux and AppArmor.

## Usage

The primary use case is to implement security software. For example, the `socket_create` hook is called when a process calls the `socket` syscall, if the eBPF program returns `0`
the socket is allowed to be created, but the eBPF program can also return an error value to block the socket creation.

The list of all LSM hooks can be found in [lsm_hook_defs.h](https://github.com/torvalds/linux/blob/457391b0380335d5e9a5babdec90ac53928b23b4/include/linux/lsm_hook_defs.h), additional documentation for these hooks lives in [lsm_hooks.h](https://github.com/torvalds/linux/blob/457391b0380335d5e9a5babdec90ac53928b23b4/include/linux/lsm_hooks.h) 

```c
// Copyright (C) 2020 Google LLC.
SEC("lsm/file_mprotect")
int BPF_PROG(mprotect_audit, struct vm_area_struct *vma,
            unsigned long reqprot, unsigned long prot, int ret)
{
    /* ret is the return value from the previous BPF program
        * or 0 if it's the first hook.
        */
    if (ret != 0)
        return ret;

    int is_heap;

    is_heap = (vma->vm_start >= vma->vm_mm->start_brk &&
            vma->vm_end <= vma->vm_mm->brk);

    /* Return an -EPERM or write information to the perf events buffer
        * for auditing
        */
    if (is_heap)
        return -EPERM;
}
```

## Context

LSM programs are invoked with an array of `__u64` values equal in length to the amount of arguments of the LSM hook, each index representing the arguments in order. The `BPF_PROG` macro defined in `tools/lib/bpf/bpf_tracing.h` is often used to make it easier to write LSM programs. The macro allows the user to write the arguments as declared on the hooks, the macro will cast the arguments. The actual arguments and their times are determined by the hook to which this program is attached.

## Attachment

LSM programs are exclusively attached via bpf links. To do so the program must be loaded with the [`BPF_LSM_MAC`](../syscall/BPF_LINK_CREATE.md#bpf_lsm_mac) expected attach type and use it as the param to [`attach_type`](../syscall/BPF_LINK_CREATE.md#attach_type). The [`target_btf_id`](../syscall/BPF_LINK_CREATE.md#target_btf_id) parameter must be populated with the BTF ID of the LSM hook point which can be extracted from the selinux BTF on the system.

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

## Helper functions

Not all helper functions are available in all program types. These are the helper calls available for LSM programs:

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [bpf_inode_storage_get](../helper-function/bpf_inode_storage_get.md)
    * [bpf_inode_storage_delete](../helper-function/bpf_inode_storage_delete.md)
    * [bpf_sk_storage_get](../helper-function/bpf_sk_storage_get.md)
    * [bpf_sk_storage_delete](../helper-function/bpf_sk_storage_delete.md)
    * [bpf_spin_lock](../helper-function/bpf_spin_lock.md)
    * [bpf_spin_unlock](../helper-function/bpf_spin_unlock.md)
    * [bpf_bprm_opts_set](../helper-function/bpf_bprm_opts_set.md)
    * [bpf_ima_inode_hash](../helper-function/bpf_ima_inode_hash.md)
    * [bpf_ima_file_hash](../helper-function/bpf_ima_file_hash.md)
    * [bpf_setsockopt](../helper-function/bpf_setsockopt.md) [:octicons-tag-24: v6.0](9113d7e48e9128522b9f5a54dfd30dff10509a92)
    * [bpf_getsockopt](../helper-function/bpf_getsockopt.md) [:octicons-tag-24: v6.0](9113d7e48e9128522b9f5a54dfd30dff10509a92)
    * [bpf_map_lookup_elem](../helper-function/bpf_map_lookup_elem.md)
    * [bpf_map_update_elem](../helper-function/bpf_map_update_elem.md)
    * [bpf_map_delete_elem](../helper-function/bpf_map_delete_elem.md)
    * [bpf_map_push_elem](../helper-function/bpf_map_push_elem.md)
    * [bpf_map_pop_elem](../helper-function/bpf_map_pop_elem.md)
    * [bpf_map_peek_elem](../helper-function/bpf_map_peek_elem.md)
    * [bpf_map_lookup_percpu_elem](../helper-function/bpf_map_lookup_percpu_elem.md)
    * [bpf_ktime_get_ns](../helper-function/bpf_ktime_get_ns.md)
    * [bpf_ktime_get_boot_ns](../helper-function/bpf_ktime_get_boot_ns.md)
    * [bpf_tail_call](../helper-function/bpf_tail_call.md)
    * [bpf_get_current_pid_tgid](../helper-function/bpf_get_current_pid_tgid.md)
    * [bpf_get_current_task](../helper-function/bpf_get_current_task.md)
    * [bpf_get_current_task_btf](../helper-function/bpf_get_current_task_btf.md)
    * [bpf_task_pt_regs](../helper-function/bpf_task_pt_regs.md)
    * [bpf_get_current_uid_gid](../helper-function/bpf_get_current_uid_gid.md)
    * [bpf_get_current_comm](../helper-function/bpf_get_current_comm.md)
    * [bpf_trace_printk](../helper-function/bpf_trace_printk.md)
    * [bpf_get_smp_processor_id](../helper-function/bpf_get_smp_processor_id.md)
    * [bpf_get_numa_node_id](../helper-function/bpf_get_numa_node_id.md)
    * [bpf_perf_event_read](../helper-function/bpf_perf_event_read.md)
    * [bpf_current_task_under_cgroup](../helper-function/bpf_current_task_under_cgroup.md)
    * [bpf_get_prandom_u32](../helper-function/bpf_get_prandom_u32.md)
    * [bpf_probe_write_user](../helper-function/bpf_probe_write_user.md)
    * [bpf_probe_read_user](../helper-function/bpf_probe_read_user.md)
    * [bpf_probe_read_kernel](../helper-function/bpf_probe_read_kernel.md)
    * [bpf_probe_read_user_str](../helper-function/bpf_probe_read_user_str.md)
    * [bpf_probe_read_kernel_str](../helper-function/bpf_probe_read_kernel_str.md)
    * [bpf_probe_read](../helper-function/bpf_probe_read.md)
    * [bpf_probe_read_str](../helper-function/bpf_probe_read_str.md)
    * [bpf_get_current_cgroup_id](../helper-function/bpf_get_current_cgroup_id.md)
    * [bpf_get_current_ancestor_cgroup_id](../helper-function/bpf_get_current_ancestor_cgroup_id.md)
    * [bpf_send_signal](../helper-function/bpf_send_signal.md)
    * [bpf_send_signal_thread](../helper-function/bpf_send_signal_thread.md)
    * [bpf_perf_event_read_value](../helper-function/bpf_perf_event_read_value.md)
    * [bpf_get_ns_current_pid_tgid](../helper-function/bpf_get_ns_current_pid_tgid.md)
    * [bpf_ringbuf_output](../helper-function/bpf_ringbuf_output.md)
    * [bpf_ringbuf_reserve](../helper-function/bpf_ringbuf_reserve.md)
    * [bpf_ringbuf_submit](../helper-function/bpf_ringbuf_submit.md)
    * [bpf_ringbuf_discard](../helper-function/bpf_ringbuf_discard.md)
    * [bpf_ringbuf_query](../helper-function/bpf_ringbuf_query.md)
    * [bpf_jiffies64](../helper-function/bpf_jiffies64.md)
    * [bpf_get_task_stack](../helper-function/bpf_get_task_stack.md)
    * [bpf_copy_from_user](../helper-function/bpf_copy_from_user.md)
    * [bpf_copy_from_user_task](../helper-function/bpf_copy_from_user_task.md)
    * [bpf_snprintf_btf](../helper-function/bpf_snprintf_btf.md)
    * [bpf_per_cpu_ptr](../helper-function/bpf_per_cpu_ptr.md)
    * [bpf_this_cpu_ptr](../helper-function/bpf_this_cpu_ptr.md)
    * [bpf_task_storage_get](../helper-function/bpf_task_storage_get.md)
    * [bpf_task_storage_delete](../helper-function/bpf_task_storage_delete.md)
    * [bpf_for_each_map_elem](../helper-function/bpf_for_each_map_elem.md)
    * [bpf_snprintf](../helper-function/bpf_snprintf.md)
    * [bpf_get_func_ip](../helper-function/bpf_get_func_ip.md)
    * [bpf_get_branch_snapshot](../helper-function/bpf_get_branch_snapshot.md)
    * [bpf_find_vma](../helper-function/bpf_find_vma.md)
    * [bpf_trace_vprintk](../helper-function/bpf_trace_vprintk.md)
    * [bpf_map_lookup_elem](../helper-function/bpf_map_lookup_elem.md)
    * [bpf_map_update_elem](../helper-function/bpf_map_update_elem.md)
    * [bpf_map_delete_elem](../helper-function/bpf_map_delete_elem.md)
    * [bpf_map_push_elem](../helper-function/bpf_map_push_elem.md)
    * [bpf_map_pop_elem](../helper-function/bpf_map_pop_elem.md)
    * [bpf_map_peek_elem](../helper-function/bpf_map_peek_elem.md)
    * [bpf_map_lookup_percpu_elem](../helper-function/bpf_map_lookup_percpu_elem.md)
    * [bpf_get_prandom_u32](../helper-function/bpf_get_prandom_u32.md)
    * [bpf_get_smp_processor_id](../helper-function/bpf_get_smp_processor_id.md)
    * [bpf_get_numa_node_id](../helper-function/bpf_get_numa_node_id.md)
    * [bpf_tail_call](../helper-function/bpf_tail_call.md)
    * [bpf_ktime_get_ns](../helper-function/bpf_ktime_get_ns.md)
    * [bpf_ktime_get_boot_ns](../helper-function/bpf_ktime_get_boot_ns.md)
    * [bpf_ringbuf_output](../helper-function/bpf_ringbuf_output.md)
    * [bpf_ringbuf_reserve](../helper-function/bpf_ringbuf_reserve.md)
    * [bpf_ringbuf_submit](../helper-function/bpf_ringbuf_submit.md)
    * [bpf_ringbuf_discard](../helper-function/bpf_ringbuf_discard.md)
    * [bpf_ringbuf_query](../helper-function/bpf_ringbuf_query.md)
    * [bpf_for_each_map_elem](../helper-function/bpf_for_each_map_elem.md)
    * [bpf_loop](../helper-function/bpf_loop.md)
    * [bpf_strncmp](../helper-function/bpf_strncmp.md)
    * [bpf_spin_lock](../helper-function/bpf_spin_lock.md)
    * [bpf_spin_unlock](../helper-function/bpf_spin_unlock.md)
    * [bpf_jiffies64](../helper-function/bpf_jiffies64.md)
    * [bpf_per_cpu_ptr](../helper-function/bpf_per_cpu_ptr.md)
    * [bpf_this_cpu_ptr](../helper-function/bpf_this_cpu_ptr.md)
    * [bpf_timer_init](../helper-function/bpf_timer_init.md)
    * [bpf_timer_set_callback](../helper-function/bpf_timer_set_callback.md)
    * [bpf_timer_start](../helper-function/bpf_timer_start.md)
    * [bpf_timer_cancel](../helper-function/bpf_timer_cancel.md)
    * [bpf_trace_printk](../helper-function/bpf_trace_printk.md)
    * [bpf_get_current_task](../helper-function/bpf_get_current_task.md)
    * [bpf_get_current_task_btf](../helper-function/bpf_get_current_task_btf.md)
    * [bpf_probe_read_user](../helper-function/bpf_probe_read_user.md)
    * [bpf_probe_read_kernel](../helper-function/bpf_probe_read_kernel.md)
    * [bpf_probe_read_user_str](../helper-function/bpf_probe_read_user_str.md)
    * [bpf_probe_read_kernel_str](../helper-function/bpf_probe_read_kernel_str.md)
    * [bpf_snprintf_btf](../helper-function/bpf_snprintf_btf.md)
    * [bpf_snprintf](../helper-function/bpf_snprintf.md)
    * [bpf_task_pt_regs](../helper-function/bpf_task_pt_regs.md)
    * [bpf_trace_vprintk](../helper-function/bpf_trace_vprintk.md)
<!-- [/PROG_HELPER_FUNC_REF] -->
