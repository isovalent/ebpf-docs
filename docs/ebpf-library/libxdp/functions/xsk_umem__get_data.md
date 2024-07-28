---
title: "Libxdp Function 'xsk_umem__get_data'"
description: "This page documents the 'xsk_umem__get_data' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xsk_umem__get_data`

## Definition

Allow to get a pointer to the packet data with the **Rx** descriptor, in **aligned mode**.

!!! note
    In unaligned mode, you need to use these functions [`xsk_umem__extract_addr`](./xsk_umem__extract_addr.md) [`xsk_umem__extract_offset`](./xsk_umem__extract_offset.md) [`xsk_umem__add_offset_to_addr`](./xsk_umem__add_offset_to_addr.md)

### Returns

No return
    
## Usage

```c
void *xsk_umem__get_data(void *umem_area, __u64 addr);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
