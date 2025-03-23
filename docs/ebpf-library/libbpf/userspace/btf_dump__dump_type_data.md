---
title: "Libbpf userspace function 'btf_dump__dump_type_data'"
description: "This page documents the 'btf_dump__dump_type_data' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf_dump__dump_type_data`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.5.0](https://github.com/libbpf/libbpf/releases/tag/v0.5.0)
<!-- [/LIBBPF_TAG] -->

Dump BTF type data.

## Definition

`#!c int btf_dump__dump_type_data(struct btf_dump *d, __u32 id, const void *data, size_t data_sz, const struct btf_dump_type_data_opts *opts);`

**Parameters**

- `d`: pointer to a `struct btf_dump` object
- `id`: BTF type ID
- `data`: pointer to BTF type data
- `data_sz`: size of BTF type data
- `opts`: pointer to a `struct btf_dump_type_data_opts` object

**Return**

`0`, on success; `-errno`, on error.

### `struct btf_dump_type_data_opts`

```c
struct btf_dump_type_data_opts {
	/* size of this struct, for forward/backward compatibility */
	size_t sz;
	const char *indent_str;
	int indent_level;
	/* below match "show" flags for bpf_show_snprintf() */
	bool compact;
	bool skip_names;
	bool emit_zeroes;
	size_t :0;
};
```

#### `indent_str`

[:octicons-tag-24: 0.5.0](https://github.com/libbpf/libbpf/commit/c0b2ceba1d25da148cdf369fb727b613d78c1f19)

#### `indent_level`

[:octicons-tag-24: 0.5.0](https://github.com/libbpf/libbpf/commit/c0b2ceba1d25da148cdf369fb727b613d78c1f19)

#### `compact`

[:octicons-tag-24: 0.5.0](https://github.com/libbpf/libbpf/commit/c0b2ceba1d25da148cdf369fb727b613d78c1f19)

no newlines/indentation

#### `skip_names`

[:octicons-tag-24: 0.5.0](https://github.com/libbpf/libbpf/commit/c0b2ceba1d25da148cdf369fb727b613d78c1f19)

skip member/type names

#### `emit_zeroes`

[:octicons-tag-24: 0.5.0](https://github.com/libbpf/libbpf/commit/c0b2ceba1d25da148cdf369fb727b613d78c1f19)

show 0-valued fields

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
