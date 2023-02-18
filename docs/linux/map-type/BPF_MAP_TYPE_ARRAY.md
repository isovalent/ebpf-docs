# Map type `BPF_MAP_TYPE_ARRAY`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_ARRAY) -->
[:octicons-tag-24: v3.19](https://github.com/torvalds/linux/commit/28fbcfa08d8ed7c5a50d41a0433aad222835e8e3)
<!-- [/FEATURE_TAG] -->

The array map type is a generic map type with no restrictions on the structure of the value. Like a normal array, the array map has a numeric key starting at 0 and incrementing.

## Attributes

While the [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) is essentially unrestricted, the [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) must always be `4` indicating the key is a 32-bit unsigned integer.

The value of [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) shouldn't exceed `KMALLOC_MAX_SIZE`. `KMALLOC_MAX_SIZE` is the maximum size which can be allocated by the kernel memory allocator, its exact value being dependant on a number of factors. If this edge case is hit a `-E2BIG` [error number](https://man7.org/linux/man-pages/man3/errno.3.html) is returned to the [map create syscall](../syscall/BPF_MAP_CREATE.md).

<!-- TODO link to generic page for attributes which are the same for every map type -->

## Syscall commands

The following syscall commands work with this map type:

* [`BPF_MAP_LOOKUP_ELEM`](../syscall/BPF_MAP_LOOKUP_ELEM.md)
* [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md)
* [`BPF_MAP_GET_NEXT_KEY`](../syscall/BPF_MAP_GET_NEXT_KEY.md)
* [`BPF_MAP_LOOKUP_BATCH`](../syscall/BPF_MAP_LOOKUP_BATCH.md)

## Helper functions

The following helper functions work with this map type:

<!-- DO NOT EDIT MANUALLY -->
<!-- [MAP_HELPER_FUNC_REF] -->
 * [bpf_map_lookup_elem](../helper-function/bpf_map_lookup_elem.md)
 * [bpf_map_update_elem](../helper-function/bpf_map_update_elem.md)
 * [bpf_map_delete_elem](../helper-function/bpf_map_delete_elem.md)
 * [bpf_map_for_each_callback](../helper-function/bpf_map_for_each_callback.md)
<!-- [/MAP_HELPER_FUNC_REF] -->

## Flags

The following flags are supported by this map type.

### `BPF_F_NUMA_NODE`

[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/96eabe7a40aa17e613cf3db2c742ee8b1fc764d0)

When set, the [`numa_node`](../syscall/BPF_MAP_CREATE.md#numa_node) attribute is respected during map creation.

### `BPF_F_MMAPABLE`

[:octicons-tag-24: v5.5](https://github.com/torvalds/linux/commit/fc9702273e2edb90400a34b3be76f7b08fa3344b)

Setting this flag on a `BPF_MAP_TYPE_ARRAY` will allow userspace programs to [mmap](https://man7.org/linux/man-pages/man2/mmap.2.html) the array values into the userspace process, effectively making a shared memory region between eBPF programs and a userspace program.

This can significantly improve read and write performance since there is no syscall overhead to access the map.

<!-- Based on the `bpf_map_mmap` function -->
There are a few limitation however:

* Maps containing spin locks can't be memory mapped. A map can be created with this flag but any `mmap` call will result in a `-ENOTSUPP` error code.
* Maps containing timers can't be memory mapped. A map can be created with this flag but any `mmap` call will result in a `-ENOTSUPP` error code.
* Maps which have been frozen can't be memory mapped. The `mmap` call will result in a `-EPERM` error code.
* Maps with the `BPF_F_RDONLY_PROG` flag set cannot be memory mapped, the `mmap` call would result in a `-EACCES` error code.

<!-- TODO link "spin lock" and "timer" as soon as we have pages for them -->

### `BPF_F_RDONLY`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/6e71b04a82248ccf13a94b85cbc674a9fefe53f5)

Setting this flag will make it so the map can only be read via the [syscall](../syscall/index.md) interface, but not written to.

For details please check the [generic description](../syscall/BPF_MAP_CREATE.md#bpf_f_rdonly).

### `BPF_F_WRONLY`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/6e71b04a82248ccf13a94b85cbc674a9fefe53f5)

Setting this flag will make it so the map can only be written to via the [syscall](../syscall/index.md) interface, but not read from.

### `BPF_F_RDONLY_PROG`

[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/591fe9888d7809d9ee5c828020b6c6ae27c37229)

Setting this flag will make it so the map can only be read via [helper functions](../helper-function/index.md), but not written to.

For details please check the [generic description](../syscall/BPF_MAP_CREATE.md#bpf_f_rdonly_prog).

### `BPF_F_WRONLY_PROG`

[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/591fe9888d7809d9ee5c828020b6c6ae27c37229)

Setting this flag will make it so the map can only be written to via [helper functions](../helper-function/index.md), but not read from.

For details please check the [generic description](../syscall/BPF_MAP_CREATE.md#bpf_f_wronly_prog).

## Global data
[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/d8eca5bbb2be9bc7546f9e733786fa2f1a594c67)

In regular C programs running in userspace it is not uncommon to use global variables. These live in the heap memory accessible by the processes, the memory location of which are relocated by the operating system when the program is started.

In eBPF we have no heap, only the stack and pointers into kernel space. To work around the issue, we use array maps. The compiler will place global data into the `.data`, `.rodata`, and `.bss` [ELF](../../elf.md) sections. The [loader](../../loader.md) will take these sections and turn them into array maps with a single key (`max_entries` set to 1) and its value the same size as the binary blob in the ELF section.

The compiler emits special `LDIMM64` instructions with the first source register set to `BPF_PSEUDO_MAP_VALUE` and an offset into the map value. The loader will relocate the correct map file descriptor into it as well. For details, check out the [instruction set](../../instruction-set.md) page.

All of this results in the ability to get a pointer into the global data in a single instruction much like in regular userspace programs.

<!-- ## Usage -->
<!-- ### Global data -->
<!-- TODO make an example -->

<!-- ### Memory mapping -->
<!-- TODO make an example -->
