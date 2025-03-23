---
title: "Libbpf userspace function 'btf__add_btf'"
description: "This page documents the 'btf__add_btf' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__add_btf`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/releases/tag/v0.6.0)
<!-- [/LIBBPF_TAG] -->

Appends all the BTF types from `src_btf` into `btf`

## Definition

`#!c int btf__add_btf(struct btf *btf, const struct btf *src_btf);`

**Parameters**

- `btf`: BTF object which all the BTF types and strings are added to
- `src_btf`: BTF object which all BTF types and referenced strings are copied from

**Return**

BTF type ID of the first appended BTF type, or negative error code

## Usage

[`btf__add_btf`](btf__add_btf.md) can be used to simply and efficiently append the entire contents of one BTF object to another one. All the BTF type data is copied over, all referenced type IDs are adjusted by adding a necessary ID offset. Only strings referenced from BTF types are copied over and deduplicated, so if there were some unused strings in `src_btf`, those won't be copied over, which is consistent with the general string deduplication semantics of BTF writing APIs.

If any error is encountered during this process, the contents of `btf` is left intact, which means that [`btf__add_btf`](btf__add_btf.md) follows the transactional semantics and the operation as a whole is all-or-nothing.

`src_btf` has to be non-split BTF, as of now copying types from split BTF is not supported and will result in `-ENOTSUP` error code returned.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
