---
title: "Libbpf eBPF macro 'BPF_SEQ_PRINTF'"
description: "This page documents the 'BPF_SEQ_PRINTF' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `BPF_SEQ_PRINTF`

[:octicons-tag-24: v0.0.9](https://github.com/libbpf/libbpf/releases/tag/v0.0.9)

!!! note
    This macro was moved from `bpf_tracing.h` to `bpf_helpers.h` in [:octicons-tag-24: v0.5.0](https://github.com/libbpf/libbpf/releases/tag/v0.5.0)

The `BPF_SEQ_PRINTF` macro is used to make printing to bpf iterator sequence files easier.

## Definition

```c
/*
 * BPF_SEQ_PRINTF to wrap bpf_seq_printf to-be-printed values
 * in a structure.
 */
#define BPF_SEQ_PRINTF(seq, fmt, args...)			\
({								\
	static const char ___fmt[] = fmt;			\
	unsigned long long ___param[___bpf_narg(args)];		\
								\
	_Pragma("GCC diagnostic push")				\
	_Pragma("GCC diagnostic ignored \"-Wint-conversion\"")	\
	___bpf_fill(___param, args);				\
	_Pragma("GCC diagnostic pop")				\
								\
	bpf_seq_printf(seq, ___fmt, sizeof(___fmt),		\
		       ___param, sizeof(___param));		\
})
```

## Usage

This macro is a wrapper around the [`bpf_seq_printf`](../../../linux/helper-function/bpf_seq_printf.md) helper. It places the literal format string in a global variable, this is necessary to get the compiler to emit code that will be accepted by the verifier.

### Example
```c hl_lines="14 25 26"
SEC("iter/task_file")
int dump_task_file(struct bpf_iter__task_file *ctx)
{
    struct seq_file *seq = ctx->meta->seq;
    struct task_struct *task = ctx->task;
    struct file *file = ctx->file;
    __u32 fd = ctx->fd;

    if (task == NULL || file == NULL)
        return 0;

    if (ctx->meta->seq_num == 0) {
        count = 0;
        BPF_SEQ_PRINTF(seq, "    tgid      gid       fd      file\n");
    }

    if (tgid == task->tgid && task->tgid != task->pid)
        count++;

    if (last_tgid != task->tgid) {
        last_tgid = task->tgid;
        unique_tgid_count++;
    }

    BPF_SEQ_PRINTF(seq, "%8d %8d %8d %lx\n", task->tgid, task->pid, fd,
            (long)file->f_op);
    return 0;
}
```
