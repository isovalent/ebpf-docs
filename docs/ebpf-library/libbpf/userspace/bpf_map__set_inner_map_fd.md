---
title: "Libbpf userspace function 'bpf_map__set_inner_map_fd'"
description: "This page documents the 'bpf_map__set_inner_map_fd' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__set_inner_map_fd`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.1](https://github.com/libbpf/libbpf/releases/tag/v0.0.1)
<!-- [/LIBBPF_TAG] -->

Set the file descriptor of an inner map for a map.

## Definition

`#!c int bpf_map__set_inner_map_fd(struct bpf_map *map, int fd);`

**Parameters**

- `map`: The bpf_map
- `fd`: The file descriptor of the inner map

**Return**

`0`, on success; negative error, otherwise

## Usage

When loading map-in-map maps, such as [`BPF_MAP_TYPE_ARRAY_OF_MAPS`](../../../linux/map-type/BPF_MAP_TYPE_ARRAY_OF_MAPS.md) or [`BPF_MAP_TYPE_HASH_OF_MAPS`](../../../linux/map-type/BPF_MAP_TYPE_HASH_OF_MAPS.md), the verifier needs what sort of maps you will be putting into it. To communicate that, a map with the same attributes as will be used as values must be loaded first, and then its file descriptor passed to the outer map before loading.

### Example

```c
// outer map
// struct {
//         __uint(type, BPF_MAP_TYPE_HASH_OF_MAPS);
//         __uint(max_entries, 8);
//         __type(key, u32);
//         __type(value, u32);
// } outer_map SEC(".maps");

// we will create this map in user space
// struct inner_map {
//         __uint(type, BPF_MAP_TYPE_HASH);
//         __uint(max_entries, 10);
//         __type(key, __u32);
//         __type(value, __u32);
// };

int create_template_map(){
    int fd = bpf_map_create(
        BPF_MAP_TYPE_HASH,
        NULL,
        sizeof(__u32),
        sizeof(__u32),
        10,
        NULL
    );
    return fd;
}
int main(){
    struct tracer_bpf *skel = tracer_bpf__open();

    int fd = create_template_map();

    bpf_map__set_inner_map_fd(skel->maps.outer_map, fd);

    tracer_bpf__load(skel);
    close(fd); //we don't need template map fd anymore

    tracer_bpf__attach(skel);

    puts("looping forever");
    while(1){}

    return 0;
}

```
