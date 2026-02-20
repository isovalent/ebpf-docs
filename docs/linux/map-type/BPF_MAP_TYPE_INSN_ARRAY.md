---
title: "Map Type 'BPF_MAP_TYPE_INSN_ARRAY'"
description: "This page documents the 'BPF_MAP_TYPE_INSN_ARRAY' eBPF map type, including its definition, usage, program types that can use it, and examples."
---
# Map type `BPF_MAP_TYPE_INSN_ARRAY`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_INSN_ARRAY) -->
[:octicons-tag-24: v6.19](https://github.com/torvalds/linux/commit/b4ce5923e780d6896d4aaf19de5a27652b8bf1ea)
<!-- [/FEATURE_TAG] -->

This map is used to track changes in instruction offsets from the provided byte code during loading to the translated bytecode and the JIT-ed bytecode. At present this is only useful to be able to programmatically see which JIT-ed instructions correspond to which instructions in the original bytecode. This would for example allow a user to profile a program with `perf`, and then to take the report and map the profiling info back to the eBPF byte code, and perhaps even to source code with BTF line info and or DWARF debug info.

This map is slated to be used in a number of future kernel features. This doc should be updated once these have landed in the kernel.

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

