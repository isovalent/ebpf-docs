---
title: "Libxdp Function 'xdp_program__from_fd'"
description: "This page documents the 'xdp_program__from_fd' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xdp_program__from_fd`

## Definition

Return the `xdp_program` associated to a file descriptor.

### Returns

`struct xdp_program` on success, or a negative error in case of failure:

**-ENOENT** if file not found
    
## Usage

```c
struct xdp_program *xdp_program__from_fd(int fd);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
