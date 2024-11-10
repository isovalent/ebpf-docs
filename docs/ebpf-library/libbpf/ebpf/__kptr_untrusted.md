---
title: "Libbpf eBPF macro '__kptr_untrusted'"
description: "This page documents the '__kptr_untrusted' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `__kptr_untrusted`

[:octicons-tag-24: v1.2.0](https://github.com/libbpf/libbpf/releases/tag/v1.2.0)

The `__kptr_untrusted` macros is used to tag a pointer to tell the verifier it holds untrusted pointers to kernel memory.

## Definition

`#!c #define __kptr_untrusted __attribute__((btf_type_tag("kptr_untrusted")))`

## Usage

This macro can used on type definitions for both global variables and fields in map values. It informs the verifier that the pointer is a untrusted kernel pointer. 

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome 

### Example

```c hl_lines="7"
// SPDX-License-Identifier: GPL-2.0
#include <vmlinux.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_helpers.h>

struct map_value {
	struct task_struct __kptr_untrusted *ptr;
};

struct {
	__uint(type, BPF_MAP_TYPE_LRU_HASH);
	__uint(max_entries, 1);
	__type(key, int);
	__type(value, struct map_value);
} lru_map SEC(".maps");

int pid = 0;
int result = 1;

SEC("fentry/bpf_ktime_get_ns")
int printk(void *ctx)
{
	struct map_value v = {};

	if (pid == bpf_get_current_task_btf()->pid)
		bpf_map_update_elem(&lru_map, &(int){0}, &v, 0);
	return 0;
}

SEC("fentry/do_nanosleep")
int nanosleep(void *ctx)
{
	struct map_value val = {}, *v;
	struct task_struct *current;

	bpf_map_update_elem(&lru_map, &(int){0}, &val, 0);
	v = bpf_map_lookup_elem(&lru_map, &(int){0});
	if (!v)
		return 0;
	bpf_map_delete_elem(&lru_map, &(int){0});
	current = bpf_get_current_task_btf();
	v->ptr = current;
	pid = current->pid;
	bpf_ktime_get_ns();
	result = !v->ptr;
	return 0;
}

char _license[] SEC("license") = "GPL";

```
