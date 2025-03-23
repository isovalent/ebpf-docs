---
title: "Libbpf userspace function 'bpf_btf_load'"
description: "This page documents the 'bpf_btf_load' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_btf_load`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.7.0](https://github.com/libbpf/libbpf/releases/tag/v0.7.0)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_BTF_LOAD`](../../../linux/syscall/BPF_BTF_LOAD.md) syscall command.

## Definition

`#!c int bpf_btf_load(const void *btf_data, size_t btf_size, struct bpf_btf_load_opts *opts);`

**Parameters**

- `btf_data`: pointer to the BTF data to load
- `btf_size`: size of the BTF data
- `opts`: options for the BTF load

### `struct bpf_btf_load_opts`

```c
struct bpf_btf_load_opts {
	size_t sz; /* size of this struct for forward/backward compatibility */

	/* kernel log options */
	char *log_buf;
	__u32 log_level;
	__u32 log_size;
	/* output: actual total log contents size (including terminating zero).
	 * It could be both larger than original log_size (if log was
	 * truncated), or smaller (if log buffer wasn't filled completely).
	 * If kernel doesn't support this feature, log_size is left unchanged.
	 */
	__u32 log_true_size;

	__u32 btf_flags;
	__u32 token_fd;
	size_t :0;
};
```

#### `log_buf`

[:octicons-tag-24: 0.7.0](https://github.com/libbpf/libbpf/commit/3c93f7ddb20f90f3f272ccbbe34339737bea0501)

#### `log_level`

[:octicons-tag-24: 0.7.0](https://github.com/libbpf/libbpf/commit/3c93f7ddb20f90f3f272ccbbe34339737bea0501)

#### `log_size`

[:octicons-tag-24: 0.7.0](https://github.com/libbpf/libbpf/commit/3c93f7ddb20f90f3f272ccbbe34339737bea0501)

#### `log_true_size`

[:octicons-tag-24: 1.2.0](https://github.com/libbpf/libbpf/commit/2f01564c50deb039896f1ce65b4d3698be3833ea)

#### `btf_flags`

[:octicons-tag-24: 1.4.0](https://github.com/libbpf/libbpf/commit/ac4a66ea12924239ae842b98da3bf94373b7d6a8)

#### `token_fd`

[:octicons-tag-24: 1.4.0](https://github.com/libbpf/libbpf/commit/ac4a66ea12924239ae842b98da3bf94373b7d6a8)

## Usage

This function should only be used if you need precise control over the BTF loading process. For most cases, program should be loaded via [`bpf_object__load`](bpf_object__load.md) or similar high level APIs instead.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
