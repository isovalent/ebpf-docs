---
title: "Libbpf userspace function 'bpf_xdp_attach'"
description: "This page documents the 'bpf_xdp_attach' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_xdp_attach`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.7.0](https://github.com/libbpf/libbpf/releases/tag/v0.7.0)
<!-- [/LIBBPF_TAG] -->

Attach a [`BPF_PROG_TYPE_XDP`](../../../linux/program-type/BPF_PROG_TYPE_XDP.md) program to a network interface. Via netlink.

## Definition

`#!c int bpf_xdp_attach(int ifindex, int prog_fd, __u32 flags, const struct bpf_xdp_attach_opts *opts);`

**Parameters**

- `ifindex`: index of the network interface to attach the program to
- `prog_fd`: file descriptor of the BPF program to attach
- `flags`: flags to control the attachment behavior
- `opts`: options to control the attachment, see [`struct bpf_xdp_attach_opts`](#struct-bpf_xdp_attach_opts)

**Flags**

* `XDP_FLAGS_UPDATE_IF_NOEXIST` = `(1U << 0)` - If set, only attach if no program is already attached to the interface.
* `XDP_FLAGS_SKB_MODE` = `(1U << 1)` - If set, force loading in SKB (generic) mode.
* `XDP_FLAGS_DRV_MODE` = `(1U << 2)` - If set, force loading in DRV (driver / native) mode.
* `XDP_FLAGS_HW_MODE` = `(1U << 3)` - If set, force loading in hardware offload mode.
* `XDP_FLAGS_REPLACE` = `(1U << 4)` - If set, replace the currently attached program, throw an error if no program is attached.

`XDP_FLAGS_SKB_MODE`, `XDP_FLAGS_DRV_MODE`, and `XDP_FLAGS_HW_MODE` are mutually exclusive. If none are set, `XDP_FLAGS_DRV_MODE` is used when the driver of the interface supports it, and falls back to `XDP_FLAGS_SKB_MODE` otherwise.

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

This function attaches a BPF program to a network interface using the older netlink-based method, unlike the BPF link based [`bpf_program__attach_xdp`](bpf_program__attach_xdp.md) function.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
