---
title: "KFunc 'bpf_sock_addr_set_sun_path'"
description: "This page documents the 'bpf_sock_addr_set_sun_path' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_sock_addr_set_sun_path`

<!-- [FEATURE_TAG](bpf_sock_addr_set_sun_path) -->
[:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4)
<!-- [/FEATURE_TAG] -->

Modify the socket address of a socket.

## Definition

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_sock_addr_set_sun_path(struct bpf_sock_addr_kern *sa_kern, const u8 *sun_path, u32 sun_path__sz)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_CGROUP_DEVICE`](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
- [`BPF_PROG_TYPE_CGROUP_SOCK`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
- [`BPF_PROG_TYPE_CGROUP_SYSCTL`](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
- [`BPF_PROG_TYPE_SOCK_OPS`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
<!-- [/KFUNC_PROG_REF] -->

### Example

```c
// SPDX-License-Identifier: GPL-2.0
/* Copyright (c) 2022 Meta Platforms, Inc. and affiliates. */

#include "vmlinux.h"

#include <string.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_core_read.h>
#include "bpf_kfuncs.h"

__u8 SERVUN_REWRITE_ADDRESS[] = "\0bpf_cgroup_unix_test_rewrite";

SEC("cgroup/connect_unix")
int connect_unix_prog(struct bpf_sock_addr *ctx)
{
	struct bpf_sock_addr_kern *sa_kern = bpf_cast_to_kern_ctx(ctx);
	struct sockaddr_un *sa_kern_unaddr;
	__u32 unaddrlen = offsetof(struct sockaddr_un, sun_path) +
			  sizeof(SERVUN_REWRITE_ADDRESS) - 1;
	int ret;

	/* Rewrite destination. */
	ret = bpf_sock_addr_set_sun_path(sa_kern, SERVUN_REWRITE_ADDRESS,
					 sizeof(SERVUN_REWRITE_ADDRESS) - 1);
	if (ret)
		return 0;

	if (sa_kern->uaddrlen != unaddrlen)
		return 0;

	sa_kern_unaddr = bpf_rdonly_cast(sa_kern->uaddr,
						bpf_core_type_id_kernel(struct sockaddr_un));
	if (memcmp(sa_kern_unaddr->sun_path, SERVUN_REWRITE_ADDRESS,
			sizeof(SERVUN_REWRITE_ADDRESS) - 1) != 0)
		return 0;

	return 1;
}

char _license[] SEC("license") = "GPL";
```

