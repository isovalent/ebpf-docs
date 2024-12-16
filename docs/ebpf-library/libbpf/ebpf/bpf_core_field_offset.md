---
title: "Libbpf eBPF macro 'bpf_core_field_offset'"
description: "This page documents the 'bpf_core_field_offset' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `bpf_core_field_offset`

[:octicons-tag-24: v0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)

The `bpf_core_field_offset` macro is used to query the offset of a struct field on the kernel the program is being loaded on. The returned offset is in bytes.

## Definition

```c
#define bpf_core_field_offset(field...)					    \
	__builtin_preserve_field_info(___bpf_field_ref(field), BPF_FIELD_BYTE_OFFSET)
```

## Usage

The `bpf_core_field_offset` macro is used to query the offset of a struct field on the kernel the program is being loaded on. The returned size is in bytes.

Supports two forms:

* field reference through variable access: `#!c bpf_core_field_offset(p->my_field)`
* field reference through type and field names: `#!c bpf_core_field_offset(struct my_type, my_field)`

This result is determined by the loader library such as libbpf, and set at load time and is considered as a constant value by the verifier.

### Example

```c
struct some_kernel_struct {
    __u16 field1; // This field might be smaller or larger on the target kernel
    __u32 field2;
}

SEC("kprobe")
int kprobe__example(struct pt_regs *ctx)
{
    struct some_kernel_struct *a = PT_REGS_PARM1(ctx);
    int field_offset = bpf_core_field_offset(a->field2)

    // Do something with field_offset
    // ...

    return 0;
}
```
