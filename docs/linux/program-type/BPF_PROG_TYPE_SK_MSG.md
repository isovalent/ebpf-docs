---
title: "Program Type 'BPF_PROG_TYPE_SK_MSG' - eBPF Docs"
description: "This page documents the 'BPF_PROG_TYPE_SK_MSG' eBPF program type, including its defintion, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_SK_MSG`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_SK_MSG) -->
[:octicons-tag-24: v4.17](https://github.com/torvalds/linux/commit/4f738adba30a7cfc006f605707e7aee847ffefa0)
<!-- [/FEATURE_TAG] -->

Socket message programs are called for every `sendmsg` or `sendfile` syscall. This program type can pass verdict on individual packets or larger L7 messages chunked over multiple syscalls.

## Usage

Socket MSG programs are attached to [`BPF_MAP_TYPE_SOCKMAP`](../map-type/BPF_MAP_TYPE_SOCKMAP.md) or [`BPF_MAP_TYPE_SOCKHASH`](../map-type/BPF_MAP_TYPE_SOCKHASH.md) maps and will be invoked `sendmsg` or `sendfile` syscalls are executed on sockets which are part of the map the program is attached to.

The program returns a verdict on what to do with the data the process wants to send.

* `SK_PASS` - The message may pass to the socket or it has been redirected with a helper.
* `SK_DROP` - The message should be dropped.

The [`bpf_msg_apply_bytes`](../helper-function/bpf_msg_apply_bytes.md) helper function can be used to indicate for which bytes the verdict applies. This has two cases. First BPF program applies
verdict to fewer bytes than in the current sendmsg/sendfile this will apply the verdict to the first N bytes of the message then run the BPF program again with data pointers recalculated to the N+1 byte. The second case is the BPF program applies a verdict to more bytes than the current sendmsg or sendfile system call. In this case the infrastructure will cache the verdict and apply it to future sendmsg/sendfile calls until the byte limit is reached. This avoids the overhead of running BPF programs on large payloads.

The helper [`bpf_msg_cork_bytes`](../helper-function/bpf_msg_cork_bytes.md) handles a different case where a BPF program can not reach a verdict on a msg until it receives more bytes AND the program doesn't want to forward the packet until it is known to be "good". The example case being a user (albeit a dumb one probably) sends messages in 1B system calls. The BPF program can call [`bpf_msg_cork_bytes`](../helper-function/bpf_msg_cork_bytes.md) with the required byte limit to reach a verdict and then the program will only be called again once N bytes are received.

## Context

Socket message programs are invoked with a `struct sk_msg_md` context. All field are readable, none are writable.

```c
struct sk_msg_md {
	__bpf_md_ptr(void *, data);
	__bpf_md_ptr(void *, data_end);

	__u32 family;
	__u32 remote_ip4;	/* Stored in network byte order */
	__u32 local_ip4;	/* Stored in network byte order */
	__u32 remote_ip6[4];	/* Stored in network byte order */
	__u32 local_ip6[4];	/* Stored in network byte order */
	__u32 remote_port;	/* Stored in network byte order */
	__u32 local_port;	/* stored in host byte order */
	__u32 size;		/* Total size of sk_msg */

	__bpf_md_ptr(struct bpf_sock *, sk); /* current socket */
};
```

## Attachment

This program type must always be loaded with the [`expected_attach_type`](../syscall/BPF_PROG_LOAD.md#expected_attach_type) of `BPF_SK_MSG_VERDICT`.

Socket message programs are attached to [`BPF_MAP_TYPE_SOCKMAP`](../map-type/BPF_MAP_TYPE_SOCKMAP.md) or [`BPF_MAP_TYPE_SOCKHASH`](../map-type/BPF_MAP_TYPE_SOCKHASH.md) using the [`BPF_PROG_ATTACH`](../syscall/BPF_PROG_ATTACH.md) syscall (`bpf_prog_attach` libbpf function).

## Example

Example of redirecting a message:

```c
// Copyright (c) 2020 Cloudflare

struct {
	__uint(type, BPF_MAP_TYPE_SOCKMAP);
	__uint(max_entries, 2);
	__type(key, __u32);
	__type(value, __u64);
} sock_map SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_SOCKHASH);
	__uint(max_entries, 2);
	__type(key, __u32);
	__type(value, __u64);
} sock_hash SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_ARRAY);
	__uint(max_entries, 2);
	__type(key, int);
	__type(value, unsigned int);
} verdict_map SEC(".maps");

SEC("sk_msg")
int prog_msg_verdict(struct sk_msg_md *msg)
{
	unsigned int *count;
	__u32 zero = 0;
	int verdict;

	if (test_sockmap)
		verdict = bpf_msg_redirect_map(msg, &sock_map, zero, 0);
	else
		verdict = bpf_msg_redirect_hash(msg, &sock_hash, &zero, 0);

	count = bpf_map_lookup_elem(&verdict_map, &verdict);
	if (count)
		(*count)++;

	return verdict;
}
```

Example of dropping based on PID and TPID:

```c
// Copyright (c) 2020 Isovalent, Inc.

struct {
	__uint(type, BPF_MAP_TYPE_SOCKMAP);
	__uint(max_entries, 2);
	__type(key, __u32);
	__type(value, __u64);
} sock_map SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_SOCKHASH);
	__uint(max_entries, 2);
	__type(key, __u32);
	__type(value, __u64);
} sock_hash SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_SK_STORAGE);
	__uint(map_flags, BPF_F_NO_PREALLOC);
	__type(key, __u32);
	__type(value, __u64);
} socket_storage SEC(".maps");

SEC("sk_msg")
int prog_msg_verdict(struct sk_msg_md *msg)
{
	struct task_struct *task = (struct task_struct *)bpf_get_current_task();
	int verdict = SK_PASS;
	__u32 pid, tpid;
	__u64 *sk_stg;

	pid = bpf_get_current_pid_tgid() >> 32;
	sk_stg = bpf_sk_storage_get(&socket_storage, msg->sk, 0, BPF_SK_STORAGE_GET_F_CREATE);
	if (!sk_stg)
		return SK_DROP;
	*sk_stg = pid;
	bpf_probe_read_kernel(&tpid , sizeof(tpid), &task->tgid);
	if (pid != tpid)
		verdict = SK_DROP;
	bpf_sk_storage_delete(&socket_storage, (void *)msg->sk);
	return verdict;
}
```


## Helper functions

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [bpf_msg_redirect_map](../helper-function/bpf_msg_redirect_map.md)
    * [bpf_msg_redirect_hash](../helper-function/bpf_msg_redirect_hash.md)
    * [bpf_msg_apply_bytes](../helper-function/bpf_msg_apply_bytes.md)
    * [bpf_msg_cork_bytes](../helper-function/bpf_msg_cork_bytes.md)
    * [bpf_msg_pull_data](../helper-function/bpf_msg_pull_data.md)
    * [bpf_msg_push_data](../helper-function/bpf_msg_push_data.md)
    * [bpf_msg_pop_data](../helper-function/bpf_msg_pop_data.md)
    * [bpf_perf_event_output](../helper-function/bpf_perf_event_output.md)
    * [bpf_get_current_uid_gid](../helper-function/bpf_get_current_uid_gid.md)
    * [bpf_get_current_pid_tgid](../helper-function/bpf_get_current_pid_tgid.md)
    * [bpf_sk_storage_get](../helper-function/bpf_sk_storage_get.md)
    * [bpf_sk_storage_delete](../helper-function/bpf_sk_storage_delete.md)
    * [bpf_get_netns_cookie](../helper-function/bpf_get_netns_cookie.md)
    * [bpf_get_current_cgroup_id](../helper-function/bpf_get_current_cgroup_id.md)
    * [bpf_get_current_ancestor_cgroup_id](../helper-function/bpf_get_current_ancestor_cgroup_id.md)
    * [bpf_get_cgroup_classid](../helper-function/bpf_get_cgroup_classid.md)
    * [bpf_skc_to_tcp6_sock](../helper-function/bpf_skc_to_tcp6_sock.md)
    * [bpf_skc_to_tcp_sock](../helper-function/bpf_skc_to_tcp_sock.md)
    * [bpf_skc_to_tcp_timewait_sock](../helper-function/bpf_skc_to_tcp_timewait_sock.md)
    * [bpf_skc_to_tcp_request_sock](../helper-function/bpf_skc_to_tcp_request_sock.md)
    * [bpf_skc_to_udp6_sock](../helper-function/bpf_skc_to_udp6_sock.md)
    * [bpf_skc_to_unix_sock](../helper-function/bpf_skc_to_unix_sock.md)
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

