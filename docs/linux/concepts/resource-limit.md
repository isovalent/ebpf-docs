---
title: "Resource limits"
description: "This page explains the concept of resource limits in eBPF. It explains what resource limits are and how resource limits work."
---
# Resource limits

The Linux kernel has protection mechanisms that prevent processes from taking up too much memory. Since BPF maps can take up a lot of space, they are also limited via these mechanisms.

## Rlimit

rlimit or "resource limit" is a system to track and limit the amount of certain resources you are allowed to use. One of the things it limits is the amount of "locked memory" https://man7.org/linux/man-pages/man2/getrlimit.2.html

Until kernel version v5.11 this mechanism was used to track and limit the memory usage of BPF maps which count towards the locked memory limit, so you commonly would have to increase or disable this rlimit which requires an additional capability `CAP_SYS_RESOURCE`.

## cGroup memory limit

In the v5.11 kernel update, [this patch set](https://lore.kernel.org/bpf/20201201215900.3569844-1-guro@fb.com/) switched the memory accounting and limiting from rlimit to cGroups. This means that all memory used adds to the "memory used" figure of the cGroup of which the process that creates it is a part. This eliminates the need to grant loaders `CAP_SYS_RESOURCE` capability. If resource limits need to be raised, it should be done so with the `memory.max` setting on the cGroup.

!!! note
    Kernel memory accounting and limiting per cGroup can be disabled by disabling the `MEMCG_KMEM` kconfig during kernel compilation which is set to `y` by default.

!!! note
    [:octicons-tag-24: v6.3](https://github.com/torvalds/linux/commit/b6c1a8af5b1eec42aabc13376f94aa90c3d765f1 ) adds a new kernel parameter `cgroup.memory=nobpf` which disables memory accounting and limiting for BPF as well.
