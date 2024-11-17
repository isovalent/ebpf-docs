---
title: "Libbpf eBPF macro '__arg_trusted'"
description: "This page documents the '__arg_trusted' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `__arg_trusted`

[:octicons-tag-24: v1.4.0](https://github.com/libbpf/libbpf/releases/tag/v1.4.0)

The `__arg_trusted` macros is used to tag a function argument to tell the verifier that its value is a trusted pointer to kernel memory.

## Definition

`#!c #define __arg_trusted __attribute((btf_decl_tag("arg:trusted")))`

## Usage

This macro can be used to tag a function argument of a [global function](../../../linux/concepts/functions.md#function-by-function-verification) to tell the verifier that its value is a trusted pointer to kernel memory. Similarly to how [`__kptr`](__kptr.md) is used on map values and global variables. By default the verifier will assume the argument is never `NULL`, this can be changed by adding the [`__arg_nullable`](__arg_nullable.md) attribute to the argument. The verifier will enforce that a valid trusted pointer is passed to the function on the callsite.

### Example

```c hl_lines="1"
__weak int subprog_nonnull_task_flavor(struct task_struct___local *task __arg_trusted)
{
    char buf[16];

    return bpf_copy_from_user_task(&buf, sizeof(buf), NULL, (void *)task, 0);
}

SEC("?uprobe.s")
int flavor_ptr_nonnull(void *ctx)
{
    struct task_struct *t = bpf_get_current_task_btf();

    return subprog_nonnull_task_flavor((void *)t);
}
```
