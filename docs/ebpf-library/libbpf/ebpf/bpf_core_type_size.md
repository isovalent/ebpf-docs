---
title: "Libbpf eBPF macro 'bpf_core_type_size'"
description: "This page documents the 'bpf_core_type_size' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `bpf_core_type_size`

[:octicons-tag-24: v0.2](https://github.com/libbpf/libbpf/releases/tag/v0.2)

The `bpf_core_type_size` macro is used to get the byte size of a provided named type (struct/union/enum/typedef) in a target kernel.

## Definition

```c
#define bpf_core_type_size(type)					    \
	__builtin_preserve_type_info(*___bpf_typeof(type), BPF_TYPE_SIZE)
```

## Usage

The `bpf_core_type_size` macro is used to get the byte size of a provided named type (struct/union/enum/typedef) in a target kernel. The returned size is in bytes.

Returns:

 * \>= 0 size (in bytes), if type is present in target kernel's BTF
 * 0, if no matching type is found

This result is determined by the loader library such as libbpf, and set at load time and is considered as a constant value by the verifier.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
