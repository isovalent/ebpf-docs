---
title: "Libbpf eBPF macro '__hidden'"
description: "This page documents the '__hidden' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `__hidden`

[:octicons-tag-24: v0.4.0](https://github.com/libbpf/libbpf/releases/tag/v0.4.0)

The `__hidden` macros is used make a symbol as hidden.

## Definition

`#!c #define __hidden __attribute__((visibility("hidden")))`

## Usage

This macro is used to mark a symbol as hidden. This can be used for two thing. First, if applied on a global function, the loader (library) will modify the BTF to tell the verifier its a static function. This triggers different BPF-to-BPF function verification, see [Function by function verification](../../../linux/concepts/functions.md#function-by-function-verification).

The second use case is that when applied to global variables, they will be hidden from userspace, effectively making them eBPF only.

### Example

In this example we use `__hidden` since userspace should not interfere with the spinlock, which is only used by eBPF programs. However, this is a very bad example, please to not use them this way in practice.

```c hl_lines="1"
struct bpf_spin_lock lockA __hidden SEC(".data.A");
int counter = 0;

SEC("xdp")
int example_prog(struct xdp_md *ctx)
{
    bpf_spin_lock(&lockA);
    counter++;
    bpf_spin_unlock(&lockA);
	return XDP_PASS;
}
```
