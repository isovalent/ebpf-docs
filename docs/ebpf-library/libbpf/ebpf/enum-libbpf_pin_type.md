---
title: "Libbpf eBPF type 'enum libbpf_pin_type'"
description: "This page documents the 'enum libbpf_pin_type' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `enum libbpf_pin_type`

[:octicons-tag-24: v0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)

The `enum libbpf_pin_type` enum specifies valid values for the `pinning` BTF map property.

## Definition

```c
enum libbpf_pin_type {
	LIBBPF_PIN_NONE,
	/* PIN_BY_NAME: pin maps by name (in /sys/fs/bpf by default) */
	LIBBPF_PIN_BY_NAME,
};
```

## Usage

This type defines the valid values for the `pinning` BTF map property. The `pinning` property is used to specify how the BPF map should be pinned in the filesystem. The `LIBBPF_PIN_NONE` value indicates that the map should not be pinned, while the `LIBBPF_PIN_BY_NAME` value indicates that the map should be pinned by name in the `/sys/fs/bpf` directory.

### Example

```c hl_lines="6"
struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, 1024);
    __uint(key_size, sizeof(int));
    __uint(value_size, sizeof(long));
    __uint(pinning, LIBBPF_PIN_BY_NAME);
} SEC(".maps") my_map;
```
