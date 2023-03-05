# BPF Syscall `BPF_MAP_DELETE_ELEM` command

<!-- [FEATURE_TAG](BPF_MAP_DELETE_ELEM) -->
[:octicons-tag-24: v3.18](https://github.com/torvalds/linux/commit/db20fd2b01087bdfbe30bce314a198eefedcc42e)
<!-- [/FEATURE_TAG] -->

This command deletes an element from a map.

!!! warning
    Not all map types (particularly array maps) support this operation, instead a zero value can be written to the map value. Check the map types page to check for support.

## Return value

This command will return `0` on success or a error number (negative integer) if something went wrong.

## Attributes
### `map_fd`

This field indices the file descriptor of the map from which you would like to delete an element.

### `key`

This fields indicates the key you wish to remove from the map. This field should contain a pointer to the key you wish delete. The size of the key should match the size of the key in the map indicated by the [`key_size`](../syscall/BPF_MAP_CREATE.md#value_size).
		
### `value`

This field is not used in this command.
			
### `flags`

This attribute is a bitmask of flags.

#### `BPF_ANY`

This flag has a value of `0`, so setting it together with another flag has no impact. It is meant to be used if no other flags are specified to explicitly state that the command should delete the element from the map regardless of if the key already exists or not.

#### `BPF_NOEXIST`

If this flag does not do anything for this command.

#### `BPF_EXISTS`

If this flag is set, the command will make sure that the given key already exists. If no entry for this key exists, the `-ENOENT` error number will be returned.

#### `BPF_F_LOCK`

If this flag is set, the command will acquire the spin-lock of the map value we are deleting. If the map contains no spin-lock in its value, `-EINVAL` will be returned by the command.
