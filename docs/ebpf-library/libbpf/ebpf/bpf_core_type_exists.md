---
title: "Libbpf eBPF macro 'bpf_core_type_exists'"
description: "This page documents the 'bpf_core_type_exists' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `bpf_core_type_exists`

[:octicons-tag-24: v0.2](https://github.com/libbpf/libbpf/releases/tag/v0.2)

The `bpf_core_type_exists` macro is used to check that provided named type (struct/union/enum/typedef) exists in a target kernel.

## Definition

```c
#define bpf_core_type_exists(type)					    \
	__builtin_preserve_type_info(*___bpf_typeof(type), BPF_TYPE_EXISTS)
```

## Usage

The `bpf_core_field_exists` macro is used to essentially ask a given type (struct/union/enum/typedef) exists in the kernel the program is being loaded on. One use case is to have fallback code in case a certain field is not present in the kernel.

Returns:

 * 1, if such type is present in target kernel's BTF
 * 0, if no matching type is found

This result is determined by the loader library such as libbpf, and set at load time. If a branch is never taken based on the result, it will not be evaluated by the verifier.

### Example

```c hl_lines="32 33"
struct some_kernel_struct {
    int a;
    int b;
};

SEC("kprobe")
int kprobe__example(struct pt_regs *ctx)
{
    int b;

    // Depending on if some_kernel_struct exists on the kernel we are running on
    // one or the other branch is taken. The verifier will only evaluate the branch
    // that is taken, and will optimize the if statement away, so this does not
    // impact the program's performance.
    if (bpf_core_field_exists(struct some_kernel_struct)) {
        struct some_kernel_struct *a = PT_REGS_PARM1(ctx);
        int b = BPF_CORE_READ(a, b);
    } else {
        int b = more_complex_fallback_to_get_b();
    }

    bpf_printk("Value of field 'a' = %d", b);

    return 0;
}
```
