---
title: "Libbpf userspace function 'btf__type_by_id'"
description: "This page documents the 'btf__type_by_id' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__type_by_id`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Get the type by its id, in the BTF object.

## Definition

`#!c const struct btf_type *btf__type_by_id(const struct btf *btf, __u32 id);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `id`: ID of the type

**Return**

Return a pointer to a `struct btf_type` object on success, or `NULL` on failure.

### `struct btf_type`

```c
struct btf_type {
    __u32 name_off;
    /* "info" bits arrangement
     * bits  0-15: vlen (e.g. # of struct's members)
     * bits 16-23: unused
     * bits 24-28: kind (e.g. int, ptr, array...etc)
     * bits 29-30: unused
     * bit     31: kind_flag, currently used by
     *             struct, union, fwd, enum and enum64.
     */
    __u32 info;
    /* "size" is used by INT, ENUM, STRUCT, UNION and ENUM64.
     * "size" tells the size of the type it is describing.
     *
     * "type" is used by PTR, TYPEDEF, VOLATILE, CONST, RESTRICT,
     * FUNC, FUNC_PROTO, DECL_TAG and TYPE_TAG.
     * "type" is a type_id referring to another type.
     */
    union {
            __u32 size;
            __u32 type;
    };
};
```

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
