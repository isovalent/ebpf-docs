# KFunc `bpf_dynptr_is_null`

<!-- [FEATURE_TAG](bpf_dynptr_is_null) -->
[:octicons-tag-24: v6.5](https://github.com/torvalds/linux/commit/540ccf96ddbc173474c32e595787d5622253be3d)
<!-- [/FEATURE_TAG] -->

Returns true if the dynptr is null / invalid.

## Definition

Returns true if the dynptr is null / invalid (determined by whether ptr->data is NULL), else false if
the dynptr is a valid dynptr.

<!-- [KFUNC_DEF] -->
`#!c bool bpf_dynptr_is_null(struct bpf_dynptr_kern *ptr)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

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

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

