# Helper function `bpf_rc_repeat`

<!-- [FEATURE_TAG](bpf_rc_repeat) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/f4364dcfc86df7c1ca47b256eaf6b6d0cdd0d936)
<!-- [/FEATURE_TAG] -->

## Definition

<!-- [HELPER_FUNC_DEF] -->
This helper is used in programs implementing IR decoding, to report a successfully decoded repeat key message. This delays the generation of a key up event for previously generated key down event.

Some IR protocols like NEC have a special IR message for repeating last button, for when a button is held down.

The _ctx_ should point to the lirc sample as passed into the program.

This helper is only available is the kernel was compiled with the **CONFIG_BPF_LIRC_MODE2** configuration option set to "**y**".

### Returns

0

`#!c static long (*bpf_rc_repeat)(void *ctx) = (void *) 77;`
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
