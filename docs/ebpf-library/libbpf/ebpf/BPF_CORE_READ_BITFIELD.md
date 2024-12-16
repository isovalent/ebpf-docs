---
title: "Libbpf eBPF macro 'BPF_CORE_READ_BITFIELD'"
description: "This page documents the 'BPF_CORE_READ_BITFIELD' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `BPF_CORE_READ_BITFIELD`

[:octicons-tag-24: v0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)

The `BPF_CORE_READ_BITFIELD` macro extracts a bitfield from a given structure in a CO-RE relocatable way.

## Definition

```c
#define BPF_CORE_READ_BITFIELD(s, field) ({				      \
	const void *p = (const void *)s + __CORE_RELO(s, field, BYTE_OFFSET); \
	unsigned long long val;						      \
									      \
	/* This is a so-called barrier_var() operation that makes specified   \
	 * variable "a black box" for optimizing compiler.		      \
	 * It forces compiler to perform BYTE_OFFSET relocation on p and use  \
	 * its calculated value in the switch below, instead of applying      \
	 * the same relocation 4 times for each individual memory load.       \
	 */								      \
	asm volatile("" : "=r"(p) : "0"(p));				      \
									      \
	switch (__CORE_RELO(s, field, BYTE_SIZE)) {			      \
	case 1: val = *(const unsigned char *)p; break;			      \
	case 2: val = *(const unsigned short *)p; break;		      \
	case 4: val = *(const unsigned int *)p; break;			      \
	case 8: val = *(const unsigned long long *)p; break;		      \
	default: val = 0; break;					      \
	}								      \
	val <<= __CORE_RELO(s, field, LSHIFT_U64);			      \
	if (__CORE_RELO(s, field, SIGNED))				      \
		val = ((long long)val) >> __CORE_RELO(s, field, RSHIFT_U64);  \
	else								      \
		val = val >> __CORE_RELO(s, field, RSHIFT_U64);		      \
	val;								      \
})
```

## Usage

`BPF_CORE_READ_BITFIELD` extract bitfield, identified by `s->field`, and return its value as u64. This is a variant of the [`BPF_CORE_READ_BITFIELD_PROBED`](BPF_CORE_READ_BITFIELD_PROBED.md) macro. This macro is using direct memory reads and should be used from BPF program types that support such functionality (e.g., typed [raw tracepoints](../../../linux/program-type/BPF_PROG_TYPE_TRACING.md#raw-tracepoint)).

### Example

```c
// SPDX-License-Identifier: GPL-2.0
/* Copyright (c) 2019 Facebook */

struct tcp_sock {
    u32	snd_cwnd;	/* Sending congestion window		*/
    /* [...] */
    u8	chrono_type : 2,	/* current chronograph type */
		repair      : 1,
		tcp_usec_ts : 1, /* TSval values in usec */
		is_sack_reneg:1,    /* in recovery from loss with SACK reneg? */
		is_cwnd_limited:1;/* forward progress limited by snd_cwnd? */
    /* [...] */
    u32	max_packets_out;  /* max packets_out in last window */
}

static inline bool tcp_is_cwnd_limited(const struct sock *sk)
{
	const struct tcp_sock *tp = tcp_sk(sk);

	/* If in slow start, ensure cwnd grows to twice what was ACKed. */
	if (tcp_in_slow_start(tp))
		return tp->snd_cwnd < 2 * tp->max_packets_out;

	return !!BPF_CORE_READ_BITFIELD(tp, is_cwnd_limited);
}
```
