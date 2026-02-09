# LIBXDP

libxdp is a light eBPF library who add 2 features for [XDP programs](../../linux/program-type/BPF_PROG_TYPE_XDP.md).

- Load multiple programs on single network device using a "dispatcher program" thanks to [`freplace`](../../linux/program-type/BPF_PROG_TYPE_EXT.md)
- Configuring [`AF_XDP`](../../linux/concepts/af_xdp.md) and functions to read and write on these sockets

You can check more information on the [libxdp readme](https://github.com/xdp-project/xdp-tools/blob/master/lib/libxdp/README.org).

## Manage multiple XDP programs

### Load and unload

Libxdp can help you to load, attach, unload and manage your XDP program.

- [`xdp_program__from_bpf_obj`](./functions/xdp_program__from_bpf_obj.md)
- [`xdp_program__find_file`](./functions/xdp_program__find_file.md)
- [`xdp_program__open_file`](./functions/xdp_program__open_file.md)
- [`xdp_program__from_fd`](./functions/xdp_program__from_fd.md)
- [`xdp_program__from_id`](./functions/xdp_program__from_id.md)
- [`xdp_program__from_pin`](./functions/xdp_program__from_pin.md)

!!! note
    `xdp_program__find_file`, will search the `bpf_object` to the path set by `LIBXDP_OBJECT_PATH`. By default it will be `/usr/lib/bpf`.

### Metadata

XDP program contains metadata to manage programs in the dispatcher.

To modify metadata you can use :

```c
#include <bpf/bpf_helpers.h>
#include <xdp/xdp_helpers.h>

struct {
	__uint(priority, 10);  // priority
	__uint(XDP_PASS, 1);   // chain call action
	__uint(XDP_DROP, 1);   // chain call action
} XDP_RUN_CONFIG(my_xdp_func);
```

Or by using xdp functions:  
This won't modify the BTF file, but these metadata will be stored with the attachment of the program.

!!! warning
    This work **only** before the attachment of the program to the dispatcher.
		
- [`xdp_program__run_prio`](./functions/xdp_program__run_prio.md)
- [`xdp_program__set_run_prio`](./functions/xdp_program__set_run_prio.md)
- [`xdp_program__chain_call_enabled`](./functions/xdp_program__chain_call_enabled.md)
- [`xdp_program__set_chain_call_enabled`](./functions/xdp_program__set_chain_call_enabled.md)
- [`xdp_program__print_chain_call_actions`](./functions/xdp_program__print_chain_call_actions.md)

#### Priority

The priority of a program is an integer used to determine the order program execution on the interface. Programs are ordered in increasing priority from low to high.
For passing packets to next program in the priority, the program should return a one of [chain call actions].

!!! note
    The default priority value is 50 if not specified.

[chain call actions]: #chain-call-action

#### Chain call action

Chain call actions are the return codes a program uses to indicate that a packet should continue in the dispatcher program. 
If a program returns one of these actions, subsequent programs in the call chain will execute. 
If a different action is returned the processing stop. 

!!! note
    By default, this is set to XDP_PASS

### The dispatcher program

To support multiple non-offloaded programs on the same network interface, libxdp uses a dispatcher program, a small wrapper that sequentially calls each component program. 
The dispatcher expects return codes and proceeds to the next program based on the [chain call actions] of the previous program.

- [`xdp_multiprog__get_from_ifindex`](./functions/xdp_multiprog__get_from_ifindex.md)
- [`xdp_multiprog__next_prog`](./functions/xdp_multiprog__next_prog.md)
- [`xdp_multiprog__close`](./functions/xdp_multiprog__close.md)
- [`xdp_multiprog__detach`](./functions/xdp_multiprog__detach.md)
- [`xdp_multiprog__attach_mode`](./functions/xdp_multiprog__attach_mode.md)
- [`xdp_multiprog__main_prog`](./functions/xdp_multiprog__main_prog.md)
- [`xdp_multiprog__hw_prog`](./functions/xdp_multiprog__hw_prog.md)
- [`xdp_multiprog__is_legacy`](./functions/xdp_multiprog__is_legacy.md)


### XDP dispatcher pinning
The kernel will automatically detach component programs from the dispatcher once their last reference disappears. To prevent this, libxdp pins the component program references in bpffs (BPF file system) before attaching the dispatcher to the network interface. The generated path names for pinning are:

- `/sys/fs/bpf/xdp/dispatch-IFINDEX-DID` : Dispatcher program for IFINDEX with BPF program ID DID.
- `/sys/fs/bpf/xdp/dispatch-IFINDEX-DID/prog0-prog` : Component program 0, program reference.
- `/sys/fs/bpf/xdp/dispatch-IFINDEX-DID/prog0-link` : Component program 0, bpf_link reference.
- `/sys/fs/bpf/xdp/dispatch-IFINDEX-DID/prog1-prog` : Component program 1, program reference.
- `/sys/fs/bpf/xdp/dispatch-IFINDEX-DID/prog1-link` : Component program 1, bpf_link reference.

The dispatcher can up max to ten programs
If set, the `LIBXDP_BPFFS` environment variable will override the default location of bpffs, though the xdp subdirectory is always used. If no bpffs is mounted, libxdp will check the LIBXDP_BPFFS_AUTOMOUNT environment variable. 
If this is set to 1, libxdp will attempt to automount a bpffs. 
If not set, libxdp will revert to loading a single program without a dispatcher, as if the kernel did not support the features required for multiprog attachment.

## AF_XDP sockets

You can find an explanation of the [AF_XDP concept here](../../linux/concepts/af_xdp.md)

AF_XDP sockets provide a high-performance mechanism for redirecting network packets to user space. The libxdp library implements helper functions for configuring these sockets and managing packet I/O.

!!! note
    Previously, this functionality was part of libbpf, but it has been moved to libxdp.  
    As of **libbpf 1.0**, AF_XDP socket support has been fully transitioned to libxdp.

### Control path

Libxdp provides utility functions to help create and manage umems and AF_XDP sockets. You need to create a *umem* area and then link an *AF_XDP* socket to it.

#### Umem Area

it's a memory region designated to store packets. It holds the packets that are received and those that need to be sent. `xsk_umem__get_data` is used to access the packet data in the umem area. But in unaligned mode, you need to use the three last function

- [`xsk_umem__create`](./functions/xsk_umem__create.md)
- [`xsk_umem__create_with_fd`](./functions/xsk_umem__create_with_fd.md)
- [`xsk_umem__delete`](./functions/xsk_umem__delete.md)
- [`xsk_umem__fd`](./functions/xsk_umem__fd.md)
- [`xsk_umem__get_data`](./functions/xsk_umem__get_data.md)
- [`xsk_umem__extract_addr`](./functions/xsk_umem__extract_addr.md)
- [`xsk_umem__extract_offset`](./functions/xsk_umem__extract_offset.md)
- [`xsk_umem__add_offset_to_addr`](./functions/xsk_umem__add_offset_to_addr.md)

#### Sockets

Once the umem is created, you can create AF_XDP sockets that are linked to this umem.
These sockets can either:
* Exclusively own the umem: This is done using the function [`xsk_socket__create()`](./functions/xsk_umem__create.md).
* Share the umem with other sockets: This is done using the function [`xsk_socket__create_shared()`](./functions/xsk_socket__create_shared.md).

- [`xsk_socket__create`](./functions/xsk_socket__create.md)
- [`xsk_socket__create_shared`](./functions/xsk_socket__create_shared.md)
- [`xsk_socket__delete`](./functions/xsk_socket__delete.md)
- [`xsk_socket__fd`](./functions/xsk_socket__fd.md)
- [`xsk_setup_xdp_prog`](./functions/xsk_setup_xdp_prog.md)
- [`xsk_socket__update_xskmap`](./functions/xsk_socket__update_xskmap.md)

The [XSK map](../../linux/map-type/BPF_MAP_TYPE_XSKMAP.md) is used by the XDP program to manage the mapping between the network interface and the user-space sockets.

### Data Path

There are four FIFO rings, categorized into two main types :

- Producer Rings: These include the fill and TX rings, using `xsk_ring_prod*` functions : 
    * _Fill ring_ : Provide buffers to the kernel.
    * _TX ring_ : Send packets.

- Consumer Rings: These include the Rx and completion rings, using `xsk_ring_cons*` functions :  
    * _Rx ring_ : Receive packets from the kernel.
    * _Completion ring_ : Acknowledge completion of transmitted packets.

The producer rings manage the supply of buffers for sending and receiving packets, while the consumer rings manage the recovery and reuse of these buffers after the packets have been processed. You can read more information about the concept in the [detailed section AF_XDP](../../linux/concepts/af_xdp.md#receiving-and-sending-packets) in this wiki.

!!! note
    All the data path functions are static inline functions.

!!! note
    To advance to the next entry, simply do `idx++`.

#### Producer rings

For producer rings, you start with **reserving** one or more slots in a producer ring and then when they have been filled out, you **submit** them so that the kernel will act on them. After this you **release** them back to the kernel so it can use them for new packets.  
Others functions **writes** entries in the fill and TX rings.

- [`xsk_ring_prod__reserve`](./functions/xsk_ring_prod__reserve.md)
- [`xsk_ring_prod__submit`](./functions/xsk_ring_prod__submit.md)
- [`xsk_ring_prod__fill_addr`](./functions/xsk_ring_prod__fill_addr.md)
- [`xsk_ring_prod__tx_desc`](./functions/xsk_ring_prod__tx_desc.md)
- [`xsk_ring_prod__needs_wakeup`](./functions/xsk_ring_prod__needs_wakeup.md)

#### Consumer rings

For a consumer ring, you **peek** if there are any new packets in the ring and if so you can read them from the ring. After this you **release** them back to the kernel so it can use them for new packets
There is also a **cancel** operation for consumer rings if the application does not want to consume all packets received with the peek operation.  
Others functions **reads** entries from the completion and Rx rings.

- [`xsk_ring_cons__peek`](./functions/xsk_ring_cons__peek.md)
- [`xsk_ring_cons__cancel`](./functions/xsk_ring_cons__cancel.md)
- [`xsk_ring_cons__release`](./functions/xsk_ring_cons__release.md)
- [`xsk_ring_cons__comp_addr`](./functions/xsk_ring_cons__comp_addr.md)
- [`xsk_ring_cons__rx_desc`](./functions/xsk_ring_cons__rx_desc.md)

## Examples

You can find example in the [bpf examples repository](https://github.com/xdp-project/bpf-examples)

## Tools

You can use [`xdp-tools`](https://github.com/xdp-project/xdp-tools) (and [bpftool](https://github.com/libbpf/bpftool)) to help you to build XDP programs.

This doc is written base on the doc provided on [libxdp readme][1]

[1]: https://github.com/xdp-project/xdp-tools/tree/master/lib/libxdp "Libxdp readme"

