---
title: "Libbpf userspace function 'perf_buffer__new_raw'"
description: "This page documents the 'perf_buffer__new_raw' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `perf_buffer__new_raw`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/releases/tag/v0.6.0)
<!-- [/LIBBPF_TAG] -->

Same as [`perf_buffer__new`](perf_buffer__new.md) but with more control over parameters which [`perf_buffer__new`](perf_buffer__new.md) auto-detects.

## Definition

`#!c struct perf_buffer * perf_buffer__new_raw(int map_fd, size_t page_cnt, struct perf_event_attr *attr, perf_buffer_event_fn event_cb, void *ctx, const struct perf_buffer_raw_opts *opts);`

**Parameters**

- `map_fd`: file descriptor of [`BPF_PERF_EVENT_ARRAY`](../../../linux/map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md) BPF map that will be used by BPF
code to send data over to user-space
- `page_cnt`: number of memory pages allocated for each per-CPU buffer
- `sample_cb`: function called on each received data record
- `lost_cb`: function called when record loss has occurred
- `ctx`: user-provided extra context passed into `sample_cb` and `lost_cb`
- `opts`: options for the perf buffer

**Return**

A new instance of `struct perf_buffer` on success, `NULL` on error with [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) containing an error code.

### `struct perf_buffer_raw_opts`

```c
struct perf_buffer_raw_opts {
	size_t sz;
	long :0;
	long :0;
	/* if cpu_cnt == 0, open all on all possible CPUs (up to the number of
	 * max_entries of given PERF_EVENT_ARRAY map)
	 */
	int cpu_cnt;
	/* if cpu_cnt > 0, cpus is an array of CPUs to open ring buffers on */
	int *cpus;
	/* if cpu_cnt > 0, map_keys specify map keys to set per-CPU FDs for */
	int *map_keys;
};
```

#### `cpu_cnt`

If `cpu_cnt == 0`, open all on all possible CPUs (up to the number of [`max_entries`](../../../linux/syscall/BPF_MAP_CREATE.md#max_entries) of given [`BPF_MAP_TYPE_PERF_EVENT_ARRAY`](../../../linux/map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md) map).

#### `cpus`

if `cpu_cnt > 0`, cpus is an array of CPUs to open ring buffers on.

#### `map_keys`

if `cpu_cnt > 0`, `map_keys` specify map keys to set per-CPU FDs for.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
