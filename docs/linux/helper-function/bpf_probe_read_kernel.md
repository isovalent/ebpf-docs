---
title: "Helper Function 'bpf_probe_read_kernel'"
description: "This page documents the 'bpf_probe_read_kernel' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_probe_read_kernel`

<!-- [FEATURE_TAG](bpf_probe_read_kernel) -->
[:octicons-tag-24: v5.5](https://github.com/torvalds/linux/commit/6ae08ae3dea2cfa03dd3665a3c8475c2d429ef47)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Safely attempt to read _size_ bytes from kernel space address _unsafe_ptr_ and store the data in _dst_.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (* const bpf_probe_read_kernel)(void *dst, __u32 size, const void *unsafe_ptr) = (void *) 113;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

This helper could be very helpful in cases, when we need to access some global kernel variable or structure. It should be also used in case, when BPF program works in cooperation with some particular kernel module and needs to access to its global variables or structures to obtain some configuration settings, counters values, etc.

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_CGROUP_DEVICE`](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md)
 * [`BPF_PROG_TYPE_CGROUP_SKB`](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCK`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCKOPT`](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md)
 * [`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [`BPF_PROG_TYPE_CGROUP_SYSCTL`](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
 * [`BPF_PROG_TYPE_FLOW_DISSECTOR`](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md)
 * [`BPF_PROG_TYPE_KPROBE`](../program-type/BPF_PROG_TYPE_KPROBE.md)
 * [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
 * [`BPF_PROG_TYPE_LWT_IN`](../program-type/BPF_PROG_TYPE_LWT_IN.md)
 * [`BPF_PROG_TYPE_LWT_OUT`](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
 * [`BPF_PROG_TYPE_LWT_SEG6LOCAL`](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
 * [`BPF_PROG_TYPE_LWT_XMIT`](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [`BPF_PROG_TYPE_NETFILTER`](../program-type/BPF_PROG_TYPE_NETFILTER.md)
 * [`BPF_PROG_TYPE_PERF_EVENT`](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE`](../program-type/BPF_PROG_TYPE_RAW_TRACEPOINT_WRITABLE.md)
 * [`BPF_PROG_TYPE_SCHED_ACT`](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [`BPF_PROG_TYPE_SK_LOOKUP`](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md)
 * [`BPF_PROG_TYPE_SK_MSG`](../program-type/BPF_PROG_TYPE_SK_MSG.md)
 * [`BPF_PROG_TYPE_SK_REUSEPORT`](../program-type/BPF_PROG_TYPE_SK_REUSEPORT.md)
 * [`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md)
 * [`BPF_PROG_TYPE_SOCKET_FILTER`](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md)
 * [`BPF_PROG_TYPE_SOCK_OPS`](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
 * [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
 * [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [`BPF_PROG_TYPE_TRACEPOINT`](../program-type/BPF_PROG_TYPE_TRACEPOINT.md)
 * [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
 * [`BPF_PROG_TYPE_XDP`](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

```c
/* Some loadable kernel module with global variable to share */

#include <linux/init.h>       // Macros for module initialization
#include <linux/module.h>     // Core header for loading modules
#include <linux/kernel.h>     // Kernel logging macros

const int test_kmod_var = 666;

static int __init test_kmod_init(void)
{
    printk(KERN_INFO "Hello, world, test_var=%d\n", test_kmod_var);

    return 0;
}

static void __exit test_kmod_exit(void)
{
    printk(KERN_INFO "Goodbye, world!\n");
}

module_init(test_kmod_init);
module_exit(test_kmod_exit);

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Your Name");
MODULE_DESCRIPTION("A simple module");
MODULE_VERSION("1.0");

/* BPF program, which uses global variables from kernel
 * and from some loadable module
 */

#include "vmlinux.h"
#include <bpf/bpf_helpers.h>
#include "bpf_tracing_net.h"

extern const void test_kmod_var __ksym;
extern const void jiffies __ksym;

SEC("xdp")
int xdp_test(struct xdp_md *xdp) {

	u64 kernel_jiffies;
	u32 loadable_kmod_var;

	bpf_probe_read_kernel(&kernel_jiffies, 8, &jiffies);

	bpf_printk(">>> %s: Can access jiffies=%lu\n", __func__, kernel_jiffies);

	bpf_probe_read_kernel(&loadable_kmod_var, 4, &test_kmod_var);

	bpf_printk(">>> %s: Can also access test_kmod_var=%lu\n", __func__, loadable_kmod_var);

	return XDP_PASS;
}
```
