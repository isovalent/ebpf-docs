---
title: "Libbpf eBPF macro 'bpf_core_type_id_local'"
description: "This page documents the 'bpf_core_type_id_local' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `bpf_core_type_id_local`

[:octicons-tag-24: v0.2](https://github.com/libbpf/libbpf/releases/tag/v0.2)

The `bpf_core_type_id_local` macro to get BTF type ID of a specified type, using a local BTF information.

## Definition

```c
#define bpf_core_type_id_local(type)					    \
	__builtin_btf_type_id(*___bpf_typeof(type), BPF_TYPE_ID_LOCAL)
```

## Usage

The `bpf_core_type_id_local` macro to get BTF type ID of a specified type, using a local BTF information. Return 32-bit unsigned integer with type ID from program's own BTF. Always succeeds.

The local type ID can be useful to communicate type info from BPF to userspace or as parameter for kfuncs such as [`bpf_obj_new_impl`](../../../linux/kfuncs/bpf_obj_new_impl.md) or [`bpf_percpu_obj_new_impl`](../../../linux/kfuncs/bpf_percpu_obj_new_impl.md).

This result is determined by the loader library such as libbpf, and set at load time and is considered as a constant value by the verifier.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
