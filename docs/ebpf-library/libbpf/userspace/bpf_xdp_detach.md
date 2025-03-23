---
title: "Libbpf userspace function 'bpf_xdp_detach'"
description: "This page documents the 'bpf_xdp_detach' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_xdp_detach`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.7.0](https://github.com/libbpf/libbpf/releases/tag/v0.7.0)
<!-- [/LIBBPF_TAG] -->

Detach a [`BPF_PROG_TYPE_XDP`](../../../linux/program-type/BPF_PROG_TYPE_XDP.md) program that was previously attached via [`bpf_xdp_attach`](bpf_xdp_attach.md).

## Definition

`#!c int bpf_xdp_detach(int ifindex, __u32 flags, const struct bpf_xdp_attach_opts *opts);`

**Parameters**

- `ifindex`: index of the network interface to detach the program from
- `flags`: flags to control the detachment behavior
- `opts`: options to control the detachment, see [`struct bpf_xdp_attach_opts`](#struct-bpf_xdp_attach_opts)

**Flags**

* `XDP_FLAGS_REPLACE` = `(1U << 4)` - If set, remove the currently attached program, throw an error if no program is attached.

**Return**

Zero on success, or a negative error code on failure.

### `struct bpf_xdp_attach_opts`

```c
struct bpf_xdp_attach_opts {
	size_t sz;
	int old_prog_fd;
	size_t :0;
};
```
## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
