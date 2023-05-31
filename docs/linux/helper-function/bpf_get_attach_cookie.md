# Helper function `bpf_get_attach_cookie`

<!-- [FEATURE_TAG](bpf_get_attach_cookie) -->
[:octicons-tag-24: v5.15](https://github.com/torvalds/linux/commit/82e6b1eee6a8875ef4eacfd60711cce6965c6b04)
<!-- [/FEATURE_TAG] -->

## Definition

<!-- [HELPER_FUNC_DEF] -->
Get bpf_cookie value provided (optionally) during the program attachment. It might be different for each individual attachment, even if BPF program itself is the same. Expects BPF program context _ctx_ as a first argument.

Supported for the following program types:

&nbsp;&nbsp;&nbsp;&nbsp;- kprobe/uprobe;
&nbsp;- tracepoint;
&nbsp;- perf_event.


### Returns

Value specified by user at BPF link creation/attachment time or 0, if it was not specified.

`#!c static __u64 (*bpf_get_attach_cookie)(void *ctx) = (void *) 174;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

This is useful for cases when the same BPF program is used for attaching and processing invocation of different tracepoints/kprobes/uprobes in a generic fashion, but such that each invocation is distinguished from each other (e.g., BPF program can look up additional information associated with a specific kernel function without having to rely on function IP lookups). This enables new use cases to be implemented simply and efficiently that previously were possible only through code generation (and thus multiple instances of almost identical BPF program) or compilation at runtime (BCC-style) on target hosts (even more expensive resource-wise). For uprobes it is not even possible in some cases to know function IP before hand (e.g., when attaching to shared library without PID filtering, in which case base load address is not known for a library).

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_KPROBE](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [BPF_PROG_TYPE_PERF_EVENT](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [BPF_PROG_TYPE_TRACEPOINT](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
