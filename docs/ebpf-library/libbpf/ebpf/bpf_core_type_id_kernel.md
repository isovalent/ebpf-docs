---
title: "Libbpf eBPF macro 'bpf_core_type_id_kernel'"
description: "This page documents the 'bpf_core_type_id_kernel' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `bpf_core_type_id_kernel`

[:octicons-tag-24: v0.2](https://github.com/libbpf/libbpf/releases/tag/v0.2)

The `bpf_core_type_id_kernel` macro to get BTF type ID of a target kernel's type that matches specified local type.

## Definition

```c
#define bpf_core_type_id_kernel(type)					    \
	__builtin_btf_type_id(*___bpf_typeof(type), BPF_TYPE_ID_TARGET)
```

## Usage

The `bpf_core_type_id_kernel` macro to get BTF type ID of a target kernel's type that matches specified local type.

The target BTF type ID can be used to construct a `struct btf_ptr` to be used as parameter to [`bpf_snprintf_btf`](../../../linux/helper-function/bpf_snprintf_btf.md).

Returns:

 * valid 32-bit unsigned type ID in kernel BTF
 * 0, if no matching type was found in a target kernel BTF

This result is determined by the loader library such as libbpf, and set at load time and is considered as a constant value by the verifier.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
