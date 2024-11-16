---
title: "Libbpf eBPF macro 'container_of'"
description: "This page documents the 'offsetof' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `container_of`

[:octicons-tag-24: v0.0.9](https://github.com/libbpf/libbpf/releases/tag/v0.0.9)

The `container_of` macro is used to get the address of the structure that contains the pointer we have.

## Definition

```c
#define container_of(ptr, type, member)                 \
    ({                                                  \
        void *__mptr = (void *)(ptr);                   \
        ((type *)(__mptr - offsetof(type, member)));    \
    })
```

## Usage

It sometimes happens that you obtain a pointer to a type, but this type is actually a member of a larger structure. The `container_of` macro allows you to get the address of the structure that contains the pointer you have, which allows you to access other members of the structure.

### Example

An artificial example of how `container_of` can be used:
```c hl_lines="12"
struct a {
    int x;
    int y;
    struct b b;
}

struct b {
    int z;
}

int some_func(struct b *b) {
    struct a *a = container_of(b, struct a, b);
    return a->y;
}
```

An example of how `container_of` can be used in an eBPF program:
```c hl_lines="11"
// SPDX-License-Identifier: GPL-2.0
/* Copyright (c) 2020 Facebook */
#include "bpf_iter.h"
#include "bpf_tracing_net.h"
#include <bpf/bpf_helpers.h>

char _license[] SEC("license") = "GPL";

static __attribute__((noinline)) struct inode *SOCK_INODE(struct socket *socket)
{
	return &container_of(socket, struct socket_alloc, socket)->vfs_inode;
}

SEC("iter/netlink")
int dump_netlink(struct bpf_iter__netlink *ctx)
{
	struct seq_file *seq = ctx->meta->seq;
	struct netlink_sock *nlk = ctx->sk;
	unsigned long group, ino;
	struct inode *inode;
	struct socket *sk;
	struct sock *s;

	if (nlk == (void *)0)
		return 0;

	if (ctx->meta->seq_num == 0)
		BPF_SEQ_PRINTF(seq, "sk               Eth Pid        Groups   "
				    "Rmem     Wmem     Dump  Locks    Drops    "
				    "Inode\n");

	s = &nlk->sk;
	BPF_SEQ_PRINTF(seq, "%pK %-3d ", s, s->sk_protocol);

	if (!nlk->groups)  {
		group = 0;
	} else {
		/* FIXME: temporary use bpf_probe_read_kernel here, needs
		 * verifier support to do direct access.
		 */
		bpf_probe_read_kernel(&group, sizeof(group), &nlk->groups[0]);
	}
	BPF_SEQ_PRINTF(seq, "%-10u %08x %-8d %-8d %-5d %-8d ",
		       nlk->portid, (u32)group,
		       s->sk_rmem_alloc.counter,
		       s->sk_wmem_alloc.refs.counter - 1,
		       nlk->cb_running, s->sk_refcnt.refs.counter);

	sk = s->sk_socket;
	if (!sk) {
		ino = 0;
	} else {
		/* FIXME: container_of inside SOCK_INODE has a forced
		 * type conversion, and direct access cannot be used
		 * with current verifier.
		 */
		inode = SOCK_INODE(sk);
		bpf_probe_read_kernel(&ino, sizeof(ino), &inode->i_ino);
	}
	BPF_SEQ_PRINTF(seq, "%-8u %-8lu\n", s->sk_drops.counter, ino);

	return 0;
}
```
