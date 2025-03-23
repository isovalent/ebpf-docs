---
title: "Libbpf userspace function 'libbpf_probe_bpf_helper'"
description: "This page documents the 'libbpf_probe_bpf_helper' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `libbpf_probe_bpf_helper`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.7.0](https://github.com/libbpf/libbpf/releases/tag/v0.7.0)
<!-- [/LIBBPF_TAG] -->

Detects if host kernel supports the use of a given BPF helper from specified BPF program type.

## Definition

`#!c int libbpf_probe_bpf_helper(enum bpf_prog_type prog_type, enum bpf_func_id helper_id, const void *opts);`

**Parameters**

- `prog_type`: BPF program type used to check the support of BPF helper
- `helper_id`: BPF helper ID (enum bpf_func_id) to check support for
- `opts`: reserved for future extensibility, should be `NULL`

**Return**

`1`, if given combination of program type and helper is supported; `0`, if the combination is not supported; negative error code if feature detection for provided input arguments failed or can't be performed

## Usage

Make sure the process has required set of `CAP_*` permissions (or runs as root) when performing feature checking.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
