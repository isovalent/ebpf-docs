# Maps

Maps provide a way for eBPF programs to communicate with each other (kernel space) and with user space.

When both kernel and user space access the same maps they will need a common understanding of the key and value structures in memory. This can work if both programs are written in C and they share a header. Otherwise, both the user space language and the kernel space structures must understand the k/v structures byte-for-byte.

Maps come in a variety of [types](../map-type/index.md), each of which working in a slightly different way, like different data structures.

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

```
#define __uint(name, val) int (*name)[val]
#define __type(name, val) typeof(val) *name
#define __array(name, val) typeof(val) *name[]
```

Valid fields:

* `type` - enum, one of `enum bpf_map_type` located in `include/uapi/linux/bpf.h`.
* `max_entries` - 
* `map_flags` -
* `numa_node` -
* `key_size` -
* `key` -
* `value_size` - 
* `value` -
* `values` -
* `pinning` -

!!! note
    With BTF style maps its possible to use either `key_size` or `key` as long as the structure defined in `key` has size information in the resulting eBPF object file's BTF.

## Creating BPF Maps

It is common for maps to be declared in the eBPF program, but maps are ultimately created from userspace. Most loader libraries pick up the map declarations from the compiled ELF file and create them automatically for the user.

However it is also possible for users to manually create maps using the [BPF_MAP_CREATE](../syscall/BPF_MAP_CREATE.md) command of the BPF syscall or to use a loader library which such capabilities.

### LibBPF

LibBPF is such a library, it provide the `bpf_map_create` function to allow for the manual creation of maps.

