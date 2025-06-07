---
title: "KFunc 'bpf_rdonly_cast'"
description: "This page documents the 'bpf_rdonly_cast' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_rdonly_cast`

<!-- [FEATURE_TAG](bpf_rdonly_cast) -->
[:octicons-tag-24: v6.2](https://github.com/torvalds/linux/commit/a35b9af4ec2c7f69286ef861fd2074a577e354cb)
<!-- [/FEATURE_TAG] -->

## Definition

This kfunc tries to cast the object to a specified type.

The function returns the same `obj` but with `PTR_TO_BTF_ID` with
`btf_id`. The verifier will ensure btf_id being a struct type.

Since the supported type cast may not reflect what the 'obj'
represents, the returned `btf_id` is marked as `PTR_UNTRUSTED`, so
the return value and subsequent pointer chasing cannot be
used as helper/kfunc arguments.

**Signature**

<!-- [KFUNC_DEF] -->
`#!c void *bpf_rdonly_cast(const void *obj__ign, u32 btf_id__k)`
<!-- [/KFUNC_DEF] -->

!!! warning "signature changed"
    The signature of this kfunc has changed in [:octicons-tag-24: v6.9](https://github.com/torvalds/linux/commit/5b268d1ebcdceacf992dfda8f9031d56005a274e). The previous signature was `#!c void *bpf_rdonly_cast(void *obj__ign, u32 btf_id__k)` weak ELF symbols can be used to support both versions.

## Usage

This tries to support use case like below:

`#!c #define skb_shinfo(SKB) ((struct skb_shared_info *)(skb_end_pointer(SKB)))`

where `skb_end_pointer(SKB)` is a `unsigned char *` and needs to
be casted to `struct skb_shared_info *`.


### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_CGROUP_DEVICE`](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
- [`BPF_PROG_TYPE_CGROUP_SOCK`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
- [`BPF_PROG_TYPE_CGROUP_SYSCTL`](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_LWT_IN`](../program-type/BPF_PROG_TYPE_LWT_IN.md)
- [`BPF_PROG_TYPE_LWT_OUT`](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
- [`BPF_PROG_TYPE_LWT_SEG6LOCAL`](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
- [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
- [`BPF_PROG_TYPE_NETFILTER`](../program-type/BPF_PROG_TYPE_NETFILTER.md)
- [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
- [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md)
- [`BPF_PROG_TYPE_SOCKET_FILTER`](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
- [`BPF_PROG_TYPE_SOCK_OPS`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/bc638d8cb5be813d4eeb9f63cce52caaa18f3960) - 
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
- [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

This example shows a use case of `bpf_rdonly_cast`, which in the example is called by the [`bpf_core_cast`](../../ebpf-library/libbpf/ebpf/bpf_core_cast.md) macro.

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

