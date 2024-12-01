---
title: "Libbpf eBPF macro 'BPF_KSYSCALL'"
description: "This page documents the 'BPF_KSYSCALL' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `BPF_KSYSCALL`

[:octicons-tag-24: v0.7.0](https://github.com/libbpf/libbpf/releases/tag/v0.7.0)

!!! note
    Originally this macro was called [`BPF_KPROBE_SYSCALL`](BPF_KPROBE_SYSCALL.md), in Libbpf [:octicons-tag-24: v1.0.0](https://github.com/libbpf/libbpf/releases/tag/v1.0.0) it was renamed. The old macro now aliases to this one for backwards compatibility.

The `BPF_KSYSCALL` macro makes it easier to write [kprobe](../../../linux/program-type/BPF_PROG_TYPE_KPROBE.md) programs that attach to syscalls.

## Definition

```c
#define BPF_KSYSCALL(name, args...)					    \
name(struct pt_regs *ctx);						    \
extern _Bool LINUX_HAS_SYSCALL_WRAPPER __kconfig;			    \
static __always_inline typeof(name(0))					    \
____##name(struct pt_regs *ctx, ##args);				    \
typeof(name(0)) name(struct pt_regs *ctx)				    \
{									    \
	struct pt_regs *regs = LINUX_HAS_SYSCALL_WRAPPER		    \
			       ? (struct pt_regs *)PT_REGS_PARM1(ctx)	    \
			       : ctx;					    \
	_Pragma("GCC diagnostic push")					    \
	_Pragma("GCC diagnostic ignored \"-Wint-conversion\"")		    \
	if (LINUX_HAS_SYSCALL_WRAPPER)					    \
		return ____##name(___bpf_syswrap_args(args));		    \
	else								    \
		return ____##name(___bpf_syscall_args(args));		    \
	_Pragma("GCC diagnostic pop")					    \
}									    \
static __always_inline typeof(name(0))					    \
____##name(struct pt_regs *ctx, ##args)
```

## Usage

This macro is useful when writing kprobe programs that attach to syscalls. On some CPU architectures the kernel will use a so called syscall wrapper (indicated by the `CONFIG_ARCH_HAS_SYSCALL_WRAPPER` kernel config). These wrappers change the syscall calling convention, so the actual `struct pt_regs` are in a different location then normally expected.

Traditionally a program author would have to use the [`PT_REGS_SYSCALL_REGS`](PT_REGS_SYSCALL_REGS.md) macros to extract a given parameter from the context and then manually cast them to the actual type.

The `BPF_KSYSCALL` macro allows you to write your program with an argument list, the macro will do the casting for you and accounts for the syscall wrapping.

!!! note
    The original context will stay available as `ctx`, if you ever wish to access it manually or need to pass it to a helper or kfunc. Therefor, the variable name `ctx` should not be reused in arguments or function body.

!!! warning
    At the moment `BPF_KSYSCALL` does not transparently handle all the calling convention quirks for the following syscalls:

    * `mmap()` - `__ARCH_WANT_SYS_OLD_MMAP`.
    * `clone()` - `CONFIG_CLONE_BACKWARDS`, `CONFIG_CLONE_BACKWARDS2`, and `CONFIG_CLONE_BACKWARDS3`.
    * socket-related syscalls - `__ARCH_WANT_SYS_SOCKETCALL`.
    * <nospell>compat</nospell> syscalls.

    This may or may not change in the future. User needs to take extra measures to handle such quirks explicitly, if necessary.

!!! note
    This macro relies on BPF [CO-RE](../../../concepts/core.md) support and virtual [`__kconfig`](__kconfig.md) `extern`s.

### Example

```c hl_lines="2"
SEC("ksyscall/write")
int BPF_KSYSCALL(bpf_prog3, unsigned int fd, const char *buf, size_t count)
{
	long init_val = 1;
	long *value;
	struct hist_key key;

	key.index = log2l(count);
	key.pid_tgid = bpf_get_current_pid_tgid();
	key.uid_gid = bpf_get_current_uid_gid();
	bpf_get_current_comm(&key.comm, sizeof(key.comm));

	value = bpf_map_lookup_elem(&my_hist_map, &key);
	if (value)
		__sync_fetch_and_add(value, 1);
	else
		bpf_map_update_elem(&my_hist_map, &key, &init_val, BPF_ANY);
	return 0;
}
```
