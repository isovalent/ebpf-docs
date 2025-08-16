---
title: "Program Type 'BPF_PROG_TYPE_CGROUP_SOCK_ADDR'"
description: "This page documents the 'BPF_PROG_TYPE_CGROUP_SOCK_ADDR' eBPF program type, including its definition, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_CGROUP_SOCK_ADDR`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_CGROUP_SOCK_ADDR) -->
[:octicons-tag-24: v4.17](https://github.com/torvalds/linux/commit/4fbac77d2d092b475dda9eea66da674369665427)
<!-- [/FEATURE_TAG] -->

cGroup socket address programs are triggered when a process in a cGroup to which the program is attached uses socket related syscalls. This program can overwrite arguments to the syscall such as addresses.

## Usage

This program type can be used to overwrite arguments to socket related syscalls or to block the call to the syscall entirely. Which syscall depends on the attach type used.

### `BPF_CGROUP_INET4_BIND` and `BPF_CGROUP_INET6_BIND`

<!-- [FEATURE_TAG](BPF_CGROUP_INET4_BIND) -->
[:octicons-tag-24: v4.17](https://github.com/torvalds/linux/commit/4fbac77d2d092b475dda9eea66da674369665427)
<!-- [/FEATURE_TAG] -->

This attach type is triggered when a process calls the `bind` syscall with an IPv4 or IPv6 address respectively. The typical ELF sections used for this attach type are: `cgroup/bind4` and `cgroup/bind6`.

!!! note
    Since [:octicons-tag-24: v5.12](https://github.com/torvalds/linux/commit/772412176fb98493158929b220fe250127f611af) the 2's bit of the return value is used to indicate that checking for the `CAP_NET_BIND_SERVICE` capability can be skipped. Normally this capability is required when binding to a privileged port (`<1024`). So when a BPF program rewrites the listening port on a process without the capability it can set this bit to prevent the kernel from blocking the call. 

### `BPF_CGROUP_INET4_CONNECT` and `BPF_CGROUP_INET6_CONNECT`    

<!-- [FEATURE_TAG](BPF_CGROUP_INET4_CONNECT) -->
[:octicons-tag-24: v4.17](https://github.com/torvalds/linux/commit/d74bad4e74ee373787a9ae24197c17b7cdc428d5)
<!-- [/FEATURE_TAG] -->

This attach type is triggered when a process calls the `connect` syscall with an IPv4 or IPv6 address respectively. The typical ELF sections used for this attach type are: `cgroup/connect4` and `cgroup/connect6`.

### `BPF_CGROUP_UDP4_SENDMSG` and `BPF_CGROUP_UDP6_SENDMSG`     

<!-- [FEATURE_TAG](BPF_CGROUP_UDP4_SENDMSG) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/1cedee13d25ab118d325f95588c1a084e9317229)
<!-- [/FEATURE_TAG] -->

This attach type is triggered when a process calls the `sendmsg` syscall with an IPv4 or IPv6 address respectively. The typical ELF sections used for this attach type are: `cgroup/sendmsg4` and `cgroup/sendmsg6`.

### `BPF_CGROUP_UDP4_RECVMSG` and `BPF_CGROUP_UDP6_RECVMSG`  

<!-- [FEATURE_TAG](BPF_CGROUP_UDP4_RECVMSG) -->
[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/983695fa676568fc0fe5ddd995c7267aabc24632)
<!-- [/FEATURE_TAG] -->

This attach type is triggered when a process calls the `recvmsg` syscall with an IPv4 or IPv6 address respectively. The typical ELF sections used for this attach type are: `cgroup/recvmsg4` and `cgroup/recvmsg6`.

### `BPF_CGROUP_INET4_GETPEERNAME` and `BPF_CGROUP_INET6_GETPEERNAME` 

<!-- [FEATURE_TAG](BPF_CGROUP_INET4_GETPEERNAME) -->
[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/1b66d253610c7f8f257103808a9460223a087469)
<!-- [/FEATURE_TAG] -->

This attach type is triggered when a process calls the `getpeername` syscall with an IPv4 or IPv6 address respectively. The typical ELF sections used for this attach type are: `cgroup/getpeername4` and `cgroup/getpeername6`.

### `BPF_CGROUP_INET4_GETSOCKNAME` and `BPF_CGROUP_INET6_GETSOCKNAME` 

<!-- [FEATURE_TAG](BPF_CGROUP_INET4_GETSOCKNAME) -->
[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/1b66d253610c7f8f257103808a9460223a087469)
<!-- [/FEATURE_TAG] -->

This attach type is triggered when a process calls the `getsockname` syscall with an IPv4 or IPv6 address respectively. The typical ELF sections used for this attach type are: `cgroup/getsockname4` and `cgroup/getsockname6`.

## Context

??? abstract "C structure"
    ```c
    struct bpf_sock_addr {
        __u32 user_family;	/* Allows 4-byte read, but no write. */
        __u32 user_ip4;		/* Allows 1,2,4-byte read and 4-byte write.
                    * Stored in network byte order.
                    */
        __u32 user_ip6[4];	/* Allows 1,2,4,8-byte read and 4,8-byte write.
                    * Stored in network byte order.
                    */
        __u32 user_port;	/* Allows 1,2,4-byte read and 4-byte write.
                    * Stored in network byte order
                    */
        __u32 family;		/* Allows 4-byte read, but no write */
        __u32 type;		/* Allows 4-byte read, but no write */
        __u32 protocol;		/* Allows 4-byte read, but no write */
        __u32 msg_src_ip4;	/* Allows 1,2,4-byte read and 4-byte write.
                    * Stored in network byte order.
                    */
        __u32 msg_src_ip6[4];	/* Allows 1,2,4,8-byte read and 4,8-byte write.
                    * Stored in network byte order.
                    */
        __bpf_md_ptr(struct bpf_sock *, sk);
    };
    ```

### `user_family` 

[:octicons-tag-24: v4.17](https://github.com/torvalds/linux/commit/4fbac77d2d092b475dda9eea66da674369665427)

This field contains the address family passed to the syscall. Its value is one of `AF_*` values defined in `include/linux/socket.h`.

The context allows 4-byte reads from the field, but no writes to it.

### `user_ip4` 

[:octicons-tag-24: v4.17](https://github.com/torvalds/linux/commit/4fbac77d2d092b475dda9eea66da674369665427)

This field contains the IPv4 address passed to the syscall. Its value is stored in network byte order. This field is only valid of `INET4` attach types.

The context allows 1,2,4-byte reads and 4-byte writes.

### `user_ip6` 

[:octicons-tag-24: v4.17](https://github.com/torvalds/linux/commit/4fbac77d2d092b475dda9eea66da674369665427)

This field contains the IPv6 address passed to the syscall. Its value is stored in network byte order. This field is only valid of `INET6` attach types.

This context allows 1,2,4,8-byte reads and 4,8-byte writes.

!!! note
    8-byte wide loads are only supported since [:octicons-tag-24: v5.3](https://github.com/torvalds/linux/commit/d4ecfeb15494ec261fef2d25d96eecba66f0b182) 


### `user_port` 

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/1cedee13d25ab118d325f95588c1a084e9317229)

This field contains the port number passed to the syscall. Its value is stored in network byte order.

This context allows 1,2,4-byte reads and 4-byte writes.

### `family` 

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/1cedee13d25ab118d325f95588c1a084e9317229)

This field contains the address family of the socket. Its value is one of `AF_*` values defined in `include/linux/socket.h`.

The context allows 4-byte reads from the field, but no writes to it.

### `type` 

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/1cedee13d25ab118d325f95588c1a084e9317229)

This field contains the socket type. Its value is one of `SOCK_*` values defined in `include/linux/socket.h`.

This context allows 4-byte reads from the field, but no writes to it.

### `protocol` 

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/1cedee13d25ab118d325f95588c1a084e9317229)

This field contains the socket protocol. Its value is one of `IPPROTO_*` values defined in `include/linux/socket.h`.

This context allows 4-byte reads from the field, but no writes to it.

### `msg_src_ip4` 

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/1cedee13d25ab118d325f95588c1a084e9317229)

This field contains a IPv4 address which is the source IP of the message about to be sent. Its value is stored in network byte order.

This field is only valid of `BPF_CGROUP_UDP4_SENDMSG` attach type.

This context allows 1,2,4-byte reads and 4-byte writes.

### `msg_src_ip6` 

[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/1cedee13d25ab118d325f95588c1a084e9317229)

This field contains a IPv6 address which is the source IP of the message about to be sent. Its value is stored in network byte order.

This field is only valid of `BPF_CGROUP_UDP6_SENDMSG` attach type.

This context allows 1,2,4,8-byte reads and 4,8-byte writes.

!!! note
    8-byte wide loads are only supported since [:octicons-tag-24: v5.3](https://github.com/torvalds/linux/commit/d4ecfeb15494ec261fef2d25d96eecba66f0b182) 

### `sk`

[:octicons-tag-24: v5.3](https://github.com/torvalds/linux/commit/fb85c4a730af221339c1dde1a434b73da0dfc3ed)

This field contains a pointer to the socket for which the program was invoked, its type being a `struct bpf_sock`.

## Attachment

cGroup socket buffer programs are attached to cGroups via the [`BPF_PROG_ATTACH`](../syscall/BPF_PROG_ATTACH.md) syscall or via [BPF link](../syscall/BPF_LINK_CREATE.md).

## Example

??? example "`BPF_CGROUP_INET4_BIND` and `BPF_CGROUP_INET6_BIND`"
    ```c
    // SPDX-License-Identifier: GPL-2.0

    #include <linux/stddef.h>
    #include <linux/bpf.h>
    #include <sys/types.h>
    #include <sys/socket.h>
    #include <bpf/bpf_helpers.h>
    #include <bpf/bpf_endian.h>

    static __always_inline int bind_prog(struct bpf_sock_addr *ctx, int family)
    {
        struct bpf_sock *sk;

        sk = ctx->sk;
        if (!sk)
            return 0;

        if (sk->family != family)
            return 0;

        if (ctx->type != SOCK_STREAM)
            return 0;

        /* Return 1 OR'ed with the first bit set to indicate
        * that CAP_NET_BIND_SERVICE should be bypassed.
        */
        if (ctx->user_port == bpf_htons(111))
            return (1 | 2);

        return 1;
    }

    SEC("cgroup/bind4")
    int bind_v4_prog(struct bpf_sock_addr *ctx)
    {
        return bind_prog(ctx, AF_INET);
    }

    SEC("cgroup/bind6")
    int bind_v6_prog(struct bpf_sock_addr *ctx)
    {
        return bind_prog(ctx, AF_INET6);
    }

    char _license[] SEC("license") = "GPL";
    ```

??? example "`BPF_CGROUP_INET4_CONNECT`, `BPF_CGROUP_INET4_GETSOCKNAME`, and `BPF_CGROUP_INET4_GETPEERNAME`"
    ```c
    // SPDX-License-Identifier: GPL-2.0
    #include <string.h>
    #include <stdbool.h>

    #include <linux/bpf.h>
    #include <linux/in.h>
    #include <linux/in6.h>
    #include <sys/socket.h>

    #include <bpf/bpf_helpers.h>
    #include <bpf/bpf_endian.h>

    #include <bpf_sockopt_helpers.h>

    char _license[] SEC("license") = "GPL";

    struct svc_addr {
        __be32 addr;
        __be16 port;
    };

    struct {
        __uint(type, BPF_MAP_TYPE_SK_STORAGE);
        __uint(map_flags, BPF_F_NO_PREALLOC);
        __type(key, int);
        __type(value, struct svc_addr);
    } service_mapping SEC(".maps");

    SEC("cgroup/connect4")
    int connect4(struct bpf_sock_addr *ctx)
    {
        struct sockaddr_in sa = {};
        struct svc_addr *orig;

        /* Force local address to 127.0.0.1:22222. */
        sa.sin_family = AF_INET;
        sa.sin_port = bpf_htons(22222);
        sa.sin_addr.s_addr = bpf_htonl(0x7f000001);

        if (bpf_bind(ctx, (struct sockaddr *)&sa, sizeof(sa)) != 0)
            return 0;

        /* Rewire service 1.2.3.4:60000 to backend 127.0.0.1:60123. */
        if (ctx->user_port == bpf_htons(60000)) {
            orig = bpf_sk_storage_get(&service_mapping, ctx->sk, 0,
                        BPF_SK_STORAGE_GET_F_CREATE);
            if (!orig)
                return 0;

            orig->addr = ctx->user_ip4;
            orig->port = ctx->user_port;

            ctx->user_ip4 = bpf_htonl(0x7f000001);
            ctx->user_port = bpf_htons(60123);
        }
        return 1;
    }

    SEC("cgroup/getsockname4")
    int getsockname4(struct bpf_sock_addr *ctx)
    {
        if (!get_set_sk_priority(ctx))
            return 1;

        /* Expose local server as 1.2.3.4:60000 to client. */
        if (ctx->user_port == bpf_htons(60123)) {
            ctx->user_ip4 = bpf_htonl(0x01020304);
            ctx->user_port = bpf_htons(60000);
        }
        return 1;
    }

    SEC("cgroup/getpeername4")
    int getpeername4(struct bpf_sock_addr *ctx)
    {
        struct svc_addr *orig;

        if (!get_set_sk_priority(ctx))
            return 1;

        /* Expose service 1.2.3.4:60000 as peer instead of backend. */
        if (ctx->user_port == bpf_htons(60123)) {
            orig = bpf_sk_storage_get(&service_mapping, ctx->sk, 0, 0);
            if (orig) {
                ctx->user_ip4 = orig->addr;
                ctx->user_port = orig->port;
            }
        }
        return 1;
    }
    ```

## Helper functions

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [`bpf_bind`](../helper-function/bpf_bind.md)
    * [`bpf_cgrp_storage_delete`](../helper-function/bpf_cgrp_storage_delete.md)
    * [`bpf_cgrp_storage_get`](../helper-function/bpf_cgrp_storage_get.md)
    * [`bpf_dynptr_data`](../helper-function/bpf_dynptr_data.md)
    * [`bpf_dynptr_from_mem`](../helper-function/bpf_dynptr_from_mem.md)
    * [`bpf_dynptr_read`](../helper-function/bpf_dynptr_read.md)
    * [`bpf_dynptr_write`](../helper-function/bpf_dynptr_write.md)
    * [`bpf_for_each_map_elem`](../helper-function/bpf_for_each_map_elem.md)
    * [`bpf_get_cgroup_classid`](../helper-function/bpf_get_cgroup_classid.md)
    * [`bpf_get_current_ancestor_cgroup_id`](../helper-function/bpf_get_current_ancestor_cgroup_id.md)
    * [`bpf_get_current_cgroup_id`](../helper-function/bpf_get_current_cgroup_id.md)
    * [`bpf_get_current_comm`](../helper-function/bpf_get_current_comm.md)
    * [`bpf_get_current_pid_tgid`](../helper-function/bpf_get_current_pid_tgid.md)
    * [`bpf_get_current_task`](../helper-function/bpf_get_current_task.md)
    * [`bpf_get_current_task_btf`](../helper-function/bpf_get_current_task_btf.md)
    * [`bpf_get_current_uid_gid`](../helper-function/bpf_get_current_uid_gid.md)
    * [`bpf_get_local_storage`](../helper-function/bpf_get_local_storage.md)
    * [`bpf_get_netns_cookie`](../helper-function/bpf_get_netns_cookie.md)
    * [`bpf_get_ns_current_pid_tgid`](../helper-function/bpf_get_ns_current_pid_tgid.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
    * [`bpf_get_numa_node_id`](../helper-function/bpf_get_numa_node_id.md)
    * [`bpf_get_prandom_u32`](../helper-function/bpf_get_prandom_u32.md)
    * [`bpf_get_smp_processor_id`](../helper-function/bpf_get_smp_processor_id.md)
    * [`bpf_get_socket_cookie`](../helper-function/bpf_get_socket_cookie.md)
    * [`bpf_getsockopt`](../helper-function/bpf_getsockopt.md) [:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/beecf11bc2188067824591612151c4dc6ec383c7)
    * [`bpf_jiffies64`](../helper-function/bpf_jiffies64.md)
    * [`bpf_kptr_xchg`](../helper-function/bpf_kptr_xchg.md)
    * [`bpf_ktime_get_boot_ns`](../helper-function/bpf_ktime_get_boot_ns.md)
    * [`bpf_ktime_get_ns`](../helper-function/bpf_ktime_get_ns.md)
    * [`bpf_ktime_get_tai_ns`](../helper-function/bpf_ktime_get_tai_ns.md)
    * [`bpf_loop`](../helper-function/bpf_loop.md)
    * [`bpf_map_delete_elem`](../helper-function/bpf_map_delete_elem.md)
    * [`bpf_map_lookup_elem`](../helper-function/bpf_map_lookup_elem.md)
    * [`bpf_map_lookup_percpu_elem`](../helper-function/bpf_map_lookup_percpu_elem.md)
    * [`bpf_map_peek_elem`](../helper-function/bpf_map_peek_elem.md)
    * [`bpf_map_pop_elem`](../helper-function/bpf_map_pop_elem.md)
    * [`bpf_map_push_elem`](../helper-function/bpf_map_push_elem.md)
    * [`bpf_map_update_elem`](../helper-function/bpf_map_update_elem.md)
    * [`bpf_per_cpu_ptr`](../helper-function/bpf_per_cpu_ptr.md)
    * [`bpf_perf_event_output`](../helper-function/bpf_perf_event_output.md)
    * [`bpf_probe_read_kernel`](../helper-function/bpf_probe_read_kernel.md)
    * [`bpf_probe_read_kernel_str`](../helper-function/bpf_probe_read_kernel_str.md)
    * [`bpf_probe_read_user`](../helper-function/bpf_probe_read_user.md)
    * [`bpf_probe_read_user_str`](../helper-function/bpf_probe_read_user_str.md)
    * [`bpf_ringbuf_discard`](../helper-function/bpf_ringbuf_discard.md)
    * [`bpf_ringbuf_discard_dynptr`](../helper-function/bpf_ringbuf_discard_dynptr.md)
    * [`bpf_ringbuf_output`](../helper-function/bpf_ringbuf_output.md)
    * [`bpf_ringbuf_query`](../helper-function/bpf_ringbuf_query.md)
    * [`bpf_ringbuf_reserve`](../helper-function/bpf_ringbuf_reserve.md)
    * [`bpf_ringbuf_reserve_dynptr`](../helper-function/bpf_ringbuf_reserve_dynptr.md)
    * [`bpf_ringbuf_submit`](../helper-function/bpf_ringbuf_submit.md)
    * [`bpf_ringbuf_submit_dynptr`](../helper-function/bpf_ringbuf_submit_dynptr.md)
    * [`bpf_setsockopt`](../helper-function/bpf_setsockopt.md) [:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/beecf11bc2188067824591612151c4dc6ec383c7)
    * [`bpf_sk_lookup_tcp`](../helper-function/bpf_sk_lookup_tcp.md)
    * [`bpf_sk_lookup_udp`](../helper-function/bpf_sk_lookup_udp.md)
    * [`bpf_sk_release`](../helper-function/bpf_sk_release.md)
    * [`bpf_sk_storage_delete`](../helper-function/bpf_sk_storage_delete.md)
    * [`bpf_sk_storage_get`](../helper-function/bpf_sk_storage_get.md)
    * [`bpf_skc_lookup_tcp`](../helper-function/bpf_skc_lookup_tcp.md)
    * [`bpf_snprintf`](../helper-function/bpf_snprintf.md)
    * [`bpf_snprintf_btf`](../helper-function/bpf_snprintf_btf.md)
    * [`bpf_spin_lock`](../helper-function/bpf_spin_lock.md)
    * [`bpf_spin_unlock`](../helper-function/bpf_spin_unlock.md)
    * [`bpf_strncmp`](../helper-function/bpf_strncmp.md)
    * [`bpf_tail_call`](../helper-function/bpf_tail_call.md)
    * [`bpf_task_pt_regs`](../helper-function/bpf_task_pt_regs.md)
    * [`bpf_this_cpu_ptr`](../helper-function/bpf_this_cpu_ptr.md)
    * [`bpf_timer_cancel`](../helper-function/bpf_timer_cancel.md)
    * [`bpf_timer_init`](../helper-function/bpf_timer_init.md)
    * [`bpf_timer_set_callback`](../helper-function/bpf_timer_set_callback.md)
    * [`bpf_timer_start`](../helper-function/bpf_timer_start.md)
    * [`bpf_trace_printk`](../helper-function/bpf_trace_printk.md)
    * [`bpf_trace_vprintk`](../helper-function/bpf_trace_vprintk.md)
    * [`bpf_user_ringbuf_drain`](../helper-function/bpf_user_ringbuf_drain.md)
<!-- [/PROG_HELPER_FUNC_REF] -->

## KFuncs

<!-- [PROG_KFUNC_REF] -->
??? abstract "Supported kfuncs"
    - [`__bpf_trap`](../kfuncs/__bpf_trap.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_arena_alloc_pages`](../kfuncs/bpf_arena_alloc_pages.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_arena_free_pages`](../kfuncs/bpf_arena_free_pages.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_cast_to_kern_ctx`](../kfuncs/bpf_cast_to_kern_ctx.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_cgroup_read_xattr`](../kfuncs/bpf_cgroup_read_xattr.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_copy_from_user_dynptr`](../kfuncs/bpf_copy_from_user_dynptr.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_copy_from_user_str`](../kfuncs/bpf_copy_from_user_str.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_copy_from_user_str_dynptr`](../kfuncs/bpf_copy_from_user_str_dynptr.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_copy_from_user_task_dynptr`](../kfuncs/bpf_copy_from_user_task_dynptr.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_copy_from_user_task_str`](../kfuncs/bpf_copy_from_user_task_str.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_copy_from_user_task_str_dynptr`](../kfuncs/bpf_copy_from_user_task_str_dynptr.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_dynptr_adjust`](../kfuncs/bpf_dynptr_adjust.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_dynptr_clone`](../kfuncs/bpf_dynptr_clone.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_dynptr_copy`](../kfuncs/bpf_dynptr_copy.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_dynptr_from_skb`](../kfuncs/bpf_dynptr_from_skb.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_dynptr_is_null`](../kfuncs/bpf_dynptr_is_null.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_dynptr_is_rdonly`](../kfuncs/bpf_dynptr_is_rdonly.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_dynptr_memset`](../kfuncs/bpf_dynptr_memset.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_dynptr_size`](../kfuncs/bpf_dynptr_size.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_dynptr_slice`](../kfuncs/bpf_dynptr_slice.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_dynptr_slice_rdwr`](../kfuncs/bpf_dynptr_slice_rdwr.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_get_kmem_cache`](../kfuncs/bpf_get_kmem_cache.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_bits_destroy`](../kfuncs/bpf_iter_bits_destroy.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_bits_new`](../kfuncs/bpf_iter_bits_new.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_bits_next`](../kfuncs/bpf_iter_bits_next.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_css_destroy`](../kfuncs/bpf_iter_css_destroy.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_css_new`](../kfuncs/bpf_iter_css_new.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_css_next`](../kfuncs/bpf_iter_css_next.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_css_task_destroy`](../kfuncs/bpf_iter_css_task_destroy.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_css_task_new`](../kfuncs/bpf_iter_css_task_new.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_css_task_next`](../kfuncs/bpf_iter_css_task_next.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_dmabuf_destroy`](../kfuncs/bpf_iter_dmabuf_destroy.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_dmabuf_new`](../kfuncs/bpf_iter_dmabuf_new.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_dmabuf_next`](../kfuncs/bpf_iter_dmabuf_next.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_kmem_cache_destroy`](../kfuncs/bpf_iter_kmem_cache_destroy.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_kmem_cache_new`](../kfuncs/bpf_iter_kmem_cache_new.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_kmem_cache_next`](../kfuncs/bpf_iter_kmem_cache_next.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_num_destroy`](../kfuncs/bpf_iter_num_destroy.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_num_new`](../kfuncs/bpf_iter_num_new.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_num_next`](../kfuncs/bpf_iter_num_next.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_task_destroy`](../kfuncs/bpf_iter_task_destroy.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_task_new`](../kfuncs/bpf_iter_task_new.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_task_next`](../kfuncs/bpf_iter_task_next.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_task_vma_destroy`](../kfuncs/bpf_iter_task_vma_destroy.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_task_vma_new`](../kfuncs/bpf_iter_task_vma_new.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_iter_task_vma_next`](../kfuncs/bpf_iter_task_vma_next.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_local_irq_restore`](../kfuncs/bpf_local_irq_restore.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_local_irq_save`](../kfuncs/bpf_local_irq_save.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_map_sum_elem_count`](../kfuncs/bpf_map_sum_elem_count.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_preempt_disable`](../kfuncs/bpf_preempt_disable.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_preempt_enable`](../kfuncs/bpf_preempt_enable.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_probe_read_kernel_dynptr`](../kfuncs/bpf_probe_read_kernel_dynptr.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_probe_read_kernel_str_dynptr`](../kfuncs/bpf_probe_read_kernel_str_dynptr.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_probe_read_user_dynptr`](../kfuncs/bpf_probe_read_user_dynptr.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_probe_read_user_str_dynptr`](../kfuncs/bpf_probe_read_user_str_dynptr.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_rcu_read_lock`](../kfuncs/bpf_rcu_read_lock.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_rcu_read_unlock`](../kfuncs/bpf_rcu_read_unlock.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_rdonly_cast`](../kfuncs/bpf_rdonly_cast.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_res_spin_lock`](../kfuncs/bpf_res_spin_lock.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_res_spin_lock_irqsave`](../kfuncs/bpf_res_spin_lock_irqsave.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_res_spin_unlock`](../kfuncs/bpf_res_spin_unlock.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_res_spin_unlock_irqrestore`](../kfuncs/bpf_res_spin_unlock_irqrestore.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_sock_addr_set_sun_path`](../kfuncs/bpf_sock_addr_set_sun_path.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_sock_ops_enable_tx_tstamp`](../kfuncs/bpf_sock_ops_enable_tx_tstamp.md)
    - [`bpf_strchr`](../kfuncs/bpf_strchr.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_strchrnul`](../kfuncs/bpf_strchrnul.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_strcmp`](../kfuncs/bpf_strcmp.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_strcspn`](../kfuncs/bpf_strcspn.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_stream_vprintk`](../kfuncs/bpf_stream_vprintk.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_strlen`](../kfuncs/bpf_strlen.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_strnchr`](../kfuncs/bpf_strnchr.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_strnlen`](../kfuncs/bpf_strnlen.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_strnstr`](../kfuncs/bpf_strnstr.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_strrchr`](../kfuncs/bpf_strrchr.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_strspn`](../kfuncs/bpf_strspn.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_strstr`](../kfuncs/bpf_strstr.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_wq_init`](../kfuncs/bpf_wq_init.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_wq_set_callback_impl`](../kfuncs/bpf_wq_set_callback_impl.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
    - [`bpf_wq_start`](../kfuncs/bpf_wq_start.md) [:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/53e380d21441909b12b6e0782b77187ae4b971c4) - 
<!-- [/PROG_KFUNC_REF] -->
