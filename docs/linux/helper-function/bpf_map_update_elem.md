# Helper call `bpf_map_update_elem`

<!-- [FEATURE_TAG](bpf_map_update_elem) -->
[:octicons-tag-24: v3.18](https://github.com/torvalds/linux/commit/0a542a86d73b1577e7d4f55fc95dcffd3fe62643)
<!-- [/FEATURE_TAG] -->

The update map element helper call is used to write values from [maps](../index.md#maps).

## Definition

`#!c static long (*bpf_map_update_elem)(void *map, const void *key, const void *value, __u64 flags) = (void *) 2;`

## Usage

Arguments of this helper are `map` which is a pointer to a map definition, `key` which is a pointer to the key you
wish to write to, `value` which is a pointer to the value you wish to write to the map, and `flags` which are described below.

The `flags` argument can be one of the following values:

* `BPF_NOEXIST` - If set the update will only happen if the key doesn't exist yet, to prevent overwriting existing data.
* `BPF_EXIST` - If set the update will only happen if the key exists, to ensure an update and no new key creation.
* `BPF_ANY` - It doesn't matter, an update will be attempted in both cases.

!!! info
    `BPF_NOEXIST` isn't supported for array type maps since all keys always exist.

The return value will be `0` on success or a negative valued error number indicating a failure.


### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_DEVICE](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md)
 * [BPF_PROG_TYPE_CGROUP_SKB](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
 * [BPF_PROG_TYPE_CGROUP_SOCKOPT](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK_ADDR](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [BPF_PROG_TYPE_CGROUP_SYSCTL](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
 * [BPF_PROG_TYPE_FLOW_DISSECTOR](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md)
 * [BPF_PROG_TYPE_KPROBE](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [BPF_PROG_TYPE_LIRC_MODE2](../program-type/BPF_PROG_TYPE_LIRC_MODE2.md)
 * [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
 * [BPF_PROG_TYPE_LWT_IN](../program-type/BPF_PROG_TYPE_LWT_IN.md)
 * [BPF_PROG_TYPE_LWT_OUT](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
 * [BPF_PROG_TYPE_LWT_SEG6LOCAL](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
 * [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [BPF_PROG_TYPE_PERF_EVENT](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [BPF_PROG_TYPE_SK_LOOKUP](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md)
 * [BPF_PROG_TYPE_SK_MSG](../program-type/BPF_PROG_TYPE_SK_MSG.md)
 * [BPF_PROG_TYPE_SK_REUSEPORT](../program-type/BPF_PROG_TYPE_SK_REUSEPORT.md)
 * [BPF_PROG_TYPE_SK_SKB](../program-type/BPF_PROG_TYPE_SK_SKB.md)
 * [BPF_PROG_TYPE_SOCKET_FILTER](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
 * [BPF_PROG_TYPE_SOCK_OPS](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
 * [BPF_PROG_TYPE_SYSCALL](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [BPF_PROG_TYPE_TRACEPOINT](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
 * [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Map types

This helper call can be used with the following map types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_MAP_REF] -->
 * [BPF_MAP_TYPE_ARRAY](../map-type/BPF_MAP_TYPE_ARRAY.md)
 * [BPF_MAP_TYPE_HASH](../map-type/BPF_MAP_TYPE_HASH.md)
 * [BPF_MAP_TYPE_LPM_TRIE](../map-type/BPF_MAP_TYPE_LPM_TRIE.md)
 * [BPF_MAP_TYPE_LRU_HASH](../map-type/BPF_MAP_TYPE_LRU_HASH.md)
 * [BPF_MAP_TYPE_LRU_PERCPU_HASH](../map-type/BPF_MAP_TYPE_LRU_PERCPU_HASH.md)
 * [BPF_MAP_TYPE_PERCPU_ARRAY](../map-type/BPF_MAP_TYPE_PERCPU_ARRAY.md)
 * [BPF_MAP_TYPE_PERCPU_HASH](../map-type/BPF_MAP_TYPE_PERCPU_HASH.md)
 * [BPF_MAP_TYPE_SOCKHASH](../map-type/BPF_MAP_TYPE_SOCKHASH.md)
 * [BPF_MAP_TYPE_SOCKMAP](../map-type/BPF_MAP_TYPE_SOCKMAP.md)
<!-- [/HELPER_FUNC_MAP_REF] -->


### Example

<!-- TODO add C / Rust example -->
