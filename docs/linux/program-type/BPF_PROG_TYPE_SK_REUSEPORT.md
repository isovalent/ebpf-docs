---
title: "Program Type 'BPF_PROG_TYPE_SK_REUSEPORT' - eBPF Docs"
description: "This page documents the 'BPF_PROG_TYPE_SK_REUSEPORT' eBPF program type, including its defintion, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_SK_REUSEPORT`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_SK_REUSEPORT) -->
[:octicons-tag-24: v4.19](https://github.com/torvalds/linux/commit/2dbb9b9e6df67d444fbe425c7f6014858d337adf)
<!-- [/FEATURE_TAG] -->

Socket reuse port programs can be attached to a `SO_REUSEPORT` socket group to replace the default socket selection mechanism.

## Usage

In [:octicons-tag-24: v3.9](https://github.com/torvalds/linux/commit/c617f398edd4db2b8567a28e899a88f8f574798d) the [`SO_REUSEPORT`](https://lwn.net/Articles/542629/) socket option was added which allows multiple sockets to listen to the same port on the same host. The original purpose of the feature being that this allows for high-efficient distribution of traffic across threads which would normally have to be done in userspace causing unnecessary delay.

By default, incoming connections and datagrams are distributed to the server sockets using a hash based on the 4-tuple of the connection—that is, the peer IP address and port plus the local IP address and port.

With the introduction of [BPF_PROG_TYPE_SK_REUSEPORT](../program-type/BPF_PROG_TYPE_SK_REUSEPORT.md) program, [BPF_MAP_TYPE_REUSEPORT_SOCKARRAY](../map-type/BPF_MAP_TYPE_REUSEPORT_SOCKARRAY.md) map, and the [bpf_sk_select_reuseport](../helper-function/bpf_sk_select_reuseport.md) helper function we can replace the default distribution behavior with a BPF program.

A key feature is that the sockets do not have to belong to the same process. This means that you can steer traffic between two processes to do A/B testing or software updates without dropping connections. For the latter scenario, the typical use case is to use a map-in-map with a [BPF_MAP_TYPE_REUSEPORT_SOCKARRAY](../map-type/BPF_MAP_TYPE_REUSEPORT_SOCKARRAY.md) as inner map, allowing userspace to switch out all sockets at once. In that scenario, any existing TCP connections would still be handled by the old sockets/process but new connections are routed to the new process. 

## Context

The context of this program type is `#!c struct sk_reuseport_md`. All fields of this context type are read-only and may not be modified by the program directly.

??? abstract "c structure"
    ```c
    struct sk_reuseport_md {
        /*
        * Start of directly accessible data. It begins from
        * the tcp/udp header.
        */
        __bpf_md_ptr(void *, data);
        /* End of directly accessible data */
        __bpf_md_ptr(void *, data_end);
        /*
        * Total length of packet (starting from the tcp/udp header).
        * Note that the directly accessible bytes (data_end - data)
        * could be less than this "len".  Those bytes could be
        * indirectly read by a helper "bpf_skb_load_bytes()".
        */
        __u32 len;
        /*
        * Eth protocol in the mac header (network byte order). e.g.
        * ETH_P_IP(0x0800) and ETH_P_IPV6(0x86DD)
        */
        __u32 eth_protocol;
        __u32 ip_protocol;	/* IP protocol. e.g. IPPROTO_TCP, IPPROTO_UDP */
        __u32 bind_inany;	/* Is sock bound to an INANY address? */
        __u32 hash;		/* A hash of the packet 4 tuples */
        /* When reuse->migrating_sk is NULL, it is selecting a sk for the
        * new incoming connection request (e.g. selecting a listen sk for
        * the received SYN in the TCP case).  reuse->sk is one of the sk
        * in the reuseport group. The bpf prog can use reuse->sk to learn
        * the local listening ip/port without looking into the skb.
        *
        * When reuse->migrating_sk is not NULL, reuse->sk is closed and
        * reuse->migrating_sk is the socket that needs to be migrated
        * to another listening socket.  migrating_sk could be a fullsock
        * sk that is fully established or a reqsk that is in-the-middle
        * of 3-way handshake.
        */
        __bpf_md_ptr(struct bpf_sock *, sk);
        __bpf_md_ptr(struct bpf_sock *, migrating_sk);
    };
    ```

### `data`

This field contain a pointer to the start of directly accessible data. It begins from the tcp/udp header.

!!! note
    This program type only has read access, it may not modify the packet data.

### `data_end`

This field contain a pointer to the end of directly accessible data.

### `len`

This field contains the total length of packet (starting from the tcp/udp header). 

!!! note 
    The directly accessible bytes (data_end - data) could be less than this "len". Those bytes could be indirectly read by a helper `bpf_skb_load_bytes`.

### `eth_protocol`

This field contains the ethernet protocol in the mac header (network byte order). e.g. `ETH_P_IP` (`0x0800`) and `ETH_P_IPV6` (`0x86DD`)

### `ip_protocol`

This field contain the IP protocol. e.g. `IPPROTO_TCP`, `IPPROTO_UDP`.

### `bind_inany`

This field is `true` if the socket group is bound to an INANY address.

### `hash`

This field is a hash of the packet 4 tuples.

### `sk` and `migrating_sk`

[:octicons-tag-24: v5.14](https://github.com/torvalds/linux/commit/e061047684af63f2d4f1338ec73140f6e29eb59f)
and
[:octicons-tag-24: v5.14](https://github.com/torvalds/linux/commit/d5e4ddaeb6ab2c3c7fbb7b247a6d34bb0b18d87e)

These fields are used together to handle socket migration. If both are `NULL` we are doing the initial selection.

When `migrating_sk` is `NULL`, it is selecting a sk for the new incoming connection request (e.g. selecting a listen sk for the received SYN in the TCP case).  `sk` is one of the sk in the reuseport group. The bpf prog can use reuse->sk to learn the local listening ip/port without looking into the skb.

When `migrating_sk` is not NULL, `sk` is closed and `migrating_sk` is the socket that needs to be migrated to another listening socket.  migrating_sk could be a fullsock sk that is fully established or a reqsk that is in-the-middle of 3-way handshake.

## Attachment

This program type can be attached to a reuse port socket group by using the [`setsockopt`](https://man7.org/linux/man-pages/man2/setsockopt.2.html) syscall on one of the sockets in the group with the `SOL_SOCKET` socket level and  `SO_ATTACH_BPF` [socket option](https://man7.org/linux/man-pages/man7/socket.7.html).

This program should be loaded with the `BPF_SK_REUSEPORT_SELECT` [`expected_attach_type`](../syscall/BPF_PROG_LOAD.md#expected_attach_type) to use it only for the selection logic or `BPF_SK_REUSEPORT_SELECT_OR_MIGRATE` if the program should also handle [socket migration](#socket-migration) logic.

## Socket migration

Before [:octicons-tag-24: v5.14](https://github.com/torvalds/linux/commit/d5e4ddaeb6ab2c3c7fbb7b247a6d34bb0b18d87e), the reuse port feature had a defect in its logic. When a SYN packet is received, the connection is tied to a listening socket. Accordingly, when the listener is closed, in-flight requests during the three-way handshake and child sockets in the accept queue are dropped even if other listeners could accept such connections.

This situation can happen when various server management tools restart server (such as nginx) processes. For instance, when we change nginx configurations and restart it, it spins up new workers that respect the new configuration and closes all listeners on the old workers, resulting in in-flight ACK of 3WHS is responded by RST.

To fix this defect, the concept of socket migration was added, which will repeat the socket selection logic to pick a new socket. When not using eBPF, the same hash logic is used, but only if the `net.ipv4.tcp_migrate_req` sysctl setting has been enabled. When using eBPF with this program type, loading the program with the `BPF_SK_REUSEPORT_SELECT_OR_MIGRATE` attachment type indicates that this program also overwrites the migration logic. No need to set the sysctl option in this case. This does mean that the the program can be called for initial selection as well as for migration. The `sk` and `sk_migration` context fields indicate for which purpose the program is invoked.

When invoked for migration, the following actions can be taken:

 * return `SK_PASS` after selecting a socket with [bpf_sk_select_reuseport](../helper-function/bpf_sk_select_reuseport.md), select it as a new listener.
 * return `SK_PASS` without calling [bpf_sk_select_reuseport](../helper-function/bpf_sk_select_reuseport.md), falls back to the random selection.
 * return `SK_DROP`, cancel the migration.

!!! note
    The kernel select a listening socket in three places, but it does not have `struct skb` at closing a listener or retransmitting a SYN+ACK. On the other hand, some helper functions do not expect skb is NULL (e.g. skb_header_pointer() in BPF_FUNC_skb_load_bytes(), skb_tail_pointer() in BPF_FUNC_skb_load_bytes_relative()). So the kernel allocates an empty skb temporarily before running the eBPF program.

## Helper functions

Not all helper functions are available in all program types. These are the helper calls available for socket reuse port programs:

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [bpf_sk_select_reuseport](../helper-function/bpf_sk_select_reuseport.md)
    * [bpf_skb_load_bytes](../helper-function/bpf_skb_load_bytes.md)
    * [bpf_skb_load_bytes_relative](../helper-function/bpf_skb_load_bytes_relative.md)
    * [bpf_get_socket_cookie](../helper-function/bpf_get_socket_cookie.md)
    * [bpf_ktime_get_coarse_ns](../helper-function/bpf_ktime_get_coarse_ns.md)
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
