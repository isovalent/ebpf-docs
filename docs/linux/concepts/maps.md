---
title: "Maps"
description: "This page explains the concept of eBPF maps. It goes into depth about how to define, create, and use maps in eBPF programs."
---
# Maps

Maps provide a way for eBPF programs to communicate with each other (kernel space) and with user space.

When both kernel and user space access the same maps they will need a common understanding of the key and value structures in memory. This can work if both programs are written in C and they share a header. Otherwise, both the user space language and the kernel space structures must understand the k/v structures byte-for-byte.

Maps come in a variety of [types](../map-type/index.md), each of which works in a slightly different way, like different data structures.

## Defining maps in eBPF programs

Before we can start using maps in our eBPF programs we have to define them.

### Legacy Maps

The legacy way of defining maps is to use the `struct bpf_map_def` type from libbpf's eBPF side library or from the linux uapi. These map declarations should reside in the `maps` ELF section. The major downside of this method is that key and value type information is lost, which is why it was replaced by [BTF style maps](#btf-style-maps).

```c
struct bpf_map_def my_map = {
	.type = BPF_MAP_TYPE_HASH,
	.key_size = sizeof(int),
	.value_size = sizeof(int),
	.max_entries = 100,
	.map_flags = BPF_F_NO_PREALLOC,
} SEC("maps");
```

### BTF Style Maps

