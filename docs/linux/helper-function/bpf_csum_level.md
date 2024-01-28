# Helper function `bpf_csum_level`

<!-- [FEATURE_TAG](bpf_csum_level) -->
[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/7cdec54f9713256bb170873a1fc5c75c9127c9d2)
<!-- [/FEATURE_TAG] -->

## Defintion

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Change the skbs checksum level by one layer up or down, or reset it entirely to none in order to have the stack perform checksum validation. The level is applicable to the following protocols: TCP, UDP, GRE, SCTP, FCOE. For example, a decap of | ETH | IP | UDP | GUE | IP | TCP | into | ETH | IP | TCP | through **bpf_skb_adjust_room**() helper with passing in

&nbsp;&nbsp;&nbsp;&nbsp;**BPF_F_ADJ_ROOM_NO_CSUM_RESET** flag would require onecall

to **bpf_csum_level**() with **BPF_CSUM_LEVEL_DEC** since the UDP header is removed. Similarly, an encap of the latter into the former could be accompanied by a helper call to **bpf_csum_level**() with **BPF_CSUM_LEVEL_INC** if the skb is still intended to be processed in higher layers of the stack instead of just egressing at tc.

There are three supported level settings at this time:

* **BPF_CSUM_LEVEL_INC**: Increases skb->csum_level for skbs
  with CHECKSUM_UNNECESSARY. * **BPF_CSUM_LEVEL_DEC**: Decreases skb->csum_level for skbs
  with CHECKSUM_UNNECESSARY. * **BPF_CSUM_LEVEL_RESET**: Resets skb->csum_level to 0 and
  sets CHECKSUM_NONE to force checksum validation by the stack. * **BPF_CSUM_LEVEL_QUERY**: No-op, returns the current
  skb->csum_level.

### Returns

0 on success, or a negative error in case of failure. In the case of **BPF_CSUM_LEVEL_QUERY**, the current skb->csum_level is returned or the error code -EACCES in case the skb is not subject to CHECKSUM_UNNECESSARY.

`#!c static long (*bpf_csum_level)(struct __sk_buff *skb, __u64 level) = (void *) 135;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
