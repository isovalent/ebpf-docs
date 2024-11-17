---
title: "Libbpf eBPF macro '__arg_nullable'"
description: "This page documents the '__arg_nullable' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `__arg_nullable`

[:octicons-tag-24: v1.4.0](https://github.com/libbpf/libbpf/releases/tag/v1.4.0)

The `__arg_nullable` macros is used to tag a function argument to tell the verifier that its value may be null.

## Definition

`#!c #define __arg_nullable __attribute((btf_decl_tag("arg:nullable")))`

## Usage

This macro can be used to tag a function argument of a [global function](../../../linux/concepts/functions.md#function-by-function-verification) to tell the verifier that it can assume the argument can be `NULL`. It was introduced alongside the [`__arg_trusted`](__arg_trusted.md) macro which tells the verifier that an argument is a trusted pointer to kernel memory. The verifier will by default assume that any trusted pointer argument is never `NULL` (the opposite of normal pointers see [`__arg_nonnull`](__arg_nonnull.md)). Adding the `__arg_nullable` attribute to a trusted pointer argument will tell the verifier that the argument can be `NULL`, which requires the function to add a `NULL` check, but allows the caller to pass a `NULL` pointer. Thus making the argument optional.

### Example

```c hl_lines="2"
__weak int subprog_nullable_task_flavor(
	struct task_struct___local *task __arg_trusted __arg_nullable)
{
	char buf[16];

	if (!task)
		return 0;

	return bpf_copy_from_user_task(&buf, sizeof(buf), NULL, (void *)task, 0);
}

SEC("?uprobe.s")
int flavor_ptr_nullable(void *ctx)
{
	struct task_struct___local *t = (void *)bpf_get_current_task_btf();

	return subprog_nullable_task_flavor(t);
}

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
