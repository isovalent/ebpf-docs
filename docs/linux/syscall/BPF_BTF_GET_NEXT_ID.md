# BPF Syscall `BPF_BTF_GET_NEXT_ID` command

<!-- [FEATURE_TAG](BPF_BTF_GET_NEXT_ID) -->
[:octicons-tag-24: v5.4](https://github.com/torvalds/linux/commit/1b9ed84ecf268904d89edf2908426a8eb3b5a4ba)
<!-- [/FEATURE_TAG] -->

This syscall command is used to iterate over all loaded BTF objects.

## Return type

This command will return `0` on success or a error number (negative integer) if something went wrong.

## Usage

This syscall command will populate the `next_id` field with the ID of the "next" BTF object which will have a higher number than `start_id`. If no BTF object IDs are known, `start_id` can be left at `0`. If no BTF objects exist higher than `start_id`, `next_id` is set to `-1` and the syscall will return an `-ENOENT` error code.

So to iterate or discover all loaded BTF objects: 

1. call this command repeatably with the same attribute pointer and the attributes initialized at zero
2. move `next_id` to `start_id` between each call
3. record all `next_id` values
4. stop when we get an error

The IDs returned by this command can be used with the [`BPF_BTF_GET_FD_BY_ID`](BPF_BTF_GET_FD_BY_ID.md) syscall command to get a file descriptor to the actual BTF object.

## Attributes

### `start_id`

The ID from which we wish to start iterating. `next_id` will always be higher than this field.

### `next_id`

This field will be set to the next BTF object ID, or `-1` if no next BTF object exists.
