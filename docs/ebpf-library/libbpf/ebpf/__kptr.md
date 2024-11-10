---
title: "Libbpf eBPF macro '__kptr'"
description: "This page documents the '__kptr' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `__kptr`

[:octicons-tag-24: v0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)

The `__kptr` macros is used to tag a pointer to tell the verifier it holds pointers to kernel memory.

## Definition

`#!c #define __kptr __attribute__((btf_type_tag("kptr")))`

## Usage

This macro can used on type definitions for both global variables and fields in map values. It informs the verifier that the pointer is a kernel pointer.

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome 

### Example

```c hl_lines="2"
struct cpumask_map_value {
        struct bpf_cpumask __kptr * cpumask;
};

struct array_map {
        __uint(type, BPF_MAP_TYPE_ARRAY);
        __type(key, int);
        __type(value, struct cpumask_map_value);
        __uint(max_entries, 65536);
} cpumask_map SEC(".maps");

static int cpumask_map_insert(struct bpf_cpumask *mask, u32 pid)
{
        struct cpumask_map_value local, *v;
        long status;
        struct bpf_cpumask *old;
        u32 key = pid;

        local.cpumask = NULL;
        status = bpf_map_update_elem(&cpumask_map, &key, &local, 0);
        if (status) {
                bpf_cpumask_release(mask);
                return status;
        }

        v = bpf_map_lookup_elem(&cpumask_map, &key);
        if (!v) {
                bpf_cpumask_release(mask);
                return -ENOENT;
        }

        old = bpf_kptr_xchg(&v->cpumask, mask);
        if (old)
                bpf_cpumask_release(old);

        return 0;
}

/**
    * A sample tracepoint showing how a task's cpumask can be queried and
    * recorded as a kptr.
    */
SEC("tp_btf/task_newtask")
int BPF_PROG(record_task_cpumask, struct task_struct *task, u64 clone_flags)
{
        struct bpf_cpumask *cpumask;
        int ret;

        cpumask = bpf_cpumask_create();
        if (!cpumask)
                return -ENOMEM;

        if (!bpf_cpumask_full(task->cpus_ptr))
                bpf_printk("task %s has CPU affinity", task->comm);

        bpf_cpumask_copy(cpumask, task->cpus_ptr);
        return cpumask_map_insert(cpumask, task->pid);
}
```
