---
title: "KFunc 'bpf_sk_assign_tcp_reqsk'"
description: "This page documents the 'bpf_sk_assign_tcp_reqsk' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_sk_assign_tcp_reqsk`

<!-- [FEATURE_TAG](bpf_sk_assign_tcp_reqsk) -->
[:octicons-tag-24: v6.11](https://github.com/torvalds/linux/commit/e472f88891abbc535a5e16a68a104073985f6061)
<!-- [/FEATURE_TAG] -->

Assign custom parameters used to validate SYN cookies.

## Definition

`bpf_sk_assign_tcp_reqsk()` takes `skb`, a listener `sk`, and `struct bpf_tcp_req_attrs` and allocates `reqsk` and configures it. Then, bpf_sk_assign_tcp_reqsk() links `reqsk` with `skb` and the listener.

<!-- [KFUNC_DEF] -->
`#!c int bpf_sk_assign_tcp_reqsk(struct __sk_buff *s, struct sock *sk, struct bpf_tcp_req_attrs *attrs, int attrs__sz)`
<!-- [/KFUNC_DEF] -->

## Usage

Under SYN Flood, the TCP stack generates SYN Cookie to remain stateless for the connection request until a valid ACK is responded to the SYN+ACK.

The cookie contains two kinds of host-specific bits, a timestamp and secrets, so only can it be validated by the generator. It means SYN Cookie consumes network resources between the client and the server; intermediate nodes must remember which nodes to route ACK for the cookie.

SYN Proxy reduces such unwanted resource allocation by handling <nospell>3WHS</nospell> at the edge network.  After SYN Proxy completes <nospell>3WHS</nospell>, it forwards SYN to the backend server and completes another <nospell>3WHS</nospell>.  However, since the server's <nospell>ISN</nospell> differs from the cookie, the proxy must manage the <nospell>ISN</nospell> mappings and fix up SEQ/ACK numbers in every packet for each connection.  If a proxy node goes down, all the connections through it are terminated.  Keeping a state at proxy is painful from that perspective.

This kfunc allows BPF to validate an arbitrary SYN Cookie on the backend server, the proxy doesn't need not restore SYN nor pass it.  After validating ACK, the proxy node just needs to forward it, and then the server can do the lightweight validation (e.g. check if ACK came from proxy nodes, etc)
and create a connection from the ACK.

The arguments supplied to the kfunc can be derived from the SYN Cookie.

See also: [patch set](https://lore.kernel.org/all/20240115205514.68364-1-kuniyu@amazon.com/)

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

```c
struct bpf_tcp_req_attrs attrs = {
    .mss = mss,
    .wscale_ok = wscale_ok,
    .rcv_wscale = rcv_wscale, /* Server's WScale < 15 */
    .snd_wscale = snd_wscale, /* Client's WScale < 15 */
    .tstamp_ok = tstamp_ok,
    .rcv_tsval = tsval,
    .rcv_tsecr = tsecr, /* Server's Initial TSval */
    .usec_ts_ok = usec_ts_ok,
    .sack_ok = sack_ok,
    .ecn_ok = ecn_ok,
}

skc = bpf_skc_lookup_tcp(...);
sk = (struct sock *)bpf_skc_to_tcp_sock(skc);
bpf_sk_assign_tcp_reqsk(skb, sk, attrs, sizeof(attrs));
bpf_sk_release(skc);
```
