---
title: "Libxdp Function 'xsk_umem__create_with_fd'"
description: "This page documents the 'xsk_umem__create_with_fd' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xsk_umem__create_with_fd`

## Definition

Create an umem area using a file descriptor.

### Returns

0 on success, or a negative error in case of failure:

**-EINVAL** if arguments are invalid  

**-EFAULT** if the memory adress is invalid

**-ENOMEM** if no space memory left
    
## Usage

```c
int xsk_umem__create_with_fd(struct xsk_umem **umem,
			     int fd, void *umem_area, __u64 size,
			     struct xsk_ring_prod *fill,
			     struct xsk_ring_cons *comp,
			     const struct xsk_umem_config *config);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
