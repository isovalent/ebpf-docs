# KFunc `bpf_iter_num_destroy`

<!-- [FEATURE_TAG](bpf_iter_num_destroy) -->
[:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/06accc8779c1d558a5b5a21f2ac82b0c95827ddd)
<!-- [/FEATURE_TAG] -->

This kfunc destroys the iterator `it`.

## Definition

This kfunc destroys the iterator `it`, the iterator can be reused, but as to be re-initialized with [`bpf_iter_num_new`](bpf_iter_num_new.md) before being used again.

<!-- [KFUNC_DEF] -->
`#!c void bpf_iter_num_destroy(struct bpf_iter_num *it)`
<!-- [/KFUNC_DEF] -->

## Usage

See [`bpf_iter_num_new`](bpf_iter_num_new.md#usage) for usage details.

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_CGROUP_SKB](../../program-types/BPF_PROG_TYPE_CGROUP_SKB.md)
- [BPF_PROG_TYPE_CGROUP_SOCK_ADDR](../../program-types/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
- [BPF_PROG_TYPE_LSM](../../program-types/BPF_PROG_TYPE_LSM.md)
- [BPF_PROG_TYPE_LWT_IN](../../program-types/BPF_PROG_TYPE_LWT_IN.md)
- [BPF_PROG_TYPE_LWT_OUT](../../program-types/BPF_PROG_TYPE_LWT_OUT.md)
- [BPF_PROG_TYPE_LWT_SEG6LOCAL](../../program-types/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
- [BPF_PROG_TYPE_LWT_XMIT](../../program-types/BPF_PROG_TYPE_LWT_XMIT.md)
- [BPF_PROG_TYPE_NETFILTER](../../program-types/BPF_PROG_TYPE_NETFILTER.md)
- [BPF_PROG_TYPE_SCHED_ACT](../../program-types/BPF_PROG_TYPE_SCHED_ACT.md)
- [BPF_PROG_TYPE_SCHED_CLS](../../program-types/BPF_PROG_TYPE_SCHED_CLS.md)
- [BPF_PROG_TYPE_SK_SKB](../../program-types/BPF_PROG_TYPE_SK_SKB.md)
- [BPF_PROG_TYPE_SOCKET_FILTER](../../program-types/BPF_PROG_TYPE_SOCKET_FILTER.md)
- [BPF_PROG_TYPE_STRUCT_OPS](../../program-types/BPF_PROG_TYPE_STRUCT_OPS.md)
- [BPF_PROG_TYPE_SYSCALL](../../program-types/BPF_PROG_TYPE_SYSCALL.md)
- [BPF_PROG_TYPE_TRACING](../../program-types/BPF_PROG_TYPE_TRACING.md)
- [BPF_PROG_TYPE_XDP](../../program-types/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

See [`bpf_iter_num_new`](bpf_iter_num_new.md#example) for examples.
