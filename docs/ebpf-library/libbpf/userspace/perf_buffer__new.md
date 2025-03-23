---
title: "Libbpf userspace function 'perf_buffer__new'"
description: "This page documents the 'perf_buffer__new' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `perf_buffer__new`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/releases/tag/v0.6.0)
<!-- [/LIBBPF_TAG] -->

Creates BPF perf buffer manager for a specified [`BPF_PERF_EVENT_ARRAY`](../../../linux/map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md) map.

## Definition

```c
typedef void (*perf_buffer_sample_fn)(void *ctx, int cpu, void *data, __u32 size);
typedef void (*perf_buffer_lost_fn)(void *ctx, int cpu, __u64 cnt);

struct perf_buffer * perf_buffer__new(int map_fd, size_t page_cnt, perf_buffer_sample_fn sample_cb, perf_buffer_lost_fn lost_cb, void *ctx, const struct perf_buffer_opts *opts);
```

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

### `struct perf_buffer_opts`

```c
struct perf_buffer_opts {
	size_t sz;
	__u32 sample_period;
	size_t :0;
};
```

#### `sample_period`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.2.0](https://github.com/libbpf/libbpf/commit/8c8243a4090dd6e8ad1f602c8e4b365cc1872620)
<!-- [/LIBBPF_TAG] -->

By default the perf buffer becomes signaled for every event that is being pushed to it. When set, this is the amount of events that need to be pushed to the perf buffer before it becomes signaled.

Processing events in batches can be more efficient than processing them one by one.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
