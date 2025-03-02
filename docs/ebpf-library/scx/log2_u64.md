---
title: "SCX eBPF function 'log2_u64'"
description: "This page documents the 'log2_u64' scx eBPF function, including its definition, usage, and examples."
---
# SCX eBPF function `log2_u64`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/2a52ca7c98960aafb0eca9ef96b2d0c932171357)

The `log2_u64` function computes the base 2 logarithm of a 64-bit exponential value.

## Definition

```c
static inline u32 log2_u64(u64 v)
{
        u32 hi = v >> 32;
        if (hi)
                return [log2_u32](log2_u32.md)(hi) + 32 + 1;
        else
                return [log2_u32](log2_u32.md)(v) + 1;
}
```

## Usage

Compute the base 2 logarithm of a 64-bit value.

**Parameters**

- `v`: The value for which we're computing the base 2 logarithm.

### Example


!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
