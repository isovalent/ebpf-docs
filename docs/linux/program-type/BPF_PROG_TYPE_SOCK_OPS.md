---
title: "Program Type 'BPF_PROG_TYPE_SOCK_OPS'"
description: "This page documents the 'BPF_PROG_TYPE_SOCK_OPS' eBPF program type, including its definition, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_SOCK_OPS`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_SOCK_OPS) -->
[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/40304b2a1567fecc321f640ee4239556dd0f3ee0)
<!-- [/FEATURE_TAG] -->

Socket ops programs are attached to cGroups and get called for multiple lifecycle events of a socket, giving the program the opportunity to changes settings per connection or to record the existence of a socket.

## Usage

Socket ops programs are called multiple times on the same socket during different parts of its lifecycle for different operations. Some operations query the program for certain parameters, others just inform the program of certain events so the program can perform some at that time.

Regardless of the type of operation, the program should always return `1` on success. A negative integer indicate a operation is not supported. For operations that query information, the [`reply`](#reply) field in the context is used to "reply" to the query, the program is expected to set it equal to the requested value.

There are a few envisioned use cases for this program type. First is to reply with certain settings like RTO, RTT and ECN (see [ops](#ops) section for details) or to set socket options using the [`bpf_setsockopt`](../helper-function/bpf_setsockopt.md) helper to tune settings/options on a per-connection basis. 

>  For example, it is easy to use Facebook's internal IPv6 addresses to determine if both hosts of a connection are in the same data center. Therefore, it is easy to write a BPF program to choose a small SYN RTO value when both hosts are in the same data center.

Secondly, socket ops programs are in an excellent position to gather detailed metrics about connections. Especially after [:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6).

Thirdly, socket ops programs can be used to implement TCP options which are not known to the kernel, both on the sending and receiving side. See [`BPF_SOCK_OPS_PARSE_HDR_OPT_CB`](#bpf_sock_ops_parse_hdr_opt_cb) and [`BPF_SOCK_OPS_WRITE_HDR_OPT_CB`](#bpf_sock_ops_write_hdr_opt_cb).

The last, but not least, envisioned use case for socket ops programs is to dynamically add sockets to [`BPF_MAP_TYPE_SOCKMAP`](../map-type/BPF_MAP_TYPE_SOCKMAP.md) or [`BPF_MAP_TYPE_SOCKHASH`](../map-type/BPF_MAP_TYPE_SOCKHASH.md) maps. Since socket ops programs are notified when sockets are connecting or listening, it allows us to add the sockets to these maps before any actual message traffic happens. This allows [`BPF_PROG_TYPE_SK_MSG`](../program-type/BPF_PROG_TYPE_SK_MSG.md) and [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md) to operate without user space needing to add sockets to the sock maps. The [`bpf_sock_map_update`](../helper-function/bpf_sock_map_update.md) and [`bpf_sock_hash_update`](../helper-function/bpf_sock_hash_update.md) helpers exist for this very purpose.

## Ops

After attaching the program, it will be invoked for multiple socket and multiple ops. The `op` field in the context indicates for which operation the program is invoked. Availability of fields in the context and the meaning of return values vary from op to op.

The ops ending with `_CB` are callbacks which are just called to notify the program of an event. Return values for these ops are ignored. Some of these callbacks are not triggered unless activated by setting flags on the socket. Setting these flags is done by the program itself with the use of the [`bpf_sock_ops_cb_flags_set`](../helper-function/bpf_sock_ops_cb_flags_set.md) helper which can both set and unset flags.

### `BPF_SOCK_OPS_TIMEOUT_INIT`
<!-- [FEATURE_TAG](BPF_SOCK_OPS_TIMEOUT_INIT) -->
[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/8550f328f45db6d37981eb2041bc465810245c03)
<!-- [/FEATURE_TAG] -->

When invoked with this `op`, the program can overwrite the default RTO (retransmission timeout) for a SYN or SYN-ACK. `-1` can be returned if default value should be used.

### `BPF_SOCK_OPS_RWND_INIT`
<!-- [FEATURE_TAG](BPF_SOCK_OPS_RWND_INIT) -->
[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/13d3b1ebe28762c79e981931a41914fae5d04386)
<!-- [/FEATURE_TAG] -->

When invoked with this `op`, the program can overwrite the default initial advertized window (in packets) or -1 if default value should be used.

### `BPF_SOCK_OPS_TCP_CONNECT_CB`
<!-- [FEATURE_TAG](BPF_SOCK_OPS_TCP_CONNECT_CB) -->
[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/9872a4bde31b0b055448e9ac1f4c9ee62d978766)
<!-- [/FEATURE_TAG] -->	

The program is invoked with this `op` when a socket is in the 'connect' state, it has sent out a SYN message, but is not yet established. This is just a notification, return value is discarded.

### `BPF_SOCK_OPS_ACTIVE_ESTABLISHED_CB`
<!-- [FEATURE_TAG](BPF_SOCK_OPS_ACTIVE_ESTABLISHED_CB) -->
[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/9872a4bde31b0b055448e9ac1f4c9ee62d978766)
<!-- [/FEATURE_TAG] -->


The program is invoked with this `op` when a active socket transitioned to have an established connection. This happens when a outgoing connection establishes. This is just a notification, return value is discarded.

### `BPF_SOCK_OPS_PASSIVE_ESTABLISHED_CB`
<!-- [FEATURE_TAG](BPF_SOCK_OPS_PASSIVE_ESTABLISHED_CB) -->
[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/9872a4bde31b0b055448e9ac1f4c9ee62d978766)
<!-- [/FEATURE_TAG] -->

The program is invoked with this `op` when a active socket transitioned to have an established connection. This happens when a incoming connection establishes. This is just a notification, return value is discarded.

### `BPF_SOCK_OPS_NEEDS_ECN`
<!-- [FEATURE_TAG](BPF_SOCK_OPS_NEEDS_ECN) -->
[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/91b5b21c7c16899abb37f4a9e4388b4e9aae0b9d)
<!-- [/FEATURE_TAG] -->

When invoked with this `op`, the program is asked if [ECN](https://en.wikipedia.org/wiki/Explicit_Congestion_Notification) (Explicit Congestion Notification) should be enabled for a given connection. The program is expected to return `0` or `1`. 

### `BPF_SOCK_OPS_BASE_RTT`
<!-- [FEATURE_TAG](BPF_SOCK_OPS_BASE_RTT) -->
[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/e6546ef6d86d0fc38e0e84ccae80e641f3fc0087)
<!-- [/FEATURE_TAG] -->

When invoked with this `op`, the program is asked for the base RTT (Round Trip Time) for a given connection. If the measured RTT goes above this value it indicates the connection is congested and the congestion control algorithm will take steps.

### `BPF_SOCK_OPS_RTO_CB`
<!-- [FEATURE_TAG](BPF_SOCK_OPS_RTO_CB) -->
[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/f89013f66d0f1a0dad44c513318efb706399a36b)
<!-- [/FEATURE_TAG] -->

When `BPF_SOCK_OPS_RTO_CB_FLAG` is set via [`bpf_sock_ops_cb_flags_set`](../helper-function/bpf_sock_ops_cb_flags_set.md), this program may be called with this `op` to indicate when an RTO (Retransmission Timeout) has triggered. This is just a notification, return value is discarded.

The arguments in the context will have the following meanings:

* `args[0]`: value of `icsk_retransmits`
* `args[1]`: value of `icsk_rto`
* `args[2]`: whether RTO has expired

### `BPF_SOCK_OPS_RETRANS_CB`
<!-- [FEATURE_TAG](BPF_SOCK_OPS_RETRANS_CB) -->
[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/a31ad29e6a30cb0b9084a9425b819cdcd97273ce)
<!-- [/FEATURE_TAG] -->

When the `BPF_SOCK_OPS_RETRANS_CB_FLAG` flag is set with [`bpf_sock_ops_cb_flags_set`](../helper-function/bpf_sock_ops_cb_flags_set.md), the program is invoked with this `op` when a packet from the skb has been retransmitted. This is just a notification, return value is discarded.

The arguments in the context will have the following meanings:

* `args[0]`: sequence number of 1st byte
* `args[1]`: # segments
* `args[2]`: return value of tcp_transmit_skb (0 => success)

### `BPF_SOCK_OPS_STATE_CB`
<!-- [FEATURE_TAG](BPF_SOCK_OPS_STATE_CB) -->
[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/d44874910a26f3a8f81edf873a2473363f07f660)
<!-- [/FEATURE_TAG] -->

When the `BPF_SOCK_OPS_STATE_CB_FLAG` flag is set with [`bpf_sock_ops_cb_flags_set`](../helper-function/bpf_sock_ops_cb_flags_set.md), the program is invoked with this `op` when the TCP state of the socket changes. This is just a notification, return value is discarded.

The arguments in the context will have the following meanings:

* `args[0]`: old_state
* `args[1]`: new_state

The states will be one of:
```
enum {
	BPF_TCP_ESTABLISHED = 1,
	BPF_TCP_SYN_SENT,
	BPF_TCP_SYN_RECV,
	BPF_TCP_FIN_WAIT1,
	BPF_TCP_FIN_WAIT2,
	BPF_TCP_TIME_WAIT,
	BPF_TCP_CLOSE,
	BPF_TCP_CLOSE_WAIT,
	BPF_TCP_LAST_ACK,
	BPF_TCP_LISTEN,
	BPF_TCP_CLOSING,	/* Now a valid state */
	BPF_TCP_NEW_SYN_RECV
};
```

### `BPF_SOCK_OPS_TCP_LISTEN_CB`
<!-- [FEATURE_TAG](BPF_SOCK_OPS_TCP_LISTEN_CB) -->
[:octicons-tag-24: v4.19](https://github.com/torvalds/linux/commit/f333ee0cdb27ba201e6cc0c99c76b1364aa29b86)
<!-- [/FEATURE_TAG] -->

The program is invoked with this `op` when the `listen` syscall is used on the socket, transitioning it to the `LISTEN` state. This is just a notification, return value is discarded.

### `BPF_SOCK_OPS_RTT_CB`
<!-- [FEATURE_TAG](BPF_SOCK_OPS_RTT_CB) -->
[:octicons-tag-24: v5.3](https://github.com/torvalds/linux/commit/23729ff23186424e54b4d6678fcd526cdacef4d3)
<!-- [/FEATURE_TAG] -->

When the `BPF_SOCK_OPS_RTT_CB_FLAG` flag is set with [`bpf_sock_ops_cb_flags_set`](../helper-function/bpf_sock_ops_cb_flags_set.md), the program is invoked with this `op` for every round trip. This is just a notification, return value is discarded.

### `BPF_SOCK_OPS_PARSE_HDR_OPT_CB`
<!-- [FEATURE_TAG](BPF_SOCK_OPS_PARSE_HDR_OPT_CB) -->
[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/0813a841566f0962a5551be7749b43c45f0022a0)
<!-- [/FEATURE_TAG] -->

The program is invoked with this `op` to parse TCP headers. If the `BPF_SOCK_OPS_PARSE_ALL_HDR_OPT_CB_FLAG` is set, the program will be invoked for all TCP headers, if `BPF_SOCK_OPS_PARSE_UNKNOWN_HDR_OPT_CB_FLAG` is set, the program is only invoked for unknown TCP headers.

The program will be invoked to handle the packets received at an already established connection.

The TCP header is question starts at `sock_ops->skb_data`, the [`bpf_load_hdr_opt`](../helper-function/bpf_load_hdr_opt.md) helper can also be used to search for a particular option.

This is just a notification, return value is discarded.

### `BPF_SOCK_OPS_HDR_OPT_LEN_CB`
<!-- [FEATURE_TAG](BPF_SOCK_OPS_HDR_OPT_LEN_CB) -->
[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/0813a841566f0962a5551be7749b43c45f0022a0)
<!-- [/FEATURE_TAG] -->

When the `BPF_SOCK_OPS_WRITE_HDR_OPT_CB_FLAG` flag is set with [`bpf_sock_ops_cb_flags_set`](../helper-function/bpf_sock_ops_cb_flags_set.md), the program is invoked with this `op` to reserve space for TCP options which will be written to the packet when the program is invoked with the `BPF_SOCK_OPS_WRITE_HDR_OPT_CB` op.

The arguments in the context will have the following meanings:

* `args[0]`: bool want_cookie. (in writing SYNACK only)

`sock_ops->skb_data`: Not available because no header has been written yet.

`sock_ops->skb_tcp_flags`: The tcp_flags of the outgoing skb. (e.g. SYN, ACK, FIN).

The [`bpf_reserve_hdr_opt`](../helper-function/bpf_reserve_hdr_opt.md) should be used to reserve space.

This is just a notification, return value is discarded.

### `BPF_SOCK_OPS_WRITE_HDR_OPT_CB`
<!-- [FEATURE_TAG](BPF_SOCK_OPS_WRITE_HDR_OPT_CB) -->
[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/0813a841566f0962a5551be7749b43c45f0022a0)
<!-- [/FEATURE_TAG] -->

When the `BPF_SOCK_OPS_WRITE_HDR_OPT_CB_FLAG` flag is set with [`bpf_sock_ops_cb_flags_set`](../helper-function/bpf_sock_ops_cb_flags_set.md), the program is invoked with this `op` to write TCP options to the packet, the room for these options has been reserved in a previous invokation of the program with the `BPF_SOCK_OPS_HDR_OPT_LEN_CB` op.

The arguments in the context will have the following meanings:

`args[0]`: bool want_cookie. (in writing SYNACK only)

`sock_ops->skb_data`: Referring to the outgoing skb. It covers the TCP header that has already been written by the kernel and the earlier BPF programs.

`sock_ops->skb_tcp_flags`: The tcp_flags of the outgoing skb. (e.g. SYN, ACK, FIN).

The [`bpf_store_hdr_opt`](../helper-function/bpf_store_hdr_opt.md) should be used to write the option.

The [`bpf_load_hdr_opt`](../helper-function/bpf_load_hdr_opt.md) can also be used to search for a particular option that has already been written by the kernel or the earlier BPF programs.

### `BPF_SOCK_OPS_TSTAMP_SCHED_CB`

<!-- [FEATURE_TAG](BPF_SOCK_OPS_TSTAMP_SCHED_CB) -->
[:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/6b98ec7e882af1c3088a88757e2226d06c8514f9)
<!-- [/FEATURE_TAG] -->

Called when skb is passing through device layer when `SK_BPF_CB_TX_TIMESTAMPING` feature is on. Which is done by setting a socket option `bpf_setsockopt(SK_BPF_CB_FLAGS, SK_BPF_CB_TX_TIMESTAMPING)` or calling [`bpf_sock_ops_enable_tx_tstamp`](../kfuncs/bpf_sock_ops_enable_tx_tstamp.md) on the socket.

!!! warning
    This sock op is called without taking a socket lock and will therefor not be able to use the following helper functions: [`bpf_setsockopt`](../helper-function/bpf_setsockopt.md), [`bpf_getsockopt`](../helper-function/bpf_getsockopt.md), [`bpf_sock_ops_cb_flags_set`](../helper-function/bpf_sock_ops_cb_flags_set.md), [`bpf_load_hdr_opt`](../helper-function/bpf_load_hdr_opt.md). They will always return `-EOPNOTSUPP` instead of failing verification.

### `BPF_SOCK_OPS_TSTAMP_SND_SW_CB`

<!-- [FEATURE_TAG](BPF_SOCK_OPS_TSTAMP_SND_SW_CB) -->
[:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/ecebb17ad818bc043e558c278a6c56d5bbaebacc)
<!-- [/FEATURE_TAG] -->

Called when skb is about to send to the NIC when `SK_BPF_CB_TX_TIMESTAMPING` feature is on. Which is done by setting a socket option `bpf_setsockopt(SK_BPF_CB_FLAGS, SK_BPF_CB_TX_TIMESTAMPING)` or calling [`bpf_sock_ops_enable_tx_tstamp`](../kfuncs/bpf_sock_ops_enable_tx_tstamp.md) on the socket.

!!! warning
    This sock op is called without taking a socket lock and will therefor not be able to use the following helper functions: [`bpf_setsockopt`](../helper-function/bpf_setsockopt.md), [`bpf_getsockopt`](../helper-function/bpf_getsockopt.md), [`bpf_sock_ops_cb_flags_set`](../helper-function/bpf_sock_ops_cb_flags_set.md), [`bpf_load_hdr_opt`](../helper-function/bpf_load_hdr_opt.md). They will always return `-EOPNOTSUPP` instead of failing verification.


### `BPF_SOCK_OPS_TSTAMP_SND_HW_CB`

<!-- [FEATURE_TAG](BPF_SOCK_OPS_TSTAMP_SND_HW_CB) -->
[:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/2deaf7f42b8c551e84da20483ca2d4a65c3623b3)
<!-- [/FEATURE_TAG] -->

Called in hardware phase when `SK_BPF_CB_TX_TIMESTAMPING` feature is on. Which is done by setting a socket option `bpf_setsockopt(SK_BPF_CB_FLAGS, SK_BPF_CB_TX_TIMESTAMPING)` or calling [`bpf_sock_ops_enable_tx_tstamp`](../kfuncs/bpf_sock_ops_enable_tx_tstamp.md) on the socket.

!!! warning
    This sock op is called without taking a socket lock and will therefor not be able to use the following helper functions: [`bpf_setsockopt`](../helper-function/bpf_setsockopt.md), [`bpf_getsockopt`](../helper-function/bpf_getsockopt.md), [`bpf_sock_ops_cb_flags_set`](../helper-function/bpf_sock_ops_cb_flags_set.md), [`bpf_load_hdr_opt`](../helper-function/bpf_load_hdr_opt.md). They will always return `-EOPNOTSUPP` instead of failing verification.


### `BPF_SOCK_OPS_TSTAMP_ACK_CB`

<!-- [FEATURE_TAG](BPF_SOCK_OPS_TSTAMP_ACK_CB) -->
[:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/b3b81e6b009dd8f85cd3b9c65eb492249c2649a8)
<!-- [/FEATURE_TAG] -->

Called when all the <nospell>SKBs</nospell> in the same `sendmsg` call are acked when [`SK_BPF_CB_TX_TIMESTAMPING`](../kfuncs/bpf_sock_ops_enable_tx_tstamp.md) feature is on. Which is done by setting a socket option `bpf_setsockopt(SK_BPF_CB_FLAGS, SK_BPF_CB_TX_TIMESTAMPING)` or calling [`bpf_sock_ops_enable_tx_tstamp`](../kfuncs/bpf_sock_ops_enable_tx_tstamp.md) on the socket.

!!! warning
    This sock op is called without taking a socket lock and will therefor not be able to use the following helper functions: [`bpf_setsockopt`](../helper-function/bpf_setsockopt.md), [`bpf_getsockopt`](../helper-function/bpf_getsockopt.md), [`bpf_sock_ops_cb_flags_set`](../helper-function/bpf_sock_ops_cb_flags_set.md), [`bpf_load_hdr_opt`](../helper-function/bpf_load_hdr_opt.md). They will always return `-EOPNOTSUPP` instead of failing verification.


### `BPF_SOCK_OPS_TSTAMP_SENDMSG_CB`

<!-- [FEATURE_TAG](BPF_SOCK_OPS_TSTAMP_SENDMSG_CB) -->
[:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/c9525d240c8117de35171ae705058ddf9667be27)
<!-- [/FEATURE_TAG] -->

Called when every `sendmsg` syscall is triggered. It's used to correlate `sendmsg` timestamp with corresponding `tskey`.

!!! warning
    This sock op is called without taking a socket lock and will therefor not be able to use the following helper functions: [`bpf_setsockopt`](../helper-function/bpf_setsockopt.md), [`bpf_getsockopt`](../helper-function/bpf_getsockopt.md), [`bpf_sock_ops_cb_flags_set`](../helper-function/bpf_sock_ops_cb_flags_set.md), [`bpf_load_hdr_opt`](../helper-function/bpf_load_hdr_opt.md). They will always return `-EOPNOTSUPP` instead of failing verification.

## Context

??? abstract "C structure"
    ```c
    /* User bpf_sock_ops struct to access socket values and specify request ops
    * and their replies.
    * Some of this fields are in network (bigendian) byte order and may need
    * to be converted before use (bpf_ntohl() defined in samples/bpf/bpf_endian.h).
    * New fields can only be added at the end of this structure
    */
    struct bpf_sock_ops {
        __u32 [op](#op);
        union {
            __u32 [args](#args)[4];		/* Optionally passed to bpf program */
            __u32 [reply](#reply);		/* Returned by bpf program	    */
            __u32 [replylong](#replylong)[4];	/* Optionally returned by bpf prog  */
        };
        __u32 [family](#family);
        __u32 [remote_ip4](#remote_ip4);	/* Stored in network byte order */
        __u32 [local_ip4](#local_ip4);	/* Stored in network byte order */
        __u32 [remote_ip6](#remote_ip6)[4];	/* Stored in network byte order */
        __u32 [local_ip6](#local_ip6)[4];	/* Stored in network byte order */
        __u32 [remote_port](#remote_port);	/* Stored in network byte order */
        __u32 [local_port](#local_port);	/* stored in host byte order */
        __u32 [is_fullsock](#is_fullsock);	/* Some TCP fields are only valid if
                    * there is a full socket. If not, the
                    * fields read as zero.
                    */
        __u32 [snd_cwnd](#snd_cwnd);
        __u32 [srtt_us](#srtt_us);		/* Averaged RTT << 3 in usecs */
        __u32 [bpf_sock_ops_cb_flags](#bpf_sock_ops_cb_flags); /* flags defined in uapi/linux/tcp.h */
        __u32 [state](#state);
        __u32 [rtt_min](#rtt_min);
        __u32 [snd_ssthresh](#snd_ssthresh);
        __u32 [rcv_nxt](#rcv_nxt);
        __u32 [snd_nxt](#snd_nxt);
        __u32 [snd_una](#snd_una);
        __u32 [mss_cache](#mss_cache);
        __u32 [ecn_flags](#ecn_flags);
        __u32 [rate_delivered](#rate_delivered);
        __u32 [rate_interval_us](#rate_interval_us);
        __u32 [packets_out](#packets_out);
        __u32 [retrans_out](#retrans_out);
        __u32 [total_retrans](#total_retrans);
        __u32 [segs_in](#segs_in);
        __u32 [data_segs_in](#data_segs_in);
        __u32 [segs_out](#segs_out);
        __u32 [data_segs_out](#data_segs_out);
        __u32 [lost_out](#lost_out);
        __u32 [sacked_out](#sacked_out);
        __u32 [sk_txhash](#sk_txhash);
        __u64 [bytes_received](#bytes_received);
        __u64 [bytes_acked](#bytes_acked);
        __bpf_md_ptr(struct bpf_sock *, [sk](#sk));
        /* [skb_data, skb_data_end) covers the whole TCP header.
        *
        * BPF_SOCK_OPS_PARSE_HDR_OPT_CB: The packet received
        * BPF_SOCK_OPS_HDR_OPT_LEN_CB:   Not useful because the
        *                                header has not been written.
        * BPF_SOCK_OPS_WRITE_HDR_OPT_CB: The header and options have
        *				  been written so far.
        * BPF_SOCK_OPS_ACTIVE_ESTABLISHED_CB:  The SYNACK that concludes
        *					the 3WHS.
        * BPF_SOCK_OPS_PASSIVE_ESTABLISHED_CB: The ACK that concludes
        *					the 3WHS.
        *
        * bpf_load_hdr_opt() can also be used to read a particular option.
        */
        __bpf_md_ptr(void *, [skb_data](#skb_data));
        __bpf_md_ptr(void *, [skb_data_end](#skb_data_end));
        __u32 [skb_len](#skb_len);		/* The total length of a packet.
                    * It includes the header, options,
                    * and payload.
                    */
        __u32 [skb_tcp_flags](#skb_tcp_flags);	/* tcp_flags of the header.  It provides
                    * an easy way to check for tcp_flags
                    * without parsing skb_data.
                    *
                    * In particular, the skb_tcp_flags
                    * will still be available in
                    * BPF_SOCK_OPS_HDR_OPT_LEN even though
                    * the outgoing header has not
                    * been written yet.
                    */
        __u64 [skb_hwtstamp](#skb_hwtstamp);
    };
    ```

### `op`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/40304b2a1567fecc321f640ee4239556dd0f3ee0)

This field will indicate the current operation, see the [ops section](#ops) for the possible values and meanings.

### `args`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/de525be2ca2734865d29c4b67ddd29913b214906)

This field is an array of 4 `__u32` values, used by some operations to provide additional information. The meaning of the arguments is dependant on the `op`.

### `reply`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/40304b2a1567fecc321f640ee4239556dd0f3ee0)

This field is used as the return value for operations that expect one. It is the only field the BPF program is allowed to modify.

### `replylong`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/40304b2a1567fecc321f640ee4239556dd0f3ee0)

This field was envisioned to be used for replies that do not fit in a single `__u32`, but in practice this has not occurred as of :octicons-tag-24: v6.3.

### `family`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/40304b2a1567fecc321f640ee4239556dd0f3ee0)

The address family of the socket for which the program is invoked. One of the [`AF_*` enums](https://elixir.bootlin.com/linux/v6.3/source/include/linux/socket.h#L188).

### `remote_ip4`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/40304b2a1567fecc321f640ee4239556dd0f3ee0)

The remote IPv4 address in network byte order if `family` == `AF_INET`.

### `local_ip4`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/40304b2a1567fecc321f640ee4239556dd0f3ee0)

The local IPv4 address in network byte order if `family` == `AF_INET`.

### `remote_ip6`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/40304b2a1567fecc321f640ee4239556dd0f3ee0)

The remote IPv6 address in network byte order if `family` == `AF_INET6`.

### `local_ip6`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/40304b2a1567fecc321f640ee4239556dd0f3ee0)

The local IPv6 address in network byte order if `family` == `AF_INET6`.

### `remote_port`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/40304b2a1567fecc321f640ee4239556dd0f3ee0)

The remote data link / layer 4 port in network byte order.

### `local_port`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/40304b2a1567fecc321f640ee4239556dd0f3ee0)

The local data link / layer 4 port in network byte order.

### `is_fullsock`
	
[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/f19397a5c65665d66e3866b42056f1f58b7a366b)

Some TCP fields are only valid if there is a full socket. If not, the fields read as zero.

### `snd_cwnd`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/f19397a5c65665d66e3866b42056f1f58b7a366b)

The sending congestion window

### `srtt_us`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/f19397a5c65665d66e3866b42056f1f58b7a366b)

The averaged/smoothed RTT (Round Trip Time), stored 3 bits shifted left in μs (microseconds).

```
actual srtt in μs = ctx->srtt_us >> 3;
```

### `bpf_sock_ops_cb_flags`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/b13d880721729384757f235166068c315326f4a1)

This field contains the flags that indicate which optional operations are enabled or not. Possible values are listed in [`include/uapi/linux/bpf.h`](https://elixir.bootlin.com/linux/v6.3/source/include/uapi/linux/bpf.h#L6476). To the change the contents of the field, the [`bpf_sock_ops_cb_flags_set`](../helper-function/bpf_sock_ops_cb_flags_set.md) helper must be used.

### `state`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

This field contains the connection state of the socket.

The states will be one of:
```
enum {
	BPF_TCP_ESTABLISHED = 1,
	BPF_TCP_SYN_SENT,
	BPF_TCP_SYN_RECV,
	BPF_TCP_FIN_WAIT1,
	BPF_TCP_FIN_WAIT2,
	BPF_TCP_TIME_WAIT,
	BPF_TCP_CLOSE,
	BPF_TCP_CLOSE_WAIT,
	BPF_TCP_LAST_ACK,
	BPF_TCP_LISTEN,
	BPF_TCP_CLOSING,	/* Now a valid state */
	BPF_TCP_NEW_SYN_RECV
};
```

### `rtt_min`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

The minimum observed RTT (Round Trip Time)

### `snd_ssthresh`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

The slow start size threshold.

### `rcv_nxt`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

The TCP sequence number we want to receive next.

### `snd_nxt`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

The TCP sequence number we will to send next.

### `snd_una`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

The first byte we want to ACK for.

### `mss_cache`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

Cached effective MSS (Maximum Segment Size), not including SACKS.

### `ecn_flags`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

[ECN](https://en.wikipedia.org/wiki/Explicit_Congestion_Notification) (Explicit Congestion Notification) status bits.

### `rate_delivered`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

Saved rate sample: packets delivered.

### `rate_interval_us`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

Saved rate sample: time elapsed.

### `packets_out`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

Number of packets which are "in flight".

### `retrans_out`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

Number of packets re-transmitted out.

### `total_retrans`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

Total # of packet re-transmits for entire connection.

### `segs_in`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

RFC4898 tcpEStatsPerfSegsIn total number of segments in.

### `data_segs_in`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

RFC4898 tcpEStatsPerfDataSegsIn total number of data segments in.

### `segs_out`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

RFC4898 tcpEStatsPerfSegsOut the total number of segments sent.

### `data_segs_out`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

RFC4898 tcpEStatsPerfDataSegsOut total number of data segments sent.

### `lost_out`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

Number of lost packets.

### `sacked_out`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

Number of <nospell>SACK'd</nospell> packets.

### `sk_txhash`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

Computed flow hash for use on transmit.

### `bytes_received`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

RFC4898 `tcpEStatsAppHCThruOctetsReceived sum(delta(rcv_nxt))`, or how many bytes were acked.

### `bytes_acked`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/44f0e43037d3a17b043843ba67610ac7c7e37db6)

RFC4898 `tcpEStatsAppHCThruOctetsAcked sum(delta(snd_una))`, or how many bytes were acked.

### `sk`

[:octicons-tag-24: v5.3](https://github.com/torvalds/linux/commit/1314ef561102e534e14cb1d37f89f5c1df0b2ea7)

Pointer to the `struct bpf_sock`.

### `skb_data`

[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/0813a841566f0962a5551be7749b43c45f0022a0)

`skb_data` to `skb_data_end` covers the whole TCP header.
    
* `BPF_SOCK_OPS_PARSE_HDR_OPT_CB` - The packet received
* `BPF_SOCK_OPS_HDR_OPT_LEN_CB` - Not useful because the header has not been written.
* `BPF_SOCK_OPS_WRITE_HDR_OPT_CB` - The header and options have been written so far.
* `BPF_SOCK_OPS_ACTIVE_ESTABLISHED_CB` - The SYNACK that concludes the 3WHS.
* `BPF_SOCK_OPS_PASSIVE_ESTABLISHED_CB` - The ACK that concludes the 3WHS.

[`bpf_load_hdr_opt`](../helper-function/bpf_load_hdr_opt.md) can also be used to read a particular option.

### `skb_data_end`

[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/0813a841566f0962a5551be7749b43c45f0022a0)

The end pointer of the TCP header.

### `skb_len`

[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/0813a841566f0962a5551be7749b43c45f0022a0)

The total length of a packet. It includes the header, options, and payload.

### `skb_tcp_flags`

[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/0813a841566f0962a5551be7749b43c45f0022a0)

`tcp_flags` of the header.  It provides an easy way to check for `tcp_flags` without parsing skb_data.

In particular, the `skb_tcp_flags` will still be available in `BPF_SOCK_OPS_HDR_OPT_LEN` even though the outgoing header has not been written yet.

### `skb_hwtstamp`

[:octicons-tag-24: v6.2](https://github.com/torvalds/linux/commit/9bb053490f1a5a0914eb9f7b4116a0e4a95d4f8e)

The timestamp at which the packet was received as reported by the hardware/NIC.

> <nospell>In sockops, the skb is also available to the bpf prog during the `BPF_SOCK_OPS_PARSE_HDR_OPT_CB` event.  There is a use case that the hwtstamp will be useful to the sockops prog to better measure the one-way-delay when the sender has put the tx timestamp in the tcp header option.</nospell>

!!! warning
    `hwtstamps` can only be compared against other `hwtstamps` from the same device.

## Attachment

Socket ops programs are attached to cGroups via the [`BPF_PROG_ATTACH`](../syscall/BPF_PROG_ATTACH.md) syscall or via [BPF link](../syscall/BPF_LINK_CREATE.md).

## Examples

??? example "Clamping a connection"
    ```c
    // Copyright (c) 2017 Facebook
    #define DEBUG 1

    SEC("sockops")
    int bpf_clamp(struct bpf_sock_ops *skops)
    {
        int bufsize = 150000;
        int to_init = 10;
        int clamp = 100;
        int rv = 0;
        int op;

        /* For testing purposes, only execute rest of BPF program
        * if neither port numberis 55601
        */
        if (bpf_ntohl(skops->remote_port) != 55601 && skops->local_port != 55601) {
            skops->reply = -1;
            return 0;
        }

        op = (int) skops->op;

    #ifdef DEBUG
        bpf_printk("BPF command: %d\n", op);
    #endif

        /* Check that both hosts are within same datacenter. For this example
        * it is the case when the first 5.5 bytes of their IPv6 addresses are
        * the same.
        */
        if (skops->family == AF_INET6 &&
            skops->local_ip6[0] == skops->remote_ip6[0] &&
            (bpf_ntohl(skops->local_ip6[1]) & 0xfff00000) ==
            (bpf_ntohl(skops->remote_ip6[1]) & 0xfff00000)) {
            switch (op) {
            case BPF_SOCK_OPS_TIMEOUT_INIT:
                rv = to_init;
                break;
            case BPF_SOCK_OPS_TCP_CONNECT_CB:
                /* Set sndbuf and rcvbuf of active connections */
                rv = bpf_setsockopt(skops, SOL_SOCKET, SO_SNDBUF,
                            &bufsize, sizeof(bufsize));
                rv += bpf_setsockopt(skops, SOL_SOCKET,
                            SO_RCVBUF, &bufsize,
                            sizeof(bufsize));
                break;
            case BPF_SOCK_OPS_ACTIVE_ESTABLISHED_CB:
                rv = bpf_setsockopt(skops, SOL_TCP,
                            TCP_BPF_SNDCWND_CLAMP,
                            &clamp, sizeof(clamp));
                break;
            case BPF_SOCK_OPS_PASSIVE_ESTABLISHED_CB:
                /* Set sndbuf and rcvbuf of passive connections */
                rv = bpf_setsockopt(skops, SOL_TCP,
                            TCP_BPF_SNDCWND_CLAMP,
                            &clamp, sizeof(clamp));
                rv += bpf_setsockopt(skops, SOL_SOCKET,
                            SO_SNDBUF, &bufsize,
                            sizeof(bufsize));
                rv += bpf_setsockopt(skops, SOL_SOCKET,
                            SO_RCVBUF, &bufsize,
                            sizeof(bufsize));
                break;
            default:
                rv = -1;
            }
        } else {
            rv = -1;
        }
    #ifdef DEBUG
        bpf_printk("Returning %d\n", rv);
    #endif
        skops->reply = rv;
        return 1;
    }
    ```

??? example "Dump statistics"
    ```c
    #define INTERVAL			1000000000ULL

    int _version SEC("version") = 1;
    char _license[] SEC("license") = "GPL";

    struct {
        __u32 type;
        __u32 map_flags;
        int *key;
        __u64 *value;
    } bpf_next_dump SEC(".maps") = {
        .type = BPF_MAP_TYPE_SK_STORAGE,
        .map_flags = BPF_F_NO_PREALLOC,
    };

    SEC("sockops")
    int _sockops(struct bpf_sock_ops *ctx)
    {
        struct bpf_tcp_sock *tcp_sk;
        struct bpf_sock *sk;
        __u64 *next_dump;
        __u64 now;

        switch (ctx->op) {
        case BPF_SOCK_OPS_TCP_CONNECT_CB:
            bpf_sock_ops_cb_flags_set(ctx, BPF_SOCK_OPS_RTT_CB_FLAG);
            return 1;
        case BPF_SOCK_OPS_RTT_CB:
            break;
        default:
            return 1;
        }

        sk = ctx->sk;
        if (!sk)
            return 1;

        next_dump = bpf_sk_storage_get(&bpf_next_dump, sk, 0,
                        BPF_SK_STORAGE_GET_F_CREATE);
        if (!next_dump)
            return 1;

        now = bpf_ktime_get_ns();
        if (now < *next_dump)
            return 1;

        tcp_sk = bpf_tcp_sock(sk);
        if (!tcp_sk)
            return 1;

        *next_dump = now + INTERVAL;

        bpf_printk("dsack_dups=%u delivered=%u\n",
            tcp_sk->dsack_dups, tcp_sk->delivered);
        bpf_printk("delivered_ce=%u icsk_retransmits=%u\n",
            tcp_sk->delivered_ce, tcp_sk->icsk_retransmits);

        return 1;
    }
    ```

??? example "Adding socket to map"
    ```c
    // Copyright (c) 2017-2018 Covalent IO
    SEC("sockops")
    int bpf_sockmap(struct bpf_sock_ops *skops)
    {
        __u32 lport, rport;
        int op, err = 0, index, key, ret;


        op = (int) skops->op;

        switch (op) {
        case BPF_SOCK_OPS_PASSIVE_ESTABLISHED_CB:
            lport = skops->local_port;
            rport = skops->remote_port;

            if (lport == 10000) {
                ret = 1;
    #ifdef SOCKMAP
                err = bpf_sock_map_update(skops, &sock_map, &ret,
                            BPF_NOEXIST);
    #else
                err = bpf_sock_hash_update(skops, &sock_map, &ret,
                            BPF_NOEXIST);
    #endif
            }
            break;
        case BPF_SOCK_OPS_ACTIVE_ESTABLISHED_CB:
            lport = skops->local_port;
            rport = skops->remote_port;

            if (bpf_ntohl(rport) == 10001) {
                ret = 10;
    #ifdef SOCKMAP
                err = bpf_sock_map_update(skops, &sock_map, &ret,
                            BPF_NOEXIST);
    #else
                err = bpf_sock_hash_update(skops, &sock_map, &ret,
                            BPF_NOEXIST);
    #endif
            }
            break;
        default:
            break;
        }

        return 0;
    }
    ```

## Helper functions

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [`bpf_cgrp_storage_delete`](../helper-function/bpf_cgrp_storage_delete.md)
    * [`bpf_cgrp_storage_get`](../helper-function/bpf_cgrp_storage_get.md)
    * [`bpf_dynptr_data`](../helper-function/bpf_dynptr_data.md)
    * [`bpf_dynptr_from_mem`](../helper-function/bpf_dynptr_from_mem.md)
    * [`bpf_dynptr_read`](../helper-function/bpf_dynptr_read.md)
    * [`bpf_dynptr_write`](../helper-function/bpf_dynptr_write.md)
    * [`bpf_for_each_map_elem`](../helper-function/bpf_for_each_map_elem.md)
    * [`bpf_get_current_pid_tgid`](../helper-function/bpf_get_current_pid_tgid.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
    * [`bpf_get_current_task`](../helper-function/bpf_get_current_task.md)
    * [`bpf_get_current_task_btf`](../helper-function/bpf_get_current_task_btf.md)
    * [`bpf_get_local_storage`](../helper-function/bpf_get_local_storage.md)
    * [`bpf_get_netns_cookie`](../helper-function/bpf_get_netns_cookie.md)
    * [`bpf_get_ns_current_pid_tgid`](../helper-function/bpf_get_ns_current_pid_tgid.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
    * [`bpf_get_numa_node_id`](../helper-function/bpf_get_numa_node_id.md)
    * [`bpf_get_prandom_u32`](../helper-function/bpf_get_prandom_u32.md)
    * [`bpf_get_smp_processor_id`](../helper-function/bpf_get_smp_processor_id.md)
    * [`bpf_get_socket_cookie`](../helper-function/bpf_get_socket_cookie.md)
    * [`bpf_getsockopt`](../helper-function/bpf_getsockopt.md)
    * [`bpf_jiffies64`](../helper-function/bpf_jiffies64.md)
    * [`bpf_kptr_xchg`](../helper-function/bpf_kptr_xchg.md)
    * [`bpf_ktime_get_boot_ns`](../helper-function/bpf_ktime_get_boot_ns.md)
    * [`bpf_ktime_get_coarse_ns`](../helper-function/bpf_ktime_get_coarse_ns.md)
    * [`bpf_ktime_get_ns`](../helper-function/bpf_ktime_get_ns.md)
    * [`bpf_ktime_get_tai_ns`](../helper-function/bpf_ktime_get_tai_ns.md)
    * [`bpf_load_hdr_opt`](../helper-function/bpf_load_hdr_opt.md)
    * [`bpf_loop`](../helper-function/bpf_loop.md)
    * [`bpf_map_delete_elem`](../helper-function/bpf_map_delete_elem.md)
    * [`bpf_map_lookup_elem`](../helper-function/bpf_map_lookup_elem.md)
    * [`bpf_map_lookup_percpu_elem`](../helper-function/bpf_map_lookup_percpu_elem.md)
    * [`bpf_map_peek_elem`](../helper-function/bpf_map_peek_elem.md)
    * [`bpf_map_pop_elem`](../helper-function/bpf_map_pop_elem.md)
    * [`bpf_map_push_elem`](../helper-function/bpf_map_push_elem.md)
    * [`bpf_map_update_elem`](../helper-function/bpf_map_update_elem.md)
    * [`bpf_per_cpu_ptr`](../helper-function/bpf_per_cpu_ptr.md)
    * [`bpf_perf_event_output`](../helper-function/bpf_perf_event_output.md)
    * [`bpf_probe_read_kernel`](../helper-function/bpf_probe_read_kernel.md)
    * [`bpf_probe_read_kernel_str`](../helper-function/bpf_probe_read_kernel_str.md)
    * [`bpf_probe_read_user`](../helper-function/bpf_probe_read_user.md)
    * [`bpf_probe_read_user_str`](../helper-function/bpf_probe_read_user_str.md)
    * [`bpf_reserve_hdr_opt`](../helper-function/bpf_reserve_hdr_opt.md)
    * [`bpf_ringbuf_discard`](../helper-function/bpf_ringbuf_discard.md)
    * [`bpf_ringbuf_discard_dynptr`](../helper-function/bpf_ringbuf_discard_dynptr.md)
    * [`bpf_ringbuf_output`](../helper-function/bpf_ringbuf_output.md)
    * [`bpf_ringbuf_query`](../helper-function/bpf_ringbuf_query.md)
    * [`bpf_ringbuf_reserve`](../helper-function/bpf_ringbuf_reserve.md)
    * [`bpf_ringbuf_reserve_dynptr`](../helper-function/bpf_ringbuf_reserve_dynptr.md)
    * [`bpf_ringbuf_submit`](../helper-function/bpf_ringbuf_submit.md)
    * [`bpf_ringbuf_submit_dynptr`](../helper-function/bpf_ringbuf_submit_dynptr.md)
    * [`bpf_setsockopt`](../helper-function/bpf_setsockopt.md)
    * [`bpf_sk_storage_delete`](../helper-function/bpf_sk_storage_delete.md)
    * [`bpf_sk_storage_get`](../helper-function/bpf_sk_storage_get.md)
    * [`bpf_skc_to_tcp6_sock`](../helper-function/bpf_skc_to_tcp6_sock.md)
    * [`bpf_skc_to_tcp_request_sock`](../helper-function/bpf_skc_to_tcp_request_sock.md)
    * [`bpf_skc_to_tcp_sock`](../helper-function/bpf_skc_to_tcp_sock.md)
    * [`bpf_skc_to_tcp_timewait_sock`](../helper-function/bpf_skc_to_tcp_timewait_sock.md)
    * [`bpf_skc_to_udp6_sock`](../helper-function/bpf_skc_to_udp6_sock.md)
    * [`bpf_skc_to_unix_sock`](../helper-function/bpf_skc_to_unix_sock.md)
    * [`bpf_snprintf`](../helper-function/bpf_snprintf.md)
    * [`bpf_snprintf_btf`](../helper-function/bpf_snprintf_btf.md)
    * [`bpf_sock_hash_update`](../helper-function/bpf_sock_hash_update.md)
    * [`bpf_sock_map_update`](../helper-function/bpf_sock_map_update.md)
    * [`bpf_sock_ops_cb_flags_set`](../helper-function/bpf_sock_ops_cb_flags_set.md)
    * [`bpf_spin_lock`](../helper-function/bpf_spin_lock.md)
    * [`bpf_spin_unlock`](../helper-function/bpf_spin_unlock.md)
    * [`bpf_store_hdr_opt`](../helper-function/bpf_store_hdr_opt.md)
    * [`bpf_strncmp`](../helper-function/bpf_strncmp.md)
    * [`bpf_tail_call`](../helper-function/bpf_tail_call.md)
    * [`bpf_task_pt_regs`](../helper-function/bpf_task_pt_regs.md)
    * [`bpf_tcp_sock`](../helper-function/bpf_tcp_sock.md)
    * [`bpf_this_cpu_ptr`](../helper-function/bpf_this_cpu_ptr.md)
    * [`bpf_timer_cancel`](../helper-function/bpf_timer_cancel.md)
    * [`bpf_timer_init`](../helper-function/bpf_timer_init.md)
    * [`bpf_timer_set_callback`](../helper-function/bpf_timer_set_callback.md)
    * [`bpf_timer_start`](../helper-function/bpf_timer_start.md)
    * [`bpf_trace_printk`](../helper-function/bpf_trace_printk.md)
    * [`bpf_trace_vprintk`](../helper-function/bpf_trace_vprintk.md)
    * [`bpf_user_ringbuf_drain`](../helper-function/bpf_user_ringbuf_drain.md)
<!-- [/PROG_HELPER_FUNC_REF] -->

## KFuncs

<!-- [PROG_KFUNC_REF] -->
??? abstract "Supported kfuncs"
    - [`bpf_arena_alloc_pages`](../kfuncs/bpf_arena_alloc_pages.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_arena_free_pages`](../kfuncs/bpf_arena_free_pages.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_cast_to_kern_ctx`](../kfuncs/bpf_cast_to_kern_ctx.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_copy_from_user_dynptr`](../kfuncs/bpf_copy_from_user_dynptr.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_copy_from_user_str`](../kfuncs/bpf_copy_from_user_str.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_copy_from_user_str_dynptr`](../kfuncs/bpf_copy_from_user_str_dynptr.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_copy_from_user_task_dynptr`](../kfuncs/bpf_copy_from_user_task_dynptr.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_copy_from_user_task_str`](../kfuncs/bpf_copy_from_user_task_str.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_copy_from_user_task_str_dynptr`](../kfuncs/bpf_copy_from_user_task_str_dynptr.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_dynptr_adjust`](../kfuncs/bpf_dynptr_adjust.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_dynptr_clone`](../kfuncs/bpf_dynptr_clone.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_dynptr_copy`](../kfuncs/bpf_dynptr_copy.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_dynptr_from_skb`](../kfuncs/bpf_dynptr_from_skb.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_dynptr_is_null`](../kfuncs/bpf_dynptr_is_null.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_dynptr_is_rdonly`](../kfuncs/bpf_dynptr_is_rdonly.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_dynptr_size`](../kfuncs/bpf_dynptr_size.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_dynptr_slice`](../kfuncs/bpf_dynptr_slice.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_dynptr_slice_rdwr`](../kfuncs/bpf_dynptr_slice_rdwr.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_get_kmem_cache`](../kfuncs/bpf_get_kmem_cache.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_iter_bits_destroy`](../kfuncs/bpf_iter_bits_destroy.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_iter_bits_new`](../kfuncs/bpf_iter_bits_new.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_iter_bits_next`](../kfuncs/bpf_iter_bits_next.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_iter_css_destroy`](../kfuncs/bpf_iter_css_destroy.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_iter_css_new`](../kfuncs/bpf_iter_css_new.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_iter_css_next`](../kfuncs/bpf_iter_css_next.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_iter_css_task_destroy`](../kfuncs/bpf_iter_css_task_destroy.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_iter_css_task_new`](../kfuncs/bpf_iter_css_task_new.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_iter_css_task_next`](../kfuncs/bpf_iter_css_task_next.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_iter_kmem_cache_destroy`](../kfuncs/bpf_iter_kmem_cache_destroy.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_iter_kmem_cache_new`](../kfuncs/bpf_iter_kmem_cache_new.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_iter_kmem_cache_next`](../kfuncs/bpf_iter_kmem_cache_next.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_iter_num_destroy`](../kfuncs/bpf_iter_num_destroy.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_iter_num_new`](../kfuncs/bpf_iter_num_new.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_iter_num_next`](../kfuncs/bpf_iter_num_next.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_iter_task_destroy`](../kfuncs/bpf_iter_task_destroy.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_iter_task_new`](../kfuncs/bpf_iter_task_new.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_iter_task_next`](../kfuncs/bpf_iter_task_next.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_iter_task_vma_destroy`](../kfuncs/bpf_iter_task_vma_destroy.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_iter_task_vma_new`](../kfuncs/bpf_iter_task_vma_new.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_iter_task_vma_next`](../kfuncs/bpf_iter_task_vma_next.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_local_irq_restore`](../kfuncs/bpf_local_irq_restore.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_local_irq_save`](../kfuncs/bpf_local_irq_save.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_map_sum_elem_count`](../kfuncs/bpf_map_sum_elem_count.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_preempt_disable`](../kfuncs/bpf_preempt_disable.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_preempt_enable`](../kfuncs/bpf_preempt_enable.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_probe_read_kernel_dynptr`](../kfuncs/bpf_probe_read_kernel_dynptr.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_probe_read_kernel_str_dynptr`](../kfuncs/bpf_probe_read_kernel_str_dynptr.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_probe_read_user_dynptr`](../kfuncs/bpf_probe_read_user_dynptr.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_probe_read_user_str_dynptr`](../kfuncs/bpf_probe_read_user_str_dynptr.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_rcu_read_lock`](../kfuncs/bpf_rcu_read_lock.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_rcu_read_unlock`](../kfuncs/bpf_rcu_read_unlock.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_rdonly_cast`](../kfuncs/bpf_rdonly_cast.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_res_spin_lock`](../kfuncs/bpf_res_spin_lock.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_res_spin_lock_irqsave`](../kfuncs/bpf_res_spin_lock_irqsave.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_res_spin_unlock`](../kfuncs/bpf_res_spin_unlock.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_res_spin_unlock_irqrestore`](../kfuncs/bpf_res_spin_unlock_irqrestore.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_sock_addr_set_sun_path`](../kfuncs/bpf_sock_addr_set_sun_path.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_sock_ops_enable_tx_tstamp`](../kfuncs/bpf_sock_ops_enable_tx_tstamp.md)
    - [`bpf_wq_init`](../kfuncs/bpf_wq_init.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_wq_set_callback_impl`](../kfuncs/bpf_wq_set_callback_impl.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
    - [`bpf_wq_start`](../kfuncs/bpf_wq_start.md) [:octicons-tag-24: v6.15](https://github.com/torvalds/linux/commit/59422464266f8baa091edcb3779f0955a21abf00) - 
<!-- [/PROG_KFUNC_REF] -->
