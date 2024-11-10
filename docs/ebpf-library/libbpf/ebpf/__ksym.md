---
title: "Libbpf eBPF macro '__ksym'"
description: "This page documents the '__ksym' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `__ksym`

[:octicons-tag-24: v0.1.0](https://github.com/libbpf/libbpf/releases/tag/v0.1.0)

The `__ksym` macros is used to instruct the loader to provide the memory address of a kernel symbol.

## Definition

`#!c #define __ksym __attribute__((section(".ksyms")))`

## Usage

This macro tells the compiler to put the variable in the `.ksyms` section of the eBPF object file. This signals the loader (library) to initialize the value of the variable to the memory address of the kernel symbol with the same name as the variable.

When the defined variable of of the type `void *`, the loader will only consider the name of the kernel symbol. If it is a pointer to an explicit type, then the loader will use BTF to verify the type of the symbol.

An example use case for this feature would be to get the address of a global variable in the kernel and use [`bpf_probe_read_kernel`](../../../linux/helper-function/bpf_probe_read_kernel.md) to read its value.

The `__ksym` macro is also part of every [kfunc](../../../linux/concepts/kfuncs.md) definition, here it purpose is very similar. It instructs the loader to provide the memory address of the kernel function with the same type.

In both the variable and function case, if the symbol is not found, the loader (library) will error out, unless the [`__weak`](__weak.md) attribute is also used.

!!! note
    The Clang eBPF backend will actually omit the actual ELF section, but the data section still exists in the BTF type info and relocation entries.

### Example

```c hl_lines="8"
// SPDX-License-Identifier: GPL-2.0
/* Copyright (C) 2022. Huawei Technologies Co., Ltd */
#include "vmlinux.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>

extern bool CONFIG_PREEMPT __kconfig __weak;
extern const int bpf_task_storage_busy __ksym;

char _license[] SEC("license") = "GPL";

int pid = 0;
int busy = 0;

struct {
	__uint(type, BPF_MAP_TYPE_TASK_STORAGE);
	__uint(map_flags, BPF_F_NO_PREALLOC);
	__type(key, int);
	__type(value, long);
} task SEC(".maps");

SEC("raw_tp/sys_enter")
int BPF_PROG(read_bpf_task_storage_busy)
{
	int *value;

	if (!CONFIG_PREEMPT)
		return 0;

	if (bpf_get_current_pid_tgid() >> 32 != pid)
		return 0;

	value = bpf_this_cpu_ptr(&bpf_task_storage_busy);
	if (value)
		busy = *value;

	return 0;
}
```
