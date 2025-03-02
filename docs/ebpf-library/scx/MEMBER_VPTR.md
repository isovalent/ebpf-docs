---
title: "SCX eBPF macro 'MEMBER_VPTR'"
description: "This page documents the 'MEMBER_VPTR' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `MEMBER_VPTR`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/2a52ca7c98960aafb0eca9ef96b2d0c932171357)

The `MEMBER_VPTR` macro obtains a verified pointer to a struct or array member

## Definition

```c
#define MEMBER_VPTR(base, member) (typeof((base) member) *)             \
({                                                                      \
    u64 __base = (u64)&(base);                                          \
    u64 __addr = (u64)&((base) member) - __base;                        \
    _Static_assert(sizeof(base) >= sizeof((base) member),               \
               "@base is smaller than @member, is @base a pointer?");   \
    asm volatile (                                                      \
        "if %0 <= %[max] goto +2\n"                                     \
        "%0 = 0\n"                                                      \
        "goto +1\n"                                                     \
        "%0 += %1\n"                                                    \
        : "+r"(__addr)                                                  \
        : "r"(__base),                                                  \
          [max]"i"(sizeof(base) - sizeof((base) member)));              \
    __addr;                                                             \
})
```

## Usage

The verifier often gets confused by the instruction sequence the compiler generates for indexing struct fields or arrays. This macro forces the compiler to generate a code sequence which first calculates the byte offset, checks it against the struct or array size and add that byte offset to generate the pointer to the member to help the verifier.

Ideally, we want to abort if the calculated offset is out-of-bounds. However, BPF currently doesn't support abort, so evaluate to `NULL` instead. The caller must check for `NULL` and take appropriate action to appease the verifier. To avoid confusing the verifier, it's best to check for `NULL` and dereference immediately.

```c
vptr = MEMBER_VPTR(my_array, [i][j]);
if (!vptr)
    return error;
*vptr = new_value;
```

`sizeof(base)` should encompass the memory area to be accessed and thus can't be a pointer to the area. Use `MEMBER_VPTR(*ptr, .member)` instead of `MEMBER_VPTR(ptr, ->member)`.

**Parameters**

 * @base: struct or array to index
 * @member: dereferenced member (e.g. `.field`, `[idx0][idx1]`, `.field[idx0]` ...)

### Example

```c hl_lines="33"
/* SPDX-License-Identifier: GPL-2.0 */
/* Copyright (c) 2022 Meta Platforms, Inc. and affiliates.
 * Copyright (c) 2022 Tejun Heo <tj@kernel.org>
 * Copyright (c) 2022 David Vernet <dvernet@meta.com>
 */

/*
 * Print out the online and possible CPU map using bpf_printk() as a
 * demonstration of using the cpumask kfuncs and ops.cpu_on/offline().
 */
static void print_cpus(void)
{
	const struct cpumask *possible, *online;
	s32 cpu;
	char buf[128] = "", *p;
	int idx;

	possible = [scx_bpf_get_possible_cpumask](../../linux/kfuncs/scx_bpf_get_possible_cpumask.md)();
	online = [scx_bpf_get_online_cpumask](../../linux/kfuncs/scx_bpf_get_online_cpumask.md)();

	idx = 0;
	[bpf_for](../libbpf/ebpf/bpf_for.md)(cpu, 0, [scx_bpf_nr_cpu_ids](../../linux/kfuncs/scx_bpf_nr_cpu_ids.md)()) {
		if (!(p = MEMBER_VPTR(buf, [idx++])))
			break;
		if ([bpf_cpumask_test_cpu](../../linux/kfuncs/bpf_cpumask_test_cpu.md)(cpu, online))
			*p++ = 'O';
		else if ([bpf_cpumask_test_cpu](../../linux/kfuncs/bpf_cpumask_test_cpu.md)(cpu, possible))
			*p++ = 'X';
		else
			*p++ = ' ';

		if ((cpu & 7) == 7) {
			if (!(p = MEMBER_VPTR(buf, [idx++])))
				break;
			*p++ = '|';
		}
	}
	buf[sizeof(buf) - 1] = '\0';

	[scx_bpf_put_cpumask](../../linux/kfuncs/scx_bpf_put_cpumask.md)(online);
	[scx_bpf_put_cpumask](../../linux/kfuncs/scx_bpf_put_cpumask.md)(possible);

	[bpf_printk](../libbpf/ebpf/bpf_printk.md)("CPUS: |%s", buf);
}

void [BPF_STRUCT_OPS](BPF_STRUCT_OPS.md)(qmap_cpu_online, s32 cpu)
{
	[bpf_printk](../libbpf/ebpf/bpf_printk.md)("CPU %d coming online", cpu);
	/* @cpu is already online at this point */
	print_cpus();
}
```
