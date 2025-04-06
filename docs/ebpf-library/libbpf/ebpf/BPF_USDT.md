---
title: "Libbpf eBPF macro 'BPF_USDT'"
description: "This page documents the 'BPF_USDT' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `BPF_USDT`

[:octicons-tag-24: v0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)

The `BPF_USDT` macro is similar to [`BPF_PROG`](BPF_PROG.md) but for [USDT](../../../linux/concepts/usdt.md) probes.

## Definition

```c
#define BPF_USDT(name, args...)						    \
name(struct pt_regs *ctx);						    \
static __always_inline typeof(name(0))					    \
____##name(struct pt_regs *ctx, ##args);				    \
typeof(name(0)) name(struct pt_regs *ctx)				    \
{									    \
        _Pragma("GCC diagnostic push")					    \
        _Pragma("GCC diagnostic ignored \"-Wint-conversion\"")		    \
        return ____##name(___bpf_usdt_args(args));			    \
        _Pragma("GCC diagnostic pop")					    \
}									    \
static __always_inline typeof(name(0))					    \
____##name(struct pt_regs *ctx, ##args)
```

## Usage

A USDT program is simply a uprobe that is expected to be attached to a USDT tracepoint. Such tracepoints can pass arguments, however, unlike arguments to a function which follow an ABI, the arguments to a tracepoint can live anywhere in the process. The tracepoint describes them using [GAS(GNU assembler) operands](https://en.wikibooks.org/wiki/X86_Assembly/GNU_assembly_syntax). This macro allows you to define your programs as these arguments were actually passed to the program like normal arguments.

The `BPF_USDT` will call [`bpf_usdt_arg`](bpf_usdt_arg.md) for every defined argument and cast it to abstract away the complexity. This functionality requires the loader to translate the GAS operands into a format understandable by [`bpf_usdt_arg`](bpf_usdt_arg.md), this is taken care of by [`bpf_program__attach_usdt`](../userspace/bpf_program__attach_usdt.md).

### Example

```c hl_lines="4"
// SPDX-License-Identifier: GPL-2.0
/* Copyright (c) 2022 Meta Platforms, Inc. and affiliates. */
SEC("usdt/./urandom_read:urand:read_without_sema")
int BPF_USDT(urand_read_without_sema, int iter_num, int iter_cnt, int buf_sz)
{
	if (urand_pid != (bpf_get_current_pid_tgid() >> 32))
		return 0;

	__sync_fetch_and_add(&urand_read_without_sema_call_cnt, 1);
	__sync_fetch_and_add(&urand_read_without_sema_buf_sz_sum, buf_sz);

	return 0;
}
```

[Source](https://github.com/torvalds/linux/blob/master/tools/testing/selftests/bpf/progs/test_urandom_usdt.c)
