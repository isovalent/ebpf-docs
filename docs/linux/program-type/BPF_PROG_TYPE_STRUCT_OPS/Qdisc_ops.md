---
title: "Struct ops 'Qdisc_ops'"
description: "This page documents the 'Qdisc_ops' struct ops, its semantics, capabilities, and limitations."
---
# Struct ops `Qdisc_ops`

[:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/c8240344956e3f0b4e8f1d40ec3435e47040cacb)

Qdisc ops is a type of struct_ops which allows the implementation of a custom [qdisc](https://man7.org/linux/man-pages/man8/tc.8.html#QDISCS) in BPF.

## Usage

BPF qdisc aims to be a flexible and easy-to-use infrastructure that allows users to quickly experiment with different scheduling algorithms/policies.

## Fields and ops

```c
struct Qdisc_ops {
        struct Qdisc_ops        *[next](#next);
        const struct Qdisc_class_ops    *[cl_ops](#cl_ops);
        char                    [id](#id)[IFNAMSIZ];
        int                     [priv_size](#priv_size);
        unsigned int            [static_flags](#static_flags);

        int                     (*[enqueue](#enqueue))([struct sk_buff](#struct-sk_buff) *skb,
                                           [struct Qdisc](#struct-qdisc) *sch,
                                           [struct sk_buff](#struct-sk_buff) **to_free);
        [struct sk_buff](#struct-sk_buff) *        (*[dequeue](#dequeue))([struct Qdisc](#struct-qdisc) *);
        [struct sk_buff](#struct-sk_buff) *        (*[peek](#peek))([struct Qdisc](#struct-qdisc) *);

        int                     (*[init](#init))([struct Qdisc](#struct-qdisc) *sch, struct nlattr *arg,
                                        struct netlink_ext_ack *extack);
        void                    (*[reset](#reset))([struct Qdisc](#struct-qdisc) *);
        void                    (*[destroy](#destroy))([struct Qdisc](#struct-qdisc) *);
        int                     (*[change](#change))([struct Qdisc](#struct-qdisc) *sch,
                                          struct nlattr *arg,
                                          struct netlink_ext_ack *extack);
        void                    (*[attach](#attach))([struct Qdisc](#struct-qdisc) *sch);
        int                     (*[change_tx_queue_len](#change_tx_queue_len))([struct Qdisc](#struct-qdisc) *, unsigned int);
        void                    (*[change_real_num_tx](#change_real_num_tx))([struct Qdisc](#struct-qdisc) *sch,
                                                      unsigned int new_real_tx);

        int                     (*[dump](#dump))([struct Qdisc](#struct-qdisc) *, [struct sk_buff](#struct-sk_buff) *);
        int                     (*[dump_stats](#dump_stats))([struct Qdisc](#struct-qdisc) *, struct gnet_dump *);

        void                    (*[ingress_block_set](#ingress_block_set))([struct Qdisc](#struct-qdisc) *sch,
                                                     u32 block_index);
        void                    (*[egress_block_set](#egress_block_set))([struct Qdisc](#struct-qdisc) *sch,
                                                    u32 block_index);
        u32                     (*[ingress_block_get](#ingress_block_get))([struct Qdisc](#struct-qdisc) *sch);
        u32                     (*[egress_block_get](#egress_block_get))([struct Qdisc](#struct-qdisc) *sch);

        struct module           *[owner](#owner);
};
```

### `next`

`#!c struct Qdisc_ops *next;`

All registered Qdisc_ops are linked together in a linked list, this field is the linked list header. It should always be unspecified by BPF as its managed by the kernel.

### `cl_ops`

`#!c const struct Qdisc_class_ops *cl_ops;`

This field are operations specific to classful qdiscs, which are not yet implemented as of v6.16.

### `id`

`#!c char id[IFNAMSIZ];`

The unique identifier of this qdisc type.

### `priv_size`

`#!c int priv_size;`

The amount of bytes stored in [`qdisc->privdata`](#struct-qdisc_ops-privdata). Typically used by builtin qdisc types, not available to BPF Qdisc as of v6.16.

### `static_flags`

`#!c unsigned int static_flags;`

A set of flags which will be the initial value of [`qdisc->flags`](#struct-qdisc-flags).

### `enqueue`

`#!c int (*enqueue)(struct sk_buff *skb, struct Qdisc *sch, struct sk_buff **to_free);`

This op is called for every packet that is directed to a network device using the `sch` instance of this BPF qdisc. The `skb` is the packet that should be added to whatever data structure is implemented by the BPF qdisc.

Recognized return values are:

* `NET_XMIT_SUCCESS` (0x00) - The packet has been added to the data structure.
* `NET_XMIT_DROP` (0x01) - The packet has been dropped.
* `NET_XMIT_CN` (0x02) - Does not guarantee that this packet is lost. It indicates that the device will soon be dropping packets, or already drops some packets of the same priority; prompting us to send less aggressively.

If adding this packet causes other packets (perhaps older or of lower priority) to be removed from the data structure, the [`bpf_qdisc_skb_drop`](../../kfuncs/bpf_qdisc_skb_drop.md) kfunc can be used to enqueue these other packets to `to_free` so memory associated with these dropped packets is freed.

### `dequeue`

`#!c struct sk_buff * (*dequeue)(struct Qdisc *sch);`

This op is called periodically when a network device with the `sch` instance of this BPF Qdisc is ready to send packets. The op should return a packet to be sent if available, or `NULL` if no packet is available. The returned packet should be removed from the data structure.

### `peek`

`#!c struct sk_buff * (*peek)(struct Qdisc *sch);`

This op can be called by a network device with the `sch` instance of this BPF Qdisc to see if any packets are available and which would be first. The op should return a packet to be sent if available, or `NULL` if no packet is available. The packet will not actually be sent, and the BPF Qdisc should hold onto the packet in its data structure.


### `init`

`#!c int (*init)(struct Qdisc *sch, struct nlattr *arg, struct netlink_ext_ack *extack);`

This op is called to initialize a qdisc instance `sch` which will use the current ops for its implementation. `arg` are the netlink arguments used to create this new qdisc instance.

`extack` is the extended acknowledge, used to carry verbose error messages, which BPF qdiscs cannot utilize as of v6.16.

### `reset`

`#!c void (*reset)(struct Qdisc *);`

This op is called on an already initialized qdisc instance to reset it to its initial state.

### `destroy`

`#!c void (*destroy)(struct Qdisc *);`

This op is called when a qdisc instance is no longer in use, such as when its switched out or the device goes away.

### `change`

`#!c int (*change)(struct Qdisc *sch, struct nlattr *arg, struct netlink_ext_ack *extack);`

This op is called when settings for qdisc instance `sch` are updated after it has been initialized. `arg` are the netlink arguments containing the new settings.

`extack` is the extended acknowledge, used to carry verbose error messages, which BPF qdiscs cannot utilize as of v6.16.

### `attach`

`#!c void (*attach)(struct Qdisc *sch);`

This op is called when qdisc instance `sch` is attached to a network device.

### `change_tx_queue_len`

`#!c int (*change_tx_queue_len)(struct Qdisc *sch, unsigned int new_len);`

This op is called when a change of transmission queue length is requested. Returning `0` indicates success, any other value indicates an error.

### `change_real_num_tx`

`#!c void (*change_real_num_tx)(struct Qdisc *sch, unsigned int new_real_tx);`

This op is called to inform the qdisc of the number of transmission queues used by the network device to which qdisc instance `sch` is attached.

### `dump`

`#!c int (*dump)(struct Qdisc *sch, struct sk_buff *skb);`

This op is called to dump information such as settings of the current qdisc instance `sch`. `skb` is a socket buffer (network packet) which will be the netlink message sent to userspace. The packet will already have other netlink data in there, this op is expected to append netlink attributes to the end, being pre-defined or custom, as long as netlink "type-length-attribute" format is used.

Returning `NULL` means success, a negative value indicates failure.

### `dump_stats`

`#!c int (*dump_stats)(struct Qdisc *sch, struct gnet_dump *d);`

This op is called to fill `d` with statistics about qdisc instance `sch`.

Returning `NULL` means success, a negative value indicates failure.

### `ingress_block_set`

`#!c void (*ingress_block_set)(struct Qdisc *sch, u32 block_index);`

This op is called when a [TC block](https://lwn.net/Articles/946802/) is specified for ingress via the `TCA_INGRESS_BLOCK` attribute.

### `egress_block_set`

`#!c void (*egress_block_set)(struct Qdisc *sch, u32 block_index);`

This op is called when a [TC block](https://lwn.net/Articles/946802/) is specified for egress via the `TCA_EGRESS_BLOCK` attribute.

### `ingress_block_get`

`#!c u32 (*ingress_block_get)(struct Qdisc *sch);`

This op is used to query the current ingress [TC block](https://lwn.net/Articles/946802/) associated with this qdisc instance `sch`.

### `egress_block_get`

`#!c u32 (*egress_block_get)(struct Qdisc *sch);`

This op is used to query the current egress [TC block](https://lwn.net/Articles/946802/) associated with this qdisc instance `sch`.

### `owner`

`#!c struct module *owner;`

This is a field internally used by the kernel to associate an owner for the ops.

## Types

### `struct Qdisc`

```c
struct Qdisc {
    int                    (*enqueue)(struct sk_buff *skb,
                                      struct Qdisc *sch,
                                      struct sk_buff **to_free);
    struct sk_buff *       (*dequeue)(struct Qdisc *sch);
    unsigned int             [flags](#struct-qdisc-flags);
    u32                      limit;
    const struct Qdisc_ops  *ops;
    struct qdisc_size_table __rcu *stab;
    struct hlist_node        hash;
    u32                      handle;
    u32                      parent;

    struct netdev_queue *dev_queue;

    struct net_rate_estimator __rcu         *rate_est;
    struct gnet_stats_basic_sync __percpu   *cpu_bstats;
    struct gnet_stats_queue __percpu        *cpu_qstats;
    int                                     pad;
    refcount_t                              refcnt;

    /*
     * For performance sake on SMP, we put highly modified fields at the end
     */
    struct sk_buff_head             gso_skb ____cacheline_aligned_in_smp;
    struct qdisc_skb_head           q;
    struct gnet_stats_basic_sync    bstats;
    struct gnet_stats_queue         qstats;
    int                             owner;
    unsigned long                   state;
    unsigned long                   state2; /* must be written under qdisc spinlock */
    struct Qdisc                   *next_sched;
    struct sk_buff_head             skb_bad_txq;

    spinlock_t busylock ____cacheline_aligned_in_smp;
    spinlock_t seqlock;

    struct rcu_head        rcu;
    netdevice_tracker      dev_tracker;
    struct lock_class_key  root_lock_key;
    /* private data */
    long privdata[] ____cacheline_aligned;
};
```

All fields are read-only accept for `limit`, `q->qlen`, and `qstats`.

#### `flags` {#struct-qdisc-flags}

```c
#define TCQ_F_BUILTIN       1
#define TCQ_F_INGRESS       2
#define TCQ_F_CAN_BYPASS    4
#define TCQ_F_MQROOT        8
#define TCQ_F_ONETXQUEUE    0x10 /* dequeue_skb() can assume all skbs are for
                                 * q->dev_queue : It can test
                                 * netif_xmit_frozen_or_stopped() before
                                 * dequeueing next packet.
                                 * Its true for MQ/MQPRIO slaves, or non
                                 * multiqueue device.
                                 */
#define TCQ_F_WARN_NONWC    (1 << 16)
#define TCQ_F_CPUSTATS      0x20 /* run using percpu statistics */
#define TCQ_F_NOPARENT      0x40 /* root of its hierarchy :
                                  * qdisc_tree_decrease_qlen() should stop.
                                  */
#define TCQ_F_INVISIBLE     0x80 /* invisible by default in dump */
#define TCQ_F_NOLOCK        0x100 /* qdisc does not require locking */
#define TCQ_F_OFFLOADED     0x200 /* qdisc is offloaded to HW */
```

### `privdata` {#struct-qdisc_ops-privdata}

### `struct sk_buff`

See [`skbuff.h`](https://elixir.bootlin.com/linux/v6.16/source/include/linux/skbuff.h#L769) for full structure.

All fields are read-only, except for `tstamp` and `cb`.

## Example

### FIFO (First In First Out)

A BPF implementation of the most basic qdisc, a simple queue without scheduling or ordering logic.

```c
// SPDX-License-Identifier: GPL-2.0
// Copyright by Amery Hung <amery.hung@bytedance.com>

#include <vmlinux.h>
#include "bpf_experimental.h"
#include "bpf_qdisc_common.h"

char _license[] SEC("license") = "GPL";

struct skb_node {
	struct sk_buff __kptr * skb;
	struct bpf_list_node node;
};

private(A) struct bpf_spin_lock q_fifo_lock;
private(A) struct bpf_list_head q_fifo __contains(skb_node, node);

bool init_called;

SEC("struct_ops/bpf_fifo_enqueue")
int BPF_PROG(bpf_fifo_enqueue, struct sk_buff *skb, struct Qdisc *sch,
	     struct bpf_sk_buff_ptr *to_free)
{
	struct skb_node *skbn;
	u32 pkt_len;

	if (sch->q.qlen == sch->limit)
		goto drop;

	skbn = bpf_obj_new(typeof(*skbn));
	if (!skbn)
		goto drop;

	pkt_len = qdisc_pkt_len(skb);

	sch->q.qlen++;
	skb = bpf_kptr_xchg(&skbn->skb, skb);
	if (skb)
		bpf_qdisc_skb_drop(skb, to_free);

	bpf_spin_lock(&q_fifo_lock);
	bpf_list_push_back(&q_fifo, &skbn->node);
	bpf_spin_unlock(&q_fifo_lock);

	sch->qstats.backlog += pkt_len;
	return NET_XMIT_SUCCESS;
drop:
	bpf_qdisc_skb_drop(skb, to_free);
	return NET_XMIT_DROP;
}

SEC("struct_ops/bpf_fifo_dequeue")
struct sk_buff *BPF_PROG(bpf_fifo_dequeue, struct Qdisc *sch)
{
	struct bpf_list_node *node;
	struct sk_buff *skb = NULL;
	struct skb_node *skbn;

	bpf_spin_lock(&q_fifo_lock);
	node = bpf_list_pop_front(&q_fifo);
	bpf_spin_unlock(&q_fifo_lock);
	if (!node)
		return NULL;

	skbn = container_of(node, struct skb_node, node);
	skb = bpf_kptr_xchg(&skbn->skb, skb);
	bpf_obj_drop(skbn);
	if (!skb)
		return NULL;

	sch->qstats.backlog -= qdisc_pkt_len(skb);
	bpf_qdisc_bstats_update(sch, skb);
	sch->q.qlen--;

	return skb;
}

SEC("struct_ops/bpf_fifo_init")
int BPF_PROG(bpf_fifo_init, struct Qdisc *sch, struct nlattr *opt,
	     struct netlink_ext_ack *extack)
{
	sch->limit = 1000;
	init_called = true;
	return 0;
}

SEC("struct_ops/bpf_fifo_reset")
void BPF_PROG(bpf_fifo_reset, struct Qdisc *sch)
{
	struct bpf_list_node *node;
	struct skb_node *skbn;
	int i;

	bpf_for(i, 0, sch->q.qlen) {
		struct sk_buff *skb = NULL;

		bpf_spin_lock(&q_fifo_lock);
		node = bpf_list_pop_front(&q_fifo);
		bpf_spin_unlock(&q_fifo_lock);

		if (!node)
			break;

		skbn = container_of(node, struct skb_node, node);
		skb = bpf_kptr_xchg(&skbn->skb, skb);
		if (skb)
			bpf_kfree_skb(skb);
		bpf_obj_drop(skbn);
	}
	sch->q.qlen = 0;
}

SEC("struct_ops")
void BPF_PROG(bpf_fifo_destroy, struct Qdisc *sch)
{
}

SEC(".struct_ops")
struct Qdisc_ops fifo = {
	.enqueue   = (void *)bpf_fifo_enqueue,
	.dequeue   = (void *)bpf_fifo_dequeue,
	.init      = (void *)bpf_fifo_init,
	.reset     = (void *)bpf_fifo_reset,
	.destroy   = (void *)bpf_fifo_destroy,
	.id        = "bpf_fifo",
};
```
