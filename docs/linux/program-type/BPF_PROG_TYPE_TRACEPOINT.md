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

We can use the [tracefs](https://www.kernel.org/doc/Documentation/trace/ftrace.txt) to list all of these available trace events. For the sake of this page we will assume the tracefs is mounted at `/sys/kernel/tracing` (which is usual for most distros). The `/sys/kernel/tracing/events/` directory contains a number of yet more directories. The events are grouped by the first word in their name, so all `kvm_*` events reside in `/sys/kernel/tracing/events/kvm`. So `xdp_exception` is located in `/sys/kernel/tracing/events/xdp/xdp_exception`. We will refer to this directory as the "event directory".

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

We start by looking up the event ID in the tracefs. Inside the event directory is located a pseudo-file called `id`, so for `xdp_exception` that would be `/sys/kernel/tracing/events/xdp/xdp_exception/id`. When reading the file a decimal number is returned.

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

We call the syscall command with the [`BPF_PERF_EVENT`](../syscall/BPF_LINK_CREATE.md#bpf_perf_event) [`attach_type`](../syscall/BPF_LINK_CREATE.md#attach_type), [`target_fd`](../syscall/BPF_LINK_CREATE.md#target_fd) set to the perf event fd, [`prog_fd`](../syscall/BPF_LINK_CREATE.md#prog_fd) to the file descriptor of the tracepoint program, and optionally a [`cookie`](../syscall/BPF_LINK_CREATE.md#cookie)

## Helper functions

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [bpf_perf_event_output](../helper-function/bpf_perf_event_output.md)
    * [bpf_get_stackid](../helper-function/bpf_get_stackid.md)
    * [bpf_get_stack](../helper-function/bpf_get_stack.md)
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
