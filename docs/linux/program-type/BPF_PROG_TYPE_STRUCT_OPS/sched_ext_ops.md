---
title: "Struct ops 'sched_ext_ops'"
description: "This page documents the 'sched_ext_ops' struct ops, its semantics, capabilities, and limitations."
---
# Struct ops `sched_ext_ops`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

Sched ext (Scheduler extension) Ops can be used to implement a custom scheduler in BPF.

## Usage

The Linux kernel provides built-in scheduler implementations like [`CFS`](https://docs.kernel.org/scheduler/sched-design-CFS.html) and [`EEVDF`](https://docs.kernel.org/scheduler/sched-eevdf.html). These schedulers are designed to provide a good balance between fairness and performance for most workloads. However, there are use cases where a custom scheduler is needed to meet specific requirements. The BPF scheduler extension provides a way to implement a custom scheduler in BPF.

See also [kernel docs](https://www.kernel.org/doc/html/next/scheduler/sched-ext.html)

## Fields and ops

A BPF scheduler can implement an arbitrary scheduling policy by implementing and loading operations in this table. Note that a userland scheduling policy can also be implemented using the BPF scheduler as a shim layer.


!!! note
    The following definition has been modified from the one found in the kernel for the sake of readability. This does not impact the definition for the purposes of implementing a BPF program.

```c
struct sched_ext_ops {
    char [name](#name)[SCX_OPS_NAME_LEN];
    u32  [dispatch_max_batch](#dispatch_max_batch);
    u64  [flags](#flags);
    u32  [timeout_ms](#timeout_ms);
    u32  [exit_dump_len](#exit_dump_len);
    u64  [hotplug_seq](#hotplug_seq);

    s32  (*[select_cpu](#select_cpu))([struct task_struct](#struct-task_struct) *p, s32 prev_cpu, u64 wake_flags);
    void (*[enqueue](#enqueue))([struct task_struct](#struct-task_struct) *p, u64 enq_flags);
    void (*[dequeue](#dequeue))([struct task_struct](#struct-task_struct) *p, u64 deq_flags);
    void (*[dispatch](#dispatch))(s32 cpu, [struct task_struct](#struct-task_struct) *prev);
    void (*[tick](#tick))([struct task_struct](#struct-task_struct) *p);
    void (*[runnable](#runnable))([struct task_struct](#struct-task_struct) *p, u64 enq_flags);
    void (*[running](#running))([struct task_struct](#struct-task_struct) *p);
    void (*[stopping](#stopping))([struct task_struct](#struct-task_struct) *p, bool runnable);
    void (*[quiescent](#quiescent))([struct task_struct](#struct-task_struct) *p, u64 deq_flags);
    bool (*[yield](#yield))([struct task_struct](#struct-task_struct) *from, [struct task_struct](#struct-task_struct) *to);
    bool (*[core_sched_before](#core_sched_before))([struct task_struct](#struct-task_struct) *a, [struct task_struct](#struct-task_struct) *b);
    void (*[set_weight](#set_weight))([struct task_struct](#struct-task_struct) *p, u32 weight);
    void (*[set_cpumask](#set_cpumask))([struct task_struct](#struct-task_struct) *p, const [struct cpumask](#struct-cpumask) *cpumask);
    void (*[update_idle](#update_idle))(s32 cpu, bool idle);
    void (*[cpu_acquire](#cpu_acquire))(s32 cpu, [struct scx_cpu_acquire_args](#struct-scx_cpu_acquire_args) *args);
    void (*[cpu_release](#cpu_release))(s32 cpu, [struct scx_cpu_release_args](#struct-scx_cpu_release_args) *args);
    
    s32  (*[init_task](#init_task))([struct task_struct](#struct-task_struct) *p, [struct scx_init_task_args](#struct-scx_init_task_args) *args);
    void (*[exit_task](#exit_task))([struct task_struct](#struct-task_struct) *p, [struct scx_exit_task_args](#struct-scx_exit_task_args) *args);
    
    void (*[enable](#enable))([struct task_struct](#struct-task_struct) *p);
    void (*[disable](#disable))([struct task_struct](#struct-task_struct) *p);
    
    void (*[dump](#dump))([struct scx_dump_ctx](#struct-scx_dump_ctx) *ctx);
    void (*[dump_cpu](#dump_cpu))([struct scx_dump_ctx](#struct-scx_dump_ctx) *ctx, s32 cpu, bool idle);
    void (*[dump_task](#dump_task))([struct scx_dump_ctx](#struct-scx_dump_ctx) *ctx, [struct task_struct](#struct-task_struct) *p);

#ifdef CONFIG_EXT_GROUP_SCHED
    s32  (*[cgroup_init](#cgroup_init))([struct cgroup](#struct-cgroup) *cgrp, [struct scx_cgroup_init_args](#struct-scx_cgroup_init_args) *args);
    void (*[cgroup_exit](#cgroup_exit))([struct cgroup](#struct-cgroup) *cgrp);
    s32  (*[cgroup_prep_move](#cgroup_prep_move))([struct task_struct](#struct-task_struct) *p, [struct cgroup](#struct-cgroup) *from, [struct cgroup](#struct-cgroup) *to);
    void (*[cgroup_move](#cgroup_move))([struct task_struct](#struct-task_struct) *p, [struct cgroup](#struct-cgroup) *from, [struct cgroup](#struct-cgroup) *to);
    void (*[cgroup_cancel_move](#cgroup_cancel_move))([struct task_struct](#struct-task_struct) *p, [struct cgroup](#struct-cgroup) *from, [struct cgroup](#struct-cgroup) *to);
    void (*[cgroup_set_weight](#cgroup_set_weight))([struct cgroup](#struct-cgroup) *cgrp, u32 weight);
#endif /* CONFIG_EXT_GROUP_SCHED */

    void (*[cpu_online](#cpu_online))(s32 cpu);
    void (*[cpu_offline](#cpu_offline))(s32 cpu);

    s32  (*[init](#init))(void);
    void (*[exit](#exit))([struct scx_exit_info](#struct-scx_exit_info) *info);
};
```

### `name`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c name[SCX_OPS_NAME_LEN]` (`SCX_OPS_NAME_LEN = 128`)

The BPF <nospell>scheduler's</nospell> name, for observability purposes.

Must be a non-zero valid BPF object name including only [`isalnum()`](https://elixir.bootlin.com/linux/v6.13/source/tools/include/linux/ctype.h#L25), `_` and `.` chars. Shows up in `kernel.sched_ext_ops` sysctl while the BPF scheduler is enabled.

### `dispatch_max_batch`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c u32 dispatch_max_batch`

Max number of tasks that [`dispatch`](#dispatch) can dispatch.

### `flags`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c u64 flags`

The flags field is a bitfield that can be set to control the behavior of the scheduler. The [`enum scx_ops_flags`](#enum-scx_ops_flags) enum defines the flags that can be set in this field.

### `timeout_ms`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c u32 timeout_ms`

The maximum amount of time, in milliseconds, that a runnable task should be able to wait before being scheduled. The maximum timeout may not exceed the default timeout of 30 seconds.

Defaults to the maximum allowed timeout value of 30 seconds.

### `exit_dump_len`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c u32 exit_dump_len`

[`scx_exit_info.dump`](#struct-scx_exit_info-dump) buffer length. If `0`, the default value of `32768` is used.

### `hotplug_seq`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c u64 hotplug_seq`

A sequence number that may be set by the scheduler to detect when a hot-plug event has occurred during the loading process. If `0`, no detection occurs. Otherwise, the scheduler will fail to load if the sequence number does not match [`scx_hotplug_seq`](https://elixir.bootlin.com/linux/v6.13/source/kernel/sched/ext.c#L899) on the enable path.

### `select_cpu`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c s32 (*select_cpu)(struct task_struct *p, s32 prev_cpu, u64 wake_flags);`

Pick the target CPU for a task which is being woken up.

Decision made here isn't final. `p` may be moved to any CPU while it is getting dispatched for execution later. However, as `p` is not on the rq at this point, getting the eventual execution CPU right here saves a small bit of overhead down the line.

If an idle CPU is returned, the CPU is kicked and will try to dispatch. While an explicit custom mechanism can be added, `select_cpu` serves as the default way to wake up idle CPUs.

`p` may be inserted into a DSQ directly by calling [`scx_bpf_dsq_insert`](../../kfuncs/scx_bpf_dsq_insert.md). If so, the [`enqueue`](#enqueue) will be skipped. Directly inserting into [`SCX_DSQ_LOCAL`](#scx_dsq_local) will put `p` in the local DSQ of the CPU returned by this operation.

!!! note
    `select_cpu` is never called for tasks that can only run on a single CPU or tasks with migration disabled, as they don't have the option to select a different CPU. See [`select_task_rq`](https://elixir.bootlin.com/linux/v6.13/source/kernel/sched/core.c#L3566) for details.

**Parameters**

`p`: task being woken upa

`prev_cpu`: the cpu `p` was on before sleeping

`wake_flags`: `SCX_WAKE_*`, possible values are:

* `SCX_WAKE_FORK` (`0x02`) - Wakeup after exec
* `SCX_WAKE_TTWU` (`0x04`) - Wakeup after fork
* `SCX_WAKE_SYNC` (`0x08`) - Wakeup

**Returns**

The ID of the CPU to be woken up.

### `enqueue`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*enqueue)(struct task_struct *p, u64 enq_flags);`

Enqueue a task on the BPF scheduler


`p` is ready to run. Insert directly into a DSQ by calling [`scx_bpf_dsq_insert`](../../kfuncs/scx_bpf_dsq_insert.md) or enqueue on the BPF scheduler. If not directly inserted, the bpf scheduler owns `p` and if it fails to dispatch `p`, the task will stall.

If `p` was inserted into a DSQ from [`select_cpu`](#select_cpu), this callback is skipped.

**Parameters**

`p`: task being enqueued

`enq_flags`: Enqueue flags, possible values defined by [`enum scx_enq_flags`](#enum-scx_enq_flags)

### `dequeue`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*dequeue)(struct task_struct *p, u64 deq_flags);`

Remove a task from the BPF scheduler

Remove `p` from the BPF scheduler. This is usually called to isolate the task while updating its scheduling properties (e.g. priority).

The ext core keeps track of whether the BPF side owns a given task or not and can gracefully ignore spurious dispatches from BPF side, which makes it safe to not implement this method. However, depending on the scheduling logic, this can lead to confusing behaviors - e.g. scheduling position not being updated across a priority change.

**Parameters**

`p`: task being dequeued

`deq_flags`: Dequeue flags, possible values defined by [`enum scx_deq_flags`](#enum-scx_deq_flags)

### `dispatch`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*dispatch)(s32 cpu, struct task_struct *prev);`

Dispatch tasks from the BPF scheduler and/or user DSQs

Called when a CPU's local DSQ is empty. The operation should dispatch one or more tasks from the BPF scheduler into the DSQs using [`scx_bpf_dsq_insert`](../../kfuncs/scx_bpf_dsq_insert.md) and/or move from user DSQs into the local DSQ using [`scx_bpf_dsq_move_to_local`](../../kfuncs/scx_bpf_dsq_move_to_local.md).

The maximum number of times [`scx_bpf_dsq_insert`](../../kfuncs/scx_bpf_dsq_insert.md) can be called without an intervening [`scx_bpf_dsq_move_to_local`](../../kfuncs/scx_bpf_dsq_move_to_local.md) is specified by ops.dispatch_max_batch. See the comments on top of the two functions for more details.

When not `NULL`, `prev` is an SCX task with its slice depleted. If `prev` is still runnable as indicated by set [`SCX_TASK_QUEUED`](#scx_task_queued) in [`prev->scx.flags`](#struct-sched_ext_entity-flags), it is not enqueued yet and will be enqueued after [`dispatch`](#dispatch) returns. To keep executing `prev`, return without dispatching or moving any tasks. Also see [`SCX_OPS_ENQ_LAST`](#scx_ops_enq_last).

**Parameters**

`cpu`: CPU to dispatch tasks for

`prev`: previous task being switched out


### `tick`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*tick)(struct task_struct *p);`

Periodic tick. This operation is called every 1/HZ seconds on CPUs which are executing an SCX task. Setting [`p->scx.slice`](#struct-sched_ext_entity-slice) to `0` will trigger an immediate dispatch cycle on the CPU.

**Parameters**

`p`: task running currently

### `runnable`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*runnable)(struct task_struct *p, u64 enq_flags);`

A task is becoming runnable on its associated CPU

This and the following three functions can be used to track a task's execution state transitions. A task becomes [`runnable`](#runnable) on a CPU, and then goes through one or more [`running`](#running) and [`stopping`](#stopping) pairs as it runs on the CPU, and eventually becomes [`quiescent`](#quiescent) when it's done running on the CPU.

`p` is becoming runnable on the CPU because it's

* waking up ([`SCX_ENQ_WAKEUP`](#scx_enq_wakeup))
* being moved from another CPU
* being restored after temporarily taken off the queue for an attribute change.

This and [`enqueue`](#enqueue) are related but not coupled. This operation notifies `p`'s state transition and may not be followed by [`enqueue`](#enqueue) e.g. when `p` is being dispatched to a remote CPU, or when `p` is being enqueued on a CPU experiencing a hotplug event. Likewise, a task may be [`enqueue`](#enqueue)'d without being preceded by this operation e.g. after exhausting its slice.

**Parameters**

`p`: task becoming runnable

`enq_flags`: Bitfield of flags, valid values defined in [`enum scx_enq_flags`](#enum-scx_enq_flags)

### `running`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*running)(struct task_struct *p);`

A task is starting to run on its associated CPU. See [`runnable`](#runnable) for explanation on the task state notifiers.

**Parameters**

`p`: task starting to run

### `stopping`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*stopping)(struct task_struct *p, bool runnable);`

A task is stopping execution. See [`runnable`](#runnable) for explanation on the task state notifiers. If !`runnable`, [`quiescent`](#quiescent) will be invoked after this operation returns.

**Parameters**

`p`: task stopping to run

`runnable`: is task `p` still runnable?


### `quiescent`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*quiescent)(struct task_struct *p, u64 deq_flags);`

A task is becoming not runnable on its associated CPU. See [`runnable`](#runnable) for explanation on the task state notifiers.

`p` is becoming quiescent on the CPU because it's

* sleeping ([`SCX_DEQ_SLEEP`](#scx_deq_sleep))
* being moved to another CPU
* being temporarily taken off the queue for an attribute change (`SCX_DEQ_SAVE`).

* This and [`dequeue`](#dequeue) are related but not coupled. This operation
* notifies `p`'s state transition and may not be preceded by [`dequeue`](#dequeue)
* e.g. when `p` is being dispatched to a remote CPU.

**Parameters**

`p`: task becoming not runnable

`deq_flags`: Bitfield of flags, valid values defined in [`enum scx_deq_flags`](#enum-scx_deq_flags)

### `yield`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c bool (*yield)(struct task_struct *from, struct task_struct *to);`

Yield CPU. If `to` is `NULL`, `from` is yielding the CPU to other runnable tasks. The BPF scheduler should ensure that other available tasks are dispatched before the yielding task. Return value is ignored in this case. 

If `to` is not-`NULL`, `from` wants to yield the CPU to `to`.

**Parameters**

`from`: yielding task

`to`: optional yield target task

**Returns**

If the bpf scheduler can implement the request, return `true`; otherwise, `false`.


### `core_sched_before`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c bool (*core_sched_before)(struct task_struct *a, struct task_struct *b);`

Task ordering for core-sched. Used by core-sched to determine the ordering between two tasks. See [`Documentation/admin-guide/hw-vuln/core-scheduling.rst`](https://www.kernel.org/doc/Documentation/admin-guide/hw-vuln/core-scheduling.rst) for details on core-sched.

Both `a` and `b` are runnable and may or may not currently be queued on the BPF scheduler. 

If not specified, the default is ordering them according to when they became runnable.

**Parameters**

`a`: task A

`b`: task B

**Returns**

Should return `true` if `a` should run before `b`. `false` if there's no required ordering or `b` should run before `a`.

### `set_weight`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*set_weight)(struct task_struct *p, u32 weight);`

Set task weight. Update `p`'s weight to `weight`.

**Parameters**

`p`: task to set weight for

`weight`: new weight `[1..10000]`

### `set_cpumask`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*set_cpumask)(struct task_struct *p, const struct cpumask *cpumask);`

Set CPU affinity. Update `p`'s CPU affinity to `cpumask`.

**Parameters**

`p`: task to set CPU affinity for

`cpumask`: [cpumask](#struct-cpumask) of cpus that `p` can run on

### `update_idle`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*update_idle)(s32 cpu, bool idle);`

Update the idle state of a CPU. This operation is called when `rq`'s CPU goes or leaves the idle state. By default, implementing this operation disables the built-in idle CPU tracking and the following helpers become unavailable:

* [`scx_bpf_select_cpu_dfl`](../../kfuncs/scx_bpf_select_cpu_dfl.md)
* [`scx_bpf_test_and_clear_cpu_idle`](../../kfuncs/scx_bpf_test_and_clear_cpu_idle.md)
* [`scx_bpf_pick_idle_cpu`](../../kfuncs/scx_bpf_pick_idle_cpu.md)

The user also must implement [`select_cpu`](#select_cpu) as the default implementation relies on [`scx_bpf_select_cpu_dfl`](../../kfuncs/scx_bpf_select_cpu_dfl.md).

Specify the [`SCX_OPS_KEEP_BUILTIN_IDLE`](#scx_ops_keep_builtin_idle) flag to keep the built-in idle tracking.

**Parameters**

`cpu`: CPU to update the idle state for

`idle`: whether entering or exiting the idle state

### `cpu_acquire`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*cpu_acquire)(s32 cpu, struct scx_cpu_acquire_args *args);`

A CPU is becoming available to the BPF scheduler. A CPU that was previously released from the BPF scheduler is now once again under its control.

**Parameters**

`cpu`: The CPU being acquired by the BPF scheduler.

`args`: Acquire arguments, see [`struct scx_cpu_acquire_args`](#struct-scx_cpu_acquire_args).

### `cpu_release`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*cpu_release)(s32 cpu, struct scx_cpu_release_args *args);`

A CPU is taken away from the BPF scheduler. The specified CPU is no longer under the control of the BPF scheduler. This could be because it was preempted by a higher priority sched_class, though there may be other reasons as well. The caller should consult [`args->reason`](#struct-scx_cpu_release_args-reason) to determine the cause.

**Parameters**

`cpu`: The CPU being released by the BPF scheduler.

`args`: Release arguments, see [`struct scx_cpu_release_args`](#struct-scx_cpu_release_args).

### `init_task`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c s32 (*init_task)(struct task_struct *p, struct scx_init_task_args *args);`

Initialize a task to run in a BPF scheduler.  Either we're loading a BPF scheduler or a new task is being forked. Initialize `p` for BPF scheduling. This operation may block and can be used for allocations, and is called exactly once for a task.

!!! note
    The BPF program assigned to this field is allowed to be [sleepable](../../syscall/BPF_PROG_LOAD.md#bpf_f_sleepable).

**Parameters**

`p`: task to initialize for BPF scheduling

`args`: init arguments, see [`struct scx_init_task_args`](#struct-scx_init_task_args)

**Returns**

`0` for success, `-errno` for failure. An error return while loading will abort loading of the BPF scheduler. During a fork, it will abort that specific fork.

### `exit_task`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*exit_task)(struct task_struct *p, struct scx_exit_task_args *args);`

Exit a previously-running task from the system. `p` is exiting or the BPF scheduler is being unloaded. Perform any necessary cleanup for `p`.

**Parameters**

`p`: task to exit

`args`: exit arguments, see [`struct scx_exit_task_args`](#struct-scx_exit_task_args)

### `enable`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*enable)(struct task_struct *p);`

Enable BPF scheduling for a task. Enable `p` for BPF scheduling. `enable` is called on `p` any time it enters SCX, and is always paired with a matching [`disable`](#disable).

**Parameters**

`p`: task to enable BPF scheduling for

### `disable`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*disable)(struct task_struct *p);`

Disable BPF scheduling for a task. `p` is exiting, leaving SCX or the BPF scheduler is being unloaded. Disable BPF scheduling for `p`. A `disable` call is always matched with a prior [`enable`](#enable) call.

**Parameters**

`p`: task to disable BPF scheduling for

### `dump`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*dump)(struct scx_dump_ctx *ctx);`

Dump BPF scheduler state on error. Use [`scx_bpf_dump`](../../../ebpf-library/scx/scx_bpf_dump.md) to generate BPF scheduler specific debug dump.

**Parameters**

`ctx`: debug dump context, see [`struct scx_dump_ctx`](#struct-scx_dump_ctx)

### `dump_cpu`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*dump_cpu)(struct scx_dump_ctx *ctx, s32 cpu, bool idle);`

Dump BPF scheduler state for a CPU on error. Use [`scx_bpf_dump`](../../../ebpf-library/scx/scx_bpf_dump.md) to generate BPF scheduler specific debug dump for `cpu`. If `idle` is `true` and this operation doesn't produce any output, `cpu` is skipped for dump.

**Parameters**

`ctx`: debug dump context, see [`struct scx_dump_ctx`](#struct-scx_dump_ctx)

`cpu`: CPU to generate debug dump for

`idle`: `cpu` is currently idle without any runnable tasks

### `dump_task`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*dump_task)(struct scx_dump_ctx *ctx, struct task_struct *p);`

Dump BPF scheduler state for a runnable task on error

Use [`scx_bpf_dump`](../../../ebpf-library/scx/scx_bpf_dump.md) to generate BPF scheduler specific debug dump for `p`.

**Parameters**

`ctx`: debug dump context, see [`struct scx_dump_ctx`](#struct-scx_dump_ctx)

`p`: runnable task to generate debug dump for

### `cgroup_init`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/8195136669661fdfe54e9a8923c33b31c92fc1da)

`#!c s32 (*cgroup_init)(struct cgroup *cgrp, struct scx_cgroup_init_args *args);`

Initialize a cGroup. Either the BPF scheduler is being loaded or `cgrp` created, initialize `cgrp` for sched_ext. This operation may block.

!!! note
    This field is only available on kernels compiled with the `CONFIG_EXT_GROUP_SCHED` Kconfig enabled.

!!! note
    The BPF program assigned to this field is allowed to be [sleepable](../../syscall/BPF_PROG_LOAD.md#bpf_f_sleepable).

**Parameters**

`cgrp`: cgroup being initialized

`args`: init arguments, see [`struct scx_cgroup_init_args`](#struct-scx_cgroup_init_args)

**Returns**

`0` for success, `-errno` for failure. An error return while loading will abort loading of the BPF scheduler. During cgroup creation, it will abort the specific cgroup creation.

### `cgroup_exit`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/8195136669661fdfe54e9a8923c33b31c92fc1da)

`#!c void (*cgroup_exit)(struct cgroup *cgrp);`

Exit a cGroup. Either the BPF scheduler is being unloaded or `cgrp` destroyed, exit `cgrp` for sched_ext. This operation my block.

!!! note
    This field is only available on kernels compiled with the `CONFIG_EXT_GROUP_SCHED` Kconfig enabled.

!!! note
    The BPF program assigned to this field is allowed to be [sleepable](../../syscall/BPF_PROG_LOAD.md#bpf_f_sleepable).

**Parameters**

`cgrp`: cgroup being exited

### `cgroup_prep_move`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/8195136669661fdfe54e9a8923c33b31c92fc1da)

`#!c s32 (*cgroup_prep_move)(struct task_struct *p, struct cgroup *from, struct cgroup *to);`

Prepare a task to be moved to a different cGroup. Prepare `p` for move from cGroup `from` to `to`. This operation may block and can be used for allocations.

!!! note
    This field is only available on kernels compiled with the `CONFIG_EXT_GROUP_SCHED` Kconfig enabled.

!!! note
    The BPF program assigned to this field is allowed to be [sleepable](../../syscall/BPF_PROG_LOAD.md#bpf_f_sleepable).

**Parameters**

`p`: task being moved

`from`: cgroup `p` is being moved from

`to`: cgroup `p` is being moved to

**Returns**

`0` for success, `-errno` for failure. An error return aborts the migration.

### `cgroup_move`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/8195136669661fdfe54e9a8923c33b31c92fc1da)

`#!c void (*cgroup_move)(struct task_struct *p, struct cgroup *from, struct cgroup *to);`

Commit cGroup move. `p` is dequeued during this operation.

!!! note
    This field is only available on kernels compiled with the `CONFIG_EXT_GROUP_SCHED` Kconfig enabled.

**Parameters**

`p`: task being moved

`from`: cgroup `p` is being moved from

`to`: cgroup `p` is being moved to

### `cgroup_cancel_move`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/8195136669661fdfe54e9a8923c33b31c92fc1da)

`#!c void (*cgroup_cancel_move)(struct task_struct *p, struct cgroup *from, struct cgroup *to);`

Cancel cGroup move. `p` was [`cgroup_prep_move`](#cgroup_prep_move)'d but failed before reaching [`cgroup_move`](#cgroup_move). Undo the preparation.

!!! note
    This field is only available on kernels compiled with the `CONFIG_EXT_GROUP_SCHED` Kconfig enabled.

**Parameters**

`p`: task whose cgroup move is being canceled

`from`: cgroup `p` was being moved from

`to`: cgroup `p` was being moved to

### `cgroup_set_weight`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/8195136669661fdfe54e9a8923c33b31c92fc1da)

`#!c void (*cgroup_set_weight)(struct cgroup *cgrp, u32 weight);`

A cGroup's weight is being changed. Update `cgrp`'s weight to `weight`.

!!! note
    This field is only available on kernels compiled with the `CONFIG_EXT_GROUP_SCHED` Kconfig enabled.

**Parameters**

`cgrp`: cgroup whose weight is being updated

`weight`: new weight `[1..10000]`

### `cpu_online`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*cpu_online)(s32 cpu);`

A CPU became online. `cpu` just came online. `cpu` will not call [`enqueue`](#enqueue) or [`dispatch`](#dispatch), nor run tasks associated with other CPUs beforehand.

!!! note
    The BPF program assigned to this field is allowed to be [sleepable](../../syscall/BPF_PROG_LOAD.md#bpf_f_sleepable).

**Parameters**

`cpu`: CPU which just came up

### `cpu_offline`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*cpu_offline)(s32 cpu);`

A CPU is going offline. `cpu` is going offline. `cpu` will not call [`enqueue`](#enqueue) or [`dispatch`](#dispatch), nor run tasks associated with other CPUs afterwards.

!!! note
    The BPF program assigned to this field is allowed to be [sleepable](../../syscall/BPF_PROG_LOAD.md#bpf_f_sleepable).

**Parameters**

`cpu`: CPU which is going offline

### `init`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c s32 (*init)(void);`

Initialize the BPF scheduler.

!!! note
    The BPF program assigned to this field is allowed to be [sleepable](../../syscall/BPF_PROG_LOAD.md#bpf_f_sleepable).

### `exit`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c void (*exit)(struct scx_exit_info *info);`

Clean up after the BPF scheduler. [`exit`](#exit) is also called on [`init`](#init) failure, which is a bit unusual. This is to allow rich reporting through `info` on how [`init`](#init) failed.

!!! note
    The BPF program assigned to this field is allowed to be [sleepable](../../syscall/BPF_PROG_LOAD.md#bpf_f_sleepable).

**Parameters**

`info`: Exit info, see [`struct scx_exit_info`](#struct-scx_exit_info)

## Types

### `enum scx_ops_flags`

This enum defines all of the flags that can be set as bitfield in [`flags`](#flags).

```c
enum scx_ops_flags {
    [SCX_OPS_KEEP_BUILTIN_IDLE](#scx_ops_keep_builtin_idle)   = 1LLU << 0,
    [SCX_OPS_ENQ_LAST](#scx_ops_enq_last)            = 1LLU << 1,
    [SCX_OPS_ENQ_EXITING](#scx_ops_enq_exiting)         = 1LLU << 2,
    [SCX_OPS_SWITCH_PARTIAL](#scx_ops_switch_partial)      = 1LLU << 3,
    [SCX_OPS_HAS_CGROUP_WEIGHT](#scx_ops_has_cgroup_weight)   = 1LLU << 16,
};
```

#### `SCX_OPS_KEEP_BUILTIN_IDLE`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

Keep built-in idle tracking even if [`update_idle`](#update_idle) is implemented.

#### `SCX_OPS_ENQ_LAST`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

By default, if there are no other task to run on the CPU, ext core keeps running the current task even after its slice expires. If this flag is specified, such tasks are passed to ops.[`enqueue`](#enqueue) with [`SCX_ENQ_LAST`](#scx_enq_last). See the comment above [`SCX_ENQ_LAST`](#scx_enq_last) for more info.

#### `SCX_OPS_ENQ_EXITING`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

An exiting task may schedule after [`PF_EXITING`](https://elixir.bootlin.com/linux/v6.13.4/source/include/linux/sched.h#L1675) is set. In such cases, [`bpf_task_from_pid`](../../kfuncs/bpf_task_from_pid.md) may not be able to find the task and if the BPF scheduler depends on PID lookup for dispatching, the task will be lost leading to various issues including RCU grace period stalls.

To mask this problem, by default, <nospell>unhashed</nospell> tasks are automatically dispatched to the local DSQ on enqueue. If the BPF scheduler doesn't depend on PID lookups and wants to handle these tasks directly, the following flag can be used.

#### `SCX_OPS_SWITCH_PARTIAL`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

If set, only tasks with policy set to [`SCHED_EXT`](https://elixir.bootlin.com/linux/v6.13.4/source/include/uapi/linux/sched.h#L121) are attached to sched_ext. If clear, SCHED_NORMAL tasks are also included.

#### `SCX_OPS_HAS_CGROUP_WEIGHT`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/8195136669661fdfe54e9a8923c33b31c92fc1da)

CPU cGroup support flags.

### `enum scx_enq_flags`

```c
enum scx_enq_flags {
    [SCX_ENQ_WAKEUP](#scx_enq_wakeup)          = 1LLU << 0,
    [SCX_ENQ_HEAD](#scx_enq_head)            = 1LLU << 4,
    [SCX_ENQ_CPU_SELECTED](#scx_enq_cpu_selected)    = 1LLU << 10,
    [SCX_ENQ_PREEMPT](#scx_enq_preempt)         = 1LLU << 32,
    [SCX_ENQ_REENQ](#scx_enq_reenq)           = 1LLU << 40,
    [SCX_ENQ_LAST](#scx_enq_last)            = 1LLU << 41,
    [SCX_ENQ_CLEAR_OPSS](#scx_enq_clear_opss)      = 1LLU << 56,
    [SCX_ENQ_DSQ_PRIQ](#scx_enq_dsq_priq)        = 1LLU << 57,
};
```

#### `SCX_ENQ_WAKEUP`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

Mark a task as runnable.

#### `SCX_ENQ_HEAD`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

Place the task at the head of the runqueue. If not set, the task is placed at the tail.

#### `SCX_ENQ_CPU_SELECTED`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/9b671793c7d95f020791415cbbcc82b9c007d19c)

This flag is set by the scheduler core internals in [`select_task_rq`](https://elixir.bootlin.com/linux/v6.13/source/kernel/sched/core.c#L3566), if the scheduler called the `select_task_rq` callback of the current class. This callback for SCX translates into [`select_cpu`](#select_cpu).

`select_task_rq`/`select_cpu` is not called when a task can only run on 1 CPU or CPU migration is disabled, since in those cases there is no decision to be made. 

#### `SCX_ENQ_PREEMPT`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

Set the following to trigger preemption when calling [`scx_bpf_dsq_insert`](../../kfuncs/scx_bpf_dsq_insert.md) with a local DSQ as the target. The slice of the current task is cleared to zero and the CPU is kicked into the scheduling path. Implies [`SCX_ENQ_HEAD`](#scx_enq_head).

#### `SCX_ENQ_REENQ`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

The task being enqueued was previously enqueued on the current CPU's [`SCX_DSQ_LOCAL`](#scx_dsq_local), but was removed from it in a call to the [`bpf_scx_reenqueue_local`](../../kfuncs/scx_bpf_reenqueue_local.md) kfunc. If [`bpf_scx_reenqueue_local`](../../kfuncs/scx_bpf_reenqueue_local.md) was invoked in a [`cpu_release`](#cpu_release) callback, and the task is again dispatched back to [`SCX_DSQ_LOCAL`](#scx_dsq_local) by this current [`enqueue`](#enqueue), the task will not be scheduled on the CPU until at least the next invocation of the [`cpu_acquire`](#cpu_acquire) callback.

#### `SCX_ENQ_LAST`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

The task being enqueued is the only task available for the CPU. By default, ext core keeps executing such tasks but when [`SCX_OPS_ENQ_LAST`](#scx_ops_enq_last) is specified, they're [`enqueue`](#enqueue)'d with the [`SCX_ENQ_LAST`](#scx_enq_last) flag set.

The BPF scheduler is responsible for triggering a follow-up scheduling event. Otherwise, Execution may stall.

#### `SCX_ENQ_CLEAR_OPSS`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

When this flag is set, to hand back control over a task from the BPF scheduler to the SCX core. Setting 
[`p->scx.ops_state`](#struct-sched_ext_entity-ops_state) to [`SCX_OPSS_NONE`](#enum-scx_ops_state)

#### `SCX_ENQ_DSQ_PRIQ`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/06e51be3d5e7a07aea5c9012773df8d5de01db6c)

This flag is set when a task is inserted into a DSQ with the [`scx_bpf_dsq_insert_vtime`](../../kfuncs/scx_bpf_dsq_insert_vtime.md) kfunc. It indicates that the ordering of the task in the DSQ is based on the virtual time of the task, not insertion order.

### `enum scx_deq_flags`

```c
enum scx_deq_flags {
    SCX_DEQ_SLEEP           = 1LLU << 0,
    SCX_DEQ_CORE_SCHED_EXEC = 1LLU << 32,
};
```

#### `SCX_DEQ_SLEEP`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

Task is no longer runnable.

#### `SCX_DEQ_CORE_SCHED_EXEC`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

The generic core-sched layer decided to execute the task even though it hasn't been dispatched yet. Dequeue from the BPF side.

### `enum scx_dsq_id_flags`

DSQ (dispatch queue) IDs are 64bit of the format:

```
Bits: [63] [62 ..  0]
      [ B] [   ID   ]
```

* `B`: 1 for IDs for built-in DSQs, 0 for ops-created user DSQs
* `ID`: 63 bit ID

Built-in IDs:

```
Bits: [63] [62] [61..32] [31 ..  0]
      [ 1] [ L] [   R  ] [    V   ]
```

* `1`: 1 for built-in DSQs.
* `L`: 1 for LOCAL_ON DSQ IDs, 0 for others
* `V`: For LOCAL_ON DSQ IDs, a CPU number. For others, a pre-defined value.

```c
enum scx_dsq_id_flags {
    SCX_DSQ_FLAG_BUILTIN    = 1LLU << 63,
    SCX_DSQ_FLAG_LOCAL_ON   = 1LLU << 62,

    SCX_DSQ_INVALID         = SCX_DSQ_FLAG_BUILTIN | 0,
    SCX_DSQ_GLOBAL          = SCX_DSQ_FLAG_BUILTIN | 1,
    SCX_DSQ_LOCAL           = SCX_DSQ_FLAG_BUILTIN | 2,
    SCX_DSQ_LOCAL_ON        = SCX_DSQ_FLAG_BUILTIN | SCX_DSQ_FLAG_LOCAL_ON,
};
```

#### `SCX_DSQ_FLAG_BUILTIN`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

This flag is part of a DSQ ID. If set indicates that the DSQ is a built-in DSQ or a DSQ created by the BPF scheduler.

#### `SCX_DSQ_FLAG_LOCAL_ON`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

This flag is part of a DSQ ID. If set indicates the DSQ is a local DSQ and the CPU number is encoded in the ID.

#### `SCX_DSQ_GLOBAL`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

Combined flags. The DSQ is builtin and global.

#### `SCX_DSQ_LOCAL`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

Combined flags. The DSQ is builtin and local to the current CPU.

#### `SCX_DSQ_LOCAL_ON`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

Combined flags. The DSQ is builtin and local to a specific CPU (encoded in the ID).

### `enum scx_ent_flags`

```c
enum scx_ent_flags {
    SCX_TASK_QUEUED             = 1 << 0,
	SCX_TASK_RESET_RUNNABLE_AT  = 1 << 2,
	SCX_TASK_DEQD_FOR_SLEEP     = 1 << 3,
	SCX_TASK_CURSOR             = 1 << 31,
};
```

#### `SCX_TASK_QUEUED`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

On ext runqueue

#### `SCX_TASK_RESET_RUNNABLE_AT`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

[`runnable_at`](#struct-sched_ext_entity-runnable_at) should be reset

#### `SCX_TASK_DEQD_FOR_SLEEP`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

Last dequeue was for SLEEP.

#### `SCX_TASK_CURSOR`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

iteration cursor, not a task

### `enum scx_task_state`

```c
enum scx_task_state {
	SCX_TASK_NONE,
	SCX_TASK_INIT,
	SCX_TASK_READY,
	SCX_TASK_ENABLED,
};
```

#### `SCX_TASK_NONE`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

[`init_task`](#init_task) not called yet

#### `SCX_TASK_INIT`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

[`init_task`](#init_task) succeeded, but task can be cancelled

#### `SCX_TASK_READY`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

fully initialized, but not in sched_ext

#### `SCX_TASK_ENABLED`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

fully initialized and in sched_ext

### `enum scx_kf_mask`

Mask bits for [`sched_ext_entity.kf_mask`](#struct-sched_ext_entity-kf_mask). Not all kfuncs can be called from everywhere and the following bits track which kfunc sets are currently allowed for [`current`](https://elixir.bootlin.com/linux/v6.13/source/include/asm-generic/current.h#L9). This simple per-task tracking works because SCX ops nest in a limited way. BPF will likely implement a way to allow and disallow kfuncs depending on the calling context which will replace this manual mechanism. See [`scx_kf_allow()`](https://elixir.bootlin.com/linux/v6.13/source/kernel/sched/ext.c#L1081).

```c
enum scx_kf_mask {
    SCX_KF_UNLOCKED     = 0,
    SCX_KF_CPU_RELEASE  = 1 << 0,
    SCX_KF_DISPATCH     = 1 << 1,
    SCX_KF_ENQUEUE      = 1 << 2,
    SCX_KF_SELECT_CPU   = 1 << 3,
    SCX_KF_REST         = 1 << 4,
};
```
#### `SCX_KF_UNLOCKED`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

sleepable and not <nospell>rq</nospell> locked

#### `SCX_KF_CPU_RELEASE`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/245254f7081dbe1c8da54675d0e4ddbe74cee61b)

This flag is set when kfuncs are enabled that may only be called from the [`cpu_release`](#cpu_release) callback. `SCX_KF_ENQUEUE` and `SCX_KF_DISPATCH` may be nested inside `SCX_KF_CPU_RELEASE`.

#### `SCX_KF_DISPATCH`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

This flag is set when kfuncs are enabled that may only be called from the [`dispatch`](#dispatch) callback. `SCX_KF_REST` may be nested inside `SCX_KF_DISPATCH`.

#### `SCX_KF_ENQUEUE`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

This flag is set when kfuncs are enabled that may only be called from the [`enqueue`](#enqueue) callback. `SCX_KF_SELECT_CPU` may be nested inside `SCX_KF_ENQUEUE`.

#### `SCX_KF_SELECT_CPU`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

This flag is set when kfuncs are enabled that may only be called from the [`select_cpu`](#select_cpu) callback.

#### `SCX_KF_REST`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

This flag is set when the rest of the kfuncs (kfuncs not under any of the other flags) are enabled.

### `enum scx_ops_state`

```c
enum scx_ops_state {
	SCX_OPSS_NONE,          // (1)!
	SCX_OPSS_QUEUEING,      // (2)!
	SCX_OPSS_QUEUED,        // (3)!
	SCX_OPSS_DISPATCHING,   // (4)!
};
```

1. [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f) owned by the SCX core
2. [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f) in transit to the BPF scheduler
3. [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f) owned by the BPF scheduler
4. [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f) in transit back to the SCX core

### `enum scx_ent_dsq_flags`

```c
enum scx_ent_dsq_flags {
	SCX_TASK_DSQ_ON_PRIQ = 1 << 0, // (1)!
};
```

1. [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/06e51be3d5e7a07aea5c9012773df8d5de01db6c) task is queued on the priority queue of a DSQ

### `enum scx_exit_kind`

```c
enum scx_exit_kind {
    SCX_EXIT_NONE,
    SCX_EXIT_DONE,          // (1)!

    SCX_EXIT_UNREG = 64,    // (2)!
    SCX_EXIT_UNREG_BPF,     // (3)!
    SCX_EXIT_UNREG_KERN,    // (4)!
    SCX_EXIT_SYSRQ,         // (5)!

    SCX_EXIT_ERROR = 1024,  // (6)!
    SCX_EXIT_ERROR_BPF,     // (7)!
    SCX_EXIT_ERROR_STALL,   // (8)!
};
```

1. [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)
2. [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f) user-space initiated unregistration 
3. [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f) BPF-initiated unregistration 
4. [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f) kernel-initiated unregistration 
5. [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f) requested by 'S' <nospell>sysrq</nospell>
6. [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f) runtime error, error message contains details 
7. [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f) ERROR but triggered through scx_bpf_error() 
8. [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f) watchdog detected stalled runnable tasks 

### `enum scx_cpu_preempt_reason`

```c
enum scx_cpu_preempt_reason {
	SCX_CPU_PREEMPT_RT,      // (1)!
	SCX_CPU_PREEMPT_DL,      // (2)!
	SCX_CPU_PREEMPT_STOP,    // (3)!
	SCX_CPU_PREEMPT_UNKNOWN, // (4)!
};
```

1. next task is being scheduled by [`&rt_sched_class`](https://elixir.bootlin.com/linux/v6.13.4/source/kernel/sched/sched.h#L2513)
2. next task is being scheduled by [`&dl_sched_class`](https://elixir.bootlin.com/linux/v6.13.4/source/kernel/sched/sched.h#L2512)
3. next task is being scheduled by [`&stop_sched_class`](https://elixir.bootlin.com/linux/v6.13.4/source/kernel/sched/sched.h#L2511)
4. unknown reason for SCX being preempted

### `struct task_struct`

This struct is the main data structure for a task in the Linux kernel. It contains all the information about a task, including its state, scheduling information, and more. Due to its size, only the fields relevant to SCX are documented here. The below definition is a simplified version of the actual struct. For the full definition see the [Linux source code](https://elixir.bootlin.com/linux/v6.13.4/source/include/linux/sched.h#L785).

```c
struct task_struct {
#ifdef CONFIG_SCHED_CLASS_EXT
	[struct sched_ext_entity](#struct-sched_ext_entity) scx;
#endif
};
```

#### `scx` {#struct-task_struct-scx}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c struct sched_ext_entity scx`

SCX specific information for the task.

### `struct cgroup`

This the representation of a cGroup in the Linux kernel. It is part of a tree of cGroups, in the hierarchy defined by the user. BPF schedulers can walk pointers provided in this struct to access cGroup related information.

As this struct is not directly related to SCX, its not documented here. For the full definition see the [Linux source code](https://elixir.bootlin.com/linux/v6.13.4/source/include/linux/cgroup-defs.h#L412).

### `struct sched_ext_entity`

This struct is embedded in [`task_struct`](#struct-task_struct) and contains all fields necessary for a task to be scheduled by SCX.

The fields on this structure are read-only unless otherwise noted.

```c
struct sched_ext_entity {
    struct scx_dispatch_q      *dsq;
    struct scx_dsq_list_node    dsq_list; 
    struct rb_node              dsq_priq; 

    u32 dsq_seq;
    u32 dsq_flags;
    u32 flags; 
    u32 weight;
    s32 sticky_cpu;
    s32 holding_cpu;
    u32 kf_mask; 
    
    [struct task_struct](#struct-task_struct)  *kf_tasks[2];
    atomic_long_t        ops_state;
    struct list_head     runnable_node;
    unsigned long        runnable_at;

#ifdef CONFIG_SCHED_CORE
    u64 core_sched_at;
#endif
    
    u64  ddsp_dsq_id;
    u64  ddsp_enq_flags;
    u64  slice;
    u64  dsq_vtime;
    bool disallow;

#ifdef CONFIG_EXT_GROUP_SCHED
    struct cgroup       *cgrp_moving_from;
#endif
    struct list_head    tasks_node;
};
```

#### `dsq` {#struct-sched_ext_entity-dsq}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c struct scx_dispatch_q *dsq`

The DSQ the task is currently on, or `NULL` if the task is not on any DSQ.

#### `dsq_list` {#struct-sched_ext_entity-dsq_list}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/d4af01c3731ff9c6e224d7183f8226a56d72b56c)

`#!c struct scx_dsq_list_node dsq_list`

The linked list node, that is part of the FIFO-DSQ the task is on. The linked list is in dispatch order.

#### `dsq_priq` {#struct-sched_ext_entity-dsq_priq}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/d4af01c3731ff9c6e224d7183f8226a56d72b56c)

`#!c struct rb_node dsq_priq`

The red-black tree node, that is part of the vtime-DSQ the task is on. The red-black priority queue is ordered by
[`p->scx.dsq_vtime`](#struct-sched_ext_entity-dsq_vtime).

#### `dsq_seq` {#struct-sched_ext_entity-dsq_seq}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/650ba21b131ed1f8ee57826b2c6295a3be221132)

`#!c u32 dsq_seq`

This is the DSQ sequence number the task is on.

#### `dsq_flags` {#struct-sched_ext_entity-dsq_flags}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/d4af01c3731ff9c6e224d7183f8226a56d72b56c)

`#!c u32 dsq_flags`

Flags related to the DSQ the task is on. See [`enum scx_ent_dsq_flags`](#enum-scx_ent_dsq_flags) for possible values.

#### `flags` {#struct-sched_ext_entity-flags}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c u32 flags`

This field contains both flags defined in [`enum scx_ent_flags`](#enum-scx_ent_flags) and a task state defined in [`enum scx_task_state`](#enum-scx_task_state). The value of the task state can be masked out with [`scx_entity.flags & SCX_TASK_STATE_MASK`](https://elixir.bootlin.com/linux/v6.13/source/include/linux/sched/ext.h#L79).

#### `weight` {#struct-sched_ext_entity-weight}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c u32 weight`

The weight of the task. A value in the range `1..10000`. The higher the weight, the more priority the task should have.

#### `sticky_cpu` {#struct-sched_ext_entity-sticky_cpu}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c s32 sticky_cpu`

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

#### `holding_cpu` {#struct-sched_ext_entity-holding_cpu}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c s32 holding_cpu`

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

#### `kf_mask` {#struct-sched_ext_entity-kf_mask}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c u32 kf_mask`

See [`scx_kf_mask`](#enum-scx_kf_mask).

#### `kf_tasks` {#struct-sched_ext_entity-kf_tasks}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/36454023f50b2aadd25b7f47c50b5edc0c59e1c0)

`#!c struct task_struct *kf_tasks[2]`

[`SCX_CALL_OP_TASK()`](https://elixir.bootlin.com/linux/v6.13/source/kernel/sched/ext.c#L1132)

#### `ops_state` {#struct-sched_ext_entity-ops_state}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c atomic_long_t ops_state`

Used to track the task ownership between the SCX core and the BPF scheduler. Valid values described by [`enum scx_ops_state`]
(#enum-scx_ops_state).

State transitions look as follows:

```
 NONE -> QUEUEING -> QUEUED -> DISPATCHING
   ^              |                 |
   |              v                 v
   \-------------------------------/
```

`QUEUEING` and `DISPATCHING` states can be waited upon. See [wait_ops_state()](https://elixir.bootlin.com/linux/v6.13/source/kernel/sched/ext.c#L1495) call sites for explanations on the conditions being waited upon and why they are safe. Transitions out of them into `NONE` or `QUEUED` must store_release and the waiters should load_acquire.

Tracking scx_ops_state enables sched_ext core to reliably determine whether any given task can be dispatched by the BPF scheduler at all times and thus relaxes the requirements on the BPF scheduler. This allows the BPF scheduler to try to dispatch any task anytime regardless of its state as the SCX core can safely reject invalid dispatches.

#### `runnable_node` {#struct-sched_ext_entity-runnable_node}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c struct list_head runnable_node`

This is a node of a runqueue linked list. This node is linked into the  [`rq->scx.runnable_list`](https://elixir.bootlin.com/linux/v6.13/source/kernel/sched/sched.h#L770) when the task becomes runnable.

#### `runnable_at` {#struct-sched_ext_entity-runnable_at}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/8a010b81b3a50b033fc3cddc613517abda586cbe)

`#!c unsigned long runnable_at`

The [jiffies](https://man7.org/linux/man-pages/man7/time.7.html) value at which the task became runnable. 

#### `core_sched_at` {#struct-sched_ext_entity-core_sched_at}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/7b0888b7cc1924b74ce660e02f6211df8dd46e7b)

`#!c u64 core_sched_at`

See [`scx_prio_less()`](https://elixir.bootlin.com/linux/v6.13/source/kernel/sched/ext.c#L3135)

!!! note
    This field is only available on kernels compiled with the `CONFIG_SCHED_CORE` Kconfig enabled.

#### `ddsp_dsq_id` {#struct-sched_ext_entity-ddsp_dsq_id}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c u64 ddsp_dsq_id`

The DSQ ID when on the direct dispatch path.

#### `ddsp_enq_flags` {#struct-sched_ext_entity-ddsp_enq_flags}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c u64 ddsp_enq_flags`

The DSQ enqueue flags when on the direct dispatch path.

#### `slice` {#struct-sched_ext_entity-slice}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c u64 slice`

!!! note
    This field can be modified by the BPF scheduler.

Runtime budget in nanoseconds. This is usually set through [`scx_bpf_dispatch`](../../kfuncs/scx_bpf_dispatch.md) but can also be modified directly by the BPF scheduler. Automatically decreased by SCX as the task executes. On depletion, a scheduling event is triggered.

This value is cleared to zero if the task is preempted by [`SCX_KICK_PREEMPT`](../../kfuncs/scx_bpf_kick_cpu.md#scx_kick_preempt) and shouldn't be used to determine how long the task ran. Use [`p->se.sum_exec_runtime`](https://elixir.bootlin.com/linux/v6.13/source/include/linux/sched.h#L557) instead.

#### `dsq_vtime` {#struct-sched_ext_entity-dsq_vtime}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/06e51be3d5e7a07aea5c9012773df8d5de01db6c)

`#!c u64 dsq_vtime`

!!! note
    This field can be modified by the BPF scheduler.

Used to order tasks when dispatching to the vtime-ordered priority queue of a DSQ. This is usually set through [`scx_bpf_dispatch_vtime`](../../kfuncs/scx_bpf_dispatch_vtime.md) but can also be modified directly by the BPF scheduler. Modifying it while a task is queued on a DSQ may mangle the ordering and is not recommended.

#### `disallow` {#struct-sched_ext_entity-disallow}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/7bb6f0810ecfb73a9d7a2ca56fb001e0201a6758)

`#!c bool disallow`

!!! note
    This field can be modified by the BPF scheduler.

Reject switching into SCX.

If set, reject future [`sched_setscheduler(2)`](https://man7.org/linux/man-pages/man2/sched_setscheduler.2.html) calls updating the policy to [`SCHED_EXT`](https://elixir.bootlin.com/linux/v6.13.4/source/include/uapi/linux/sched.h#L121) with `-EACCES`.

Can be set from [`init_task`](#init_task) while the BPF scheduler is being loaded (![`scx_init_task_args->fork`](#struct-scx_init_task_args-fork)). If set and the task's policy is already [`SCHED_EXT`](https://elixir.bootlin.com/linux/v6.13.4/source/include/uapi/linux/sched.h#L121), the task's policy is rejected and forcefully reverted to [`SCHED_NORMAL`](https://elixir.bootlin.com/linux/v6.13.4/source/include/uapi/linux/sched.h#L114). The number of such events are reported through `/sys/kernel/debug/sched_ext::nr_rejected`. Setting this flag during fork is not allowed.

#### `cgrp_moving_from` {#struct-sched_ext_entity-cgrp_moving_from}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/8195136669661fdfe54e9a8923c33b31c92fc1da)

`#!c struct cgroup *cgrp_moving_from`

!!! note
    This field is only available on kernels compiled with the `CONFIG_EXT_GROUP_SCHED` Kconfig enabled.

#### `tasks_node` {#struct-sched_ext_entity-tasks_node}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c struct list_head tasks_node`

This is a node of a task list linked list. This node is linked into the [`scx_tasks`](https://elixir.bootlin.com/linux/v6.13.4/source/kernel/sched/ext.c#L864) after [forking](https://man7.org/linux/man-pages/man2/fork.2.html).

### `struct cpumask`

This structure is a bitmap, one bit for every CPU.

```c
struct cpumask { 
    [DECLARE_BITMAP](https://elixir.bootlin.com/linux/v6.13.4/source/include/linux/types.h#L10)(bits, NR_CPUS); 
};
```

### `struct scx_cpu_acquire_args`

Argument container for [`cpu_acquire`](#cpu_acquire). Currently empty, but may be expanded in the future.

```c
struct scx_cpu_acquire_args {};
```

### `struct scx_cpu_release_args`

argument container for [`cpu_release`](#cpu_release)

```c
struct scx_cpu_release_args {
	[enum scx_cpu_preempt_reason](#enum-scx_cpu_preempt_reason) reason;
	[struct task_struct](#struct-task_struct)         *task;
};
```

#### `reason` {#struct-scx_cpu_release_args-reason}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/245254f7081dbe1c8da54675d0e4ddbe74cee61b)

The reason the CPU was preempted. See [`enum scx_cpu_preempt_reason`](#enum-scx_cpu_preempt_reason) for possible values.

#### `task` {#struct-scx_cpu_release_args-task}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/245254f7081dbe1c8da54675d0e4ddbe74cee61b)

The task that's going to be scheduled on the CPU.

### `struct scx_init_task_args`

Argument container for [`init_task`](#init_task)

```c
struct scx_init_task_args {
	bool fork;
#ifdef CONFIG_EXT_GROUP_SCHED
	struct cgroup *cgroup;
#endif
};
```

#### `fork` {#struct-scx_init_task_args-fork}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

Set if [`init_task`](#init_task) is being invoked on the fork path, as opposed to the scheduler transition path.

#### `cgroup` {#struct-scx_init_task_args-cgroup}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/8195136669661fdfe54e9a8923c33b31c92fc1da)

The cGroup the task is joining.

### `struct scx_exit_task_args`

Argument container for [`exit_task`](#exit_task)

```c
struct scx_exit_task_args {
	bool cancelled;
};
```

#### `cancelled` {#struct-scx_exit_task_args-cancelled}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

Whether the task exited before running on sched_ext.

### `struct scx_dump_ctx`

Informational context provided to dump operations.

```c
struct scx_dump_ctx {
	[enum scx_exit_kind](#enum-scx_exit_kind)  kind;
	s64                 exit_code;
	const char         *reason;
	u64                 at_ns;
	u64                 at_jiffies;
};
```

### `struct scx_cgroup_init_args`

Argument container for [`cgroup_init`](#cgroup_init)

```c
struct scx_cgroup_init_args {
	u32 weight;
};
```

#### `weight` {#struct-scx_cgroup_init_args-weight}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/8195136669661fdfe54e9a8923c33b31c92fc1da)

The weight of the cGroup [1..10000].

### `struct scx_exit_info`

This struct is passed to [`exit`](#exit) to describe why the BPF scheduler is being disabled.

```c
struct scx_exit_info {
    enum scx_exit_kind  kind;
    s64                 exit_code;
    const char          *reason;
    unsigned long       *bt;
    u32                 bt_len;
    char                *msg;
    char                *dump;
};
```

#### `kind` {#struct-scx_exit_info-kind}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c enum scx_exit_kind kind`

Broad category of the exit reason, one of [`enum scx_exit_kind`](#enum-scx_exit_kind)

#### `exit_code` {#struct-scx_exit_info-exit_code}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c s64 exit_code`

Exit code if gracefully exiting

#### `reason` {#struct-scx_exit_info-reason}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c const char *reason;`

Textual representation of the above

#### `bt` {#struct-scx_exit_info-bt}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c unsigned long *bt;`

Backtrace if exiting due to an error

#### `bt_len` {#struct-scx_exit_info-bt_len}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c u32 bt_len;`

#### `msg` {#struct-scx_exit_info-msg}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`#!c char *msg;`

Informational message

#### `dump` {#struct-scx_exit_info-dump}

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/f0e1a0643a59bf1f922fa209cec86a170b784f3f)

`char *dump;`

debug dump
