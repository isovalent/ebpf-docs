---
title: "Libbpf eBPF macro 'KERNEL_VERSION'"
description: "This page documents the 'KERNEL_VERSION' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `KERNEL_VERSION`

[:octicons-tag-24: v0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)

The `KERNEL_VERSION` macros is used to convert from three part {major}.{minor}.{patch} version number to a single integer.

## Definition

`#!c #define KERNEL_VERSION(a, b, c) (((a) << 16) + ((b) << 8) + ((c) > 255 ? 255 : (c)))`

## Usage

This macro can be used together with the kernel version of the machine the program is loaded on. To get this kernel version you define `extern int LINUX_KERNEL_VERSION __kconfig;`. The `LINUX_KERNEL_VERSION` variable has a special meaning an will be resolved by the loader (library). It is encoded the same way as `KERNEL_VERSION`, allowing you to compare the two values. This allows you to write CO-RE like code that can adapt to different kernel versions.

!!! warning
    Version numbers do not always reflect the actual features available in the kernel. Some distributions backport features to older kernels without reflecting this in the version number. If possible, it is always better to probe for the availability of a feature directly instead of inferring it from the kernel version.

### Example

```c hl_lines="6"
extern int LINUX_KERNEL_VERSION __kconfig;

SEC("xdp")
int example_prog(struct xdp_md *ctx)
{
    if (LINUX_KERNEL_VERSION < KERNEL_VERSION(5, 10, 0)) {
        // Fall back to old behavior
    } else {
        // Use new features
    }
}
```
