---
title: "KFunc 'bpf_copy_from_user_str'"
description: "This page documents the 'bpf_copy_from_user_str' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_copy_from_user_str`

<!-- [FEATURE_TAG](bpf_copy_from_user_str) -->
[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/65ab5ac4df012388481d0414fcac1d5ac1721fb3)
<!-- [/FEATURE_TAG] -->

This function copies a string from an unsafe user address.

## Definition

Copies a `NULL`-terminated string from userspace to BPF space. If user string is too long this will still ensure zero termination in the `dst` buffer unless buffer size is `0`.

**Parameters**

`dst`: Destination address, in kernel space. This buffer must be at least `dst__sz` bytes long.

`dst__sz`: Maximum number of bytes to copy, includes the trailing `NULL`.

`unsafe_ptr__ign`: Source address, in user space.

`flags`: The only supported flag is `BPF_F_PAD_ZEROS`

**Flags**

If `BPF_F_PAD_ZEROS` flag is set, `memset` the tail of `dst` to `0` on success and `memset` all of @dst on failure.

**Returns**

On success, the length of the `dst` string, including the trailing NULL character. On error, a negative value. If `dst__sz` is 0, then the return value will be 0.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_copy_from_user_str(void *dst, u32 dst__sz, const void *unsafe_ptr__ign, u64 flags)`

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

```c
// SPDX-License-Identifier: GPL-2.0
// Copyright (c) 2017 Facebook

#include "vmlinux.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_core_read.h>
#include <errno.h>
#include "bpf_misc.h"

u32 dynamic_sz = 1;
int uprobe_byname3_str_sleepable_res = 0;
void *user_ptr = 0;

int bpf_copy_from_user_str(void *dst, u32, const void *, u64) __weak __ksym;

static __always_inline bool verify_sleepable_user_copy_str(void)
{
	int ret;
	char data_long[20];
	char data_long_pad[20];
	char data_long_err[20];
	char data_short[4];
	char data_short_pad[4];

	ret = bpf_copy_from_user_str(data_short, sizeof(data_short), user_ptr, 0);

	if (bpf_strncmp(data_short, 4, "tes\0") != 0 || ret != 4)
		return false;

	ret = bpf_copy_from_user_str(data_short_pad, sizeof(data_short_pad), user_ptr, BPF_F_PAD_ZEROS);

	if (bpf_strncmp(data_short, 4, "tes\0") != 0 || ret != 4)
		return false;

	/* Make sure this passes the verifier */
	ret = bpf_copy_from_user_str(data_long, dynamic_sz & sizeof(data_long), user_ptr, 0);

	if (ret != 0)
		return false;

	ret = bpf_copy_from_user_str(data_long, sizeof(data_long), user_ptr, 0);

	if (bpf_strncmp(data_long, 10, "test_data\0") != 0 || ret != 10)
		return false;

	ret = bpf_copy_from_user_str(data_long_pad, sizeof(data_long_pad), user_ptr, BPF_F_PAD_ZEROS);

	if (bpf_strncmp(data_long_pad, 10, "test_data\0") != 0 || ret != 10 || data_long_pad[19] != '\0')
		return false;

	ret = bpf_copy_from_user_str(data_long_err, sizeof(data_long_err), (void *)data_long, BPF_F_PAD_ZEROS);

	if (ret > 0 || data_long_err[19] != '\0')
		return false;

	ret = bpf_copy_from_user_str(data_long, sizeof(data_long), user_ptr, 2);

	if (ret != -EINVAL)
		return false;

	return true;
}

SEC("uprobe.s//proc/self/exe:trigger_func3")
int handle_uprobe_byname3_sleepable(struct pt_regs *ctx)
{
	if (verify_sleepable_user_copy_str())
		uprobe_byname3_str_sleepable_res = 10;
	return 0;
}
```
