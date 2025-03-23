---
title: "Libbpf userspace function 'libbpf_set_memlock_rlim'"
description: "This page documents the 'libbpf_set_memlock_rlim' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `libbpf_set_memlock_rlim`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.7.0](https://github.com/libbpf/libbpf/releases/tag/v0.7.0)
<!-- [/LIBBPF_TAG] -->

Set the <nospell>memlock</nospell> resource limit for the current process.

## Definition

`#!c int libbpf_set_memlock_rlim(size_t memlock_bytes);`

**Parameters**

- `memlock_bytes`: the number of bytes to set the <nospell>memlock</nospell> resource limit to.

**Return**

`0`, on success; `<0`, on error.

## Usage

See [resource limit concept](../../../linux/concepts/resource-limit.md) page for more information.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
