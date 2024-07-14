---
title: "Libxdp Function 'xsk_ring_prod__reserve'"
description: "This page documents the 'xsk_ring_prod__reserve' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xsk_ring_prod__reserve`

## Definition

Reserve one or more slots in a **producer** ring.

### Returns

`__u32` number of slots that were successfully reserved (`idx`) on success, or a 0 in case of failure.

!!! note
    It can be less than or equal to the number of packets requested to reserve.

## Usage

```c
__u32 xsk_ring_prod__reserve(struct xsk_ring_prod *prod, __u32 nb, __u32 *idx);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
