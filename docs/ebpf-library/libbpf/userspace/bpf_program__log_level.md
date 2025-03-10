---
title: "Libbpf userspace function 'bpf_program__log_level'"
description: "This page documents the 'bpf_program__log_level' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__log_level`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.7.0](https://github.com/libbpf/libbpf/releases/tag/v0.7.0)
<!-- [/LIBBPF_TAG] -->

Get the verifier log level with which the BPF program was or will be loaded.

## Definition

`#!c __u32 bpf_program__log_level(const struct bpf_program *prog);`

**Parameters**

- `prog`: BPF program to get the verifier log level of.

**Return**

The verifier log level of the BPF program.

```c
#define BPF_LOG_LEVEL1	1
#define BPF_LOG_LEVEL2	2
#define BPF_LOG_STATS	4
#define BPF_LOG_FIXED	8
```

The log level is considered one value, so `BPF_LOG_LEVEL1` or `BPF_LOG_LEVEL2`. `BPF_LOG_STATS` and `BPF_LOG_FIXED` are flags that can be combined with the log level.

### `BPF_LOG_LEVEL1`

Print every instruction and verifier state at branch points.

### `BPF_LOG_LEVEL2`

[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/06ee7115b0d1742de745ad143fb5e06d77d27fba)

Print every instruction and verifier state at every instruction

### `BPF_LOG_STATS`

[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/06ee7115b0d1742de745ad143fb5e06d77d27fba)

Print verifier error and stats at the end of verification

### `BPF_LOG_FIXED`

[:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/1216640938035e63bdbd32438e91c9bcc1fd8ee1)

Since kernel v6.4, the verifier log automatically rotates when it reaches the buffer size, writing over the start of the log, so the last lines are always in the log. 

This flag disables the automatic rotation, so we see the start of the log until the buffer is full.

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
