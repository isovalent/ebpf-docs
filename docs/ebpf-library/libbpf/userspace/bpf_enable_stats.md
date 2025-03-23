---
title: "Libbpf userspace function 'bpf_enable_stats'"
description: "This page documents the 'bpf_enable_stats' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_enable_stats`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.9](https://github.com/libbpf/libbpf/releases/tag/v0.0.9)
<!-- [/LIBBPF_TAG] -->

Enable eBPF statistics collection system wide. A wrapper around the [`BPF_ENABLE_STATS`](../../../linux/syscall/BPF_ENABLE_STATS.md) syscall command.

## Definition

`#!c int bpf_enable_stats(enum bpf_stats_type type);`

**Parameters**

- `type`: type of statistics to enable

**Return**

`0`, on success; negative error code, otherwise

### `enum bpf_stats_type`

```c
enum bpf_stats_type {
	/* enabled run_time_ns and run_cnt */
	BPF_STATS_RUN_TIME = 0,
};
```

## Usage

This function enables eBPF statistics collection system wide. The statistics can be retrieved as part of the information gotten from [`bpf_prog_get_info_by_fd`](bpf_prog_get_info_by_fd.md).

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
