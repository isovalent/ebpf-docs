---
title: "Libbpf userspace function 'bpf_program__set_attach_target'"
description: "This page documents the 'bpf_program__set_attach_target' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__set_attach_target`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.8](https://github.com/libbpf/libbpf/releases/tag/v0.0.8)
<!-- [/LIBBPF_TAG] -->

Sets BTF-based attach target for supported BPF program types:

- BTF-aware raw tracepoints ([`tp_btf`](../../../linux/program-type/BPF_PROG_TYPE_TRACING.md#raw-tracepoint));   
- [`fentry`](../../../linux/program-type/BPF_PROG_TYPE_TRACING.md#fentry)/[`fexit`](../../../linux/program-type/BPF_PROG_TYPE_TRACING.md#fexit)/[`fmod_ret`](../../../linux/program-type/BPF_PROG_TYPE_TRACING.md#modify-return);   
- [`lsm`](../../../linux/program-type/BPF_PROG_TYPE_LSM.md)   
- [`freplace`](../../../linux/program-type/BPF_PROG_TYPE_EXT.md)

## Definition

`#!c int bpf_program__set_attach_target(struct bpf_program *prog, int attach_prog_fd, const char *attach_func_name);`

**Parameters**

- `prog`: BPF program to set the attach type for
- `attach_prog_fd`: the file descriptor of the BPF program to attach to, `0` if not attaching to another BPF program
- `attach_func_name`: the name of the function to attach to

**Return**

error code; or 0 if no error occurred.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
