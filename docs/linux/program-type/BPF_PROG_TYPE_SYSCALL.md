---
title: "Program Type 'BPF_PROG_TYPE_SYSCALL'"
description: "This page documents the 'BPF_PROG_TYPE_SYSCALL' eBPF program type, including its definition, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_SYSCALL`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_SYSCALL) -->
[:octicons-tag-24: v5.14](https://github.com/torvalds/linux/commit/79a7f8bdb159d9914b58740f3d31d602a6e4aca8)
<!-- [/FEATURE_TAG] -->

Syscall programs can be used to execute syscalls from eBPF.

## Usage

The abstract purpose of the syscall program type is to execute syscalls from eBPF. The initial use case for this program type was to offload some of the work of loader libraries to syscall eBPF programs. The program type can also be used by for "HID-BPF" to register a BPF program as a HID device driver.

### Loading with light skeletons

This use case revolves around using a `BPF_PROG_TYPE_SYSCALL` program to load one or more eBPF programs. The reason behind this is two-fold. First, with a bit of automation in the form of generation tools, loading a program can be made easier. Second, this new structure would make it easier to implement a form of code signing for eBPF programs. However, the code signing use case so far has not been successful.

The way this works is that you write and compile your primary eBPF program as normal. You then feed it to `bpftool` with the `gen skeleton -L {prog}.o > {prog}.skel.h` command. This will generate a "light skeleton" for the program. Essentially a header file which can be included by a custom userspace program as dependency. It exposes pre-defined function to then load the eBPF program. The header file embeds the essential parts of the primary ELF file and a generated `BPF_PROG_TYPE_SYSCALL` program. Parts of the primary program such as its instructions, map definitions, and initial keys/values are part of the generated program or provided as data via existing mechanisms. The syscall program then uses a series of [`bpf_sys_bpf`](../helper-function/bpf_sys_bpf.md) helper calls to load the primary program just like a loader would normally do from userspace.

### HID-BPF

The use case of HID-BPF is to implement HID device drivers in eBPF, at least partially. This allows HID drivers implemented this way for new devices to work on older kernels without the need for a kernel module.

No special program type was created for this use case, rather the [`FMOD_RET`](BPF_PROG_TYPE_TRACING.md#modify-return) tracing program type is repurposed. However, normally these attach to a single instance of a kernel function. For the HID-BPF use case, we want to attach to a specific HID device. This is done by using the [`hid_bpf_attach_prog`](../kfuncs/hid_bpf_attach_prog.md) kfunc to attach the program to the HID device. Which bring us to the `BPF_PROG_TYPE_SYSCALL` program which is used to actually execute this kfunc.

## Context

This program type does not have a set context type, so as long as your eBPF program and userspace are aligned, you can use any context type you want.

## Attachment

Syscall programs are never attached to any hook. They can only be executed from the [`BPF_PROG_RUN`](../syscall/BPF_PROG_TEST_RUN.md) syscall command.

## Example

!!! example "BPF-HID"
    ```c
    // SPDX-License-Identifier: GPL-2.0-only
    /* Copyright (c) 2022 Benjamin Tissoires
    */

    #include "vmlinux.h"
    #include <bpf/bpf_helpers.h>
    #include <bpf/bpf_tracing.h>
    #include "hid_bpf_helpers.h"

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
    ```

## Helper functions

Not all helper functions are available in all program types. These are the helper calls available for syscall programs:

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [`bpf_btf_find_by_name_kind`](../helper-function/bpf_btf_find_by_name_kind.md)
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
    * [`bpf_get_task_stack`](../helper-function/bpf_get_task_stack.md)
    * [`bpf_jiffies64`](../helper-function/bpf_jiffies64.md)
    * [`bpf_kallsyms_lookup_name`](../helper-function/bpf_kallsyms_lookup_name.md)
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
    * [`bpf_snprintf`](../helper-function/bpf_snprintf.md)
    * [`bpf_snprintf_btf`](../helper-function/bpf_snprintf_btf.md)
    * [`bpf_spin_lock`](../helper-function/bpf_spin_lock.md)
    * [`bpf_spin_unlock`](../helper-function/bpf_spin_unlock.md)
    * [`bpf_strncmp`](../helper-function/bpf_strncmp.md)
    * [`bpf_sys_bpf`](../helper-function/bpf_sys_bpf.md)
    * [`bpf_sys_close`](../helper-function/bpf_sys_close.md)
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
    - [`bpf_copy_from_user_str`](../kfuncs/bpf_copy_from_user_str.md)
    - [`bpf_copy_from_user_task_str`](../kfuncs/bpf_copy_from_user_task_str.md)
    - [`bpf_crypto_ctx_acquire`](../kfuncs/bpf_crypto_ctx_acquire.md)
    - [`bpf_crypto_ctx_create`](../kfuncs/bpf_crypto_ctx_create.md)
    - [`bpf_crypto_ctx_release`](../kfuncs/bpf_crypto_ctx_release.md)
    - [`bpf_dynptr_adjust`](../kfuncs/bpf_dynptr_adjust.md)
    - [`bpf_dynptr_clone`](../kfuncs/bpf_dynptr_clone.md)
    - [`bpf_dynptr_copy`](../kfuncs/bpf_dynptr_copy.md)
    - [`bpf_dynptr_is_null`](../kfuncs/bpf_dynptr_is_null.md)
    - [`bpf_dynptr_is_rdonly`](../kfuncs/bpf_dynptr_is_rdonly.md)
    - [`bpf_dynptr_size`](../kfuncs/bpf_dynptr_size.md)
    - [`bpf_dynptr_slice`](../kfuncs/bpf_dynptr_slice.md)
    - [`bpf_dynptr_slice_rdwr`](../kfuncs/bpf_dynptr_slice_rdwr.md)
    - [`bpf_get_kmem_cache`](../kfuncs/bpf_get_kmem_cache.md)
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
    - [`bpf_iter_scx_dsq_destroy`](../kfuncs/bpf_iter_scx_dsq_destroy.md)
    - [`bpf_iter_scx_dsq_new`](../kfuncs/bpf_iter_scx_dsq_new.md)
    - [`bpf_iter_scx_dsq_next`](../kfuncs/bpf_iter_scx_dsq_next.md)
    - [`bpf_iter_task_destroy`](../kfuncs/bpf_iter_task_destroy.md)
    - [`bpf_iter_task_new`](../kfuncs/bpf_iter_task_new.md)
    - [`bpf_iter_task_next`](../kfuncs/bpf_iter_task_next.md)
    - [`bpf_iter_task_vma_destroy`](../kfuncs/bpf_iter_task_vma_destroy.md)
    - [`bpf_iter_task_vma_new`](../kfuncs/bpf_iter_task_vma_new.md)
    - [`bpf_iter_task_vma_next`](../kfuncs/bpf_iter_task_vma_next.md)
    - [`bpf_local_irq_restore`](../kfuncs/bpf_local_irq_restore.md)
    - [`bpf_local_irq_save`](../kfuncs/bpf_local_irq_save.md)
    - [`bpf_map_sum_elem_count`](../kfuncs/bpf_map_sum_elem_count.md)
    - [`bpf_preempt_disable`](../kfuncs/bpf_preempt_disable.md)
    - [`bpf_preempt_enable`](../kfuncs/bpf_preempt_enable.md)
    - [`bpf_rcu_read_lock`](../kfuncs/bpf_rcu_read_lock.md)
    - [`bpf_rcu_read_unlock`](../kfuncs/bpf_rcu_read_unlock.md)
    - [`bpf_rdonly_cast`](../kfuncs/bpf_rdonly_cast.md)
    - [`bpf_res_spin_lock`](../kfuncs/bpf_res_spin_lock.md)
    - [`bpf_res_spin_lock_irqsave`](../kfuncs/bpf_res_spin_lock_irqsave.md)
    - [`bpf_res_spin_unlock`](../kfuncs/bpf_res_spin_unlock.md)
    - [`bpf_res_spin_unlock_irqrestore`](../kfuncs/bpf_res_spin_unlock_irqrestore.md)
    - [`bpf_wq_init`](../kfuncs/bpf_wq_init.md)
    - [`bpf_wq_set_callback_impl`](../kfuncs/bpf_wq_set_callback_impl.md)
    - [`bpf_wq_start`](../kfuncs/bpf_wq_start.md)
    - [`hid_bpf_allocate_context`](../kfuncs/hid_bpf_allocate_context.md)
    - [`hid_bpf_hw_output_report`](../kfuncs/hid_bpf_hw_output_report.md)
    - [`hid_bpf_hw_request`](../kfuncs/hid_bpf_hw_request.md)
    - [`hid_bpf_input_report`](../kfuncs/hid_bpf_input_report.md)
    - [`hid_bpf_release_context`](../kfuncs/hid_bpf_release_context.md)
    - [`scx_bpf_cpu_node`](../kfuncs/scx_bpf_cpu_node.md)
    - [`scx_bpf_cpu_rq`](../kfuncs/scx_bpf_cpu_rq.md)
    - [`scx_bpf_cpuperf_cap`](../kfuncs/scx_bpf_cpuperf_cap.md)
    - [`scx_bpf_cpuperf_cur`](../kfuncs/scx_bpf_cpuperf_cur.md)
    - [`scx_bpf_cpuperf_set`](../kfuncs/scx_bpf_cpuperf_set.md)
    - [`scx_bpf_create_dsq`](../kfuncs/scx_bpf_create_dsq.md)
    - [`scx_bpf_destroy_dsq`](../kfuncs/scx_bpf_destroy_dsq.md)
    - [`scx_bpf_dsq_move`](../kfuncs/scx_bpf_dsq_move.md)
    - [`scx_bpf_dsq_move_set_slice`](../kfuncs/scx_bpf_dsq_move_set_slice.md)
    - [`scx_bpf_dsq_move_set_vtime`](../kfuncs/scx_bpf_dsq_move_set_vtime.md)
    - [`scx_bpf_dsq_move_vtime`](../kfuncs/scx_bpf_dsq_move_vtime.md)
    - [`scx_bpf_dsq_nr_queued`](../kfuncs/scx_bpf_dsq_nr_queued.md)
    - [`scx_bpf_dump_bstr`](../kfuncs/scx_bpf_dump_bstr.md)
    - [`scx_bpf_error_bstr`](../kfuncs/scx_bpf_error_bstr.md)
    - [`scx_bpf_events`](../kfuncs/scx_bpf_events.md)
    - [`scx_bpf_exit_bstr`](../kfuncs/scx_bpf_exit_bstr.md)
    - [`scx_bpf_get_idle_cpumask`](../kfuncs/scx_bpf_get_idle_cpumask.md)
    - [`scx_bpf_get_idle_cpumask_node`](../kfuncs/scx_bpf_get_idle_cpumask_node.md)
    - [`scx_bpf_get_idle_smtmask`](../kfuncs/scx_bpf_get_idle_smtmask.md)
    - [`scx_bpf_get_idle_smtmask_node`](../kfuncs/scx_bpf_get_idle_smtmask_node.md)
    - [`scx_bpf_get_online_cpumask`](../kfuncs/scx_bpf_get_online_cpumask.md)
    - [`scx_bpf_get_possible_cpumask`](../kfuncs/scx_bpf_get_possible_cpumask.md)
    - [`scx_bpf_kick_cpu`](../kfuncs/scx_bpf_kick_cpu.md)
    - [`scx_bpf_now`](../kfuncs/scx_bpf_now.md)
    - [`scx_bpf_nr_cpu_ids`](../kfuncs/scx_bpf_nr_cpu_ids.md)
    - [`scx_bpf_nr_node_ids`](../kfuncs/scx_bpf_nr_node_ids.md)
    - [`scx_bpf_pick_any_cpu`](../kfuncs/scx_bpf_pick_any_cpu.md)
    - [`scx_bpf_pick_any_cpu_node`](../kfuncs/scx_bpf_pick_any_cpu_node.md)
    - [`scx_bpf_pick_idle_cpu`](../kfuncs/scx_bpf_pick_idle_cpu.md)
    - [`scx_bpf_pick_idle_cpu_node`](../kfuncs/scx_bpf_pick_idle_cpu_node.md)
    - [`scx_bpf_put_cpumask`](../kfuncs/scx_bpf_put_cpumask.md)
    - [`scx_bpf_put_idle_cpumask`](../kfuncs/scx_bpf_put_idle_cpumask.md)
    - [`scx_bpf_task_cgroup`](../kfuncs/scx_bpf_task_cgroup.md)
    - [`scx_bpf_task_cpu`](../kfuncs/scx_bpf_task_cpu.md)
    - [`scx_bpf_task_running`](../kfuncs/scx_bpf_task_running.md)
    - [`scx_bpf_test_and_clear_cpu_idle`](../kfuncs/scx_bpf_test_and_clear_cpu_idle.md)
<!-- [/PROG_KFUNC_REF] -->
