---
title: "Libbpf eBPF macro 'BPF_KRETPROBE'"
description: "This page documents the 'BPF_KRETPROBE' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `BPF_KRETPROBE`

[:octicons-tag-24: v0.0.8](https://github.com/libbpf/libbpf/releases/tag/v0.0.8)

The `BPF_KRETPROBE` macro makes it easier to write [kretprobe](../../../linux/program-type/BPF_PROG_TYPE_KPROBE.md) programs.

## Definition

```c
#define BPF_KRETPROBE(name, args...)					    \
name(struct pt_regs *ctx);						    \
static __always_inline typeof(name(0))					    \
____##name(struct pt_regs *ctx, ##args);				    \
typeof(name(0)) name(struct pt_regs *ctx)				    \
{									    \
	_Pragma("GCC diagnostic push")					    \
	_Pragma("GCC diagnostic ignored \"-Wint-conversion\"")		    \
	return ____##name(___bpf_kretprobe_args(args));			    \
	_Pragma("GCC diagnostic pop")					    \
}									    \
static __always_inline typeof(name(0)) ____##name(struct pt_regs *ctx, ##args)
```

## Usage

This macro is useful when writing kprobe programs that attach at the start of a function. Traditionally a program author would have to use the [`PT_REGS_RC`](PT_REGS_RC.md) macro to extract the return value and then manually cast them to the actual type.

The `BPF_KRETPROBE` macro allows you to write your program with an argument list, the macro will do the casting for you. Unlike the [`BPF_KPROBE`](BPF_KPROBE.md) this macro only provides the optional return value. (and the original `struct pt_regs *` context).

!!! note
    The original context will stay available as `ctx`, if you ever wish to access it manually or need to pass it to a helper or kfunc. Therefor, the variable name `ctx` should not be reused in arguments or function body.

### Example

```c hl_lines="2"
SEC("kretprobe/do_unlinkat")
int BPF_KRETPROBE(do_unlinkat_exit, long ret)
{
    pid_t pid;

    pid = bpf_get_current_pid_tgid() >> 32;
    bpf_printk("KPROBE EXIT: pid = %d, ret = %ld\n", pid, ret);
    return 0;
}
```
