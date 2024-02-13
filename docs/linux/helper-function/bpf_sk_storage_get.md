# Helper function `bpf_sk_storage_get`

<!-- [FEATURE_TAG](bpf_sk_storage_get) -->
[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/6ac99e8f23d4b10258406ca0dd7bffca5f31da9d)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Get a bpf-local-storage from a _sk_.

Logically, it could be thought of getting the value from a _map_ with _sk_ as the **key**.  From this perspective,  the usage is not much different from **bpf_map_lookup_elem**(_map_, **&**_sk_) except this helper enforces the key must be a full socket and the map must be a **BPF_MAP_TYPE_SK_STORAGE** also.

Underneath, the value is stored locally at _sk_ instead of the _map_.  The _map_ is used as the bpf-local-storage "type". The bpf-local-storage "type" (i.e. the _map_) is searched against all bpf-local-storages residing at _sk_.

_sk_ is a kernel **struct sock** pointer for LSM program. _sk_ is a **struct bpf_sock** pointer for other program types.

An optional _flags_ (**BPF_SK_STORAGE_GET_F_CREATE**) can be used such that a new bpf-local-storage will be created if one does not exist.  _value_ can be used together with **BPF_SK_STORAGE_GET_F_CREATE** to specify the initial value of a bpf-local-storage.  If _value_ is **NULL**, the new bpf-local-storage will be zero initialized.

### Returns

A bpf-local-storage pointer is returned on success.

**NULL** if not found or there was an error in adding a new bpf-local-storage.

`#!c static void *(*bpf_sk_storage_get)(void *map, void *sk, void *value, __u64 flags) = (void *) 107;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_SKB](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
 * [BPF_PROG_TYPE_CGROUP_SOCKOPT](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK_ADDR](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [BPF_PROG_TYPE_SK_MSG](../program-type/BPF_PROG_TYPE_SK_MSG.md)
 * [BPF_PROG_TYPE_SOCK_OPS](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
