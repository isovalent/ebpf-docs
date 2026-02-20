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

The perf event program can be temporarily disabled with the `PERF_EVENT_IOC_DISABLE` ioctl option. Otherwise the perf event program stays attached until the perf_event goes away due to the closing of the perf_event FD or the program exiting. The perf event holds a reference to the BPF program so it will stay loaded until no more perf event program reference it.

### `perf_event_open` PMU

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### BPF link

This is the newest and most recommended method of attaching perf event programs. 

After we have gotten the perf event file descriptor we attach the program by making a bpf link via the [link create syscall command](../syscall/BPF_LINK_CREATE.md).

We call the syscall command with the [`BPF_PERF_EVENT`](../syscall/BPF_LINK_CREATE.md#bpf_perf_event) [`attach_type`](../syscall/BPF_LINK_CREATE.md#attach_type), [`target_fd`](../syscall/BPF_LINK_CREATE.md#target_fd) set to the perf event file descriptor, [`prog_fd`](../syscall/BPF_LINK_CREATE.md#prog_fd) to the file descriptor of the tracepoint program, and optionally a [`cookie`](../syscall/BPF_LINK_CREATE.md#tracing-cookie)


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
    * [`bpf_get_attach_cookie`](../helper-function/bpf_get_attach_cookie.md)
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
    * [`bpf_perf_prog_read_value`](../helper-function/bpf_perf_prog_read_value.md)
    * [`bpf_probe_read`](../helper-function/bpf_probe_read.md)
    * [`bpf_probe_read_kernel`](../helper-function/bpf_probe_read_kernel.md)
    * [`bpf_probe_read_kernel_str`](../helper-function/bpf_probe_read_kernel_str.md)
    * [`bpf_probe_read_str`](../helper-function/bpf_probe_read_str.md)
    * [`bpf_probe_read_user`](../helper-function/bpf_probe_read_user.md)
    * [`bpf_probe_read_user_str`](../helper-function/bpf_probe_read_user_str.md)
    * [`bpf_probe_write_user`](../helper-function/bpf_probe_write_user.md)
    * [`bpf_read_branch_records`](../helper-function/bpf_read_branch_records.md)
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
    * [`bpf_strncmp`](../helper-function/bpf_strncmp.md)
    * [`bpf_tail_call`](../helper-function/bpf_tail_call.md)
    * [`bpf_task_pt_regs`](../helper-function/bpf_task_pt_regs.md)
    * [`bpf_task_storage_delete`](../helper-function/bpf_task_storage_delete.md)
    * [`bpf_task_storage_get`](../helper-function/bpf_task_storage_get.md)
    * [`bpf_this_cpu_ptr`](../helper-function/bpf_this_cpu_ptr.md)
    * [`bpf_trace_printk`](../helper-function/bpf_trace_printk.md)
    * [`bpf_trace_vprintk`](../helper-function/bpf_trace_vprintk.md)
    * [`bpf_user_ringbuf_drain`](../helper-function/bpf_user_ringbuf_drain.md)
<!-- [/PROG_HELPER_FUNC_REF] -->

## KFuncs

<!-- [PROG_KFUNC_REF] -->
??? abstract "Supported kfuncs"
    - [`__bpf_trap`](../kfuncs/__bpf_trap.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`__scx_bpf_select_cpu_and`](../kfuncs/__scx_bpf_select_cpu_and.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_arena_alloc_pages`](../kfuncs/bpf_arena_alloc_pages.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_arena_free_pages`](../kfuncs/bpf_arena_free_pages.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_arena_reserve_pages`](../kfuncs/bpf_arena_reserve_pages.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cast_to_kern_ctx`](../kfuncs/bpf_cast_to_kern_ctx.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cgroup_acquire`](../kfuncs/bpf_cgroup_acquire.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cgroup_ancestor`](../kfuncs/bpf_cgroup_ancestor.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cgroup_from_id`](../kfuncs/bpf_cgroup_from_id.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cgroup_read_xattr`](../kfuncs/bpf_cgroup_read_xattr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cgroup_release`](../kfuncs/bpf_cgroup_release.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_copy_from_user_dynptr`](../kfuncs/bpf_copy_from_user_dynptr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_copy_from_user_str`](../kfuncs/bpf_copy_from_user_str.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_copy_from_user_str_dynptr`](../kfuncs/bpf_copy_from_user_str_dynptr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_copy_from_user_task_dynptr`](../kfuncs/bpf_copy_from_user_task_dynptr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_copy_from_user_task_str`](../kfuncs/bpf_copy_from_user_task_str.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_copy_from_user_task_str_dynptr`](../kfuncs/bpf_copy_from_user_task_str_dynptr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_acquire`](../kfuncs/bpf_cpumask_acquire.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_and`](../kfuncs/bpf_cpumask_and.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_any_and_distribute`](../kfuncs/bpf_cpumask_any_and_distribute.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_any_distribute`](../kfuncs/bpf_cpumask_any_distribute.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_clear`](../kfuncs/bpf_cpumask_clear.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_clear_cpu`](../kfuncs/bpf_cpumask_clear_cpu.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_copy`](../kfuncs/bpf_cpumask_copy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_create`](../kfuncs/bpf_cpumask_create.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_empty`](../kfuncs/bpf_cpumask_empty.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_equal`](../kfuncs/bpf_cpumask_equal.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_first`](../kfuncs/bpf_cpumask_first.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_first_and`](../kfuncs/bpf_cpumask_first_and.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_first_zero`](../kfuncs/bpf_cpumask_first_zero.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_full`](../kfuncs/bpf_cpumask_full.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_intersects`](../kfuncs/bpf_cpumask_intersects.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_or`](../kfuncs/bpf_cpumask_or.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_populate`](../kfuncs/bpf_cpumask_populate.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_release`](../kfuncs/bpf_cpumask_release.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_set_cpu`](../kfuncs/bpf_cpumask_set_cpu.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_setall`](../kfuncs/bpf_cpumask_setall.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_subset`](../kfuncs/bpf_cpumask_subset.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_test_and_clear_cpu`](../kfuncs/bpf_cpumask_test_and_clear_cpu.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_test_and_set_cpu`](../kfuncs/bpf_cpumask_test_and_set_cpu.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_test_cpu`](../kfuncs/bpf_cpumask_test_cpu.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_weight`](../kfuncs/bpf_cpumask_weight.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cpumask_xor`](../kfuncs/bpf_cpumask_xor.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_dynptr_adjust`](../kfuncs/bpf_dynptr_adjust.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_dynptr_clone`](../kfuncs/bpf_dynptr_clone.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_dynptr_copy`](../kfuncs/bpf_dynptr_copy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_dynptr_file_discard`](../kfuncs/bpf_dynptr_file_discard.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_dynptr_from_file`](../kfuncs/bpf_dynptr_from_file.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_dynptr_from_skb`](../kfuncs/bpf_dynptr_from_skb.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_dynptr_is_null`](../kfuncs/bpf_dynptr_is_null.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_dynptr_is_rdonly`](../kfuncs/bpf_dynptr_is_rdonly.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_dynptr_memset`](../kfuncs/bpf_dynptr_memset.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_dynptr_size`](../kfuncs/bpf_dynptr_size.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_dynptr_slice`](../kfuncs/bpf_dynptr_slice.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_dynptr_slice_rdwr`](../kfuncs/bpf_dynptr_slice_rdwr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_get_dentry_xattr`](../kfuncs/bpf_get_dentry_xattr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_get_file_xattr`](../kfuncs/bpf_get_file_xattr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_get_fsverity_digest`](../kfuncs/bpf_get_fsverity_digest.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_get_kmem_cache`](../kfuncs/bpf_get_kmem_cache.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_get_task_exe_file`](../kfuncs/bpf_get_task_exe_file.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_bits_destroy`](../kfuncs/bpf_iter_bits_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_bits_new`](../kfuncs/bpf_iter_bits_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_bits_next`](../kfuncs/bpf_iter_bits_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_css_destroy`](../kfuncs/bpf_iter_css_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_css_new`](../kfuncs/bpf_iter_css_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_css_next`](../kfuncs/bpf_iter_css_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_css_task_destroy`](../kfuncs/bpf_iter_css_task_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_css_task_new`](../kfuncs/bpf_iter_css_task_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_css_task_next`](../kfuncs/bpf_iter_css_task_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_dmabuf_destroy`](../kfuncs/bpf_iter_dmabuf_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_dmabuf_new`](../kfuncs/bpf_iter_dmabuf_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_dmabuf_next`](../kfuncs/bpf_iter_dmabuf_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_kmem_cache_destroy`](../kfuncs/bpf_iter_kmem_cache_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_kmem_cache_new`](../kfuncs/bpf_iter_kmem_cache_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_kmem_cache_next`](../kfuncs/bpf_iter_kmem_cache_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_num_destroy`](../kfuncs/bpf_iter_num_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_num_new`](../kfuncs/bpf_iter_num_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_num_next`](../kfuncs/bpf_iter_num_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_scx_dsq_destroy`](../kfuncs/bpf_iter_scx_dsq_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_scx_dsq_new`](../kfuncs/bpf_iter_scx_dsq_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_scx_dsq_next`](../kfuncs/bpf_iter_scx_dsq_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_task_destroy`](../kfuncs/bpf_iter_task_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_task_new`](../kfuncs/bpf_iter_task_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_task_next`](../kfuncs/bpf_iter_task_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_task_vma_destroy`](../kfuncs/bpf_iter_task_vma_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_task_vma_new`](../kfuncs/bpf_iter_task_vma_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_iter_task_vma_next`](../kfuncs/bpf_iter_task_vma_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_key_put`](../kfuncs/bpf_key_put.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_local_irq_restore`](../kfuncs/bpf_local_irq_restore.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_local_irq_save`](../kfuncs/bpf_local_irq_save.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_lookup_system_key`](../kfuncs/bpf_lookup_system_key.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_lookup_user_key`](../kfuncs/bpf_lookup_user_key.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_map_sum_elem_count`](../kfuncs/bpf_map_sum_elem_count.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_obj_drop_impl`](../kfuncs/bpf_obj_drop_impl.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_obj_new_impl`](../kfuncs/bpf_obj_new_impl.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_path_d_path`](../kfuncs/bpf_path_d_path.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_percpu_obj_drop_impl`](../kfuncs/bpf_percpu_obj_drop_impl.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_percpu_obj_new_impl`](../kfuncs/bpf_percpu_obj_new_impl.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_preempt_disable`](../kfuncs/bpf_preempt_disable.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_preempt_enable`](../kfuncs/bpf_preempt_enable.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_probe_read_kernel_dynptr`](../kfuncs/bpf_probe_read_kernel_dynptr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_probe_read_kernel_str_dynptr`](../kfuncs/bpf_probe_read_kernel_str_dynptr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_probe_read_user_dynptr`](../kfuncs/bpf_probe_read_user_dynptr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_probe_read_user_str_dynptr`](../kfuncs/bpf_probe_read_user_str_dynptr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_put_file`](../kfuncs/bpf_put_file.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_rbtree_remove`](../kfuncs/bpf_rbtree_remove.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_rcu_read_lock`](../kfuncs/bpf_rcu_read_lock.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_rcu_read_unlock`](../kfuncs/bpf_rcu_read_unlock.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_rdonly_cast`](../kfuncs/bpf_rdonly_cast.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_refcount_acquire_impl`](../kfuncs/bpf_refcount_acquire_impl.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_remove_dentry_xattr`](../kfuncs/bpf_remove_dentry_xattr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_send_signal_task`](../kfuncs/bpf_send_signal_task.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_set_dentry_xattr`](../kfuncs/bpf_set_dentry_xattr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_sock_destroy`](../kfuncs/bpf_sock_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_strcasecmp`](../kfuncs/bpf_strcasecmp.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_strcasestr`](../kfuncs/bpf_strcasestr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_strchr`](../kfuncs/bpf_strchr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_strchrnul`](../kfuncs/bpf_strchrnul.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_strcmp`](../kfuncs/bpf_strcmp.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_strcspn`](../kfuncs/bpf_strcspn.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_stream_vprintk_impl`](../kfuncs/bpf_stream_vprintk_impl.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_strlen`](../kfuncs/bpf_strlen.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_strncasestr`](../kfuncs/bpf_strncasestr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_strnchr`](../kfuncs/bpf_strnchr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_strnlen`](../kfuncs/bpf_strnlen.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_strnstr`](../kfuncs/bpf_strnstr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_strrchr`](../kfuncs/bpf_strrchr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_strspn`](../kfuncs/bpf_strspn.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_strstr`](../kfuncs/bpf_strstr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_task_acquire`](../kfuncs/bpf_task_acquire.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_task_from_pid`](../kfuncs/bpf_task_from_pid.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_task_from_vpid`](../kfuncs/bpf_task_from_vpid.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_task_get_cgroup1`](../kfuncs/bpf_task_get_cgroup1.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_task_release`](../kfuncs/bpf_task_release.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_task_under_cgroup`](../kfuncs/bpf_task_under_cgroup.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_task_work_schedule_resume_impl`](../kfuncs/bpf_task_work_schedule_resume_impl.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_task_work_schedule_signal_impl`](../kfuncs/bpf_task_work_schedule_signal_impl.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_throw`](../kfuncs/bpf_throw.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_verify_pkcs7_signature`](../kfuncs/bpf_verify_pkcs7_signature.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`crash_kexec`](../kfuncs/crash_kexec.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`css_rstat_flush`](../kfuncs/css_rstat_flush.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`css_rstat_updated`](../kfuncs/css_rstat_updated.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_cpu_curr`](../kfuncs/scx_bpf_cpu_curr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_cpu_node`](../kfuncs/scx_bpf_cpu_node.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_cpu_rq`](../kfuncs/scx_bpf_cpu_rq.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_cpuperf_cap`](../kfuncs/scx_bpf_cpuperf_cap.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_cpuperf_cur`](../kfuncs/scx_bpf_cpuperf_cur.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_cpuperf_set`](../kfuncs/scx_bpf_cpuperf_set.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_destroy_dsq`](../kfuncs/scx_bpf_destroy_dsq.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_dsq_nr_queued`](../kfuncs/scx_bpf_dsq_nr_queued.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_dsq_peek`](../kfuncs/scx_bpf_dsq_peek.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_dump_bstr`](../kfuncs/scx_bpf_dump_bstr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_error_bstr`](../kfuncs/scx_bpf_error_bstr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_events`](../kfuncs/scx_bpf_events.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_exit_bstr`](../kfuncs/scx_bpf_exit_bstr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_get_idle_cpumask`](../kfuncs/scx_bpf_get_idle_cpumask.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_get_idle_cpumask_node`](../kfuncs/scx_bpf_get_idle_cpumask_node.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_get_idle_smtmask`](../kfuncs/scx_bpf_get_idle_smtmask.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_get_idle_smtmask_node`](../kfuncs/scx_bpf_get_idle_smtmask_node.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_get_online_cpumask`](../kfuncs/scx_bpf_get_online_cpumask.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_get_possible_cpumask`](../kfuncs/scx_bpf_get_possible_cpumask.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_kick_cpu`](../kfuncs/scx_bpf_kick_cpu.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_locked_rq`](../kfuncs/scx_bpf_locked_rq.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_now`](../kfuncs/scx_bpf_now.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_nr_cpu_ids`](../kfuncs/scx_bpf_nr_cpu_ids.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_nr_node_ids`](../kfuncs/scx_bpf_nr_node_ids.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_pick_any_cpu`](../kfuncs/scx_bpf_pick_any_cpu.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_pick_any_cpu_node`](../kfuncs/scx_bpf_pick_any_cpu_node.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_pick_idle_cpu`](../kfuncs/scx_bpf_pick_idle_cpu.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_pick_idle_cpu_node`](../kfuncs/scx_bpf_pick_idle_cpu_node.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_put_cpumask`](../kfuncs/scx_bpf_put_cpumask.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_put_idle_cpumask`](../kfuncs/scx_bpf_put_idle_cpumask.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_reenqueue_local___v2`](../kfuncs/scx_bpf_reenqueue_local___v2.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_select_cpu_and`](../kfuncs/scx_bpf_select_cpu_and.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_task_cgroup`](../kfuncs/scx_bpf_task_cgroup.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_task_cpu`](../kfuncs/scx_bpf_task_cpu.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_task_running`](../kfuncs/scx_bpf_task_running.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_task_set_dsq_vtime`](../kfuncs/scx_bpf_task_set_dsq_vtime.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_task_set_slice`](../kfuncs/scx_bpf_task_set_slice.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_test_and_clear_cpu_idle`](../kfuncs/scx_bpf_test_and_clear_cpu_idle.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
<!-- [/PROG_KFUNC_REF] -->
