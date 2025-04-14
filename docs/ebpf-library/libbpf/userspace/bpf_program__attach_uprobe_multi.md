---
title: "Libbpf userspace function 'bpf_program__attach_uprobe_multi'"
description: "This page documents the 'bpf_program__attach_uprobe_multi' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_uprobe_multi`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/releases/tag/v1.3.0)
<!-- [/LIBBPF_TAG] -->

Attaches a BPF program to multiple uprobes with uprobe_multi link.

## Definition

`#!c struct bpf_link * bpf_program__attach_uprobe_multi(const struct bpf_program *prog, pid_t pid, const char *binary_path, const char *func_pattern, const struct bpf_uprobe_multi_opts *opts);`

**Parameters**

- `prog`: BPF program to attach
- `pid`: Process ID to attach the uprobe to, 0 for self (own process),
`-1` for all processes
- `binary_path`: Path to binary
- `func_pattern`: Regular expression to specify functions to attach
BPF program to
- `opts`: Additional options

**Return**

`0`, on success; negative error code, otherwise

### `struct bpf_uprobe_multi_opts`

```c
struct bpf_uprobe_multi_opts {
	/* size of this struct, for forward/backward compatibility */
	size_t sz;
	/* array of function symbols to attach to */
	const char **syms;
	/* array of function addresses to attach to */
	const unsigned long *offsets;
	/* optional, array of associated ref counter offsets */
	const unsigned long *ref_ctr_offsets;
	/* optional, array of associated BPF cookies */
	const __u64 *cookies;
	/* number of elements in syms/addrs/cookies arrays */
	size_t cnt;
	/* create return uprobes */
	bool retprobe;
	/* create session kprobes */
	bool session;
	size_t :0;
};
```


#### `syms`

[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/commit/9f76dd6dd09e324f34c2450533bae99ea599601f)

Array of function symbols to attach. `cnt` must be set to the number of `syms`.

#### `offsets`

[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/commit/9f76dd6dd09e324f34c2450533bae99ea599601f)

Array of offsets from the start of the function to the instruction to add the uprobe to. `cnt` must be set to the number of `offsets`.

#### `cookies`

[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/commit/9f76dd6dd09e324f34c2450533bae99ea599601f)

Array of user-provided values fetchable through [`bpf_get_attach_cookie`](../../../linux/helper-function/bpf_get_attach_cookie.md). The number of cookies must be equal to the number of `syms` or `addrs`.

This allows the program to know for which attach point it is being called.

This field is optional, and can be `NULL` if you do not need to pass any cookies.

#### `ref_ctr_offsets`

[:octicons-tag-24: 1.3.0](https://github.com/libbpf/libbpf/commit/9f76dd6dd09e324f34c2450533bae99ea599601f)

Array of offsets to USDT reference counter fields. See:
- [https://sourceware.org/systemtap/wiki/UserSpaceProbeImplementation](https://sourceware.org/systemtap/wiki/UserSpaceProbeImplementation)
- [https://github.com/torvalds/linux/commit/1cc33161a83d](https://github.com/torvalds/linux/commit/1cc33161a83d)
- [https://github.com/torvalds/linux/commit/a6ca88b241d5](https://github.com/torvalds/linux/commit/a6ca88b241d5)

This field is optional, and can be `NULL` if you do not need to pass any USDT reference counters.

#### `session`

[:octicons-tag-24: 1.6.0](https://github.com/libbpf/libbpf/commit/c975e0261208c7e97592d53bbbfe6e4a1b7673dd)

Created the uprobe in session mode. In session mode, a uprobe and uretprobe are associated with each other. The uprobe can decide if the uretprobe will be triggered upon function return. Both programs also share the same cookie.

## Usage

User can specify 2 mutually exclusive set of inputs:

  1) use only `path`/`func_pattern`/`pid` arguments

  2) use `path`/`pid` with allowed combinations of
     `syms`/`offsets`/`ref_ctr_offsets`/`cookies`/`cnt`

     - `syms` and `offsets` are mutually exclusive
     - `ref_ctr_offsets` and `cookies` are optional

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
