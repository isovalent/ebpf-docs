---
title: "Syscall command 'BPF_MAP_FREEZE'"
description: "This page documents the 'BPF_MAP_FREEZE' eBPF syscall command, including its definition, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_MAP_FREEZE` command

<!-- [FEATURE_TAG](BPF_MAP_FREEZE) -->
[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/87df15de441bd4add7876ef584da8cabdd9a042a)
<!-- [/FEATURE_TAG] -->

This syscall command "freezes" a map, making its contents read-only from then on. This operation cannot be undone.

## Return value

This command will return a zero on success or an error number (negative integer) if something went wrong.

## Attributes

### `map_fd`

File descriptor of the map to freeze.

## Usage

The primary purpose of freezing a map is to allow a loader to provide map contents to the program (and verifier), and to give the guarantee that the contents of that map will not change for the duration of the lifetime of the program. 

A specific example of where this is useful is for global constants. When a global constant is define, the compiler puts it in the `.rodata` ELF section, this section is turned into an array map, which gets frozen before loading. The verifier will see that this map is frozen, and that the values within are thus truly constant and will treat the map contents as scalar values. This is significant because it allows the compiler to do [dead code elimination](../concepts/verifier.md#dead-code-elimination) at load time based on the contents of the map.
