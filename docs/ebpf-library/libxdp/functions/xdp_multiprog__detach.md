---
title: "Libxdp Function 'xdp_multiprog__detach'"
description: "This page documents the 'xdp_multiprog__detach' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xdp_multiprog__detach`

## Definition

This function will detach programs inside a dispatcher, but It will not close them.

### Returns

0 on success, or a negative error in case of failure:

**-EINVAL** if argument is invalid
    
## Usage

```c
int xdp_multiprog__detach(struct xdp_multiprog *mp, int ifindex);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
