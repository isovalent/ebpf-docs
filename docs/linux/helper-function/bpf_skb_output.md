---
title: "Helper Function 'bpf_skb_output'"
description: "This page documents the 'bpf_skb_output' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_skb_output`

<!-- [FEATURE_TAG](bpf_skb_output) -->
[:octicons-tag-24: v5.5](https://github.com/torvalds/linux/commit/a7658e1a4164ce2b9eb4a11aadbba38586e93bd6)
<!-- [/FEATURE_TAG] -->

This helper writes a raw `data` blob into a special BPF perf event held by `map` of type [`BPF_MAP_TYPE_PERF_EVENT_ARRAY`](../map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md).

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


**Returns**
0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_skb_output)(void *ctx, void *map, __u64 flags, void *data, __u64 size) = (void *) 111;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

The perf event must have the following attributes: `PERF_SAMPLE_RAW` as `sample_type`, `PERF_TYPE_SOFTWARE` as `type`, and `PERF_COUNT_SW_BPF_OUTPUT` as `config`.

The `flags` are used to indicate the index in `map` for which the value must be put, masked with `BPF_F_INDEX_MASK`. Alternatively, `flags` can be set to `BPF_F_CURRENT_CPU` to indicate that the index of the current CPU core should be used.

The value to write, of `size`, is passed through eBPF stack and
pointed by `data`.

`ctx` is a pointer to in-kernel `struct sk_buff`.

This helper is similar to [`bpf_perf_event_output`](bpf_perf_event_output.md) but restricted to raw_tracepoint bpf programs.

### Program types

This helper call can be used in the following program types:

 * [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)

### Map types

This helper call can be used with the following map types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_MAP_REF] -->
 * [`BPF_MAP_TYPE_PERF_EVENT_ARRAY`](../map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md)
<!-- [/HELPER_FUNC_MAP_REF] -->

### Example

```c
#include "vmlinux.h"
#include <linux/version.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_core_read.h>

struct {
    __uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
    __uint(key_size, sizeof(int));
    __uint(value_size, sizeof(u32));
    __uint(max_entries, 2);
} my_map SEC(".maps");

SEC("raw_tp/sched_wakeup")
int BPF_PROG(handle_sched_wakeup, struct task_struct *p)
{
    struct S {
        u64 pid;
        u64 cookie;
    } data;

    data.pid = BPF_CORE_READ(p, pid);
    data.cookie = 0x12345678;

    bpf_skb_output(ctx, &my_map, 0, &data, sizeof(data));

    return 0;
}

char _license[] SEC("license") = "GPL";
u32 _version SEC("version") = LINUX_VERSION_CODE;
```
