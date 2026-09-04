---
title: "KFunc 'bpf_iter_dmabuf_new'"
description: "This page documents the 'bpf_iter_dmabuf_new' eBPF kfunc, including its definition, usage, program types that can use it, and examples."
---
# KFunc `bpf_iter_dmabuf_new`

<!-- [FEATURE_TAG](bpf_iter_dmabuf_new) -->
[:octicons-tag-24: v6.16](https://github.com/torvalds/linux/commit/6eab7ac7c5eea7628b92cd5f9427bbd963a954ec)
<!-- [/FEATURE_TAG] -->

This function initializes a iterator for DMA buffers.

## Definition

**Parameters**

`it`: A pointer to a stack allocated `struct bpf_iter_dmabuf` that is used to iterate over DMA buffers. 

**Returns**

`0` on success, a negative error code on failure

**Signature**

<!-- [KFUNC_DEF] -->
`#!c int bpf_iter_dmabuf_new(struct bpf_iter_dmabuf *it)`

!!! note
    This function may sleep, and therefore can only be used from [sleepable programs](../syscall/BPF_PROG_LOAD.md/#bpf_f_sleepable).
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
- [`BPF_PROG_TYPE_STRUCT_OPS`](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [`BPF_PROG_TYPE_SYSCALL`](../program-type/BPF_PROG_TYPE_SYSCALL.md)
- [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

```c
// SPDX-License-Identifier: GPL-2.0
/* Copyright (c) 2025 Google LLC */
#include <vmlinux.h>
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_helpers.h>

/* From uapi/linux/dma-buf.h */
#define DMA_BUF_NAME_LEN 32

char _license[] SEC("license") = "GPL";

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(key_size, DMA_BUF_NAME_LEN);
	__type(value, bool);
	__uint(max_entries, 5);
} testbuf_hash SEC(".maps");

SEC("syscall")
int iter_dmabuf_for_each(const void *ctx)
{
	struct dma_buf *d;

	[bpf_for_each](../../ebpf-library/libbpf/ebpf/bpf_for_each.md)(dmabuf, d) {
		char name[DMA_BUF_NAME_LEN];
		const char *pname;
		bool *found;
		long len;
		int i;

		if ([bpf_core_read](../../ebpf-library/libbpf/ebpf/bpf_core_read.md)(&pname, sizeof(pname), &d->name))
			return 1;

		/* Buffers are not required to be named */
		if (!pname)
			continue;

		len = [bpf_probe_read_kernel_str](../helper-function/bpf_probe_read_kernel_str.md)(name, sizeof(name), pname);
		if (len < 0)
			return 1;

		/*
		 * The entire name buffer is used as a map key.
		 * Zeroize any uninitialized trailing bytes after the NUL.
		 */
		[bpf_for](../../ebpf-library/libbpf/ebpf/bpf_for.md)(i, len, DMA_BUF_NAME_LEN)
			name[i] = 0;

		found = [bpf_map_lookup_elem](../helper-function/bpf_map_lookup_elem.md)(&testbuf_hash, name);
		if (found) {
			bool t = true;

			[bpf_map_update_elem](../helper-function/bpf_map_update_elem.md)(&testbuf_hash, name, &t, BPF_EXIST);
		}
	}

	return 0;
}
```
