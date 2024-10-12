---
title: "Helper Function 'bpf_probe_write_user'"
description: "This page documents the 'bpf_probe_write_user' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_probe_write_user`

<!-- [FEATURE_TAG](bpf_probe_write_user) -->
[:octicons-tag-24: v4.8](https://github.com/torvalds/linux/commit/96ae52279594470622ff0585621a13e96b700600)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Attempt in a safe way to write _len_ bytes from the buffer _src_ to _dst_ in memory. It only works for threads that are in user context, and _dst_ must be a valid user space address.

This helper should not be used to implement any kind of security mechanism because of TOC-TOU attacks, but rather to debug, divert, and manipulate execution of semi-cooperative processes.

Keep in mind that this feature is meant for experiments, and it has a risk of crashing the system and running programs. Therefore, when an eBPF program using this helper is attached, a warning including PID and process name is printed to kernel logs.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_probe_write_user)(void *dst, const void *src, __u32 len) = (void *) 36;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_KPROBE`](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
 * [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

```c
#include <vmlinux.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_core_read.h>

// We do it in the exit to not alter the syscall behavior. The userspace program
// will see the new filename only after the syscall execution.
SEC("fexit/__x64_sys_open")
int BPF_PROG(p_open, struct pt_regs *regs, long ret) {
  // If it is our example program overwrite the open path.
  struct task_struct *task = (struct task_struct *)bpf_get_current_task_btf();
  if (bpf_strncmp(task->comm, TASK_COMM_LEN, "example") != 0) {
    return 0;
  }

  // SYSCALL_DEFINE3(open, const char __user *, filename, int, flags, umode_t, mode)
  // first param is the pointer to filename.
  void *filename_ptr = (void *)PT_REGS_PARM1_CORE_SYSCALL(regs);
  const char filename[16] = "/tmp/new";
  if (bpf_probe_write_user(filename_ptr, filename, 16)) {
    bpf_printk("Failed to write new filename\n");
  }
  return 0;
}
```
