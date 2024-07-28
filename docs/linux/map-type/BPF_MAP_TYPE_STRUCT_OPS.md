---
title: "Map Type 'BPF_MAP_TYPE_STRUCT_OPS'"
description: "This page documents the 'BPF_MAP_TYPE_STRUCT_OPS' eBPF map type, including its definition, usage, program types that can use it, and examples."
---
# Map type `BPF_MAP_TYPE_STRUCT_OPS`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_STRUCT_OPS) -->
[:octicons-tag-24: v5.6](https://github.com/torvalds/linux/commit/85d33df357b634649ddbe0a20fd2d0fc5732c3cb)
<!-- [/FEATURE_TAG] -->

Struct ops maps are specialized maps that act as implementations of "struct ops" structures defined in the kernel.

## Usage

The kernel has the concept of of "struct ops" which are function pointers inside a struct, this is the kernels way of implementing polymorphism. A call site defines a structure and contract for how the implementations should behave. A new implementation can then create an instance of this struct and set its own functions.

The struct ops map type is meant to serve as a way to allocate memory for an instant of struct ops. BPF programs are then set as field values to implement the functions in the struct ops.

Struct ops maps only have 1 key, that being `0`. The value is the struct formatted following C struct layout rules. In place of function pointers, file descriptors to BPF programs are used (the kernel converts to a memory address under the hood). Fields can be left empty assuming the contract allows it. Not all fields are functions, some fields can also be other C data type such as integers or nested structures.

If the `BPF_F_LINK` flag is not used, the struct ops is attached when the value of the map is set using `BPF_MAP_UPDATE_ELEM` and can be detached using `BPF_MAP_DELETE_ELEM`. If the `BPF_F_LINK` flag is used, the struct ops is attached using a BPF link and `BPF_MAP_DELETE_ELEM` can't be used to detach.

## Attributes

The [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) of the map must be `4` and the [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) must be equal to the size of target struct type. The [`max_entries`](../syscall/BPF_MAP_CREATE.md#max_entries) must be set to `1`. 

[`btf_vmlinux_value_type_id`](../syscall/BPF_MAP_CREATE.md#btf_vmlinux_value_type_id) must be set to the BTF type id of the target struct type in vmlinux. Since [:octicons-tag-24: v6.9](https://github.com/torvalds/linux/commit/fcc2c1fb0651477c8ed78a3a293c175ccd70697a) it is possible to attach to struct ops defined in kernel modules. In that case the `BPF_F_VTYPE_BTF_OBJ_FD` flag must be set and the `value_type_btf_obj_fd` must be set to the file descriptor of the BTF object file that contains the target struct type. This is required since BTF ids above the max ID found in vmlinux are not unique across kernel modules. 

## Syscall commands

* [`BPF_MAP_LOOKUP_ELEM`](../syscall/BPF_MAP_LOOKUP_ELEM.md)
* [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md)
* [`BPF_MAP_DELETE_ELEM`](../syscall/BPF_MAP_DELETE_ELEM.md)

## Helper functions

The following helper functions work with this map type:

<!-- DO NOT EDIT MANUALLY -->
<!-- [MAP_HELPER_FUNC_REF] -->
<!-- [/MAP_HELPER_FUNC_REF] -->

## Flags

### `BPF_F_LINK` 

[:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/68b04864ca425d1894c96b8141d4fba1181f11cb)

If this flag is specified, the defined struct ops isn't directly attached to the call site of the struct ops. Instead, a BPF link must be used to attach the struct ops to the call site. 

### `BPF_F_VTYPE_BTF_OBJ_FD`

[:octicons-tag-24: v6.9](https://github.com/torvalds/linux/commit/fcc2c1fb0651477c8ed78a3a293c175ccd70697a)

This flag is set to indicate that the BTF ID provided in `btf_vmlinux_value_type_id` is to be found in a kernel module. The `value_type_btf_obj_fd` attribute must be set to the file descriptor of the BTF object of the kernel module that contains the target struct type.

