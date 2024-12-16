---
title: "Libbpf eBPF macro 'bpf_core_enum_value'"
description: "This page documents the 'bpf_core_enum_value' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `bpf_core_enum_value`

[:octicons-tag-24: v0.2](https://github.com/libbpf/libbpf/releases/tag/v0.2)

The `bpf_core_enum_value` macro is used to get the integer value of an enumerator value in a target kernel.

## Definition

```c
#ifdef __clang__
#define bpf_core_enum_value(enum_type, enum_value)			    \
	__builtin_preserve_enum_value(*(typeof(enum_type) *)enum_value, BPF_ENUMVAL_VALUE)
#else
#define bpf_core_enum_value(enum_type, enum_value)			    \
	__builtin_preserve_enum_value(___bpf_typeof(enum_type), enum_value, BPF_ENUMVAL_VALUE)
#endif
```

## Usage

The `bpf_core_enum_value` macro is used to get the integer value of an enumerator value in a target kernel.

Returns:

 * 64-bit value, if specified enum type and its enumerator value are present in target kernel's BTF
 * 0, if no matching enum and/or enum value within that enum is found

This result is determined by the loader library such as libbpf, and set at load time and is considered as a constant value by the verifier.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
