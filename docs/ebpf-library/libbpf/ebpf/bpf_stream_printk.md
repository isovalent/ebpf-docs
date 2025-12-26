---
title: "Libbpf eBPF macro 'bpf_stream_printk'"
description: "This page documents the 'bpf_stream_printk' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `bpf_stream_printk`

[:octicons-tag-24: v1.6.0](https://github.com/libbpf/libbpf/releases/tag/v1.6.0)

The `bpf_stream_printk` macro is used to make writing to streams easier.

## Definition

```c
#define bpf_stream_printk(stream_id, fmt, args...)				\
({										\
	static const char ___fmt[] = fmt;					\
	unsigned long long ___param[___bpf_narg(args)];				\
										\
	_Pragma("GCC diagnostic push")						\
	_Pragma("GCC diagnostic ignored \"-Wint-conversion\"")			\
	___bpf_fill(___param, args);						\
	_Pragma("GCC diagnostic pop")						\
										\
	bpf_stream_vprintk_impl(stream_id, ___fmt, ___param, sizeof(___param), NULL);\
})
```

## Usage

This macro is a wrapper around the [`bpf_stream_vprintk_impl`](../../../linux/kfuncs/bpf_stream_vprintk_impl.md) helper. It places the literal format string in a global variable, this is necessary to get the compiler to emit code that will be accepted by the verifier.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
