---
title: "Syscall command 'BPF_MAP_LOOKUP_AND_DELETE_BATCH' - eBPF Docs"
description: "This page documents the 'BPF_MAP_LOOKUP_AND_DELETE_BATCH' eBPF syscall command, including its defintion, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_MAP_LOOKUP_AND_DELETE_BATCH` command

<!-- [FEATURE_TAG](BPF_MAP_LOOKUP_AND_DELETE_BATCH) -->
[:octicons-tag-24: v5.6](https://github.com/torvalds/linux/commit/057996380a42bb64ccc04383cfa9c0ace4ea11f0)
<!-- [/FEATURE_TAG] -->

!!! warning
    You would expect, due to the similar naming to the `BPF_MAP_LOOKUP_AND_DELETE_ELEM` command that this would apply to stack and queue maps, however, it does not. This command can only be used on htab (hash map, per-cpu-hash-map, lru-hash-map, per-cpu-lru-hashmap) maps as of v6.2.

<!-- TODO -->
