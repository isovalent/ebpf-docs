---
title: "Libbpf userspace function 'btf_ext__set_endianness'"
description: "This page documents the 'btf_ext__set_endianness' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf_ext__set_endianness`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.5.0](https://github.com/libbpf/libbpf/releases/tag/v1.5.0)
<!-- [/LIBBPF_TAG] -->

Set the endianness of the BTF extension object.

## Definition

`#!c int btf_ext__set_endianness(struct btf_ext *btf_ext, enum btf_endianness endian);`

**Parameters**

- `btf_ext`: pointer to a `struct btf_ext` object
- `endian`: endianness to set

**Return**

Return 0 on success, or a negative error code on failure.

### `enum btf_endianness`

```c
enum btf_endianness {
	BTF_LITTLE_ENDIAN = 0,
	BTF_BIG_ENDIAN = 1,
};
```

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
