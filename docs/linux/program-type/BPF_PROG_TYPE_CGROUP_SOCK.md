---
title: "Program Type 'BPF_PROG_TYPE_CGROUP_SOCK'"
description: "This page documents the 'BPF_PROG_TYPE_CGROUP_SOCK' eBPF program type, including its definition, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_CGROUP_SOCK`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_CGROUP_SOCK) -->
[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/61023658760032e97869b07d54be9681d2529e77)
<!-- [/FEATURE_TAG] -->

cGroup socket programs are attached to cGroups and triggered when sockets are created, released or bound by a process in the cGroup.

## Usage

cGroup socket programs are invoked when a socket is created, released or bound depending on the attach type which can be one of the following:

* `BPF_CGROUP_INET_SOCK_CREATE`
* `BPF_CGROUP_INET_SOCK_RELEASE`
* `BPF_CGROUP_INET4_POST_BIND`
* `BPF_CGROUP_INET6_POST_BIND`

The ELF sections typically used for the respective attach types are:

* `cgroup/sock_create` 
* `cgroup/sock_release`
* `cgroup/post_bind4` 
* `cgroup/post_bind6` 

All attach types can be used for monitoring. The create and release attach types can modify the `bound_dev_if`, `mark` and `priority` fields, the rest of the attach types can only read the fields. Lastly, all attach types can block the operation by returning `0`, returning `1` allows the operation to proceed.

## Context

The context for cGroup socket programs is a `struct bpf_sock`.

??? abstract "C structure"
    ```c
    struct bpf_sock {
        __u32 bound_dev_if;
        __u32 family;
        __u32 type;
        __u32 protocol;
        __u32 mark;
        __u32 priority;
        /* IP address also allows 1 and 2 bytes access */
        __u32 src_ip4;
        __u32 src_ip6[4];
        __u32 src_port;		/* host byte order */
        __be16 dst_port;	/* network byte order */
        __u16 :16;		/* zero padding */
        __u32 dst_ip4;
        __u32 dst_ip6[4];
        __u32 state;
        __s32 rx_queue_mapping;
    };
    ```

### `bound_dev_if`

[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/61023658760032e97869b07d54be9681d2529e77)

This field contains the device index of the network device the socket is bound to. `BPF_CGROUP_INET_SOCK_CREATE` and `BPF_CGROUP_INET_SOCK_RELEASE` attached programs can modify this field.

### `family`

[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/aa4c1037a30f4e88f444e83d42c2befbe0d5caf5)

This field contains the address family of the socket. Its value is one of `AF_*` values defined in `include/linux/socket.h`.

### `type`

[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/aa4c1037a30f4e88f444e83d42c2befbe0d5caf5)

This field contains the socket type. Its value is one of `SOCK_*` values defined in `include/linux/net.h`.

### `protocol`

[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/aa4c1037a30f4e88f444e83d42c2befbe0d5caf5)

This field contains the socket protocol. Its value is one of `IPPROTO_*` values defined in `include/uapi/linux/in.h`.

### `mark`

[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/482dca939fb7ee35ba20b944b4c2476133dbf0df)

This field contains the socket mark. `BPF_CGROUP_INET_SOCK_CREATE` and `BPF_CGROUP_INET_SOCK_RELEASE` attached programs can modify this field.

### `priority`

[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/482dca939fb7ee35ba20b944b4c2476133dbf0df)

This field contains the socket priority. `BPF_CGROUP_INET_SOCK_CREATE` and `BPF_CGROUP_INET_SOCK_RELEASE` attached programs can modify this field.

### `src_ip4`

[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/aa65d6960a98fc15a96ce361b26e9fd55c9bccc5)

This field contains the source IPv4 address of the socket. `BPF_CGROUP_INET4_POST_BIND` and `BPF_CGROUP_INET6_POST_BIND` attached programs can read this field. Other attach types are not allowed to read or write this field.

### `src_ip6`

[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/aa65d6960a98fc15a96ce361b26e9fd55c9bccc5)

This field contains the source IPv6 address of the socket. `BPF_CGROUP_INET4_POST_BIND` and `BPF_CGROUP_INET6_POST_BIND` attached programs can read this field. Other attach types are not allowed to read or write this field.

### `src_port`

[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/aa65d6960a98fc15a96ce361b26e9fd55c9bccc5)

This field contains the source port of the socket. `BPF_CGROUP_INET4_POST_BIND` and `BPF_CGROUP_INET6_POST_BIND` attached programs can read this field. Other attach types are not allowed to read or write this field.

### `dst_port`

[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/aa65d6960a98fc15a96ce361b26e9fd55c9bccc5)

This field contains the destination port of the socket. `BPF_CGROUP_INET4_POST_BIND` and `BPF_CGROUP_INET6_POST_BIND` attached programs can read this field. Other attach types are not allowed to read or write this field.

