---
title: "Libbpf userspace function 'bpf_prog_load'"
description: "This page documents the 'bpf_prog_load' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_prog_load`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/releases/tag/v0.6.0)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_PROG_LOAD`](../../../linux/syscall/BPF_PROG_LOAD.md) syscall command.

## Definition

`#!c int bpf_prog_load(enum bpf_prog_type prog_type, const char *prog_name, const char *license, const struct bpf_insn *insns, size_t insn_cnt, struct bpf_prog_load_opts *opts);`

**Parameters**

- `prog_type`: type of the program to load
- `prog_name`: name of the program
- `license`: license of the program
- `insns`: pointer to the program instructions
- `insn_cnt`: number of instructions in the program
- `opts`: options for the program load

### `struct bpf_prog_load_opts`

```c
struct bpf_prog_load_opts {
	size_t sz; /* size of this struct for forward/backward compatibility */

	/* libbpf can retry BPF_PROG_LOAD command if bpf() syscall returns
	 * -EAGAIN. This field determines how many attempts libbpf has to
	 *  make. If not specified, libbpf will use default value of 5.
	 */
	int attempts;

	enum bpf_attach_type expected_attach_type;
	__u32 prog_btf_fd;
	__u32 prog_flags;
	__u32 prog_ifindex;
	__u32 kern_version;

	__u32 attach_btf_id;
	__u32 attach_prog_fd;
	__u32 attach_btf_obj_fd;

	const int *fd_array;

	/* .BTF.ext func info data */
	const void *func_info;
	__u32 func_info_cnt;
	__u32 func_info_rec_size;

	/* .BTF.ext line info data */
	const void *line_info;
	__u32 line_info_cnt;
	__u32 line_info_rec_size;

	/* verifier log options */
	__u32 log_level;
	__u32 log_size;
	char *log_buf;
	/* output: actual total log contents size (including terminating zero).
	 * It could be both larger than original log_size (if log was
	 * truncated), or smaller (if log buffer wasn't filled completely).
	 * If kernel doesn't support this feature, log_size is left unchanged.
	 */
	__u32 log_true_size;
	__u32 token_fd;

	/* if set, provides the length of fd_array */
	__u32 fd_array_cnt;
	size_t :0;
};
```

#### `attempts`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/65cdd0c73d0b7d10e72f0dc8f4438594eed5368e)

#### `expected_attach_type`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/65cdd0c73d0b7d10e72f0dc8f4438594eed5368e)

#### `prog_btf_fd`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/65cdd0c73d0b7d10e72f0dc8f4438594eed5368e)

#### `prog_flags`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/65cdd0c73d0b7d10e72f0dc8f4438594eed5368e)

#### `prog_ifindex`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/65cdd0c73d0b7d10e72f0dc8f4438594eed5368e)

#### `kern_version`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/65cdd0c73d0b7d10e72f0dc8f4438594eed5368e)

#### `attach_btf_id`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/65cdd0c73d0b7d10e72f0dc8f4438594eed5368e)

#### `attach_prog_fd`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/65cdd0c73d0b7d10e72f0dc8f4438594eed5368e)

#### `attach_btf_obj_fd`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/65cdd0c73d0b7d10e72f0dc8f4438594eed5368e)

#### `fd_array`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/65cdd0c73d0b7d10e72f0dc8f4438594eed5368e)

#### `func_info`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/65cdd0c73d0b7d10e72f0dc8f4438594eed5368e)

#### `func_info_cnt`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/65cdd0c73d0b7d10e72f0dc8f4438594eed5368e)

#### `func_info_rec_size`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/65cdd0c73d0b7d10e72f0dc8f4438594eed5368e)

#### `line_info`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/65cdd0c73d0b7d10e72f0dc8f4438594eed5368e)

#### `line_info_cnt`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/65cdd0c73d0b7d10e72f0dc8f4438594eed5368e)

#### `line_info_rec_size`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/65cdd0c73d0b7d10e72f0dc8f4438594eed5368e)

#### `log_level`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/65cdd0c73d0b7d10e72f0dc8f4438594eed5368e)

#### `log_size`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/65cdd0c73d0b7d10e72f0dc8f4438594eed5368e)

#### `log_buf`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/65cdd0c73d0b7d10e72f0dc8f4438594eed5368e)

#### `log_true_size`

[:octicons-tag-24: 1.2.0](https://github.com/libbpf/libbpf/commit/c2fe7adb336b344c0024b0ed7d8280ce3e829e7d)

#### `token_fd`

[:octicons-tag-24: 1.4.0](https://github.com/libbpf/libbpf/commit/8082a311d393c9d3be5f6799dc2c106ae0bb4d93)

#### `fd_array_cnt`

[:octicons-tag-24: 1.6.0](https://github.com/libbpf/libbpf/commit/48c771c4ce53cf045187a6efce96cc24d50ab282)

## Usage

This function should only be used if you need precise control over the program loading process. For most cases, program should be loaded via [`bpf_object__load`](bpf_object__load.md) or similar high level APIs instead.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
