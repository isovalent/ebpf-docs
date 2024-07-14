---
title: "Libxdp Function 'xsk_setup_xdp_prog'"
description: "This page documents the 'xsk_setup_xdp_prog' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xsk_setup_xdp_prog`

## Definition

This function sets up an XDP program on a specified network interface.  

### Returns

0 on success, or a negative error in case of failure:

**-ENOENT** if file not found
    
## Usage

```c
int xsk_setup_xdp_prog(int ifindex, int *xsks_map_fd);
```

### Example

You can find an exemple at [AF_XDP-example](https://github.com/xdp-project/bpf-examples/tree/master/AF_XDP-example)

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
