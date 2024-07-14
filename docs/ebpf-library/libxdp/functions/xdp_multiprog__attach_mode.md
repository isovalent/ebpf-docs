---
title: "Libxdp Function 'xdp_multiprog__attach_mode'"
description: "This page documents the 'xdp_multiprog__attach_mode' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `xdp_multiprog__attach_mode`

## Definition

It give the `xdp_attach_mode` of the `xdp_multiprog` passed in parameter.

### Returns

Return the `enum xdp_attach_mode` <!-- Maybe list them -->

In case of not found, It will return `XDP_MODE_UNSPEC`.

## Usage

```c
enum xdp_attach_mode xdp_multiprog__attach_mode(const struct xdp_multiprog *mp);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
