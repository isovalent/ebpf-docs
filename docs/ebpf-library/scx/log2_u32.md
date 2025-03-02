---
title: "SCX eBPF function 'log2_u32'"
description: "This page documents the 'log2_u32' scx eBPF function, including its definition, usage, and examples."
---
# SCX eBPF function `log2_u32`

[:octicons-tag-24: v6.12](https://github.com/torvalds/linux/commit/2a52ca7c98960aafb0eca9ef96b2d0c932171357)

The `log2_u32` function computes the base 2 logarithm of a 32-bit exponential value.

## Definition

```c
static inline u32 log2_u32(u32 v)
{
    u32 r;
    u32 shift;

    r = (v > 0xFFFF) << 4; v >>= r;
    shift = (v > 0xFF) << 3; v >>= shift; r |= shift;
    shift = (v > 0xF) << 2; v >>= shift; r |= shift;
    shift = (v > 0x3) << 1; v >>= shift; r |= shift;
    r |= (v >> 1);
    return r;
}
```

## Usage

Compute the base 2 logarithm of a 32-bit value.

**Parameters**

- `v`: The value for which we're computing the base 2 logarithm.

### Example


!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
