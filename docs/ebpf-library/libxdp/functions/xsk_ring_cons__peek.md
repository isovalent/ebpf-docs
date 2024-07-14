---
title: "Libxdp Function 'xsk_ring_cons__peek'"
description: "This page documents the 'xsk_ring_cons__peek' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xsk_ring_cons__peek`

## Definition

Check for new packets in the ring.

### Returns

`__u32` returns the number of packets that are available in the consumer ring (`idx`).

!!! note
    It can be less than or equal to the number of packets requested to peek.

## Usage

```c
__u32 xsk_ring_cons__peek(struct xsk_ring_cons *cons, __u32 nb, __u32 *idx);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
