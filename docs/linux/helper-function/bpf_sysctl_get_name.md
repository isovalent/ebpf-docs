# Helper function `bpf_sysctl_get_name`

<!-- [FEATURE_TAG](bpf_sysctl_get_name) -->
[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/808649fb787d918a48a360a668ee4ee9023f0c11)
<!-- [/FEATURE_TAG] -->

## Defintion

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Get name of sysctl in /proc/sys/ and copy it into provided by program buffer _buf_ of size _buf_len_.

The buffer is always NUL terminated, unless it's zero-sized.

If _flags_ is zero, full name (e.g. "net/ipv4/tcp_mem") is copied. Use **BPF_F_SYSCTL_BASE_NAME** flag to copy base name only (e.g. "tcp_mem").

### Returns

Number of character copied (not including the trailing NUL).

**-E2BIG** if the buffer wasn't big enough (_buf_ will contain truncated name in this case).

`#!c static long (*bpf_sysctl_get_name)(struct bpf_sysctl *ctx, char *buf, unsigned long buf_len, __u64 flags) = (void *) 101;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_SYSCTL](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
