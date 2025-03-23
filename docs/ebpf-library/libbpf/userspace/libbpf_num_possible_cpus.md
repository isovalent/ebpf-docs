---
title: "Libbpf userspace function 'libbpf_num_possible_cpus'"
description: "This page documents the 'libbpf_num_possible_cpus' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `libbpf_num_possible_cpus`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.4](https://github.com/libbpf/libbpf/releases/tag/v0.0.4)
<!-- [/LIBBPF_TAG] -->

Is a helper function to get the number of possible CPUs that the host kernel supports and expects.

## Definition

`#!c int libbpf_num_possible_cpus(void);`

**Return**

Number of possible CPUs; or error code on failure.

## Usage

This function is useful when working with per-CPU maps, as it allows you to allocate the right amount of memory for the values array or to know the max index of maps like [`BPF_MAP_TYPE_PERF_EVENT_ARRAY`](../../../linux/map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md) which are index based on CPU index.

### Example

```c
int ncpus = libbpf_num_possible_cpus();
if (ncpus < 0) {
        // error handling
}
long values[ncpus];
bpf_map_lookup_elem(per_cpu_map_fd, key, values);
```
