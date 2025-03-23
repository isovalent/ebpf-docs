---
title: "Libbpf userspace function 'bpf_prog_query_opts'"
description: "This page documents the 'bpf_prog_query_opts' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_prog_query_opts`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 1.0.0](https://github.com/libbpf/libbpf/releases/tag/v1.0.0)
<!-- [/LIBBPF_TAG] -->

Queries the BPF programs and BPF links which are attached to `target` which can represent a file descriptor or netdevice ifindex.

## Definition

`#!c int bpf_prog_query_opts(int target, enum bpf_attach_type type, struct bpf_prog_query_opts *opts);`

**Parameters**

- `target`: query location file descriptor or ifindex
- `type`: attach type for the BPF program
- `opts`: options for configuring the query

**Return**

`0`, on success; negative error code, otherwise ([`errno`](https://man7.org/linux/man-pages/man3/errno.3.html) is also set to the error code)

### `struct bpf_prog_query_opts`

```c
struct bpf_prog_query_opts {
	size_t sz; /* size of this struct for forward/backward compatibility */
	__u32 query_flags;
	__u32 attach_flags; /* output argument */
	__u32 *prog_ids;
	union {
		/* input+output argument */
		__u32 prog_cnt;
		__u32 count;
	};
	__u32 *prog_attach_flags;
	__u32 *link_ids;
	__u32 *link_attach_flags;
	__u64 revision;
	size_t :0;
};
```

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
