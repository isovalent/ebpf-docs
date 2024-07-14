---
title: "Libxdp Function 'xdp_program__print_chain_call_actions'"
description: "This page documents the 'xdp_program__print_chain_call_actions' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xdp_program__print_chain_call_actions`

## Definition

Return to the buffer the chain call actions associated to a `xdp_program`.

### Returns

0 on success, or a negative error in case of failure:

**-EINVAL** if arguments are invalid
    
## Usage

```c
int xdp_program__print_chain_call_actions(const struct xdp_program *prog, char *buf, size_t buf_len);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
