---
title: "Program Type 'BPF_PROG_TYPE_LIRC_MODE2'"
description: "This page documents the 'BPF_PROG_TYPE_LIRC_MODE2' eBPF program type, including its definition, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_LIRC_MODE2`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_LIRC_MODE2) -->
[:octicons-tag-24: v4.18](https://github.com/torvalds/linux/commit/f4364dcfc86df7c1ca47b256eaf6b6d0cdd0d936)
<!-- [/FEATURE_TAG] -->

`BPF_PROG_TYPE_LIRC_MODE2` are eBPF programs that can attach to a LIRC (Linux Infrared Remote Control) device to perform IR protocol decoding.

## Usage

The primary use case is to implement IR protocol decoding. 
Many common IR protocols are supported by the kernel already, but some of them are not.
The hook allows users to write their own protocol decoding.


## Context

The context for this program type is a pointer to an unsigned int that follows the format described [here](https://docs.kernel.org/userspace-api/media/rc/lirc-dev-intro.html#lirc-mode-mode2) for `LIRC_MODE_MODE2`.

## Attachment

Attachment is through the [`BPF_PROG_ATTACH`](../syscall/BPF_PROG_ATTACH.md) syscall.
The `target_fd` should be the LIRC device you want to attach the program to.
A maximum of 64 programs can be attached to a single LIRC device at a time.

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
    if (LIRC_IS_PULSE(*sample)) {
        unsigned int duration = LIRC_VALUE(*sample);

        if (duration & 0x10000)
            bpf_rc_keydown(sample, 0x40, duration & 0xffff, 0);
        if (duration & 0x20000)
            bpf_rc_pointer_rel(sample, (duration >> 8) & 0xff,
                       duration & 0xff);
    }

    return 0;
}

char _license[] SEC("license") = "GPL";
```

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
