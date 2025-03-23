---
title: "Libbpf userspace function 'btf__add_var'"
description: "This page documents the 'btf__add_var' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__add_var`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Append new `BTF_KIND_VAR` type to BTF object.

## Definition

`#!c int btf__add_var(struct btf *btf, const char *name, int linkage, int type_id);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `name`: name of the variable, can't be `NULL` or empty;
- `linkage`: linkage type of the variable;

**Return**

`>0`, type ID of newly added BTF type; `<0`, on error.

### `enum btf_var_linkage`

```c
enum {
	BTF_VAR_STATIC = 0,
	BTF_VAR_GLOBAL_ALLOCATED = 1,
	BTF_VAR_GLOBAL_EXTERN = 2,
};
```

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
