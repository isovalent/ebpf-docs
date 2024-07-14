---
title: "Libxdp Function 'xsk_socket__update_xskmap'"
description: "This page documents the 'xsk_socket__update_xskmap' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xsk_socket__update_xskmap`

## Definition

This function updates an XSK map with a new AF_XDP socket.

### Returns

0 on success, or a negative error in case of failure:

**-EINVAL** if arguments are invalid

**-ENOENT** if entry not found
    
## Usage

```c
int xsk_socket__update_xskmap(struct xsk_socket *xsk, int xsks_map_fd);
```

### Example

You can find an exemple at [AF_XDP-example](https://github.com/xdp-project/bpf-examples/tree/master/AF_XDP-example)

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
