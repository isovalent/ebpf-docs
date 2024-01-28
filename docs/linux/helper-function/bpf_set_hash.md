# Helper function `bpf_set_hash`

<!-- [FEATURE_TAG](bpf_set_hash) -->
[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/ded092cd73c2c56a394b936f86897f29b2e131c0)
<!-- [/FEATURE_TAG] -->

## Defintion

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Set the full hash for _skb_ (set the field _skb_**->hash**) to value _hash_.

### Returns

0

`#!c static long (*bpf_set_hash)(struct __sk_buff *skb, __u32 hash) = (void *) 48;`
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
