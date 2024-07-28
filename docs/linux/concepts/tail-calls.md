---
title: "Tail calls"
description: "This page explains the concept of tail calls in eBPF. It explains what tail calls are, how to use them, and when to use them."
---
# Tail calls

A tail call is a form mechanism that allows eBPF authors to break up their logic into multiple parts and go from one to the other. Unlike traditional function calls, control flow never returns to the code making a tail call, it works more like a `goto` statement.

To use tail calls, an author would add a [`BPF_MAP_TYPE_PROG_ARRAY`](../map-type/BPF_MAP_TYPE_PROG_ARRAY.md) map to their program. The map can be filled with references to other programs (given a few conditions). And the program can then use the [`bpf_tail_call`](../helper-function/bpf_tail_call.md) helper call with a reference to the map and an index to perform the actual tail call.

A popular use for tail calls is to spread "complexity" over several programs. Each target of a tail call is seen as a separate eBPF program, starting with a zero stack and only a context at R1. Thus each program has to pass the verifier independently and also gets its own complexity budget. Now we can make our program multiple times more complex by breaking it up into multiple pieces.

Another use case is for replacing or extending logic. By replacing the contents of the program array while it is in use. For example to update a program version without downtime or to enable/disable logic.

## Limitations

To prevent infinite loops or very long running programs, the kernel limits the amount of tail calls per initial invocation to `32` so `33` programs can execute in total before the tail call helper will refuse to jump anymore.

If a program array is associated with a program, any program added to the map should "match" the program. So they have to have the same `type`, `expected_attach_type`, `attached_btf`, etc.

While the same stack frame is shared, the verifier will block you from using any existing stack state without re-initializing it, the same goes for the registers. Thus, there is no straightforward way to shared state. Common workarounds for this issue are to use opaque fields in metadata such as [`__sk_buff->cb`](../program-context/__sk_buff.md#cb) or [`xdp_md->data_meta`](../program-type/BPF_PROG_TYPE_XDP.md#data_meta) memory. Alternatively, a per-CPU map with a single entry can be used to share data, which works since eBPF programs never migrate to a different CPU even between tail calls. However on RT (real time) kernels eBPF programs might be interrupted and re-started at a later time, so these maps should only be shared between tail calls on the same task, not globally.

When tail calls are combined with BPF-to-BPF function calls, the available stack size per program will shrink from `512` bytes to `256` bytes. This is to limit the stack allocation required by the kernel, as explained by the following comment from the kernel:

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
