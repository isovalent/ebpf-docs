---
title: "Helper Function 'bpf_trace_printk'"
description: "This page documents the 'bpf_trace_printk' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_trace_printk`

<!-- [FEATURE_TAG](bpf_trace_printk) -->
[:octicons-tag-24: v4.1](https://github.com/torvalds/linux/commit/9c959c863f8217a2ff3d7c296e8223654d240569)
<!-- [/FEATURE_TAG] -->

This helper prints messages to the trace log of the kernel.

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


**Returns**
The number of bytes written to the buffer, or a negative error
in case of failure.

`#!c static long (* const bpf_trace_printk)(const char *fmt, __u32 fmt_size, ...) = (void *) 6;`

## Usage

This helper is a "printk()-like" facility for debugging. It prints a message defined by format `fmt` (of size `fmt_size`) to file `/sys/kernel/tracing/trace` from TraceFS, if available. It can take up to three additional `u64` arguments (as an eBPF helpers, the total number of arguments is limited to five).

!!! warning
    A commonly made mistake is to call `bpf_trace_printk` with a literal string like 
    
    ```c
    char *fmt = "some log";
    bpf_trace_printk(fmt, 9);
    ``` 
    
    The compiler will place such a literal string in a ELF section meant for the heap, which does not exist in eBPF programs. Rater the `fmt` should be defined as static const 
    
    ```c 
    static const char fmt[] = "some log"; 
    bpf_trace_printk(fmt, sizeof(fmt))
    ```
    
    which will result in the format being stack allocated. A good alternative would be to use the `bpf_printk` macro provided by the libbpf's [bpf_helpers.h](https://github.com/libbpf/libbpf/blob/master/src/bpf_helpers.h) file which does this step for you. `#!c bpf_printk("some log");`

### Trace output

Each time the helper is called, it appends a line to the trace. Lines are discarded while `/sys/kernel/tracing/trace` is open, use `/sys/kernel/tracing/trace_pipe` to avoid this. The format of the trace is customizable, and the exact output one will get depends on the options set in `/sys/kernel/tracing/trace_options` (see also the `README` file under the same directory). However, it usually defaults to something like:

`telnet-470   [001] .N.. 419421.045894: 0x00000001: <formatted msg>`

In the above:

`telnet` is the name of the current task. `470` is the PID of the current task. `001` is the CPU number on which the task is running. In `.N..`, each character refers to a set of options (whether irqs are enabled, scheduling options, whether hard/softirqs are running, level of preempt_disabled respectively). `N` means that `TIF_NEED_RESCHED` and `PREEMPT_NEED_RESCHED` are set. `419421.045894` is a timestamp. `0x00000001` is a fake value used by BPF for the instruction pointer register. `<formatted msg>` is the message formatted with `fmt`.

### String format

The conversion specifiers supported by `fmt` are similar, but more limited than for `printk()`. They are:

* `%d`, `%i` - Signed integer decimal
* `%u` - Unsigned integer decimal
* `%x` - Integer hexadecimal 
* `%ld`, `%li` - Signed long decimal
* `%lu` - Unsigned long decimal
* `%lx` - Long hexadecimal
* `%lld`, `%lli` - Long long decimal
* `%llu` - Unsigned long long decimal
* `%llx` - Long long hexadecimal
* `%p` - Pointer as decimal
* `%s` - String

No modifier (size of field, padding with zeroes, etc.) is available, and the helper will return `-EINVAL` (but print nothing) if it encounters an unknown specifier.

The above is true for pre [:octicons-tag-24: v5.13](https://github.com/torvalds/linux/commit/d9c9e4db186ab4d81f84e6f22b225d333b9424e3) kernels. Since then the following changes were made:

* `%%` - Produces a percent char `%`
* Modifiers such as `%03d` or `%+d` no longer cause `-EINVAL` but are ignored.
* Uppercase `%x` (`%X`, `%lX`, `%llX`) now prints the hexadecimal output but with uppercase chars instead.
* `%pK` - Kernel pointer which should be hidden from unprivileged
users
* `%px` - Pointer as hexadecimal
* `%pB` - Pointer to a symbol, prints the name with offsets and should be used when printing stack backtraces
* `%pi4`, `%pI4` - Pointer to an IPv4 address
* `%pi6`, `%pI6` - Pointer to an IPv6 address

!!! note 
    `bpf_trace_printk()` is slow, and should only be used for debugging purposes. For this reason, a notice block (spanning several lines) is printed to kernel logs and states that the helper should not be used "for production use" the first time this helper is used (or more precisely, when `trace_printk()` buffers are allocated). For passing values to user space, perf events should be preferred.

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_CGROUP_DEVICE`](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md)
 * [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCK`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [`BPF_PROG_TYPE_CGROUP_SYSCTL`](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
 * [`BPF_PROG_TYPE_FLOW_DISSECTOR`](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md)
 * [`BPF_PROG_TYPE_KPROBE`](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [`BPF_PROG_TYPE_LIRC_MODE2`](../program-type/BPF_PROG_TYPE_LIRC_MODE2.md)
 * [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
 * [`BPF_PROG_TYPE_LWT_IN`](../program-type/BPF_PROG_TYPE_LWT_IN.md)
 * [`BPF_PROG_TYPE_LWT_OUT`](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
 * [`BPF_PROG_TYPE_LWT_SEG6LOCAL`](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
 * [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [`BPF_PROG_TYPE_SK_LOOKUP`](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md)
 * [`BPF_PROG_TYPE_SK_MSG`](../program-type/BPF_PROG_TYPE_SK_MSG.md)
 * [`BPF_PROG_TYPE_SK_REUSEPORT`](../program-type/BPF_PROG_TYPE_SK_REUSEPORT.md)
 * [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md)
 * [`BPF_PROG_TYPE_SOCKET_FILTER`](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
 * [`BPF_PROG_TYPE_SOCK_OPS`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
 * [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
 * [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
 * [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

```c
static const char fmt[] = "some log"; 
bpf_trace_printk(fmt, sizeof(fmt));

// 
static const char fmt[] = "some int: %d"; 
bpf_trace_printk(fmt, sizeof(fmt), 123);

// 
static const char fmt[] = "big number: %lld"; 
long long abc = 123456789;
bpf_trace_printk(fmt, sizeof(fmt), abc);
```
