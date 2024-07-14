---
title: "Libxdp Function 'xsk_ring_prod__needs_wakeup'"
description: "This page documents the 'xsk_ring_prod__needs_wakeup' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xsk_ring_prod__needs_wakeup`

## Definition

This function function checks if the kernel needs to be woken up to process the **producer** ring.

!!! note
    It encouraged to enable the flag `XDP_USE_NEED_WAKEUP` in `xdp_bind_flags` provided in `xsk_socket_create_*shared*`

!!! note
    If the function returns non-zero value, you should call `recvmsg()` when receiving, `sendto()` when sending, or `poll() `for both operations, although this last is slower than the other two.

### Returns

return non-zero value if the kernel need to be wake up, else it will return 0.

## Usage

```c
int xsk_ring_prod__needs_wakeup(const struct xsk_ring_prod *r);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
