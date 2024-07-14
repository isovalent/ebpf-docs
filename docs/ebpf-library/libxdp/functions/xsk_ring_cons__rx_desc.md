---
title: "Libxdp Function 'xsk_ring_cons__rx_desc'"
description: "This page documents the 'xsk_ring_cons__rx_desc' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xsk_ring_cons__rx_desc`

## Definition

This function is used to retrieve the receive descriptor at a specific index in the **Rx** ring. 

### Returns

`struct xdp_desc` represents the receive descriptor at the given index in the Rx ring on success, or a negative error in case of failure:

## Usage

```c
const struct xdp_desc *xsk_ring_cons__rx_desc(const struct xsk_ring_cons *rx, __u32 idx);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
