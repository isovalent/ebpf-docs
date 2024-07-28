---
title: "Syscall command 'BPF_OBJ_GET_INFO_BY_FD'"
description: "This page documents the 'BPF_OBJ_GET_INFO_BY_FD' eBPF syscall command, including its definition, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_OBJ_GET_INFO_BY_FD` command

<!-- [FEATURE_TAG](BPF_OBJ_GET_INFO_BY_FD) -->
[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/1e270976908686ec25fb91b8a34145be54137976)
<!-- [/FEATURE_TAG] -->

This syscall command is used to get information about BPF objects.

## Return type

This command will return `0` on success or a error number (negative integer) if something went wrong.

## Usage

This syscall command returns information about the BPF object indicated by `bpf_fd`. The structure of the returned information is dependant on the object type. More on this in the [info structures](#info-structures) section.

This syscall is typically used by inspection tools to get information about loaded objects that you would normally "know" if you were the author or loader of a given object. 

This command can also be used in a monitoring or benchmarking scenario in combination with the [`BPF_ENABLE_STATS`](BPF_ENABLE_STATS.md) syscall command.

## Attributes

??? abstract "C structure"
    ```c
	struct {
		__u32		bpf_fd;
		__u32		info_len;
		__aligned_u64	info;
	};
    ```

### `bpf_fd`

This field indicates the file descriptor of the BPF object for which we want to request the information.

### `info_len`

This field indicates the size of the buffer which `info` points to and will be changed by the syscall command to the actual amount of data written to the info buffer.

### `info`

This field indicates a memory region to which the kernel will write the requested information, the structure of which is dependant on the object type to which `bpf_fd` refers (see [info structures](#info-structures)). This field should be a pointer.

## Info structures

### `struct bpf_prog_info`

??? abstract "C structure"
    ```c
    struct bpf_prog_info {
        __u32 type;
        __u32 id;
        __u8  tag[BPF_TAG_SIZE];
        __u32 jited_prog_len;
        __u32 xlated_prog_len;
        __aligned_u64 jited_prog_insns;
        __aligned_u64 xlated_prog_insns;
        __u64 load_time;	/* ns since boottime */
        __u32 created_by_uid;
        __u32 nr_map_ids;
        __aligned_u64 map_ids;
        char name[BPF_OBJ_NAME_LEN];
        __u32 ifindex;
        __u32 gpl_compatible:1;
        __u32 :31; /* alignment pad */
        __u64 netns_dev;
        __u64 netns_ino;
        __u32 nr_jited_ksyms;
        __u32 nr_jited_func_lens;
        __aligned_u64 jited_ksyms;
        __aligned_u64 jited_func_lens;
        __u32 btf_id;
        __u32 func_info_rec_size;
        __aligned_u64 func_info;
        __u32 nr_func_info;
        __u32 nr_line_info;
        __aligned_u64 line_info;
        __aligned_u64 jited_line_info;
        __u32 nr_jited_line_info;
        __u32 line_info_rec_size;
        __u32 jited_line_info_rec_size;
        __u32 nr_prog_tags;
        __aligned_u64 prog_tags;
        __u64 run_time_ns;
        __u64 run_cnt;
        __u64 recursion_misses;
        __u32 verified_insns;
        __u32 attach_btf_obj_id;
        __u32 attach_btf_id;
    } __attribute__((aligned(8)));
    ```

#### `type`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/1e270976908686ec25fb91b8a34145be54137976)

This field indicates the type of the program, values is one of [`program types`](../program-type/index.md).

#### `id`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/1e270976908686ec25fb91b8a34145be54137976)

This field indicates the unique ID of the program, as seen in [`BPF_PROG_GET_NEXT_ID`](BPF_PROG_GET_NEXT_ID.md)

#### `tag`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/1e270976908686ec25fb91b8a34145be54137976)

This field indicates the tag of the program. A tag is a hash of the program instructions. Its main purpose is to check compare programs, like a sort of checksum. It is explicitly not meant as a security feature, hence the small size of 8 bytes. [Commit#f1f7714e](https://github.com/torvalds/linux/commit/f1f7714ea51c56b7163fb1a5acf39c6a204dd758)

#### `jited_prog_len`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/1e270976908686ec25fb91b8a34145be54137976)

This field indicates the length of the buffer provided by `jited_prog_insns` when calling the command and will be set to the actual number of bytes available after calling the command.

#### `xlated_prog_len`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/1e270976908686ec25fb91b8a34145be54137976)

This field indicates the length of the buffer provided by `xlated_prog_len` when calling the command and will be set to the actual number of bytes available after calling the command.

#### `jited_prog_insns`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/1e270976908686ec25fb91b8a34145be54137976)

This field indicates a memory area where the kernel can write the JIT-ed program instructions to. This field should be a pointer to a memory buffer.

The JIT-ed instructions are machine code in the host architecture, the actual code that runs on the CPU.

#### `xlated_prog_insns`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/1e270976908686ec25fb91b8a34145be54137976)

This field indicates a memory area where the kernel can write the <nospell>xlated</nospell>/translated program instructions to. This field should be a pointer to a memory buffer.

The translated program instructions are still use the eBPF instruction set, but have been modified by the verifier.

* Some helper calls are inlined
* IMM64 instructions with file descriptors are changed to actual pointers
* Accesses into ctx are rewritten
* Some helper calls like map access are now specialized

This list is not exhaustive.

#### `load_time`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/cb4d2b3f03d8eed90be3a194e5b54b734ec4bbe9)

This field indicates when the program was loaded in nanoseconds since boot time.

#### `created_by_uid`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/cb4d2b3f03d8eed90be3a194e5b54b734ec4bbe9)

This field indicates the User ID of the process who loaded the program.

#### `nr_map_ids`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/cb4d2b3f03d8eed90be3a194e5b54b734ec4bbe9)

This field indicates the number of map IDs that were written to `map_ids`.

#### `map_ids`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/cb4d2b3f03d8eed90be3a194e5b54b734ec4bbe9)

This field indicates the list of maps that are used by this program directly. Its value should be a pointer to an array of 32-bit map IDs.

#### `name`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/067cae47771c864604969fd902efe10916e0d79c)

This field indicates the name of the program which was set via the [`prog_name`](BPF_PROG_LOAD.md#prog_name) attribute while loading.

#### `ifindex`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/675fc275a3a2d905535207237402c6d8dcb5fa4b)

This field indicates the network interface index of the device on which this program is offloaded. Or `0` if the program is not offloaded.

#### `gpl_compatible`

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/b85fab0e67b162014cd328cb4e2a8e8ae382cb8a)

This field indicates if the program was loaded with a GPL compatible [`license`](BPF_PROG_LOAD.md#license).

#### `netns_dev`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/675fc275a3a2d905535207237402c6d8dcb5fa4b)

This field indicates the device number of the device on which this program is offloaded. Or `0` if the program is not offloaded.

#### `netns_ino`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/675fc275a3a2d905535207237402c6d8dcb5fa4b)

This field indicates the inode number of the device on which this program is offloaded. Or `0` if the program is not offloaded.

#### `nr_jited_ksyms`

This field indicates the number of kernel symbols available via `jited_ksyms`

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/dbecd7388476aedeb66389febea84d5450d28773)

#### `nr_jited_func_lens`

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/815581c11cc29f74af252b6306ea1ec94160231a)

This field indicates the number of entries available via `jited_func_lens`

#### `jited_ksyms`

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/dbecd7388476aedeb66389febea84d5450d28773)

This field indicates a list of 64-bit memory addresses. The value of this field should be a pointer to an array of 64-bit numbers.

These memory addresses can be translated to kernel symbols using the `/prog/kallsyms` mapping.

#### `jited_func_lens`

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/815581c11cc29f74af252b6306ea1ec94160231a)

This field indicates a list of function lengths. This value should be a pointer to an array of 32-bit numbers.

These function lengths can be used in combination with `jited_ksyms` and `jited_prog_insns` to separate the single blob of instructions back into separate BPF-to-BPF functions.

#### `btf_id`

[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/838e96904ff3fc6c30e5ebbc611474669856e3c0)

This field indicates the id of the BTF object which is associated with this program.

#### `func_info_rec_size`

[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/838e96904ff3fc6c30e5ebbc611474669856e3c0)

This field indicates the size of the records available via `func_info`

#### `func_info`

[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/838e96904ff3fc6c30e5ebbc611474669856e3c0)

This field indicates a list of function info records. This should be a pointer to an array of function info blobs the size of `func_info_rec_size` with `nr_func_info` elements.

These are the function info records supplied during program loading via [`func_info`](BPF_PROG_LOAD.md#func_info). And are aligned to the `xlated_prog_insns`.

#### `nr_func_info`

[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/11d8b82d2222cade12caad2c125f23023777dcbc)

This field indicates the amount of function info records are available via `func_info`.

#### `nr_line_info`

[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/11d8b82d2222cade12caad2c125f23023777dcbc)

This field indicates the amount of function info records are available via `line_info`.

#### `line_info`

[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/c454a46b5efd8eff8880e88ece2976e60a26bf35)

This field indicates a list of function info records. This should be a pointer to an array of function info blobs the size of `line_info_rec_size` with `nr_line_info` elements.

These are the line info records supplied during program loading via [`line_info`](BPF_PROG_LOAD.md#line_info). And are aligned to the `xlated_prog_insns`.

#### `jited_line_info`

[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/c454a46b5efd8eff8880e88ece2976e60a26bf35)

This field indicates a list of function info records. This should be a pointer to an array of function info blobs the size of `jited_line_info_rec_size` with `nr_jited_line_info` elements.

These are the line info records supplied during program loading via [`line_info`](BPF_PROG_LOAD.md#line_info). And are aligned to the `jitted_prog_insns`.

#### `nr_jited_line_info`

[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/11d8b82d2222cade12caad2c125f23023777dcbc)

This field indicates the number of line info records available via `jited_line_info`

#### `line_info_rec_size`

[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/c454a46b5efd8eff8880e88ece2976e60a26bf35)

This field indicates the size of the line info records available via `line_info`

#### `jited_line_info_rec_size`

[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/c454a46b5efd8eff8880e88ece2976e60a26bf35)

This field indicates the size of the line info records available via `jited_line_info`

#### `nr_prog_tags`

[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/c872bdb38febb4c31ece3599c52cf1f833b89f4e)

This field indicates the number of sub-program tags in `prog_tags`

#### `prog_tags`

[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/c872bdb38febb4c31ece3599c52cf1f833b89f4e)

This field indicates a list of tags for the sub-programs. This value should be a pointer to an array of 8-byte tags.

These are tags for the sub-programs/bpf-to-bpf functions. Their value is derived the same way as the `tag` field.

#### `run_time_ns`

[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/5f8f8b93aeb8371c54af08bece2bd04bc2d48707)

This field indicates the total amount of nanoseconds this program has been running accumulatively. This field is only updated by the kernel if when enabled via the [`BPF_ENABLE_STATS`](BPF_ENABLE_STATS.md) syscall command.

#### `run_cnt`

[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/5f8f8b93aeb8371c54af08bece2bd04bc2d48707)

This field indicates the total amount of times this program has been called/executed accumulatively. This field is only updated by the kernel if when enabled via the [`BPF_ENABLE_STATS`](BPF_ENABLE_STATS.md) syscall command.

#### `recursion_misses`

[:octicons-tag-24: v5.12](https://github.com/torvalds/linux/commit/9ed9e9ba2337205311398a312796c213737bac35)

This field indicates how often the "recursion prevention" mechanism has kicked in.

This mechanism was introduced in this [commit](https://github.com/torvalds/linux/commit/ca06f55b90020cd97f4cc6d52db95436162e7dcf). Its purposes seems to be to prevent the case where an [fentry/fexit tracing program starts code paths it uses itself](https://lore.kernel.org/bpf/20210210033634.62081-6-alexei.starovoitov@gmail.com/) which cause cause issue.

#### `verified_insns`

[:octicons-tag-24: v5.16](https://github.com/torvalds/linux/commit/aba64c7da98330141dcdadd5612f088043a83696)

This field indicates the amount of verified instructions. This is the same number as would be logged by the verifier when initially loading the program.

#### `attach_btf_obj_id`

[:octicons-tag-24: v6.0](https://github.com/torvalds/linux/commit/b79c9fc9551b45953a94abf550b7bd3b00e3a0f9)

This field indicates the id of the BTF object which contains the type indicated by `attach_btf_id` to which this program is attached.

#### `attach_btf_id`

[:octicons-tag-24: v6.0](https://github.com/torvalds/linux/commit/b79c9fc9551b45953a94abf550b7bd3b00e3a0f9)

This field indicates the type id of the BTF function type to which this program is attached. Currently used for `BPF_LSM_CGROUP` programs.

### `struct bpf_map_info`

??? abstract "C structure"
    ```c
    struct bpf_map_info {
        __u32 type;
        __u32 id;
        __u32 key_size;
        __u32 value_size;
        __u32 max_entries;
        __u32 map_flags;
        char  name[BPF_OBJ_NAME_LEN];
        __u32 ifindex;
        __u32 btf_vmlinux_value_type_id;
        __u64 netns_dev;
        __u64 netns_ino;
        __u32 btf_id;
        __u32 btf_key_type_id;
        __u32 btf_value_type_id;
        __u32 :32;	/* alignment pad */
        __u64 map_extra;
    } __attribute__((aligned(8)));
    ```

#### `type`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/1e270976908686ec25fb91b8a34145be54137976)

This field indicates the map type which should be one of [map types](../map-type/index.md)

#### `id`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/1e270976908686ec25fb91b8a34145be54137976)

This field indicates the unique ID of the program, as seen in [`BPF_MAP_GET_NEXT_ID`](BPF_MAP_GET_NEXT_ID.md)

#### `key_size`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/1e270976908686ec25fb91b8a34145be54137976)

This field indicates the size of the map key in bytes.

#### `value_size`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/1e270976908686ec25fb91b8a34145be54137976)

This field indicates the size of the map value in bytes

#### `max_entries`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/1e270976908686ec25fb91b8a34145be54137976)

This field indicates the maximum amount of elements that the map can hold.

#### `map_flags`

[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/1e270976908686ec25fb91b8a34145be54137976)

This field indicates the flags with which the map has been loaded, see [flags](BPF_MAP_CREATE.md#flags) for possible values.

#### `name`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/067cae47771c864604969fd902efe10916e0d79c)

This field indicates the name of the map given to it when it was created.

#### `ifindex`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/52775b33bb5072fbc07b02c0cf4fe8da1f7ee7cd)

This field indicates the network interface index of the network interface onto which the map has been offloaded.

#### `btf_vmlinux_value_type_id`

[:octicons-tag-24: v5.6](https://github.com/torvalds/linux/commit/85d33df357b634649ddbe0a20fd2d0fc5732c3cb)

This field indicates the type id of the struct which the values of the map are replacing the function pointers. This field is specifically used for map of type [`BPF_MAP_TYPE_STRUCT_OPS`](../map-type/BPF_MAP_TYPE_STRUCT_OPS.md)

#### `netns_dev`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/52775b33bb5072fbc07b02c0cf4fe8da1f7ee7cd)

This field indicates the device number of the network device onto which the maps has been offloaded.

#### `netns_ino`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/52775b33bb5072fbc07b02c0cf4fe8da1f7ee7cd)

This field indicates the inode number of the network device onto which the maps has been offloaded.

#### `btf_id`

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/78958fca7ead2f81b60a6827881c4866d1ed0c52)

This field indicates the ID of the BTF object which contain the types for `btf_key_type_id` and `btf_value_type_id`.

#### `btf_key_type_id`

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/9b2cf328b2eccf761537a06bef914d2a0700fba7)

This field indicates the BTF type id representing the keys of this map.

#### `btf_value_type_id`

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/9b2cf328b2eccf761537a06bef914d2a0700fba7)

This field indicates the BTF type id representing the values of this map.

#### `map_extra`

[:octicons-tag-24: v5.16](https://github.com/torvalds/linux/commit/9330986c03006ab1d33d243b7cfe598a7a3c1baa)

This field indicates any extra settings a map might have. Currently this is only used for [`BPF_MAP_TYPE_BLOOM_FILTER`](../map-type/BPF_MAP_TYPE_BLOOM_FILTER.md) maps.

### `struct bpf_btf_info`


??? abstract "C structure"
    ```c
    struct bpf_btf_info {
        __aligned_u64 btf;
        __u32 btf_size;
        __u32 id;
        __aligned_u64 name;
        __u32 name_len;
        __u32 kernel_btf;
    } __attribute__((aligned(8)));
    ```

#### `btf`

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/62dab84c81a487d946a5fc37c6df541dd95cca38)

This field indicates the serialized BTF information. This value should be a pointer to a memory buffer where the kernel will write the BTF to.

#### `btf_size`

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/62dab84c81a487d946a5fc37c6df541dd95cca38)

This field indicates the size of `btf` in bytes.

#### `id`

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/62dab84c81a487d946a5fc37c6df541dd95cca38)

This field indicates the unique ID of the BTF object.

#### `name`

[:octicons-tag-24: v5.11](https://github.com/torvalds/linux/commit/5329722057d41aebc31e391907a501feaa42f7d9)

This field indicates the name of the BTF object. Which in the case of the kernels builtin BTF will be "vmlinux" and may be different for the BTF objects of kernel modules.

#### `name_len`

[:octicons-tag-24: v5.11](https://github.com/torvalds/linux/commit/5329722057d41aebc31e391907a501feaa42f7d9)

This field indicates the length of the `name` field.

#### `kernel_btf`

[:octicons-tag-24: v5.11](https://github.com/torvalds/linux/commit/5329722057d41aebc31e391907a501feaa42f7d9)

This field indicates if the BTF is the kernels own BTF (vmlinux, `1`) or a user loaded BTF blob (`0`).

### `struct bpf_link_info`

??? abstract "C structure"
    ```c
    struct bpf_link_info {
        __u32 type;
        __u32 id;
        __u32 prog_id;
        union {
            struct {
                __aligned_u64 tp_name; /* in/out: tp_name buffer ptr */
                __u32 tp_name_len;     /* in/out: tp_name buffer len */
            } raw_tracepoint;
            struct {
                __u32 attach_type;
                __u32 target_obj_id; /* prog_id for PROG_EXT, otherwise btf object id */
                __u32 target_btf_id; /* BTF type id inside the object */
            } tracing;
            struct {
                __u64 cgroup_id;
                __u32 attach_type;
            } cgroup;
            struct {
                __aligned_u64 target_name; /* in/out: target_name buffer ptr */
                __u32 target_name_len;	   /* in/out: target_name buffer len */

                /* If the iter specific field is 32 bits, it can be put
                * in the first or second union. Otherwise it should be
                * put in the second union.
                */
                union {
                    struct {
                        __u32 map_id;
                    } map;
                };
                union {
                    struct {
                        __u64 cgroup_id;
                        __u32 order;
                    } cgroup;
                    struct {
                        __u32 tid;
                        __u32 pid;
                    } task;
                };
            } iter;
            struct  {
                __u32 netns_ino;
                __u32 attach_type;
            } netns;
            struct {
                __u32 ifindex;
            } xdp;
        };
    } __attribute__((aligned(8)));
    ```

#### `type`

[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/f2e10bff16a0fdd41ba278c84da9813700e356af)

This field indicates the link type, which is one of:

* `BPF_LINK_TYPE_RAW_TRACEPOINT` = `1`
* `BPF_LINK_TYPE_TRACING` = `2`
* `BPF_LINK_TYPE_CGROUP` = `3`
* `BPF_LINK_TYPE_ITER` = `4`
* `BPF_LINK_TYPE_NETNS` = `5`
* `BPF_LINK_TYPE_XDP` = `6`
* `BPF_LINK_TYPE_PERF_EVENT` = `7`
* `BPF_LINK_TYPE_KPROBE_MULTI` = `8`

#### `id`

[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/f2e10bff16a0fdd41ba278c84da9813700e356af)

This field indicates the unique ID of the link object.

#### `prog_id`

[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/f2e10bff16a0fdd41ba278c84da9813700e356af)

This field indicates the ID of the program which is linked.

#### `raw_tracepoint`

These fields apply for the `BPF_LINK_TYPE_RAW_TRACEPOINT` link type.

##### `tp_name`

[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/f2e10bff16a0fdd41ba278c84da9813700e356af)

This field indicates the name of the tracepoint to which this link is attached.

##### `tp_name_len`

[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/f2e10bff16a0fdd41ba278c84da9813700e356af)

This field indicates the length of `tp_name`

#### `tracing`

These fields apply for the `BPF_LINK_TYPE_TRACING` link type.

##### `attach_type`

[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/f2e10bff16a0fdd41ba278c84da9813700e356af)

This field indicates the [attach type](BPF_LINK_CREATE.md#attach-types) of the link.

##### `target_obj_id`

[:octicons-tag-24: v5.13](https://github.com/torvalds/linux/commit/441e8c66b23e027c00ccebd70df9fd933918eefe)

This field indicates the id of the program we are attached to for `BPF_PROG_TYPE_EXT` program types or the BTF object ID of other tracing programs.

##### `target_btf_id`

[:octicons-tag-24: v5.13](https://github.com/torvalds/linux/commit/441e8c66b23e027c00ccebd70df9fd933918eefe)

This field indicates the BTF type id of the type to which we are attached. The type id is within the BTF objected indicated by `target_obj_id`.

#### `cgroup`

These fields apply for the `BPF_LINK_TYPE_CGROUP` link type.

##### `cgroup_id`

[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/f2e10bff16a0fdd41ba278c84da9813700e356af)

This field indicates the ID of the cGroup we are attached to.

##### `attach_type`

[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/f2e10bff16a0fdd41ba278c84da9813700e356af)

This field indicates the [attach type](BPF_LINK_CREATE.md#attach-types) of the link.

#### `iter`

These fields apply for the `BPF_LINK_TYPE_ITER` link type.

[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/6b0a249a301e2af9adda84adbced3a2988248b95)

##### `target_name`

[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/6b0a249a301e2af9adda84adbced3a2988248b95)

<!-- TODO what is this field? -->

##### `target_name_len`

[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/6b0a249a301e2af9adda84adbced3a2988248b95)

<!-- TODO what is this field? -->

##### `map_id`

[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/b0c9eb37817943840a1a82dbc998c491609a0afd)

<!-- TODO what is this field? -->

##### `cgroup_id`

[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/6b0a249a301e2af9adda84adbced3a2988248b95)

This field indicates the cGroup ID.

##### `order`

[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/6b0a249a301e2af9adda84adbced3a2988248b95)

<!-- TODO what is this field? -->

##### `tid`

[:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/21fb6f2aa3890b0d0abf88b7756d0098e9367a7c)

This field indicates the thread ID.

##### `pid`

[:octicons-tag-24: v6.1](https://github.com/torvalds/linux/commit/21fb6f2aa3890b0d0abf88b7756d0098e9367a7c)

This field indicates the process ID.

#### `netns`

[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/7f045a49fee04b5662cbdeaf0838f9322ae8c63a)

These fields apply for the `BPF_LINK_TYPE_RAW_NETNS` link type.

##### `netns_ino`

[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/7f045a49fee04b5662cbdeaf0838f9322ae8c63a)

##### `attach_type`

[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/7f045a49fee04b5662cbdeaf0838f9322ae8c63a)

This field indicates the [attach type](BPF_LINK_CREATE.md#attach-types) of the link.

#### `xdp`

These fields apply for the `BPF_LINK_TYPE_XDP` link type.

[:octicons-tag-24: v5.9](https://github.com/torvalds/linux/commit/c1931c9784ebb5787c0784c112fb8baa5e8455b3)
##### `ifindex`

[:octicons-tag-24: v5.9](https://github.com/torvalds/linux/commit/c1931c9784ebb5787c0784c112fb8baa5e8455b3)

This field indicates the network interface index to which the XDP program is attached.
