---
title: "SCX eBPF macro 'ARRAY_ELEM_PTR'"
description: "This page documents the 'ARRAY_ELEM_PTR' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `ARRAY_ELEM_PTR`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/2a52ca7c98960aafb0eca9ef96b2d0c932171357)

The `ARRAY_ELEM_PTR` macro obtains a verified pointer to an array element.

## Definition

```c
#define ARRAY_ELEM_PTR(arr, i, n) (typeof(arr[i]) *)    \
({                                                      \
    u64 __base = (u64)arr;                              \
    u64 __addr = (u64)&(arr[i]) - __base;               \
    asm volatile (                                      \
        "if %0 <= %[max] goto +2\n"                     \
        "%0 = 0\n"                                      \
        "goto +1\n"                                     \
        "%0 += %1\n"                                    \
        : "+r"(__addr)                                  \
        : "r"(__base),                                  \
          [max]"r"(sizeof(arr[0]) * ((n) - 1)));        \
    __addr;                                             \
})
```

## Usage

Similar to [`MEMBER_VPTR`](MEMBER_VPTR.md) but is intended for use with arrays where the element count needs to be explicit. It can be used in cases where a global array is defined with an initial size but is intended to be be resized before loading the BPF program. Without this version of the macro, [`MEMBER_VPTR`](MEMBER_VPTR.md) will use the compile time size of the array to compute the max, which will result in rejection by the verifier.

**Parameters**

`arr`: array to index into

`i`: array index

`n`: number of elements in array


### Example

```c hl_lines="12"
/* SPDX-License-Identifier: GPL-2.0 */
/* Copyright (c) 2022 Meta Platforms, Inc. and affiliates.
 * Copyright (c) 2022 Tejun Heo <tj@kernel.org>
 * Copyright (c) 2022 David Vernet <dvernet@meta.com>
 */

u64 [RESIZABLE_ARRAY](RESIZABLE_ARRAY.md)(data, cpu_started_at);

void [BPF_STRUCT_OPS](BPF_STRUCT_OPS.md)(central_running, struct task_struct *p)
{
    s32 cpu = [scx_bpf_task_cpu](../../linux/kfuncs/scx_bpf_task_cpu.md)(p);
    u64 *started_at = ARRAY_ELEM_PTR(cpu_started_at, cpu, nr_cpu_ids);
    if (started_at)
        *started_at = [scx_bpf_now](../../linux/kfuncs/scx_bpf_now.md)() ?: 1;	/* 0 indicates idle */
}
```
