---
title: "Libbpf userspace function 'btf_ext__endianness'"
description: "This page documents the 'btf_ext__endianness' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf_ext__endianness`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.5.0](https://github.com/libbpf/libbpf/releases/tag/v1.5.0)
<!-- [/LIBBPF_TAG] -->

Get the endianness of the BTF extension object.

## Definition

`#!c enum btf_endianness btf_ext__endianness(const struct btf_ext *btf_ext);`

**Parameters**

- `btf_ext`: BTF extension object

**Return**

The endianness of the BTF extension object.

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
