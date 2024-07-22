---
title: "Helper Function 'bpf_getsockopt'"
description: "This page documents the 'bpf_getsockopt' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_getsockopt`

<!-- [FEATURE_TAG](bpf_getsockopt) -->
[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/cd86d1fd21025fdd6daf23d1288da405e7ad0ec6)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Emulate a call to **getsockopt()** on the socket associated to _bpf_socket_, which must be a full socket. The _level_ at which the option resides and the name _optname_ of the option must be specified, see **getsockopt(2)** for more information. The retrieved value is stored in the structure pointed by _opval_ and of length _optlen_.

_bpf_socket_ should be one of the following:

* **struct bpf_sock_ops** for **BPF_PROG_TYPE_SOCK_OPS**.
* **struct bpf_sock_addr** for **BPF_CGROUP_INET4_CONNECT**,
  **BPF_CGROUP_INET6_CONNECT** and **BPF_CGROUP_UNIX_CONNECT**.

This helper actually implements a subset of **getsockopt()**. It supports the same set of _optname_s that is supported by the **bpf_setsockopt**() helper.  The exceptions are **TCP_BPF_*** is **bpf_setsockopt**() only and **TCP_SAVED_SYN** is **bpf_getsockopt**() only.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_getsockopt)(void *bpf_socket, int level, int optname, void *optval, int optlen) = (void *) 57;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md) [:octicons-tag-24: v5.15](https://github.com/torvalds/linux/commit/2c531639deb5e3ddfd6e8123b82052b2d9fbc6e5)
 * [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md) [:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/beecf11bc2188067824591612151c4dc6ec383c7)
 * [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md) [:octicons-tag-24: v6.0](https://github.com/torvalds/linux/commit/9113d7e48e9128522b9f5a54dfd30dff10509a92)
 * [`BPF_PROG_TYPE_SOCK_OPS`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
