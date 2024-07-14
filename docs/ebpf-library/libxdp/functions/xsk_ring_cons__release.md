---
title: "Libxdp Function 'xsk_ring_cons__release'"
description: "This page documents the 'xsk_ring_cons__release' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xsk_ring_cons__release`

## Definition

This function releases a specified number of packets that have been processed from the **consumer** ring back to the kernel. Indicates to the kernel that these packets have been consumed and the buffers can be reused for new incoming packets.

### Returns

No return

## Usage

```c
void xsk_ring_cons__release(struct xsk_ring_cons *cons, __u32 nb);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
