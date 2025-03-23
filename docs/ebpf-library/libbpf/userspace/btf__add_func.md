---
title: "Libbpf userspace function 'btf__add_func'"
description: "This page documents the 'btf__add_func' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__add_func`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Append new `BTF_KIND_FUNC` type to BTF object.

## Definition

`#!c int btf__add_func(struct btf *btf, const char *name, enum btf_func_linkage linkage, int proto_type_id);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `name`: name of the function, can't be `NULL` or empty;
- `linkage`: linkage type of the function;
- `proto_type_id`: `FUNC_PROTO`'s type ID, it might not exist yet;

**Return**

`>0`, type ID of newly added BTF type; `<0`, on error.

### `enum btf_func_linkage`

```c
enum btf_func_linkage {
	BTF_FUNC_STATIC = 0,
	BTF_FUNC_GLOBAL = 1,
	BTF_FUNC_EXTERN = 2,
};
```

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
