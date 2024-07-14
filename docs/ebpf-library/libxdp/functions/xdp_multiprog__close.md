---
title: "Libxdp Function 'xdp_multiprog__close'"
description: "This page documents the 'xdp_multiprog__close' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xdp_multiprog__close`

## Definition

This function will close all programs inside a dispatcher program, linked to an interface.

### Returns

0 on success, or a negative error in case of failure:

**-EINVAL** if arguments are invalid
    
## Usage

```c
void xdp_multiprog__close(struct xdp_multiprog *mp);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
