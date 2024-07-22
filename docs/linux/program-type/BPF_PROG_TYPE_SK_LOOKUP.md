---
title: "Program Type 'BPF_PROG_TYPE_SK_LOOKUP'"
description: "This page documents the 'BPF_PROG_TYPE_SK_LOOKUP' eBPF program type, including its definition, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_SK_LOOKUP`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_SK_LOOKUP) -->
[:octicons-tag-24: v5.9](https://github.com/torvalds/linux/commit/e9ddbb7707ff5891616240026062b8c1e29864ca)
<!-- [/FEATURE_TAG] -->

The socket lookup program allows an eBPF program to pick which socket to send traffic to irrespective of how that target socket has been [bound](https://man7.org/linux/man-pages/man2/bind.2.html).

The primary use case for this program type is to allow a single program to handle traffic for network patterns which cannot be expressed with the normal [bind](https://man7.org/linux/man-pages/man2/bind.2.html) syscall. For example, a single socket can be bound to a whole `/24` network CIDR (bind only allows for single IPs, or you have to set it to `0.0.0.0` which is not desirable if another application should answer a different range of IPs). Or a single socket can listen to any port for a given IP.

## Usage

Socket lookup programs are typically put into an [ELF](../../concepts/elf.md) section prefixed with `sk_lookup`. Socket lookup programs  are invoked by the transport layer when looking up a listening socket for a new connection request for connection oriented protocols, or when looking up an unconnected socket for a packet for connection-less protocols.

The socket lookup program acts as a filter, if it returns `SK_DROP` (`0`) the connection or packet is dropped. If it returns `SK_PASS` (`1`) without setting a socket, the normal resolve behavior is used. However, the program can also chose to assign a specific socket with the [`bpf_sk_assign`](../helper-function/bpf_sk_assign.md) helper function.

## Context

Socket lookup programs are called with the `struct bpf_sk_lookup` context.

!!! abstract "c structure"
    ```c
    union {
        __bpf_md_ptr(struct bpf_sock *, sk); /* Selected socket */
        __u64 cookie; /* Non-zero if socket was selected in PROG_TEST_RUN */
    };

    __u32 family;		    /* Protocol family (AF_INET, AF_INET6) */
    __u32 protocol;		    /* IP protocol (IPPROTO_TCP, IPPROTO_UDP) */
    __u32 remote_ip4;	    /* Network byte order */
    __u32 remote_ip6[4];    /* Network byte order */
    __be16 remote_port;	    /* Network byte order */
    __u16 :16;		        /* Zero padding */
    __u32 local_ip4;	    /* Network byte order */
    __u32 local_ip6[4];	    /* Network byte order */
    __u32 local_port;	    /* Host byte order */
    __u32 ingress_ifindex;  /* The arriving interface. Determined by inet_iif. */
    ```

### `sk`

This field is a pointer to a selected socket, the field is read-only, but can be updated via the [`bpf_sk_assign`](../helper-function/bpf_sk_assign.md) helper function.

### `cookie`

This field is is set to the cookie of the assigned socket if the program assigns one during a [`PROG_TEST_RUN`](../syscall/BPF_PROG_TEST_RUN.md).

### `family`

The address family of the connection/packet for which the program is invoked. Can be [`AF_INET`](https://elixir.bootlin.com/linux/v6.2.8/source/include/linux/socket.h#L191) or [`AF_INET6`](https://elixir.bootlin.com/linux/v6.2.8/source/include/linux/socket.h#L199)

### `protocol`

The transport layer protocol of the connection/packet for which the program is invoked. Can be [`IPPROTO_TCP`](https://elixir.bootlin.com/linux/v6.2.8/source/include/uapi/linux/in.h#L38) or [`IPPROTO_UDP`](https://elixir.bootlin.com/linux/v6.2.8/source/include/uapi/linux/in.h#L44)

### `remote_ip4`

The remote IPv4 address of the connection/packet for which the program is invoked.

### `remote_ip6`

The remote IPv6 address of the connection/packet for which the program is invoked.

### `remote_port`

The remote port of the connection/packet for which the program is invoked.

### `local_ip4`

The local IPv4 address of the connection/packet for which the program is invoked.

### `local_ip6`

The local IPv6 address of the connection/packet for which the program is invoked.

### `local_port`

The local port of the connection/packet for which the program is invoked.

### `ingress_ifindex`

The network interface index of the network interface on which the packet arrived.

## Attachment

This program type must always be loaded with the [`expected_attach_type`](../syscall/BPF_PROG_LOAD.md#expected_attach_type) of `BPF_SK_LOOKUP`.

Socket lookup programs are attached to a network namespace using a link. When [creating the link](../syscall/BPF_LINK_CREATE.md) the `prog_fd` to the file descriptor of the program, `target_fd` should be set to the file descriptor of a network namespace, and the `attach_type` to `BPF_SK_LOOKUP`.

## Example

```c
// Copyright (c) 2020 Cloudflare
struct {
	__uint(type, BPF_MAP_TYPE_SOCKMAP);
	__uint(max_entries, 32);
	__type(key, __u32);
	__type(value, __u64);
} redir_map SEC(".maps");

static const __u16 DST_PORT = 7007; /* Host byte order */
static const __u32 DST_IP4 = IP4(127, 0, 0, 1);
static const __u32 KEY_SERVER_A = 0;

/* Redirect packets destined for DST_IP4 address to socket at redir_map[0]. */
SEC("sk_lookup")
int redir_ip4(struct bpf_sk_lookup *ctx)
{
	struct bpf_sock *sk;
	int err;

	if (ctx->family != AF_INET)
		return SK_PASS;
	if (ctx->local_port != DST_PORT)
		return SK_PASS;
	if (ctx->local_ip4 != DST_IP4)
		return SK_PASS;

	sk = bpf_map_lookup_elem(&redir_map, &KEY_SERVER_A);
	if (!sk)
		return SK_PASS;

	err = bpf_sk_assign(ctx, sk, 0);
	bpf_sk_release(sk);
	return err ? SK_DROP : SK_PASS;
}
```

## Helper functions

Not all helper functions are available in all program types. These are the helper calls available for socket filter programs:

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [`bpf_perf_event_output`](../helper-function/bpf_perf_event_output.md)
    * [`bpf_sk_assign`](../helper-function/bpf_sk_assign.md)
    * [`bpf_sk_release`](../helper-function/bpf_sk_release.md)
    * [`bpf_skc_to_tcp6_sock`](../helper-function/bpf_skc_to_tcp6_sock.md)
    * [`bpf_skc_to_tcp_sock`](../helper-function/bpf_skc_to_tcp_sock.md)
    * [`bpf_skc_to_tcp_timewait_sock`](../helper-function/bpf_skc_to_tcp_timewait_sock.md)
    * [`bpf_skc_to_tcp_request_sock`](../helper-function/bpf_skc_to_tcp_request_sock.md)
    * [`bpf_skc_to_udp6_sock`](../helper-function/bpf_skc_to_udp6_sock.md)
    * [`bpf_skc_to_unix_sock`](../helper-function/bpf_skc_to_unix_sock.md)
    * [`bpf_ktime_get_coarse_ns`](../helper-function/bpf_ktime_get_coarse_ns.md)
    * [`bpf_map_lookup_elem`](../helper-function/bpf_map_lookup_elem.md)
    * [`bpf_map_update_elem`](../helper-function/bpf_map_update_elem.md)
    * [`bpf_map_delete_elem`](../helper-function/bpf_map_delete_elem.md)
    * [`bpf_map_push_elem`](../helper-function/bpf_map_push_elem.md)
    * [`bpf_map_pop_elem`](../helper-function/bpf_map_pop_elem.md)
    * [`bpf_map_peek_elem`](../helper-function/bpf_map_peek_elem.md)
    * [`bpf_map_lookup_percpu_elem`](../helper-function/bpf_map_lookup_percpu_elem.md)
    * [`bpf_get_prandom_u32`](../helper-function/bpf_get_prandom_u32.md)
    * [`bpf_get_smp_processor_id`](../helper-function/bpf_get_smp_processor_id.md)
    * [`bpf_get_numa_node_id`](../helper-function/bpf_get_numa_node_id.md)
    * [`bpf_tail_call`](../helper-function/bpf_tail_call.md)
    * [`bpf_ktime_get_ns`](../helper-function/bpf_ktime_get_ns.md)
    * [`bpf_ktime_get_boot_ns`](../helper-function/bpf_ktime_get_boot_ns.md)
    * [`bpf_ringbuf_output`](../helper-function/bpf_ringbuf_output.md)
    * [`bpf_ringbuf_reserve`](../helper-function/bpf_ringbuf_reserve.md)
    * [`bpf_ringbuf_submit`](../helper-function/bpf_ringbuf_submit.md)
    * [`bpf_ringbuf_discard`](../helper-function/bpf_ringbuf_discard.md)
    * [`bpf_ringbuf_query`](../helper-function/bpf_ringbuf_query.md)
    * [`bpf_for_each_map_elem`](../helper-function/bpf_for_each_map_elem.md)
    * [`bpf_loop`](../helper-function/bpf_loop.md)
    * [`bpf_strncmp`](../helper-function/bpf_strncmp.md)
    * [`bpf_spin_lock`](../helper-function/bpf_spin_lock.md)
    * [`bpf_spin_unlock`](../helper-function/bpf_spin_unlock.md)
    * [`bpf_jiffies64`](../helper-function/bpf_jiffies64.md)
    * [`bpf_per_cpu_ptr`](../helper-function/bpf_per_cpu_ptr.md)
    * [`bpf_this_cpu_ptr`](../helper-function/bpf_this_cpu_ptr.md)
    * [`bpf_timer_init`](../helper-function/bpf_timer_init.md)
    * [`bpf_timer_set_callback`](../helper-function/bpf_timer_set_callback.md)
    * [`bpf_timer_start`](../helper-function/bpf_timer_start.md)
    * [`bpf_timer_cancel`](../helper-function/bpf_timer_cancel.md)
    * [`bpf_trace_printk`](../helper-function/bpf_trace_printk.md)
    * [`bpf_get_current_task`](../helper-function/bpf_get_current_task.md)
    * [`bpf_get_current_task_btf`](../helper-function/bpf_get_current_task_btf.md)
    * [`bpf_probe_read_user`](../helper-function/bpf_probe_read_user.md)
    * [`bpf_probe_read_kernel`](../helper-function/bpf_probe_read_kernel.md)
    * [`bpf_probe_read_user_str`](../helper-function/bpf_probe_read_user_str.md)
    * [`bpf_probe_read_kernel_str`](../helper-function/bpf_probe_read_kernel_str.md)
    * [`bpf_snprintf_btf`](../helper-function/bpf_snprintf_btf.md)
    * [`bpf_snprintf`](../helper-function/bpf_snprintf.md)
    * [`bpf_task_pt_regs`](../helper-function/bpf_task_pt_regs.md)
    * [`bpf_trace_vprintk`](../helper-function/bpf_trace_vprintk.md)
<!-- [/PROG_HELPER_FUNC_REF] -->

## KFuncs

<!-- [PROG_KFUNC_REF] -->
There are currently no kfuncs supported for this program type
<!-- [/PROG_KFUNC_REF] -->
