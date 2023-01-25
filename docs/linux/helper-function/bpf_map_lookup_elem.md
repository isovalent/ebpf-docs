# Helper call `bpf_map_lookup_elem`

The lookup map element helper call is used to read values from [maps](../index.md#maps).

!!! note
    Not all [map types](../map-type/index.md) support this helper call due to their implementation, check the map type page for details.

## Usage

`#!c static void *(*bpf_map_lookup_elem)(void *map, const void *key) = 1;`
<!-- TODO rust signature? -->

Argument to this helper are `map` which is a pointer to a map definition and `key` which is a pointer to they key you
wish to lookup.

The return value will be a pointer to the map value or `NULL`. The value is a direct reference to the kernel memory where this map value is stored, not a copy. Therefor any modifications made to the value are automatically persisted without the need to call any additional helpers.

!!! warning
    modifying map values of non per-CPU maps is subject to race conditions, atomics or spinlocks must be utilized to prevent race conditions if they are detrimental to your use case.
    <!-- TODO link to guide on memory access serialization -->

### Example

<!-- TODO add C / Rust example -->
