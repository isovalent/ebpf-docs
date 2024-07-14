---
title: "Libxdp Function 'xsk_umem__fd'"
description: "This page documents the 'xsk_umem__fd' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xsk_umem__fd`

## Definition

This function return the file descriptor of the umem passed in parameter.

### Returns

file descriptor of the umem on success, or a negative error in case of failure:

**-EINVAL** if arguments are invalid  
    
## Usage

```c
int xsk_umem__fd(const struct xsk_umem *umem);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
