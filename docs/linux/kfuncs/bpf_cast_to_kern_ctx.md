# KFunc `bpf_cast_to_kern_ctx`

<!-- [FEATURE_TAG](bpf_cast_to_kern_ctx) -->
[:octicons-tag-24: v6.2](https://github.com/torvalds/linux/commit/fd264ca020948a743e4c36731dfdecc4a812153c)
<!-- [/FEATURE_TAG] -->

## Definition

<!-- [KFUNC_DEF] -->
`#!c void *bpf_cast_to_kern_ctx(void *obj)`
<!-- [/KFUNC_DEF] -->

## Usage

The purpose of this kfunc is to cast the uAPI context programs get by default, into a kernel pointer
that is allowed to access the kernel type.

So for example an `BPF_PROG_TYPE_SCHED_CLS` program would get a `struct __sk_buff*` as the context, passing it to this kfunc would return a `struct sk_buff*` which is less stable but has more fields. CO-RE should be used to access the fields of the `struct sk_buff*` to ensure the program is compatible with different kernel versions.

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

