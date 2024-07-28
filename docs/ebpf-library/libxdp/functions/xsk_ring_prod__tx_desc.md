---
title: "Libxdp Function 'xsk_ring_prod__tx_desc'"
description: "This page documents the 'xsk_ring_prod__tx_desc' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xsk_ring_prod__tx_desc`

## Definition

This function allow to access a specific transmit descriptor in the **TX** ring.

### Returns

`struct xdp_desc`, informations about the packet to be transmitted.

## Usage

```c
struct xdp_desc *xsk_ring_prod__tx_desc(struct xsk_ring_prod *tx, __u32 idx);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
