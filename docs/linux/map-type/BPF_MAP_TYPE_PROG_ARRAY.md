# Map type `BPF_MAP_TYPE_PROG_ARRAY`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_PROG_ARRAY) -->
[:octicons-tag-24: v4.2](https://github.com/torvalds/linux/commit/04fd61ab36ec065e194ab5e74ae34a5240d992bb)
<!-- [/FEATURE_TAG] -->

The program array map type is a specialized map type which holds pointers to other eBPF programs and is used to facilitate tail-calls.

## Usage

Program array maps are used to perform tail-calls. Tail-calls allows one program to call into another, handing over flow control. In contrast to BPF-to-BPF function, tail calls do not return to the call site but instead run as if they were invoked by the kernel directly.

To perform a tail-call, a program should define a program array map. The loader should create the map and link it to program during the loading process. Afterwards the loader can load additional programs without attaching them to any kernel hooks. The loader can write the file descriptors of the programs as values to the program array map. The initial program can from then on use the [`bpf_tail_call`](../helper-function/bpf_tail_call.md) helper call to perform the tail-call. Check out the page of the helper for details of tail-calls.

## Attributes

Both the [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) and [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) must be exactly `4` bytes. The key is a 32-bit unsigned integer since this is an array map type.

<!-- TODO link to generic page for attributes which are the same for every map type -->

## Syscall commands

The following syscall commands work with this map type:

* [`BPF_MAP_LOOKUP_ELEM`](../syscall/BPF_MAP_LOOKUP_ELEM.md)
* [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md)
* [`BPF_MAP_GET_NEXT_KEY`](../syscall/BPF_MAP_GET_NEXT_KEY.md)

!!! note
    When writing to the map, the file descriptor of a program is expected. However, `BPF_MAP_LOOKUP_ELEM` will return the ID of a program, which you can turn into a file descriptor using the `BPF_PROG_GET_FD_BY_ID` syscall command.

## Helper functions

<!-- DO NOT EDIT MANUALLY -->
<!-- [MAP_HELPER_FUNC_REF] -->
 * [bpf_tail_call](../helper-function/bpf_tail_call.md)
<!-- [/MAP_HELPER_FUNC_REF] -->

## Flags

The following flags are supported by this map type.

### `BPF_F_NUMA_NODE`

[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/96eabe7a40aa17e613cf3db2c742ee8b1fc764d0)

When set, the [`numa_node`](../syscall/BPF_MAP_CREATE.md#numa_node) attribute is respected during map creation.

### `BPF_F_RDONLY`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/6e71b04a82248ccf13a94b85cbc674a9fefe53f5)

Setting this flag will make it so the map can only be read via the [syscall](../syscall/index.md) interface, but not written to.

For details please check the [generic description](../syscall/BPF_MAP_CREATE.md#bpf_f_rdonly).

### `BPF_F_WRONLY`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/6e71b04a82248ccf13a94b85cbc674a9fefe53f5)

Setting this flag will make it so the map can only be written to via the [syscall](../syscall/index.md) interface, but not read from.
