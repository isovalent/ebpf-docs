---
title: "Helper Function 'bpf_user_ringbuf_drain'"
description: "This page documents the 'bpf_user_ringbuf_drain' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_user_ringbuf_drain`

<!-- [FEATURE_TAG](bpf_user_ringbuf_drain) -->
[:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/20571567384428dfc9fe5cf9f2e942e1df13c2dd)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Drain samples from the specified user ring buffer, and invoke the provided callback for each such sample:

long (\_callback_fn)(const struct bpf_dynptr \_dynptr, void \_ctx);

If **callback_fn** returns 0, the helper will continue to try and drain the next sample, up to a maximum of BPF_MAX_USER_RINGBUF_SAMPLES samples. If the return value is 1, the helper will skip the rest of the samples and return. Other return values are not used now, and will be rejected by the verifier.

### Returns

The number of drained samples if no error was encountered while draining samples, or 0 if no samples were present in the ring buffer. If a user-space producer was epoll-waiting on this map, and at least one sample was drained, they will receive an event notification notifying them of available space in the ring buffer. If the BPF_RB_NO_WAKEUP flag is passed to this function, no wakeup notification will be sent. If the BPF_RB_FORCE_WAKEUP flag is passed, a wakeup notification will be sent even if no sample was drained.

On failure, the returned value is one of the following:

**-EBUSY** if the ring buffer is contended, and another calling context was concurrently draining the ring buffer.

**-EINVAL** if user-space is not properly tracking the ring buffer due to the producer position not being aligned to 8 bytes, a sample not being aligned to 8 bytes, or the producer position not matching the advertised length of a sample.

**-E2BIG** if user-space has tried to publish a sample which is larger than the size of the ring buffer, or which cannot fit within a struct bpf_dynptr.

`#!c static long (* const bpf_user_ringbuf_drain)(void *map, void *callback_fn, void *ctx, __u64 flags) = (void *) 209;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_DEVICE](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md)
 * [BPF_PROG_TYPE_CGROUP_SKB](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
 * [BPF_PROG_TYPE_CGROUP_SOCKOPT](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK_ADDR](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [BPF_PROG_TYPE_CGROUP_SYSCTL](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
 * [BPF_PROG_TYPE_FLOW_DISSECTOR](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md)
 * [BPF_PROG_TYPE_KPROBE](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
 * [BPF_PROG_TYPE_LWT_IN](../program-type/BPF_PROG_TYPE_LWT_IN.md)
 * [BPF_PROG_TYPE_LWT_OUT](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
 * [BPF_PROG_TYPE_LWT_SEG6LOCAL](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
 * [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [BPF_PROG_TYPE_PERF_EVENT](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [BPF_PROG_TYPE_SK_LOOKUP](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md)
 * [BPF_PROG_TYPE_SK_MSG](../program-type/BPF_PROG_TYPE_SK_MSG.md)
 * [BPF_PROG_TYPE_SK_REUSEPORT](../program-type/BPF_PROG_TYPE_SK_REUSEPORT.md)
 * [BPF_PROG_TYPE_SK_SKB](../program-type/BPF_PROG_TYPE_SK_SKB.md)
 * [BPF_PROG_TYPE_SOCKET_FILTER](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
 * [BPF_PROG_TYPE_SOCK_OPS](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
 * [BPF_PROG_TYPE_SYSCALL](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [BPF_PROG_TYPE_TRACEPOINT](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
 * [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
