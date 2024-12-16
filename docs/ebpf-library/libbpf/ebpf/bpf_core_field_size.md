---
title: "Libbpf eBPF macro 'bpf_core_field_size'"
description: "This page documents the 'bpf_core_field_size' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `bpf_core_field_size`

[:octicons-tag-24: v0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)

The `bpf_core_field_size` macro is used to query the size of a struct field on the kernel the program is being loaded on. The returned size is in bytes.

## Definition

```c
#define bpf_core_field_size(field...)					    \
	__builtin_preserve_field_info(___bpf_field_ref(field), BPF_FIELD_BYTE_SIZE)
```

## Usage

The `bpf_core_field_size` macro is used to query the size of a struct field on the kernel the program is being loaded on. The returned size is in bytes. The macro works for integers, struct/unions, pointers, arrays, and enums.

Supports two forms:

* field reference through variable access: `#!c bpf_core_field_size(p->my_field)`
* field reference through type and field names: `#!c bpf_core_field_size(struct my_type, my_field)`

This result is determined by the loader library such as libbpf, and set at load time and is considered as a constant value by the verifier.

### Example

```c
struct some_kernel_struct {
    __u16 field1; // This field might be smaller or larger on the target kernel
}

SEC("kprobe")
int kprobe__example(struct pt_regs *ctx)
{
    struct some_kernel_struct *a = PT_REGS_PARM1(ctx);
	__u64 tmp;

	tmp = 0;
	bpf_core_read_int(&tmp, bpf_core_field_size(a->field1), &a->field1);

    // Do something with tmp
    // ...

    return 0;
}
```
