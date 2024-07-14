---
title: "Libxdp Function 'xdp_multiprog__is_legacy'"
description: "This page documents the 'xdp_multiprog__is_legacy' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xdp_multiprog__is_legacy`

## Definition

Indicate whether the dispatcher program has been identified.

### Returns

- **true** if the dispatcher is legacy, unrecognized, or another program that is not the dispatcher is loaded.
- **false** if the dispatcher is identified.
    
## Usage

```c
bool xdp_multiprog__is_legacy(const struct xdp_multiprog *mp);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
