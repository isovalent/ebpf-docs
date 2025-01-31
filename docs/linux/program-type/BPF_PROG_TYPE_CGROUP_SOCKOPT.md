---
title: "Program Type 'BPF_PROG_TYPE_CGROUP_SOCKOPT'"
description: "This page documents the 'BPF_PROG_TYPE_CGROUP_SOCKOPT' eBPF program type, including its definition, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_CGROUP_SOCKOPT`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_CGROUP_SOCKOPT) -->
[:octicons-tag-24: v5.3](https://github.com/torvalds/linux/commit/0d01da6afc5402f60325c5da31b22f7d56689b49)
<!-- [/FEATURE_TAG] -->

cGroup socket ops programs are executed when a process in the cGroup to which the program is attached uses the [`getsockopt`](https://linux.die.net/man/2/getsockopt) or [`setsockopt`](https://linux.die.net/man/2/setsockopt) syscall depending on the attach type and modify or block the operation.

## Usage

cGroup socket ops programs are typically located in the `cgroup/getsockopt` or `cgroup/setsockopt` ELF section to indicate the `BPF_CGROUP_GETSOCKOPT` and `BPF_CGROUP_SETSOCKOPT` attach types respectively.

### `BPF_CGROUP_SETSOCKOPT`

`BPF_CGROUP_SETSOCKOPT` is triggered *before* the kernel handling of sockopt and it has writable context: it can modify the supplied arguments before passing them down to the kernel. This hook has access to the cGroup and socket local storage.

If BPF program sets [`optlen`](#optlen) to -1, the control will be returned back to the userspace after all other BPF programs in the cGroup chain finish (i.e. kernel [`setsockopt`](https://linux.die.net/man/2/setsockopt) handling will *not* be executed).

!!! note 
     [`optlen`](#optlen) can not be increased beyond the user-supplied value. It can only be decreased or set to -1. Any other value will trigger `EFAULT`.

Return Type:

* `0` - reject the syscall, `EPERM` will be returned to the userspace.
* `1` - success, continue with next BPF program in the cgroup chain.

### `BPF_CGROUP_GETSOCKOPT`

`BPF_CGROUP_GETSOCKOPT` is triggered *after* the kernel handing of sockopt. The BPF hook can observe [`optval`](#optval), [`optlen`](#optlen) and [`retval`](#retval) if it's interested in whatever kernel has returned. BPF hook can override the values above, adjust  [`optlen`](#optlen) and reset [`retval`](#retval) to 0. If [`optlen`](#optlen) has been increased above initial [`getsockopt`](https://linux.die.net/man/2/getsockopt) value (i.e. userspace buffer is too small), `EFAULT` is returned.

This hook has access to the cGroup and socket local storage.

!!! note
    The only acceptable value to set to [`retval`](#retval) is 0 and the original value that the kernel returned. Any other value will trigger `EFAULT`.

Return Type:

* `0` - reject the syscall, `EPERM` will be returned to the userspace.
* `1` - success: copy [`optval`](#optval) and [`optlen`](#optlen) to userspace, return[`retval`](#retval) from the syscall (note that this can be overwritten by the BPF program from the parent cGroup).

### cGroup Inheritance

Suppose, there is the following cGroup hierarchy where each cGroup has `BPF_CGROUP_GETSOCKOPT` attached at each level with `BPF_F_ALLOW_MULTI`

```
  A (root, parent)
   \
    B (child)
```

When the application calls [`getsockopt`](https://linux.die.net/man/2/getsockopt) syscall from the cGroup B, the programs are executed from the bottom up: B, A. First program (B) sees the result of kernel's [`getsockopt`](https://linux.die.net/man/2/getsockopt). It can optionally adjust [`optval`](#optval), [`optlen`](#optlen) and reset [`retval`](#retval) to 0. After that control will be passed to the second (A) program which will see the same context as B including any potential modifications.

Same for `BPF_CGROUP_SETSOCKOPT`: if the program is attached to A and B, the trigger order is B, then A. If B does any changes to the input arguments ([`level`](#level), [`optname`](#optname), [`optval`](#optval), [`optlen`](#optlen)), then the next program in the chain (A) will see those changes, *not* the original input [`setsockopt`](https://linux.die.net/man/2/setsockopt) arguments. The potentially modified values will be then passed down to the kernel.

### Large `optval`

When the [`optval`](#optval) is greater than the `PAGE_SIZE`, the BPF program can access only the first `PAGE_SIZE` of that data. So it has to options:

* Set [`optlen`](#optlen) to zero, which indicates that the kernel should use the original buffer from the userspace. Any modifications done by the BPF program to the [`optval`](#optval) are ignored.
* Set [`optlen`](#optlen) to the value less than `PAGE_SIZE`, which indicates that the kernel should use BPF's trimmed [`optval`](#optval).

When the BPF program returns with the [`optlen`](#optlen) greater than `PAGE_SIZE`, the userspace will receive original kernel buffers without any modifications that the BPF program might have applied.

## Context

`struct bpf_sockopt`

??? abstract "C structure"
    ```c
    struct bpf_sockopt {
        __bpf_md_ptr(struct bpf_sock *, sk);
        __bpf_md_ptr(void *, optval);
        __bpf_md_ptr(void *, optval_end);

        __s32	level;
        __s32	optname;
        __s32	optlen;
        __s32	retval;
    };
    ```

### `sk`

Pointer to the socket for which the syscall is invoked.

### `optval`

Pointer to the start of the option value, the end pointer being `optval_end`. The program must perform bounds check with `optval_end` before accessing the memory.

For `BPF_CGROUP_SETSOCKOPT` the opt value contains the option the process wants to set. For `BPF_CGROUP_GETSOCKOPT` the opt value contains the option the syscall returned.

### `optval_end`

This is the end pointer of the option value.

### `level`

This field indicates the socket level for which the syscall is invoked. Values are one of `SOL_*` constants. Typically `SOL_SOCKET`, `SOL_IP`, `SOL_IPV6`, `SOL_TCP`, or `SOL_UDP` unless dealing with more specialized protocols. Only `BPF_CGROUP_SETSOCKOPT` programs are allowed to modify this field.

### `optname`

This field indicates the name of the socket option. Valid options depend on the socket level. More info can be found in the man pages such as [`socket(7)`](https://linux.die.net/man/7/socket), [`ip(7)`](https://linux.die.net/man/7/ip), [`tcp(7)`](https://linux.die.net/man/7/tcp), [`udp(7)`](https://linux.die.net/man/7/udp), etc.
Only `BPF_CGROUP_SETSOCKOPT` programs are allowed to modify this field.

### `optlen`

This field indicates the length of the socket option, which should be smaller or equal to `optval_end - optval`. The program can modify this value to trim the option value. Both `BPF_CGROUP_SETSOCKOPT` and `BPF_CGROUP_GETSOCKOPT` programs are allowed to modify this field.

### `retval`

This field indicates the return value of the syscall. Only `BPF_CGROUP_GETSOCKOPT` programs can read and/or modify this value to override the return value of the syscall.

## Attachment

cGroup socket buffer programs are attached to cGroups via the [`BPF_PROG_ATTACH`](../syscall/BPF_PROG_ATTACH.md) syscall or via [BPF link](../syscall/BPF_LINK_CREATE.md).

## Example

```c
SEC("cgroup/getsockopt")
int getsockopt(struct bpf_sockopt *ctx)
{
    /* Custom socket option. */
    if (ctx->level == MY_SOL && ctx->optname == MY_OPTNAME) {
        ctx->retval = 0;
        optval[0] = ...;
        ctx->optlen = 1;
        return 1;
    }

    /* Modify kernel's socket option. */
    if (ctx->level == SOL_IP && ctx->optname == IP_FREEBIND) {
        ctx->retval = 0;
        optval[0] = ...;
        ctx->optlen = 1;
        return 1;
    }

    /* optval larger than PAGE_SIZE use kernel's buffer. */
    if (ctx->optlen > PAGE_SIZE)
        ctx->optlen = 0;

    return 1;
}

SEC("cgroup/setsockopt")
int setsockopt(struct bpf_sockopt *ctx)
{
    /* Custom socket option. */
    if (ctx->level == MY_SOL && ctx->optname == MY_OPTNAME) {
        /* do something */
        ctx->optlen = -1;
        return 1;
    }

    /* Modify kernel's socket option. */
    if (ctx->level == SOL_IP && ctx->optname == IP_FREEBIND) {
        optval[0] = ...;
        return 1;
    }

    /* optval larger than PAGE_SIZE use kernel's buffer. */
    if (ctx->optlen > PAGE_SIZE)
        ctx->optlen = 0;

    return 1;
}
```

## Helper functions

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [`bpf_cgrp_storage_delete`](../helper-function/bpf_cgrp_storage_delete.md)
    * [`bpf_cgrp_storage_get`](../helper-function/bpf_cgrp_storage_get.md)
    * [`bpf_dynptr_data`](../helper-function/bpf_dynptr_data.md)
    * [`bpf_dynptr_from_mem`](../helper-function/bpf_dynptr_from_mem.md)
    * [`bpf_dynptr_read`](../helper-function/bpf_dynptr_read.md)
    * [`bpf_dynptr_write`](../helper-function/bpf_dynptr_write.md)
    * [`bpf_for_each_map_elem`](../helper-function/bpf_for_each_map_elem.md)
    * [`bpf_get_current_cgroup_id`](../helper-function/bpf_get_current_cgroup_id.md)
    * [`bpf_get_current_pid_tgid`](../helper-function/bpf_get_current_pid_tgid.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
    * [`bpf_get_current_task`](../helper-function/bpf_get_current_task.md)
    * [`bpf_get_current_task_btf`](../helper-function/bpf_get_current_task_btf.md)
    * [`bpf_get_current_uid_gid`](../helper-function/bpf_get_current_uid_gid.md)
    * [`bpf_get_local_storage`](../helper-function/bpf_get_local_storage.md)
    * [`bpf_get_netns_cookie`](../helper-function/bpf_get_netns_cookie.md)
    * [`bpf_get_ns_current_pid_tgid`](../helper-function/bpf_get_ns_current_pid_tgid.md) [:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/eb166e522c77699fc19bfa705652327a1e51a117)
    * [`bpf_get_numa_node_id`](../helper-function/bpf_get_numa_node_id.md)
    * [`bpf_get_prandom_u32`](../helper-function/bpf_get_prandom_u32.md)
    * [`bpf_get_retval`](../helper-function/bpf_get_retval.md)
    * [`bpf_get_smp_processor_id`](../helper-function/bpf_get_smp_processor_id.md)
    * [`bpf_getsockopt`](../helper-function/bpf_getsockopt.md) [:octicons-tag-24: v5.15](https://github.com/torvalds/linux/commit/2c531639deb5e3ddfd6e8123b82052b2d9fbc6e5)
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
    * [`bpf_set_retval`](../helper-function/bpf_set_retval.md)
    * [`bpf_setsockopt`](../helper-function/bpf_setsockopt.md) [:octicons-tag-24: v5.15](https://github.com/torvalds/linux/commit/2c531639deb5e3ddfd6e8123b82052b2d9fbc6e5)
    * [`bpf_sk_storage_delete`](../helper-function/bpf_sk_storage_delete.md)
    * [`bpf_sk_storage_get`](../helper-function/bpf_sk_storage_get.md)
    * [`bpf_snprintf`](../helper-function/bpf_snprintf.md)
    * [`bpf_snprintf_btf`](../helper-function/bpf_snprintf_btf.md)
    * [`bpf_spin_lock`](../helper-function/bpf_spin_lock.md)
    * [`bpf_spin_unlock`](../helper-function/bpf_spin_unlock.md)
    * [`bpf_strncmp`](../helper-function/bpf_strncmp.md)
    * [`bpf_tail_call`](../helper-function/bpf_tail_call.md)
    * [`bpf_task_pt_regs`](../helper-function/bpf_task_pt_regs.md)
    * [`bpf_tcp_sock`](../helper-function/bpf_tcp_sock.md)
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
    - [`bpf_arena_alloc_pages`](../kfuncs/bpf_arena_alloc_pages.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_arena_free_pages`](../kfuncs/bpf_arena_free_pages.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_cast_to_kern_ctx`](../kfuncs/bpf_cast_to_kern_ctx.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_copy_from_user_str`](../kfuncs/bpf_copy_from_user_str.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_dynptr_adjust`](../kfuncs/bpf_dynptr_adjust.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_dynptr_clone`](../kfuncs/bpf_dynptr_clone.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_dynptr_from_skb`](../kfuncs/bpf_dynptr_from_skb.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_dynptr_is_null`](../kfuncs/bpf_dynptr_is_null.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_dynptr_is_rdonly`](../kfuncs/bpf_dynptr_is_rdonly.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_dynptr_size`](../kfuncs/bpf_dynptr_size.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_dynptr_slice`](../kfuncs/bpf_dynptr_slice.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_dynptr_slice_rdwr`](../kfuncs/bpf_dynptr_slice_rdwr.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_get_kmem_cache`](../kfuncs/bpf_get_kmem_cache.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_bits_destroy`](../kfuncs/bpf_iter_bits_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_bits_new`](../kfuncs/bpf_iter_bits_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_bits_next`](../kfuncs/bpf_iter_bits_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_css_destroy`](../kfuncs/bpf_iter_css_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_css_new`](../kfuncs/bpf_iter_css_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_css_next`](../kfuncs/bpf_iter_css_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_css_task_destroy`](../kfuncs/bpf_iter_css_task_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_css_task_new`](../kfuncs/bpf_iter_css_task_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_css_task_next`](../kfuncs/bpf_iter_css_task_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_kmem_cache_destroy`](../kfuncs/bpf_iter_kmem_cache_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_kmem_cache_new`](../kfuncs/bpf_iter_kmem_cache_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_kmem_cache_next`](../kfuncs/bpf_iter_kmem_cache_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_num_destroy`](../kfuncs/bpf_iter_num_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_num_new`](../kfuncs/bpf_iter_num_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_num_next`](../kfuncs/bpf_iter_num_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_task_destroy`](../kfuncs/bpf_iter_task_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_task_new`](../kfuncs/bpf_iter_task_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_task_next`](../kfuncs/bpf_iter_task_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_task_vma_destroy`](../kfuncs/bpf_iter_task_vma_destroy.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_task_vma_new`](../kfuncs/bpf_iter_task_vma_new.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_iter_task_vma_next`](../kfuncs/bpf_iter_task_vma_next.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_local_irq_restore`](../kfuncs/bpf_local_irq_restore.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_local_irq_save`](../kfuncs/bpf_local_irq_save.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_map_sum_elem_count`](../kfuncs/bpf_map_sum_elem_count.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_preempt_disable`](../kfuncs/bpf_preempt_disable.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_preempt_enable`](../kfuncs/bpf_preempt_enable.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_rcu_read_lock`](../kfuncs/bpf_rcu_read_lock.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_rcu_read_unlock`](../kfuncs/bpf_rcu_read_unlock.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_rdonly_cast`](../kfuncs/bpf_rdonly_cast.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_sock_addr_set_sun_path`](../kfuncs/bpf_sock_addr_set_sun_path.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_wq_init`](../kfuncs/bpf_wq_init.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_wq_set_callback_impl`](../kfuncs/bpf_wq_set_callback_impl.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
    - [`bpf_wq_start`](../kfuncs/bpf_wq_start.md) [:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/67666479edf1e2b732f4d0ac797885e859a78de4) - 
<!-- [/PROG_KFUNC_REF] -->
