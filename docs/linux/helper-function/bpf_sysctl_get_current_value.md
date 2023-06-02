# Helper function `bpf_sysctl_get_current_value`

<!-- [FEATURE_TAG](bpf_sysctl_get_current_value) -->
[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/1d11b3016cec4ed9770b98e82a61708c8f4926e7)
<!-- [/FEATURE_TAG] -->

## Definition

<!-- [HELPER_FUNC_DEF] -->
Get current value of sysctl as it is presented in /proc/sys (incl. newline, etc), and copy it as a string into provided by program buffer _buf_ of size _buf_len_.

The whole value is copied, no matter what file position user space issued e.g. sys_read at.

The buffer is always NUL terminated, unless it's zero-sized.

### Returns

Number of character copied (not including the trailing NUL).

**-E2BIG** if the buffer wasn't big enough (_buf_ will contain truncated name in this case).

**-EINVAL** if current value was unavailable, e.g. because sysctl is uninitialized and read returns -EIO for it.

`#!c static long (*bpf_sysctl_get_current_value)(struct bpf_sysctl *ctx, char *buf, unsigned long buf_len) = (void *) 102;`
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
