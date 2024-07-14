---
title: "Libxdp Function 'xdp_program__set_run_prio'"
description: "This page documents the 'xdp_program__set_run_prio' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xdp_program__set_run_prio`

## Definition

This function allow to set a [priority](../libxdp.md#priority) the program when loaded to the dispatcher.
The priority of a program is an integer used to determine the order in which programs are executed on the interface. 

!!! note
    The higher the value, the later the program will run.
    The lower the value, the earlier the program will run.

!!! note
    The default priority value is 50 if not specified.

!!! warning
    This only work before the load of the program to the dispatcher.

### Returns

0 on success, or a negative error in case of failure:

**-EINVAL** if arguments are invalid

## Usage

```c
int xdp_program__set_run_prio(struct xdp_program *xdp_prog, unsigned int run_prio);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
