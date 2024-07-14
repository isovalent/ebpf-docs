---
title: "Libxdp Function 'xsk_ring_prod__fill_addr'"
description: "This page documents the 'xsk_ring_prod__fill_addr' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xsk_ring_prod__fill_addr`

## Definition

Use this function to get a pointer to a slot in the **fill** ring to set the address of a packet buffer.

### Returns

`__u64` address of the packet.

## Usage

```c
__u64 *xsk_ring_prod__fill_addr(struct xsk_ring_prod *fill, __u32 idx);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
