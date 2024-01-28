# Helper function `bpf_inode_storage_get`

<!-- [FEATURE_TAG](bpf_inode_storage_get) -->
[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/8ea636848aca35b9f97c5b5dee30225cf2dd0fe6)
<!-- [/FEATURE_TAG] -->

## Defintion

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Get a bpf_local_storage from an _inode_.

Logically, it could be thought of as getting the value from a _map_ with _inode_ as the **key**.  From this perspective,  the usage is not much different from **bpf_map_lookup_elem**(_map_, **&**_inode_) except this helper enforces the key must be an inode and the map must also be a **BPF_MAP_TYPE_INODE_STORAGE**.

Underneath, the value is stored locally at _inode_ instead of the _map_.  The _map_ is used as the bpf-local-storage "type". The bpf-local-storage "type" (i.e. the _map_) is searched against all bpf_local_storage residing at _inode_.

An optional _flags_ (**BPF_LOCAL_STORAGE_GET_F_CREATE**) can be used such that a new bpf_local_storage will be created if one does not exist.  _value_ can be used together with **BPF_LOCAL_STORAGE_GET_F_CREATE** to specify the initial value of a bpf_local_storage.  If _value_ is **NULL**, the new bpf_local_storage will be zero initialized.

### Returns

A bpf_local_storage pointer is returned on success.

**NULL** if not found or there was an error in adding a new bpf_local_storage.

`#!c static void *(*bpf_inode_storage_get)(void *map, void *inode, void *value, __u64 flags) = (void *) 145;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
