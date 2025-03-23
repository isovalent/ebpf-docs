---
title: "Libbpf userspace function 'bpf_prog_linfo__new'"
description: "This page documents the 'bpf_prog_linfo__new' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_prog_linfo__new`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Get the line info for a BPF program.

## Definition

`#!c struct bpf_prog_linfo * bpf_prog_linfo__new(const struct bpf_prog_info *info);`

**Parameters**

- `info`: BPF program info to get line info from

**Returns**

Pointer to the line info for the BPF program, or `NULL` on error.

### `struct bpf_prog_linfo`

```c
struct bpf_prog_linfo {
	void *raw_linfo;
	void *raw_jited_linfo;
	__u32 *nr_jited_linfo_per_func;
	__u32 *jited_linfo_func_idx;
	__u32 nr_linfo;
	__u32 nr_jited_func;
	__u32 rec_size;
	__u32 jited_rec_size;
};
```

#### `raw_linfo`

An array of line info records of type `struct bpf_line_info`. The instruction offset is that of the BPF program before loading.

```c
#define BPF_LINE_INFO_LINE_NUM(line_col)	((line_col) >> 10)
#define BPF_LINE_INFO_LINE_COL(line_col)	((line_col) & 0x3ff)

struct bpf_line_info {
	__u32	insn_off;
	__u32	file_name_off;
	__u32	line_off;
	__u32	line_col;
};
```

#### `raw_jited_linfo`

An array of line info records of type `struct bpf_line_info`. The instruction offset is that of the JITed BPF program.

#### `nr_jited_linfo_per_func`

An array of the number of `JIT`ed line info records per function.

#### `jited_linfo_func_idx`

An array of the function index for each `JIT`ed line info record.

#### `nr_linfo`

The number of line info records.

#### `nr_jited_func`

The number of `JIT`ed functions.

#### `rec_size`

The size of each line info record. (in case the structure changes in the future)

#### `jited_rec_size`

The size of each `JIT`ed line info record. (in case the structure changes in the future)

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
