---
title: "Syscall command 'BPF_PROG_TEST_RUN' - eBPF Docs"
description: "This page documents the 'BPF_PROG_TEST_RUN' eBPF syscall command, including its defintion, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_PROG_TEST_RUN` command

<!-- [FEATURE_TAG](BPF_PROG_TEST_RUN) -->
[:octicons-tag-24: v4.12](https://github.com/torvalds/linux/commit/1cf1cae963c2e6032aebe1637e995bc2f5d330f4)
<!-- [/FEATURE_TAG] -->

This command runs a loaded eBPF program in the kernel one or multiple times with a supplied input and records the output. This can be used to test or benchmark a program.

## Return value

This command will return `0` on success or a error number (negative integer) if something went wrong.

## Usage

This command can be used to test or benchmark programs. Not all program types support this feature, only the following program types can be tested or benchmarked:

* [`BPF_PROG_TYPE_SOCK_OPS`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
* [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
* [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
* [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
* [`BPF_PROG_TYPE_SOCKET_FILTER`](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
* [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
* [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
* [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
* [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
* [`BPF_PROG_TYPE_LWT_IN`](../program-type/BPF_PROG_TYPE_LWT_IN.md)
* [`BPF_PROG_TYPE_LWT_OUT`](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
* [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
* [`BPF_PROG_TYPE_LWT_SEG6LOCAL`](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
* [`BPF_PROG_TYPE_FLOW_DISSECTOR`](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md)

Programs of [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md) are an exception, in the sense that this mechanism is the only way for the program to actually execute. Upon calling it can do further bpf syscalls for the purposes of acting a loader. For more details checkout its page.

!!! note
    The test framework for [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md) programs seems to execute some internal self tests and populate the `retval` with `1` or `2` depending on the outcome, but [does not actually execute the program](https://github.com/torvalds/linux/commit/da00d2f117a08fbca262db5ea422c80a568b112b) like you would expect to. This is as of kernel v6.2.

## Attributes

### `prog_fd`

[:octicons-tag-24: v4.12](https://github.com/torvalds/linux/commit/1cf1cae963c2e6032aebe1637e995bc2f5d330f4)

This field indicates the file descriptor for the program you would like to test.

### `retval`

[:octicons-tag-24: v4.12](https://github.com/torvalds/linux/commit/1cf1cae963c2e6032aebe1637e995bc2f5d330f4)

This field will be set to the return value, returned by the program after calling the command. The meaning of the return value depends on the program type.

### `data_size_in`

[:octicons-tag-24: v4.12](https://github.com/torvalds/linux/commit/1cf1cae963c2e6032aebe1637e995bc2f5d330f4)

This field indicates the size of the data passed into `data_in`

This field is not supported for the following program types:

* [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
* [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)

### `data_size_out`

[:octicons-tag-24: v4.12](https://github.com/torvalds/linux/commit/1cf1cae963c2e6032aebe1637e995bc2f5d330f4)

This field indicates the size of the buffer provided with `data_out`. If the size is smaller than the data outputted by the program, the syscall command will return a `-ENOSPC` value.

This field is not supported for the following program types:

* [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
* [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)

### `data_in`

[:octicons-tag-24: v4.12](https://github.com/torvalds/linux/commit/1cf1cae963c2e6032aebe1637e995bc2f5d330f4)

This field indicates the data to provide to the program. It should be a pointer to a memory buffer in the calling program. The format depends on the data the program expects, it is not the context but rater the data referred to by the context such as [`xdp_md->data`](../program-type/BPF_PROG_TYPE_XDP.md#data) or [`__sk_buff->data`](../program-context/__sk_buff.md#data).

This field is not supported for the following program types:

* [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
* [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
  
### `data_out`

[:octicons-tag-24: v4.12](https://github.com/torvalds/linux/commit/1cf1cae963c2e6032aebe1637e995bc2f5d330f4)

This field indicates the data after the program has potentially modified it. It should be a pointer to a memory buffer in the calling program.

This field is not supported for the following program types:

* [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
* [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)

### `repeat`

[:octicons-tag-24: v4.12](https://github.com/torvalds/linux/commit/1cf1cae963c2e6032aebe1637e995bc2f5d330f4)

This field indicates how often the program should be ran for a single syscall command. This is useful when benchmarking a program to increase the significance of the `duration` value.

This field is not supported for the following program types:

* [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
* [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)

### `duration`

[:octicons-tag-24: v4.12](https://github.com/torvalds/linux/commit/1cf1cae963c2e6032aebe1637e995bc2f5d330f4)

This field indicates how long the execution of the program took in nanoseconds. If `repeat` is larger than 1, this value should be divided by `repeat` to get the average per-invocation time.

This field is not supported for the following program types:

* [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
* [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)

### `ctx_size_in`

[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/b0b9395d865e3060d97658fbc9ba3f77fecc8da1)

This field indicates the size of the `ctx_in` buffer.

This field is not supported for programs of type [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md).

### `ctx_size_out`

[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/b0b9395d865e3060d97658fbc9ba3f77fecc8da1)

This field indicates the size of the `ctx_out` buffer. If this size is smaller than the actual context the kernel want to write back, the syscall command return an `-ENOSPEC` error code.

This field is not supported for the following program types:

* [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
* [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)

### `ctx_in`

[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/b0b9395d865e3060d97658fbc9ba3f77fecc8da1)

This field indicates the context with which the program should be invoked. This should be a pointer to a memory buffer. The structure of the memory depends on the context of the program type under test.

For programs of type [`BPF_PROG_TYPE_SOCK_OPS`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md) the context must be the an array of 64-bit integers of the size equal to the amount of arguments of the operation.

For programs of type [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md) the value of this field is available to to the program one to one.

For programs of type [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md) the value of this field is available to the program one to one, and should thus match the structure of the tracepoint.

For programs of type [`BPF_PROG_TYPE_FLOW_DISSECTOR`](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md) the value of this field should match `struct bpf_flow_keys`.

For programs of type [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md) the value of this field should match [`struct xdp_md`](../program-type/BPF_PROG_TYPE_XDP.md#context). The `egress_ifindex` field may not be set, the `ingress_ifindex` and `rx_queue_index` fields are mutually exclusive, and `data` and `data_end` are ignored and set automatically based on the supplied `data_in`.

For all other program types (which all are skb-based) the value of this field should match [`struct __sk_buff`](../program-context/__sk_buff.md). However not all fields are allowed and some are generated, the following rules apply:

* `mark` can be set
* `priority` can be set
* `ingress_ifindex` can be set
* `ifindex` can be set
* `cb` can be set
* `tstamp` can be set
* `wire_len` can be set
* `gso_segs` can be set
* `gso_size` can be set
* `hwtstamp` can be set
* `protocol` is inferred from the `data` passed in and can only be `ETH_P_IP` or `ETH_P_IPV6`.
* All other fields may not be set, and should be zero

### `ctx_out`

[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/b0b9395d865e3060d97658fbc9ba3f77fecc8da1)

This field indicates the context after the program has executed. The kernel will write the modified context back to the memory indicated by this field. This field should be a pointer to a memory buffer in the calling program.

This field is not supported for the following program types:

* [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
* [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)

### `flags`

[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/1b4d60ec162f82ea29a2e7a907b5c6cc9f926321)

This field contains flags to modify behavior of the test. More details in the [flags](#flags) section.

This field is not supported for the following program types:

* [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
* [`BPF_PROG_TYPE_FLOW_DISSECTOR`](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md)

### `cpu`

[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/1b4d60ec162f82ea29a2e7a907b5c6cc9f926321)

This field indicates the logical CPU on which the test program should be executed.

This field is only honoured when the `BPF_F_TEST_RUN_ON_CPU` flag is set.

This field can only be used with [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md) programs.

### `batch_size`

[:octicons-tag-24: v5.18](https://github.com/torvalds/linux/commit/b530e9e1063ed2b817eae7eec6ed2daa8be11608)

This field indicates the size of the batch of network packets to be be sent to the kernel at once.

This field only has meaning when the `BPF_F_TEST_XDP_LIVE_FRAMES` flag is set.

This field only works for [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md) programs.

## Flags

### `BPF_F_TEST_RUN_ON_CPU`

When set, this flag enable running the program on a specific CPU. This is necessary since the default value of `0` is a valid CPU index.

### `BPF_F_TEST_XDP_LIVE_FRAMES`

When set, the packet after being processed by the program are injected into the network stack as if having arrived. This can be used to inject packets into the kernel for a number of tests.
