---
title: "Libxdp Function 'xdp_program__find_file'"
description: "This page documents the 'xdp_program__find_file' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xdp_program__find_file`

## Definition

Return the `xdp_program` associated to a filename and section names. 
The filename is **without** the PATH, it will automaticcaly look at `LIBXDP_OBJECT_PATH`. 

!!! note
    By default, `LIBXDP_OBJECT_PATH` is set to `/usr/lib/bpf` (alternatively `/usr/lib64/bpf` or `/usr/local/lib/bpf`, depending on the PREFIX libxdp was compiled with)

!!! note
    When libxdp is compiled with the DEBUG flag, it will additionally look for BPF object files in the current directory, before checking the system-wide directory. This should not be used in production, as shadowing of system BPF programs can be a security issue.

!!! note
    If you want to use a specific path, you can use [`xdp_program__open_file`](./xdp_program__open_file.md)

### Returns

`struct xdp_program` on success, or a negative error in case of failure
    
## Usage

```c
struct xdp_program *xdp_program__find_file(const char *filename, const char *section_name, struct bpf_object_open_opts *opts);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
