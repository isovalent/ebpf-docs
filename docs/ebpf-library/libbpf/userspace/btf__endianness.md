---
title: "Libbpf userspace function 'btf__endianness'"
description: "This page documents the 'btf__endianness' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__endianness`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.2.0](https://github.com/libbpf/libbpf/releases/tag/v0.2.0)
<!-- [/LIBBPF_TAG] -->

Get the endianness of the BTF object.

## Definition

`#!c enum btf_endianness btf__endianness(const struct btf *btf);`

**Parameters**

- `btf`: pointer to a `struct btf` object

**Return**

Return the endianness of the BTF object.

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
