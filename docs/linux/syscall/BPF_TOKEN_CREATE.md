---
title: "Syscall command 'BPF_TOKEN_CREATE'"
description: "This page documents the 'BPF_TOKEN_CREATE' eBPF syscall command, including its definition, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_TOKEN_CREATE` command

<!-- [FEATURE_TAG](BPF_TOKEN_CREATE) -->
[:octicons-tag-24: v6.9](https://github.com/torvalds/linux/commit/35f96de04127d332a5c5e8a155d31f452f88c76d)
<!-- [/FEATURE_TAG] -->

This syscall command creates a [token](../concepts/token.md) from a BPF file system with delegated privileges.

## Return value

This command will return the file descriptor of the token (positive integer) or an error number (negative integer) if something went wrong.

## Attributes

### `flags`

A field for flags, currently unused.

### `bpffs_fd`

The file descriptor to the root of a BPF file system.

## Usage

Tokens allow a privileged process (with `CAP_SYS_ADMIN` running in the init user namespace) to delegate permission to do certain BPF related operations to a non-privileged process. Normally unprivileged processes have little capabilities, the can load a few program types and maps, unless the `unprivileged_bpf_disabled` syscall is set. But with tokens, unprivileged processes gain near root level capabilities. A token can only be obtained if the process has `CAP_BPF`, which is the only capability the process needs to be able to use BPF with tokens.

A token has its own "policy" for which actions can or cannot be executed, set by the delegating process. A token can allow or disallow access to individual syscall commands, program types, map types, and attach types. So for example, a process can be given permission to load [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md) programs, but not any other program types.

Normally, using certain BPF features requires `CAP_NET_ADMIN` or `CAP_PERFMON`, having a token allows you to use these features (if the token permits) without having the capabilities.

So the use case is isolation. Before tokens, the loader processes needed to have capabilities to do certain actions. However, these capabilities are to broad. Giving a process `CAP_NET_ADMIN` for example also allows it to do a great deal of modifying and configuration of network interfaces, something we may not want. So a token allows us to setup an environment where a process can still perform BPF operations, but without the overly broad capabilities. In addition, the token allows for scoping down what can or can not be done, allowing us to permit only exactly what a process requires, and nothing more, to minimize what attackers could leverage if a process were to get compromised.
