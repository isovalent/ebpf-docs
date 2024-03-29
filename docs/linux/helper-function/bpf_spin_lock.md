---
title: "Helper Function 'bpf_spin_lock'"
description: "This page documents the 'bpf_spin_lock' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_spin_lock`

<!-- [FEATURE_TAG](bpf_spin_lock) -->
[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/d83525ca62cf8ebe3271d14c36fb900c294274a2)
<!-- [/FEATURE_TAG] -->

Acquire a spinlock represented by the pointer `lock`, which is
stored as part of a value of a map. Taking the lock allows to
safely update the rest of the fields in that value. The
spinlock can (and must) later be released with a call to
[`bpf_spin_unlock`](bpf_spin_unlock.md).

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


`#!c static long (* const bpf_spin_lock)(struct bpf_spin_lock *lock) = (void *) 93;`

## Usage

Spinlocks in BPF programs come with a number of restrictions
and constraints:

* `bpf_spin_lock` objects are only allowed inside maps of types [`BPF_MAP_TYPE_HASH`](../map-type/BPF_MAP_TYPE_HASH.md) and [`BPF_MAP_TYPE_ARRAY`](../map-type/BPF_MAP_TYPE_ARRAY.md) (this list could be extended in the future).
* BTF description of the map is mandatory.
* The BPF program can take **ONE** lock at a time, since taking two
or more could cause dead locks.
* Only one `struct bpf_spin_lock` is allowed per map element.
* When the lock is taken, calls (either BPF to BPF or helpers) are not allowed.
* The `BPF_LD_ABS` and `BPF_LD_IND` instructions are not allowed inside a spinlock-ed region.
* The BPF program **MUST** call `bpf_spin_unlock()` to release the lock, on all execution paths, before it returns.
* The BPF program can access `struct bpf_spin_lock` only via the `bpf_spin_lock()` and `bpf_spin_unlock()` helpers. Loading or storing data into the `struct bpf_spin_lock lock;` field of a map is not allowed.
* To use the `bpf_spin_lock()` helper, the BTF description of the map value must be a struct and have `struct bpf_spin_lock anyname;` field at the top level. Nested lock inside another struct is not allowed.
* The `struct bpf_spin_lock lock` field in a map value must be aligned on a multiple of 4 bytes in that value.
* Syscall with command [`BPF_MAP_LOOKUP_ELEM`](../syscall/BPF_MAP_LOOKUP_ELEM.md) does not copy the `bpf_spin_lock` field to user space.
* Syscall with command [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md), or update from a BPF program, do not update the `bpf_spin_lock` field.
* `bpf_spin_lock` cannot be on the stack or inside a networking packet (it can only be inside of a map values).
* `bpf_spin_lock` is available to root only.
* Tracing programs and [socket filter programs](../program-type/BPF_PROG_TYPE_SOCKET_FILTER.md) cannot use `bpf_spin_lock()` due to insufficient preemption checks (but this may change in the future).
* `bpf_spin_lock` is not allowed in inner maps of map-in-map.

### Program types

This helper call can be used in the following program types:

<!-- TODO use generated list as soon as we can exclude functions from inherited groups -->
 * [BPF_PROG_TYPE_CGROUP_DEVICE](../program-type/BPF_PROG_TYPE_CGROUP_DEVICE.md)
 * [BPF_PROG_TYPE_CGROUP_SKB](../program-type/BPF_PROG_TYPE_CGROUP_SKB.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK](../program-type/BPF_PROG_TYPE_CGROUP_SOCK.md)
 * [BPF_PROG_TYPE_CGROUP_SOCKOPT](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md)
 * [BPF_PROG_TYPE_CGROUP_SOCK_ADDR](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
 * [BPF_PROG_TYPE_CGROUP_SYSCTL](../program-type/BPF_PROG_TYPE_CGROUP_SYSCTL.md)
 * [BPF_PROG_TYPE_FLOW_DISSECTOR](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md)
 * [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
 * [BPF_PROG_TYPE_LWT_IN](../program-type/BPF_PROG_TYPE_LWT_IN.md)
 * [BPF_PROG_TYPE_LWT_OUT](../program-type/BPF_PROG_TYPE_LWT_OUT.md)
 * [BPF_PROG_TYPE_LWT_SEG6LOCAL](../program-type/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
 * [BPF_PROG_TYPE_LWT_XMIT](../program-type/BPF_PROG_TYPE_LWT_XMIT.md)
 * [BPF_PROG_TYPE_SCHED_ACT](../program-type/BPF_PROG_TYPE_SCHED_ACT.md)
 * [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
 * [BPF_PROG_TYPE_SK_LOOKUP](../program-type/BPF_PROG_TYPE_SK_LOOKUP.md)
 * [BPF_PROG_TYPE_SK_MSG](../program-type/BPF_PROG_TYPE_SK_MSG.md)
 * [BPF_PROG_TYPE_SK_REUSEPORT](../program-type/BPF_PROG_TYPE_SK_REUSEPORT.md)
 * [BPF_PROG_TYPE_SK_SKB](../program-type/BPF_PROG_TYPE_SK_SKB.md)
 * [BPF_PROG_TYPE_SOCK_OPS](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
 * [BPF_PROG_TYPE_SYSCALL](../program-type/BPF_PROG_TYPE_SYSCALL.md)
 * [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
 * [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)

!!! note
    `bpf_spin_lock` also can't be used if the program has been loaded with the `BPF_F_SLEEPABLE` flag.
    <!-- https://elixir.bootlin.com/linux/v6.1/source/kernel/bpf/verifier.c#L12691 -->


### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
