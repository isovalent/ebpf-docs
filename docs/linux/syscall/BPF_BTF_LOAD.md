---
title: "Syscall command 'BPF_BTF_LOAD' - eBPF Docs"
description: "This page documents the 'BPF_BTF_LOAD' eBPF syscall command, including its defintion, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_BTF_LOAD` command

<!-- [FEATURE_TAG](BPF_BTF_LOAD) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/f56a653c1fd13a197076dec4461c656fd2adec73)
<!-- [/FEATURE_TAG] -->

This command loads a BTF object into the kernel.

## Return value

This command will return a file descriptor to the created BTF object on success (positive integer) or an error number (negative integer) if something went wrong.

## Attributes

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

This field is the level of detail contained in the log. Possible values are:

* `0` = no log
* `1` = basic logging   (`BPF_LOG_LEVEL1`)
* `2` = verbose logging (`BPF_LOG_LEVEL2`)

Additionally the 3rd bit is a flag, if set the kernel will output statistics to the log (`BPF_LOG_STATS`).
