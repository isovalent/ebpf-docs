---
title: "KFunc 'bpf_send_signal_task'"
description: "This page documents the 'bpf_send_signal_task' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_send_signal_task`

<!-- [FEATURE_TAG](bpf_send_signal_task) -->
[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/6280cf718db0c557b5fe44e2d2e8ad8e832696a7)
<!-- [/FEATURE_TAG] -->

This function allows for the sending of signals to threads and processes.

## Definition

`bpf_send_signal_task` is a kfunc that is similar to the [`bpf_send_signal_thread`](../helper-function/bpf_send_signal_thread.md) and [`bpf_send_signal`](../helper-function/bpf_send_signal.md) helpers, but can be used to send signals to other threads and processes. It also supports sending a cookie with the signal similar to [`sigqueue()`](https://man7.org/linux/man-pages/man3/sigqueue.3.html).

If the receiving process establishes a handler for the signal using the `SA_SIGINFO` flag to [`sigaction()`](https://man7.org/linux/man-pages/man2/sigaction.2.html), then it can obtain this cookie via the `si_value` field of the `siginfo_t` structure passed as the second argument to the handler.

**Parameters**

`task`: Pointer to the task_struct of the thread or process to send the signal to.

`sig`: Signal number to send.

`type`: Specifies to send the signal to a specific process or a thread. Possible values: `PIDTYPE_PID`(0) and `PIDTYPE_TGID`(1)

`value`: Cookie to send with the signal.

**Returns**

`0` on success, a negative error code on failure

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_send_signal_task(struct task_struct *task, int sig, pid_type type, u64 value)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
- [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

```c
// SPDX-License-Identifier: GPL-2.0
// Copyright (c) 2019 Facebook
#include <vmlinux.h>
#include <linux/version.h>
#include <bpf/bpf_helpers.h>

struct task_struct *bpf_task_from_pid(int pid) __ksym;
void bpf_task_release(struct task_struct *p) __ksym;
int bpf_send_signal_task(struct task_struct *task, int sig, enum pid_type type, u64 value) __ksym;

__u32 sig = 0, pid = 0, status = 0, signal_thread = 0, target_pid = 0;

static __always_inline int bpf_send_signal_test(void *ctx)
{
	struct task_struct *target_task = NULL;
	int ret;
	u64 value;

	if (status != 0 || pid == 0)
		return 0;

	if ((bpf_get_current_pid_tgid() >> 32) == pid) {
		if (target_pid) {
			target_task = bpf_task_from_pid(target_pid);
			if (!target_task)
				return 0;
			value = 8;
		}

		if (signal_thread) {
			if (target_pid)
				ret = bpf_send_signal_task(target_task, sig, PIDTYPE_PID, value);
			else
				ret = bpf_send_signal_thread(sig);
		} else {
			if (target_pid)
				ret = bpf_send_signal_task(target_task, sig, PIDTYPE_TGID, value);
			else
				ret = bpf_send_signal(sig);
		}
		if (ret == 0)
			status = 1;
	}

	if (target_task)
		bpf_task_release(target_task);

	return 0;
}
```
