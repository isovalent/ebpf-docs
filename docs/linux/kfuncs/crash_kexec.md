---
title: "KFunc 'crash_kexec'"
description: "This page documents the 'crash_kexec' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `crash_kexec`

<!-- [FEATURE_TAG](crash_kexec) -->
[:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/133790596406ce2658f0864eb7eac64987c2b12f)
<!-- [/FEATURE_TAG] -->

Crash the kernel at a specific point in the code.

## Definition

<!-- [KFUNC_DEF] -->
`#!c void crash_kexec(struct pt_regs *regs)`

!!! warning
	This kfunc is destructive to the system. For example such a call can result in system rebooting or panicking. 
	Due to this additional restrictions apply to these calls. At the moment they only require CAP_SYS_BOOT capability, 
	but more can be added later.
<!-- [/KFUNC_DEF] -->

## Usage

eBPF is often used for kernel debugging, and one of the widely used and
powerful debugging techniques is post-mortem debugging with a full memory dump.

This kfunc allows to trigger a kernel panic at a specific point in the kernels 
execution, this allows for the inspection of the memory dump at the exact point
a program detected a certain condition.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
- [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [BPF_PROG_TYPE_STRUCT_OPS](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
- [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

