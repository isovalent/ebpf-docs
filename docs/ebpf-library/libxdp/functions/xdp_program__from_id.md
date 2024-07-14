---
title: "Libxdp Function 'xdp_program__from_id'"
description: "This page documents the 'xdp_program__from_id' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xdp_program__from_id`

## Definition

Return the `xdp_program` associated to an id.

### Returns

`struct xdp_program` on success, or a negative error in case of failure:

**-ENOENT** If no such file or directory

**-EINVAL** If arguments are invalid
    
## Usage

```c
struct xdp_program *xdp_program__from_id(__u32 prog_id);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
