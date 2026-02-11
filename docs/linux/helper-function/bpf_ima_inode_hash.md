---
title: "Helper Function 'bpf_ima_inode_hash'"
description: "This page documents the 'bpf_ima_inode_hash' eBPF helper function, including its definition, usage, program types that can use it, and examples."
---
# Helper function `bpf_ima_inode_hash`

<!-- [FEATURE_TAG](bpf_ima_inode_hash) -->
[:octicons-tag-24: v5.11](https://github.com/torvalds/linux/commit/27672f0d280a3f286a410a8db2004f46ace72a17)
<!-- [/FEATURE_TAG] -->

## Definition

> Copyright (c) 2015 The Libbpf Authors. All rights reserved.


<!-- [HELPER_FUNC_DEF] -->
Returns the stored IMA hash of the _inode_ (if it's available). If the hash is larger than _size_, then only _size_ bytes will be copied to _dst_

### Returns

The **hash_algo** is returned on success, **-EOPNOTSUPP** if IMA is disabled or **-EINVAL** if invalid arguments are passed.

`#!c static long (* const bpf_ima_inode_hash)(struct inode *inode, void *dst, __u32 size) = (void *) 161;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [`BPF_PROG_TYPE_LSM`](../program-type/BPF_PROG_TYPE_LSM.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example
The following program prints hash of files just before they are being executed.
Kernel command line is `ima_policy=tcb ima_hash=sha256`.

```c
#include "vmlinux.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>

static void print_sha256(__u8 *buf) {
    bpf_printk("IMA Hash Part 1: %02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x",
               buf[0], buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7], buf[8], buf[9], buf[10], buf[11]);
    bpf_printk("IMA Hash Part 2: %02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x",
               buf[12], buf[13], buf[14], buf[15], buf[16], buf[17], buf[18], buf[19], buf[20], buf[21], buf[22], buf[23]);
    bpf_printk("IMA Hash Part 3: %02x%02x%02x%02x%02x%02x%02x%02x",
               buf[24], buf[25], buf[26], buf[27], buf[28], buf[29], buf[30], buf[31]);
}
SEC("lsm.s/bprm_creds_for_exec")
int BPF_PROG(test_func, struct linux_binprm *b)
{
    // We are expecting SHA-256
    __u8 buf[32 / sizeof(__u8)] = {0};
    enum hash_algo algo = 0;

    algo = bpf_ima_inode_hash(b->file->f_inode, buf, sizeof(buf));
    if (algo < 0)
        return 0;
    /*just to showcase enum hash_algo*/
    if(algo != HASH_ALGO_SHA256){
        bpf_printk("algo mismatch");
        return 0;
    }
    bpf_printk("%s", b->filename);
    print_sha256(buf);
    return 0;
}

char __license[] SEC("license") = "GPL";
```

Output should be something like this:
```
<...>-20230   [008] ...11  9707.708954: bpf_trace_printk: /usr/bin/figlet
<...>-20230   [008] ...11  9707.708957: bpf_trace_printk: IMA Hash Part 1: 1748eeb53c9479fb923fb772
<...>-20230   [008] ...11  9707.708957: bpf_trace_printk: IMA Hash Part 2: c21bd9c9f5c27aa4e81c66cd
<...>-20230   [008] ...11  9707.708957: bpf_trace_printk: IMA Hash Part 3: 59886d7b339e70d0
<...>-20231   [000] ...11  9707.709873: bpf_trace_printk: /usr/bin/python3
<...>-20231   [000] ...11  9707.709876: bpf_trace_printk: IMA Hash Part 1: e59d0124ff06c248546876e0
<...>-20231   [000] ...11  9707.709876: bpf_trace_printk: IMA Hash Part 2: 1fcfb1ea3cda63534949f94a
<...>-20231   [000] ...11  9707.709877: bpf_trace_printk: IMA Hash Part 3: 9372bfcfe3bfc3f5
```
