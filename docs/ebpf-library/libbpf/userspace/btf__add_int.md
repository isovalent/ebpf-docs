---
title: "Libbpf userspace function 'btf__add_int'"
description: "This page documents the 'btf__add_int' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__add_int`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Append new `BTF_KIND_INT` type to the BTF object.

## Definition

`#!c int btf__add_int(struct btf *btf, const char *name, size_t byte_sz, int encoding);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `name`: name of the type; non-empty, non-NULL type name;
- `byte_sz`: power-of-2 (1, 2, 4, ..) size of the type, in bytes;
- `encoding`: combination of `BTF_INT_SIGNED`, `BTF_INT_CHAR`, `BTF_INT_BOOL`.

```c
/* Attributes stored in the BTF_INT_ENCODING */
#define BTF_INT_SIGNED	(1 << 0)
#define BTF_INT_CHAR	(1 << 1)
#define BTF_INT_BOOL	(1 << 2)
```

**Return**

`>0`, type ID of newly added BTF type; `<0`, on error.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
