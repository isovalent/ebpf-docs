---
title: "Libbpf eBPF macro 'BPF_KPROBE'"
description: "This page documents the 'BPF_KPROBE' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `BPF_KPROBE`

[:octicons-tag-24: v0.0.8](https://github.com/libbpf/libbpf/releases/tag/v0.0.8)

The `BPF_KPROBE` macro makes it easier to write [kprobe](../../../linux/program-type/BPF_PROG_TYPE_KPROBE.md) programs.

!!! note
    The original context will stay available as `ctx`, if you ever wish to access it manually or need to pass it to a helper or kfunc. Therefor, the variable name `ctx` should not be reused in arguments or function body.

## Definition

```c
#define BPF_KPROBE(name, args...)					    \
name(struct pt_regs *ctx);						    \
static __always_inline typeof(name(0))					    \
____##name(struct pt_regs *ctx, ##args);				    \
typeof(name(0)) name(struct pt_regs *ctx)				    \
{									    \
	_Pragma("GCC diagnostic push")					    \
	_Pragma("GCC diagnostic ignored \"-Wint-conversion\"")		    \
	return ____##name(___bpf_kprobe_args(args));			    \
	_Pragma("GCC diagnostic pop")					    \
}									    \
static __always_inline typeof(name(0))					    \
____##name(struct pt_regs *ctx, ##args)
```

## Usage

This macro is useful when writing kprobe programs that attach at the start of a function. Traditionally a program author would have to use the [`PT_REGS_PARAM`](PT_REGS_PARM.md) macros to extract a given parameter from the context and then manually cast them to the actual type.

The `BPF_KPROBE` macro allows you to write your program with an argument list, the macro will do the casting for you. This makes reading and writing kprobes easier.

### Example

```c hl_lines="2"
SEC("kprobe/bpf_map_copy_value")
int BPF_KPROBE(bpf_prog2, struct bpf_map *map)
{
	u32 key = bpf_get_smp_processor_id();
	struct bpf_perf_event_value *val, buf;
	enum bpf_map_type type;
	int error;

	type = BPF_CORE_READ(map, map_type);
	if (type != BPF_MAP_TYPE_HASH)
		return 0;

	error = bpf_perf_event_read_value(&counters, key, &buf, sizeof(buf));
	if (error)
		return 0;

	val = bpf_map_lookup_elem(&values2, &key);
	if (val)
		*val = buf;
	else
		bpf_map_update_elem(&values2, &key, &buf, BPF_NOEXIST);

	return 0;
}
```
