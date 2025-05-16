---
title: "Libbpf eBPF macro 'PT_REGS_SYSCALL_REGS'"
description: "This page documents the 'PT_REGS_SYSCALL_REGS' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `PT_REGS_SYSCALL_REGS`

[:octicons-tag-24: v0.7.0](https://github.com/libbpf/libbpf/releases/tag/v0.7.0)

The `PT_REGS_SYSCALL_REGS` macro that ensures consistent access to `struct pt_regs` when using a kprobe on a syscall.


## Definition

```c
#define PT_REGS_SYSCALL_REGS(ctx) ((struct pt_regs *)PT_REGS_PARM1(ctx))
```

### Usage

This macro is useful when placing kprobes on syscalls. On some CPU architectures the kernel will use a so called syscall wrapper (indicated by the `CONFIG_ARCH_HAS_SYSCALL_WRAPPER` kernel config). The syscall calling convention (which registers are used for what parameter) can differ from the normal calling convention. A syscall wrapper translates syscall calling convention to normal calling convention before invoking the actual syscall function.

You might need to attach to this wrapper since the function that is called once in the right calling convention might be inlined. The wrapper is handed a `struct pt_regs *` as its first argument by the syscall interrupt handler, which contains the values passed by userspace. When a kprobe executes it also causes an interrupt yielding a second `struct pt_regs *` with the registers as the wrapper gets them.

```
kprobe arg1 -> pt_regs (wrapper) arg1 -> pt_regs (syscall)
```

Since its likely the syscall arguments that are of interest we need to unwrap. Which is what `PT_REGS_SYSCALL_REGS` does. You give it the kprobe `struct pt_regs *` and it returns the syscall `struct pt_regs *`.

The architecture for which the eBPF program is compiled is determined by setting one of the `__TARGET_ARCH_{arch}` macros. These are typically set by passing a flag to the compiler, such as `-D__TARGET_ARCH_x86` for x86. This allows for easy cross-compilation of eBPF programs for different architectures by changing the compiler invocation.

### Example

This example places a kprobe on the open syscall wrapper. We get the syscall registers with `PT_REGS_SYSCALL_REGS`.
In order to access the second parameter of the syscall we need to probe the syscall `struct pt_regs *`. We use the [`PT_REGS_PARM2_SYSCALL`](PT_REGS_PARM_SYSCALL.md) macro to get the correct register for param 2 according to the syscall calling convention.

```c hl_lines="4"
SEC("kprobe/" SYS_PREFIX "open")
int trace_sys_open(struct pt_regs *ctx)
{
    struct pt_regs *realregs = PT_REGS_SYSCALL_REGS(ctx);
    
    int flags;
    [bpf_probe_read_kernel](../../../linux/helper-function/bpf_probe_read_kernel.md)(&flags, sizeof(flags), &[PT_REGS_PARM2_SYSCALL](PT_REGS_PARM_SYSCALL.md)(realregs));
    if (flags & O_CREAT) {
        bpf_printk("Creating a file");
    }

    return 0;
}
```


