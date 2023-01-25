# Helper call `bpf_map_delete_elem`

The delete map element helper call is used to delete values from [maps](../index.md#maps).

!!! note
    Not all [map types](../map-type/index.md) support this helper call due to their implementation, check the map type page for details.

## Usage

`#!c static long (*bpf_map_delete_elem)(void *map, const void *key) = (void *) 3;`
<!-- TODO rust signature? -->

Arguments of this helper are `map` which is a pointer to a map definition and `key` which is a pointer to the key you
wish to delete.

The return value will be `0` on success or a negative valued error number indicating a failure.

### Example

<!-- TODO add C / Rust example -->
