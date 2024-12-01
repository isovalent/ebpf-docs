---
title: "Libbpf eBPF macro 'BPF_PROG2'"
description: "This page documents the 'BPF_PROG2' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `BPF_PROG2`

[:octicons-tag-24: v1.1.0](https://github.com/libbpf/libbpf/releases/tag/v1.1.0)

The `BPF_PROG2` macro makes it easier to write programs for program types that receive `[]u64` contexts such as [`BPF_PROG_TYPE_TRACING`](../../../linux/program-type/BPF_PROG_TYPE_TRACING.md) programs. It improves upon the older [`BPF_PROG`](BPF_PROG.md) macro.

## Definition

```c
#define BPF_PROG2(name, args...)						\
name(unsigned long long *ctx);							\
static __always_inline typeof(name(0))						\
____##name(unsigned long long *ctx ___bpf_ctx_decl(args));			\
typeof(name(0)) name(unsigned long long *ctx)					\
{										\
	return ____##name(ctx ___bpf_ctx_arg(args));				\
}										\
static __always_inline typeof(name(0))						\
____##name(unsigned long long *ctx ___bpf_ctx_decl(args))
```

## Usage

This macro is useful when using program types that have a `[]u64` context type (typically written as `unsigned long long *`). 

Conventionally with these program contexts, the arguments to the program are put in this array. So the first argument would be in `ctx[0]`, the second in `ctx[1]`. It is up to the program author to cast them into their actual type.

The `BPF_PROG2` macro allows you to write your program with a normal function signature, the macro will then do the casting for you.

This macro also accounts for an edge case in the <nospell>Sys V</nospell> calling convention. When a function call is made, every argument will be put in specific registers. When a variable is to large to put in a register (8 bytes) it is often put on the stack and a pointer to it is passed instead (pointers being 8 bytes). However, the <nospell>Sys V</nospell> calling convention specifies that if a variable is between 8 and 16 bytes, it may be transferred using 2 registers instead, for a `struct{u64, u64}` for example. Since the context is a translation of the arguments passed, it to can use one or two slots depending on the type. The `BPF_PROG2` handles this in the background, which is what improved over the [`BPF_PROG`](BPF_PROG.md) version of this macro.

!!! note
    The original context will stay available as `ctx`, if you ever wish to access it manually or need to pass it to a helper or kfunc. Therefor, the variable name `ctx` should not be reused in arguments or function body.

### Example

The `bpfptr_t` is actually the following type:
```c
struct {
	union {
		void		*kernel;
		void __user	*user;
	};
	bool		is_kernel : 1;
}
```

Which, with padding is 16 bytes and thus requires `BPF_PROG2` to correctly cast the arguments.

```c hl_lines="2"
SEC("fexit/__sys_bpf")
int BPF_PROG2(sys_bpf, int, cmd, bpfptr_t, uattr, unsigned int, size, int, ret)
{
    bpf_printf("BPF syscall returned with: %d", ret);
    return 0;
}
```
