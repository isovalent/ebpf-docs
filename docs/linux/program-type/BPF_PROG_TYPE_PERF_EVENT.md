---
title: "Program Type 'BPF_PROG_TYPE_PERF_EVENT'"
description: "This page documents the 'BPF_PROG_TYPE_PERF_EVENT' eBPF program type, including its definition, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_PERF_EVENT`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_PERF_EVENT) -->
[:octicons-tag-24: v4.9](https://github.com/torvalds/linux/commit/0515e5999a466dfe6e1924f460da599bb6821487)
<!-- [/FEATURE_TAG] -->

Perf event programs that can be attached to hardware and software perf events. Once attached the BPF program is executed each time the perf event is triggered. 

## Usage

Perf event programs are typically used for profiling and tracing. These programs are called with the CPU register state at the time of the event. This allows the programs to collect information for each event and aggregate it in a customized way. 

Perf event programs are typically placed in the `perf_event` ELF header.

## Context

??? abstract "C Structure"
    ```c
    struct bpf_perf_event_data {
        bpf_user_pt_regs_t regs;
        __u64 sample_period;
        __u64 addr;
    };
    ```

### `regs`

This field contains the CPU registers at the time of the event. The type of the field is different for each architecture since each architecture has different registers. The helpers in `tools/lib/bpf/bpf_tracing.h` can be used to access the registers in a portable way.

### `sample_period`

This field contains the amount of times this perf even has been triggered.

### `addr`

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

## Attachment

here are three methods of attaching perf event programs, from oldest and least recommended to newest and most recommended, however, all methods have this first part in common. 

Next step is to open a new perf event using the [`perf_event_open`](https://man7.org/linux/man-pages/man2/perf_event_open.2.html) syscall:

```c
struct perf_event_attr attr = {
    .sample_freq = SAMPLE_FREQ,
    .freq = 1,
    .type = PERF_TYPE_HARDWARE,
    .config = PERF_COUNT_HW_CPU_CYCLES,
};

syscall(SYS_perf_event_open, 
    &attr,  /* struct perf_event_attr * */
    -1,     /* pid_t pid */
    0       /* int cpu */
    -1,     /* int group_fd */
    PERF_FLAG_FD_CLOEXEC /* unsigned long flags */
);
```

This syscall will return a file descriptor on success. Perf event programs can be attached to any event, as long as it is of type `PERF_TYPE_HARDWARE` or `PERF_TYPE_SOFTWARE`.

### ioctl method

This is the oldest and least recommended method. After we have the perf event file descriptor we execute two [`ioctl`](https://man7.org/linux/man-pages/man2/ioctl.2.html) syscalls to attach our BPF program to the trace event and to enable the trace.

`#!c ioctl(perf_event_fd, PERF_EVENT_IOC_SET_BPF, bpf_prog_fd);` to attach.

`#!c ioctl(perf_event_fd, PERF_EVENT_IOC_ENABLE, 0);` to enable.

The perf event program can be temporality disabled with the `PERF_EVENT_IOC_DISABLE` ioctl option. Otherwise the perf event program stays attached until the perf_event goes away due to the closing of the perf_event FD or the program exiting. The perf event holds a reference to the BPF program so it will stay loaded until no more perf event program reference it.

### `perf_event_open` PMU

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### BPF link

This is the newest and most recommended method of attaching perf event programs. 

After we have gotten the perf event file descriptor we attach the program by making a bpf link via the [link create syscall command](../syscall/BPF_LINK_CREATE.md).

We call the syscall command with the [`BPF_PERF_EVENT`](../syscall/BPF_LINK_CREATE.md#bpf_perf_event) [`attach_type`](../syscall/BPF_LINK_CREATE.md#attach_type), [`target_fd`](../syscall/BPF_LINK_CREATE.md#target_fd) set to the perf event fd, [`prog_fd`](../syscall/BPF_LINK_CREATE.md#prog_fd) to the file descriptor of the tracepoint program, and optionally a [`cookie`](../syscall/BPF_LINK_CREATE.md#cookie)


## Examples

??? Example "profiling example"
    ```c
    /* Copyright (c) 2016 Facebook
    *
    * This program is free software; you can redistribute it and/or
    * modify it under the terms of version 2 of the GNU General Public
    * License as published by the Free Software Foundation.
    */
    #include <linux/ptrace.h>
    #include <uapi/linux/bpf.h>
    #include <uapi/linux/bpf_perf_event.h>
    #include <uapi/linux/perf_event.h>
    #include <bpf/bpf_helpers.h>
    #include <bpf/bpf_tracing.h>

    struct key_t {
        char comm[TASK_COMM_LEN];
        u32 kernstack;
        u32 userstack;
    };

    struct {
        __uint(type, BPF_MAP_TYPE_HASH);
        __type(key, struct key_t);
        __type(value, u64);
        __uint(max_entries, 10000);
    } counts SEC(".maps");

    struct {
        __uint(type, BPF_MAP_TYPE_STACK_TRACE);
        __uint(key_size, sizeof(u32));
        __uint(value_size, PERF_MAX_STACK_DEPTH * sizeof(u64));
        __uint(max_entries, 10000);
    } stackmap SEC(".maps");

    #define KERN_STACKID_FLAGS (0 | BPF_F_FAST_STACK_CMP)
    #define USER_STACKID_FLAGS (0 | BPF_F_FAST_STACK_CMP | BPF_F_USER_STACK)

    SEC("perf_event")
    int bpf_prog1(struct bpf_perf_event_data *ctx)
    {
        char time_fmt1[] = "Time Enabled: %llu, Time Running: %llu";
        char time_fmt2[] = "Get Time Failed, ErrCode: %d";
        char addr_fmt[] = "Address recorded on event: %llx";
        char fmt[] = "CPU-%d period %lld ip %llx";
        u32 cpu = bpf_get_smp_processor_id();
        struct bpf_perf_event_value value_buf;
        struct key_t key;
        u64 *val, one = 1;
        int ret;

        if (ctx->sample_period < 10000)
            /* ignore warmup */
            return 0;
        bpf_get_current_comm(&key.comm, sizeof(key.comm));
        key.kernstack = bpf_get_stackid(ctx, &stackmap, KERN_STACKID_FLAGS);
        key.userstack = bpf_get_stackid(ctx, &stackmap, USER_STACKID_FLAGS);
        if ((int)key.kernstack < 0 && (int)key.userstack < 0) {
            bpf_trace_printk(fmt, sizeof(fmt), cpu, ctx->sample_period,
                    PT_REGS_IP(&ctx->regs));
            return 0;
        }

        ret = bpf_perf_prog_read_value(ctx, (void *)&value_buf, sizeof(struct bpf_perf_event_value));
        if (!ret)
        bpf_trace_printk(time_fmt1, sizeof(time_fmt1), value_buf.enabled, value_buf.running);
        else
        bpf_trace_printk(time_fmt2, sizeof(time_fmt2), ret);

        if (ctx->addr != 0)
        bpf_trace_printk(addr_fmt, sizeof(addr_fmt), ctx->addr);

        val = bpf_map_lookup_elem(&counts, &key);
        if (val)
            (*val)++;
        else
            bpf_map_update_elem(&counts, &key, &one, BPF_NOEXIST);
        return 0;
    }

    char _license[] SEC("license") = "GPL";
    ```

??? example "recording instruction pointer"
    ```c
    /* Copyright 2016 Netflix, Inc.
    *
    * This program is free software; you can redistribute it and/or
    * modify it under the terms of version 2 of the GNU General Public
    * License as published by the Free Software Foundation.
    */
    #include <linux/ptrace.h>
    #include <uapi/linux/bpf.h>
    #include <uapi/linux/bpf_perf_event.h>
    #include <bpf/bpf_helpers.h>
    #include <bpf/bpf_tracing.h>

    #define MAX_IPS		8192

    struct {
        __uint(type, BPF_MAP_TYPE_HASH);
        __type(key, u64);
        __type(value, u32);
        __uint(max_entries, MAX_IPS);
    } ip_map SEC(".maps");

    SEC("perf_event")
    int do_sample(struct bpf_perf_event_data *ctx)
    {
        u64 ip;
        u32 *value, init_val = 1;

        ip = PT_REGS_IP(&ctx->regs);
        value = bpf_map_lookup_elem(&ip_map, &ip);
        if (value)
            *value += 1;
        else
            /* E2BIG not tested for this example only */
            bpf_map_update_elem(&ip_map, &ip, &init_val, BPF_NOEXIST);

        return 0;
    }
    char _license[] SEC("license") = "GPL";
    ```

## Helper functions

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [bpf_perf_event_output](../helper-function/bpf_perf_event_output.md)
    * [bpf_get_stackid](../helper-function/bpf_get_stackid.md)
    * [bpf_get_stack](../helper-function/bpf_get_stack.md)
    * [bpf_perf_prog_read_value](../helper-function/bpf_perf_prog_read_value.md)
    * [bpf_read_branch_records](../helper-function/bpf_read_branch_records.md)
    * [bpf_get_attach_cookie](../helper-function/bpf_get_attach_cookie.md)
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

## KFuncs

<!-- [PROG_KFUNC_REF] -->
There are currently no kfuncs supported for this program type
<!-- [/PROG_KFUNC_REF] -->
