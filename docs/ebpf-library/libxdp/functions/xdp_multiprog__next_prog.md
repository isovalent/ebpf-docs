---
title: "Libxdp Function 'xdp_multiprog__next_prog'"
description: "This page documents the 'xdp_multiprog__next_prog' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xdp_multiprog__next_prog`

## Definition

Allow to get the next loaded in the dispatcher.

### Returns

`struct xdp_program` pointer of the next program if success, or a negative error in case of failure
    
## Usage

```c
struct xdp_program *xdp_multiprog__next_prog(const struct xdp_program *prog, const struct xdp_multiprog *mp);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
