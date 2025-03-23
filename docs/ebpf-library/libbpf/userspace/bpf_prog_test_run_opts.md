---
title: "Libbpf userspace function 'bpf_prog_test_run_opts'"
description: "This page documents the 'bpf_prog_test_run_opts' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_prog_test_run_opts`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_PROG_TEST_RUN`](../../../linux/syscall/BPF_PROG_TEST_RUN.md) syscall command.

## Definition

`#!c int bpf_prog_test_run_opts(int prog_fd, struct bpf_test_run_opts *opts);`

**Parameters**

- `prog_fd`: BPF program file descriptor
- `opts`: options for configuring the test run

### `struct bpf_test_run_opts`

```c
struct bpf_test_run_opts {
	size_t sz; /* size of this struct for forward/backward compatibility */
	const void *data_in; /* optional */
	void *data_out;      /* optional */
	__u32 data_size_in;
	__u32 data_size_out; /* in: max length of data_out
			      * out: length of data_out
			      */
	const void *ctx_in; /* optional */
	void *ctx_out;      /* optional */
	__u32 ctx_size_in;
	__u32 ctx_size_out; /* in: max length of ctx_out
			     * out: length of cxt_out
			     */
	__u32 retval;        /* out: return code of the BPF program */
	int repeat;
	__u32 duration;      /* out: average per repetition in ns */
	__u32 flags;
	__u32 cpu;
	__u32 batch_size;
};
```

## Usage

This function can be used to execute a loaded program for testing purposes or to perform some action in eBPF when the userspace application requests it.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
