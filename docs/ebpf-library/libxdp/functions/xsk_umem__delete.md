---
title: "Libxdp Function 'xsk_umem__delete'"
description: "This page documents the 'xsk_umem__delete' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xsk_umem__delete`

## Definition

Delete an umem area.

### Returns

0 on success, or a negative error in case of failure:

**-EINVAL** if arguments are invalid  

**-EBUSY** if the umem is busy
    
## Usage

```c
int xsk_umem__delete(struct xsk_umem *umem);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
