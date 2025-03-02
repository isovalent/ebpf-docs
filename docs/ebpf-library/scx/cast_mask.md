---
title: "SCX eBPF macro 'cast_mask'"
description: "This page documents the 'cast_mask' scx eBPF macro, including its definition, usage, and examples."
---
# SCX eBPF macro `cast_mask`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/2a52ca7c98960aafb0eca9ef96b2d0c932171357)

The `cast_mask` macro casts a BPF cpumask to a regular cpumask.

## Definition

```c
static __always_inline const struct cpumask *cast_mask(struct bpf_cpumask *mask)
{
	return (const struct cpumask *)mask;
}
```

## Usage

For when you have a `struct bpf_cpumask` and need a `struct cpumask`.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
