---
title: "Libbpf eBPF macro 'BPF_SNPRINTF'"
description: "This page documents the 'BPF_SNPRINTF' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `BPF_SNPRINTF`

[:octicons-tag-24: v0.4.0](https://github.com/libbpf/libbpf/releases/tag/v0.4.0)

!!! note
    This macro was moved from `bpf_tracing.h` to `bpf_helpers.h` in [:octicons-tag-24: v0.5.0](https://github.com/libbpf/libbpf/releases/tag/v0.5.0)

The `BPF_SNPRINTF` macro is used to make in-BPF string formatting easier.

## Definition

```c
/*
 * BPF_SNPRINTF wraps the bpf_snprintf helper with variadic arguments instead of
 * an array of u64.
 */
#define BPF_SNPRINTF(out, out_size, fmt, args...)		\
({								\
	static const char ___fmt[] = fmt;			\
	unsigned long long ___param[___bpf_narg(args)];		\
								\
	_Pragma("GCC diagnostic push")				\
	_Pragma("GCC diagnostic ignored \"-Wint-conversion\"")	\
	___bpf_fill(___param, args);				\
	_Pragma("GCC diagnostic pop")				\
								\
	bpf_snprintf(out, out_size, ___fmt,			\
		     ___param, sizeof(___param));		\
})
```

## Usage

This macro is a wrapper around the [`bpf_snprintf`](../../../linux/helper-function/bpf_snprintf.md) helper. It places the literal format string in a global variable, this is necessary to get the compiler to emit code that will be accepted by the verifier.

### Example
```c hl_lines="17"
struct {
	__uint(type, BPF_MAP_TYPE_RINGBUF);
	__uint(max_entries, 4096);
} ringbuf SEC(".maps");

SEC("tc")
int example_prog(struct __sk_buff *ctx)
{
    const int size = 100;
    char *str_out = bpf_ringbuf_reserve(&ringbuf, size, 0);
    if (!str_out)
        return TC_ACT_OK;

    // `len` is the length of the formatted string including NULL-termination char.
    // if `len` > `size` then the string was truncated.
    // `len` can also be -EBUSY if the per-CPU memory copy buffer is busy.
    long len = BPF_SNPRINTF(str_out, size,
        "Got a packet from interface %d, src: %pi4, dst: %pi4, src port: %d, dst port: %d\\n", 
        ctx->ingress_ifindex, 
        ctx->remote_ip4, 
        ctx->local_ip4,
        ctx->remote_port,
        ctx->local_port
    );

    bpf_ringbuf_submit(str_out, 0);

    return TC_ACT_OK;
}
```
