# BPF Syscall `BPF_MAP_GET_NEXT_KEY` command

<!-- [FEATURE_TAG](BPF_MAP_GET_NEXT_KEY) -->
[:octicons-tag-24: v3.18](https://github.com/torvalds/linux/commit/db20fd2b01087bdfbe30bce314a198eefedcc42e)
<!-- [/FEATURE_TAG] -->


The `BPF_MAP_GET_NEXT_KEY` command is used to iterate over the keys of a map.

## Behavior and return value

If the given `key` isn't found within the map, `0` is returned and the value pointed to by `next_key` will be set to the first key in the map. 

If the given `key` is found withing the map, `0` is returned and the value pointed to by `next_key` will be set to the key after the current key.

If the given `key` is found within the map and it is the last key, an error number of `-ENOENT` will be returned.

## Attributes

### `map_fd`

This attribute specifies the file descriptor of the map in which you wish to lookup a value.

### `key`

This attribute holds a **pointer** to the current key, a null pointer can be passed in to indicate you have no current value. The size of the key is derived from the [`key_size`](BPF_MAP_CREATE.md#key_size) attribute of the map you specified with `map_fd`.

The memory indicated by this field will not be modified.

### `next_key`

This attribute holds a **pointer** to a memory location where the kernel will write the next key to. The size of the key is derived from the [`key_size`](BPF_MAP_CREATE.md#key_size) attribute of the map you specified with `map_fd`.

