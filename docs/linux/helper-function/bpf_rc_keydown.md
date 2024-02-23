---
title: "Helper Function 'bpf_rc_keydown' - eBPF Docs"
description: "This page documents the 'bpf_rc_keydown' eBPF helper function, including its defintion, usage, program types that can use it, and examples."
---
# Helper function `bpf_rc_keydown`

<!-- [FEATURE_TAG](bpf_rc_keydown) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/f4364dcfc86df7c1ca47b256eaf6b6d0cdd0d936)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
This helper is used in programs implementing IR decoding, to report a successfully decoded key press with _scancode_, _toggle_ value in the given _protocol_. The scancode will be translated to a keycode using the rc keymap, and reported as an input key down event. After a period a key up event is generated. This period can be extended by calling either **bpf_rc_keydown**() again with the same values, or calling **bpf_rc_repeat**().

Some protocols include a toggle bit, in case the button was released and pressed again between consecutive scancodes.

The _ctx_ should point to the lirc sample as passed into the program.

The _protocol_ is the decoded protocol number (see **enum rc_proto** for some predefined values).

This helper is only available is the kernel was compiled with the **CONFIG_BPF_LIRC_MODE2** configuration option set to "**y**".

### Returns

0

`#!c static long (*bpf_rc_keydown)(void *ctx, __u32 protocol, __u64 scancode, __u32 toggle) = (void *) 78;`
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
