---
title: "Libxdp Function 'xsk_umem__add_offset_to_addr'"
description: "This page documents the 'xsk_umem__add_offset_to_addr' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xsk_umem__add_offset_to_addr`

## Definition

This function add the offset to the address in **unaligned mode**.

!!! note
    You need to use this function with these functions [`xsk_umem__extract_addr`](./xsk_umem__extract_addr.md) [`xsk_umem__extract_offset`](./xsk_umem__extract_offset.md)

!!! note
    In aligned mode, you need to use [`xsk_umem_get_data`](./xsk_umem__get_data.md).

### Returns

`__u64` of the new address
    
## Usage

```c
__u64 xsk_umem__add_offset_to_addr(__u64 addr);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
