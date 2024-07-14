---
title: "Libxdp Function 'xdp_program__run_prio'"
description: "This page documents the 'xdp_program__run_prio' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xdp_program__run_prio`

## Definition

It allow to retrieve the value of the priorty of the program.

!!! note
    The higher the value, the later the program will run.
    The lower the value, the earlier the program will run.

### Returns

`unsigned int` on success, or a negative error in case of failure:

**-EINVAL** if arguments are invalid
    
## Usage

```c
unsigned int xdp_program__run_prio(const struct xdp_program *xdp_prog);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
