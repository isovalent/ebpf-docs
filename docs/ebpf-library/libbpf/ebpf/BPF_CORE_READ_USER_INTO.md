---
title: "Libbpf eBPF macro 'BPF_CORE_READ_USER_INTO'"
description: "This page documents the 'BPF_CORE_READ_USER_INTO' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `BPF_CORE_READ_USER_INTO`

[:octicons-tag-24: v0.4.0](https://github.com/libbpf/libbpf/releases/tag/v0.4.0)

The `BPF_CORE_READ_USER_INTO` macro is a more performance-conscious variant of [`BPF_CORE_READ_USER`](BPF_CORE_READ.md), in which final field is read into user-provided storage.

## Definition

```c
#define BPF_CORE_READ_USER_INTO(dst, src, a, ...) ({			    \
	___core_read(bpf_core_read_user, bpf_core_read_user,		    \
		     dst, (src), a, ##__VA_ARGS__)			    \
})
```

## Usage

`BPF_CORE_READ_USER_INTO` is very similar to [`BPF_CORE_READ_USER`](BPF_CORE_READ_USER.md), but instead of returning the value, it writes the value into the provided destination.

The following two code snippets are equivalent:

```c
int x = BPF_CORE_READ_USER(s, a.b.c, d.e, f, g);
```

```c
int x;
BPF_CORE_READ_USER_INTO(&x, s, a.b.c, d.e, f, g);
```

Please refer to the [`BPF_CORE_READ_USER`](BPF_CORE_READ_USER.md) documentation for more details on usage of it and its variants like this macros.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
