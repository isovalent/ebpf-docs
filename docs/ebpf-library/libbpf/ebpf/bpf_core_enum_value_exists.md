---
title: "Libbpf eBPF macro 'bpf_core_enum_value_exists'"
description: "This page documents the 'bpf_core_enum_value_exists' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `bpf_core_enum_value_exists`

[:octicons-tag-24: v0.2](https://github.com/libbpf/libbpf/releases/tag/v0.2)

The `bpf_core_enum_value_exists` macro is used to check that provided enumerator value is defined in a target kernel.

## Definition

```c
#ifdef __clang__
#define bpf_core_enum_value_exists(enum_type, enum_value)		    \
	__builtin_preserve_enum_value(*(typeof(enum_type) *)enum_value, BPF_ENUMVAL_EXISTS)
#else
#define bpf_core_enum_value_exists(enum_type, enum_value)		    \
	__builtin_preserve_enum_value(___bpf_typeof(enum_type), enum_value, BPF_ENUMVAL_EXISTS)
#endif
```

## Usage

The `bpf_core_enum_value_exists` macro is used to check that provided enumerator value is defined in a target kernel.

Returns:

 * 1, if specified enum type and its enumerator value are present in target kernel's BTF
 * 0, if no matching enum and/or enum value within that enum is found

This result is determined by the loader library such as libbpf, and set at load time. If a branch is never taken based on the result, it will not be evaluated by the verifier.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
