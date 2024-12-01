---
title: "Libbpf eBPF macro 'PT_REGS_PARM'"
description: "This page documents the 'PT_REGS_PARM' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `PT_REGS_PARM`

[:octicons-tag-24: v0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)

The `PT_REGS_PARM{1-8}` macros make it easy to extract an argument from `struct pt_regs` style contexts in an architecture-independent way.

## Usage

Since the `struct pt_regs` type represents the state of the CPU registers, it is different for every architecture. The `PT_REGS_PARM{1-8}` macros translates the argument number to the correct register in the `struct pt_regs` type depending on the calling convention of the architecture.

The architecture for which the eBPF program is compiled is determined by setting one of the `__TARGET_ARCH_{arch}` macros. These are typically set by passing a flag to the compiler, such as `-D__TARGET_ARCH_x86` for x86. This allows for easy cross-compilation of eBPF programs for different architectures by changing the compiler invocation.

### Example

```c hl_lines="20"
/* Copyright (c) 2013-2015 PLUMgrid, http://plumgrid.com
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of version 2 of the GNU General Public
 * License as published by the Free Software Foundation.
 */

SEC("kprobe.multi/__netif_receive_skb_core*")
int bpf_prog1(struct pt_regs *ctx)
{
	/* attaches to kprobe __netif_receive_skb_core,
	 * looks for packets on loobpack device and prints them
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
```
