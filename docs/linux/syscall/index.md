---
title: Syscall commands
hide: toc
---
# Syscall commands

## Object creation commands

These commands create new objects in the kernel, returning their file descriptors.

* [BPF_MAP_CREATE](BPF_MAP_CREATE.md)
* [BPF_PROG_LOAD](BPF_PROG_LOAD.md)
* [BPF_BTF_LOAD](BPF_BTF_LOAD.md)
* [BPF_LINK_CREATE](BPF_LINK_CREATE.md)
* BPF_ITER_CREATE
* BPF_RAW_TRACEPOINT_OPEN

## Map commands

Commands related to BPF maps.

* [BPF_MAP_CREATE](BPF_MAP_CREATE.md)
* [BPF_MAP_LOOKUP_ELEM](BPF_MAP_LOOKUP_ELEM.md)
* [BPF_MAP_UPDATE_ELEM](BPF_MAP_UPDATE_ELEM.md)
* [BPF_MAP_DELETE_ELEM](BPF_MAP_DELETE_ELEM.md)
* [BPF_MAP_GET_NEXT_KEY](BPF_MAP_GET_NEXT_KEY.md)
* [BPF_MAP_LOOKUP_BATCH](BPF_MAP_LOOKUP_BATCH.md)
* BPF_MAP_LOOKUP_AND_DELETE_BATCH
* BPF_MAP_UPDATE_BATCH
* BPF_MAP_DELETE_BATCH
* BPF_MAP_LOOKUP_AND_DELETE_ELEM
* BPF_MAP_FREEZE

## Pin commands

Commands related to the pinning of BPF objects.

* BPF_OBJ_PIN
* BPF_OBJ_GET

## Program commands

Commands related to BPF programs.

* [BPF_PROG_LOAD](BPF_PROG_LOAD.md)
* BPF_PROG_ATTACH
* BPF_PROG_DETACH
* [BPF_PROG_TEST_RUN](BPF_PROG_TEST_RUN.md)
* [BPF_PROG_RUN](BPF_PROG_TEST_RUN.md)
* BPF_PROG_BIND_MAP

## Object discovery commands

Commands used to find existing objects or iterate over them.

* [BPF_PROG_GET_NEXT_ID](BPF_PROG_GET_NEXT_ID.md)
* [BPF_MAP_GET_NEXT_ID](BPF_MAP_GET_NEXT_ID.md)
* [BPF_PROG_GET_FD_BY_ID](BPF_PROG_GET_FD_BY_ID.md)
* [BPF_MAP_GET_FD_BY_ID](BPF_MAP_GET_FD_BY_ID.md)
* [BPF_OBJ_GET_INFO_BY_FD](BPF_OBJ_GET_INFO_BY_FD.md)
* BPF_PROG_QUERY
* [BPF_BTF_GET_FD_BY_ID](BPF_BTF_GET_FD_BY_ID.md)
* BPF_TASK_FD_QUERY
* [BPF_BTF_GET_NEXT_ID](BPF_BTF_GET_NEXT_ID.md)
* [BPF_LINK_GET_FD_BY_ID](BPF_LINK_GET_FD_BY_ID.md)
* [BPF_LINK_GET_NEXT_ID](BPF_LINK_GET_NEXT_ID.md)

## Link commands 

Commands related to links.

* [BPF_LINK_CREATE](BPF_LINK_CREATE.md)
* BPF_LINK_UPDATE
* BPF_LINK_DETACH

## Statistics commands

* [BPF_ENABLE_STATS](BPF_ENABLE_STATS.md)