[`/tools/lib/bpf/bpf.h`](https://elixir.bootlin.com/linux/v6.2.2/source/tools/lib/bpf/bpf.h#L40)
```c
LIBBPF_API int bpf_map_create(enum bpf_map_type map_type,                           
│   │   │   │     const char *map_name,                                                                                                                                                                                                                        
│   │   │   │     __u32 key_size,                                                                                                                                                                                                                              
│   │   │   │     __u32 value_size,                                                                                                                                                                                                                            
│   │   │   │     __u32 max_entries,                                                                                                                                                                                                                           
│   │   │   │     const struct bpf_map_create_opts *opts);                                                                                                                                                                                                     
                                                                                    
struct bpf_create_map_attr {                                                        
│   const char *name;                                                                                                                                                                                                                                          
│   enum bpf_map_type map_type;                                                                                                                                                                                                                                
│   __u32 map_flags;                                                                                                                                                                                                                                           
│   __u32 key_size;                                                                                                                                                                                                                                            
│   __u32 value_size;                                                                                                                                                                                                                                          
│   __u32 max_entries;                                                                                                                                                                                                                                         
│   __u32 numa_node;                                                                                                                                                                                                                                           
│   __u32 btf_fd;                                                                                                                                                                                                                                              
│   __u32 btf_key_type_id;                                                                                                                                                                                                                                     
│   __u32 btf_value_type_id;                                                                                                                                                                                                                                   
│   __u32 map_ifindex;                                                                                                                                                                                                                                         
│   union {                                                                                                                                                                                                                                                    
│   │   __u32 inner_map_fd;                                                                                                                                                                                                                                    
│   │   __u32 btf_vmlinux_value_type_id;                                                                                                                                                                                                                       
│   };                                                                                                                                                                                                                                                         
};  
```

The `bpf_map_create` function in `libbpf` can be used to create maps during runtime. 


## Using maps

Maps are manipulated differently from kernel space then user space.

### Using from kernel space

eBPF programs can interact with maps via [helper functions](../helper-function/index.md), these are defined in `tools/lib/bpf/bpf_helpers.h`. The exact helper functions that can be used to interact with maps depends on the [map type](../map-type/index.md), you can reference the list of supported helper calls on the page of a given map type.

Elements of generic maps can we read using the [`bpf_map_lookup_elem`](../helper-function/bpf_map_lookup_elem.md) helper, updated with the [`bpf_map_update_elem`](../helper-function/bpf_map_update_elem.md), and deleted with [`bpf_map_delete_elem`](../helper-function/bpf_map_delete_elem.md).

Some of these generic map types can be iterated using the [`bpf_for_each_map_elem`](../helper-function/bpf_for_each_map_elem.md) helper function.

More specialized map types might require special helpers such as [`bpf_redirect_map`](../helper-function/bpf_redirect_map.md) to perform packet redirection based on the contents of a map. Or [`bpf_perf_event_output`](../helper-function/bpf_perf_event_output.md) to send a message via a [`BPF_MAP_TYPE_PERF_EVENT_ARRAY`](../map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md) map.

### Using from user space

From user space we can also use maps in a number of ways. Most [map types](../map-type/index.md) support reading via the [`BPF_MAP_LOOKUP_ELEM`](../syscall/BPF_MAP_LOOKUP_ELEM.md) syscall command, writing or updating with the [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md) syscall command, and deleting with the [`BPF_MAP_DELETE_ELEM`](../syscall/BPF_MAP_DELETE_ELEM.md) syscall command. However, this does not hold true for all map types, check the page of the specific map type to see what syscall commands it supports.

Besides the single key versions, there are also batch variants of those syscall commands: [`BPF_MAP_LOOKUP_BATCH`](../syscall/BPF_MAP_LOOKUP_BATCH.md), [`BPF_MAP_UPDATE_BATCH`](../syscall/BPF_MAP_UPDATE_BATCH.md), and [`BPF_MAP_DELETE_BATCH`](../syscall/BPF_MAP_DELETE_BATCH.md). These work for a smaller sub-set of maps, again, check the specific map type for compatibility.

Most map types support iterating over keys using the [`BPF_MAP_GET_NEXT_KEY`](../syscall/BPF_MAP_GET_NEXT_KEY.md) syscall command.

Some map types like the [`BPF_MAP_TYPE_PERF_EVENT_ARRAY`](../map-type/BPF_MAP_TYPE_PERF_EVENT_ARRAY.md) require the usage of additional mechanisms like perf_event and ring buffers to read the actual data sent via the [`bpf_perf_event_output`](../helper-function/bpf_perf_event_output.md) helper from the kernel side.

## BPF Virtual Filesystem
When a program which creates an eBPF map exits its maps are removed from the kernel as well. 

To prevent this you can "pin" a map to the eBPF virtual filesystem. 

Two new syscalls exist to support this. 

Default `bpf` file-system is located at `sys/fs/bpf`. 

```
# mount -t bpf /sys/fs/bpf /sys/fs/bpf
```

This directory can be organized as the user sees fit. Example: `sys/fs/bpf/shared/ips` can hold IP information maps.

These maps can hold eBPF maps and other eBPF programs. They are referenced as file descriptors. 

`BPF_PIN_FD` - command to pin an  eBPF map via `bpf` syscall.
`BPF_OBJ_GET` - command to get a `fd` of a pin map.

Helpers exist to abstract the syscalls:
`tools/lib/bpf/bpf.h`
```c
LIBBPF_API int bpf_obj_pin(int fd, const char *pathname);                           
LIBBPF_API int bpf_obj_get(const char *pathname);    
```

`libbpf` is used to load a eBPF object file with this signature:

```c
struct bpf_object *
bpf_object__open_file(const char *path, const struct bpf_object_open_opts *opts)
{
```

The option struct provides this option:

```c
(abbreviated)
struct bpf_object_open_opts {
	...
	/* maps that set the 'pinning' attribute in their definition will have
	 * their pin_path attribute set to a file in this directory, and be
	 * auto-pinned to that path on load; defaults to "/sys/fs/bpf".
	 */
	const char *pin_root_path;
```

Alternatively you can overwrite this by defining a map value specifically. 

## Example pinning program
`Makefile`
```
all: loader printer 

loader: loader.o
	gcc -lbpf -g -o $@ loader.o

loader.o: loader.c
	gcc -g -c -o $@ loader.c
	
printer: printer.o
	gcc -lbpf -g -o $@ printer.o

printer.o: printer.c
	gcc -g -c -o $@ printer.c

```

`loader.c`
```c
#include <stdio.h>
#include <errno.h>
#include <bpf/bpf.h>

static const char * file_path = "/sys/fs/bpf/my_array";

int main(int argc, char *argv[]) {
    int key, value, fd, added, pinned;

    fd = bpf_map_create(
            BPF_MAP_TYPE_ARRAY, "my_array", sizeof(int), sizeof(int), 100, 0
    );

    if (fd < 0) {
        printf("Failed to create map: %d (%s)\n", fd, strerror(errno));
        return -1;
    }

    key = 1, value = 1;  
    added = bpf_map_update_elem(fd, &key, &value, BPF_ANY);
    if (added < 0) {
        printf("Failed to update map: %d (%s) \n", added, strerror(errno));
        return -1;
    }

    pinned = bpf_obj_pin(fd, file_path);
    if (pinned < 0) {
        printf("Failed to pin map: %d (%s) \n", pinned, strerror(errno));
        return -1;
    }

    return 0;
};
```

`printer.c`
```c
#include <stdio.h>
#include <errno.h>
#include <bpf/bpf.h>

static const char *file_path = "/sys/fs/bpf/my_array";

int main (int argc, char *argv[]) {
    int fd, key, value, result;


    fd = bpf_obj_get(file_path);
    if (fd < 0) {
        printf("Failed to fetch map: %d (%s)", fd, strerror(fd));
        return -1;
    }

    key = 1;
    result = bpf_map_lookup_elem(fd, &key, &value);
    if (result < 0){
        printf("Failed to fetch map: %d (%s)", result, strerror(fd));
        return -1;
    }

    printf("Value read from the map: '%d'\n", value);
    return 0;
}
```

