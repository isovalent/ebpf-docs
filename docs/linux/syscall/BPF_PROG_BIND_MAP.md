---
title: "Syscall command 'BPF_PROG_BIND_MAP'"
description: "This page documents the 'BPF_PROG_BIND_MAP' eBPF syscall command, including its definition, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_PROG_BIND_MAP` command

<!-- [FEATURE_TAG](BPF_PROG_BIND_MAP) -->
[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/ef15314aa5de955c6afd87d512e8b00f5ac08d06)
<!-- [/FEATURE_TAG] -->

This syscall binds a map to a program.

## Return value

## Attributes

### `prog_fd`

The file descriptor for the program to bind the map to.

### `map_fd`

The file descriptor of the map to be bound.

### `flags`

Flags, currently not use, so always zero.

## Usage

Normally a map is bound to a program when a program uses a map directly. What this syscall allows us to do is to associate (bind) the map to an already loaded program in the same way. So the program will increment the reference count on the map so it will stay loaded as long as the program is loaded.

The intended use case is to bind array maps containing metadata to programs. These would not be referenced by the program itself, hence the need for the binding. Userspace can use these maps to store information such as commit hashes and other auxiliary information there.
