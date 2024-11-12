---
title: "Program Type 'BPF_PROG_TYPE_KPROBE'"
description: "This page documents the 'BPF_PROG_TYPE_KPROBE' eBPF program type, including its definition, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_KPROBE`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_KPROBE) -->
[:octicons-tag-24: v4.1](https://github.com/torvalds/linux/commit/2541517c32be2531e0da59dfd7efc1ce844644f5)
<!-- [/FEATURE_TAG] -->

`BPF_PROG_TYPE_KPROBE` are eBPF programs that can attach to [kprobes](https://docs.kernel.org/trace/kprobes.html). KProbes are not a eBPF specific feature, but they do work very well together. Traditionally, one would have to write a custom kernel module which could be invoked from a kprobe or be content with just the trace log output. eBPF makes this process easier.

## Usage

Probes come in 4 different flavors: `kprobe`, `kretprobe`, `uprobe`, and `uretprobe`. `kprobe` and `kretprobe` are used to probe the kernel, `uprobe` and `uretprobe` are used to probe userspace. The normal probes are invoked when the probed location is executed. The `ret` variants will execute once the function returns, allowing for the capture of the return value.

<!-- TODO explain ELF section conventions -->

All of these probe types work with the kprobe program type, it is the attach method which determines how the program is executed.

The return value of kprobes programs doesn't do anything.

## Context

The context passed to kprobe programs is `struct pt_regs`. This structure is different for each CPU architecture since it contains a copy of the CPU registers at the time the kprobe was invoked.

It is common for kprobe programs to use the macros from the Libbpf `bpf_tracing.h` header file which defines `PT_REGS_PARM1` ... `PT_REGS_PARM5` as well as a number of others. These macros will translate to the correct field in `struct pt_regs` depending on the current architecture. Communicating the architecture you are compiling the BPF program for is done by defining one of the `__TARGET_ARCH_*` values in your program or via the command line while compiling.

The same header file also provides the `BPF_KPROBE(name, args...)` macro which allows program authors to define the function signatures in the same fashion as the functions they are tracing with type info and all. The macro will cast the correct argument numbers to the given argument names. For example:

```c
SEC("kprobe/proc_sys_write")
int BPF_KPROBE(my_kprobe_example,
		   struct file* filp, const char* buf,
		   size_t count, loff_t* ppos) {
    ...
}
```

Similar macros also exists for kprobes intended to attach to syscalls: `BPF_KSYSCALL(name, args...)` and kretprobes: `BPF_KRETPROBE(name, args...)`

## Attachment

There are two methods of attaching probe programs with variations for uprobes. The "legacy" way involves the manual creation of a `k{ret}probe` or `u{ret}probe` event via the [`DebugFS`](https://www.kernel.org/doc/html/next/filesystems/debugfs.html) and then attaching a BPF program to that event via the `perf_event_open` syscall.

The newer method uses BPF links to do both the probe event creation and attaching in one for multiple probes. Single probes can, however, still be attached via the `perf_event_open` syscall but require different parameters and need to utilize BPF links afterwards.

### Legacy kprobe attaching

First step is to create a kprobe or kretprobe trace event. To do so we can use the <nospell>DebugFS</nospell>, which we will assume is mounted at `/sys/kernel/debug` for the purposes of this document.

Existing kprobe events can be listed by printing `/sys/kernel/debug/tracing/kprobe_events`. And we can create new events by writing to this pseudo-file. For example executing `echo 'p:myprobe do_sys_open' > /sys/kernel/debug/tracing/kprobe_events`
will make a new kprobe (`p:`) called `myprobe` at the `do_sys_open` function in the kernel. For details on the full syntax, checkout [this link](https://docs.kernel.org/trace/kprobetrace.html). kretprobes are created by specifying a `r:` prefix.

After the probe has been created, a new directory will appear in `/sys/kernel/debug/tracing/events/kprobes/` with the same name as we have given our probe, `/sys/kernel/debug/tracing/events/kprobes/myprobe` in this case. This directory contains a few pseudo-files, for us `id` is important. The contents of `/sys/kernel/debug/tracing/events/kprobes/myprobe/id` contains a unique identifier we will need in the next step.

Next step is to open a new perf event using the [`perf_event_open`](https://man7.org/linux/man-pages/man2/perf_event_open.2.html) syscall:

```c
struct perf_event_attr attr = {
    .type = PERF_TYPE_TRACEPOINT,
    .size = sizeof(struct perf_event_attr),
    .config = kprobe_id, /* The ID of your kprobe */
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

This syscall will return a file descriptor on success. The final step are two [`ioctl`](https://man7.org/linux/man-pages/man2/ioctl.2.html) syscalls to attach our BPF program to the kprobe event and to enable the kprobe.

`#!c ioctl(perf_event_fd, PERF_EVENT_IOC_SET_BPF, bpf_prog_fd);` to attach.

`#!c ioctl(perf_event_fd, PERF_EVENT_IOC_ENABLE, 0);` to enable.

The kprobe can be temporality disabled with the `PERF_EVENT_IOC_DISABLE` ioctl option. Otherwise the kprobe stays attached until the perf_event goes away due to the closing of the perf_event FD or the program exiting. The perf event holds a reference to the BPF program so it will stay loaded until no more kprobes reference it.

<!-- TODO uprobe variation -->

### Link kprobe attaching

The more modern and preferred way of attaching is using the [link create command](../syscall/BPF_LINK_CREATE.md) of the BPF syscall. 
For single probes, open a new perf event using the [`perf_event_open`](https://man7.org/linux/man-pages/man2/perf_event_open.2.html) syscall. Note that the values of the attributes of the perf event structure are a little different here compared to the legacy way. 

<!-- TODO First use PMU instead of DebugFS -->

```c
 struct perf_event_attr attr = {
    .type = 8; /* read type from /sys/bus/event_source/devices/kprobe/type or uprobe/type */
    .sample_type = PERF_SAMPLE_RAW;
    .sample_period = 1;
    .wakeup_events = 1;
    .size = sizeof(attr);
    .config |= 1 << 0;
    .kprobe_func = ((uint64_t)"symbol_name");  /* symbol name in string, valid names can be found in /proc/kallsyms */
    .config1 = ((uint64_t)"symbol_name"); 
    .config2 = 0x0;
    .probe_offset = 0x0; /* offset must be a valid instruction, here it is just the start of the kernel symbol*/
};

syscall(SYS_perf_event_open, 
    &attr,  /* struct perf_event_attr * */
    -1,     /* pid_t pid */
    0       /* int cpu */
    -1,     /* int group_fd */
    PERF_FLAG_FD_CLOEXEC /* unsigned long flags */
);
```
<!-- TODO then use Link instead of perf_event -->

After the perf event syscall is successful, the valid file descriptor returned can be used to set the link_create.target_fd attribute in the bpf structure before the [link create command](../syscall/BPF_LINK_CREATE.md) is called.

```c
union bpf_attr attr = {
    .link_create.prog_fd = prog_fd; /* valid fd to bpf program of type KPROBE */
    .link_create.target_fd = perf_fd; /* valid fd to PMU event */
    .link_create.attach_type = BPF_PERF_EVENT;
    .link_create.flags = 0;
    .link_create.perf_event.bpf_cookie = 0;
};

syscall(SYS_bpf,
    BPF_LINK_CREATE,
    &attr,
    sizeof(attr)
);
```   
For multiple probes, [link create command](../syscall/BPF_LINK_CREATE.md) can be used to combine the creation and linking of the probes. [Fprobes](https://lore.kernel.org/bpf/20220316122419.933957-4-jolsa@kernel.org/) are used under the hood for multiple kprobes.

```c

union bpf_attr attr = {
    attr.link_create.prog_fd = prog_fd;
    attr.link_create.target_fd = 0;
    attr.link_create.attach_type = BPF_TRACE_KPROBE_MULTI;
    attr.link_create.flags = 0;
    attr.link_create.kprobe_multi.cnt = target_count; 
    attr.link_create.kprobe_multi.cookies = ((uint64_t)targets); 
    attr.link_create.kprobe_multi.flags = BPF_F_KPROBE_MULTI_RETURN;
    attr.link_create.kprobe_multi.syms = ptr_to_u64(targets);
);

syscall(SYS_bpf,
    BPF_LINK_CREATE,
    &attr,
    sizeof(attr)
);
```   

<!-- TODO upbrobe variation -->

## Helper functions

Not all helper functions are available in all program types. These are the helper calls available for socket filter programs:

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
    * [`bpf_override_return`](../helper-function/bpf_override_return.md)
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
    - [`bpf_session_cookie`](../kfuncs/bpf_session_cookie.md)
    - [`bpf_session_is_return`](../kfuncs/bpf_session_is_return.md)
<!-- [/PROG_KFUNC_REF] -->
