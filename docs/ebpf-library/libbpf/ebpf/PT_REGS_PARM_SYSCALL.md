---
title: "Libbpf eBPF macro 'PT_REGS_PARM_SYSCALL'"
description: "This page documents the 'PT_REGS_PARM_SYSCALL' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `PT_REGS_PARM_SYSCALL`

[:octicons-tag-24: v0.7.0](https://github.com/libbpf/libbpf/releases/tag/v0.7.0)

The `PT_REGS_PARM{1-8}_SYSCALL` macros make it easy to extract an argument from `struct pt_regs` following the syscall calling convention in an architecture-independent way.

## Usage

These macro are variants of the [`PT_REGS_PARAM{1-8}`](PT_REGS_PARM.md) macros that translate a parameter number to the correct register according to the syscall calling convention, which can be different from the normal calling convention. So when reading arguments to a syscall, these should be used.

The architecture for which the eBPF program is compiled is determined by setting one of the `__TARGET_ARCH_{arch}` macros. These are typically set by passing a flag to the compiler, such as `-D__TARGET_ARCH_x86` for x86. This allows for easy cross-compilation of eBPF programs for different architectures by changing the compiler invocation.

### Example

```c hl_lines="10 11"
// SPDX-License-Identifier: GPL-2.0

/*
 * Copyright 2020 Google LLC.
 */

SEC("fentry.s/" SYS_PREFIX "sys_setdomainname")
int [BPF_PROG](BPF_PROG.md)(test_sys_setdomainname, struct pt_regs *regs)
{
	void *ptr = (void *)PT_REGS_PARM1_SYSCALL(regs);
	int len = PT_REGS_PARM2_SYSCALL(regs);
	int buf = 0;
	long ret;

	ret = [bpf_copy_from_user](../../../linux/helper-function/bpf_copy_from_user.md)(&buf, sizeof(buf), ptr);
	if (len == -2 && ret == 0 && buf == 1234)
		copy_test++;
	if (len == -3 && ret == -EFAULT)
		copy_test++;
	if (len == -4 && ret == -EFAULT)
		copy_test++;
	return 0;
}
```
