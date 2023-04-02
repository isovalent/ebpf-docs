---
title: Map types (Linux)
hide: toc
---

# Map types (Linux)

## Generic map types

These map types are not limited to a very specific use case but can be used in a number of different use cases to store data.

  * [BPF_MAP_TYPE_HASH](BPF_MAP_TYPE_HASH.md)
  * [BPF_MAP_TYPE_ARRAY](BPF_MAP_TYPE_ARRAY.md)
  * [BPF_MAP_TYPE_PERCPU_HASH](BPF_MAP_TYPE_PERCPU_HASH.md)
  * [BPF_MAP_TYPE_PERCPU_ARRAY](BPF_MAP_TYPE_PERCPU_ARRAY.md)
  * [BPF_MAP_TYPE_QUEUE](BPF_MAP_TYPE_QUEUE.md)
  * [BPF_MAP_TYPE_STACK](BPF_MAP_TYPE_STACK.md)
  * BPF_MAP_TYPE_LRU_HASH
  * BPF_MAP_TYPE_LRU_PERCPU_HASH
  * BPF_MAP_TYPE_LPM_TRIE
  * BPF_MAP_TYPE_BLOOM_FILTER

## Map in map

These map types hold references to other map types as their values.

  * [BPF_MAP_TYPE_ARRAY_OF_MAPS](BPF_MAP_TYPE_ARRAY_OF_MAPS.md)
  * BPF_MAP_TYPE_HASH_OF_MAPS

## Streaming

These map types allow for the streaming of large amounts of data across the kernel-userspace boundary.

  * [BPF_MAP_TYPE_PERF_EVENT_ARRAY](BPF_MAP_TYPE_PERF_EVENT_ARRAY.md)
  * BPF_MAP_TYPE_RINGBUF
  * BPF_MAP_TYPE_USER_RINGBUF

## Packet redirection

These map types facilitate the steering or redirection of packets, between network devices, logical CPUs, or sockets.

  * [BPF_MAP_TYPE_DEVMAP](BPF_MAP_TYPE_DEVMAP.md)
  * [BPF_MAP_TYPE_SOCKMAP](BPF_MAP_TYPE_SOCKMAP.md)
  * BPF_MAP_TYPE_CPUMAP
  * BPF_MAP_TYPE_XSKMAP
  * [BPF_MAP_TYPE_SOCKHASH](BPF_MAP_TYPE_SOCKHASH.md)
  * [BPF_MAP_TYPE_DEVMAP_HASH](BPF_MAP_TYPE_DEVMAP_HASH.md)
  * [BPF_MAP_TYPE_REUSEPORT_SOCKARRAY](BPF_MAP_TYPE_REUSEPORT_SOCKARRAY.md)

## Tail call map

This map types is used to perform tail calls.

  * [BPF_MAP_TYPE_PROG_ARRAY](BPF_MAP_TYPE_PROG_ARRAY.md)

## Object attached storage

These map types provide storage related to kernel objects for/on which a BPF program is called.

  * BPF_MAP_TYPE_CGROUP_STORAGE
  * BPF_MAP_TYPE_PERCPU_CGROUP_STORAGE
  * BPF_MAP_TYPE_SK_STORAGE
  * BPF_MAP_TYPE_INODE_STORAGE
  * BPF_MAP_TYPE_TASK_STORAGE
  * BPF_MAP_TYPE_CGRP_STORAGE

## cGroup array

This map type contains references to cGroups.

  * BPF_MAP_TYPE_CGROUP_ARRAY

## Stack trace

This map type allows eBPF programs to read the stack trace.

  * BPF_MAP_TYPE_STACK_TRACE

## Struct ops

This map type replaces a structure containing function pointers in the kernel and allows authors to replace existing function pointers with pointers to BPF programs.

  * BPF_MAP_TYPE_STRUCT_OPS
  
## Per CPU maps

There are a number of per-CPU map types like `BPF_MAP_TYPE_PERCPU_HASH` and `BPF_MAP_TYPE_PERCPU_ARRAY`. These are per-CPU variants of their base map types. Like the name implies these maps consists of multiple copies, one for each logical CPU on the host. Programs running in the context of CPU#0 for example will see different map contents as a program running on CPU#2. 

Since multiple CPUs will never read or write to memory being accessed by another CPU, it is impossible for race conditions to occur, and thus programs don't need to waste cycles on mechanisms like spin-locks or atomics to synchronize access. It also improves speed due to better cache locality.

Another interesting use of these map types are as scratch buffers. Since eBPF programs always execute on the same logical CPU for their entire execution including tail calls, these maps can be used to transfer information between tail calls, something that is difficult to do otherwise. In the same spirit, these maps can also be used to hold data without counting towards the stack limit of the eBPF program.

When per-CPU maps are accessed via user-space all copies are always accesses at the same time. In fact the values read and written to these maps via the syscalls are like arrays with a size equal to the logical CPU count of the host. For more details, see the `BPF_MAP_LOOKUP_ELEM` and `BPF_MAP_UPDATE_ELEM` pages.

The downside is that programs cannot share information across CPU boundaries. So these kinds of maps are typically better suited for use-cases where data flows from the programs to user space like counters or use-cases where related events always happen on the same logical CPU like incoming network packets of the same flow due to RSS
