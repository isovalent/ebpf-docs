---
title: "SCX eBPF macro 'scx_bpf_bstr_preamble'"
description: "This page documents the 'scx_bpf_bstr_preamble' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `scx_bpf_bstr_preamble`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/2a52ca7c98960aafb0eca9ef96b2d0c932171357)

The `scx_bpf_bstr_preamble` macro initializes the `fmt` and variadic argument inputs to [`scx_bpf_dump_bstr`](../../linux/kfuncs/scx_bpf_dump_bstr.md), [`scx_bpf_error_bstr`](../../linux/kfuncs/scx_bpf_error_bstr.md) and [`scx_bpf_exit_bstr`](../../linux/kfuncs/scx_bpf_exit_bstr.md) kfuncs. Callers to this function should use `___fmt` and `___param` to refer to the initialized list of inputs to the bstr kfunc.

!!! note
    Note that `__param[]` must have at least one element to keep the verifier happy.	

## Definition

```c
#define scx_bpf_bstr_preamble(fmt, args...)                         \
	static char ___fmt[] = fmt;                                     \
	unsigned long long ___param[___bpf_narg(args) ?: 1] = {};       \
                                                                    \
	_Pragma("GCC diagnostic push")                                  \
	_Pragma("GCC diagnostic ignored \"-Wint-conversion\"")          \
	[___bpf_fill](../libbpf/ebpf/___bpf_fill.md)(___param, args);                                    \
	_Pragma("GCC diagnostic pop")
```

## Usage

This macro is used internally in [`scx_bpf_exit`](scx_bpf_exit.md), [`scx_bpf_error`](scx_bpf_error.md)
and [`scx_bpf_dump`](scx_bpf_dump.md). These prepare the parameters and call the kfuncs. But you could use this macro manually if you wish.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
