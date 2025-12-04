---
title: "Map Type 'BPF_MAP_TYPE_HASH_OF_MAPS'"
description: "This page documents the 'BPF_MAP_TYPE_HASH_OF_MAPS' eBPF map type, including its definition, usage, program types that can use it, and examples."
---
# Map type `BPF_MAP_TYPE_HASH_OF_MAPS`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_HASH_OF_MAPS) -->
[:octicons-tag-24: v4.12](https://github.com/torvalds/linux/commit/bcc6b1b7ebf857a9fe56202e2be3361131588c15)
<!-- [/FEATURE_TAG] -->

The hash of maps map type contains references to other maps.

# Usage
This map type is a map-in-map type. The map values contain references to other BPF maps. We will refer to map-in-map as the "outer map" and the maps referenced as the "inner map(s)". The key advantage of using a map-in-map is that the outer map is directly referenced by any programs that use it, but the inner maps are not.

Users should be aware of the read/write asymmetry of this map type:

* The [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md) syscall command takes *file descriptor* of the BPF map you wish to insert into the map.
* The [`bpf_map_lookup_elem`](../helper-function/bpf_map_lookup_elem.md) returns a pointer to the inner map or `NULL`. This pointer can be used for any other helper that takes map pointers.

# Attributes
The [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) must always be `4` indicating a 32-bit unsigned integer.

The [`inner_map_fd`](../syscall/BPF_MAP_CREATE.md#inner_map_fd) attribute must be set to the file descriptor of another map. This other map will serve as a template for the inner maps. After loading, during insertion of values, the kernel will verify that the spec of the inner map values you are attempting to insert match the spec of the map provided by this field. The map used to indicate the type is not linked to the map-in-map type in any way, it is just used to transfer type info. A common technique for loaders is to build a temporary map just for the purpose of providing the type info and freeing that map as soon as the outer map has been created.

# Example
The following bpf program logs different syscall invokes for different user ID's, in this case we logged `openat` and `chdir` syscalls for user ID 1000, and `chmod` and `bind` for user ID 0.
```c
//tracer.bpf.c

#include "vmlinux.h"
#include <bpf/bpf_tracing.h>

struct {
        __uint(type, BPF_MAP_TYPE_HASH_OF_MAPS);
        __uint(max_entries, 8);
        __type(key, u32);
        __type(value, u32);
} outer_map_instance SEC(".maps");


struct sys_enter_ctx {
    struct trace_entry ent;
    long id;
    unsigned long args[6];
};

SEC("tp/raw_syscalls/sys_enter")
int trace_users(struct sys_enter_ctx *ctx){
	u32 uid = bpf_get_current_uid_gid() & 0xffffffff;
	void *this_uid_map = bpf_map_lookup_elem(&outer_map_instance, &uid);
	if(this_uid_map == NULL)
		return 1;
	
    long id = ctx->id;
	void *do_log = bpf_map_lookup_elem(this_uid_map, &id);
	if(do_log == NULL || *((int *)do_log) == 0){
		return 1;
    }
		
    bpf_printk("syscall %d called by uid %d", ctx->id, uid);

    return 0;
}

char __license[] SEC("license") = "GPL";

```

```c
//loader.c

#include <linux/types.h>
#include <bpf/bpf.h>
#include <sys/syscall.h>
#include <errno.h>
#include <stdio.h>
#include <unistd.h>

#include "tracer.bpf.skel.h"

// we will create this map in userspace
// struct inner_map {
//         __uint(type, BPF_MAP_TYPE_HASH);
//         __uint(max_entries, 10);
//         __type(key, __u32);
//         __type(value, __u32);
// };

#define INNER_MAP_MAX_ENTRY 10
const int one = 1;

int create_template_map(){
    int fd = bpf_map_create(
        BPF_MAP_TYPE_HASH,
        NULL,
        sizeof(__u32),
        sizeof(__u32),
        INNER_MAP_MAX_ENTRY,
        NULL
    );
    return fd;
}

int create_and_init_example_inner_map(int *syscall_ids_to_track, char *name){
    int fd = bpf_map_create(
        BPF_MAP_TYPE_HASH,
        name,
        sizeof(__u32),
        sizeof(__u32),
        INNER_MAP_MAX_ENTRY,
        NULL
    );

    for(int i = 0; i < INNER_MAP_MAX_ENTRY; i++){
        bpf_map_update_elem(fd, &syscall_ids_to_track[i], &one, BPF_ANY);
    }
    return fd;
}

int main(){
    int uid;
    int syscall_ids_to_track[INNER_MAP_MAX_ENTRY];
    memset(syscall_ids_to_track, -1, INNER_MAP_MAX_ENTRY * sizeof(int));
    int fd = create_template_map();

    struct tracer_bpf *skel = tracer_bpf__open();

    bpf_map__set_inner_map_fd(skel->maps.outer_map_instance, fd);

    tracer_bpf__load(skel);
    close(fd); //we don't need template map fd anymore

    tracer_bpf__attach(skel);

    uid = 1000;
    syscall_ids_to_track[0] = SYS_openat;
    syscall_ids_to_track[1] = SYS_chdir;
    fd = create_and_init_example_inner_map(syscall_ids_to_track, "uid_1000_example");

    bpf_map_update_elem(
        bpf_map__fd(skel->maps.outer_map_instance), 
        &uid,
        &fd,
        BPF_NOEXIST
    );
    close(fd);

    uid = 0;
    syscall_ids_to_track[0] = SYS_chmod;
    syscall_ids_to_track[1] = SYS_bind;
    fd = create_and_init_example_inner_map(syscall_ids_to_track, "uid_0_example");

    bpf_map_update_elem(
        bpf_map__fd(skel->maps.outer_map_instance), 
        &uid,
        &fd,
        BPF_NOEXIST
    );
    close(fd);

    puts("looping forever");
    while(1){}

    return 0;
}
```
