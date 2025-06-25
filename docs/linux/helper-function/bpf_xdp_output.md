---
title: "Helper Function 'bpf_xdp_output'"
description: "This page documents the 'bpf_xdp_output' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_xdp_output`

<!-- [FEATURE_TAG](bpf_xdp_output) -->
[:octicons-tag-24: v5.7](https://github.com/torvalds/linux/commit/d831ee84bfc9173eecf30dbbc2553ae81b996c60)
<!-- [/FEATURE_TAG] -->

This helper writes a raw `data` blob into a special BPF perf event held by `map` of type [`BPF_MAP_TYPE_PERF_EVENT_ARRAY`](../map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md).

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


**Returns**
0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_xdp_output)(void *ctx, void *map, __u64 flags, void *data, __u64 size) = (void *) 121;`

## Usage

The perf event must have the following attributes: `PERF_SAMPLE_RAW` as `sample_type`, `PERF_TYPE_SOFTWARE` as `type`, and `PERF_COUNT_SW_BPF_OUTPUT` as `config`.

The `flags` are used to indicate the index in `map` for which the value must be put, masked with `BPF_F_INDEX_MASK`. Alternatively, `flags` can be set to `BPF_F_CURRENT_CPU` to indicate that the index of the current CPU core should be used.

The value to write, of `size`, is passed through eBPF stack and pointed by `data`.

`ctx` is a pointer to in-kernel struct xdp_buff.

This helper is similar to [`bpf_perf_event_output`](bpf_perf_event_output.md) but restricted to raw_tracepoint bpf programs.

### Program types

This helper call can be used in the following program types:

 * [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)

### Map types

This helper call can be used with the following map types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_MAP_REF] -->
 * [`BPF_MAP_TYPE_PERF_EVENT_ARRAY`](../map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md)
<!-- [/HELPER_FUNC_MAP_REF] -->

### Example

```
// SPDX-License-Identifier: GPL-2.0
#include <linux/bpf.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_helpers.h>

char _license[] SEC("license") = "GPL";

struct net_device {
	/* Structure does not need to contain all entries,
	 * as "preserve_access_index" will use BTF to fix this...
	 */
	int ifindex;
} __attribute__((preserve_access_index));

struct xdp_rxq_info {
	/* Structure does not need to contain all entries,
	 * as "preserve_access_index" will use BTF to fix this...
	 */
	struct net_device *dev;
	__u32 queue_index;
} __attribute__((preserve_access_index));

struct xdp_buff {
	void *data;
	void *data_end;
	void *data_meta;
	void *data_hard_start;
	unsigned long handle;
	struct xdp_rxq_info *rxq;
} __attribute__((preserve_access_index));

struct meta {
	int ifindex;
	int pkt_len;
};

struct {
	__uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
	__type(key, int);
	__type(value, int);
} perf_buf_map SEC(".maps");

__u64 test_result_fentry = 0;
SEC("fentry/FUNC")
int BPF_PROG(trace_on_entry, struct xdp_buff *xdp)
{
	struct meta meta;

	meta.ifindex = xdp->rxq->dev->ifindex;
	meta.pkt_len = bpf_xdp_get_buff_len((struct xdp_md *)xdp);
	bpf_xdp_output(xdp, &perf_buf_map,
		       ((__u64) meta.pkt_len << 32) |
		       BPF_F_CURRENT_CPU,
		       &meta, sizeof(meta));

	test_result_fentry = xdp->rxq->dev->ifindex;
	return 0;
}

__u64 test_result_fexit = 0;
SEC("fexit/FUNC")
int BPF_PROG(trace_on_exit, struct xdp_buff *xdp, int ret)
{
	test_result_fexit = ret;
	return 0;
}
```
