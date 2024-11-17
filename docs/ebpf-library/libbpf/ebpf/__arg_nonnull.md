---
title: "Libbpf eBPF macro '__arg_nonnull'"
description: "This page documents the '__arg_nonnull' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `__arg_nonnull`

[:octicons-tag-24: v1.4.0](https://github.com/libbpf/libbpf/releases/tag/v1.4.0)

The `__arg_nonnull` macros is used to tag a function argument to tell the verifier that its value may never be null.

## Definition

`#!c #define __arg_nonnull __attribute((btf_decl_tag("arg:nonnull")))`

## Usage

This macro can be used to tag a function argument of a [global function](../../../linux/concepts/functions.md#function-by-function-verification) to tell the verifier that it can assume the argument is never null, and to enforce this on the call site. Since global functions can be verifier out of order, the verifier will always assume that a pointer argument may contain a `NULL` value and will force you to implement a check for this. As program author you may know that you always do a `NULL` check on all of your callsites. In that case you can add the `__arg_nonnull` attribute to the function argument, the verifier will assume the argument is never `NULL` while verifying the function and will not enforce that any pointer passed into the function can not be `NULL`.

### Example

Note how the first function is not `static`, thus global. The `__noinline __weak` attributes are added to force the compiler to emit a separate function instead of inlining it. That is not needed in actual usage.

```c hl_lines="1"
__noinline __weak int subprog_nonnull_ptr_good(int *p1 __arg_nonnull, int *p2 __arg_nonnull)
{
	return (*p1) * (*p2); /* good, no need for NULL checks */
}

int x = 47;

SEC("?raw_tp")
int arg_tag_nonnull_ptr_good(void *ctx)
{
	int y = 74;

	return subprog_nonnull_ptr_good(&x, &y);
}
```
