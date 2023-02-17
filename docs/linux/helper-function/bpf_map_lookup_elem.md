# Helper call `bpf_map_lookup_elem`

The lookup map element helper call is used to read values from [maps](../index.md#maps).

!!! note
    Not all [map types](../map-type/index.md) support this helper call due to their implementation, check the map type page for details.

## Definition

<!-- [HELPER_FUNC_DEF] -->
Delete entry with *key* from *map*.

**Returns**
0 on success, or a negative error in case of failure.

`#!c static long (*bpf_map_delete_elem)(void *map, const void *key) = (void *) 3;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

The `map` argument must be a pointer to a map definition and `key` must be a pointer to the key you
wish to lookup.

The return value will be a pointer to the map value or `NULL`. The value is a direct reference to the kernel memory where this map value is stored, not a copy. Therefor any modifications made to the value are automatically persisted without the need to call any additional helpers.

!!! warning
    modifying map values of non per-CPU maps is subject to race conditions, atomics or spinlocks must be utilized to prevent race conditions if they are detrimental to your use case.
    <!-- TODO link to guide on memory access serialization -->


### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [BPF_PROG_KRPOBE](../program-type/BPF_PROG_KRPOBE.md)
 * [BPF_PROG_KRPOBE](../program-type/BPF_PROG_KRPOBE.md)
 * [BPF_PROG_TYPE_CGROUP_SKB](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [BPF_PROG_TYPE_LWT_SEG6LOCAL](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
 * [BPF_PROG_TYPE_LIRC_MODE2](../program-type/BPF_PROG_TYPE_LIRC_MODE2.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
 * [BPF_PROG_TYPE_LWT_OUT](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
 * [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
 * [BPF_PROG_TYPE_SK_LOOKUP](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md)
 * [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
 * [BPF_PROG_TYPE_PERF_EVENT](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [BPF_PROG_TYPE_PERF_EVENT](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [BPF_PROG_TYPE_FLOW_DISSECTOR](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md)
 * [BPF_PROG_TYPE_CGROUP_SOCKOPT](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md)
 * [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
 * [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
 * [BPF_PROG_TYPE_TRACEPOINT](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [BPF_PROG_TYPE_TRACEPOINT](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [BPF_PROG_TYPE_LWT_IN](../program-type/BPF_PROG_TYPE_LWT_IN.md)
 * [BPF_PROG_TYPE_SOCK_OPS](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
 * [BPF_PROG_TYPE_SK_SKB](../program-type/BPF_PROG_TYPE_SK_SKB.md)
 * [BPF_PROG_TYPE_SK_MSG](../program-type/BPF_PROG_TYPE_SK_MSG.md)
 * [BPF_PROG_TYPE_SYSCALL](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [BPF_PROG_TYPE_SYSCALL](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [BPF_PROG_TYPE_SOCKET_FILTER](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
 * [BPF_PROG_TYPE_CGROUP_DEVICE](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK_ADDR](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [BPF_PROG_TYPE_SK_REUSEPORT](../program-type/BPF_PROG_TYPE_SK_REUSEPORT.md)
 * [BPF_PROG_TYPE_CGROUP_SYSCTL](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Map types

This helper call can be used with the following map types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_MAP_REF] -->
 * [BPF_MAP_TYPE_SOCKMAP](../map-type/BPF_MAP_TYPE_SOCKMAP.md)
 * [BPF_MAP_TYPE_LRU_HASH](../map-type/BPF_MAP_TYPE_LRU_HASH.md)
 * [BPF_MAP_TYPE_HASH_OF_MAPS](../map-type/BPF_MAP_TYPE_HASH_OF_MAPS.md)
 * [BPF_MAP_TYPE_LPM_TRIE](../map-type/BPF_MAP_TYPE_LPM_TRIE.md)
 * [BPF_MAP_TYPE_ARRAY_OF_MAPS](../map-type/BPF_MAP_TYPE_ARRAY_OF_MAPS.md)
 * [BPF_MAP_TYPE_ARRAY](../map-type/BPF_MAP_TYPE_ARRAY.md)
 * [BPF_MAP_TYPE_PERCPU_HASH](../map-type/BPF_MAP_TYPE_PERCPU_HASH.md)
 * [BPF_MAP_TYPE_LRU_PERCPU_HASH](../map-type/BPF_MAP_TYPE_LRU_PERCPU_HASH.md)
 * [BPF_MAP_TYPE_SOCKHASH](../map-type/BPF_MAP_TYPE_SOCKHASH.md)
 * [BPF_MAP_TYPE_PERCPU_ARRAY](../map-type/BPF_MAP_TYPE_PERCPU_ARRAY.md)
 * [BPF_MAP_TYPE_XSKMAP](../map-type/BPF_MAP_TYPE_XSKMAP.md)
 * [BPF_MAP_TYPE_HASH](../map-type/BPF_MAP_TYPE_HASH.md)
<!-- [/HELPER_FUNC_MAP_REF] -->

### Example

<!-- TODO add C / Rust example -->
