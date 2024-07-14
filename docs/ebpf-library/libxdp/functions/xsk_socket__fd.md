---
title: "Libxdp Function 'xsk_socket__fd'"
description: "This page documents the 'xsk_socket__fd' libxdp function, including its definition, usage, program types that can use it, and examples."
---
# Libxdp function `xsk_socket__fd`

## Definition

Allow you to get the file descriptor of the socket.

### Returns

0 on success, or a negative error in case of failure:

**-EINVAL** if arguments are invalid
    
## Usage

```c
int xsk_socket__fd(const struct xsk_socket *xsk);
```

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
