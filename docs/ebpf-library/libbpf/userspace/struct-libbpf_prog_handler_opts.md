---
title: "Libbpf userspace type 'struct libbpf_prog_handler_opts'"
description: "This page documents the 'struct libbpf_prog_handler_opts' libbpf userspace type, including its definition, usage, and examples."
---
# Libbpf userspace function `struct libbpf_prog_handler_opts`

[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)

Gives read-only access to BPF program's underlying BPF instructions.

## Definition

```c
struct libbpf_prog_handler_opts {
	size_t  [sz](#sz);
	long    [cookie](#cookie);
    
	libbpf_prog_setup_fn_t          [prog_setup_fn](#prog_setup_fn);
	libbpf_prog_prepare_load_fn_t   [prog_prepare_load_fn](#prog_prepare_load_fn);
	libbpf_prog_attach_fn_t         [prog_attach_fn](#prog_attach_fn);
};
```

### `sz`

Size of this struct, for forward/backward compatibility.

### `cookie`

[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)

User-provided value that is passed to [`prog_setup_fn`](#prog_setup_fn), [`prog_prepare_load_fn`](#prog_prepare_load_fn), and [`prog_attach_fn`](#prog_attach_fn) callbacks. Allows user to register one set of callbacks for multiple [`SEC`](../ebpf/SEC.md) definitions and still be able to distinguish them, if necessary. 

For example, libbpf itself is using this to pass necessary flags (e.g., [`sleepable`](../../../linux/syscall/BPF_PROG_LOAD.md#bpf_f_sleepable) flag) to a common internal [`SEC`](../ebpf/SEC.md) handler.

### `prog_setup_fn`

[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)

BPF program initialization callback. Callback is optional, pass `NULL` if it's not necessary. Called during [`bpf_object__open`](bpf_object__open.md) for each recognized BPF program. Callback can use various `bpf_program__set_*()` setters to adjust whatever properties are necessary.

```c
typedef int (*libbpf_prog_setup_fn_t)(struct bpf_program *prog, long cookie);
```

### `prog_prepare_load_fn`

[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)

BPF program loading callback. Callback is optional, pass `NULL` if it's not necessary. Called right before libbpf performs [`bpf_prog_load`](bpf_prog_load.md) to load BPF program into the kernel. Callback can adjust opts as necessary.

```c
typedef int (*libbpf_prog_prepare_load_fn_t)(struct bpf_program *prog,
					     struct bpf_prog_load_opts *opts, long cookie);
```

### `prog_attach_fn`

[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)

BPF program attach callback. Callback is optional, pass `NULL` if it's not necessary. Called during skeleton attach or through [`bpf_program__attach`](bpf_program__attach.md). If auto-attach is not supported, callback should return `0` and set link to `NULL` (it's not considered an error during skeleton attach, but it will be an error for [`bpf_program__attach`](bpf_program__attach.md) calls). 

On error, error should be returned directly and link set to `NULL`. On success, return `0` and set link to a valid `struct bpf_link`.

```c
typedef int (*libbpf_prog_attach_fn_t)(const struct bpf_program *prog, long cookie,
				       struct bpf_link **link);
```

## Usage

This struct is passed as argument to [`libbpf_register_prog_handler`](libbpf_register_prog_handler.md). It holds callbacks to implement custom handling for programs in a certain ELF section.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
