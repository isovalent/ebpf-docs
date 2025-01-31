---
title: "Program Type 'BPF_PROG_TYPE_LSM'"
description: "This page documents the 'BPF_PROG_TYPE_LSM' eBPF program type, including its definition, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_LSM`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_LSM) -->
[:octicons-tag-24: v5.7](https://github.com/torvalds/linux/commit/fc611f47f2188ade2b48ff6902d5cce8baac0c58)
<!-- [/FEATURE_TAG] -->

`BPF_PROG_TYPE_LSM` are eBPF programs that can attach to LSM (Linux Security Module) hooks. These are the same hooks as used by programs such as SELinux and AppArmor.

## Usage

The primary use case is to implement security software. For example, the `socket_create` hook is called when a process calls the `socket` syscall, if the eBPF program returns `0`
the socket is allowed to be created, but the eBPF program can also return an error value to block the socket creation.

The list of all LSM hooks can be found in [`lsm_hook_defs.h`](https://github.com/torvalds/linux/blob/457391b0380335d5e9a5babdec90ac53928b23b4/include/linux/lsm_hook_defs.h), additional documentation for these hooks lives in [`lsm_hooks.h`](https://github.com/torvalds/linux/blob/457391b0380335d5e9a5babdec90ac53928b23b4/include/linux/lsm_hooks.h) 

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

LSM programs are exclusively attached via bpf links. To do so the program must be loaded with the [`BPF_LSM_MAC`](../syscall/BPF_LINK_CREATE.md#bpf_lsm_mac) expected attach type and use it as the parameter to [`attach_type`](../syscall/BPF_LINK_CREATE.md#attach_type). The [`target_btf_id`](../syscall/BPF_LINK_CREATE.md#target_btf_id) parameter must be populated with the BTF ID of the LSM hook point which can be extracted from the SELinux BTF on the system.

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

## Helper functions

Not all helper functions are available in all program types. These are the helper calls available for LSM programs:

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [`bpf_bprm_opts_set`](../helper-function/bpf_bprm_opts_set.md)
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
    * [`bpf_get_attach_cookie`](../helper-function/bpf_get_attach_cookie.md) [:octicons-tag-24: v5.19](https://github.com/torvalds/linux/commit/2fcc82411e74e5e6aba336561cf56fb899bfae4e)
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
    * [`bpf_get_task_stack`](../helper-function/bpf_get_task_stack.md)
    * [`bpf_getsockopt`](../helper-function/bpf_getsockopt.md) [:octicons-tag-24: v6.0](https://github.com/torvalds/linux/commit/9113d7e48e9128522b9f5a54dfd30dff10509a92)
    * [`bpf_ima_file_hash`](../helper-function/bpf_ima_file_hash.md)
    * [`bpf_ima_inode_hash`](../helper-function/bpf_ima_inode_hash.md)
    * [`bpf_inode_storage_delete`](../helper-function/bpf_inode_storage_delete.md)
    * [`bpf_inode_storage_get`](../helper-function/bpf_inode_storage_get.md)
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
    * [`bpf_setsockopt`](../helper-function/bpf_setsockopt.md) [:octicons-tag-24: v6.0](https://github.com/torvalds/linux/commit/9113d7e48e9128522b9f5a54dfd30dff10509a92)
    * [`bpf_sk_storage_delete`](../helper-function/bpf_sk_storage_delete.md)
    * [`bpf_sk_storage_get`](../helper-function/bpf_sk_storage_get.md)
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
??? abstract "Supported kfuncs"
    - [`bpf_arena_alloc_pages`](../kfuncs/bpf_arena_alloc_pages.md)
    - [`bpf_arena_free_pages`](../kfuncs/bpf_arena_free_pages.md)
    - [`bpf_cast_to_kern_ctx`](../kfuncs/bpf_cast_to_kern_ctx.md)
    - [`bpf_cgroup_acquire`](../kfuncs/bpf_cgroup_acquire.md)
    - [`bpf_cgroup_ancestor`](../kfuncs/bpf_cgroup_ancestor.md)
    - [`bpf_cgroup_from_id`](../kfuncs/bpf_cgroup_from_id.md)
    - [`bpf_cgroup_release`](../kfuncs/bpf_cgroup_release.md)
    - [`bpf_copy_from_user_str`](../kfuncs/bpf_copy_from_user_str.md)
    - [`bpf_cpumask_acquire`](../kfuncs/bpf_cpumask_acquire.md)
    - [`bpf_cpumask_and`](../kfuncs/bpf_cpumask_and.md)
    - [`bpf_cpumask_any_and_distribute`](../kfuncs/bpf_cpumask_any_and_distribute.md)
    - [`bpf_cpumask_any_distribute`](../kfuncs/bpf_cpumask_any_distribute.md)
    - [`bpf_cpumask_clear`](../kfuncs/bpf_cpumask_clear.md)
    - [`bpf_cpumask_clear_cpu`](../kfuncs/bpf_cpumask_clear_cpu.md)
    - [`bpf_cpumask_copy`](../kfuncs/bpf_cpumask_copy.md)
    - [`bpf_cpumask_create`](../kfuncs/bpf_cpumask_create.md)
    - [`bpf_cpumask_empty`](../kfuncs/bpf_cpumask_empty.md)
    - [`bpf_cpumask_equal`](../kfuncs/bpf_cpumask_equal.md)
    - [`bpf_cpumask_first`](../kfuncs/bpf_cpumask_first.md)
    - [`bpf_cpumask_first_and`](../kfuncs/bpf_cpumask_first_and.md)
    - [`bpf_cpumask_first_zero`](../kfuncs/bpf_cpumask_first_zero.md)
    - [`bpf_cpumask_full`](../kfuncs/bpf_cpumask_full.md)
    - [`bpf_cpumask_intersects`](../kfuncs/bpf_cpumask_intersects.md)
    - [`bpf_cpumask_or`](../kfuncs/bpf_cpumask_or.md)
    - [`bpf_cpumask_release`](../kfuncs/bpf_cpumask_release.md)
    - [`bpf_cpumask_set_cpu`](../kfuncs/bpf_cpumask_set_cpu.md)
    - [`bpf_cpumask_setall`](../kfuncs/bpf_cpumask_setall.md)
    - [`bpf_cpumask_subset`](../kfuncs/bpf_cpumask_subset.md)
    - [`bpf_cpumask_test_and_clear_cpu`](../kfuncs/bpf_cpumask_test_and_clear_cpu.md)
    - [`bpf_cpumask_test_and_set_cpu`](../kfuncs/bpf_cpumask_test_and_set_cpu.md)
    - [`bpf_cpumask_test_cpu`](../kfuncs/bpf_cpumask_test_cpu.md)
    - [`bpf_cpumask_weight`](../kfuncs/bpf_cpumask_weight.md)
    - [`bpf_cpumask_xor`](../kfuncs/bpf_cpumask_xor.md)
    - [`bpf_dynptr_adjust`](../kfuncs/bpf_dynptr_adjust.md)
    - [`bpf_dynptr_clone`](../kfuncs/bpf_dynptr_clone.md)
    - [`bpf_dynptr_is_null`](../kfuncs/bpf_dynptr_is_null.md)
    - [`bpf_dynptr_is_rdonly`](../kfuncs/bpf_dynptr_is_rdonly.md)
    - [`bpf_dynptr_size`](../kfuncs/bpf_dynptr_size.md)
    - [`bpf_dynptr_slice`](../kfuncs/bpf_dynptr_slice.md)
    - [`bpf_dynptr_slice_rdwr`](../kfuncs/bpf_dynptr_slice_rdwr.md)
    - [`bpf_get_dentry_xattr`](../kfuncs/bpf_get_dentry_xattr.md)
    - [`bpf_get_file_xattr`](../kfuncs/bpf_get_file_xattr.md)
    - [`bpf_get_fsverity_digest`](../kfuncs/bpf_get_fsverity_digest.md)
    - [`bpf_get_kmem_cache`](../kfuncs/bpf_get_kmem_cache.md)
    - [`bpf_get_task_exe_file`](../kfuncs/bpf_get_task_exe_file.md)
    - [`bpf_iter_bits_destroy`](../kfuncs/bpf_iter_bits_destroy.md)
    - [`bpf_iter_bits_new`](../kfuncs/bpf_iter_bits_new.md)
    - [`bpf_iter_bits_next`](../kfuncs/bpf_iter_bits_next.md)
    - [`bpf_iter_css_destroy`](../kfuncs/bpf_iter_css_destroy.md)
    - [`bpf_iter_css_new`](../kfuncs/bpf_iter_css_new.md)
    - [`bpf_iter_css_next`](../kfuncs/bpf_iter_css_next.md)
    - [`bpf_iter_css_task_destroy`](../kfuncs/bpf_iter_css_task_destroy.md)
    - [`bpf_iter_css_task_new`](../kfuncs/bpf_iter_css_task_new.md)
    - [`bpf_iter_css_task_next`](../kfuncs/bpf_iter_css_task_next.md)
    - [`bpf_iter_kmem_cache_destroy`](../kfuncs/bpf_iter_kmem_cache_destroy.md)
    - [`bpf_iter_kmem_cache_new`](../kfuncs/bpf_iter_kmem_cache_new.md)
    - [`bpf_iter_kmem_cache_next`](../kfuncs/bpf_iter_kmem_cache_next.md)
    - [`bpf_iter_num_destroy`](../kfuncs/bpf_iter_num_destroy.md)
    - [`bpf_iter_num_new`](../kfuncs/bpf_iter_num_new.md)
    - [`bpf_iter_num_next`](../kfuncs/bpf_iter_num_next.md)
    - [`bpf_iter_task_destroy`](../kfuncs/bpf_iter_task_destroy.md)
    - [`bpf_iter_task_new`](../kfuncs/bpf_iter_task_new.md)
    - [`bpf_iter_task_next`](../kfuncs/bpf_iter_task_next.md)
    - [`bpf_iter_task_vma_destroy`](../kfuncs/bpf_iter_task_vma_destroy.md)
    - [`bpf_iter_task_vma_new`](../kfuncs/bpf_iter_task_vma_new.md)
    - [`bpf_iter_task_vma_next`](../kfuncs/bpf_iter_task_vma_next.md)
    - [`bpf_list_pop_back`](../kfuncs/bpf_list_pop_back.md)
    - [`bpf_list_pop_front`](../kfuncs/bpf_list_pop_front.md)
    - [`bpf_list_push_back_impl`](../kfuncs/bpf_list_push_back_impl.md)
    - [`bpf_list_push_front_impl`](../kfuncs/bpf_list_push_front_impl.md)
    - [`bpf_local_irq_restore`](../kfuncs/bpf_local_irq_restore.md)
    - [`bpf_local_irq_save`](../kfuncs/bpf_local_irq_save.md)
    - [`bpf_map_sum_elem_count`](../kfuncs/bpf_map_sum_elem_count.md)
    - [`bpf_obj_drop_impl`](../kfuncs/bpf_obj_drop_impl.md)
    - [`bpf_obj_new_impl`](../kfuncs/bpf_obj_new_impl.md)
    - [`bpf_path_d_path`](../kfuncs/bpf_path_d_path.md)
    - [`bpf_percpu_obj_drop_impl`](../kfuncs/bpf_percpu_obj_drop_impl.md)
    - [`bpf_percpu_obj_new_impl`](../kfuncs/bpf_percpu_obj_new_impl.md)
    - [`bpf_preempt_disable`](../kfuncs/bpf_preempt_disable.md)
    - [`bpf_preempt_enable`](../kfuncs/bpf_preempt_enable.md)
    - [`bpf_put_file`](../kfuncs/bpf_put_file.md)
    - [`bpf_rbtree_add_impl`](../kfuncs/bpf_rbtree_add_impl.md)
    - [`bpf_rbtree_first`](../kfuncs/bpf_rbtree_first.md)
    - [`bpf_rbtree_remove`](../kfuncs/bpf_rbtree_remove.md)
    - [`bpf_rcu_read_lock`](../kfuncs/bpf_rcu_read_lock.md)
    - [`bpf_rcu_read_unlock`](../kfuncs/bpf_rcu_read_unlock.md)
    - [`bpf_rdonly_cast`](../kfuncs/bpf_rdonly_cast.md)
    - [`bpf_refcount_acquire_impl`](../kfuncs/bpf_refcount_acquire_impl.md)
    - [`bpf_send_signal_task`](../kfuncs/bpf_send_signal_task.md)
    - [`bpf_task_acquire`](../kfuncs/bpf_task_acquire.md)
    - [`bpf_task_from_pid`](../kfuncs/bpf_task_from_pid.md)
    - [`bpf_task_from_vpid`](../kfuncs/bpf_task_from_vpid.md)
    - [`bpf_task_get_cgroup1`](../kfuncs/bpf_task_get_cgroup1.md)
    - [`bpf_task_release`](../kfuncs/bpf_task_release.md)
    - [`bpf_task_under_cgroup`](../kfuncs/bpf_task_under_cgroup.md)
    - [`bpf_throw`](../kfuncs/bpf_throw.md)
    - [`bpf_wq_init`](../kfuncs/bpf_wq_init.md)
    - [`bpf_wq_set_callback_impl`](../kfuncs/bpf_wq_set_callback_impl.md)
    - [`bpf_wq_start`](../kfuncs/bpf_wq_start.md)
    - [`cgroup_rstat_flush`](../kfuncs/cgroup_rstat_flush.md)
    - [`cgroup_rstat_updated`](../kfuncs/cgroup_rstat_updated.md)
    - [`crash_kexec`](../kfuncs/crash_kexec.md)
    - [`hid_bpf_allocate_context`](../kfuncs/hid_bpf_allocate_context.md) -  [:octicons-tag-24: v6.11](https://github.com/torvalds/linux/commit/ebc0d8093e8c97de459615438edefad1a4ac352c)
    - [`hid_bpf_get_data`](../kfuncs/hid_bpf_get_data.md) -  [:octicons-tag-24: v6.11](https://github.com/torvalds/linux/commit/ebc0d8093e8c97de459615438edefad1a4ac352c)
    - [`hid_bpf_hw_output_report`](../kfuncs/hid_bpf_hw_output_report.md) -  [:octicons-tag-24: v6.11](https://github.com/torvalds/linux/commit/ebc0d8093e8c97de459615438edefad1a4ac352c)
    - [`hid_bpf_hw_request`](../kfuncs/hid_bpf_hw_request.md) -  [:octicons-tag-24: v6.11](https://github.com/torvalds/linux/commit/ebc0d8093e8c97de459615438edefad1a4ac352c)
    - [`hid_bpf_input_report`](../kfuncs/hid_bpf_input_report.md) -  [:octicons-tag-24: v6.11](https://github.com/torvalds/linux/commit/ebc0d8093e8c97de459615438edefad1a4ac352c)
    - [`hid_bpf_release_context`](../kfuncs/hid_bpf_release_context.md) -  [:octicons-tag-24: v6.11](https://github.com/torvalds/linux/commit/ebc0d8093e8c97de459615438edefad1a4ac352c)
    - [`hid_bpf_try_input_report`](../kfuncs/hid_bpf_try_input_report.md) -  [:octicons-tag-24: v6.11](https://github.com/torvalds/linux/commit/ebc0d8093e8c97de459615438edefad1a4ac352c)
<!-- [/PROG_KFUNC_REF] -->

