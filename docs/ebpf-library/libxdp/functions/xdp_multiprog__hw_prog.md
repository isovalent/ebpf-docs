---
title: "Libxdp Function 'xdp_multiprog__hw_prog'"
description: "This page documents the 'xdp_multiprog__hw_prog' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xdp_multiprog__hw_prog`

## Definition

It return a reference to the program loaded in the interface (useful is it's not the dispatcher).

### Returns

`struct xdp_program` on success, or a negative error in case of failure.

## Usage

```c
struct xdp_program *xdp_multiprog__hw_prog(const struct xdp_multiprog *mp);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