### `dst_ip4`

[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/aa65d6960a98fc15a96ce361b26e9fd55c9bccc5)

This field contains the destination IPv4 address of the socket. `BPF_CGROUP_INET4_POST_BIND` and `BPF_CGROUP_INET6_POST_BIND` attached programs can read this field. Other attach types are not allowed to read or write this field.

### `dst_ip6`

[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/aa65d6960a98fc15a96ce361b26e9fd55c9bccc5)

This field contains the destination IPv6 address of the socket. `BPF_CGROUP_INET4_POST_BIND` and `BPF_CGROUP_INET6_POST_BIND` attached programs can read this field. Other attach types are not allowed to read or write this field.

### `state`

[:octicons-tag-24: v5.1](https://github.com/torvalds/linux/commit/aa65d6960a98fc15a96ce361b26e9fd55c9bccc5)

This field contains the connection state of the socket.

The states will be one of:
```
enum {
	BPF_TCP_ESTABLISHED = 1,
	BPF_TCP_SYN_SENT,
	BPF_TCP_SYN_RECV,
	BPF_TCP_FIN_WAIT1,
	BPF_TCP_FIN_WAIT2,
	BPF_TCP_TIME_WAIT,
	BPF_TCP_CLOSE,
	BPF_TCP_CLOSE_WAIT,
	BPF_TCP_LAST_ACK,
	BPF_TCP_LISTEN,
	BPF_TCP_CLOSING,	/* Now a valid state */
	BPF_TCP_NEW_SYN_RECV
};
```

### `rx_queue_mapping`

[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/c3c16f2ea6d20159903cf93afbb1155f3d8348d5)

This field contains the receive queue number for the connection. The Rx queue
is marked in `tcp_finish_connect()` and is otherwise `-1`.

## Attachment

cGroup socket buffer programs are attached to cgroups via the [`BPF_PROG_ATTACH`](../syscall/BPF_PROG_ATTACH.md) syscall or via [BPF link](../syscall/BPF_LINK_CREATE.md).

## Example

Example BPF program:

```c
// SPDX-License-Identifier: GPL-2.0-only

#include <sys/socket.h>
#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>

int invocations = 0, in_use = 0;

struct {
	__uint(type, BPF_MAP_TYPE_SK_STORAGE);
	__uint(map_flags, BPF_F_NO_PREALLOC);
	__type(key, int);
	__type(value, int);
} sk_map SEC(".maps");

SEC("cgroup/sock_create")
int sock(struct bpf_sock *ctx)
{
	int *sk_storage;

	if (ctx->type != SOCK_DGRAM)
		return 1;

	sk_storage = bpf_sk_storage_get(&sk_map, ctx, 0,
					BPF_SK_STORAGE_GET_F_CREATE);
	if (!sk_storage)
		return 0;
	*sk_storage = 0xdeadbeef;

	__sync_fetch_and_add(&invocations, 1);

	if (in_use > 0) {
		/* BPF_CGROUP_INET_SOCK_RELEASE is _not_ called
		 * when we return an error from the BPF
		 * program!
		 */
		return 0;
	}

	__sync_fetch_and_add(&in_use, 1);
	return 1;
}

SEC("cgroup/sock_release")
int sock_release(struct bpf_sock *ctx)
{
	int *sk_storage;

	if (ctx->type != SOCK_DGRAM)
		return 1;

	sk_storage = bpf_sk_storage_get(&sk_map, ctx, 0, 0);
	if (!sk_storage || *sk_storage != 0xdeadbeef)
		return 0;

	__sync_fetch_and_add(&invocations, 1);
	__sync_fetch_and_add(&in_use, -1);
	return 1;
}
```

## Helper functions

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [`bpf_get_current_uid_gid`](../helper-function/bpf_get_current_uid_gid.md)
    * [`bpf_get_local_storage`](../helper-function/bpf_get_local_storage.md)
    * [`bpf_get_socket_cookie`](../helper-function/bpf_get_socket_cookie.md)
    * [`bpf_get_netns_cookie`](../helper-function/bpf_get_netns_cookie.md)
    * [`bpf_perf_event_output`](../helper-function/bpf_perf_event_output.md)
    * [`bpf_get_current_pid_tgid`](../helper-function/bpf_get_current_pid_tgid.md)
    * [`bpf_get_current_comm`](../helper-function/bpf_get_current_comm.md)
    * [`bpf_get_current_cgroup_id`](../helper-function/bpf_get_current_cgroup_id.md)
    * [`bpf_get_current_ancestor_cgroup_id`](../helper-function/bpf_get_current_ancestor_cgroup_id.md)
    * [`bpf_get_cgroup_classid`](../helper-function/bpf_get_cgroup_classid.md)
    * [`bpf_sk_storage_get`](../helper-function/bpf_sk_storage_get.md)
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
