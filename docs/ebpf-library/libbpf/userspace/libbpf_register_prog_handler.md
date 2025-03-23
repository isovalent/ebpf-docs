---
title: "Libbpf userspace function 'libbpf_register_prog_handler'"
description: "This page documents the 'libbpf_register_prog_handler' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `libbpf_register_prog_handler`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)
<!-- [/LIBBPF_TAG] -->

Registers a custom BPF program [`SEC`](../ebpf/SEC.md) handler.

## Definition

`#!c int libbpf_register_prog_handler(const char *sec, enum bpf_prog_type prog_type, enum bpf_attach_type exp_attach_type, const struct libbpf_prog_handler_opts *opts);`

**Parameters**

- `sec`: section prefix for which custom handler is registered
- `prog_type`: BPF program type associated with specified section
- `exp_attach_type`: Expected BPF attach type associated with specified section
- `opts`: optional cookie, callbacks, and other extra options

**Return**

Non-negative handler ID is returned on success. This handler ID has to be passed to [`libbpf_unregister_prog_handler()`](libbpf_unregister_prog_handler.md) to unregister such custom handler. Negative error code is returned on error.

## Usage

`sec` defines which `SEC()` definitions are handled by this custom handler
registration. `sec` can have few different forms:
  - if `sec` is just a plain string (e.g., "abc"), it will match only `SEC("abc")`. If BPF program specifies `SEC("abc/whatever")` it will result in an error;
  - if `sec` is of the form "abc/", proper `SEC()` form is `SEC("abc/something")`, where acceptable "something" should be checked by `prog_init_fn` callback, if there are additional restrictions;
  - if `sec` is of the form "abc+", it will successfully match both `SEC("abc")` and `SEC("abc/whatever")` forms;
  - if `sec` is `NULL`, custom handler is registered for any BPF program that doesn't match any of the registered (custom or libbpf's own) `SEC()` handlers. There could be only one such generic custom handler registered at any given time.

All custom handlers (except the one with `sec` == `NULL`) are processed before libbpf's own `SEC()` handlers. It is allowed to "override" libbpf's `SEC()` handlers by registering custom ones for the same section prefix (i.e., it's possible to have custom `SEC("perf_event/LLC-load-misses")` handler).

!!! note
  like much of global libbpf APIs (e.g., [`libbpf_set_print`](libbpf_set_print.md), [`libbpf_set_strict_mode`](libbpf_set_strict_mode.md), etc)) these APIs are not thread-safe. User needs to ensure synchronization if there is a risk of running this API from multiple threads simultaneously.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
