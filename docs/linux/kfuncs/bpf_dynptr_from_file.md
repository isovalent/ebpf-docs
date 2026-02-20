---
title: "KFunc 'bpf_dynptr_from_file'"
description: "This page documents the 'bpf_dynptr_from_file' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_dynptr_from_file`

<!-- [FEATURE_TAG](bpf_dynptr_from_file) -->
[:octicons-tag-24: v6.19](https://github.com/torvalds/linux/commit/8d8771dc03e48300e80b43744dd3c320ccaf746a)
<!-- [/FEATURE_TAG] -->

Create a [dynptr](../concepts/dynptrs.md) from a file.

## Definition

**Parameters**

`file`: The file to create a dynptr from, for later reading.

`flags`: Potential future flags, currently always `0`.

`ptr__uninit`: Pointer to an uninitialized dynptr, to be initialized by this call.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_dynptr_from_file(struct file *file, u32 flags, struct bpf_dynptr *ptr__uninit)`
<!-- [/KFUNC_DEF] -->

## Usage

This kfunc creates a read-only [dynptr](../concepts/dynptrs.md) for a given `file`. This allows a BPF program to directly do file reads, which can be useful in a number of use cases.

One use case is to implement ELF symbol parsing directly in eBPF, a process that is part of what a profiler does to transform observed events from memory addresses to human readable info. Something that traditionally had to be done in userspace and then communicated to eBPF via maps.

When a read is attempted on a section of a file that is not paged into memory a page fault occurs, which triggers the kernel to retrieve that bit of the file. When a read on a file dynptr is performed from a sleepable context, the program sleeps until the requested data is available. But in non-sleepable contexts, the read will result in a `-EFAULT` error.

One way to work around this, is to use the [`bpf_task_work_schedule_signal_impl`](bpf_task_work_schedule_signal_impl.md) kfunc to schedule a callback. This callback will run right before the scheduler returns execution to a task to invoke its signal handler. The callback is ran in a sleepable context. This is the approach taken in the [example](#example).

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_CGROUP_DEVICE`](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
- [`BPF_PROG_TYPE_CGROUP_SOCK`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
- [`BPF_PROG_TYPE_CGROUP_SYSCTL`](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_LWT_IN`](../program-type/BPF_PROG_TYPE_LWT_IN.md)
- [`BPF_PROG_TYPE_LWT_OUT`](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
- [`BPF_PROG_TYPE_LWT_SEG6LOCAL`](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
- [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
- [`BPF_PROG_TYPE_NETFILTER`](../program-type/BPF_PROG_TYPE_NETFILTER.md)
- [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
- [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md)
- [`BPF_PROG_TYPE_SOCKET_FILTER`](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
- [`BPF_PROG_TYPE_SOCK_OPS`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
- [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

Example of using a work queue to get a sleepable context to do file reads.

```c
// SPDX-License-Identifier: GPL-2.0
/* Copyright (c) 2025 Meta Platforms, Inc. and affiliates. */

#include <vmlinux.h>
#include <string.h>
#include <stdbool.h>
#include <bpf/bpf_tracing.h>
#include "bpf_misc.h"
#include "errno.h"

char _license[] [SEC](../../ebpf-library/libbpf/ebpf/SEC.md)("license") = "GPL";

struct {
	[__uint](../../ebpf-library/libbpf/ebpf/__uint.md)(type, BPF_MAP_TYPE_ARRAY);
	[__uint](../../ebpf-library/libbpf/ebpf/__uint.md)(max_entries, 1);
	[__type](../../ebpf-library/libbpf/ebpf/__type.md)(key, int);
	[__type](../../ebpf-library/libbpf/ebpf/__type.md)(value, struct elem);
} arrmap [SEC](../../ebpf-library/libbpf/ebpf/SEC.md)(".maps");

struct elem {
	struct file *file;
	struct bpf_task_work tw;
};

char user_buf[256000];
char tmp_buf[256000];

int pid = 0;
int err, run_success = 0;

static int validate_file_read(struct file *file);
static int task_work_callback(struct bpf_map *map, void *key, void *value);

SEC("lsm/file_open")
int on_open_validate_file_read(void *c)
{
	struct task_struct *task = [bpf_get_current_task_btf](../helper-function/bpf_get_current_task_btf.md)();
	struct elem *work;
	int key = 0;

	if ([bpf_get_current_pid_tgid](../helper-function/bpf_get_current_pid_tgid.md)() >> 32 != pid)
		return 0;

	work = [bpf_map_lookup_elem](../helper-function/bpf_map_lookup_elem.md)(&arrmap, &key);
	if (!work) {
		err = 1;
		return 0;
	}
	[bpf_task_work_schedule_signal_impl](bpf_task_work_schedule_signal_impl.md)(task, &work->tw, &arrmap, task_work_callback, NULL);
	return 0;
}

/* Called in a sleepable context, read 256K bytes, cross check with user space read data */
static int task_work_callback(struct bpf_map *map, void *key, void *value)
{
	struct task_struct *task = [bpf_get_current_task_btf](../helper-function/bpf_get_current_task_btf.md)();
	struct file *file = [bpf_get_task_exe_file](bpf_get_task_exe_file.md)(task);

	if (!file)
		return 0;

	err = validate_file_read(file);
	if (!err)
		run_success = 1;
	bpf_put_file(file);
	return 0;
}

static int verify_dynptr_read(struct bpf_dynptr *ptr, u32 off, char *user_buf, u32 len)
{
	int i;

	if ([bpf_dynptr_read](../helper-function/bpf_dynptr_read.md)(tmp_buf, len, ptr, off, 0))
		return 1;

	/* Verify file contents read from BPF is the same as the one read from userspace */
	[bpf_for](../../ebpf-library/libbpf/ebpf/bpf_for.md)(i, 0, len)
	{
		if (tmp_buf[i] != user_buf[i])
			return 1;
	}
	return 0;
}

static int validate_file_read(struct file *file)
{
	struct bpf_dynptr dynptr;
	int loc_err = 1, off;
	__u32 user_buf_sz = sizeof(user_buf);

	if (bpf_dynptr_from_file(file, 0, &dynptr))
		goto cleanup;

	loc_err = verify_dynptr_read(&dynptr, 0, user_buf, user_buf_sz);
	off = 1;
	loc_err = loc_err ?: verify_dynptr_read(&dynptr, off, user_buf + off, user_buf_sz - off);
	off = user_buf_sz - 1;
	loc_err = loc_err ?: verify_dynptr_read(&dynptr, off, user_buf + off, user_buf_sz - off);
	/* Read file with random offset and length */
	off = 4097;
	loc_err = loc_err ?: verify_dynptr_read(&dynptr, off, user_buf + off, 100);

	/* Adjust dynptr, verify read */
	loc_err = loc_err ?: bpf_dynptr_adjust(&dynptr, off, off + 1);
	loc_err = loc_err ?: verify_dynptr_read(&dynptr, 0, user_buf + off, 1);
	/* Can't read more than 1 byte */
	loc_err = loc_err ?: verify_dynptr_read(&dynptr, 0, user_buf + off, 2) == 0;
	/* Can't read with far offset */
	loc_err = loc_err ?: verify_dynptr_read(&dynptr, 1, user_buf + off, 1) == 0;

cleanup:
	[bpf_dynptr_file_discard](bpf_dynptr_file_discard.md)(&dynptr);
	return loc_err;
}
```
