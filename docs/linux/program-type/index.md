---
title: Program types (Linux)
description: This page lists all program types that are available in the Linux kernel. They are categorized based on their functionality.
hide: toc
---
# Program types (Linux)

eBPF programs can be used for large and ever growing variety of different purposes. Different types of eBPF programs exist to accommodate these different use-cases. The Linux kernel may restrict or allow certain features depending on the program type, not all types of programs can do the same things because of where they are executed in the kernel. The verifier will enforce such restrictions.

## Network program types

These program types are triggered by network events

* [`BPF_PROG_TYPE_SOCKET_FILTER`](BPF_PROG_TYPE_SOCKET_FILTER.md)
* [`BPF_PROG_TYPE_SCHED_CLS`](BPF_PROG_TYPE_SCHED_CLS.md)
* [`BPF_PROG_TYPE_SCHED_ACT`](BPF_PROG_TYPE_SCHED_ACT.md)
* [`BPF_PROG_TYPE_XDP`](BPF_PROG_TYPE_XDP.md)
* [`BPF_PROG_TYPE_SOCK_OPS`](BPF_PROG_TYPE_SOCK_OPS.md)
* [`BPF_PROG_TYPE_SK_SKB`](BPF_PROG_TYPE_SK_SKB.md)
* [`BPF_PROG_TYPE_SK_MSG`](BPF_PROG_TYPE_SK_MSG.md)
* [`BPF_PROG_TYPE_SK_LOOKUP`](BPF_PROG_TYPE_SK_LOOKUP.md)
* [`BPF_PROG_TYPE_SK_REUSEPORT`](BPF_PROG_TYPE_SK_REUSEPORT.md)
* [`BPF_PROG_TYPE_FLOW_DISSECTOR`](BPF_PROG_TYPE_FLOW_DISSECTOR.md)
* [`BPF_PROG_TYPE_NETFILTER`](BPF_PROG_TYPE_NETFILTER.md)

### Light weight tunnel program types

These program types are used to implement custom light weight tunneling protocols

* [`BPF_PROG_TYPE_LWT_IN`](BPF_PROG_TYPE_LWT_IN.md)
* [`BPF_PROG_TYPE_LWT_OUT`](BPF_PROG_TYPE_LWT_OUT.md)
* [`BPF_PROG_TYPE_LWT_XMIT`](BPF_PROG_TYPE_LWT_XMIT.md)
* `BPF_PROG_TYPE_LWT_SEG6LOCAL`

## cGroup program types

These program types are triggered by events of cGroups to which the program is attached

