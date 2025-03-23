---
title: "Libbpf userspace function 'btf_dump__new'"
description: "This page documents the 'btf_dump__new' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf_dump__new`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/releases/tag/v0.6.0)
<!-- [/LIBBPF_TAG] -->

Create a new BTF dumper object.

## Definition

```c
typedef void (*btf_dump_printf_fn_t)(void *ctx, const char *fmt, va_list args);

struct btf_dump *btf_dump__new(const struct btf *btf, btf_dump_printf_fn_t printf_fn, void *ctx, const struct btf_dump_opts *opts);
```

### `struct btf_dump_opts`

```c
struct btf_dump_opts {
	size_t sz;
};
```

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
