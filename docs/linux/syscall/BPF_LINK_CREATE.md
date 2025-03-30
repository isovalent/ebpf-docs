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

The attributes for this syscall are particularly complex. Ultimately the `attach_type` field will determine which fields are used and how.

```c
union bpf_attr {
	struct {
		union {
			__u32	[prog_fd](#prog_fd);
			__u32	[map_fd](#map_fd);
		};
		union {
			__u32	[target_fd](#target_fd);
			__u32	[target_ifindex](#target_ifindex);
		};
		__u32		[attach_type](#attach_type);
		__u32		[flags](#flags);
		union {
			__u32	[target_btf_id](#target_btf_id);
			struct {
				__aligned_u64	[iter_info](#iter_info);
				__u32			[iter_info_len](#iter_info_len);
			};
			struct {
				[...]
			} [perf_event](#perf_event);
			struct {
				[...]
			} [kprobe_multi](#kprobe_multi);
			struct {
				[...]
			} [tracing](#tracing);
			struct {
				[...]
			} [netfilter](#netfilter);
			struct {
				[...]
			} [tcx](#tcx);
			struct {
				[...]
			} [uprobe_multi](#uprobe_multi);
			struct {
				[...]
			} [netkit](#netkit);
		};
	} link_create;
};
```

### `prog_fd`

