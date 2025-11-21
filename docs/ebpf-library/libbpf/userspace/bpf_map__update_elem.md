---
title: "Libbpf userspace function 'bpf_map__update_elem'"
description: "This page documents the 'bpf_map__update_elem' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__update_elem`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)
<!-- [/LIBBPF_TAG] -->

Allows to insert or update value in BPF map that corresponds to provided key.

## Definition

`#!c int bpf_map__update_elem(const struct bpf_map *map, const void *key, size_t key_sz, const void *value, size_t value_sz, __u64 flags);`

**Parameters**

- `map`: BPF map to insert to or update element in
- `key`: pointer to memory containing bytes of the key
- `key_sz`: size in bytes of key data, needs to match BPF map definition's `key_size`
- `value`: pointer to memory containing bytes of the value
- `value_sz`: size in byte of value data memory; it has to match BPF map definition's `value_size`. For per-CPU BPF maps value size has to be a product of BPF map value size and number of possible CPUs in the system (could be fetched with [`libbpf_num_possible_cpus`](libbpf_num_possible_cpus.md)). Note also that for per-CPU values value size has to be aligned up to closest 8 bytes for alignment reasons, so expected size is: `round_up(value_size, 8) * libbpf_num_possible_cpus()`.
- `flags`: flags passed to kernel for this operation

**Return**

`0`, on success; negative error, otherwise

## Usage

[`bpf_map__update_elem()`](bpf_map__update_elem.md) is high-level equivalent of [`bpf_map_update_elem()`](bpf_map_update_elem.md) API with added check for key and value size.

### Example

```c
#include <stdio.h>
#include "log_syscalls.skel.h"

#define STRING_LEN 128

// struct {
//     __uint(type, BPF_MAP_TYPE_HASH);
//     __type(key, char[STRING_LEN]);
//     __type(value, int);
//     __uint(max_entries, 32);
// } syscall_count_map SEC(".maps");

const int zero = 0;
int main(){

    char program_to_log[STRING_LEN] = "cat";

    struct log_syscalls *skel = log_syscalls__open_and_load();

    int err = bpf_map__update_elem(
    skel->maps.syscall_count_map,
        program_to_log,
        STRING_LEN,
        &zero,
        sizeof(int),
        BPF_NOEXIST
    );
    if(err){
        printf("err inserting value: %s\n", strerror(errno));
        return 1;
    }

    log_syscalls__attach(skel);
    while(1){};
}
```
