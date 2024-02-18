# KFunc `bpf_rdonly_cast`

<!-- [FEATURE_TAG](bpf_rdonly_cast) -->
[:octicons-tag-24: v6.2](https://github.com/torvalds/linux/commit/a35b9af4ec2c7f69286ef861fd2074a577e354cb)
<!-- [/FEATURE_TAG] -->

## Definition

This kfunc tries to cast the object to a specified type.

The function returns the same `obj` but with PTR_TO_BTF_ID with
btf_id. The verifier will ensure btf_id being a struct type.

Since the supported type cast may not reflect what the 'obj'
represents, the returned btf_id is marked as PTR_UNTRUSTED, so
the return value and subsequent pointer chasing cannot be
used as helper/kfunc arguments.

<!-- [KFUNC_DEF] -->
`#!c void *bpf_rdonly_cast(void *obj__ign, u32 btf_id__k)`
<!-- [/KFUNC_DEF] -->

## Usage

This tries to support use case like below:

`#!c #define skb_shinfo(SKB) ((struct skb_shared_info *)(skb_end_pointer(SKB)))`

where `skb_end_pointer(SKB)` is a `unsigned char *` and needs to
be casted to `struct skb_shared_info *`.


### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_CGROUP_SKB](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
- [BPF_PROG_TYPE_CGROUP_SOCK_ADDR](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
- [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
- [BPF_PROG_TYPE_LWT_IN](../program-type/BPF_PROG_TYPE_LWT_IN.md)
- [BPF_PROG_TYPE_LWT_OUT](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
- [BPF_PROG_TYPE_LWT_SEG6LOCAL](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
- [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
- [BPF_PROG_TYPE_NETFILTER](../program-type/BPF_PROG_TYPE_NETFILTER.md)
- [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
- [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [BPF_PROG_TYPE_SK_SKB](../program-type/BPF_PROG_TYPE_SK_SKB.md)
- [BPF_PROG_TYPE_SOCKET_FILTER](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
- [BPF_PROG_TYPE_STRUCT_OPS](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [BPF_PROG_TYPE_SYSCALL](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
- [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

