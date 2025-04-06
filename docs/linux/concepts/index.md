---
title: Linux eBPF concepts
description: An index of Linux specific eBPF concepts.
hide: toc
---
# Linux eBPF concepts

This is an index of Linux specific eBPF concepts and features. For more generic eBPF concepts that are not Linux specific, see the [eBPF concepts](../../concepts/index.md) page.

<div class="grid cards" markdown>

-   __Maps__

    ---

    Maps allow for data storage and communication

    [:octicons-arrow-right-24: Maps](./maps.md)

-   __Verifier__

    ---

    The verifier checks the safety of eBPF programs

    [:octicons-arrow-right-24: Verifier](./verifier.md)

-  __Functions__

    ---

    This page explains how functions work for eBPF on Linux

    [:octicons-arrow-right-24: Functions](./functions.md)

-  __Concurrency__

    ---

    This page explains the effects of concurrency on eBPF programs and how to handle it

    [:octicons-arrow-right-24: Concurrency](./concurrency.md)

-  __Pinning__

    ---

    Pinning allows the file system to reference eBPF objects and keep them alive

    [:octicons-arrow-right-24: Pinning](./pinning.md)

-  __Tail calls__

    ---

    Tail calls allow for the chaining of eBPF programs

    [:octicons-arrow-right-24: Tail calls](./tail-calls.md)

-  __Loops__

    ---

    Loops in eBPF are not trivial, this page explains how to use different types of loops

    [:octicons-arrow-right-24: Loops](./loops.md)

- __Timers__

    ---

    Timers allow for the scheduling of eBPF functions to execute at a later time

    [:octicons-arrow-right-24: Timers](./timers.md)

- __Resource Limit__

    ---

    This page explains how the Linux kernel counts and restricts the resources used by eBPF

    [:octicons-arrow-right-24: Resource Limit](./resource-limit.md)

- __AF_XDP__

    ---

    AF_XDP allows you to bypass the kernel network stack and process packets in userspace

    [:octicons-arrow-right-24: AF_XDP](./af_xdp.md)

- __KFuncs__

    ---

    KFuncs allow for the calling of kernel functions from eBPF programs

    [:octicons-arrow-right-24: KFuncs](./kfuncs.md)

- __Dynamic pointers__

    ---

    Dynamic pointers are pointers with metadata, moving memory safety checks to runtime

    [:octicons-arrow-right-24: Dynptrs](./dynptrs.md)

- __eBPF Tokens__

    ---

    eBPF tokens are like authentication tokens for eBPF operations

    [:octicons-arrow-right-24: Token](./token.md)

- __Trampolines__

    ---

    Trampolines are used to attach eBPF programs to kernel functions

    [:octicons-arrow-right-24: Trampolines](./trampolines.md)

- __USDT__

    ---

    This page explain USDT (User Statically-Defined Tracing)

    [:octicons-arrow-right-24: USDT](./usdt.md)

</div>
