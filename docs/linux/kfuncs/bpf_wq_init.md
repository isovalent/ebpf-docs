---
title: "KFunc 'bpf_wq_init'"
description: "This page documents the 'bpf_wq_init' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_wq_init`

<!-- [FEATURE_TAG](bpf_wq_init) -->
[:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb48f6cd41a0f7803770a76bbffb6bd5b1b2ae2f)
<!-- [/FEATURE_TAG] -->

Initialize a work-queue.

## Definition

This kfunc initializes a work-queue which allows eBPF programs to schedule work to be executed asynchronously.

`wq`: A pointer to a `struct bpf_wq` which must reside in a map value.

`p__map`: A pointer to a map that contains the `wq` as value.

`flags`: Flags to allow for future extensions.

**Returns**

Return `0` on success, or a negative error code on failure.

<!-- [KFUNC_DEF] -->
`#!c int bpf_wq_init(struct bpf_wq *wq, void *p__map, unsigned int flags)`
<!-- [/KFUNC_DEF] -->

## Usage

This is the first step in using a work-queue. A work-queue is a mechanism to schedule work to be executed asynchronously. After initialization a callback function can be associated with the work-queue using the [`bpf_wq_set_callback_impl`](bpf_wq_set_callback_impl.md) kfunc and the work can be started using the [`bpf_wq_start`](bpf_wq_start.md) kfunc.

The callback will be called asynchronously sometime after the current eBPF program has finished executing whenever the scheduler decides to run the work-queue.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
- [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_LWT_IN`](../program-type/BPF_PROG_TYPE_LWT_IN.md)
- [`BPF_PROG_TYPE_LWT_OUT`](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
- [`BPF_PROG_TYPE_LWT_SEG6LOCAL`](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
- [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
- [`BPF_PROG_TYPE_NETFILTER`](../program-type/BPF_PROG_TYPE_NETFILTER.md)
- [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
- [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md)
- [`BPF_PROG_TYPE_SOCKET_FILTER`](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
- [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example
    ```c
    /* Copyright (c) 2024 Benjamin Tissoires */

    #include <bpf/bpf_helpers.h>

    char _license[] SEC("license") = "GPL";

    struct elem {
        struct bpf_wq w;
    };

    struct {
        __uint(type, BPF_MAP_TYPE_ARRAY);
        __uint(max_entries, 2);
        __type(key, int);
        __type(value, struct elem);
    } array SEC(".maps");

    __u32 ok;
    __u32 ok_sleepable;
    void bpf_kfunc_common_test(void) __ksym;

    static int test_elem_callback(void *map, int *key,
            int (callback_fn)(void *map, int *key, struct bpf_wq *wq))
    {
        struct elem init = {}, *val;
        struct bpf_wq *wq;

        if ((ok & (1 << *key) ||
            (ok_sleepable & (1 << *key))))
            return -22;

        if (map == &lru &&
            bpf_map_update_elem(map, key, &init, 0))
            return -1;

        val = bpf_map_lookup_elem(map, key);
        if (!val)
            return -2;

        wq = &val->w;
        if (bpf_wq_init(wq, map, 0) != 0)
            return -3;

        if (bpf_wq_set_callback(wq, callback_fn, 0))
            return -4;

        if (bpf_wq_start(wq, 0))
            return -5;

        return 0;
    }

    /* callback for non sleepable workqueue */
    static int wq_callback(void *map, int *key, struct bpf_wq *work)
    {
        bpf_kfunc_common_test();
        ok |= (1 << *key);
        return 0;
    }

    SEC("tc")
    long test_call_array_sleepable(void *ctx)
    {
        int key = 0;

        return test_elem_callback(&array, &key, wq_cb_sleepable);
    }
    ```
