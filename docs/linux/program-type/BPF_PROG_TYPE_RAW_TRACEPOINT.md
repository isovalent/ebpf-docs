---
title: "Program Type 'BPF_PROG_TYPE_RAW_TRACEPOINT'"
description: "This page documents the 'BPF_PROG_TYPE_RAW_TRACEPOINT' eBPF program type, including its defintion, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_RAW_TRACEPOINT`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_RAW_TRACEPOINT) -->
[:octicons-tag-24: v4.17](https://github.com/torvalds/linux/commit/c4f6699dfcb8558d138fe838f741b2c10f416cf9)
<!-- [/FEATURE_TAG] -->

Raw tracepoint programs are similar to [tracepoint programs](BPF_PROG_TYPE_TRACEPOINT.md), but the kernel does no pre-processing on the arguments and passes the raw arguments directly to the tracepoint program.

## Usage

Raw tracepoint programs are typically put into an [ELF](../../elf.md) section prefixed with `raw_tp/` or in a `raw_tracepoint` section. When loading as a `BPF_PROG_TYPE_TRACING` program, the raw tracepoint is typically located in a section prefixed with `tp_btf/`.

Raw tracepoints are attached to the same tracepoints as normal tracepoint programs. The reason why you might want to use raw tracepoints over normal tracepoints is due to the performance improvement. For normal tracepoints, the kernel will cast or transform arguments even if the arguments are never used. By taking the raw arguments, the BPF program can do the casting or transformation only if the arguments are used, thereby making a more efficient tracepoint program.

## Context

The context for raw tracepoint programs is a pointer to a `struct bpf_raw_tracepoint_args`:

```c
struct bpf_raw_tracepoint_args {
       __u64 args[0];
};
```

The `args` array contains the raw arguments to the tracepoint. The number of arguments is determined by the tracepoint. The verifier will enforce that the number of arguments matches the number of arguments expected by the tracepoint. The BPF program can cast the u64 values to the expected types or use the [`bpf_probe_read`](../helper-function/bpf_probe_read.md)/[`bpf_probe_read_kernel`](../helper-function/bpf_probe_read_kernel.md) helper function to read the arguments.

## Attachment

Raw tracepoints can be attached in two ways, first is with a dedicated syscall, the second method is with the more generic [BPF link](../syscall/BPF_LINK_CREATE.md) syscall.

### Syscall

The dedicated syscall `BPF_RAW_TRACEPOINT_OPEN` can be used to attach the raw tracepoint. This requires the `name` field to be set to a string containing the name of the tracepoint to which the user whishes to attach to. The `prog_fd` attribute field should be set to the file descriptor of the BPF program to attach.

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
    <!-- An libbpf example would be nice -->

### BPF link

<!-- [FEATURE_TAG](BPF_PROG_TYPE_TRACING) -->
[:octicons-tag-24: v5.5](https://github.com/torvalds/linux/commit/f1b9509c2fb0ef4db8d22dac9aef8e856a5d81f6)
<!-- [/FEATURE_TAG] -->

A BPF link can also be used to attach a raw tracepoint program. To do so the raw tracepoint must be loaded with `BPF_PROG_TYPE_TRACING` program type instead of the `BPF_PROG_TYPE_RAW_TRACEPOINT` program type. The `expected_attach_type` should be `BPF_TRACE_RAW_TP` and the `attach_btf_id` attribute set to the BTF ID of the tracepoint the program should be attached to.

After that a link should be created via the [link create syscall command](../syscall/BPF_LINK_CREATE.md) syscall. The attach type set to `BPF_TRACE_RAW_TP`.

## Example

??? example "raw tracepoint"
    ```c
    // SPDX-License-Identifier: (GPL-2.0-only OR BSD-2-Clause)
    // Copyright (c) 2021 Facebook
    #include "vmlinux.h"
    #include <bpf/bpf_helpers.h>
    #include <bpf/bpf_tracing.h>

    struct {
        __uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
        __uint(key_size, sizeof(__u32));
        __uint(value_size, sizeof(int));
        __uint(map_flags, BPF_F_PRESERVE_ELEMS);
    } events SEC(".maps");

    struct {
        __uint(type, BPF_MAP_TYPE_PERCPU_ARRAY);
        __uint(key_size, sizeof(__u32));
        __uint(value_size, sizeof(struct bpf_perf_event_value));
        __uint(max_entries, 1);
    } prev_readings SEC(".maps");

    struct {
        __uint(type, BPF_MAP_TYPE_PERCPU_ARRAY);
        __uint(key_size, sizeof(__u32));
        __uint(value_size, sizeof(struct bpf_perf_event_value));
        __uint(max_entries, 1);
    } diff_readings SEC(".maps");

    SEC("raw_tp/sched_switch")
    int BPF_PROG(on_switch)
    {
        struct bpf_perf_event_value val, *prev_val, *diff_val;
        __u32 key = bpf_get_smp_processor_id();
        __u32 zero = 0;
        long err;

        prev_val = bpf_map_lookup_elem(&prev_readings, &zero);
        if (!prev_val)
            return 0;

        diff_val = bpf_map_lookup_elem(&diff_readings, &zero);
        if (!diff_val)
            return 0;

        err = bpf_perf_event_read_value(&events, key, &val, sizeof(val));
        if (err)
            return 0;

        diff_val->counter = val.counter - prev_val->counter;
        diff_val->enabled = val.enabled - prev_val->enabled;
        diff_val->running = val.running - prev_val->running;
        *prev_val = val;
        return 0;
    }

    char LICENSE[] SEC("license") = "Dual BSD/GPL";
    ```

??? example "tracing program"
    ```c
    // SPDX-License-Identifier: GPL-2.0
    // Copyright (c) 2019 Facebook
    #include "vmlinux.h"
    #include <bpf/bpf_helpers.h>
    #include "runqslower.h"

    #define TASK_RUNNING 0
    #define BPF_F_CURRENT_CPU 0xffffffffULL

    const volatile __u64 min_us = 0;
    const volatile pid_t targ_pid = 0;

    struct {
        __uint(type, BPF_MAP_TYPE_TASK_STORAGE);
        __uint(map_flags, BPF_F_NO_PREALLOC);
        __type(key, int);
        __type(value, u64);
    } start SEC(".maps");

    struct {
        __uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
        __uint(key_size, sizeof(u32));
        __uint(value_size, sizeof(u32));
    } events SEC(".maps");

    /* record enqueue timestamp */
    __always_inline
    static int trace_enqueue(struct task_struct *t)
    {
        u32 pid = t->pid;
        u64 *ptr;

        if (!pid || (targ_pid && targ_pid != pid))
            return 0;

        ptr = bpf_task_storage_get(&start, t, 0,
                    BPF_LOCAL_STORAGE_GET_F_CREATE);
        if (!ptr)
            return 0;

        *ptr = bpf_ktime_get_ns();
        return 0;
    }

    SEC("tp_btf/sched_wakeup")
    int handle__sched_wakeup(u64 *ctx)
    {
        /* TP_PROTO(struct task_struct *p) */
        struct task_struct *p = (void *)ctx[0];

        return trace_enqueue(p);
    }

    SEC("tp_btf/sched_wakeup_new")
    int handle__sched_wakeup_new(u64 *ctx)
    {
        /* TP_PROTO(struct task_struct *p) */
        struct task_struct *p = (void *)ctx[0];

        return trace_enqueue(p);
    }

    SEC("tp_btf/sched_switch")
    int handle__sched_switch(u64 *ctx)
    {
        /* TP_PROTO(bool preempt, struct task_struct *prev,
        *	    struct task_struct *next)
        */
        struct task_struct *prev = (struct task_struct *)ctx[1];
        struct task_struct *next = (struct task_struct *)ctx[2];
        struct runq_event event = {};
        u64 *tsp, delta_us;
        long state;
        u32 pid;

        /* ivcsw: treat like an enqueue event and store timestamp */
        if (prev->__state == TASK_RUNNING)
            trace_enqueue(prev);

        pid = next->pid;

        /* For pid mismatch, save a bpf_task_storage_get */
        if (!pid || (targ_pid && targ_pid != pid))
            return 0;

        /* fetch timestamp and calculate delta */
        tsp = bpf_task_storage_get(&start, next, 0, 0);
        if (!tsp)
            return 0;   /* missed enqueue */

        delta_us = (bpf_ktime_get_ns() - *tsp) / 1000;
        if (min_us && delta_us <= min_us)
            return 0;

        event.pid = pid;
        event.delta_us = delta_us;
        bpf_get_current_comm(&event.task, sizeof(event.task));

        /* output */
        bpf_perf_event_output(ctx, &events, BPF_F_CURRENT_CPU,
                    &event, sizeof(event));

        bpf_task_storage_delete(&start, next);
        return 0;
    }

    char LICENSE[] SEC("license") = "GPL";
    ```

## Helper functions

Not all helper functions are available in all program types. These are the helper calls available for raw tracepoint programs:

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [bpf_perf_event_output](../helper-function/bpf_perf_event_output.md)
    * [bpf_get_stackid](../helper-function/bpf_get_stackid.md)
    * [bpf_get_stack](../helper-function/bpf_get_stack.md)
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
