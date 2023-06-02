# Helper function `bpf_sysctl_set_new_value`

<!-- [FEATURE_TAG](bpf_sysctl_set_new_value) -->
[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/4e63acdff864654cee0ac5aaeda3913798ee78f6)
<!-- [/FEATURE_TAG] -->

## Definition

<!-- [HELPER_FUNC_DEF] -->
Override new value being written by user space to sysctl with value provided by program in buffer _buf_ of size _buf_len_.

_buf_ should contain a string in same form as provided by user space on sysctl write.

User space may write new value at file position > 0. To override the whole sysctl value file position should be set to zero.

### Returns

0 on success.

**-E2BIG** if the _buf_len_ is too big.

**-EINVAL** if sysctl is being read.

`#!c static long (*bpf_sysctl_set_new_value)(struct bpf_sysctl *ctx, const char *buf, unsigned long buf_len) = (void *) 104;`
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
