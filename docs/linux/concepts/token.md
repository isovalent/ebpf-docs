---
title: "Token"
description: "This page explains the concept of BPF Token."
---
# BPF Token 

<!-- [FEATURE_TAG](bpf_token) -->
[:octicons-tag-24: v6.9](https://github.com/torvalds/linux/commit/35f96de04127d332a5c5e8a155d31f452f88c76d)
<!-- [/FEATURE_TAG] -->

BPF Token is a mechanism for delegating some of the BPF subsystem functionalities to an unprivileged process (e.g., container) within [user-namespace](https://man7.org/linux/man-pages/man7/user_namespaces.7.html) from a privileged process (e.g., container runtime) in the init-namespace.

## eBPF and the Linux Capabilities

When eBPF was first introduced in the Linux kernel, it required `CAP_SYS_ADMIN` to load programs, create maps, etc. However, having `CAP_SYS_ADMIN` for eBPF applications gives too many privileges beyond just interacting with the BPF subsystem.

That's why `CAP_BPF` has been introduced since v5.8. It allows more granular control of which eBPF functionalities the process can use. For example, to load network-related eBPF programs such as TC or XDP, it requires `CAP_BPF + CAP_NET_ADMIN`; to load tracing-related eBPF programs like kprobe or raw_tracepoint, it requires `CAP_BPF + CAP_PERFMON` and so on.

## eBPF and User Namespaces

[User Namespace](https://man7.org/linux/man-pages/man7/user_namespaces.7.html) is a namespace that isolates UID, GID, keys, and capabilities. Within the User Namespace, a process can behave like having the root privilege, but only for the resources governed by the User Namespaces.

For example, the process within the User Namespace can even have a `CAP_NET_ADMIN` and create a new Network Namespace and network devices within it. However, it cannot connect a `veth` device with the init namespace because creating a device within the init namespace requires a `CAP_NET_ADMIN` in the init namespace.

Then what about the eBPF? Can the process that owns `CAP_BPF + CAP_PERFMON` within the User Namespace load tracing eBPF programs? No, it's not allowed. eBPF, by its nature, can see many sensitive behaviors within the kernel. Therefore, it's unsafe to let unprivileged users load eBPF programs.

## Privilege Delegation by BPF Token

It's unacceptable for the kernel to let unprivileged users load the eBPF program. However, if a privileged process delegates certain privileges, that's acceptable. This is where BPF Token comes in.

Here's a rough procedure for using a BPF Token.

1. The unprivileged process creates and enters into the User Namespace.
2. The unprivileged process creates Mount Namespace and obtains BPFFS File Descriptor (FD) using `fsopen(2)`
3. The unprivileged process sends BPFFS FD over the Unix Domain Socket to the privileged process.
4. The privileged process configures BPFFS FD with the delegated privileges with `fsconfig(2)` and instantiates it with the `FSCONFIG_CMD_CREATE`.
5. The privileged process creates a Mount FD using `fsmount(2)` and sends it back to the unprivileged process.
6. The unprivileged process creates a BPF Token FD using the `BPF_TOKEN_CREATE` command of the `bpf(2)`.
7. The unprivileged process calls `bpf(2)` commands like `BPF_PROG_LOAD` with BPF Token FD.

It's a highly complex procedure, but eBPF applications usually don't have to perform it. Steps 1 - 5 would be performed by the privileged process, such as container runtime, and eBPF loader libraries like `libbpf` take care of step 6 and transparently pass it to the `bpf(2)`. Note that the BPF Token File Descriptor is bound to the User Namespace that creates it and can't be used on the outside.

## Delegation Options

Currently, the following options are available for `fsconfig(2)`.

- `delegate_cmds`: Allows performing specific BPF system call commands
- `delegate_maps`: Allows creating specific types of maps
- `delegate_progs`: Allows loading specific types of programs
- `delegate_attachs`: Allows specific attach types

## References

- [Kernel support patch set](https://patchwork.kernel.org/project/linux-fsdevel/cover/20230919214800.3803828-1-andrii@kernel.org/)
- [Test code](https://github.com/torvalds/linux/blob/master/tools/testing/selftests/bpf/prog_tests/token.c)
