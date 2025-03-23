---
title: "Libbpf userspace function 'btf_dump__dump_type'"
description: "This page documents the 'btf_dump__dump_type' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf_dump__dump_type`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.4](https://github.com/libbpf/libbpf/releases/tag/v0.0.4)
<!-- [/LIBBPF_TAG] -->

Dump BTF type in a compilable C syntax.

## Definition

`#!c int btf_dump__dump_type(struct btf_dump *d, __u32 id);`

**Parameters**

- `d`: pointer to a `struct btf_dump` object
- `id`: BTF type ID to dump

**Return**

`0` on success; `<0`, otherwise.

## Usage

Dump BTF type in a compilable C syntax, including all the necessary dependent types, necessary for compilation. If some of the dependent types were already emitted as part of previous `btf_dump__dump_type` invocation for another type, they won't be emitted again. This API allows callers to filter out BTF types according to user-defined criteria and emitted only minimal subset of types, necessary to compile everything. Full struct/union definitions will still be emitted, even if the only usage is through pointer and could be satisfied with just a forward declaration.
 
Dumping is done in two high-level passes:

1. Topologically sort type definitions to satisfy C rules of compilation.
2. Emit type definitions in C syntax.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
