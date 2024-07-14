---
title: "Libxdp Function 'xdp_program__chain_call_enabled'"
description: "This page documents the 'xdp_program__chain_call_enabled' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xdp_program__chain_call_enabled`

## Definition

Return, true or false, if the `xdp_action` passed in parameter is set to be call chain action.

### Returns

`true` if the action is set to be call chain action, else false

## Usage

```c
bool xdp_program__chain_call_enabled(const struct xdp_program *xdp_prog, enum xdp_action action);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
