---
title: "Libbpf userspace function 'bpf_prog_linfo__lfind_addr_func'"
description: "This page documents the 'bpf_prog_linfo__lfind_addr_func' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_prog_linfo__lfind_addr_func`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Find a line info record for a given instruction offset and function index.

## Definition

`#!c const struct bpf_line_info * bpf_prog_linfo__lfind_addr_func(const struct bpf_prog_linfo *prog_linfo, __u64 addr, __u32 func_idx, __u32 nr_skip);`

**Parameters**

- `prog_linfo`: line info object to search in
- `addr`: instruction address to search for
- `func_idx`: function index to search for
- `nr_skip`: number of records to skip before returning the result

**Returns**

Pointer to the line info record for the given instruction offset and function index, or `NULL` if not found.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
