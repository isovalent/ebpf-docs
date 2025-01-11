---
title: "Syscall command 'BPF_LINK_CREATE'"
description: "This page documents the 'BPF_LINK_CREATE' eBPF syscall command, including its definition, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_LINK_CREATE` command

<!-- [FEATURE_TAG](BPF_LINK_CREATE) -->
[:octicons-tag-24: v5.7](https://github.com/torvalds/linux/commit/af6eea57437a830293eab56246b6025cc7d46ee7)
<!-- [/FEATURE_TAG] -->

This syscall command create a new BPF link. BPF links are the newest and thus preferred way to attach BPF programs to their hook locations within the kernel. This syscall command is intended to replace both `BPF_PROG_ATTACH` and other legacy attachment methods such as netlink and perf events.

## Return type

If successful, this syscall command will return a file descriptor to the newly created link. The link is a reference counted object just like other BPF objects. The link is destroyed once no more references to it exist, which might happen if the loader exits without pinning the link or if the pin gets deleted. A loader might also chose to forcefully cause the link to detach from the hook point with the [`BPF_LINK_DETACH`](BPF_LINK_DETACH.md) command.

The returned file descriptor can be used with the [`BPF_LINK_UPDATE`](BPF_LINK_UPDATE.md) and [`BPF_LINK_DETACH`](BPF_LINK_DETACH.md) commands.

## Attributes

??? abstract "C structure"
    ```c
    struct { /* struct used by BPF_LINK_CREATE command */
		__u32		prog_fd;	/* eBPF program to attach */
		union {
			__u32		target_fd;	/* object to attach to */
			__u32		target_ifindex; /* target ifindex */
		};
		__u32		attach_type;	/* attach type */
		__u32		flags;		/* extra flags */
		union {
			__u32		target_btf_id;	/* btf_id of target to attach to */
			struct {
				__aligned_u64	iter_info;	/* extra bpf_iter_link_info */
				__u32		iter_info_len;	/* iter_info length */
			};
			struct {
				/* black box user-provided value passed through
				 * to BPF program at the execution time and
				 * accessible through bpf_get_attach_cookie() BPF helper
				 */
				__u64		bpf_cookie;
			} perf_event;
			struct {
				__u32		flags;
				__u32		cnt;
				__aligned_u64	syms;
				__aligned_u64	addrs;
				__aligned_u64	cookies;
			} kprobe_multi;
			struct {
				/* this is overlaid with the target_btf_id above. */
				__u32		target_btf_id;
				/* black box user-provided value passed through
				 * to BPF program at the execution time and
				 * accessible through bpf_get_attach_cookie() BPF helper
				 */
				__u64		cookie;
			} tracing;
		};
	}
    ```

### `prog_fd`

[:octicons-tag-24: v5.7](https://github.com/torvalds/linux/commit/af6eea57437a830293eab56246b6025cc7d46ee7)

This field specifies the file descriptor for the BPF program to be linked.

### `target_fd`

[:octicons-tag-24: v5.9](https://github.com/torvalds/linux/commit/aa8d3a716b59db6c1ad6c68fb8aa05e31980da60)

This field specifies the file descriptor of the target you wish to attach the program to. The kind of file descriptor varies per program type.

For cGroup programs (`BPF_PROG_TYPE_CGROUP_SKB`, `BPF_PROG_TYPE_CGROUP_SOCK`,
`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`,`BPF_PROG_TYPE_SOCK_OPS`,`BPF_PROG_TYPE_CGROUP_DEVICE`,`BPF_PROG_TYPE_CGROUP_SYSCTL`,`BPF_PROG_TYPE_CGROUP_SOCKOPT`) the file descriptor should be a cGroup directory in the cGroup FS, commonly mounted at `/sys/fs/cgroup`, for example `/sys/fs/cgroup/test.slice/`. Such a fd can be obtained by using the [`open`](https://man7.org/linux/man-pages/man2/open.2.html) syscall on the desired path.

For `BPF_PROG_TYPE_EXT` programs this should be a file descriptor to another BPF program.

For `BPF_PROG_TYPE_TRACING` programs with the attach type `BPF_LSM_CGROUP` it should also be a cGroup directory as described above.

For `BPF_PROG_TYPE_TRACING` programs with the attach type `BPF_TRACE_FENTRY`, `BPF_TRACE_FEXIT` or `BPF_MODIFY_RETURN` this should be a file descriptor to an existing BPF program.

For `BPF_PROG_TYPE_LIRC2` programs this should be a file descriptor to a infrared device in `/dev`.

For `BPF_PROG_TYPE_FLOW_DISSECTOR` and `BPF_PROG_TYPE_SK_LOOKUP` programs this should be a file descriptor to a network namespace. Named network namespaces are represented as objects in `/var/run/netns`, a file descriptor to a namespace can be obtained [`open`](https://man7.org/linux/man-pages/man2/open.2.html) syscall on one of these objects (`/var/run/netns/{name}`).

### `target_ifindex`

[:octicons-tag-24: v5.9](https://github.com/torvalds/linux/commit/aa8d3a716b59db6c1ad6c68fb8aa05e31980da60)

This field specifies the network interface index of the network device to attach the program to. This field is only used for `BPF_PROG_TYPE_XDP` programs.

### `attach_type`

[:octicons-tag-24: v5.7](https://github.com/torvalds/linux/commit/af6eea57437a830293eab56246b6025cc7d46ee7)

Attach type specifies the attach type. For more information about possible values and their meaning checkout the [Attach types](#attach-types) section.

### `flags`

[:octicons-tag-24: v5.7](https://github.com/torvalds/linux/commit/af6eea57437a830293eab56246b6025cc7d46ee7)

<!-- TODO figure out -->

### `target_btf_id`

[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/4a1e7c0c63e02daad751842b7880f9bbcdfb6e89)

This field specifies the BTF id of the target to attach to, used to specify the kernel function to hook to when attaching `BPF_PROG_TYPE_LSM` programs.

### `iter_info`

[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/4a1e7c0c63e02daad751842b7880f9bbcdfb6e89)

This field specifies over what kind of information a iterator program should iterate. It is a pointer to an instance of `union bpf_iter_link_info`

??? abstract "`union bpf_iter_link_info` C structure"
	```c
	union bpf_iter_link_info {
		struct {
			__u32	map_fd;
		} map;
		struct {
			enum bpf_cgroup_iter_order order;

			/* At most one of cgroup_fd and cgroup_id can be non-zero. If
			* both are zero, the walk starts from the default cgroup v2
			* root. For walking v1 hierarchy, one should always explicitly
			* specify cgroup_fd.
			*/
			__u32	cgroup_fd;
			__u64	cgroup_id;
		} cgroup;
		/* Parameters of task iterators. */
		struct {
			__u32	tid;
			__u32	pid;
			__u32	pid_fd;
		} task;
	};
	```

### `iter_info_len`

[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/4a1e7c0c63e02daad751842b7880f9bbcdfb6e89)

This field specifies the length of the given `iter_info` structure, for the purposes of compatibility in case new kernels add additional fields.

### `bpf_cookie`

[:octicons-tag-24: v5.15](https://github.com/torvalds/linux/commit/82e6b1eee6a8875ef4eacfd60711cce6965c6b04)

This field is an optional opaque value which is reported back to tracing programs via the [`bpf_get_attach_cookie`](../helper-function/bpf_get_attach_cookie.md) helper.

The idea behind this cookie is that if the same program gets attached to multiple locations in the kernel, this value can be used to distinguish for which attach point the program is called. This ID value can for example be used as the key for a map which contains additional data the program needs or as key when collecting statistics.

### `kprobe_multi`

[:octicons-tag-24: v5.18](https://github.com/torvalds/linux/commit/0dcac272540613d41c05e89679e4ddb978b612f1)

This sub-struct is a collection of fields which specify one or multiple kprobe attachment points to attach the same program to multiple locations with a single syscall.

#### `flags`

[:octicons-tag-24: v5.18](https://github.com/torvalds/linux/commit/0dcac272540613d41c05e89679e4ddb978b612f1)

<!-- TODO figure out -->

#### `cnt`

[:octicons-tag-24: v5.18](https://github.com/torvalds/linux/commit/0dcac272540613d41c05e89679e4ddb978b612f1)

This field is the number of `syms`, `addrs` and `cookies` to follow

#### `syms`

[:octicons-tag-24: v5.18](https://github.com/torvalds/linux/commit/0dcac272540613d41c05e89679e4ddb978b612f1)

This field specifies a list of kernel symbols to attach the kprobe to. The value should be a pointer to an array of null-terminated string pointers.

`#!c [cnt][]char`

This field is mutually exclusive with `addrs`. Only one can be used at a time.

#### `addrs`

[:octicons-tag-24: v5.18](https://github.com/torvalds/linux/commit/0dcac272540613d41c05e89679e4ddb978b612f1)

This field specifies a list of kernel addresses to attach the kprobe to. The value should be a pointer to an array of memory addresses.

This field is mutually exclusive with `syms`. Only one can be used at a time.

#### `cookies`

[:octicons-tag-24: v5.18](https://github.com/torvalds/linux/commit/ca74823c6e16dd42b7cf60d9fdde80e2a81a67bb)

This field specifies a list of cookies([`bpf_cookie`](#bpf_cookie)) values for each attachment point. The value should be a pointer to an array of 64-bit cookie values or `0` if you do not want to specify cookies.

### `tracing`

[:octicons-tag-24: v5.19](https://github.com/torvalds/linux/commit/2fcc82411e74e5e6aba336561cf56fb899bfae4e)

#### `target_btf_id`

[:octicons-tag-24: v5.19](https://github.com/torvalds/linux/commit/2fcc82411e74e5e6aba336561cf56fb899bfae4e)

#### `cookie`

[:octicons-tag-24: v5.19](https://github.com/torvalds/linux/commit/2fcc82411e74e5e6aba336561cf56fb899bfae4e)

## Attach types

This section describes the possible values and meanings for the `attach_type` attribute. These values are the same as used in the [`BPF_PROG_ATTACH`](BPF_PROG_ATTACH.md) command and the [`expected_attach_type`](BPF_PROG_LOAD.md#expected_attach_type) field of the [`BPF_PROG_LOAD`](BPF_PROG_LOAD.md) command.

The attach type is often used to communicate a specialization for a program type, for example if the program should attach to the ingress or egress. Since the hook locations will differ, the capabilities of the program may as well. Please check the pages of the program types for details about these attach type dependant limitations.

### `BPF_CGROUP_INET_INGRESS`

<!-- [FEATURE_TAG](BPF_CGROUP_INET_INGRESS) -->
[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/0e33661de493db325435d565a4a722120ae4cbf3)
<!-- [/FEATURE_TAG] -->

### `BPF_CGROUP_INET_EGRESS`

<!-- [FEATURE_TAG](BPF_CGROUP_INET_EGRESS) -->
[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/0e33661de493db325435d565a4a722120ae4cbf3)
<!-- [/FEATURE_TAG] -->


### `BPF_CGROUP_INET_SOCK_CREATE`

<!-- [FEATURE_TAG](BPF_CGROUP_INET_SOCK_CREATE) -->
[:octicons-tag-24: v4.10](https://github.com/torvalds/linux/commit/61023658760032e97869b07d54be9681d2529e77)
<!-- [/FEATURE_TAG] -->


### `BPF_CGROUP_SOCK_OPS`

<!-- [FEATURE_TAG](BPF_CGROUP_SOCK_OPS) -->
[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/40304b2a1567fecc321f640ee4239556dd0f3ee0)
<!-- [/FEATURE_TAG] -->


### `BPF_SK_SKB_STREAM_PARSER`

<!-- [FEATURE_TAG](BPF_SK_SKB_STREAM_PARSER) -->
[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/464bc0fd6273d518aee79fbd37211dd9bc35d863)
<!-- [/FEATURE_TAG] -->


### `BPF_SK_SKB_STREAM_VERDICT`

<!-- [FEATURE_TAG](BPF_SK_SKB_STREAM_VERDICT) -->
[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/464bc0fd6273d518aee79fbd37211dd9bc35d863)
<!-- [/FEATURE_TAG] -->


### `BPF_CGROUP_DEVICE`

<!-- [FEATURE_TAG](BPF_CGROUP_DEVICE) -->
[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/ebc614f687369f9df99828572b1d85a7c2de3d92)
<!-- [/FEATURE_TAG] -->


### `BPF_SK_MSG_VERDICT`

<!-- [FEATURE_TAG](BPF_SK_MSG_VERDICT) -->
[:octicons-tag-24: v4.17](https://github.com/torvalds/linux/commit/4f738adba30a7cfc006f605707e7aee847ffefa0)
<!-- [/FEATURE_TAG] -->


### `BPF_CGROUP_INET4_BIND`

<!-- [FEATURE_TAG](BPF_CGROUP_INET4_BIND) -->
[:octicons-tag-24: v4.17](https://github.com/torvalds/linux/commit/4fbac77d2d092b475dda9eea66da674369665427)
<!-- [/FEATURE_TAG] -->


### `BPF_CGROUP_INET6_BIND`

<!-- [FEATURE_TAG](BPF_CGROUP_INET6_BIND) -->
[:octicons-tag-24: v4.17](https://github.com/torvalds/linux/commit/4fbac77d2d092b475dda9eea66da674369665427)
<!-- [/FEATURE_TAG] -->


### `BPF_CGROUP_INET4_CONNECT`

<!-- [FEATURE_TAG](BPF_CGROUP_INET4_CONNECT) -->
[:octicons-tag-24: v4.17](https://github.com/torvalds/linux/commit/d74bad4e74ee373787a9ae24197c17b7cdc428d5)
<!-- [/FEATURE_TAG] -->


### `BPF_CGROUP_INET6_CONNECT`

<!-- [FEATURE_TAG](BPF_CGROUP_INET6_CONNECT) -->
[:octicons-tag-24: v4.17](https://github.com/torvalds/linux/commit/d74bad4e74ee373787a9ae24197c17b7cdc428d5)
<!-- [/FEATURE_TAG] -->


### `BPF_CGROUP_INET4_POST_BIND`

<!-- [FEATURE_TAG](BPF_CGROUP_INET4_POST_BIND) -->
[:octicons-tag-24: v4.17](https://github.com/torvalds/linux/commit/aac3fc320d9404f2665a8b1249dc3170d5fa3caf)
<!-- [/FEATURE_TAG] -->


### `BPF_CGROUP_INET6_POST_BIND`

<!-- [FEATURE_TAG](BPF_CGROUP_INET6_POST_BIND) -->
[:octicons-tag-24: v4.17](https://github.com/torvalds/linux/commit/aac3fc320d9404f2665a8b1249dc3170d5fa3caf)
<!-- [/FEATURE_TAG] -->


### `BPF_CGROUP_UDP4_SENDMSG`

<!-- [FEATURE_TAG](BPF_CGROUP_UDP4_SENDMSG) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/1cedee13d25ab118d325f95588c1a084e9317229)
<!-- [/FEATURE_TAG] -->


### `BPF_CGROUP_UDP6_SENDMSG`

<!-- [FEATURE_TAG](BPF_CGROUP_UDP6_SENDMSG) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/1cedee13d25ab118d325f95588c1a084e9317229)
<!-- [/FEATURE_TAG] -->


### `BPF_LIRC_MODE2`

<!-- [FEATURE_TAG](BPF_LIRC_MODE2) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/f4364dcfc86df7c1ca47b256eaf6b6d0cdd0d936)
<!-- [/FEATURE_TAG] -->


### `BPF_FLOW_DISSECTOR`

<!-- [FEATURE_TAG](BPF_FLOW_DISSECTOR) -->
[:octicons-tag-24: v4.20](https://github.com/torvalds/linux/commit/d58e468b1112dcd1d5193c0a89ff9f98b5a3e8b9)
<!-- [/FEATURE_TAG] -->


### `BPF_CGROUP_SYSCTL`

<!-- [FEATURE_TAG](BPF_CGROUP_SYSCTL) -->
[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/7b146cebe30cb481b0f70d85779da938da818637)
<!-- [/FEATURE_TAG] -->


### `BPF_CGROUP_UDP4_RECVMSG`

<!-- [FEATURE_TAG](BPF_CGROUP_UDP4_RECVMSG) -->
[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/983695fa676568fc0fe5ddd995c7267aabc24632)
<!-- [/FEATURE_TAG] -->


### `BPF_CGROUP_UDP6_RECVMSG`

<!-- [FEATURE_TAG](BPF_CGROUP_UDP6_RECVMSG) -->
[:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/983695fa676568fc0fe5ddd995c7267aabc24632)
<!-- [/FEATURE_TAG] -->


### `BPF_CGROUP_GETSOCKOPT`

<!-- [FEATURE_TAG](BPF_CGROUP_GETSOCKOPT) -->
[:octicons-tag-24: v5.3](https://github.com/torvalds/linux/commit/0d01da6afc5402f60325c5da31b22f7d56689b49)
<!-- [/FEATURE_TAG] -->


### `BPF_CGROUP_SETSOCKOPT`

<!-- [FEATURE_TAG](BPF_CGROUP_SETSOCKOPT) -->
[:octicons-tag-24: v5.3](https://github.com/torvalds/linux/commit/0d01da6afc5402f60325c5da31b22f7d56689b49)
<!-- [/FEATURE_TAG] -->


### `BPF_TRACE_RAW_TP`

<!-- [FEATURE_TAG](BPF_TRACE_RAW_TP) -->
[:octicons-tag-24: v5.5](https://github.com/torvalds/linux/commit/f1b9509c2fb0ef4db8d22dac9aef8e856a5d81f6)
<!-- [/FEATURE_TAG] -->


### `BPF_TRACE_FENTRY`

<!-- [FEATURE_TAG](BPF_TRACE_FENTRY) -->
[:octicons-tag-24: v5.5](https://github.com/torvalds/linux/commit/fec56f5890d93fc2ed74166c397dc186b1c25951)
<!-- [/FEATURE_TAG] -->


### `BPF_TRACE_FEXIT`

<!-- [FEATURE_TAG](BPF_TRACE_FEXIT) -->
[:octicons-tag-24: v5.5](https://github.com/torvalds/linux/commit/fec56f5890d93fc2ed74166c397dc186b1c25951)
<!-- [/FEATURE_TAG] -->


### `BPF_MODIFY_RETURN`

<!-- [FEATURE_TAG](BPF_MODIFY_RETURN) -->
[:octicons-tag-24: v5.7](https://github.com/torvalds/linux/commit/ae24082331d9bbaae283aafbe930a8f0eb85605a)
<!-- [/FEATURE_TAG] -->


### `BPF_LSM_MAC`

<!-- [FEATURE_TAG](BPF_LSM_MAC) -->
[:octicons-tag-24: v5.7](https://github.com/torvalds/linux/commit/fc611f47f2188ade2b48ff6902d5cce8baac0c58)
<!-- [/FEATURE_TAG] -->


### `BPF_TRACE_ITER`

<!-- [FEATURE_TAG](BPF_TRACE_ITER) -->
[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/15d83c4d7cef5c067a8b075ce59e97df4f60706e)
<!-- [/FEATURE_TAG] -->


### `BPF_CGROUP_INET4_GETPEERNAME`

<!-- [FEATURE_TAG](BPF_CGROUP_INET4_GETPEERNAME) -->
[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/1b66d253610c7f8f257103808a9460223a087469)
<!-- [/FEATURE_TAG] -->


### `BPF_CGROUP_INET6_GETPEERNAME`

<!-- [FEATURE_TAG](BPF_CGROUP_INET6_GETPEERNAME) -->
[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/1b66d253610c7f8f257103808a9460223a087469)
<!-- [/FEATURE_TAG] -->


### `BPF_CGROUP_INET4_GETSOCKNAME`

<!-- [FEATURE_TAG](BPF_CGROUP_INET4_GETSOCKNAME) -->
[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/1b66d253610c7f8f257103808a9460223a087469)
<!-- [/FEATURE_TAG] -->


### `BPF_CGROUP_INET6_GETSOCKNAME`

<!-- [FEATURE_TAG](BPF_CGROUP_INET6_GETSOCKNAME) -->
[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/1b66d253610c7f8f257103808a9460223a087469)
<!-- [/FEATURE_TAG] -->


### `BPF_XDP_DEVMAP`

<!-- [FEATURE_TAG](BPF_XDP_DEVMAP) -->
[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/fbee97feed9b3e4acdf9590e1f6b4a2eefecfffe)
<!-- [/FEATURE_TAG] -->


### `BPF_CGROUP_INET_SOCK_RELEASE`

<!-- [FEATURE_TAG](BPF_CGROUP_INET_SOCK_RELEASE) -->
[:octicons-tag-24: v5.9](https://github.com/torvalds/linux/commit/f5836749c9c04a10decd2742845ad4870965fdef)
<!-- [/FEATURE_TAG] -->


### `BPF_XDP_CPUMAP`

<!-- [FEATURE_TAG](BPF_XDP_CPUMAP) -->
[:octicons-tag-24: v5.9](https://github.com/torvalds/linux/commit/9216477449f33cdbc9c9a99d49f500b7fbb81702)
<!-- [/FEATURE_TAG] -->


### `BPF_SK_LOOKUP`

<!-- [FEATURE_TAG](BPF_SK_LOOKUP) -->
[:octicons-tag-24: v5.9](https://github.com/torvalds/linux/commit/e9ddbb7707ff5891616240026062b8c1e29864ca)
<!-- [/FEATURE_TAG] -->


### `BPF_XDP`

<!-- [FEATURE_TAG](BPF_XDP) -->
[:octicons-tag-24: v5.9](https://github.com/torvalds/linux/commit/aa8d3a716b59db6c1ad6c68fb8aa05e31980da60)
<!-- [/FEATURE_TAG] -->


### `BPF_SK_SKB_VERDICT`

<!-- [FEATURE_TAG](BPF_SK_SKB_VERDICT) -->
[:octicons-tag-24: v5.13](https://github.com/torvalds/linux/commit/a7ba4558e69a3c2ae4ca521f015832ef44799538)
<!-- [/FEATURE_TAG] -->


### `BPF_SK_REUSEPORT_SELECT`

<!-- [FEATURE_TAG](BPF_SK_REUSEPORT_SELECT) -->
[:octicons-tag-24: v5.14](https://github.com/torvalds/linux/commit/d5e4ddaeb6ab2c3c7fbb7b247a6d34bb0b18d87e)
<!-- [/FEATURE_TAG] -->


### `BPF_SK_REUSEPORT_SELECT_OR_MIGRATE`

<!-- [FEATURE_TAG](BPF_SK_REUSEPORT_SELECT_OR_MIGRATE) -->
[:octicons-tag-24: v5.14](https://github.com/torvalds/linux/commit/d5e4ddaeb6ab2c3c7fbb7b247a6d34bb0b18d87e)
<!-- [/FEATURE_TAG] -->


### `BPF_PERF_EVENT`

<!-- [FEATURE_TAG](BPF_PERF_EVENT) -->
[:octicons-tag-24: v5.15](https://github.com/torvalds/linux/commit/b89fbfbb854c9afc3047e8273cc3a694650b802e)
<!-- [/FEATURE_TAG] -->


### `BPF_TRACE_KPROBE_MULTI`

<!-- [FEATURE_TAG](BPF_TRACE_KPROBE_MULTI) -->
[:octicons-tag-24: v5.18](https://github.com/torvalds/linux/commit/0dcac272540613d41c05e89679e4ddb978b612f1)
<!-- [/FEATURE_TAG] -->


### `BPF_LSM_CGROUP`

<!-- [FEATURE_TAG](BPF_LSM_CGROUP) -->
[:octicons-tag-24: v6.0](https://github.com/torvalds/linux/commit/69fd337a975c7e690dfe49d9cb4fe5ba1e6db44e)
<!-- [/FEATURE_TAG] -->


## Flags

<!-- TODO does this command have flags? -->
