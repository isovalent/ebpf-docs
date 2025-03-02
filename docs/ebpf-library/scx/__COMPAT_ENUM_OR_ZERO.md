---
title: "SCX eBPF macro '__COMPAT_ENUM_OR_ZERO'"
description: "This page documents the '__COMPAT_ENUM_OR_ZERO' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `__COMPAT_ENUM_OR_ZERO`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/2a52ca7c98960aafb0eca9ef96b2d0c932171357)

The `__COMPAT_ENUM_OR_ZERO` macro returns the value of an enum member or zero if the member is not defined.

## Definition

```c
#define __COMPAT_ENUM_OR_ZERO(__type, __ent)        \
({                                                  \
	__type __ret = 0;                               \
	if ([bpf_core_enum_value_exists](../libbpf/ebpf/bpf_core_enum_value_exists.md)(__type, __ent))  \
		__ret = __ent;                              \
	__ret;                                          \
})
```

## Usage

This macro checks at runtime if the enum member `__ent` exists in the enum type `__type`. If it does, it returns the value of the that was known at compile time, or zero. It does not return the value of the enum on the target kernel, if it differs from the value known at compile time.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
