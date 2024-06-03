---
title: "Helper Function 'bpf_ktime_get_tai_ns'"
description: "This page documents the 'bpf_ktime_get_tai_ns' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_ktime_get_tai_ns`

<!-- [FEATURE_TAG](bpf_ktime_get_tai_ns) -->
[:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/c8996c98f703b09afe77a1d247dae691c9849dc1)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
A nonsettable system-wide clock derived from wall-clock time but ignoring leap seconds.  This clock does not experience discontinuities and backwards jumps caused by NTP inserting leap seconds as CLOCK_REALTIME does.

See: **clock_gettime**(**CLOCK_TAI**)

### Returns

Current _ktime_.

`#!c static __u64 (* const bpf_ktime_get_tai_ns)(void) = (void *) 208;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->

<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_PERF_EVENT](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
