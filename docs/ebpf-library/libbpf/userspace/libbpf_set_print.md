---
title: "Libbpf userspace function 'libbpf_set_print'"
description: "This page documents the 'libbpf_set_print' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `libbpf_set_print`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Sets user-provided log callback function to be used for libbpf warnings and informational messages. If the user callback is not set, messages are logged to `stderr` by default. The verbosity of these messages can be controlled by setting the environment variable `LIBBPF_LOG_LEVEL` to either `warn`, `info`, or `debug`.

## Definition

```c
typedef int (*libbpf_print_fn_t)(enum libbpf_print_level level, const char *, va_list ap);

libbpf_print_fn_t libbpf_set_print(libbpf_print_fn_t fn);
```

**Parameters**

- `fn`: The log print function. If `NULL`, libbpf won't print anything.

**Return**

Pointer to old print function.

## Usage

This function is thread-safe.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
