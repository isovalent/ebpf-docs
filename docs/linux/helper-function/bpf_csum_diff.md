# Helper function `bpf_csum_diff`

<!-- [FEATURE_TAG](bpf_csum_diff) -->
[:octicons-tag-24: v4.6](https://github.com/torvalds/linux/commit/7d672345ed295b1356a5d9f7111da1d1d7d65867)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Compute a checksum difference, from the raw buffer pointed by _from_, of length _from_size_ (that must be a multiple of 4), towards the raw buffer pointed by _to_, of size _to_size_ (same remark). An optional _seed_ can be added to the value (this can be cascaded, the seed may come from a previous call to the helper).

This is flexible enough to be used in several ways:

* With _from_size_ == 0, _to_size_ > 0 and _seed_ set to
  checksum, it can be used when pushing new data. * With _from_size_ > 0, _to_size_ == 0 and _seed_ set to
  checksum, it can be used when removing data from a packet. * With _from_size_ > 0, _to_size_ > 0 and _seed_ set to 0, it
  can be used to compute a diff. Note that _from_size_ and   _to_size_ do not need to be equal.

This helper can be used in combination with **bpf_l3_csum_replace**() and **bpf_l4_csum_replace**(), to which one can feed in the difference computed with **bpf_csum_diff**().

### Returns

The checksum result, or a negative error code in case of failure.

`#!c static __s64 (*bpf_csum_diff)(__be32 *from, __u32 from_size, __be32 *to, __u32 to_size, __wsum seed) = (void *) 28;`
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
 * [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
