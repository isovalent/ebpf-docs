---
title: "Libbpf eBPF macro 'BPF_PROG'"
description: "This page documents the 'BPF_PROG' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `BPF_PROG`

[:octicons-tag-24: v0.0.8](https://github.com/libbpf/libbpf/releases/tag/v0.0.8)

The `BPF_PROG` macro makes it easier to write programs for program types that receive `[]u64` contexts such as [`BPF_PROG_TYPE_TRACING`](../../../linux/program-type/BPF_PROG_TYPE_TRACING.md) programs.

## Definition

```c
#define BPF_PROG(name, args...)						    \
name(unsigned long long *ctx);						    \
static __always_inline typeof(name(0))					    \
____##name(unsigned long long *ctx, ##args);				    \
typeof(name(0)) name(unsigned long long *ctx)				    \
{									    \
	_Pragma("GCC diagnostic push")					    \
	_Pragma("GCC diagnostic ignored \"-Wint-conversion\"")		    \
	return ____##name(___bpf_ctx_cast(args));			    \
	_Pragma("GCC diagnostic pop")					    \
}									    \
static __always_inline typeof(name(0))					    \
____##name(unsigned long long *ctx, ##args)
```

## Usage

This macro is useful when using program types that have a `[]u64` context type (typically written as `unsigned long long *`). 

Conventionally with these program contexts, the arguments to the program are put in this array. So the first argument would be in `ctx[0]`, the second in `ctx[1]`. It is up to the program author to cast them into their actual type.

The `BPF_PROG` macro allows you to write your program with a normal function signature, the macro will then do the casting for you.

!!! note
    The original context will stay available as `ctx`, if you ever wish to access it manually or need to pass it to a helper or kfunc. Therefor, the variable name `ctx` should not be reused in arguments or function body.

!!! warning
    This macro assumes a 1 to 1 conversion between a `u64` and argument. However, the Sys V calling convention allows types such as structs of up to 16 bytes to be passed over 2 registers and thus two `u64`s in the context. That breaks the assumption and may lead to hard to resolve bugs. The [`BPF_PROG2`](BPF_PROG2.md) macro is the improved version of this one which does account for this. Its recommend to use the second version when you might be dealing with arguments larger than 8 bytes.

### Example

```c hl_lines="2"
SEC("struct_ops/hid_device_event")
int BPF_PROG(filter_switch, struct hid_bpf_ctx *hid_ctx)
{
    __u8 *data = hid_bpf_get_data(hid_ctx, 0 /* offset */, 192 /* size */);
    __u8 *buf;

    if (!data)
        return 0; /* EPERM check */

    if (current_value != data[152]) {
        buf = bpf_ringbuf_reserve(&ringbuf, 1, 0);
        if (!buf)
            return 0;

        *buf = data[152];

        bpf_ringbuf_commit(buf, 0);

        current_value = data[152];
    }

    return 0;
}
```
