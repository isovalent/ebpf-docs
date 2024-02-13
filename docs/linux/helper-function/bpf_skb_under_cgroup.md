# Helper function `bpf_skb_under_cgroup`

<!-- [FEATURE_TAG](bpf_skb_under_cgroup) -->
[:octicons-tag-24: v4.8](https://github.com/torvalds/linux/commit/747ea55e4f78fd980350c39570a986b8c1c3e4aa)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Check whether _skb_ is a descendant of the cgroup2 held by _map_ of type **BPF_MAP_TYPE_CGROUP_ARRAY**, at _index_.

### Returns

The return value depends on the result of the test, and can be:

* 0, if the _skb_ failed the cgroup2 descendant test.
* 1, if the _skb_ succeeded the cgroup2 descendant test.
* A negative error code, if an error occurred.


`#!c static long (*bpf_skb_under_cgroup)(struct __sk_buff *skb, void *map, __u32 index) = (void *) 33;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_LWT_IN](../program-type/BPF_PROG_TYPE_LWT_IN.md)
 * [BPF_PROG_TYPE_LWT_OUT](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
 * [BPF_PROG_TYPE_LWT_SEG6LOCAL](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
 * [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
