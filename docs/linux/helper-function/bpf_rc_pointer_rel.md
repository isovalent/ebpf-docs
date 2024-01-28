# Helper function `bpf_rc_pointer_rel`

<!-- [FEATURE_TAG](bpf_rc_pointer_rel) -->
[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/01d3240a04f4c09392e13c77b54d4423ebce2d72)
<!-- [/FEATURE_TAG] -->

## Defintion

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
This helper is used in programs implementing IR decoding, to report a successfully decoded pointer movement.

The _ctx_ should point to the lirc sample as passed into the program.

This helper is only available is the kernel was compiled with the **CONFIG_BPF_LIRC_MODE2** configuration option set to "**y**".

### Returns

0

`#!c static long (*bpf_rc_pointer_rel)(void *ctx, __s32 rel_x, __s32 rel_y) = (void *) 92;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_LIRC_MODE2](../program-type/BPF_PROG_TYPE_LIRC_MODE2.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
