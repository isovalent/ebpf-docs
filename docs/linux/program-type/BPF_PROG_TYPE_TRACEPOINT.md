---
title: "Program Type 'BPF_PROG_TYPE_TRACEPOINT'"
description: "This page documents the 'BPF_PROG_TYPE_TRACEPOINT' eBPF program type, including its definition, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_TRACEPOINT`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_TRACEPOINT) -->
[:octicons-tag-24: v4.7](https://github.com/torvalds/linux/commit/98b5c2c65c2951772a8fc661f50d675e450e8bce)
<!-- [/FEATURE_TAG] -->

`BPF_PROG_TYPE_TRACEPOINT` programs are eBPF programs that attach to pre-defined trace points in the linux kernel. These tracepoint are often placed in locations which are interesting or common locations to measure performance.

## Usage

Tracepoint programs can attach to trace events. These events are declared with the [`TRACE_EVENT`](https://elixir.bootlin.com/linux/v6.2.2/source/include/linux/tracepoint.h#L436) macro. Take for example the [`xdp_exception`](https://elixir.bootlin.com/linux/v6.2.2/source/include/trace/events/xdp.h#L28) trace event. With a combination of `TP_*` macros a function prototype for the tracepoint is defined, a structure which will be passed to any handlers and a conversion method for going from the arguments to the structure. 

The `TRACE_EVENT` macro will make a tracepoint available via a function with the `trace_` prefix followed by the name. So `trace_xdp_exception` will fire the `xdp_exception` event, which can happen from any number of locations in the code. The attached eBPF program will be called for all invocations of the trace program.

We can use the [`tracefs`](https://www.kernel.org/doc/Documentation/trace/ftrace.txt) to list all of these available trace events. For the sake of this page we will assume the `tracefs` is mounted at `/sys/kernel/tracing` (which is usual for most distros). The `/sys/kernel/tracing/events/` directory contains a number of yet more directories. The events are grouped by the first word in their name, so all `kvm_*` events reside in `/sys/kernel/tracing/events/kvm`. So `xdp_exception` is located in `/sys/kernel/tracing/events/xdp/xdp_exception`. We will refer to this directory as the "event directory".

## Context

The context for a tracepoint program is a pointer to a structure, the type of which is different for each trace event. The event directory contains a pseudo-file called `format` so for `xdp_exception` that would be `/sys/kernel/tracing/events/xdp/xdp_exception/format`. We can read this file to get the layout of the struct type:

`#!bash $ cat /sys/kernel/tracing/events/xdp/xdp_exception/format`
```
name: xdp_exception
ID: 488
format:
	field:unsigned short common_type;	offset:0;	size:2;	signed:0;
	field:unsigned char common_flags;	offset:2;	size:1;	signed:0;
	field:unsigned char common_preempt_count;	offset:3;	size:1;	signed:0;
	field:int common_pid;	offset:4;	size:4;	signed:1;

	field:int prog_id;	offset:8;	size:4;	signed:1;
	field:u32 act;	offset:12;	size:4;	signed:0;
	field:int ifindex;	offset:16;	size:4;	signed:1;

print fmt: "prog_id=%d action=%s ifindex=%d", REC->prog_id, __print_symbolic(REC->act, { 0, "ABORTED" }, { 1, "DROP" }, { 2, "PASS" }, { 3, "TX" }, { 4, "REDIRECT" }, { -1, ((void *)0) }), REC->ifindex
```

From this output we can reconstruct the context, which as C struct would look like:

```c
struct xdp_exception_ctx {
    __u16 common_type;
    __u8 flags;
    __u8 common_preempt_count;
    __s32 common_pid;

    __s32 prog_int;
    __u32 act;
    __s32 ifindex;
};
```

## Attachment

There are three methods of attaching tracepoint programs, from oldest and least recommended to newest and most recommended, however, all methods have this first part in common. 

We start by looking up the event ID in the `tracefs`. Inside the event directory is located a pseudo-file called `id`, so for `xdp_exception` that would be `/sys/kernel/tracing/events/xdp/xdp_exception/id`. When reading the file a decimal number is returned.

Next step is to open a new perf event using the [`perf_event_open`](https://man7.org/linux/man-pages/man2/perf_event_open.2.html) syscall:

```c
struct perf_event_attr attr = {
    .type = PERF_TYPE_TRACEPOINT,
    .size = sizeof(struct perf_event_attr),
    .config = event_id, /* The ID of your trace event */
    .sample_period = 1,
    .sample_type = PERF_SAMPLE_RAW,
    .wakeup_events = 1,
};

syscall(SYS_perf_event_open, 
    &attr,  /* struct perf_event_attr * */
    -1,     /* pid_t pid */
    0       /* int cpu */
    -1,     /* int group_fd */
    PERF_FLAG_FD_CLOEXEC /* unsigned long flags */
);
```

This syscall will return a file descriptor on success. 

### ioctl method

This is the oldest and least recommended method. After we have the perf event file descriptor we execute two [`ioctl`](https://man7.org/linux/man-pages/man2/ioctl.2.html) syscalls to attach our BPF program to the trace event and to enable the trace.

`#!c ioctl(perf_event_fd, PERF_EVENT_IOC_SET_BPF, bpf_prog_fd);` to attach.

`#!c ioctl(perf_event_fd, PERF_EVENT_IOC_ENABLE, 0);` to enable.

The tracepoint can be temporality disabled with the `PERF_EVENT_IOC_DISABLE` ioctl option. Otherwise the tracepoint stays attached until the perf_event goes away due to the closing of the perf_event FD or the program exiting. The perf event holds a reference to the BPF program so it will stay loaded until no more tracepoint reference it.

### `perf_event_open` PMU

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### BPF link

This is the newest and most recommended method of attaching tracepoint programs. 

After we have gotten the perf event file descriptor we attach the program by making a bpf link via the [link create syscall command](../syscall/BPF_LINK_CREATE.md).

We call the syscall command with the [`BPF_PERF_EVENT`](../syscall/BPF_LINK_CREATE.md#bpf_perf_event) [`attach_type`](../syscall/BPF_LINK_CREATE.md#attach_type), [`target_fd`](../syscall/BPF_LINK_CREATE.md#target_fd) set to the perf event file descriptor, [`prog_fd`](../syscall/BPF_LINK_CREATE.md#prog_fd) to the file descriptor of the tracepoint program, and optionally a [`cookie`](../syscall/BPF_LINK_CREATE.md#tracing-cookie)

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
??? abstract "Supported kfuncs"
    - [`__bpf_trap`](../kfuncs/__bpf_trap.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_arena_alloc_pages`](../kfuncs/bpf_arena_alloc_pages.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_arena_free_pages`](../kfuncs/bpf_arena_free_pages.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cast_to_kern_ctx`](../kfuncs/bpf_cast_to_kern_ctx.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cgroup_acquire`](../kfuncs/bpf_cgroup_acquire.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cgroup_ancestor`](../kfuncs/bpf_cgroup_ancestor.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_cgroup_from_id`](../kfuncs/bpf_cgroup_from_id.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
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
    - [`bpf_list_back`](../kfuncs/bpf_list_back.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_list_front`](../kfuncs/bpf_list_front.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_list_pop_back`](../kfuncs/bpf_list_pop_back.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_list_pop_front`](../kfuncs/bpf_list_pop_front.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_list_push_back_impl`](../kfuncs/bpf_list_push_back_impl.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_list_push_front_impl`](../kfuncs/bpf_list_push_front_impl.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
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
    - [`bpf_rbtree_add_impl`](../kfuncs/bpf_rbtree_add_impl.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_rbtree_first`](../kfuncs/bpf_rbtree_first.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_rbtree_left`](../kfuncs/bpf_rbtree_left.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_rbtree_remove`](../kfuncs/bpf_rbtree_remove.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_rbtree_right`](../kfuncs/bpf_rbtree_right.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_rbtree_root`](../kfuncs/bpf_rbtree_root.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_rcu_read_lock`](../kfuncs/bpf_rcu_read_lock.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_rcu_read_unlock`](../kfuncs/bpf_rcu_read_unlock.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_rdonly_cast`](../kfuncs/bpf_rdonly_cast.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_refcount_acquire_impl`](../kfuncs/bpf_refcount_acquire_impl.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_remove_dentry_xattr`](../kfuncs/bpf_remove_dentry_xattr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_res_spin_lock`](../kfuncs/bpf_res_spin_lock.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_res_spin_lock_irqsave`](../kfuncs/bpf_res_spin_lock_irqsave.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_res_spin_unlock`](../kfuncs/bpf_res_spin_unlock.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_res_spin_unlock_irqrestore`](../kfuncs/bpf_res_spin_unlock_irqrestore.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_send_signal_task`](../kfuncs/bpf_send_signal_task.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_set_dentry_xattr`](../kfuncs/bpf_set_dentry_xattr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_sock_destroy`](../kfuncs/bpf_sock_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_strchr`](../kfuncs/bpf_strchr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_strchrnul`](../kfuncs/bpf_strchrnul.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_strcmp`](../kfuncs/bpf_strcmp.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_strcspn`](../kfuncs/bpf_strcspn.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_strlen`](../kfuncs/bpf_strlen.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
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
    - [`bpf_throw`](../kfuncs/bpf_throw.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_verify_pkcs7_signature`](../kfuncs/bpf_verify_pkcs7_signature.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_wq_init`](../kfuncs/bpf_wq_init.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_wq_set_callback_impl`](../kfuncs/bpf_wq_set_callback_impl.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`bpf_wq_start`](../kfuncs/bpf_wq_start.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`crash_kexec`](../kfuncs/crash_kexec.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`css_rstat_flush`](../kfuncs/css_rstat_flush.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`css_rstat_updated`](../kfuncs/css_rstat_updated.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_cpu_node`](../kfuncs/scx_bpf_cpu_node.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_cpu_rq`](../kfuncs/scx_bpf_cpu_rq.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_cpuperf_cap`](../kfuncs/scx_bpf_cpuperf_cap.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_cpuperf_cur`](../kfuncs/scx_bpf_cpuperf_cur.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_cpuperf_set`](../kfuncs/scx_bpf_cpuperf_set.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_destroy_dsq`](../kfuncs/scx_bpf_destroy_dsq.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_dsq_nr_queued`](../kfuncs/scx_bpf_dsq_nr_queued.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
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
    - [`scx_bpf_now`](../kfuncs/scx_bpf_now.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_nr_cpu_ids`](../kfuncs/scx_bpf_nr_cpu_ids.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_nr_node_ids`](../kfuncs/scx_bpf_nr_node_ids.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_pick_any_cpu`](../kfuncs/scx_bpf_pick_any_cpu.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_pick_any_cpu_node`](../kfuncs/scx_bpf_pick_any_cpu_node.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_pick_idle_cpu`](../kfuncs/scx_bpf_pick_idle_cpu.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_pick_idle_cpu_node`](../kfuncs/scx_bpf_pick_idle_cpu_node.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_put_cpumask`](../kfuncs/scx_bpf_put_cpumask.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_put_idle_cpumask`](../kfuncs/scx_bpf_put_idle_cpumask.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_task_cgroup`](../kfuncs/scx_bpf_task_cgroup.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_task_cpu`](../kfuncs/scx_bpf_task_cpu.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_task_running`](../kfuncs/scx_bpf_task_running.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
    - [`scx_bpf_test_and_clear_cpu_idle`](../kfuncs/scx_bpf_test_and_clear_cpu_idle.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
<!-- [/PROG_KFUNC_REF] -->
