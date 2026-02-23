---
title: "Syscall command 'BPF_PROG_ATTACH'"
description: "This page documents the 'BPF_PROG_ATTACH' eBPF syscall command, including its definition, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_PROG_ATTACH` command

<!-- [FEATURE_TAG](BPF_PROG_ATTACH) -->
[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/f4324551489e8781d838f941b7aee4208e52e8bf)
<!-- [/FEATURE_TAG] -->

This syscall command is an outdated method of attaching select program types. Where possible, [BPF links](BPF_LINK_CREATE.md) should be used instead.

## Return value

This command will return a zero on success or an error number (negative integer) if something went wrong.

## Attributes

### `target_fd`

The file descriptor to attach the program to. The type of file descriptor changes per program type.

### `target_ifindex`

The network interface index of the network device to attach the program to.

### `attach_bpf_fd`

The file descriptor of the BPF program to attach to the `target_fd`/`target_ifindex`.

### `attach_type`

The attach type of the program attachment. Used to for example specify if a program should be installed on a networks ingress or egress path.

### `attach_flags`

Any flags relevant to attaching.

### `replace_bpf_fd`

The file descriptor of the BPF program to replace with the program specified at `attach_bpf_fd`. Can for some program types be used to replace an explicitly specified program to avoid accidents where the wrong program may be replaced.

### `relative_fd`

The file descriptor of another program to attach relative to, in the case of attach points that support multiple programs being attached at the same time.

### `relative_id`

The BPF ID of another program to attach relative to, in the case of attach points that support multiple programs being attached at the same time.

### `expected_revision`

The expected revision of the collection of programs,  in the case of attach points that support multiple programs being attached at the same time.

## Usage

This syscall command was an early attempt at allowing programs to be attached via a BPF syscall, instead of needing to use external systems such as [`ioctl`](https://man7.org/linux/man-pages/man2/ioctl.2.html) and [netlink](https://man7.org/linux/man-pages/man7/netlink.7.html) to attach BPF programs.

For the most part it has been succeeded by [BPF links](BPF_LINK_CREATE.md) are the current state of the art when it comes to attaching programs. But there are still some program types that have not received link support yet.

Below are section for each program type that can still be attached with this syscall.

### `BPF_PROG_TYPE_SK_SKB` and `BPF_PROG_TYPE_SK_MSG`

[`BPF_PROG_TYPE_SK_SKB`](../program-type/BPF_PROG_TYPE_SK_SKB.md) and [`BPF_PROG_TYPE_SK_MSG`](../program-type/BPF_PROG_TYPE_SK_MSG.md) programs are attached to a BPF socket map (map of type `BPF_MAP_TYPE_SOCKMAP` or `BPF_MAP_TYPE_SOCKHASH`). As of v6.19 there exists no link support for these program types.

`attach_bpf_fd` should be the file descriptor of a BPF socket map.

No `attach_flags` exist for these programs.

### `BPF_PROG_TYPE_LIRC_MODE2`

[`BPF_PROG_TYPE_LIRC_MODE2`](../program-type/BPF_PROG_TYPE_LIRC_MODE2.md) programs are attached to LIRC devices. As of v6.19 there exists no link support for this program types.

`attach_bpf_fd` should be the file descriptor of a LIRC device (obtained by opening a device in the `/dev` pseudo file system).

No `attach_flags` exist for this program type.

### `BPF_PROG_TYPE_FLOW_DISSECTOR`

[`BPF_PROG_TYPE_FLOW_DISSECTOR`](../program-type/BPF_PROG_TYPE_FLOW_DISSECTOR.md) programs attach to a network namespace. Flow dissectors can be attached via [BPF links](BPF_LINK_CREATE.md) which is the preferred method. But doing it via `BPF_PROG_ATTACH` is still possible for compatibility reasons.

When called, the program is attached to the network namespace of the calling process / thread. 

`attach_bpf_fd` should be zero.

No `attach_flags` exist for this program type.


### `BPF_PROG_TYPE_SCHED_CLS`

[`BPF_PROG_TYPE_SCHED_CLS`](../program-type/BPF_PROG_TYPE_SCHED_CLS.md) (Traffic Control) programs attach to network interfaces. Traffic control programs can be attached via [BPF links](BPF_LINK_CREATE.md) which is the preferred method. But doing it via `BPF_PROG_ATTACH` is still possible for compatibility reasons.

`attach_ifindex` should contain the index of the network interface the program should be attached to.

If `attach_flags` contains `BPF_F_REPLACE`, then the new program will replace the program specified with `replace_bpf_fd`.

If `attach_flags` contains `BPF_F_BEFORE`, then the new program will be inserted before `relative_fd` / `relative_id`.

If `attach_flags` contains `BPF_F_AFTER`, then the new program will be inserted after `relative_fd` / `relative_id`.

`expected_revision` must be equal to the current revision which can be obtained via the [`BPF_PROG_QUERY`](BPF_PROG_QUERY.md) syscall.
