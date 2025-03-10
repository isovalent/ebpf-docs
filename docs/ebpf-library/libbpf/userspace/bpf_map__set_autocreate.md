---
title: "Libbpf userspace function 'bpf_map__set_autocreate'"
description: "This page documents the 'bpf_map__set_autocreate' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_map__set_autocreate`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.8.0](https://github.com/libbpf/libbpf/releases/tag/v0.8.0)
<!-- [/LIBBPF_TAG] -->

Sets whether libbpf has to auto-create BPF map during BPF object load phase.

## Definition

`#!c int bpf_map__set_autocreate(struct bpf_map *map, bool autocreate);`

**Parameters**

- `map`: the BPF map instance
- `autocreate`: whether to create BPF map during BPF object load

**Return**

0 on success; `-EBUSY` if BPF object was already loaded

## Usage


[`bpf_map__set_autocreate()`](bpf_map__set_autocreate.md) allows to opt-out from libbpf auto-creating
BPF map. By default, libbpf will attempt to create every single BPF map
defined in BPF object file using BPF_MAP_CREATE command of bpf() syscall
and fill in map FD in BPF instructions.

This API allows to opt-out of this process for specific map instance. This
can be useful if host kernel doesn't support such BPF map type or used
combination of flags and user application wants to avoid creating such
a map in the first place. User is still responsible to make sure that their
BPF-side code that expects to use such missing BPF map is recognized by BPF
verifier as dead code, otherwise BPF verifier will reject such BPF program.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
