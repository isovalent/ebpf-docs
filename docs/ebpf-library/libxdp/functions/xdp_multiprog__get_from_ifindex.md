---
title: "Libxdp Function 'xdp_multiprog__get_from_ifindex'"
description: "This page documents the 'xdp_multiprog__get_from_ifindex' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xdp_multiprog__get_from_ifindex`

## Definition

This allow to get the `xdp_multiprog`, the dispatcher program, from the index of an interface.

### Returns

0 on success, or a negative error in case of failure:

**-EBUSY** if the dispatcher can't be reach

**-ENOENT** if _ifindex_ dosen't exist
    
## Usage

```c
struct xdp_multiprog *xdp_multiprog__get_from_ifindex(int ifindex);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
