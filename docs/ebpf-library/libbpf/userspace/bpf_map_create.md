---
title: "Libbpf userspace function 'bpf_map_create'"
description: "This page documents the 'bpf_map_create' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map_create`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/releases/tag/v0.6.0)
<!-- [/LIBBPF_TAG] -->

Low level wrapper around the [`BPF_MAP_CREATE`](../../../linux/syscall/BPF_MAP_CREATE.md) syscall command.

## Definition

`#!c int bpf_map_create(enum bpf_map_type map_type, const char *map_name, __u32 key_size, __u32 value_size, __u32 max_entries, const struct bpf_map_create_opts *opts);`

**Parameters**

- `map_type`: type of the map to create
- `map_name`: name of the map
- `key_size`: size of the key in bytes
- `value_size`: size of the value in bytes
- `max_entries`: maximum number of entries in the map
- `opts`: options for the map creation

### `struct bpf_map_create_opts`

```c
struct bpf_map_create_opts {
	size_t sz; /* size of this struct for forward/backward compatibility */

	__u32 btf_fd;
	__u32 btf_key_type_id;
	__u32 btf_value_type_id;
	__u32 btf_vmlinux_value_type_id;

	__u32 inner_map_fd;
	__u32 map_flags;
	__u64 map_extra;

	__u32 numa_node;
	__u32 map_ifindex;
	__s32 value_type_btf_obj_fd;

	__u32 token_fd;
	size_t :0;
};
```
#### `btf_fd`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/6cfb97c561adaecb3e245d2a73cb1e12f6afc115)

#### `btf_key_type_id`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/6cfb97c561adaecb3e245d2a73cb1e12f6afc115)

#### `btf_value_type_id`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/6cfb97c561adaecb3e245d2a73cb1e12f6afc115)

#### `btf_vmlinux_value_type_id`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/6cfb97c561adaecb3e245d2a73cb1e12f6afc115)

#### `inner_map_fd`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/6cfb97c561adaecb3e245d2a73cb1e12f6afc115)

#### `map_flags`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/6cfb97c561adaecb3e245d2a73cb1e12f6afc115)

#### `map_extra`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/6cfb97c561adaecb3e245d2a73cb1e12f6afc115)

#### `numa_node`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/6cfb97c561adaecb3e245d2a73cb1e12f6afc115)

#### `map_ifindex`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/commit/6cfb97c561adaecb3e245d2a73cb1e12f6afc115)

#### `value_type_btf_obj_fd`

[:octicons-tag-24: 1.4.0](https://github.com/libbpf/libbpf/commit/0b0dfaf1bebe45e9e82e60a4eaac675cfba3b1c6)

#### `token_fd`

[:octicons-tag-24: 1.5.0](https://github.com/libbpf/libbpf/commit/8002c052f3d388ce4f5c5726e1a648f19729c640)

**Return**

`>=0`, file descriptor of the created map; `<0`, on error.

## Usage

This function should only be used if you need precise control over the map creation process. For most cases, map should be created via [`bpf_object__load`](bpf_object__load.md) or similar high level APIs instead.

### Example

```c
// struct example_map{
//         __uint(type, BPF_MAP_TYPE_HASH);
//         __uint(max_entries, 10);
//         __type(key, __u32);
//         __type(value, __u32);
// };

int main(){
    int fd = bpf_map_create(
        BPF_MAP_TYPE_HASH,
        NULL,
        sizeof(__u32),
        sizeof(__u32),
        INNER_MAP_MAX_ENTRY,
        NULL
    );
    if(fd < 0)
        puts("error creating map");
}
```
