# Helper function `bpf_get_stackid`

<!-- [FEATURE_TAG](bpf_get_stackid) -->
[:octicons-tag-24: v4.6](https://github.com/torvalds/linux/commit/d5a3b1f691865be576c2bffa708549b8cdccda19)
<!-- [/FEATURE_TAG] -->

## Definition

<!-- [HELPER_FUNC_DEF] -->
Walk a user or a kernel stack and return its id. To achieve
this, the helper needs `ctx`, which is a pointer to the context
on which the tracing program is executed, and a pointer to a
`map` of type `BPF_MAP_TYPE_STACK_TRACE`.

The last argument, `flags`, holds the number of stack frames to
skip (from 0 to 255), masked with
`BPF_F_SKIP_FIELD_MASK`. The next bits can be used to set
a combination of the following flags:

`BPF_F_USER_STACK`
Collect a user space stack instead of a kernel stack.
`BPF_F_FAST_STACK_CMP`
Compare stacks by hash only.
`BPF_F_REUSE_STACKID`
If two different stacks hash into the same `stackid`,
discard the old one.

The stack id retrieved is a 32 bit long integer handle which
can be further combined with other data (including other stack
ids) and used as a key into maps. This can be useful for
generating a variety of graphs (such as flame graphs or off-cpu
graphs).

For walking a stack, this helper is an improvement over
`bpf_probe_read`\ (), which can be used with unrolled loops
but is not efficient and consumes a lot of eBPF instructions.
Instead, `bpf_get_stackid`\ () can collect up to
`PERF_MAX_STACK_DEPTH` both kernel and user frames. Note that
this limit can be controlled with the `sysctl` program, and
that it should be manually increased in order to profile long
user stacks (such as stacks for Java programs). To do so, use:

::

# sysctl kernel.perf_event_max_stack=<new value>


**Returns**
The positive or null stack id on success, or a negative error
in case of failure.

`#!c static long (*bpf_get_stackid)(void *ctx, void *map, __u64 flags) = (void *) 27;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_KPROBE](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [BPF_PROG_TYPE_PERF_EVENT](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [BPF_PROG_TYPE_TRACEPOINT](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