The new way of defining eBPF maps which utilize BTF type information.
See [mailing list link](https://lwn.net/ml/netdev/20190531202132.379386-7-andriin@fb.com/) for implementation details.

These maps should be located in the `.maps` section for loaders to properly pick them up.

```c
struct my_value { int x, y, z; };

struct {
	__uint(type, BPF_MAP_TYPE_ARRAY);
	__type(key, int);
	__type(value, struct my_value);
	__uint(max_entries, 16);
} icmpcnt SEC(".maps");

```

The `__uint` and `__type` macros used in the above example are typically used to make the type definition easier to read.
They are defined in [`tools/lib/bpf/bpf_helpers.h`](https://elixir.bootlin.com/linux/v6.2.2/source/tools/lib/bpf/bpf_helpers.h).

```c
#define __uint(name, val) int (*name)[val]
#define __type(name, val) typeof(val) *name
#define __array(name, val) typeof(val) *name[]
```

The `name` part of these macros refers to field names of the to be created structure. Not all names are recognized by libbpf and compatible libraries. However, the following are:

* `type` (`__uint`) - enum, see the [map types](../map-type/index.md) index for all valid options.
* `max_entries` (`__uint`) - int indicating the maximum amount of entries.
* `map_flags` (`__uint`) - a bitfield of flags, see [flags section](../syscall/BPF_MAP_CREATE.md#flags) in map load syscall command for valid options. 
* `numa_node` (`__uint`) - the ID of the NUMA node on which to place the map.
* `key_size` (`__uint`) - the size of the key in bytes. This field is mutually exclusive with the `key` field.
* `key` (`__type`) - the type of the key. This field is mutually exclusive with the `key_size` field.
* `value_size` (`__uint`) - the size of the value in bytes. This field is mutually exclusive with the `value` and `values` fields.
* `value` (`__type`) - the type of the value. This field is mutually exclusive with the `value` and `value_size` fields.
* `values` (`__array`) - see [static values section](#static-values). This field is mutually exclusive with the `value` and `value_size` field.
* `pinning` (`__uint`) - `LIBBPF_PIN_BY_NAME` or `LIBBPF_PIN_NONE` see [pinning page](pinning.md) for details.
* `map_extra` (`__uint`) - Addition settings, currently only used by bloom filters which use the lowest 4 bits to indicate the amount of hashes used in the bloom filter.

Typically, only the `type`, `key`/`key_size`, `value`/`values`/`value_size`, and `max_entries` fields are required.

#### Static values

The `values` map field has a syntax when used, it is the only field to use the `__array` macro and requires us to initialize our map constant with a value. Its purpose is to populate the contents of the map during loading without having to do so manually via a userspace application. This is especially handy for users who use `ip`, `tc`, or `bpftool` to load their programs.

The `val` part of the `__array` parameter should contain a type describing the individual array elements. The values we would like to pre-populate should go into the value part of the struct initialization.

The following examples show how to pre-populate a map-in-map:

```c
struct inner_map {
	__uint(type, BPF_MAP_TYPE_ARRAY);
	__uint(max_entries, INNER_MAX_ENTRIES);
	__type(key, __u32);
	__type(value, __u32);
} inner_map SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_ARRAY_OF_MAPS);
	__uint(max_entries, MAX_ENTRIES);
	__type(key, __u32);
	__type(value, __u32);
	__array(values, struct {
		__uint(type, BPF_MAP_TYPE_ARRAY);
		__uint(max_entries, INNER_MAX_ENTRIES);
		__type(key, __u32);
		__type(value, __u32);
	});
} m_array_of_maps SEC(".maps") = {
	.values = { (void *)&inner_map, 0, 0, 0, 0, 0, 0, 0, 0 },
};
```

Another common use is to pre-populate a tail-call map:

```c
struct {
	__uint(type, BPF_MAP_TYPE_PROG_ARRAY);
	__uint(max_entries, 2);
	__uint(key_size, sizeof(__u32));
	__array(values, int (void *));
} prog_array_init SEC(".maps") = {
	.values = {
		[1] = (void *)&tailcall_1,
	},
};
```

## Creating BPF Maps

It is common for maps to be declared in the eBPF program, but maps are ultimately created from userspace. Most loader libraries pick up the map declarations from the compiled ELF file and create them automatically for the user.

However, it is also possible for users to manually create maps using the [BPF_MAP_CREATE](../syscall/BPF_MAP_CREATE.md) command of the BPF syscall or to use a loader library with such capabilities.

### LibBPF

LibBPF is such a library, it provides the `bpf_map_create` function to allow for the manual creation of maps.

[`/tools/lib/bpf/bpf.h`](https://elixir.bootlin.com/linux/v6.2.2/source/tools/lib/bpf/bpf.h#L40)
```c
LIBBPF_API int bpf_map_create(enum bpf_map_type map_type,
			      const char *map_name,
			      __u32 key_size,
			      __u32 value_size,
			      __u32 max_entries,
			      const struct bpf_map_create_opts *opts);                                                                                                                                                                                                   
                                                                                    
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
};
```

The `bpf_map_create` function in `libbpf` can be used to create maps during runtime. 

## Using maps

Maps are manipulated differently from kernel space than user space.

### Using from kernel space

eBPF programs can interact with maps via [helper functions](../helper-function/index.md), these are defined in `tools/lib/bpf/bpf_helpers.h`. The exact helper functions that can be used to interact with maps depend on the [map type](../map-type/index.md), you can reference the list of supported helper calls on the page of a given map type.

Elements of generic maps can be read using the [`bpf_map_lookup_elem`](../helper-function/bpf_map_lookup_elem.md) helper, updated with the [`bpf_map_update_elem`](../helper-function/bpf_map_update_elem.md), and deleted with [`bpf_map_delete_elem`](../helper-function/bpf_map_delete_elem.md).

Some of these generic map types can be iterated using the [`bpf_for_each_map_elem`](../helper-function/bpf_for_each_map_elem.md) helper function.

More specialized map types might require special helpers such as [`bpf_redirect_map`](../helper-function/bpf_redirect_map.md) to perform packet redirection based on the contents of a map. Or [`bpf_perf_event_output`](../helper-function/bpf_perf_event_output.md) to send a message via a [`BPF_MAP_TYPE_PERF_EVENT_ARRAY`](../map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md) map.

### Using from user space

From user space we can also use maps in a number of ways. Most [map types](../map-type/index.md) support reading via the [`BPF_MAP_LOOKUP_ELEM`](../syscall/BPF_MAP_LOOKUP_ELEM.md) syscall command, writing or updating with the [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md) syscall command, and deleting with the [`BPF_MAP_DELETE_ELEM`](../syscall/BPF_MAP_DELETE_ELEM.md) syscall command. However, this does not hold true for all map types, check the page of the specific map type to see what syscall commands it supports.

Besides the single key versions, there are also batch variants of those syscall commands: [`BPF_MAP_LOOKUP_BATCH`](../syscall/BPF_MAP_LOOKUP_BATCH.md), [`BPF_MAP_UPDATE_BATCH`](../syscall/BPF_MAP_UPDATE_BATCH.md), and [`BPF_MAP_DELETE_BATCH`](../syscall/BPF_MAP_DELETE_BATCH.md). These work for a smaller sub-set of maps, again, check the specific map type for compatibility.

Most map types support iterating over keys using the [`BPF_MAP_GET_NEXT_KEY`](../syscall/BPF_MAP_GET_NEXT_KEY.md) syscall command.

Some map types like the [`BPF_MAP_TYPE_PERF_EVENT_ARRAY`](../map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md) require the usage of additional mechanisms like perf_event and ring buffers to read the actual data sent via the [`bpf_perf_event_output`](../helper-function/bpf_perf_event_output.md) helper from the kernel side.
