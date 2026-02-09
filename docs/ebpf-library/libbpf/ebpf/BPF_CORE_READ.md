---
title: "Libbpf eBPF macro 'BPF_CORE_READ'"
description: "This page documents the 'BPF_CORE_READ' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `BPF_CORE_READ`

[:octicons-tag-24: v0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)

The `BPF_CORE_READ` macro is used to simplify BPF CO-RE relocatable read, especially when there are few pointer chasing steps.

## Definition

```c
#define BPF_CORE_READ(src, a, ...) ({					    \
	___type((src), a, ##__VA_ARGS__) __r;				    \
	BPF_CORE_READ_INTO(&__r, (src), a, ##__VA_ARGS__);		    \
	__r;								    \
})
```

## Usage

`BPF_CORE_READ` is used to simplify BPF CO-RE relocatable read, especially when there are few pointer chasing steps. E.g., what in non-BPF world (or in BPF w/ BCC) would be something like:

`#!c int x = s->a.b.c->d.e->f->g;`

can be succinctly achieved using `BPF_CORE_READ` as:

`#!c int x = BPF_CORE_READ(s, a.b.c, d.e, f, g);`

`BPF_CORE_READ` will decompose above statement into 4 [`bpf_core_read`](bpf_core_read.md) (BPF CO-RE relocatable [`bpf_probe_read_kernel`](../../../linux/helper-function/bpf_probe_read_kernel.md) wrapper) calls, logically equivalent to:

 1. `const void *__t = s->a.b.c;`
 2. `__t = __t->d.e;`
 3. `__t = __t->f;`
 4. `return __t->g;`

Equivalence is logical, because there is a heavy type casting/preservation involved, as well as all the reads are happening through [`bpf_probe_read_kernel`](../../../linux/helper-function/bpf_probe_read_kernel.md) calls using `__builtin_preserve_access_index`() to emit CO-RE relocations.

!!! note
    Only up to 9 "field accessors" are supported, which should be more than enough for any practical purpose.

### Example

```c hl_lines="32 33"
/* Copyright (c) 2013-2015 PLUMgrid, http://plumgrid.com
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of version 2 of the GNU General Public
 * License as published by the Free Software Foundation.
 */
#include "vmlinux.h"
#include "net_shared.h"
#include <linux/version.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_tracing.h>

/* kprobe is NOT a stable ABI
 * kernel functions can be removed, renamed or completely change semantics.
 * Number of arguments and their positions can change, etc.
 * In such case this bpf+kprobe example will no longer be meaningful
 */
SEC("kprobe.multi/__netif_receive_skb_core*")
int bpf_prog1(struct pt_regs *ctx)
{
	/* attaches to kprobe __netif_receive_skb_core,
	 * looks for packets on loopback device and prints them
	 * (wildcard is used for avoiding symbol mismatch due to optimization)
	 */
	char devname[IFNAMSIZ];
	struct net_device *dev;
	struct sk_buff *skb;
	int len;

	bpf_core_read(&skb, sizeof(skb), (void *)PT_REGS_PARM1(ctx));
	dev = BPF_CORE_READ(skb, dev);
	len = BPF_CORE_READ(skb, len);

	BPF_CORE_READ_STR_INTO(&devname, dev, name);

	if (devname[0] == 'l' && devname[1] == 'o') {
		char fmt[] = "skb %p len %d\n";
		/* using bpf_trace_printk() for DEBUG ONLY */
		bpf_trace_printk(fmt, sizeof(fmt), skb, len);
	}

	return 0;
}

char _license[] SEC("license") = "GPL";
u32 _version SEC("version") = LINUX_VERSION_CODE;
```