* [`BPF_PROG_TYPE_CGROUP_SKB`](BPF_PROG_TYPE_CGROUP_SKB.md)
* [`BPF_PROG_TYPE_CGROUP_SOCK`](BPF_PROG_TYPE_CGROUP_SOCK.md)
* [`BPF_PROG_TYPE_CGROUP_DEVICE`](BPF_PROG_TYPE_CGROUP_DEVICE.md)
* [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
* [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](BPF_PROG_TYPE_CGROUP_SOCKOPT.md)
* [`BPF_PROG_TYPE_CGROUP_SYSCTL`](BPF_PROG_TYPE_CGROUP_SYSCTL.md)

## Tracing program types

These program types are triggered by tracing events from the kernel or userspace

* [`BPF_PROG_TYPE_KPROBE`](BPF_PROG_TYPE_KPROBE.md)
* [`BPF_PROG_TYPE_TRACEPOINT`](BPF_PROG_TYPE_TRACEPOINT.md)
* [`BPF_PROG_TYPE_PERF_EVENT`](BPF_PROG_TYPE_PERF_EVENT.md)
* [`BPF_PROG_TYPE_RAW_TRACEPOINT`](BPF_PROG_TYPE_RAW_TRACEPOINT.md)
* `BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE`
* [`BPF_PROG_TYPE_TRACING`](BPF_PROG_TYPE_TRACING.md)

## Misc

These program types have unique purposes and do not fit neatly in any of the larger categories

* [`BPF_PROG_TYPE_LIRC_MODE2`](BPF_PROG_TYPE_LIRC_MODE2.md)
* [`BPF_PROG_TYPE_LSM`](BPF_PROG_TYPE_LSM.md)
* [`BPF_PROG_TYPE_EXT`](BPF_PROG_TYPE_EXT.md)
* [`BPF_PROG_TYPE_STRUCT_OPS`](BPF_PROG_TYPE_STRUCT_OPS.md)
* [`BPF_PROG_TYPE_SYSCALL`](BPF_PROG_TYPE_SYSCALL.md) 

## ELF sections

The concept of a program type only exists at the kernel/syscall level. There is no standardized way of marking which program type a particular program within an [ELF](../../concepts/elf.md) is. The industry standard that most [loaders](../../concepts/loader.md) follow the example set out by Libbpf which is to patterns in the [ELF](../../concepts/elf.md) section names to convey the program type.

Section names supported by libbpf consist of one or more parts, separated by '/'. The first part identifies the program type of the program contained in the section. Subsequent parts (called `extras` in libppf documentation) may specify the [attach type](../syscall/BPF_LINK_CREATE.md#attach-types) if applicable, or the specific event to attach to. Extras, if present, provide details of how to auto-attach the program.

## Index of section names

| Program Type | Attach Type | ELF Section Name |
| --- | --- | --- |
| `BPF_PROG_TYPE_CGROUP_DEVICE` | `BPF_CGROUP_DEVICE` | `cgroup/dev` |
| `BPF_PROG_TYPE_CGROUP_SKB` || `cgroup/skb` |
| `BPF_PROG_TYPE_CGROUP_SKB` | `BPF_CGROUP_INET_EGRESS` | `cgroup_skb/egress` |
| `BPF_PROG_TYPE_CGROUP_SKB` | `BPF_CGROUP_INET_INGRESS` | `cgroup_skb/ingress` |
| `BPF_PROG_TYPE_CGROUP_SOCKOPT` | `BPF_CGROUP_GETSOCKOPT` | `cgroup/getsockopt` |
| `BPF_PROG_TYPE_CGROUP_SOCKOPT` | `BPF_CGROUP_SETSOCKOPT` | `cgroup/setsockopt` |
| `BPF_PROG_TYPE_CGROUP_SOCK_ADDR` | `BPF_CGROUP_INET4_BIND` | `cgroup/bind4` |
| `BPF_PROG_TYPE_CGROUP_SOCK_ADDR` | `BPF_CGROUP_INET4_CONNECT` | `cgroup/connect4` |
| `BPF_PROG_TYPE_CGROUP_SOCK_ADDR` | `BPF_CGROUP_INET4_GETPEERNAME` | `cgroup/getpeername4` |
| `BPF_PROG_TYPE_CGROUP_SOCK_ADDR` | `BPF_CGROUP_INET4_GETSOCKNAME` | `cgroup/getsockname4` |
| `BPF_PROG_TYPE_CGROUP_SOCK_ADDR` | `BPF_CGROUP_INET6_BIND` | `cgroup/bind6` |
| `BPF_PROG_TYPE_CGROUP_SOCK_ADDR` | `BPF_CGROUP_INET6_CONNECT` | `cgroup/connect6` |
| `BPF_PROG_TYPE_CGROUP_SOCK_ADDR` | `BPF_CGROUP_INET6_GETPEERNAME` | `cgroup/getpeername6` |
| `BPF_PROG_TYPE_CGROUP_SOCK_ADDR` | `BPF_CGROUP_INET6_GETSOCKNAME` | `cgroup/getsockname6` |
| `BPF_PROG_TYPE_CGROUP_SOCK_ADDR` | `BPF_CGROUP_UDP4_RECVMSG` | `cgroup/recvmsg4` |
| `BPF_PROG_TYPE_CGROUP_SOCK_ADDR` | `BPF_CGROUP_UDP4_SENDMSG` | `cgroup/sendmsg4` |
| `BPF_PROG_TYPE_CGROUP_SOCK_ADDR` | `BPF_CGROUP_UDP6_RECVMSG` | `cgroup/recvmsg6` |
| `BPF_PROG_TYPE_CGROUP_SOCK_ADDR` | `BPF_CGROUP_UDP6_SENDMSG` | `cgroup/sendmsg6` |
| `BPF_PROG_TYPE_CGROUP_SOCK_ADDR` | `BPF_CGROUP_UNIX_CONNECT` | `cgroup/connect_unix` |
| `BPF_PROG_TYPE_CGROUP_SOCK_ADDR` | `BPF_CGROUP_UNIX_SENDMSG` | `cgroup/sendmsg_unix` |
| `BPF_PROG_TYPE_CGROUP_SOCK_ADDR` | `BPF_CGROUP_UNIX_RECVMSG` | `cgroup/recvmsg_unix` |
| `BPF_PROG_TYPE_CGROUP_SOCK_ADDR` | `BPF_CGROUP_UNIX_GETPEERNAME` | `cgroup/getpeername_unix` |
| `BPF_PROG_TYPE_CGROUP_SOCK_ADDR` | `BPF_CGROUP_UNIX_GETSOCKNAME` | `cgroup/getsockname_unix` |
| `BPF_PROG_TYPE_CGROUP_SOCK` | `BPF_CGROUP_INET4_POST_BIND` | `cgroup/post_bind4` |
| `BPF_PROG_TYPE_CGROUP_SOCK` | `BPF_CGROUP_INET6_POST_BIND` | `cgroup/post_bind6` |
| `BPF_PROG_TYPE_CGROUP_SOCK` | `BPF_CGROUP_INET_SOCK_CREATE` | `cgroup/sock_create` |
| `BPF_PROG_TYPE_CGROUP_SOCK` | `BPF_CGROUP_INET_SOCK_CREATE` | `cgroup/sock` |
| `BPF_PROG_TYPE_CGROUP_SOCK` | `BPF_CGROUP_INET_SOCK_RELEASE` | `cgroup/sock_release` |
| `BPF_PROG_TYPE_CGROUP_SYSCTL` | `BPF_CGROUP_SYSCTL` | `cgroup/sysctl` |
| `BPF_PROG_TYPE_EXT` || `freplace`  or `freplace/<function>`  [^1]|
| `BPF_PROG_TYPE_FLOW_DISSECTOR` | `BPF_FLOW_DISSECTOR` | `flow_dissector` |
| `BPF_PROG_TYPE_KPROBE` || `kprobe`  or `kprobe/<function>` or `kprobe/<function>+<offset>`  [^2]|
| `BPF_PROG_TYPE_KPROBE` || `kretprobe`  or `kprobe/<function>` or `kprobe/<function>+<offset>`  [^2]|
| `BPF_PROG_TYPE_KPROBE` || `ksyscall`  or `ksyscall/<syscall>`  [^3]|
| `BPF_PROG_TYPE_KPROBE` || `kretsyscall`  or `ksyscall/<syscall>`  [^3]|
| `BPF_PROG_TYPE_KPROBE` || `uprobe`  or `uprobe/<path>:<function>` or `uprobe:/<path>:<function>+<offset>`  [^4]|
| `BPF_PROG_TYPE_KPROBE` || `uprobe.s`  or `uprobe.s/<path>:<function>` or `uprobe.s:/<path>:<function>+<offset>`  [^4]|
| `BPF_PROG_TYPE_KPROBE` || `uretprobe`  or `uretprobe/<path>:<function>` or `uretprobe:/<path>:<function>+<offset>`  [^4]|
| `BPF_PROG_TYPE_KPROBE` || `uretprobe.s`  or `uretprobe.s/<path>:<function>` or `uretprobe.s:/<path>:<function>+<offset>`  [^4]|
| `BPF_PROG_TYPE_KPROBE` || `usdt`  or `usdt/<path>:<provider>:<name>`  [^5]|
| `BPF_PROG_TYPE_KPROBE` | `BPF_TRACE_KPROBE_MULTI` | `kprobe.multi`  or `kprobe.multi/<pattern>`  [^6]|
| `BPF_PROG_TYPE_KPROBE` | `BPF_TRACE_KPROBE_MULTI` | `kretprobe.multi`  or  `kretprobe.multi/<pattern>`  [^6]|
| `BPF_PROG_TYPE_LIRC_MODE2` | `BPF_LIRC_MODE2` | `lirc_mode2` |
| `BPF_PROG_TYPE_LSM` | `BPF_LSM_CGROUP` | `lsm_cgroup` |     |
| `BPF_PROG_TYPE_LSM` | `BPF_LSM_MAC` | `lsm`  or `lsm/<hook>`  [^7]|
| `BPF_PROG_TYPE_LSM` | `BPF_LSM_MAC` | `lsm.s`  or `lsm.s/<hook>`  [^7]|
| `BPF_PROG_TYPE_LWT_IN` || `lwt_in` |
| `BPF_PROG_TYPE_LWT_OUT` || `lwt_out` |
| `BPF_PROG_TYPE_LWT_SEG6LOCAL` || `lwt_seg6local` |
| `BPF_PROG_TYPE_LWT_XMIT` || `lwt_xmit` |
| `BPF_PROG_TYPE_NETFILTER` || `netfilter` |
| `BPF_PROG_TYPE_PERF_EVENT` || `perf_event` |
| `BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE` || `raw_tp.w`  or `raw_tp.w/<tracepoint>`  [^8]|
| `BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE` || `raw_tracepoint.w` or `raw_tracepoint.w/<tracepoint>` |
| `BPF_PROG_TYPE_RAW_TRACEPOINT` || `raw_tp`  or `raw_tp.w/<tracepoint>`  [^8]|
| `BPF_PROG_TYPE_RAW_TRACEPOINT` || `raw_tracepoint` or `raw_tracepoint/<tracepoint>` |
| `BPF_PROG_TYPE_SCHED_ACT` || `action`   [^9]|
| `BPF_PROG_TYPE_SCHED_CLS` || `classifier`   [^9]|
| `BPF_PROG_TYPE_SCHED_CLS` || `tc`   [^9]|
| `BPF_PROG_TYPE_SCHED_CLS` | `BPF_NETKIT_PRIMARY` | `netkit/primary` |
| `BPF_PROG_TYPE_SCHED_CLS` | `BPF_NETKIT_PEER` | `netkit/peer` |
| `BPF_PROG_TYPE_SCHED_CLS` | `BPF_TCX_INGRESS` | `tc/ingress` |
| `BPF_PROG_TYPE_SCHED_CLS` | `BPF_TCX_EGRESS` | `tc/egress` |
| `BPF_PROG_TYPE_SCHED_CLS` | `BPF_TCX_INGRESS` | `tcx/ingress` |
| `BPF_PROG_TYPE_SCHED_CLS` | `BPF_TCX_EGRESS` | `tcx/egress` |
| `BPF_PROG_TYPE_SK_LOOKUP` | `BPF_SK_LOOKUP` | `sk_lookup` |
| `BPF_PROG_TYPE_SK_MSG` | `BPF_SK_MSG_VERDICT` | `sk_msg` |
| `BPF_PROG_TYPE_SK_REUSEPORT` | `BPF_SK_REUSEPORT_SELECT_OR_MIGRATE` | `sk_reuseport/migrate` |
| `BPF_PROG_TYPE_SK_REUSEPORT` | `BPF_SK_REUSEPORT_SELECT` | `sk_reuseport` |
| `BPF_PROG_TYPE_SK_SKB` || `sk_skb` |
| `BPF_PROG_TYPE_SK_SKB` | `BPF_SK_SKB_STREAM_PARSER` | `sk_skb/stream_parser` |
| `BPF_PROG_TYPE_SK_SKB` | `BPF_SK_SKB_STREAM_VERDICT` | `sk_skb/stream_verdict` |
| `BPF_PROG_TYPE_SOCKET_FILTER` || `socket` |
| `BPF_PROG_TYPE_SOCK_OPS` | `BPF_CGROUP_SOCK_OPS` | `sockops` |
| `BPF_PROG_TYPE_STRUCT_OPS` || `struct_ops`  or `struct_ops/<name>`  [^10]|
| `BPF_PROG_TYPE_STRUCT_OPS` || `struct_ops.s`  or `struct_ops.s/<name>`  [^10]|
| `BPF_PROG_TYPE_SYSCALL` || `syscall` |
| `BPF_PROG_TYPE_TRACEPOINT` || `tp`  or `tp/<category>/<name>`  [^11]|
| `BPF_PROG_TYPE_TRACEPOINT` || `tracepoint`  or `tracepoint/<category>/<name>`  [^11]|
| `BPF_PROG_TYPE_TRACING` | `BPF_MODIFY_RETURN` | `fmod_ret`  or `fmod_ret/<function>`  [^1]|
| `BPF_PROG_TYPE_TRACING` | `BPF_MODIFY_RETURN` | `fmod_ret.s`  or `fmod_ret.s/<function>`  [^1]|
| `BPF_PROG_TYPE_TRACING` | `BPF_TRACE_FENTRY` | `fentry`  or `fentry/<function>`  [^1]|
| `BPF_PROG_TYPE_TRACING` | `BPF_TRACE_FENTRY` | `fentry.s`  or `fentry.s/<function>`  [^1]|
| `BPF_PROG_TYPE_TRACING` | `BPF_TRACE_FEXIT` | `fexit`  or `fexit/<function>`  [^1]|
| `BPF_PROG_TYPE_TRACING` | `BPF_TRACE_FEXIT` | `fexit.s`  or `fexit.s/<function>`  [^1]|
| `BPF_PROG_TYPE_TRACING` | `BPF_TRACE_ITER` | `iter`  or ` iter/<struct-name>`  [^12]|
| `BPF_PROG_TYPE_TRACING` | `BPF_TRACE_ITER` | `iter.s`  or ` iter.s/<struct-name>`  [^12]|
| `BPF_PROG_TYPE_TRACING` | `BPF_TRACE_RAW_TP` | `tp_btf`  or `tp_btf/<function>`  [^1]|
| `BPF_PROG_TYPE_XDP` | `BPF_XDP_CPUMAP` | `xdp.frags/cpumap` |
| `BPF_PROG_TYPE_XDP` | `BPF_XDP_CPUMAP` | `xdp/cpumap` |
| `BPF_PROG_TYPE_XDP` | `BPF_XDP_DEVMAP` | `xdp.frags/devmap` |
| `BPF_PROG_TYPE_XDP` | `BPF_XDP_DEVMAP` | `xdp/devmap` |
| `BPF_PROG_TYPE_XDP` | `BPF_XDP` | `xdp.frags` |
| `BPF_PROG_TYPE_XDP` | `BPF_XDP` | `xdp` |

The table above was sourced from the [Program Types and ELF Sections](https://docs.kernel.org/bpf/libbpf/program_types.html) page (`Copyright (c) 2022 Donald Hunter . All rights reserved.`) in the [Linux Kernel documentation](https://docs.kernel.org/index.html).

[^1]: `<function>` is the symbol name of a function. This may be architecture-specific, such as `__x64_sys_getpid` for the `getpid` syscall on the x86_64 architecture. Valid characters for `<function>` are `a-zA-Z0-9_`.
[^2]: `<offset>` is an address offset from the symbol name. It must be a valid non-negative integer.
[^3]: `<syscall>` is the name of a system call, such as `getpid`. It is not architecture-specific.
[^4]:  `<path>` is a path to an executable or library.
[^5]:  `<path>` is a path to an executable or library that provides the USDT probe, `<provider>` is the USDT provider, and `<name>` is the USDT probe name.
[^6]:  `<pattern>` is used to match kernel function names, which may be architecture-specific. `<pattern>` supports `*` and `?` wildcards. Valid characters for `<pattern>` are `a-zA-Z0-9_.*?`.
[^7]:  `<hook>` is the name of an LSM (Linux Security Module) hook. See [Program type BPF_PROG_TYPE_LSM](./BPF_PROG_TYPE_LSM.md) for details.
[^8]:  `<tracepoint>` is the name of a trace event. See [Program type BPF_PROG_TYPE_TRACEPOINT](./BPF_PROG_TYPE_TRACEPOINT.md) and [Program Type BPF_PROG_TYPE_RAW_TRACEPOINT](./BPF_PROG_TYPE_RAW_TRACEPOINT.md) for details.
[^9]: The `tc`, `classifier` and `action` attach types are deprecated, use `tcx/*` instead.
[^10]: `<name>` is the value of the `.name` member of a struct defined in the `.struct_ops` section. See [Program type BPF_PROG_TYPE_STRUCT_OPS](./BPF_PROG_TYPE_STRUCT_OPS.md) for details.
[^11]: `<category>` is the name of a subsystem, and `<name>` is the name of an event as per [event tracing](https://docs.kernel.org/trace/events.html) convention.
[^12]: `<struct_name>` is the name of a _tracing program iterator_. See [Iterator in Program type BPF_PROG_TYPE_TRACING](./BPF_PROG_TYPE_TRACING.md#iterator) for details.
