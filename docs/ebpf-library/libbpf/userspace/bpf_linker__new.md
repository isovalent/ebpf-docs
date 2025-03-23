---
title: "Libbpf userspace function 'bpf_linker__new'"
description: "This page documents the 'bpf_linker__new' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_linker__new`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.4.0](https://github.com/libbpf/libbpf/releases/tag/v0.4.0)
<!-- [/LIBBPF_TAG] -->

Create a new `struct bpf_linker` object from a file.

## Definition

`#!c struct bpf_linker *bpf_linker__new(const char *filename, struct bpf_linker_opts *opts);`

**Parameters**

- `filename`: The name of the first file to load the BPF object from. On [`bpf_linker__finalize`](bpf_linker__finalize.md), the linker will write the linked ELF file to this path.
- `opts`: A pointer to a `struct bpf_linker_opts` object that contains options for the linker.

### `struct bpf_linker_opts`

```c
struct bpf_linker_opts {
	/* size of this struct, for forward/backward compatibility */
	size_t sz;
};
```

## Usage

This function creates a new BPF linker. The linker is used to link statically link multiple BPF objects together into a single ELF file.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
