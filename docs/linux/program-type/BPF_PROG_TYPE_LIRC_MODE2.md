---
title: "Program Type 'BPF_PROG_TYPE_LIRC_MODE2'"
description: "This page documents the 'BPF_PROG_TYPE_LIRC_MODE2' eBPF program type, including its definition, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_LIRC_MODE2`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_LIRC_MODE2) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/f4364dcfc86df7c1ca47b256eaf6b6d0cdd0d936)
<!-- [/FEATURE_TAG] -->

`BPF_PROG_TYPE_LIRC_MODE2` are eBPF programs that can attach to a [LIRC](https://docs.kernel.org/userspace-api/media/rc/lirc-dev-intro.html) (Linux Infrared Remote Control) device to perform IR protocol decoding.

## Usage

A LIRC device is presented to userspace as a character device like `/dev/lirc0`. Multiple modes exists, depending on the mode the kernel will decode the IR signal and allows userspace to read the scan codes or the IR signal can be communicated to userspace so applications can decode protocols the kernel does not know about.

This program type allows you to implement IR decoding in eBPF instead, so userspace can just read the decoded scan codes.

A receiver detect IR light or the absence of it. The duration of light and absence of light is used to communicate. A period where the IR light was on is called a pulse/flash, and a period where the IR light was off is called a space/gap.

The eBPF program is called every time a transition is detected or a special event occurs. So for a single "message" the program may be called many times. Programs typically have to keep state between calls to decode the message.

When a program decodes a message, it can report the meaning to the OS via helpers. [`bpf_rc_keydown`](../helper-function/bpf_rc_keydown.md) reports a key press (typically on something like a TV remote). [`bpf_rc_repeat`](../helper-function/bpf_rc_repeat.md) signals to repeat the previous key press (typically used when holding down a button on a TV remote). [`bpf_rc_pointer_rel`](../helper-function/bpf_rc_pointer_rel.md) reports a relative change in pointer position.

The [cir](https://gitlab.freedesktop.org/linux-media/tools/cir) project is a good example user of this program type. It implements userspace tooling that convert <nospell>keymaps</nospell> and <nospell>[IRP](https://gitlab.freedesktop.org/linux-media/tools/cir/blob/main/irp/doc/irp_introduction.md)</nospell> notation into eBPF programs to decode IR signals.

## Context

The context for this program type is a pointer to an unsigned 32 bit integer. The upper 8 bits determine the packet type, and the lower 24 bits the payload. 

The [`include/uapi/linux/lirc.h`](https://github.com/torvalds/linux/blob/master/include/uapi/linux/lirc.h) header file contains a number of useful macros to work with the context.

The [`LIRC_VALUE`](https://elixir.bootlin.com/linux/v6.13.7/source/include/uapi/linux/lirc.h#L30) macro can be used to get the payload, and the [LIRC_MODE2](https://elixir.bootlin.com/linux/v6.13.7/source/include/uapi/linux/lirc.h#L31) macro will give you the type, which is one of:

* `LIRC_MODE2_PULSE` - Signifies the presence of IR in microseconds, also known as flash.
* `LIRC_MODE2_SPACE` - Signifies absence of IR in microseconds, also known as gap.
* `LIRC_MODE2_FREQUENCY` - If measurement of the carrier frequency was enabled with [`ioctl`](https://man7.org/linux/man-pages/man2/ioctl.2.html) [`LIRC_SET_MEASURE_CARRIER_MODE`](https://docs.kernel.org/userspace-api/media/rc/lirc-set-measure-carrier-mode.html#lirc-set-measure-carrier-mode) then this packet gives you the carrier frequency in Hertz.
* `LIRC_MODE2_TIMEOUT` - When the timeout set with [`ioctl`](https://man7.org/linux/man-pages/man2/ioctl.2.html) [`LIRC_GET_REC_TIMEOUT` and `LIRC_SET_REC_TIMEOUT`](https://docs.kernel.org/userspace-api/media/rc/lirc-set-rec-timeout.html#lirc-set-rec-timeout) expires due to no IR being detected, this packet will be sent, with the number of microseconds with no IR.
* `LIRC_MODE2_OVERFLOW` - Signifies that the IR receiver encounter an overflow, and some IR is missing. The IR data after this should be correct again. The actual value is not important, but this is set to `0xffffff` by the kernel for compatibility with [lircd](https://linux.die.net/man/8/lircd).

## Attachment

Attachment is through the [`BPF_PROG_ATTACH`](../syscall/BPF_PROG_ATTACH.md) syscall. The `target_fd` should be the LIRC device you want to attach the program to. A maximum of 64 programs can be attached to a single LIRC device at a time.

## Example 

Example BPF program:

```c
// SPDX-License-Identifier: GPL-2.0
// test ir decoder
//
// Copyright (C) 2018 Sean Young <sean@mess.org>

#include <linux/bpf.h>
#include <linux/lirc.h>
#include <bpf/bpf_helpers.h>

SEC("lirc_mode2")
int bpf_decoder(unsigned int *sample)
{
    if ([LIRC_IS_PULSE](https://elixir.bootlin.com/linux/v6.13.7/source/include/uapi/linux/lirc.h#L34)(*sample)) {
        unsigned int duration = [LIRC_VALUE](https://elixir.bootlin.com/linux/v6.13.7/source/include/uapi/linux/lirc.h#L30)(*sample);

        if (duration & 0x10000)
            [bpf_rc_keydown](../helper-function/bpf_rc_keydown.md)(sample, 0x40, duration & 0xffff, 0);
        if (duration & 0x20000)
            [bpf_rc_pointer_rel](../helper-function/bpf_rc_pointer_rel.md)(sample, (duration >> 8) & 0xff,
                       duration & 0xff);
    }

    return 0;
}

char _license[] SEC("license") = "GPL";
```

## Helper functions

Not all helper functions are available in all program types. These are the helper calls available for LIRC mode 2 programs:

<!-- DO NOT EDIT MANUALLY -->
<!-- [PROG_HELPER_FUNC_REF] -->
??? abstract "Supported helper functions"
    * [`bpf_get_prandom_u32`](../helper-function/bpf_get_prandom_u32.md)
    * [`bpf_ktime_get_boot_ns`](../helper-function/bpf_ktime_get_boot_ns.md)
    * [`bpf_ktime_get_ns`](../helper-function/bpf_ktime_get_ns.md)
    * [`bpf_map_delete_elem`](../helper-function/bpf_map_delete_elem.md)
    * [`bpf_map_lookup_elem`](../helper-function/bpf_map_lookup_elem.md)
    * [`bpf_map_peek_elem`](../helper-function/bpf_map_peek_elem.md)
    * [`bpf_map_pop_elem`](../helper-function/bpf_map_pop_elem.md)
    * [`bpf_map_push_elem`](../helper-function/bpf_map_push_elem.md)
    * [`bpf_map_update_elem`](../helper-function/bpf_map_update_elem.md)
    * [`bpf_rc_keydown`](../helper-function/bpf_rc_keydown.md)
    * [`bpf_rc_pointer_rel`](../helper-function/bpf_rc_pointer_rel.md)
    * [`bpf_rc_repeat`](../helper-function/bpf_rc_repeat.md)
    * [`bpf_tail_call`](../helper-function/bpf_tail_call.md)
    * [`bpf_trace_printk`](../helper-function/bpf_trace_printk.md)
<!-- [/PROG_HELPER_FUNC_REF] -->

## KFuncs

<!-- [PROG_KFUNC_REF] -->
There are currently no kfuncs supported for this program type
<!-- [/PROG_KFUNC_REF] -->
