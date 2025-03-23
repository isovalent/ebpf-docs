---
title: "Libbpf userspace function 'btf__find_by_name_kind'"
description: "This page documents the 'btf__find_by_name_kind' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__find_by_name_kind`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)
<!-- [/LIBBPF_TAG] -->

Find the type id of a type by its name and kind, in the BTF object.

## Definition

`#!c __s32 btf__find_by_name_kind(const struct btf *btf, const char *type_name, __u32 kind);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `type_name`: name of the type
- `kind`: kind of the type, one of the following:

```c
#define BTF_KIND_INT            1       /* Integer      */
#define BTF_KIND_PTR            2       /* Pointer      */
#define BTF_KIND_ARRAY          3       /* Array        */
#define BTF_KIND_STRUCT         4       /* Struct       */
#define BTF_KIND_UNION          5       /* Union        */
#define BTF_KIND_ENUM           6       /* Enumeration up to 32-bit values */
#define BTF_KIND_FWD            7       /* Forward      */
#define BTF_KIND_TYPEDEF        8       /* Typedef      */
#define BTF_KIND_VOLATILE       9       /* Volatile     */
#define BTF_KIND_CONST          10      /* Const        */
#define BTF_KIND_RESTRICT       11      /* Restrict     */
#define BTF_KIND_FUNC           12      /* Function     */
#define BTF_KIND_FUNC_PROTO     13      /* Function Proto       */
#define BTF_KIND_VAR            14      /* Variable     */
#define BTF_KIND_DATASEC        15      /* Section      */
#define BTF_KIND_FLOAT          16      /* Floating point       */
#define BTF_KIND_DECL_TAG       17      /* Decl Tag     */
#define BTF_KIND_TYPE_TAG       18      /* Type Tag     */
#define BTF_KIND_ENUM64         19      /* Enumeration up to 64-bit values */
```

**Return**

Return the type id of the type on success, or a negative error code on failure.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
