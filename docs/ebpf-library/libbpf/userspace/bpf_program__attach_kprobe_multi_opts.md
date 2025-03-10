---
title: "Libbpf userspace function 'bpf_program__attach_kprobe_multi_opts'"
description: "This page documents the 'bpf_program__attach_kprobe_multi_opts' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_kprobe_multi_opts`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)
<!-- [/LIBBPF_TAG] -->

Attach a [`BPF_PROG_TYPE_KPROBE`](../../../linux/program-type/BPF_PROG_TYPE_KPROBE.md) program at multiple places at once.

## Definition

`#!c struct bpf_link * bpf_program__attach_kprobe_multi_opts(const struct bpf_program *prog, const char *pattern, const struct bpf_kprobe_multi_opts *opts);`

**Parameters**

- `prog`: pointer to the BPF program object.
- `pattern`: pattern to match the kernel functions to attach the probe to.
- `opts`: options to attach the probe.

### `struct bpf_kprobe_multi_opts`

```c
struct bpf_kprobe_multi_opts {
	/* size of this struct, for forward/backward compatibility */
	size_t sz;
	const char **syms;
	const unsigned long *addrs;
	const __u64 *cookies;
	/* number of elements in syms/addrs/cookies arrays */
	size_t cnt;
	bool retprobe;
	bool session;
	bool unique_match;
	size_t :0;
};
```

#### `syms`

[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/commit/05acce9e03d9b25ea909be35fd67782d88a21ba3)

Array of function symbols to attach. `cnt` must be set to the number of `syms`. Mutually exclusive with `addrs` and `pattern`.

#### `addrs`

[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/commit/05acce9e03d9b25ea909be35fd67782d88a21ba3)

Array of function addresses to attach. `cnt` must be set to the number of `addrs`. Mutually exclusive with `syms` and `pattern`.

#### `cookies`

[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/commit/05acce9e03d9b25ea909be35fd67782d88a21ba3)

Array of user-provided values fetchable through [`bpf_get_attach_cookie`](../../../linux/helper-function/bpf_get_attach_cookie.md). The number of cookies must be equal to the number of `syms` or `addrs`.

This allows the program to know for which attach point it is being called.

This field is optional, and can be `NULL` if you do not need to pass any cookies.

#### `session`

[:octicons-tag-24: 1.5.0](https://github.com/libbpf/libbpf/commit/6c3cf5108efff1b4a3bf4c8119abd2eb8a669aea)

Created the kprobe in session mode. In session mode, a kprobe and kretprobe are associated with each other. The kprobe can decide if the kretprobe will be triggered upon function return. Both programs also share the same cookie.

#### `unique_match`

[:octicons-tag-24: 1.6.0](https://github.com/libbpf/libbpf/commit/32792ec66c9daadf3740a373aaad6d2c526c4ca2)

When set, libbpf will error if the given `pattern` matches more than one function.

This is useful when you want to attach a kprobe to a single function, which has been renamed or can have different names in different kernel versions. For `try_to_wake_up()` or `try_to_wake_up.llvm.<hash>()`, you can use `try_to_wake_up*` as the pattern and set `unique_match` to `true`.

## Usage

User can specify functions to attach with `pattern` argument that allows wildcards (*?' supported) or provide symbols or addresses directly through opts argument. These 3 options are mutually exclusive.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
