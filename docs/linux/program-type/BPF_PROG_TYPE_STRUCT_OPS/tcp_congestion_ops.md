---
title: "Struct ops 'tcp_congestion_ops'"
description: "This page documents the 'tcp_congestion_ops' struct ops, its semantics, capabilities, and limitations."
---
# Struct ops `tcp_congestion_ops`

[:octicons-tag-24: v5.6](https://github.com/torvalds/linux/commit/0baf26b0fcd74bbfcef53c5d5e8bad2b99c8d0d2)

TCP congestion ops allows you to implement a TCP congestion control algorithm in BPF.

## Usage

Implementing a congestion control algorithm via BPF can be useful for testing new congestion control algorithms without having to modify the kernel or implementing an algorithm that is optimized for a specific use case but not generic enough to be included in the kernel.

## Fields and ops

```c
struct tcp_congestion_ops {
    char     [name](#name)[TCP_CA_NAME_MAX];
	u32      [key](#key);
	u32      [flags](#flags);
	u32    (*[ssthresh](#ssthresh))(struct sock *sk);
	void   (*[cong_avoid](#cong_avoid))(struct sock *sk, u32 ack, u32 acked);
	void   (*[set_state](#set_state))(struct sock *sk, u8 new_state);
	void   (*[cwnd_event](#cwnd_event))(struct sock *sk, [enum tcp_ca_event](#enum-tcp_ca_event) ev);
	void   (*[in_ack_event](#in_ack_event))(struct sock *sk, u32 flags);
	void   (*[pkts_acked](#pkts_acked))(struct sock *sk, const [struct ack_sample](#struct-ack_sample) *sample);
	u32    (*[min_tso_segs](#min_tso_segs))(struct sock *sk);
	void   (*[cong_control](#cong_control))(struct sock *sk, u32 ack, int flag, const [struct rate_sample](#struct-rate_sample) *rs);
	u32    (*[undo_cwnd](#undo_cwnd))(struct sock *sk);
	u32    (*[sndbuf_expand](#sndbuf_expand))(struct sock *sk);
	size_t (*[get_info](#get_info))(struct sock *sk, u32 ext, int *attr, [union tcp_cc_info](#union-tcp_cc_info) *info);
	void   (*[init](#init))(struct sock *sk);
	void   (*[release](#release))(struct sock *sk);
};
```

### `name`

[:octicons-tag-24: v2.6.13](https://github.com/torvalds/linux/commit/317a76f9a44b437d6301718f4e5d08bd93f98da7)

`#!c char name[TCP_CA_NAME_MAX]`

This field contains the name of the algorithm to be attached.

### `key`

[:octicons-tag-24: v4.0](https://github.com/torvalds/linux/commit/c5c6a8ab45ec0f18733afb4aaade0d4a139d80b3)

`#!c u32 key`

This field is read-only. It is calculated by the kernel when loading the algorithm by <nospell>jhash'ing</nospell> the algorithm name. This key can then be used to identify the algorithm when selecting an algorithm per route.

### `flags`

[:octicons-tag-24: v2.6.22](https://github.com/torvalds/linux/commit/164891aadf1721fca4dce473bb0e0998181537c6)

This field contains flags for the congestion control algorithm.

Bitmask values:

- `TCP_CONG_NON_RESTRICTED` - Algorithm can be set on socket without CAP_NET_ADMIN privileges
- `TCP_CONG_NEEDS_ECN` - Requires ECN/ECT set on all packets

### `ssthresh`

[:octicons-tag-24: v2.6.13](https://github.com/torvalds/linux/commit/317a76f9a44b437d6301718f4e5d08bd93f98da7)

`#!c u32 (*ssthresh)(struct sock *sk)`

This function/program returns a the slow start threshold for a given socket. This function/program must be set in order to be able to attach the congestion control algorithm.

**Parameters**

`sk`: The socket for which the slow start threshold is requested.

**Returns**

The amount of packets after which to disable [slow start](https://developer.mozilla.org/en-US/docs/Glossary/TCP_slow_start).

### `cong_avoid`

[:octicons-tag-24: v2.6.13](https://github.com/torvalds/linux/commit/317a76f9a44b437d6301718f4e5d08bd93f98da7)

`#!c void (*cong_avoid)(struct sock *sk, u32 ack, u32 acked)`

This function/program is called for every ACK received by the socket. It is used to adjust the congestion window based on the number of packets acknowledged.

**Parameters**

`sk`: The socket for which the congestion window should be adjusted.

`ack`: The sequence number of the ACK.

`acked`: The number of packets acknowledged by the ACK.

### `set_state`

[:octicons-tag-24: v2.6.13](https://github.com/torvalds/linux/commit/317a76f9a44b437d6301718f4e5d08bd93f98da7)

`#!c void (*set_state)(struct sock *sk, u8 new_state)`

This function/program is called when the congestion algorithm state of the socket changes.

**Parameters**

`sk`: The socket for which the state should be changed.

`new_state`: The new state of the congestion control algorithm. See [`enum tcp_ca_state`](#enum-tcp_ca_state) for possible values.

### `cwnd_event`

[:octicons-tag-24: v2.6.13](https://github.com/torvalds/linux/commit/317a76f9a44b437d6301718f4e5d08bd93f98da7)

`#!c void (*cwnd_event)(struct sock *sk, enum tcp_ca_event ev)`

This function/program is called when a congestion window event occurs.

**Parameters**

`sk`: The socket for which the congestion window event should be triggered.

`ev`: The event that triggered the congestion window event. See [`enum tcp_ca_event`](#enum-tcp_ca_event) for possible values.

### `in_ack_event`

[:octicons-tag-24: v3.18](https://github.com/torvalds/linux/commit/7354c8c389d18719dd71cc810da70b0921d66694)

`#!c void (*in_ack_event)(struct sock *sk, u32 flags)`

This function/program is called when an ACK event occurs.

**Parameters**

`sk`: The socket for which the ACK event should be triggered.

`flags`: Bitmask of flags for the ACK event. See [`enum tcp_ca_ack_event_flags`](#enum-tcp_ca_ack_event_flags) for possible values.

### `pkts_acked`

[:octicons-tag-24: v4.7](https://github.com/torvalds/linux/commit/756ee1729b2feb3a45767da29e338f70f2086ba3)

`#!c void (*pkts_acked)(struct sock *sk, const struct ack_sample *sample)`

This function/program is called once every time the receive queue is drained.

**Parameters**

`sk`: The socket for which the packets are acknowledged.

`sample`: A [`struct ack_sample`](#struct-ack_sample) from the received packets.

### `min_tso_segs`

[:octicons-tag-24: v4.7](https://github.com/torvalds/linux/commit/dcb8c9b4373a583451b1b8a3e916d33de273633d)

`#!c u32 (*min_tso_segs)(struct sock *sk)`

This function/program is called to query the number of segments we want in the skb we are transmitting.

**Parameters**

`sk`: The socket for which the number of segments is requested.

**Returns**

The number of segments we want in the skb we are transmitting.

### `cong_control`

[:octicons-tag-24: v4.9](https://github.com/torvalds/linux/commit/c0402760f565ae066621ebf8720a32fba074d538)

`#!c void (*cong_control)(struct sock *sk, u32 ack, int flag, const struct rate_sample *rs)`

!!! warning
    The signature of this function was changed in [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/57bfc7605ca5b102ba336779ae9adbc5bbba1d96).
    The before that the signature was `void (*cong_control)(struct sock *sk, const struct rate_sample *rs)`

This function/program is called for every ACK packet received to update <nospell>cwnd</nospell> and pacing rate, after all the ca_state processing.

**Parameters**

`sk`: The socket for which the congestion control should be updated.

`ack`: The sequence number of the ACK.

`flag`: Flags for the congestion control.

Bitmask values for `flag`:

```c
#define FLAG_DATA               0x01    // (1)!
#define FLAG_WIN_UPDATE         0x02    // (2)!
#define FLAG_DATA_ACKED         0x04    // (3)!
#define FLAG_RETRANS_DATA_ACKED	0x08    // (4)!
#define FLAG_SYN_ACKED          0x10    // (5)!
#define FLAG_DATA_SACKED        0x20    // (6)!
#define FLAG_ECE                0x40    // (7)!
#define FLAG_LOST_RETRANS       0x80    // (8)!
#define FLAG_SLOWPATH           0x100   // (9)!
#define FLAG_ORIG_SACK_ACKED    0x200   // (10)!
#define FLAG_SND_UNA_ADVANCED   0x400   // (11)!
#define FLAG_DSACKING_ACK       0x800   // (12)!
#define FLAG_SET_XMIT_TIMER     0x1000  // (13)!
#define FLAG_SACK_RENEGING      0x2000  // (14)!
#define FLAG_UPDATE_TS_RECENT   0x4000  // (15)!
#define FLAG_NO_CHALLENGE_ACK   0x8000  // (16)!
#define FLAG_ACK_MAYBE_DELAYED  0x10000 // (17)!
#define FLAG_DSACK_TLP          0x20000 // (18)!
```

1. Incoming frame contained data.
2. Incoming <nospell>ACK</nospell> was a window update.
3. This <nospell>ACK</nospell> acknowledged new data.
4. This <nospell>ACK</nospell> acknowledged new data some of which was retransmitted.
5. This <nospell>ACK</nospell> acknowledged <nospell>SYN</nospell>.
6. New <nospell>SACK</nospell>.
7. <nospell>ECE</nospell> in this <nospell>ACK</nospell>
8. This <nospell>ACK</nospell> marks some retransmission lost
9. Do not skip RFC checks for window update
10. Never retransmitted data are (s)acked
11. `snd_una` was changed (!= `FLAG_DATA_ACKED`)
12. SACK blocks contained <nospell>D-SACK</nospell> info
13. Set <nospell>TLP</nospell> or <nospell>RTO</nospell> timer
14. `snd_una` advanced to a sacked seq
15. `tcp_replace_ts_recent()`
16. do not call `tcp_send_challenge_ack()`
17. Likely a delayed ACK
18. [:octicons-tag-24: v5.15](https://github.com/torvalds/linux/commit/63f367d9de77b30f58722c1be9e334fb0f5f342d) <nospell>DSACK</nospell> for tail loss probe

`rs`: A [`struct rate_sample`](#struct-rate_sample) containing the rate sample.

### `undo_cwnd`

[:octicons-tag-24: v2.6.13](https://github.com/torvalds/linux/commit/317a76f9a44b437d6301718f4e5d08bd93f98da7)

`#!c u32 (*undo_cwnd)(struct sock *sk)`

This function/program is called when packet loss is detected.

**Parameters**

`sk`: The socket for which the congestion window should be undone.

**Returns**

The new size of the congestion window.

### `sndbuf_expand`

[:octicons-tag-24: v4.9](https://github.com/torvalds/linux/commit/77bfc174c38e558a3425d3b069aa2762b2fedfdd)

`#!c u32 (*sndbuf_expand)(struct sock *sk)`

This function/program is called to query the multiplier for the send buffer.

**Parameters**

`sk`: The socket for which the send buffer multiplier is requested.

**Returns**

The multiplier for the send buffer.

### `get_info`

[:octicons-tag-24: v2.6.13](https://github.com/torvalds/linux/commit/317a76f9a44b437d6301718f4e5d08bd93f98da7)

`#!c size_t (*get_info)(struct sock *sk, u32 ext, int *attr, union tcp_cc_info *info)`

This function/program is called to get information about the congestion control state.

**Parameters**

`sk`: The socket for which the congestion control information is requested.

`ext`: A bitmask of extensions for which information is requested. Values are:

```c
enum {
	INET_DIAG_NONE,
	INET_DIAG_MEMINFO,
	INET_DIAG_INFO,
	INET_DIAG_VEGASINFO,
	INET_DIAG_CONG,
	INET_DIAG_TOS,
	INET_DIAG_TCLASS,
	INET_DIAG_SKMEMINFO,
	INET_DIAG_SHUTDOWN,
	INET_DIAG_DCTCPINFO, // (1)!
	INET_DIAG_PROTOCOL,  // (2)!
	INET_DIAG_SKV6ONLY,
	INET_DIAG_LOCALS,
	INET_DIAG_PEERS,
	INET_DIAG_PAD,
	INET_DIAG_MARK,		// (3)!
	INET_DIAG_BBRINFO,	// (4)!
	INET_DIAG_CLASS_ID,	// (5)!
	INET_DIAG_MD5SIG,
	INET_DIAG_ULP_INFO,
	INET_DIAG_SK_BPF_STORAGES,
	INET_DIAG_CGROUP_ID,
	INET_DIAG_SOCKOPT,
};
```

1. request as `INET_DIAG_VEGASINFO` 
2. response attribute only
3. only with `CAP_NET_ADMIN`
4. request as `INET_DIAG_VEGASINFO`
5. request as `INET_DIAG_TCLASS`

`attr`: The attribute for the congestion control information.

`info`: The congestion control information, its type being [`union tcp_cc_info`](#union-tcp_cc_info).


### `init`

[:octicons-tag-24: v2.6.13](https://github.com/torvalds/linux/commit/317a76f9a44b437d6301718f4e5d08bd93f98da7)

`#!c void (*init)(struct sock *sk)`

This function/program is called when the congestion control algorithm is initialized.

**Parameters**

`sk`: The socket for which the congestion control algorithm should be initialized.

### `release`

[:octicons-tag-24: v2.6.13](https://github.com/torvalds/linux/commit/317a76f9a44b437d6301718f4e5d08bd93f98da7)

`#!c void (*release)(struct sock *sk)`

This function/program is called when the congestion control algorithm is released.

**Parameters**

`sk`: The socket for which the congestion control algorithm should be released.

## Types

### `struct rate_sample`

A rate sample measures the number of (original/retransmitted) data packets delivered "delivered" over an interval of time "interval_us". The `tcp_rate.c` code fills in the rate sample, and congestion control modules that define a cong_control function to run at the end of ACK processing can optionally chose to consult this sample when setting cwnd and pacing rate. A sample is invalid if "delivered" or "interval_us" is negative.

```c
struct rate_sample {
	u64  prior_mstamp;      // (1)!  
	u32  prior_delivered;   // (2)!  
	u32  prior_delivered_ce;// (3)!  
	s32  delivered;         // (4)!  
	s32  delivered_ce;      // (5)!  
	long interval_us;       // (6)!  
	u32 snd_interval_us;    // (7)!  
	u32 rcv_interval_us;    // (8)!  
	long rtt_us;            // (9)!  
	int  losses;            // (10)! 
	u32  acked_sacked;      // (11)! 
	u32  prior_in_flight;   // (12)! 
	u32  last_end_seq;      // (13)! 
	bool is_app_limited;    // (14)! 
	bool is_retrans;        // (15)! 
	bool is_ack_delayed;    // (16)! 
};
```

1. Starting timestamp for interval
2. `tp->delivered` at `prior_mstamp`
3. `tp->delivered_ce` at `prior_mstamp`
4. Number of packets delivered over interval
5. Number of packets delivered with CE mark
6. Time for `tp->delivered` to incr "delivered"
7. `snd` interval for delivered packets
8. `rcv` interval for delivered packets
9. <nospell>RTT</nospell> of last <nospell>(S)ACKed</nospell> packet (or `-1`)
10. Number of packets marked lost upon ACK
11. Number of packets newly <nospell>(S)ACKed</nospell> upon <nospell>ACK</nospell>
12. In flight before this <nospell>ACK</nospell>
13. end_seq of most recently <nospell>ACKed</nospell> packet
14. Is sample from packet with bubble in pipe?
15. Is sample from retransmission?
16. Is this (likely) a delayed <nospell>ACK</nospell>?

### `struct ack_sample`

```c
struct ack_sample {
	u32 pkts_acked;
	s32 rtt_us;
	u32 in_flight;
};
```

### `union tcp_cc_info`

The value of the union is determined based on the `INET_DIAG_*` value provided in the `ext` parameter of the `get_info` function.

```c
union tcp_cc_info {
	struct tcpvegas_info	vegas;
	struct tcp_dctcp_info	dctcp;
	struct tcp_bbr_info		bbr;
};
```

### `struct tcpvegas_info`

The values for `INET_DIAG_VEGASINFO`

```c
struct tcpvegas_info {
	__u32	tcpv_enabled;
	__u32	tcpv_rttcnt;
	__u32	tcpv_rtt;
	__u32	tcpv_minrtt;
};
```

### `struct tcp_dctcp_info`

The values for `INET_DIAG_DCTCPINFO`

```c
struct tcp_dctcp_info {
	__u16	dctcp_enabled;
	__u16	dctcp_ce_state;
	__u32	dctcp_alpha;
	__u32	dctcp_ab_ecn;
	__u32	dctcp_ab_tot;
};
```

### `struct tcp_bbr_info`

The values for `INET_DIAG_BBRINFO`

```c
struct tcp_bbr_info {
	/* u64 bw: max-filtered BW (app throughput) estimate in Byte per sec: */
	__u32	bbr_bw_lo;		/* lower 32 bits of bw */
	__u32	bbr_bw_hi;		/* upper 32 bits of bw */
	__u32	bbr_min_rtt;		/* min-filtered RTT in uSec */
	__u32	bbr_pacing_gain;	/* pacing gain shifted left 8 bits */
	__u32	bbr_cwnd_gain;		/* cwnd gain shifted left 8 bits */
};
```

### `enum tcp_ca_state`

Sender's congestion state indicating normal or abnormal situations in the last round of packets sent. The state is driven by the ACK information and timer events.

```c
enum tcp_ca_state {
	TCP_CA_Open     = 0,
	TCP_CA_Disorder = 1,
	TCP_CA_CWR      = 2,
	TCP_CA_Recovery = 3,
	TCP_CA_Loss     = 4
};
```

#### `TCP_CA_Open`

Nothing bad has been observed recently. No apparent reordering, packet loss, or <nospell>ECN</nospell> marks.

#### `TCP_CA_Disorder`

The sender enters disordered state when it has received <nospell>DUPACKs</nospell> or <nospell>SACKs</nospell> in the last round of packets sent. This could be due to packet loss or reordering but needs further information to confirm packets have been lost.

#### `TCP_CA_CWR`

The sender enters Congestion Window Reduction (<nospell>CWR</nospell>) state when it has received ACKs with <nospell>ECN-ECE</nospell> marks, or has experienced congestion or packet discard on the sender host (e.g. qdisc).

#### `TCP_CA_Recovery`

The sender is in fast recovery and retransmitting lost packets,
typically triggered by ACK events.

#### `TCP_CA_Loss`

The sender is in loss recovery triggered by retransmission timeout.

### `enum tcp_ca_event`

Events passed to congestion control interface

```c
enum tcp_ca_event {
	CA_EVENT_TX_START,      // (1)!
	CA_EVENT_CWND_RESTART,  // (2)!
	CA_EVENT_COMPLETE_CWR,  // (3)!
	CA_EVENT_LOSS,          // (4)!
	CA_EVENT_ECN_NO_CE,     // (5)!
	CA_EVENT_ECN_IS_CE,     // (6)!
};
```

1. First transmit when no packets in flight
2. Congestion window restart
3. End of congestion recovery
4. Loss timeout
5. <nospell>ECT</nospell> set, but not CE marked
6. Received CE marked IP packet

### `enum tcp_ca_ack_event_flags`

Information about inbound ACK, passed to [`in_ack_event`](#in_ack_event)

```c
enum tcp_ca_ack_event_flags {
	CA_ACK_SLOWPATH     = (1 << 0), // (1)!
	CA_ACK_WIN_UPDATE   = (1 << 1), // (1)!
	CA_ACK_ECE          = (1 << 2), // (1)!
};
```

1. In slow path processing
2. ACK updated window
3. <nospell>ECE</nospell> bit is set on ack

## Example

This example is sourced from [https://elixir.bootlin.com/linux/v6.13/source/tools/testing/selftests/bpf/progs/bpf_dctcp.c](https://elixir.bootlin.com/linux/v6.13/source/tools/testing/selftests/bpf/progs/bpf_dctcp.c)

```c
// SPDX-License-Identifier: GPL-2.0
/* Copyright (c) 2019 Facebook */

/* WARNING: This implementation is not necessarily the same
 * as the tcp_dctcp.c.  The purpose is mainly for testing
 * the kernel BPF logic.
 */

#include "bpf_tracing_net.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>

#ifndef EBUSY
#define EBUSY 16
#endif
#define min(a, b) ((a) < (b) ? (a) : (b))
#define max(a, b) ((a) > (b) ? (a) : (b))
#define min_not_zero(x, y) ({			\
	typeof(x) __x = (x);			\
	typeof(y) __y = (y);			\
	__x == 0 ? __y : ((__y == 0) ? __x : min(__x, __y)); })
static bool before(__u32 seq1, __u32 seq2)
{
	return (__s32)(seq1-seq2) < 0;
}

char _license[] SEC("license") = "GPL";

volatile const char fallback_cc[TCP_CA_NAME_MAX];
const char bpf_dctcp[] = "bpf_dctcp";
const char tcp_cdg[] = "cdg";
char cc_res[TCP_CA_NAME_MAX];
int tcp_cdg_res = 0;
int stg_result = 0;
int ebusy_cnt = 0;

struct {
	__uint(type, BPF_MAP_TYPE_SK_STORAGE);
	__uint(map_flags, BPF_F_NO_PREALLOC);
	__type(key, int);
	__type(value, int);
} sk_stg_map SEC(".maps");

#define DCTCP_MAX_ALPHA	1024U

struct bpf_dctcp {
	__u32 old_delivered;
	__u32 old_delivered_ce;
	__u32 prior_rcv_nxt;
	__u32 dctcp_alpha;
	__u32 next_seq;
	__u32 ce_state;
	__u32 loss_cwnd;
};

static unsigned int dctcp_shift_g = 4; /* g = 1/2^4 */
static unsigned int dctcp_alpha_on_init = DCTCP_MAX_ALPHA;

static void dctcp_reset(const struct tcp_sock *tp, struct bpf_dctcp *ca)
{
	ca->next_seq = tp->snd_nxt;

	ca->old_delivered = tp->delivered;
	ca->old_delivered_ce = tp->delivered_ce;
}

SEC("struct_ops")
void BPF_PROG(bpf_dctcp_init, struct sock *sk)
{
	const struct tcp_sock *tp = tcp_sk(sk);
	struct bpf_dctcp *ca = inet_csk_ca(sk);
	int *stg;

	if (!(tp->ecn_flags & TCP_ECN_OK) && fallback_cc[0]) {
		/* Switch to fallback */
		if (bpf_setsockopt(sk, SOL_TCP, TCP_CONGESTION,
				   (void *)fallback_cc, sizeof(fallback_cc)) == -EBUSY)
			ebusy_cnt++;

		/* Switch back to myself and the recurred bpf_dctcp_init()
		 * will get -EBUSY for all bpf_setsockopt(TCP_CONGESTION),
		 * except the last "cdg" one.
		 */
		if (bpf_setsockopt(sk, SOL_TCP, TCP_CONGESTION,
				   (void *)bpf_dctcp, sizeof(bpf_dctcp)) == -EBUSY)
			ebusy_cnt++;

		/* Switch back to fallback */
		if (bpf_setsockopt(sk, SOL_TCP, TCP_CONGESTION,
				   (void *)fallback_cc, sizeof(fallback_cc)) == -EBUSY)
			ebusy_cnt++;

		/* Expecting -ENOTSUPP for tcp_cdg_res */
		tcp_cdg_res = bpf_setsockopt(sk, SOL_TCP, TCP_CONGESTION,
					     (void *)tcp_cdg, sizeof(tcp_cdg));
		bpf_getsockopt(sk, SOL_TCP, TCP_CONGESTION,
			       (void *)cc_res, sizeof(cc_res));
		return;
	}

	ca->prior_rcv_nxt = tp->rcv_nxt;
	ca->dctcp_alpha = min(dctcp_alpha_on_init, DCTCP_MAX_ALPHA);
	ca->loss_cwnd = 0;
	ca->ce_state = 0;

	stg = bpf_sk_storage_get(&sk_stg_map, (void *)tp, NULL, 0);
	if (stg) {
		stg_result = *stg;
		bpf_sk_storage_delete(&sk_stg_map, (void *)tp);
	}
	dctcp_reset(tp, ca);
}

SEC("struct_ops")
__u32 BPF_PROG(bpf_dctcp_ssthresh, struct sock *sk)
{
	struct bpf_dctcp *ca = inet_csk_ca(sk);
	struct tcp_sock *tp = tcp_sk(sk);

	ca->loss_cwnd = tp->snd_cwnd;
	return max(tp->snd_cwnd - ((tp->snd_cwnd * ca->dctcp_alpha) >> 11U), 2U);
}

SEC("struct_ops")
void BPF_PROG(bpf_dctcp_update_alpha, struct sock *sk, __u32 flags)
{
	const struct tcp_sock *tp = tcp_sk(sk);
	struct bpf_dctcp *ca = inet_csk_ca(sk);

	/* Expired RTT */
	if (!before(tp->snd_una, ca->next_seq)) {
		__u32 delivered_ce = tp->delivered_ce - ca->old_delivered_ce;
		__u32 alpha = ca->dctcp_alpha;

		/* alpha = (1 - g) * alpha + g * F */

		alpha -= min_not_zero(alpha, alpha >> dctcp_shift_g);
		if (delivered_ce) {
			__u32 delivered = tp->delivered - ca->old_delivered;

			/* If dctcp_shift_g == 1, a 32bit value would overflow
			 * after 8 M packets.
			 */
			delivered_ce <<= (10 - dctcp_shift_g);
			delivered_ce /= max(1U, delivered);

			alpha = min(alpha + delivered_ce, DCTCP_MAX_ALPHA);
		}
		ca->dctcp_alpha = alpha;
		dctcp_reset(tp, ca);
	}
}

static void dctcp_react_to_loss(struct sock *sk)
{
	struct bpf_dctcp *ca = inet_csk_ca(sk);
	struct tcp_sock *tp = tcp_sk(sk);

	ca->loss_cwnd = tp->snd_cwnd;
	tp->snd_ssthresh = max(tp->snd_cwnd >> 1U, 2U);
}

SEC("struct_ops")
void BPF_PROG(bpf_dctcp_state, struct sock *sk, __u8 new_state)
{
	if (new_state == TCP_CA_Recovery &&
	    new_state != BPF_CORE_READ_BITFIELD(inet_csk(sk), icsk_ca_state))
		dctcp_react_to_loss(sk);
	/* We handle RTO in bpf_dctcp_cwnd_event to ensure that we perform only
	 * one loss-adjustment per RTT.
	 */
}

static void dctcp_ece_ack_cwr(struct sock *sk, __u32 ce_state)
{
	struct tcp_sock *tp = tcp_sk(sk);

	if (ce_state == 1)
		tp->ecn_flags |= TCP_ECN_DEMAND_CWR;
	else
		tp->ecn_flags &= ~TCP_ECN_DEMAND_CWR;
}

/* Minimal DCTP CE state machine:
 *
 * S:	0 <- last pkt was non-CE
 *	1 <- last pkt was CE
 */
static void dctcp_ece_ack_update(struct sock *sk, enum tcp_ca_event evt,
				 __u32 *prior_rcv_nxt, __u32 *ce_state)
{
	__u32 new_ce_state = (evt == CA_EVENT_ECN_IS_CE) ? 1 : 0;

	if (*ce_state != new_ce_state) {
		/* CE state has changed, force an immediate ACK to
		 * reflect the new CE state. If an ACK was delayed,
		 * send that first to reflect the prior CE state.
		 */
		if (inet_csk(sk)->icsk_ack.pending & ICSK_ACK_TIMER) {
			dctcp_ece_ack_cwr(sk, *ce_state);
			bpf_tcp_send_ack(sk, *prior_rcv_nxt);
		}
		inet_csk(sk)->icsk_ack.pending |= ICSK_ACK_NOW;
	}
	*prior_rcv_nxt = tcp_sk(sk)->rcv_nxt;
	*ce_state = new_ce_state;
	dctcp_ece_ack_cwr(sk, new_ce_state);
}

SEC("struct_ops")
void BPF_PROG(bpf_dctcp_cwnd_event, struct sock *sk, enum tcp_ca_event ev)
{
	struct bpf_dctcp *ca = inet_csk_ca(sk);

	switch (ev) {
	case CA_EVENT_ECN_IS_CE:
	case CA_EVENT_ECN_NO_CE:
		dctcp_ece_ack_update(sk, ev, &ca->prior_rcv_nxt, &ca->ce_state);
		break;
	case CA_EVENT_LOSS:
		dctcp_react_to_loss(sk);
		break;
	default:
		/* Don't care for the rest. */
		break;
	}
}

SEC("struct_ops")
__u32 BPF_PROG(bpf_dctcp_cwnd_undo, struct sock *sk)
{
	const struct bpf_dctcp *ca = inet_csk_ca(sk);

	return max(tcp_sk(sk)->snd_cwnd, ca->loss_cwnd);
}

extern void tcp_reno_cong_avoid(struct sock *sk, __u32 ack, __u32 acked) __ksym;

SEC("struct_ops")
void BPF_PROG(bpf_dctcp_cong_avoid, struct sock *sk, __u32 ack, __u32 acked)
{
	tcp_reno_cong_avoid(sk, ack, acked);
}

SEC(".struct_ops")
struct tcp_congestion_ops dctcp_nouse = {
	.init		= (void *)bpf_dctcp_init,
	.set_state	= (void *)bpf_dctcp_state,
	.flags		= TCP_CONG_NEEDS_ECN,
	.name		= "bpf_dctcp_nouse",
};

SEC(".struct_ops")
struct tcp_congestion_ops dctcp = {
	.init		= (void *)bpf_dctcp_init,
	.in_ack_event   = (void *)bpf_dctcp_update_alpha,
	.cwnd_event	= (void *)bpf_dctcp_cwnd_event,
	.ssthresh	= (void *)bpf_dctcp_ssthresh,
	.cong_avoid	= (void *)bpf_dctcp_cong_avoid,
	.undo_cwnd	= (void *)bpf_dctcp_cwnd_undo,
	.set_state	= (void *)bpf_dctcp_state,
	.flags		= TCP_CONG_NEEDS_ECN,
	.name		= "bpf_dctcp",
};
```
