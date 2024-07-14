---
title: "Libxdp Function 'xdp_program__set_chain_call_enabled'"
description: "This page documents the 'xdp_program__set_chain_call_enabled' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xdp_program__set_chain_call_enabled`

## Definition

This allows adding actions to the chain call. If the program returns this `xdp_action`, the packet will be sent to the next program (in order of the [priority](../libxdp.md#priority)); otherwise, it will not.

!!! note
    By default, this is set to XDP_PASS

!!! warning
    This only work before the load of the program to the dispatcher.

### Returns

0 on success, or a negative error in case of failure:

**-EINVAL** if arguments are invalid

## Usage

```c
int xdp_program__set_chain_call_enabled(struct xdp_program *prog, unsigned int action, bool enabled);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
