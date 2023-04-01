# Program type `BPF_PROG_TYPE_XDP`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_XDP) -->
[:octicons-tag-24: v4.8](https://github.com/torvalds/linux/commit/6a773a15a1e8874e5eccd2f29190c31085912c95)
<!-- [/FEATURE_TAG] -->

XDP (Express Data Path) programs can attach to network devices and are called for every incoming (ingress) packet received by that network device. XDP programs can take quite a large number of actions, most prominent of which are manipulation of the packet, dropping the packet, redirecting it and letting it pass to the network stack.

Notable use cases for XDP programs are for DDoS protection, Load Balancing, and high-throughput packet filtering. If loaded with native driver support, XDP programs will be called just after receiving the packet but before allocating memory for a socket buffer. This callsite makes XDP programs extremely performant, especially in use cases where traffic is forwarded or dropped a lot in comparison to other eBPF program types or techniques which run after the relatively expensive socket buffer allocation process has taken place, only to discard it.

## Usage

XDP programs are typically put into an [ELF](../../elf.md) section prefixed with `xdp`. The XDP program is called by the kernel with a [xdp_md](../program-context/xdp_md.md) context. The return value indicates what action the kernel should take with the packet, the following values are permitted:

* `XDP_ABORTED` - Signals that a unrecoverable error has taken place. Returning this action will cause the kernel to trigger the `xdp_exception` tracepoint and print a line to the trace log. This allows for debugging of such occurrences. It is also expensive, so should not be used without consideration in production.
* `XDP_DROP` - Discards the packet. It should be noted that since we drop the packet very early, it will be invisible to tools like `tcpdump`. Consider recording drops using a custom feedback mechanism to maintain visibility.
* `XDP_PASS` - Pass the packet to the network stack. The packet can be manipulated before hand
* `XDP_TX` - Send the packet back out the same network port it arrived on. The packet can be manipulated before hand.
* `XDP_REDIRECT` - Redirect the packet to one of a number of locations. The packet can be manipulated before hand.

`XDP_REDIRECT` should not be returned by itself, always in combination with a helper function call. A number of helper functions can be used to redirect the current packet. These annotate hidden values in the context to inform the kernel what actual redirection action to take after the program exists.

Packets can be redirected in the following ways:

* The packet can be redirected to egress on a different interface than where it entered (like `XDP_TX` but for a different interface). This can be done using the [`bpf_redirect`](../helper-function/bpf_redirect.md) helper (not recommended) or the [`bpf_redirect_map`](../helper-function/bpf_redirect_map.md) helper in combination with a [`BPF_MAP_TYPE_DEVMAP`](../map-type/BPF_MAP_TYPE_DEVMAP.md) or [`BPF_MAP_TYPE_DEVMAP_HASH`](../map-type/BPF_MAP_TYPE_DEVMAP_HASH.md) map.
* The packet can be redirected to another CPU for further processing using the [`bpf_redirect_map`](../helper-function/bpf_redirect_map.md) helper in combination with a [`BPF_MAP_TYPE_CPUMAP`](../map-type/BPF_MAP_TYPE_CPUMAP.md) map.
* The packet can be redirected to userspace, bypassing the kernel network stack using the [`bpf_redirect_map`](../helper-function/bpf_redirect_map.md) helper in combination with a [`BPF_MAP_TYPE_XSKMAP`](../map-type/BPF_MAP_TYPE_XSKMAP.md) map

## Context

XDP programs are called with the `struct xdp_md` context. This is a very simple context representing a single packet.

### `data`

[:octicons-tag-24: v4.8](https://github.com/torvalds/linux/commit/6a773a15a1e8874e5eccd2f29190c31085912c95)

This field contains a pointer to the start of packet data. The XDP program can read from this region between `data` and `data_end`, as long as it always performs bounds checks.

### `data_end`

[:octicons-tag-24: v4.8](https://github.com/torvalds/linux/commit/6a773a15a1e8874e5eccd2f29190c31085912c95)

This field contains a pointer to the end of the packet data. The verifier will enforce that any XDP program checks that offsets from `data` are less then `data_end` before the program attempts to read from it.

### `data_meta`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/de8f3a83b0a0fddb2cf56e7a718127e9619ea3da)

This field contains a pointer to the start of a metadata region in the packet memory. By default, no metadata room is available, so the value of `data_meta` and `data` will be the same. The XDP program can request metadata with the [`bpf_xdp_adjust_meta`](../helper-function/bpf_xdp_adjust_meta.md) helper, on success `data_meta` is updated so it is not less then `data`. The room between `data_meta` and `data` is freely useable by the XDP program.

If the packet with metadata is passed to the kernel, that metadata will be available in the [`__sk_buff`](../program-context/__sk_buff.md) via its [`data_meta`](../program-context/__sk_buff.md#data_meta) and `data` fields.

This means that XDP programs can communicate information to for example `BPF_PROG_TYPE_SCHED_CLS` programs which can then manipulate the socket buffer to change `__sk_buff->mark` or `__sk_buff->priority` on behalf of an XDP program.

### `ingress_ifindex`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/02dd3291b2f095bbc88e1d2628fd5bf2e92de69b)

This field contains the network interface index the packet arrived on.

### `rx_queue_index`

[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/02dd3291b2f095bbc88e1d2628fd5bf2e92de69b)

This field contains the queue index within the NIC on which the packet was received.

!!! note
    While this field is normally read-only, offloaded XDP programs are allowed to write to it to perform custom RSS (Receive-Side Scaling) in the network device [:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/0d8300325660f81787892a1c58dc1f9428a67143)

### `egress_ifindex`

[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/64b59025c15b244c0954cf52b24fbabfcf5ed8f6)

This field is read-only and contains the network interface index the packet has been redirected out of. This field is only ever set after an initial XDP program redirected a packet to another device with a [`BPF_MAP_TYPE_DEVMAP`](../map-type/BPF_MAP_TYPE_DEVMAP.md) and the value of the devmap contained a file descriptor of a secondary XDP program. This secondary program will be invoked with a context that has `egress_ifindex`, `rx_queue_index`, and `ingress_ifindex` set so it can modify fields in the packet to match the redirection.

### XDP fragments

[:octicons-tag-24: v5.18](https://github.com/torvalds/linux/commit/c2f2cdbeffda7b153c19e0f3d73149c41026c0db)

An increasingly common performance optimization technique is to use larger packets and to bulk process them (Jumbo packets, GRO, BIG-TCP). It might therefor happen that packets get larger than a single memory page or that we want to glue multiple already allocated packets together. This breaks the existing assumption XDP programs have of all the packet data living in a linear area between `data` and `data_end`. 

In order to offer support and not break existing programs, the concept of "XDP fragment aware" programs was introduced. XDP program authors writing such programs can compare the length between the `data` and `data_end` pointer and the output of [`bpf_xdp_get_buff_len`](../helper-function/bpf_xdp_get_buff_len.md). If the XDP program needs to work with data beyond the linear portion it should use the [`bpf_xdp_load_bytes`](../helper-function/bpf_xdp_load_bytes.md) and [`bpf_xdp_store_bytes`](../helper-function/bpf_xdp_store_bytes.md) helpers.

To indicate that a program is "XDP Fragment aware" the program should be loaded with the [`BPF_F_XDP_HAS_FRAGS`](../syscall/BPF_PROG_LOAD.md#bpf_f_xdp_has_frags) flag. Program authors can indicate that they wish libraries like libbpf to load programs with this flag by placing their program in a `xdp.frags/` ELF section instead of a `xdp/` section.

!!! note
    If a program is both "XDP Fragment aware" and should be attached to a CPUMAP or DEVMAP the two ELF naming conventions are combined: `xdp.frags/cpumap/` or `xdp.frags/devmap`.

## Attachment

There are two ways of attaching XDP programs to network devices, the legacy way of doing is is via a [netlink](https://man7.org/linux/man-pages/man7/netlink.7.html) socket the details of which are complex. Examples of libraries that implement netlink XDP attaching are [vishvananda/netlink](https://github.com/vishvananda/netlink/blob/afa2eb2a66aac1f8f370287f236ba93d4c078dd6/link_linux.go#L934) and [libbpf](https://github.com/libbpf/libbpf/blob/ea284299025bf85b85b4923191de6463cd43ccd6/src/netlink.c#L321).

The modern and recommended way is to use BPF links. Doing so is as easy as calling [`BPF_LINK_CREATE`](../syscall/BPF_LINK_CREATE.md) with the `target_ifindex` set to the network interface target, `attach_type` set to `BPF_LINK_TYPE_XDP` and the same `flags` as would be used for the netlink approach.

There are some subtile differences. The netlink method will give the network interface a reference to the program, which means that after attaching, the program will stay attached until it is detached by a program, even if the original loader exists. This is in contrast to kprobes for example which will stop as soon as the loader exists (assuming we are not pinning the program). With links however, this referencing doesn't occur, the creation of the link returns a file descriptor which is used to manage the lifecycle, if the link fd is closed or the loader exists without pinning it, the program will be detached from the network interface.

### Flags

#### `XDP_FLAGS_UPDATE_IF_NOEXIST`

If set, the kernel will only attach the XDP program if the network interface doesn't have a XDP program attached already.

!!! note
    This flag is only used with the netlink attach method, the link attach method handles this behavior more generically.

#### `XDP_FLAGS_SKB_MODE`

If set, the kernel will attach the program in SKB (Socket buffer) mode. This mode is also known as "Generic mode". This always works regardless of driver support. It works by calling the XDP program after a socket buffer has already been allocated further up the stack that an XDP program would normally be called. This negates the speed advantage of XDP programs. This mode also lacks full feature support since some actions cannot be taken this high up the network stack anymore. 

It is recommended to use `BPF_PROG_TYPE_SCHED_CLS` prog types instead if driver support isn't available since it offers more capabilities with roughtly the same performance.

This flag is mutually exclusive with `XDP_FLAGS_DRV_MODE` and `XDP_FLAGS_HW_MODE`

#### `XDP_FLAGS_DRV_MODE`

If set, the kernel will attach the program in driver mode. This does require support from the network driver, but most predominant network card vendors have support in the latest kernel.

This flag is mutually exclusive with `XDP_FLAGS_SKB_MODE` and `XDP_FLAGS_HW_MODE`

#### `XDP_FLAGS_HW_MODE`

If set, the kernel will attach the program in hardware offload mode. This requires both driver and hardware support for XDP offloading. Currently only select Netronome devices [support offloading](https://www.netronome.com/media/documents/eBPF_HW_OFFLOAD_HNiMne8_2_.pdf). However, it should be noted that only a subset of normal features are supported. 

#### `XDP_FLAGS_REPLACE`

If set, the kernel will atomically replace the existing program for this new program. You will also have to pass the file descriptor of the old program via the netlink request.

!!! note
    This flag is only used with the netlink attach method, the link attach method handles this behavior more generically.

### Device map program

[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/281920b7e0b31e0a7706433ff58e7d52ac97c327)

XDP programs can be attached to map values of a [`BPF_MAP_TYPE_DEVMAP`](../map-type/BPF_MAP_TYPE_DEVMAP.md) map. Once attached this program will run after the first program concluded but before the packet is sent of to the new network device. These programs are called with additional context, see [`egress_ifindex`](#egress_ifindex).

Only XDP programs that have been loaded with the `BPF_XDP_DEVMAP` value in [`expected_attach_type`](../syscall/BPF_PROG_LOAD.md#expected_attach_type) are allowed to be attached in this way.

Program authors can indicate to loaders like libbpf that a given program should be loaded with this expected attach type by placing the program in a `xdp/devmap/` ELF section.

### CPU map program

[:octicons-tag-24: v5.9](https://github.com/torvalds/linux/commit/9216477449f33cdbc9c9a99d49f500b7fbb81702).

XDP programs can be attached to map values of a [`BPF_MAP_TYPE_CPUMAP`](../map-type/BPF_MAP_TYPE_CPUMAP.md) map. Once attached this program will run on the new logical CPU. The idea being that you would spend minimal time in the first XDP program and only schedule it and perform the more CPU intensive tasks in this second program.

Only XDP programs that have been loaded with the `BPF_XDP_CPUMAP` value in [`expected_attach_type`](../syscall/BPF_PROG_LOAD.md#expected_attach_type) are allowed to be attached in this way.

Program authors can indicate to loaders like libbpf that a given program should be loaded with this expected attach type by placing the program in a `xdp/cpumap/` ELF section.

## Helper functions

Not all helper functions are available in all program types. These are the helper calls available for XDP programs:

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [bpf_perf_event_output](../helper-function/bpf_perf_event_output.md)
    * [bpf_get_smp_processor_id](../helper-function/bpf_get_smp_processor_id.md)
    * [bpf_csum_diff](../helper-function/bpf_csum_diff.md)
    * [bpf_xdp_adjust_head](../helper-function/bpf_xdp_adjust_head.md)
    * [bpf_xdp_adjust_meta](../helper-function/bpf_xdp_adjust_meta.md)
    * [bpf_redirect](../helper-function/bpf_redirect.md)
    * [bpf_redirect_map](../helper-function/bpf_redirect_map.md)
    * [bpf_xdp_adjust_tail](../helper-function/bpf_xdp_adjust_tail.md)
    * [bpf_xdp_get_buff_len](../helper-function/bpf_xdp_get_buff_len.md)
    * [bpf_xdp_load_bytes](../helper-function/bpf_xdp_load_bytes.md)
    * [bpf_xdp_store_bytes](../helper-function/bpf_xdp_store_bytes.md)
    * [bpf_fib_lookup](../helper-function/bpf_fib_lookup.md)
    * [bpf_check_mtu](../helper-function/bpf_check_mtu.md)
    * [bpf_sk_lookup_udp](../helper-function/bpf_sk_lookup_udp.md)
    * [bpf_sk_lookup_tcp](../helper-function/bpf_sk_lookup_tcp.md)
    * [bpf_sk_release](../helper-function/bpf_sk_release.md)
    * [bpf_skc_lookup_tcp](../helper-function/bpf_skc_lookup_tcp.md)
    * [bpf_tcp_check_syncookie](../helper-function/bpf_tcp_check_syncookie.md)
    * [bpf_tcp_gen_syncookie](../helper-function/bpf_tcp_gen_syncookie.md)
    * [bpf_map_lookup_elem](../helper-function/bpf_map_lookup_elem.md)
    * [bpf_map_update_elem](../helper-function/bpf_map_update_elem.md)
    * [bpf_map_delete_elem](../helper-function/bpf_map_delete_elem.md)
    * [bpf_map_push_elem](../helper-function/bpf_map_push_elem.md)
    * [bpf_map_pop_elem](../helper-function/bpf_map_pop_elem.md)
    * [bpf_map_peek_elem](../helper-function/bpf_map_peek_elem.md)
    * [bpf_map_lookup_percpu_elem](../helper-function/bpf_map_lookup_percpu_elem.md)
    * [bpf_get_prandom_u32](../helper-function/bpf_get_prandom_u32.md)
    * [bpf_get_smp_processor_id](../helper-function/bpf_get_smp_processor_id.md)
    * [bpf_get_numa_node_id](../helper-function/bpf_get_numa_node_id.md)
    * [bpf_tail_call](../helper-function/bpf_tail_call.md)
    * [bpf_ktime_get_ns](../helper-function/bpf_ktime_get_ns.md)
    * [bpf_ktime_get_boot_ns](../helper-function/bpf_ktime_get_boot_ns.md)
    * [bpf_ringbuf_output](../helper-function/bpf_ringbuf_output.md)
    * [bpf_ringbuf_reserve](../helper-function/bpf_ringbuf_reserve.md)
    * [bpf_ringbuf_submit](../helper-function/bpf_ringbuf_submit.md)
    * [bpf_ringbuf_discard](../helper-function/bpf_ringbuf_discard.md)
    * [bpf_ringbuf_query](../helper-function/bpf_ringbuf_query.md)
    * [bpf_for_each_map_elem](../helper-function/bpf_for_each_map_elem.md)
    * [bpf_loop](../helper-function/bpf_loop.md)
    * [bpf_strncmp](../helper-function/bpf_strncmp.md)
    * [bpf_spin_lock](../helper-function/bpf_spin_lock.md)
    * [bpf_spin_unlock](../helper-function/bpf_spin_unlock.md)
    * [bpf_jiffies64](../helper-function/bpf_jiffies64.md)
    * [bpf_per_cpu_ptr](../helper-function/bpf_per_cpu_ptr.md)
    * [bpf_this_cpu_ptr](../helper-function/bpf_this_cpu_ptr.md)
    * [bpf_timer_init](../helper-function/bpf_timer_init.md)
    * [bpf_timer_set_callback](../helper-function/bpf_timer_set_callback.md)
    * [bpf_timer_start](../helper-function/bpf_timer_start.md)
    * [bpf_timer_cancel](../helper-function/bpf_timer_cancel.md)
    * [bpf_trace_printk](../helper-function/bpf_trace_printk.md)
    * [bpf_get_current_task](../helper-function/bpf_get_current_task.md)
    * [bpf_get_current_task_btf](../helper-function/bpf_get_current_task_btf.md)
    * [bpf_probe_read_user](../helper-function/bpf_probe_read_user.md)
    * [bpf_probe_read_kernel](../helper-function/bpf_probe_read_kernel.md)
    * [bpf_probe_read_user_str](../helper-function/bpf_probe_read_user_str.md)
    * [bpf_probe_read_kernel_str](../helper-function/bpf_probe_read_kernel_str.md)
    * [bpf_snprintf_btf](../helper-function/bpf_snprintf_btf.md)
    * [bpf_snprintf](../helper-function/bpf_snprintf.md)
    * [bpf_task_pt_regs](../helper-function/bpf_task_pt_regs.md)
    * [bpf_trace_vprintk](../helper-function/bpf_trace_vprintk.md)
<!-- [/PROG_HELPER_FUNC_REF] -->