[:octicons-tag-24: v5.7](https://github.com/torvalds/linux/commit/af6eea57437a830293eab56246b6025cc7d46ee7)

This field specifies the file descriptor for the BPF program to be linked.

### `map_fd`

[:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/68b04864ca425d1894c96b8141d4fba1181f11cb)

This field specifies a BPF map for the BPF program to be linked to.

### `target_fd`

[:octicons-tag-24: v5.9](https://github.com/torvalds/linux/commit/aa8d3a716b59db6c1ad6c68fb8aa05e31980da60)

This field specifies the file descriptor of the target you wish to attach the program to. The kind of file descriptor varies per program type.

For cGroup programs (`BPF_PROG_TYPE_CGROUP_SKB`, `BPF_PROG_TYPE_CGROUP_SOCK`,
`BPF_PROG_TYPE_CGROUP_SOCK_ADDR`,`BPF_PROG_TYPE_SOCK_OPS`,`BPF_PROG_TYPE_CGROUP_DEVICE`,`BPF_PROG_TYPE_CGROUP_SYSCTL`,`BPF_PROG_TYPE_CGROUP_SOCKOPT`) the file descriptor should be a cGroup directory in the cGroup FS, commonly mounted at `/sys/fs/cgroup`, for example `/sys/fs/cgroup/test.slice/`. Such a fd can be obtained by using the [`open`](https://man7.org/linux/man-pages/man2/open.2.html) syscall on the desired path.

For `BPF_PROG_TYPE_EXT` programs this should be a file descriptor to another BPF program.

For `BPF_PROG_TYPE_LSM` programs with the attach type `BPF_LSM_CGROUP` it should also be a cGroup directory as described above.

For `BPF_PROG_TYPE_TRACING` programs with the attach type `BPF_TRACE_FENTRY`, `BPF_TRACE_FEXIT` or `BPF_MODIFY_RETURN` this should be a file descriptor to an existing BPF program.

For `BPF_PROG_TYPE_LIRC2` programs this should be a file descriptor to a infrared device in `/dev`.

For `BPF_PROG_TYPE_FLOW_DISSECTOR` and `BPF_PROG_TYPE_SK_LOOKUP` programs this should be a file descriptor to a network namespace. Named network namespaces are represented as objects in `/var/run/netns`, a file descriptor to a namespace can be obtained [`open`](https://man7.org/linux/man-pages/man2/open.2.html) syscall on one of these objects (`/var/run/netns/{name}`).

### `target_ifindex`

[:octicons-tag-24: v5.9](https://github.com/torvalds/linux/commit/aa8d3a716b59db6c1ad6c68fb8aa05e31980da60)

This field specifies the network interface index of the network device to attach the program to. This field is only used for `BPF_PROG_TYPE_XDP` programs.

### `attach_type`

[:octicons-tag-24: v5.7](https://github.com/torvalds/linux/commit/af6eea57437a830293eab56246b6025cc7d46ee7)

Attach type specifies the attach type. For more information about possible values and their meaning checkout the [Attach types](#attach-types) section.

### `flags` {attributes-flags}

[:octicons-tag-24: v5.7](https://github.com/torvalds/linux/commit/af6eea57437a830293eab56246b6025cc7d46ee7)

This field specifies flags to instruct how to interpret other attributes. See [Flags](#flags).

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

### `perf_event`

[:octicons-tag-24: v5.15](https://github.com/torvalds/linux/commit/82e6b1eee6a8875ef4eacfd60711cce6965c6b04)

```c
union bpf_attr {
	struct {
		[...]
		union {
			struct {
				__u64		[bpf_cookie](#bpf_cookie);
			} perf_event;
		}
	}
}
```

#### `bpf_cookie`

[:octicons-tag-24: v5.15](https://github.com/torvalds/linux/commit/82e6b1eee6a8875ef4eacfd60711cce6965c6b04)

This field is an optional opaque value which is reported back to tracing programs via the [`bpf_get_attach_cookie`](../helper-function/bpf_get_attach_cookie.md) helper.

The idea behind this cookie is that if the same program gets attached to multiple locations in the kernel, this value can be used to distinguish for which attach point the program is called. This ID value can for example be used as the key for a map which contains additional data the program needs or as key when collecting statistics.

### `kprobe_multi`

[:octicons-tag-24: v5.18](https://github.com/torvalds/linux/commit/0dcac272540613d41c05e89679e4ddb978b612f1)

```c
union bpf_attr {
	struct {
		[...]
		union {
			struct {
				__u32			[flags](#kprobe_multi-flags);
				__u32			[cnt](#cnt);
				__aligned_u64	[syms](#syms);
				__aligned_u64	[addrs](#addrs);
				__aligned_u64	[cookies](#cookies);
			} kprobe_multi;
		}
	}
}
```

This sub-struct is a collection of fields which specify one or multiple kprobe attachment points to attach the same program to multiple locations with a single syscall.

#### `flags` {#kprobe_multi-flags}

[:octicons-tag-24: v5.18](https://github.com/torvalds/linux/commit/0dcac272540613d41c05e89679e4ddb978b612f1)

Bitfield of flags, possible values are:

* `BPF_F_KPROBE_MULTI_RETURN` - When set, the kprobes are created as return probes.

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

```c
union bpf_attr {
	struct {
		[...]
		union {
			struct {
				/* this is overlaid with the target_btf_id above. */
				__u32		[target_btf_id](#tracing-target_btf_id);
				__u64		[cookie](#tracing-cookie);
			} tracing;
		}
	}
}
```

#### `target_btf_id` {#tracing-target_btf_id}

[:octicons-tag-24: v5.19](https://github.com/torvalds/linux/commit/2fcc82411e74e5e6aba336561cf56fb899bfae4e)

The definition in `tracing` is overlaid with [`target_btf_id`](#target_btf_id) in memory, and has the same meaning.

#### `cookie` {#tracing-cookie}

[:octicons-tag-24: v5.19](https://github.com/torvalds/linux/commit/2fcc82411e74e5e6aba336561cf56fb899bfae4e)

Same as [`bpf_cookie`](#bpf_cookie) but for tracing programs.

### `netfilter`

[:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/84601d6ee68ae820dec97450934797046d62db4b)

```c
union bpf_attr {
	struct {
		[...]
		union {
			struct {
				__u32		[pf](#pf);
				__u32		[hooknum](#hooknum);
				__s32		[priority](#priority);
				__u32		[flags](#netfilter-flags);
			} netfilter;
		}
	}
}
```

#### `pf`

[:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/84601d6ee68ae820dec97450934797046d62db4b)

The protocol family, supported values are `NFPROTO_IPV4` (2) and `NFPROTO_IPV6` (10).

#### `hooknum`

[:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/84601d6ee68ae820dec97450934797046d62db4b)

The hook number, supported values are `NF_INET_PRE_ROUTING` (0), `NF_INET_LOCAL_IN` (1), `NF_INET_FORWARD` (2), `NF_INET_LOCAL_OUT` (3), and `NF_INET_POST_ROUTING` (4).

#### `priority`

[:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/84601d6ee68ae820dec97450934797046d62db4b)

The priority of the hook, lower values are called first. `NF_IP_PRI_FIRST` (-2147483648) and `NF_IP_PRI_LAST` (2147483647) are not allowed.

#### `flags` {#netfilter-flags}

[:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/84601d6ee68ae820dec97450934797046d62db4b)

A bitmask of flags. Supported flags are:

* `BPF_F_NETFILTER_IP_DEFRAG` - Enable defragmentation of IP fragments, this hook will only see defragmented packets. If the `BPF_F_NETFILTER_IP_DEFRAG` [:octicons-tag-24: v6.6](https://github.com/torvalds/linux/commit/91721c2d02d3a0141df8a4787c7079b89b0d0607) flag is set, the priority must be higher than `NF_IP_PRI_CONNTRACK_DEFRAG` (-400) for ensuring the prog runs after nf_defrag.

### `tcx`

[:octicons-tag-24: v6.6](https://github.com/torvalds/linux/commit/e420bed025071a623d2720a92bc2245c84757ecb)

```c
union bpf_attr {
	struct {
		[...]
		union {
			struct {
				union {
					__u32	[relative_fd](#tcx-relative_fd);
					__u32	[relative_id](#tcx-relative_id);
				};
				__u64		[expected_revision](#tcx-expected_revision);
			} tcx;
		}
	}
}
```

#### `relative_fd` {#tcx-relative_fd}

[:octicons-tag-24: v6.6](https://github.com/torvalds/linux/commit/e420bed025071a623d2720a92bc2245c84757ecb)

The file descriptor of the program or link to attach relative to. 

* If `BPF_F_BEFORE` is set, the program is attached before the program/link indicated by this field. 
* If `BPF_F_AFTER` is set, the program is attached after the program/link indicated by this field.
* If `BPF_F_REPLACE` is set, the program replaced the program/link indicated by this field.

The above flags are mutually exclusive.

This field is used over [`relative_id`](#tcx-relative_id) when [`BPF_F_ID`](#bpf_f_id) is not set.

#### `relative_id` {#tcx-relative_id}

[:octicons-tag-24: v6.6](https://github.com/torvalds/linux/commit/e420bed025071a623d2720a92bc2245c84757ecb)

The ID of the program or link to attach relative to.

* If `BPF_F_BEFORE` is set, the program is attached before the program/link indicated by this field. 
* If `BPF_F_AFTER` is set, the program is attached after the program/link indicated by this field.
* If `BPF_F_REPLACE` is set, the program replaced the program/link indicated by this field.

The above flags are mutually exclusive.

This field is used over [`relative_fd`](#tcx-relative_fd) when [`BPF_F_ID`](#bpf_f_id) is set.

#### `expected_revision` {#tcx-expected_revision}

[:octicons-tag-24: v6.6](https://github.com/torvalds/linux/commit/e420bed025071a623d2720a92bc2245c84757ecb)

The expected <nospell>mprog</nospell> revision, to avoid unexpected behavior in case two links are created at the same time.

### `uprobe_multi`

[:octicons-tag-24: v6.6](https://github.com/torvalds/linux/commit/89ae89f53d201143560f1e9ed4bfa62eee34f88e)

```c
union bpf_attr {
	struct {
		[...]
		union {
			struct {
				__aligned_u64	[path](#uprobe_multi-path);
				__aligned_u64	[offsets](#uprobe_multi-offsets);
				__aligned_u64	[ref_ctr_offsets](#uprobe_multi-ref_ctr_offsets);
				__aligned_u64	[cookies](#uprobe_multi-cookies);
				__u32			[cnt](#uprobe_multi-cnt);
				__u32			[flags](#uprobe_multi-flags);
				__u32			[pid](#uprobe_multi-pid);
			} uprobe_multi;
		}
	}
}
```

#### `path` {#uprobe_multi-path}

[:octicons-tag-24: v6.6](https://github.com/torvalds/linux/commit/89ae89f53d201143560f1e9ed4bfa62eee34f88e)

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome


#### `offsets` {#uprobe_multi-offsets}

[:octicons-tag-24: v6.6](https://github.com/torvalds/linux/commit/89ae89f53d201143560f1e9ed4bfa62eee34f88e)

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome


#### `ref_ctr_offsets` {#uprobe_multi-ref_ctr_offsets}

[:octicons-tag-24: v6.6](https://github.com/torvalds/linux/commit/89ae89f53d201143560f1e9ed4bfa62eee34f88e)

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome


#### `cookies` {#uprobe_multi-cookies}

[:octicons-tag-24: v6.6](https://github.com/torvalds/linux/commit/89ae89f53d201143560f1e9ed4bfa62eee34f88e)

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome


#### `cnt` {#uprobe_multi-cnt}

[:octicons-tag-24: v6.6](https://github.com/torvalds/linux/commit/89ae89f53d201143560f1e9ed4bfa62eee34f88e)

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome


#### `flags` {#uprobe_multi-flags}

[:octicons-tag-24: v6.6](https://github.com/torvalds/linux/commit/89ae89f53d201143560f1e9ed4bfa62eee34f88e)

Bitfield of flags, possible values are:

* `BPF_F_UPROBE_MULTI_RETURN` - When set, the kprobes are created as return probes.


#### `pid` {#uprobe_multi-pid}

[:octicons-tag-24: v6.6](https://github.com/torvalds/linux/commit/b733eeade4204423711793595c3c8d78a2fa8b2e)

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### `netkit`

[:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/35dfaad7188cdc043fde31709c796f5a692ba2bd)

```c
union bpf_attr {
	struct {
		[...]
		union {
			struct {
				union {
					__u32	[relative_fd](#netkit-relative_fd);
					__u32	[relative_id](#netkit-relative_id);
				};
				__u64		[expected_revision](#netkit-expected_revision);
			} netkit;
		}
	}
}
```

#### `relative_fd` {#netkit-relative_fd}

[:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/35dfaad7188cdc043fde31709c796f5a692ba2bd)

The file descriptor of the program or link to attach relative to.

* If `BPF_F_BEFORE` is set, the program is attached before the program/link indicated by this field. 
* If `BPF_F_AFTER` is set, the program is attached after the program/link indicated by this field.
* If `BPF_F_REPLACE` is set, the program replaced the program/link indicated by this field.

The above flags are mutually exclusive.

This field is used over [`relative_id`](#netkit-relative_id) when [`BPF_F_ID`](#bpf_f_id) is not set.

#### `relative_id` {#netkit-relative_id}

[:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/35dfaad7188cdc043fde31709c796f5a692ba2bd)

The ID of the program or link to attach relative to.

* If `BPF_F_BEFORE` is set, the program is attached before the program/link indicated by this field. 
* If `BPF_F_AFTER` is set, the program is attached after the program/link indicated by this field.
* If `BPF_F_REPLACE` is set, the program replaced the program/link indicated by this field.

The above flags are mutually exclusive.

This field is used over [`relative_fd`](#netkit-relative_fd) when [`BPF_F_ID`](#bpf_f_id) is set.

#### `expected_revision` {#netkit-expected_revision}

[:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/35dfaad7188cdc043fde31709c796f5a692ba2bd)

The expected <nospell>mprog</nospell> revision, to avoid unexpected behavior in case two links are created at the same time.

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

### `BPF_TCX_INGRESS`

<!-- [FEATURE_TAG](BPF_TCX_INGRESS) -->
[:octicons-tag-24: v6.6](https://github.com/torvalds/linux/commit/e420bed025071a623d2720a92bc2245c84757ecb)
<!-- [/FEATURE_TAG] -->


### `BPF_TCX_EGRESS`

<!-- [FEATURE_TAG](BPF_TCX_EGRESS) -->
[:octicons-tag-24: v6.6](https://github.com/torvalds/linux/commit/e420bed025071a623d2720a92bc2245c84757ecb)
<!-- [/FEATURE_TAG] -->


### `BPF_NETKIT_PRIMARY`

<!-- [FEATURE_TAG](BPF_NETKIT_PRIMARY) -->
[:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/35dfaad7188cdc043fde31709c796f5a692ba2bd)
<!-- [/FEATURE_TAG] -->


### `BPF_NETKIT_PEER`

<!-- [FEATURE_TAG](BPF_NETKIT_PEER) -->
[:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/35dfaad7188cdc043fde31709c796f5a692ba2bd)
<!-- [/FEATURE_TAG] -->


### `BPF_TRACE_KPROBE_SESSION`

<!-- [FEATURE_TAG](BPF_TRACE_KPROBE_SESSION) -->
[:octicons-tag-24: v6.10](https://github.com/torvalds/linux/commit/535a3692ba7245792e6f23654507865d4293c850)
<!-- [/FEATURE_TAG] -->

### `BPF_TRACE_UPROBE_SESSION`

[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/d920179b3d4842a0e27cae54fdddbe5ef3977e73)

## Flags

### `BPF_F_REPLACE`

[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/7dd68b3279f1792103d12e69933db3128c6d416e)

### `BPF_F_BEFORE`

[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/7dd68b3279f1792103d12e69933db3128c6d416e)

### `BPF_F_AFTER`

[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/7dd68b3279f1792103d12e69933db3128c6d416e)

### `BPF_F_ID`

[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/7dd68b3279f1792103d12e69933db3128c6d416e)

### `BPF_F_LINK`

[:octicons-tag-24: v6.13](https://github.com/torvalds/linux/commit/7dd68b3279f1792103d12e69933db3128c6d416e)

