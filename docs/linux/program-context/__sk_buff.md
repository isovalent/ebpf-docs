# Program context `__sk_buff`

The socket buffer context is provided to program types that deal with network packets when there already is a socket buffer created/allocated. The `struct __sk_buff` is a "mirror" of the `struct sk_buff` program type which is actually used by the kernel. 

Accesses to the `struct __sk_buff` pointer are seamlessly transformed into accesses into the real socket buffer. This indirection exists to provide a stable ABI for programs since the `struct sk_buff` may change between kernel versions and to provide a layer of checks. Not all program types are allowed to read and/or write to certain fields for a number of reasons.
<!-- TODO: What reasons are those? -->

## Direct packet access

<!-- TODO: A fairly significant change, should be documented -->

## Fields

### `len`
[:octicons-tag-24: v4.1](https://github.com/torvalds/linux/commit/9bac3d6d548e5cc925570b263f35b70a00a00ffd)

This field holds the **total** length of the packet. It is important to know that this doesn't indicate the amount of data that is available via [direct packet access](#direct-packet-access). In some cases the packet is larger than a single memory page, in which case the packet data lives in non-linear in which case the `len` might be larger than `data_end`-`data` and specialized [helpers](../helper-function/index.md) are needed to access the rest of the memory.
<!-- TODO link the actual helper functions in question -->

### `pkt_type`
[:octicons-tag-24: v4.1](https://github.com/torvalds/linux/commit/9bac3d6d548e5cc925570b263f35b70a00a00ffd)

This field indicates the type of the packet which informs "who" the packet is for. Possible values of this field are
the `PACKET_*` values defined in `include/uapi/linux/if_packet.h`.

* `PACKET_HOST` - indicates the packet is addresses to the MAC address of this host
* `PACKET_BROADCAST` - indicates the packet is addressed to a broadcast address.
* `PACKET_MULTICAST` - indicates the packet is addressed to a multicast address.
* `PACKET_OTHERHOST` - indicates the packet to addressed to some other host that it has been caught by a device driver in promiscuous mode
* `PACKET_OUTGOING` - indicates the packet originating from the local host that is looped back to a packet socket

!!! note
    This is not an exhaustive list of possible values.

### `mark`
[:octicons-tag-24: v4.1](https://github.com/torvalds/linux/commit/9bac3d6d548e5cc925570b263f35b70a00a00ffd)

This field is a general purpose 32 bit tag used in the network subsystem to carry metadata with global implications across network sub-subsystem. As an example, a driver could mark on incoming packet to be used by the ingress tc classifier-action sub-subsystem, netfilter, ipsec all to execute provisioned policies.[^1]

### `queue_mapping`
[:octicons-tag-24: v4.1](https://github.com/torvalds/linux/commit/9bac3d6d548e5cc925570b263f35b70a00a00ffd)

This field indicates via which TX queue on the NIC this packet should be sent. Typically this field is set by TC but can be overwritten by certain eBPF programs to implement custom balancing logic. [^2]

### `protocol`
[:octicons-tag-24: v4.1](https://github.com/torvalds/linux/commit/c24973957975403521ca76a776c2dfd12fbe9add)

This field indicates the Layer 3 protocol of the packet and is one of the `ETH_P_*` values defined in `include/uapi/linux/if_ether.h`.

### `vlan_present`
[:octicons-tag-24: v4.1](https://github.com/torvalds/linux/commit/c24973957975403521ca76a776c2dfd12fbe9add)

This field is a boolean `0` or `1` and indicates if the packet has a VLAN header.

### `vlan_tci`
[:octicons-tag-24: v4.1](https://github.com/torvalds/linux/commit/c24973957975403521ca76a776c2dfd12fbe9add)

This field contains the VLAN TCI (Tag Control Information), if the packet included a VLAN header.

### `vlan_proto`
[:octicons-tag-24: v4.1](https://github.com/torvalds/linux/commit/27cd5452476978283decb19e429e81fc6c71e74b)

This field contains the protocol ID of the used VLAN protocol which will be one of the `ETH_P_*` values defined in `include/uapi/linux/if_ether.h`.

### `priority`
[:octicons-tag-24: v4.1](https://github.com/torvalds/linux/commit/bcad57182425426dd4aa14deb27f97acb329f3cd)

This field indicates the queueing priority of the packet. Packets with higher priority will be send out first. Only values between `0` and `63` are effective, values of `64` and above will be converted to `63`. This field only takes effect if the `skbprio` queueing discipline has been configured in TC. [^3]

This only effects egress traffic since ingress traffic is never queued.

### `ingress_ifindex`
[:octicons-tag-24: v4.2](https://github.com/torvalds/linux/commit/37e82c2f974b72c9ab49c787ef7b5bb1aec12768)

This field contains the interface index of the network devices this packet arrived on. It may be `0` if a process on the host originated the packet.

### `ifindex`
[:octicons-tag-24: v4.2](https://github.com/torvalds/linux/commit/37e82c2f974b72c9ab49c787ef7b5bb1aec12768)

This field contains the interface index of the network device the packet is currently "on", so if a packet has been redirected to another device and a eBPF program is invoked on it again, this field should be updated to the new device.

On egress this will be the device picked for sending the packet.

### `tc_index`
[:octicons-tag-24: v4.2](https://github.com/torvalds/linux/commit/d691f9e8d4405c334aa10d556e73c8bf44cb0e01)

This field is used to carry Type of Service (TOS) information. This field is populated by the `dsmark` qdisc and can subsequently be used with [`tcindex`](https://man7.org/linux/man-pages/man8/tc-tcindex.8.html) filters to classify packets based on their TOS value.

The `dsmark` uses the differentiated services (DS) fields in IPv4 (aka DSCP) and IPv6 (aka traffic class) headers.

[`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md) programs can also modify this value to implement a custom TOS value extraction from packets.

### `cb`
[:octicons-tag-24: v4.2](https://github.com/torvalds/linux/commit/d691f9e8d4405c334aa10d556e73c8bf44cb0e01)

This field is an array of 5 u32 values with no pre-defined meaning. Network subsystems and eBPF programs can read from and write to this field to share information associated with the socket buffer across programs and subsystem boundaries.

### `hash`
[:octicons-tag-24: v4.3](https://github.com/torvalds/linux/commit/ba7591d8b28bd16a2eface5d009ab0b60c7629a4)

This field contains the calculated from the flow information of the packet. The fields used to calculate the hash can differ depending on the protocol. This hash is optionally calculated by network interface devices that support it. [^4]

### `tc_classid`
[:octicons-tag-24: v4.4](https://github.com/torvalds/linux/commit/045efa82ff563cd4e656ca1c2e354fa5bf6bbda4)

This field can be used by [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md) in direct action mode to set the class id. This value is only useful if the program returns a `TC_ACT_OK` and the qdisc has classes.

### `data`
[:octicons-tag-24: v4.4](https://github.com/torvalds/linux/commit/969bf05eb3cedd5a8d4b7c346a85c2ede87a6d6d)

This field contains the pointer to the start address of the linear packet data. This will be the first byte of the layer 3 header the type of which is indicated by `protocol`.

<!-- TODO more details around bounds checks and rewriting -->

### `data_end`
[:octicons-tag-24: v4.4](https://github.com/torvalds/linux/commit/969bf05eb3cedd5a8d4b7c346a85c2ede87a6d6d)

This field contains the pointer to the last address of the packet data linear packet data. This pointer is used in combination with `data` to indicate accessible data.

### `napi_id`
[:octicons-tag-24: v4.12](https://github.com/torvalds/linux/commit/b1d9fc41aab11f9520b2e0d57ae872e2ec5d6f32)

This field contains the id of the NAPI struct this socket buffer came from.

### `family`
[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/8a31db5615667956c513d205cfb06885c3ec6d0b)

This field contains the address family of the socket associated this this socket buffer. Its value is one of `AF_*` values defined in `include/linux/socket.h`.

### `remote_ip4`
[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/8a31db5615667956c513d205cfb06885c3ec6d0b)

<!-- TODO -->

### `local_ip4`
[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/8a31db5615667956c513d205cfb06885c3ec6d0b)

<!-- TODO -->

### `remote_ip6`
[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/8a31db5615667956c513d205cfb06885c3ec6d0b)

<!-- TODO -->

### `local_ip6`
[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/8a31db5615667956c513d205cfb06885c3ec6d0b)

<!-- TODO -->

### `remote_port`
[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/8a31db5615667956c513d205cfb06885c3ec6d0b)

<!-- TODO -->

### `local_port`
[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/8a31db5615667956c513d205cfb06885c3ec6d0b)

<!-- TODO -->

### `data_meta`
[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/de8f3a83b0a0fddb2cf56e7a718127e9619ea3da)

<!-- TODO -->

### `flow_keys`
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/d58e468b1112dcd1d5193c0a89ff9f98b5a3e8b9)

This field is a pointer to a `struct bpf_flow_keys` which like the name implies hold the keys that identify the network flow of the socket buffer. More details can be found in the [dedicated section](#flow-keys).

This field is only accessible from within [`BPF_PROG_TYPE_FLOW_DISSECTOR`](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md) programs.

### `tstamp`
[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/f11216b24219ab26d8d159fbfa12dff886b16e32)

This field indicates the time when this packet should be transmitted in nanoseconds since boot. [`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md) programs can set this time to some time in the future to add delay to packets for the purposes of bandwidth limiting or simulating latency. Setting this value only works on egress if the `fq` (Fair Queue) qdisc is used.

!!! note
    The `fq` qdisc has a "drop horizon" if packets are set to transmit to far into the future they will be dropped to avoid queueing to many packets.

### `wire_len`
[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/e3da08d057002f9d0831949d51666c3e15dc6b29)

<!-- TODO -->

### `gso_segs`
[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/d9ff286a0f59fa7843549e49bd240393dd7d8b87)

<!-- TODO -->

### `sk`
[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/46f8bc92758c6259bcf945e9216098661c1587cd)

This field is a pointer to a `struct bpf_sock` which holds information about the socket associated with this socket buffer. More details can be found in the [dedicated section](#socket)

This field is always read-only.

<!-- TODO -->

### `gso_size`
[:octicons-tag-24: v5.7](https://github.com/torvalds/linux/commit/cf62089b0edd7e74a1f474844b4d9f7b5697fb5c)

<!-- TODO -->

### `tstamp_type`
[:octicons-tag-24: v5.18](https://github.com/torvalds/linux/commit/9bb984f28d5bcb917d35d930fcfb89f90f9449fd)

<!-- TODO -->

### `hwtstamp`
[:octicons-tag-24: v5.16](https://github.com/torvalds/linux/commit/f64c4acea51fbe2c08c0b0f48b7f5d1657d7a5e4)

<!-- TODO -->

## Flow keys

This section describes the fields of the `struct bpf_flow_keys` type.

### `nhoff`
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/d58e468b1112dcd1d5193c0a89ff9f98b5a3e8b9)

<!-- TODO -->

### `thoff`
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/d58e468b1112dcd1d5193c0a89ff9f98b5a3e8b9)

<!-- TODO -->

### `addr_proto`
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/d58e468b1112dcd1d5193c0a89ff9f98b5a3e8b9)

<!-- TODO -->

### `is_frag`
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/d58e468b1112dcd1d5193c0a89ff9f98b5a3e8b9)

<!-- TODO -->

### `is_first_frag`
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/d58e468b1112dcd1d5193c0a89ff9f98b5a3e8b9)

<!-- TODO -->

### `is_encap`
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/d58e468b1112dcd1d5193c0a89ff9f98b5a3e8b9)

<!-- TODO -->

### `ip_proto`
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/d58e468b1112dcd1d5193c0a89ff9f98b5a3e8b9)

<!-- TODO -->

### `n_proto`
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/d58e468b1112dcd1d5193c0a89ff9f98b5a3e8b9)

<!-- TODO -->

### `sport`
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/d58e468b1112dcd1d5193c0a89ff9f98b5a3e8b9)

<!-- TODO -->

### `dport`
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/d58e468b1112dcd1d5193c0a89ff9f98b5a3e8b9)

<!-- TODO -->

### `ipv4_src`
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/d58e468b1112dcd1d5193c0a89ff9f98b5a3e8b9)

<!-- TODO -->

### `ipv4_dst`
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/d58e468b1112dcd1d5193c0a89ff9f98b5a3e8b9)

<!-- TODO -->

### `ipv6_src`
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/d58e468b1112dcd1d5193c0a89ff9f98b5a3e8b9)

<!-- TODO -->

### `ipv6_dst`
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/d58e468b1112dcd1d5193c0a89ff9f98b5a3e8b9)

<!-- TODO -->

### `flags`
[:octicons-tag-24: v5.4](https://github.com/torvalds/linux/commit/086f95682114fd2d1790bd3226e76cbae9a2d192)

<!-- TODO -->

### `flow_label`
[:octicons-tag-24: v5.4](https://github.com/torvalds/linux/commit/71c99e32b926159ea628352751f66383d7d04d17)

<!-- TODO -->


## Socket

This section describes the fields of the `struct bpf_sock` type which is a mirror of the kernels `struct sock` type.

### `bound_dev_if`
[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/61023658760032e97869b07d54be9681d2529e77)

<!-- TODO -->

### `family`
[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/aa4c1037a30f4e88f444e83d42c2befbe0d5caf5)

<!-- TODO -->

### `type`
[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/aa4c1037a30f4e88f444e83d42c2befbe0d5caf5)

<!-- TODO -->

### `protocol`
[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/aa4c1037a30f4e88f444e83d42c2befbe0d5caf5)

<!-- TODO -->

### `mark`
[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/482dca939fb7ee35ba20b944b4c2476133dbf0df)

<!-- TODO -->

### `priority`
[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/482dca939fb7ee35ba20b944b4c2476133dbf0df)

<!-- TODO -->

### `src_ip4`
[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/aa65d6960a98fc15a96ce361b26e9fd55c9bccc5)

<!-- TODO -->

### `src_ip6`
[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/aa65d6960a98fc15a96ce361b26e9fd55c9bccc5)

<!-- TODO -->

### `src_port`
[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/aa65d6960a98fc15a96ce361b26e9fd55c9bccc5)

<!-- TODO -->

### `dst_port`
[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/aa65d6960a98fc15a96ce361b26e9fd55c9bccc5)

<!-- TODO -->

### `dst_ip4`
[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/aa65d6960a98fc15a96ce361b26e9fd55c9bccc5)

<!-- TODO -->

### `dst_ip6`
[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/aa65d6960a98fc15a96ce361b26e9fd55c9bccc5)

<!-- TODO -->

### `state`
[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/aa65d6960a98fc15a96ce361b26e9fd55c9bccc5)

<!-- TODO -->

### `rx_queue_mapping`
[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/c3c16f2ea6d20159903cf93afbb1155f3d8348d5)

<!-- TODO -->

[^1]: [https://www.spinics.net/lists/netdev/msg235744.html](https://www.spinics.net/lists/netdev/msg235744.html)
[^2]: [https://www.kernel.org/doc/Documentation/networking/multiqueue.txt](https://www.kernel.org/doc/Documentation/networking/multiqueue.txt)
[^3]: [https://man7.org/linux/man-pages/man8/tc-skbprio.8.html](https://man7.org/linux/man-pages/man8/tc-skbprio.8.html)
[^4]: [https://www.kernel.org/doc/Documentation/networking/scaling.txt](https://www.kernel.org/doc/Documentation/networking/scaling.txt)
