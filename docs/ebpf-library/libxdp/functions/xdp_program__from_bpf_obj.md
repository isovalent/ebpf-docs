---
title: "Libxdp Function 'xdp_program__from_bpf_obj'"
description: "This page documents the 'xdp_program__from_bpf_obj' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xdp_program__from_bpf_obj`

## Definition

Return the `xdp_program` associated to a bpf object with the section name.

### Returns

`struct xdp_program` on success, or a negative error in case of failure
    
## Usage

```c
struct xdp_program *xdp_program__from_bpf_obj(struct bpf_object *obj, const char *section_name);
```
### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
