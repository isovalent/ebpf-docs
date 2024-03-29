---
title: "Helper Function 'bpf_tail_call'"
description: "This page documents the 'bpf_tail_call' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_tail_call`

<!-- [FEATURE_TAG](bpf_tail_call) -->
[:octicons-tag-24: v4.2](https://github.com/torvalds/linux/commit/04fd61ab36ec065e194ab5e74ae34a5240d992bb)
<!-- [/FEATURE_TAG] -->

This special helper is used to trigger a "tail call", or in other words, to jump into another eBPF program.

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


Upon call of this helper, the program attempts to jump into a program referenced at index `index` in `prog_array_map`, a special map of type [`BPF_MAP_TYPE_PROG_ARRAY`](../program-type/BPF_MAP_TYPE_PROG_ARRAY.md), and passes
`ctx`, a pointer to the context.

**Returns**
0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_tail_call)(void *ctx, void *prog_array_map, __u32 index) = (void *) 12;`

## Usage

When jumping, The same stack frame is used (but values on stack and in registers for the caller are not accessible to the callee). This mechanism allows for program chaining, either for raising the maximum number of available eBPF instructions, or to execute given programs in conditional blocks. For security reasons, there is an upper limit to the number of successive tail calls that can be performed. This limit is defined in the kernel by the macro `MAX_TAIL_CALL_CNT` (not accessible to user space), which is currently set to 33.

If the call succeeds, the kernel immediately runs the first instruction of the new program. This is not a function call, and it never returns to the previous program. If the call fails, then the helper has no effect, and the caller continues to run its subsequent instructions. A call can fail if the destination program for the jump does not exist (i.e. `index` is superior to the number of entries in `prog_array_map`), or if the maximum number of tail calls has been reached for this chain of programs. 

!!! warning
    Using tail calls in combination with BPF-to-BPF function calls effects the maximum amount of stack memory your programs are allowed to use. Without tailcalls a total stack of 512 bytes is allowed, with tail-calls only a total stack size of 256 bytes is allowed. This quote from the verifier explains why:
    > protect against potential stack overflow that might happen when
	> bpf2bpf calls get combined with tailcalls. Limit the caller's stack
	> depth for such case down to 256 so that the worst case scenario
	> would result in 8k stack size (32 which is tailcall limit * 256 =
	> 8k).
	>
	> To get the idea what might happen, see an example:
    > ```
	> func1 -> sub rsp, 128
	>  subfunc1 -> sub rsp, 256
	>  tailcall1 -> add rsp, 256
	>   func2 -> sub rsp, 192 (total stack size = 128 + 192 = 320)
	>   subfunc2 -> sub rsp, 64
	>   subfunc22 -> sub rsp, 128
	>   tailcall2 -> add rsp, 128
	>    func3 -> sub rsp, 32 (total stack size 128 + 192 + 64 + 32 = 416)
	> ```
	> tailcall will unwind the current stack frame but it will not get rid
	> of caller's stack as shown on the example above.

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_DEVICE](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md)
 * [BPF_PROG_TYPE_CGROUP_SKB](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
 * [BPF_PROG_TYPE_CGROUP_SOCKOPT](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK_ADDR](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [BPF_PROG_TYPE_CGROUP_SYSCTL](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
 * [BPF_PROG_TYPE_FLOW_DISSECTOR](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md)
 * [BPF_PROG_TYPE_KPROBE](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [BPF_PROG_TYPE_LIRC_MODE2](../program-type/BPF_PROG_TYPE_LIRC_MODE2.md)
 * [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
 * [BPF_PROG_TYPE_LWT_IN](../program-type/BPF_PROG_TYPE_LWT_IN.md)
 * [BPF_PROG_TYPE_LWT_OUT](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
 * [BPF_PROG_TYPE_LWT_SEG6LOCAL](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
 * [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [BPF_PROG_TYPE_PERF_EVENT](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [BPF_PROG_TYPE_SK_LOOKUP](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md)
 * [BPF_PROG_TYPE_SK_MSG](../program-type/BPF_PROG_TYPE_SK_MSG.md)
 * [BPF_PROG_TYPE_SK_REUSEPORT](../program-type/BPF_PROG_TYPE_SK_REUSEPORT.md)
 * [BPF_PROG_TYPE_SK_SKB](../program-type/BPF_PROG_TYPE_SK_SKB.md)
 * [BPF_PROG_TYPE_SOCKET_FILTER](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
 * [BPF_PROG_TYPE_SOCK_OPS](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
 * [BPF_PROG_TYPE_SYSCALL](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [BPF_PROG_TYPE_TRACEPOINT](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
 * [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
