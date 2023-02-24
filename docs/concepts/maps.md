# Maps
Maps provide a way for eBPF programs to communicate with each other (kernel space) and with user space.

When both kernel and user space access the same maps they will need a common understanding of the key and value structures in memory. This can work if both programs are written in C and they share a header. Otherwise, both the user space language and the kernel space structures must understand the k/v structures byte-for-byte.

There are two ways to define maps.

## Internal representation
The internal representation of the map definitions.

`tools/lib/bpf/libbpf_internal.h:229`
```c
struct btf_map_def {
	enum map_def_parts parts;
	__u32 map_type;
	__u32 key_type_id;
	__u32 key_size;
	__u32 value_type_id;
	__u32 value_size;
	__u32 max_entries;
	__u32 map_flags;
	__u32 numa_node;
	__u32 pinning;
	__u64 map_extra;
};
```

## Legacy Maps
The below map definition and system call were the way to define a map before [[#BTF Style Maps]]

```c
union bpf_attr my_map {
	.map_type = BPF_MAP_TYPE_HASH,
	.key_size = sizeof(int),
	.value_size = sizeof(int),
	.max_entries = 100,
	.map_flags = BPF_F_NO_PREALLOC,
};
int fd = bpf(BPF_MAP_CREATE, &my_map, sizeof(my_map));
```

## BTF Style Maps
The new way of defining eBPF maps which utilize BTF type information.
See https://lwn.net/ml/netdev/20190531202132.379386-7-andriin@fb.com/ implementation and details. 

```c
struct my_value { int x, y, z; };

struct {
	int type;
	int max_entries;
	int *key;
	struct my_value *value;
} btf_map SEC(".maps") = {
	.type = BPF_MAP_TYPE_ARRAY,
	.max_entries = 16,
};
```

Example: Maps are global/static variables thrown into the `.maps` ELF section.

Typically, macros in `tools/lib/bpf/bpf_helpers.h` are used to shrink the map's definition syntax.

```
#define __uint(name, val) int (*name)[val]
#define __type(name, val) typeof(val) *name
#define __array(name, val) typeof(val) *name[]
```

Therefore above map would traditionally look as follows:
```c
struct my_value { int x, y, z; };

struct {
	__uint(type, BPF_MAP_TYPE_ARRAY);
	__type(key, int);
	__type(value, struct my_value);
	__uint(max_entries, 16);
} icmpcnt SEC(".maps");

```

## Creating BPF Maps

### bpf syscall

A syscall exists which performs many `bpf` related actions. 

It can be used to create a map directly in the kernel.

```c
union bpf_attr my_map {
.map_type = BPF_MAP_TYPE_HASH,
.key_size = sizeof(int),
.value_size = sizeof(int),
.max_entries = 100,
.map_flags = BPF_F_NO_PREALLOC,
};
int fd = bpf(BPF_MAP_CREATE, &my_map, sizeof(my_map));
```

```ad-note

See: https://elixir.bootlin.com/linux/latest/source/include/uapi/linux/bpf.h#L1272
```

### libbpf  `bpf_map_create`
`/tools/lib/bpf/bpf.h`
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

### Global static definition

If a map is globally and statically defined in your eBPF program the maps will be created on load of the eBPF program. 

These map definitions look just like the sections [[#Legacy Maps]] and [[#BTF Style Maps]]. 

When defined like this a global variable will contain the map's details and `fd`:
`map_data[0]`

`map_data` is an array which map data is placed in, in the order they are defined. 

## Global static definitions and pinning

When defining maps statically you can also specify a "pinning" field to the global variable. 

By default maps are pinned to `/sys/fs/bpf`.

This will instruct the kernel to also pin this map to communicate with user space on loading. 

```c
  struct {                                                                                                                                                                         
  |   __uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);                                                                                                                                 
  |   __uint(key_size, sizeof(u32));                                                                                                                                               
  |   __uint(value_size, sizeof(u32));                                                                                                                                            
  |   __uint(pinning, LIBBPF_PIN_BY_NAME);                                                                                                                                         
  } events SEC(".maps"); 
```


> [!NOTE] 
> With BTF style maps its possible to use either `key_size` or `key` as long as the structure defined in `key` has size information in the resulting eBPF object file's BTF.

The valid fields are parsed out in:
```c
tools/lib/bpf/libbpf.c
int parse_btf_map_def(const char *map_name, struct btf *btf,
		      const struct btf_type *def_t, bool strict,
		      struct btf_map_def *map_def, struct btf_map_def *inner_def)
{
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



## Kernel Space vs User Space API
Maps are manipulated differently from kernel space then user space.

`tools/lib/bpf/bpf_helpers.h` - prototypes for manipulating maps within eBPF programs running in the kernel. Allows access to the map object and kernel memory. Map updates are atomic. 

`tools/lib/bpf/bpf.h` - prototypes for manipulating maps from user space, works with maps by `fd`. Map updates are not atomic and involve a copy between user and kernel space.

## Manipulating maps

### Updating Elements
`bpf_helpers.h`
```c
   *                                                                                 
   * bpf_map_update_elem                                                             
   *                                                                                 
   *  Add or update the value of the entry associated to *key* in                    
   *  *map* with *value*. *flags* is one of:                                         
   *                                                                                 
   *  **BPF_NOEXIST**                                                                
   *      The entry for *key* must not exist in the map.                             
   *  **BPF_EXIST**                                                                  
   *      The entry for *key* must already exist in the map.                         
   *  **BPF_ANY**                                                                    
   *      No condition on the existence of the entry for *key*.                      
   *                                                                                 
   *  Flag value **BPF_NOEXIST** cannot be used for maps of types                    
   *  **BPF_MAP_TYPE_ARRAY** or **BPF_MAP_TYPE_PERCPU_ARRAY**  (all                  
   *  elements always exist), the helper would return an error.                      
   *                                                                                 
   * Returns                                                                         
   *  0 on success, or a negative error in case of failure.                          
   */                                                                                
static long (*bpf_map_update_elem)(void *map, const void *key, const void *value, __u64 flags) = (void *) 2;   
```

example:
```c
int key, value, result;

key = 1, value = 5678;

result = bpf_map_update_elem(&my_map, &key, &value, BPF_NOEXIST);
if (result == 0)
	printf("Map updated with new element\n");
else
	printf("Failed to update map with new value: %d (%s)\n",
		result, strerror(errno));
```

### Reading Elements

`bpf_helpers.h`
```c
  /*                                                                              
   * bpf_map_lookup_elem                                                          
   *                                                                              
   *  Perform a lookup in *map* for an entry associated to *key*.                 
   *                                                                              
   * Returns                                                                      
   *  Map value associated to *key*, or **NULL** if no entry was                  
   *  found.                                                                      
   */                                                                             
  static void *(*bpf_map_lookup_elem)(void *map, const void *key) = (void *) 1; 
```

example:
```c
int key, value, result; // value is going to store the expected element's value
key = 1;
result = bpf_map_lookup_elem(&my_map, &key, &value);
if (result == 0)
	printf("Value read from the map: '%d'\n", value);
else
	printf("Failed to read value from the map: %d (%s)\n",
		result, strerror(errno));
```

### Removing Element
`bpf_helpers.h`
```c
  /*                                                                              
   * bpf_map_delete_elem                                                          
   *                                                                              
   *  Delete entry with *key* from *map*.                                         
   *                                                                              
   * Returns                                                                      
   *  0 on success, or a negative error in case of failure.                       
   */                                                                             
  static long (*bpf_map_delete_elem)(void *map, const void *key) = (void *) 3; 
```

example:
```c
int key, result;
key = 1;
result = bpf_map_delete_element(&my_map, &key);
if (result == 0)
	printf("Element deleted from the map\n");
else
	printf("Failed to delete element from the map: %d (%s)\n",
		result, strerror(errno));
```

`bpf.c`
```c
int bpf_map_lookup_and_delete_elem(int fd, const void *key, void *value)        
{                                                                               
|   union bpf_attr attr;                                                        
|   int ret;                                                                    
|                                                                               
|   memset(&attr, 0, sizeof(attr));                                             
|   attr.map_fd = fd;                                                           
|   attr.key = ptr_to_u64(key);                                                 
|   attr.value = ptr_to_u64(value);                                             
|                                                                               
|   ret = sys_bpf(BPF_MAP_LOOKUP_AND_DELETE_ELEM, &attr, sizeof(attr));                                                                                                                      
|   return libbpf_err_errno(ret);                                                                                                                                                            
}    
```

### Iterating Over Elements
`bpf_helpers.h`
```c
/*                                                                              
* bpf_for_each_map_elem                                                        
*                                                                              
*  For each element in **map**, call **callback_fn** function with             
*  **map**, **callback_ctx** and other map-specific parameters.                
*  The **callback_fn** should be a static function and                         
*  the **callback_ctx** should be a pointer to the stack.                      
*  The **flags** is used to control certain aspects of the helper.             
*  Currently, the **flags** must be 0.                                         
*                                                                              
*  The following are a list of supported map types and their                   
*  respective expected callback signatures:                                    
*                                                                              
*  BPF_MAP_TYPE_HASH, BPF_MAP_TYPE_PERCPU_HASH,                                
*  BPF_MAP_TYPE_LRU_HASH, BPF_MAP_TYPE_LRU_PERCPU_HASH,                        
*  BPF_MAP_TYPE_ARRAY, BPF_MAP_TYPE_PERCPU_ARRAY                               
*                                                                              
*  long (\*callback_fn)(struct bpf_map \*map, const void \*key, void \*value, void \*ctx);
*                                                                              
*  For per_cpu maps, the map_value is the value on the cpu where the           
*  bpf_prog is running.                                                        
*                                                                              
*  If **callback_fn** return 0, the helper will continue to the next           
*  element. If return value is 1, the helper will skip the rest of             
*  elements and return. Other return values are not used now.                  
*                                                                              
*                                                                              
* Returns                                                                      
*  The number of traversed map elements for success, **-EINVAL** for           
*  invalid **flags**.                                                          
*/                                                                             
static long (*bpf_for_each_map_elem)(void *map, void *callback_fn, void *callback_ctx, __u64 flags) = (void *) 164;
```

`bpf.c`
```c
int bpf_map_get_next_key(int fd, const void *key, void *next_key)               
{                                                                               
|   union bpf_attr attr;                                                        
|   int ret;                                                                    
|                                                                               
|   memset(&attr, 0, sizeof(attr));                                             
|   attr.map_fd = fd;                                                           
|   attr.key = ptr_to_u64(key);                                                 
|   attr.next_key = ptr_to_u64(next_key);                                       
|                                                                               
|   ret = sys_bpf(BPF_MAP_GET_NEXT_KEY, &attr, sizeof(attr));                   
|   return libbpf_err_errno(ret);                                               
}    
```

## Concurrency
eBPF spin locks lock a map element while an eBPF program accesses it.

In user space a flag `BPF_F_LOCK` can be used with `bpf_map_update_elem` and `bpf_map_lookup_elem_flags` helper functions. 

In eBPF kernel space a semaphore can be added to map elements. 

```c
struct concurrent_element {
	struct bpf_spin_lock semaphore;
	int count;
}

...

struct bpf_map_def SEC("maps") concurrent_map = {
	.type = BPF_MAP_TYPE_HASH,
	.key_size = sizeof(int),
	.value_size = sizeof(struct concurrent_element),
	.max_entries = 100,
};
// this is legacy, this whole map definition should be BTF style and this macro
// won't be necessary.
BPF_ANNOTATE_KV_PAIR(concurrent_map, int, struct concurrent_element);

...

int bpf_program(struct pt_regs *ctx) {
	int key = 0;
	struct concurrent_element init_value = {};
	struct concurrent_element *read_value;

	bpf_map_create_elem(&concurrent_map, &key, &init_value, BPF_NOEXIST);

	read_value = bpf_map_lookup_elem(&concurrent_map, &key);
	bpf_spin_lock(&read_value->semaphore);
	read_value->count += 100;
	bpf_spin_unlock(&read_value->semaphore);
}
```

## Map Types
`include/uapi/linux/bpf.h:880`

`BPF_MAP_TYPE_HASH` - general hash map, allocations and frees handled by kernel. `bpf_map_update_elem` is atomic. 

`BPF_MAP_TYPE_ARRAY`  - elements pre-allocated to their zero values. cannot delete values and remains at fixed size. `bpf_map_update_elem` is not atomic. keys are 4 bytes.

`BPF_MAP_TYPE_PROG_ARRAY` - stores file descriptors to other eBPF programs. used with `bpf_tail_call` to jump between programs. keys and values are 4 byte.

`BPF_MAP_TYPE_PERF_EVENT_ARRAY` - used to forward event traces from the kernel to user space. used with `bpf_perf_event_output` which can write events to a map of this type without worrying about keys. 

`BPF_MAP_TYPE_PERCPU_HASH` - similar to `BPF_MAP_TYPE_HASH` but each CPU maintains its own version of the map. 

`BPF_MAP_TYPE_PERCPU_ARRAY` - similar to `BPF_MAP_TYPE_ARRAY` but each CPU maintains its own version of the map.

`BPF_MAP_TYPE_STACK_TRACE` - holds stack traces from kernel or user space programs

`BPF_MAP_TYPE_CGROUP_ARRAY` - stores references to cgroups, useful for sharing cgroup references between eBPF programs. 

`BPF_MAP_TYPE_LRU_HASH and BPF_MAP_TYPE_LRU_PERCPU_HASH` - similar to `BPF_MAP_TYPE_HASH` but with `LRU` semantics. 

`BPF_MAP_TYPE_LPM_TRIE` - longest prefix matches maps. map keys must be multiple of 8s and between 8-2048. can be used with `bpf_lpm_trie_key` to create a key. 

`BPF_MAP_TYPE_ARRAY_OF_MAPS and BPF_MAP_TYPE_HASH_OF_MAPS` - maps which store pointers to other maps. only one level of nesting. the kernel ensures updates to elements in these maps don't take place until all references to the element are dropped. 

`BPF_MAP_TYPE_DEVMAP` - holds references to network devices. used with `bpf_redirect_map` to redirect packets to elements in this map. 

`BPF_MAP_TYPE_CPUMAP` - allows redirection of network traffic to specific CPUs. useful for scalability and performance isolation of CPUs.

`BPF_MAP_TYPE_XSKMAP` - stores open sockets for forwarding packets. 

`BPF_MAP_TYPE_SOCKMAP and BPF_MAP_TYPE_SOCKHASH` - can be used to forward packets from the current XDP program to a different socket. 

`BPF_MAP_TYPE_CGROUP_STORAGE and BPF_MAP_TYPE_PERCPU_CGROUP_STORAGE` - can be used with `bpf_cgroup_storage_key`. These maps provide restricted access to values by cgroups. 

`BPF_MAP_TYPE_REUSEPORT_SOCKARRAY` - typically used with `BPF_PROG_TYPE_SK_REUSEPORT`. Can help define which sockets gets which packets even tho both sockets share a port. 

`BPF_MAP_TYPE_QUEUE` - FIFO style maps. map keys are not used for this map and the size should be 0. When using bpf helper methods keys should be null.

`BPF_MAP_TYPE_STACK` - like  `BPF_MAP_TYPE_QUEUE` but LIFO semantics. 

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

