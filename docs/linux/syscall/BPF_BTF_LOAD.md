---
title: "Syscall command 'BPF_BTF_LOAD'"
description: "This page documents the 'BPF_BTF_LOAD' eBPF syscall command, including its definition, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_BTF_LOAD` command

<!-- [FEATURE_TAG](BPF_BTF_LOAD) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/f56a653c1fd13a197076dec4461c656fd2adec73)
<!-- [/FEATURE_TAG] -->

This command loads a BTF object into the kernel.

## Return value

This command will return a file descriptor to the created BTF object on success (positive integer) or an error number (negative integer) if something went wrong.

## Attributes

```c
union bpf_attr {
    struct {
		__aligned_u64   [btf](#btf);
		__aligned_u64   [btf_log_buf](#btf_log_buf);
		__u32           [btf_size](#btf_size);
		__u32           [btf_log_size](#btf_log_size);
		__u32           [btf_log_level](#btf_log_level);
		__u32           [btf_log_true_size](#btf_log_true_size);
		__u32           [btf_flags](#btf_flags);
		__s32           [btf_token_fd](#btf_token_fd);
	};
};
```

### `btf`

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/f56a653c1fd13a197076dec4461c656fd2adec73)

This field is a pointer to the BTF information to be loaded.

### `btf_log_buf`

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/f56a653c1fd13a197076dec4461c656fd2adec73)

This field is a pointer to a reserved piece of memory where the kernel will write the log to, if enabled.

### `btf_size`

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/f56a653c1fd13a197076dec4461c656fd2adec73)

This field is the size of the info indicated by `btf`.

### `btf_log_size`

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/f56a653c1fd13a197076dec4461c656fd2adec73)

This field is the size of the memory region indicated by `btf_log_buf`

### `btf_log_level`

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/f56a653c1fd13a197076dec4461c656fd2adec73)

The lower 2 bits of this value are the log level:

* `0` = no log
* `1` = basic logging   (`BPF_LOG_LEVEL1`)
* `2` = verbose logging (`BPF_LOG_LEVEL2`)

The remaining bits are flags:

* `1 << 3` (`BPF_LOG_STATS`) If set the kernel will output statistics to the log. Flags can be used since [:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/06ee7115b0d1742de745ad143fb5e06d77d27fba)

* `1 << 4` (`BPF_LOG_FIXED`) since [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/1216640938035e63bdbd32438e91c9bcc1fd8ee1), the verifier log rotates instead of truncating. When `log_size` is exceeded. Setting this flag preserves the old behavior of truncating the log to `log_size` bytes.

### `btf_flags`

[:octicons-tag-24: v6.9](https://github.com/torvalds/linux/commit/9ea7c4bf17e39d463eb4782f948f401d9764b1b3)

This field is a bitmask of flags that control the behavior of the BTF loading process. See the [Flags](#flags) section for more information.

### `btf_token_fd`

[:octicons-tag-24: v6.9](https://github.com/torvalds/linux/commit/9ea7c4bf17e39d463eb4782f948f401d9764b1b3)

The file descriptor of a [BPF token](../../linux/concepts/token.md) can be passed to this attribute. If the BPF token grants permission to create a BTF object, the kernel will allow the program to be loaded for a user without `CAP_BPF`.

## Flags

### `BPF_F_TOKEN_FD`

[:octicons-tag-24: v6.9](https://github.com/torvalds/linux/commit/9ea7c4bf17e39d463eb4782f948f401d9764b1b3)

This flag indicates that the [`btf_token_fd`](#btf_token_fd) attribute is being used. If this flag is not set, the kernel will ignore the [`btf_token_fd`](#btf_token_fd) attribute.
