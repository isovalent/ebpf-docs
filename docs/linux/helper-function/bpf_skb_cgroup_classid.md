# Helper function `bpf_skb_cgroup_classid`

<!-- [FEATURE_TAG](bpf_skb_cgroup_classid) -->
[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/b426ce83baa7dff947fb354118d3133f2953aac8)
<!-- [/FEATURE_TAG] -->

## Defintion

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
See **bpf_get_cgroup_classid**() for the main description. This helper differs from **bpf_get_cgroup_classid**() in that the cgroup v1 net_cls class is retrieved only from the _skb_'s associated socket instead of the current process.

### Returns

The id is returned or 0 in case the id could not be retrieved.

`#!c static __u64 (*bpf_skb_cgroup_classid)(struct __sk_buff *skb) = (void *) 151;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
