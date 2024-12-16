---
title: "Libbpf eBPF macro 'bpf_core_cast'"
description: "This page documents the 'bpf_core_cast' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `bpf_core_cast`

[:octicons-tag-24: v1.4.0](https://github.com/libbpf/libbpf/releases/tag/v1.4.0)

The `bpf_core_cast` macro abstracts away [`bpf_probe_read_kernel`](../../../linux/helper-function/bpf_probe_read_kernel.md) call and captures offset relocation.

## Definition

```c
#define bpf_core_cast(ptr, type)					    \
	((typeof(type) *)bpf_rdonly_cast((ptr), bpf_core_type_id_kernel(type)))
```

## Usage

The `bpf_core_cast` macro casts the provided pointer `ptr` into a pointer to a specified `type` in such a way that BPF verifier will become aware of associated kernel-side BTF type. This allows to access members of kernel types directly without the need to use [`BPF_CORE_READ`](BPF_CORE_READ.md) macros.

This macro calls the [`bpf_rdonly_cast`](../../../linux/kfuncs/bpf_rdonly_cast.md) kfunc with the kernel-side BTF type ID of the provided `type`.

### Example

```c hl_lines="29"
// SPDX-License-Identifier: GPL-2.0
// Copyright (c) 2018 Facebook

#include <vmlinux.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_core_read.h>
#include "bpf_tracing_net.h"
#define NUM_CGROUP_LEVELS	4

__u64 cgroup_ids[NUM_CGROUP_LEVELS];
__u16 dport;

static __always_inline void log_nth_level(struct __sk_buff *skb, __u32 level)
{
	/* [1] &level passed to external function that may change it, it's
	 *     incompatible with loop unroll.
	 */
	cgroup_ids[level] = bpf_skb_ancestor_cgroup_id(skb, level);
}

SEC("tc")
int log_cgroup_id(struct __sk_buff *skb)
{
	struct sock *sk = (void *)skb->sk;

	if (!sk)
		return TC_ACT_OK;

	sk = bpf_core_cast(sk, struct sock);
	if (sk->sk_protocol == IPPROTO_UDP && sk->sk_dport == dport) {
		log_nth_level(skb, 0);
		log_nth_level(skb, 1);
		log_nth_level(skb, 2);
		log_nth_level(skb, 3);
	}

	return TC_ACT_OK;
}

char _license[] SEC("license") = "GPL";
```
