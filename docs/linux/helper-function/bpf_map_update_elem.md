# Helper call `bpf_map_update_elem`

The update map element helper call is used to write values from [maps](../index.md#maps).

!!! note
    Not all [map types](../map-type/index.md) support this helper call due to their implementation, check the map type page for details.

## Usage

`#!c static long (*bpf_map_update_elem)(void *map, const void *key, const void *value, __u64 flags) = (void *) 2;`
<!-- TODO rust signature? -->

Arguments of this helper are `map` which is a pointer to a map definition, `key` which is a pointer to the key you
wish to write to, `value` which is a pointer to the value you wish to write to the map, and `flags` which are described below.

The `flags` argument can be one of the following values:

* `BPF_NOEXIST` - If set the update will only happen if the key doesn't exist yet, to prevent overwriting existing data.
* `BPF_EXIST` - If set the update will only happen if the key exists, to ensure an update and no new key creation.
* `BPF_ANY` - It doesn't matter, an update will be attempted in both cases.

!!! info
    `BPF_NOEXIST` isn't supported for array type maps since all keys always exist.

The return value will be `0` on success or a negative valued error number indicating a failure.

### Example

<!-- TODO add C / Rust example -->
