---
title: "Libxdp Function 'xsk_socket__create_shared'"
description: "This page documents the 'xsk_socket__create_shared' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xsk_socket__create_shared`

## Definition

Creates an AF_XDP socket and share the ownership of the umem between multiple sockets.

### Returns

0 on success, or a negative error in case of failure:

**-EINVAL** if arguments are invalid  

**-EFAULT**  if memory address is invalid

**-ENOMEM**  if no data space available

**-ENOPROTOOPT**  if option is not supported by the protocol
    
## Usage

```c
int xsk_socket__create_shared(struct xsk_socket **xsk_ptr,
			      const char *ifname,
			      __u32 queue_id, struct xsk_umem *umem,
			      struct xsk_ring_cons *rx,
			      struct xsk_ring_prod *tx,
			      struct xsk_ring_prod *fill,
			      struct xsk_ring_cons *comp,
			      const struct xsk_socket_config *config);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
