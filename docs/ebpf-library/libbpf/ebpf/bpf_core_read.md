---
title: "Libbpf eBPF macro 'bpf_core_read'"
description: "This page documents the 'bpf_core_read' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `bpf_core_read`

[:octicons-tag-24: v0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)

The `bpf_core_read` macro abstracts away [`bpf_probe_read_kernel`](../../../linux/helper-function/bpf_probe_read_kernel.md) call and captures offset relocation.

## Definition

```c
#define bpf_core_read(dst, sz, src)					    \
	bpf_probe_read_kernel(dst, sz, (const void *)__builtin_preserve_access_index(src))
```

## Usage

The `bpf_core_read` abstracts away [`bpf_probe_read_kernel`](../../../linux/helper-function/bpf_probe_read_kernel.md) call and captures offset relocation. 

This relocation allows libbpf to adjust BPF instruction to use correct actual field offset, based on target kernel BTF type that matches original (local) BTF, used to record relocation.

`dst` is a pointer to the destination buffer, `sz` is the size of the buffer, and `src` an expression that evaluates to the pointer to kernel memory.

Given a structure like:

```c
struct a {
    int b;
    struct {
        struct some_value *d;
    } c;
};
```

Where you wish to get the value of `d`, you would call `bpf_core_read` as:

```c
struct some_value dst;
bpf_core_read(&dst, sizeof(dst), a.c.d);
```

Since `src` contains the field accesses, these fill be stored in the relocation entry so the actual offset can be adjusted. It is therefore important to not do the field access outside of the `bpf_core_read`:

```c
/* Incorrect */
struct some_value dst;
struct some_value *ptr = a.c.d;
bpf_core_read(&dst, sizeof(dst), ptr);
```

It is also important to remember that you should avoid using expressions that implicitly dereference pointers. For example, an expression like `a->b.c` will dereference the pointer `a` without the use of [`bpf_probe_read_kernel`](../../../linux/helper-function/bpf_probe_read_kernel.md). Whereas `&a->b` will not dereference the pointer `a`.

To do a CO-RE read with dereferences, use the [`BPF_CORE_READ`](BPF_CORE_READ.md) macro, which chains multiple `bpf_core_read` calls together for you like: `BPF_CORE_READ(a, b.c)`.


### Example

```c hl_lines="5"
struct task_struct *task = (void *)bpf_get_current_task();
struct task_struct *parent_task;
int err;

err = bpf_core_read(&parent_task, sizeof(void *), &task->parent);
if (err) {
    /* handle error */
}

/* parent_task contains the value of task->parent pointer */
```
