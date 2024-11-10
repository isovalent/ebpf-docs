---
title: "Map Type 'BPF_MAP_TYPE_PERF_EVENT_ARRAY'"
description: "This page documents the 'BPF_MAP_TYPE_PERF_EVENT_ARRAY' eBPF map type, including its definition, usage, program types that can use it, and examples."
---
# Map type `BPF_MAP_TYPE_PERF_EVENT_ARRAY`

<!-- [FEATURE_TAG](BPF_MAP_TYPE_PERF_EVENT_ARRAY) -->
[:octicons-tag-24: v4.3](https://github.com/torvalds/linux/commit/ea317b267e9d03a8241893aa176fba7661d07579)
<!-- [/FEATURE_TAG] -->

This is a specialized map type which holds file descriptors to perf events. It is most commonly used by eBPF programs to efficiently send large amounts of data from kernel space to userspace, but it also has other uses.

## Usage

There are two different ways to use this map type. The first and most popular way is to allows programs to send arbitrary data, for example debug messages, network flows, parts of the program context, etcetera.

The second way is with the [bpf_perf_event_read](../helper-function/bpf_perf_event_read.md) and [bpf_perf_event_read_value](../helper-function/bpf_perf_event_read_value.md) helpers to allow tracing programs to read performance counters.

### Data transfer

This usage scenario allows eBPF logic to piggy-back on the existing perf-subsystem implementation of ring-buffers to transfer data from the kernel to userspace. Sending data from an eBPF program is simple. Define a `BPF_MAP_TYPE_PERF_EVENT_ARRAY`, prepare some data to be sent, then call the [bpf_perf_event_output](../helper-function/bpf_perf_event_output.md) helper.

Receiving this data on the userspace side will need to do some initial setup work.

!!! note
    The following details are typically covered by loader libraries such as libbpf

After creating the map we can populate it with perf events. Such perf events are created with the [`perf_event_open`](https://man7.org/linux/man-pages/man2/perf_event_open.2.html) syscall. In our data transfer use-case we want to create a very specific type of perf event, with the following attributes: `PERF_SAMPLE_RAW` as `sample_type`, `PERF_TYPE_SOFTWARE` as `type`, and `PERF_COUNT_SW_BPF_OUTPUT` as `config`.

The following block would be an example of which attributes and parameters to use to invoke the syscall:

```c
struct perf_event_attr attr = {
    .type = PERF_TYPE_SOFTWARE,
    .size = sizeof(struct perf_event_attr),
    .config = PERF_COUNT_SW_BPF_OUTPUT,
    .watermark = true,
    .sample_type = PERF_SAMPLE_RAW,
    .wakeup_watermark = {watermark},
};

syscall(SYS_perf_event_open, 
    &attr,  /* struct perf_event_attr * */
    -1,     /* pid_t pid */
    {cpu}   /* int cpu */
    -1,     /* int group_fd */
    PERF_FLAG_FD_CLOEXEC /* unsigned long flags */
);
```

Note that `{watermark}` and `{cpu}` are not pre-set. In this example we set `.watermark` to true, this has the effect that the ringbuffer will buffer until `{watermark}` amount of bytes are ready before it signals userspace that there is data. This is not required, but does improve efficiently, if using the feature `{watermark}` should be tuned for balance between latency, performance and memory usage. The `{cpu}` value should be the index of the logical CPU for which you want to make the perf event. It should match the key of the map in which you intend to put the event.

After creating the perf event, it is recommended to make it non-blocking, if using used with a signaling mechanism such as `epoll`. This can be done by calling [`fcntl`](https://man7.org/linux/man-pages/man2/fcntl.2.html) with the `O_NONBLOCK` option on the file descriptor.

Next, we need to setup the actual ring-buffer which we will use to read data from the perf event. To do so we first have to request a shared memory region from the kernel using [`mmap`](https://man7.org/linux/man-pages/man2/mmap.2.html) `#!c mmap(nil, length, PROT_READ|PROT_WRITE, MAP_SHARED, fd, 0)`. `length` can be picked and tuned depending on the use case but should always be `1+2^n` memory pages and `n` should not be `0`. So Assuming a 4k page size, valid values would be `12288`, `20480`, `28672` and so on. The `fd` should be the file descriptor of the perf event.

The `mmap` call, if all is well, should return a pointer to a memory address. The first page will contain the following data structure populated by the perf subsystem:

??? abstract "C structure"
    ```c
    /*
    * Performance events:
    *
    *    Copyright (C) 2008-2009, Thomas Gleixner <tglx@linutronix.de>
    *    Copyright (C) 2008-2011, Red Hat, Inc., Ingo Molnar
    *    Copyright (C) 2008-2011, Red Hat, Inc., Peter Zijlstra
    *
    * Data type definitions, declarations, prototypes.
    *
    *    Started by: Thomas Gleixner and Ingo Molnar
    *
    * For licencing details see kernel-base/COPYING
    */
    /*
    * Structure of the page that can be mapped via mmap
    */
    struct perf_event_mmap_page {
        __u32	version;		/* version number of this structure */
        __u32	compat_version;		/* lowest version this is compat with */

        /*
        * Bits needed to read the hw events in user-space.
        *
        *   u32 seq, time_mult, time_shift, index, width;
        *   u64 count, enabled, running;
        *   u64 cyc, time_offset;
        *   s64 pmc = 0;
        *
        *   do {
        *     seq = pc->lock;
        *     barrier()
        *
        *     enabled = pc->time_enabled;
        *     running = pc->time_running;
        *
        *     if (pc->cap_usr_time && enabled != running) {
        *       cyc = rdtsc();
        *       time_offset = pc->time_offset;
        *       time_mult   = pc->time_mult;
        *       time_shift  = pc->time_shift;
        *     }
        *
        *     index = pc->index;
        *     count = pc->offset;
        *     if (pc->cap_user_rdpmc && index) {
        *       width = pc->pmc_width;
        *       pmc = rdpmc(index - 1);
        *     }
        *
        *     barrier();
        *   } while (pc->lock != seq);
        *
        * NOTE: for obvious reason this only works on self-monitoring
        *       processes.
        */
        __u32	lock;			/* seqlock for synchronization */
        __u32	index;			/* hardware event identifier */
        __s64	offset;			/* add to hardware event value */
        __u64	time_enabled;		/* time event active */
        __u64	time_running;		/* time event on cpu */
        union {
            __u64	capabilities;
            struct {
                __u64	cap_bit0		: 1, /* Always 0, deprecated, see commit 860f085b74e9 */
                    cap_bit0_is_deprecated	: 1, /* Always 1, signals that bit 0 is zero */

                    cap_user_rdpmc		: 1, /* The RDPMC instruction can be used to read counts */
                    cap_user_time		: 1, /* The time_{shift,mult,offset} fields are used */
                    cap_user_time_zero	: 1, /* The time_zero field is used */
                    cap_user_time_short	: 1, /* the time_{cycle,mask} fields are used */
                    cap_____res		: 58;
            };
        };

        /*
        * If cap_user_rdpmc this field provides the bit-width of the value
        * read using the rdpmc() or equivalent instruction. This can be used
        * to sign extend the result like:
        *
        *   pmc <<= 64 - width;
        *   pmc >>= 64 - width; // signed shift right
        *   count += pmc;
        */
        __u16	pmc_width;

        /*
        * If cap_usr_time the below fields can be used to compute the time
        * delta since time_enabled (in ns) using rdtsc or similar.
        *
        *   u64 quot, rem;
        *   u64 delta;
        *
        *   quot = (cyc >> time_shift);
        *   rem = cyc & (((u64)1 << time_shift) - 1);
        *   delta = time_offset + quot * time_mult +
        *              ((rem * time_mult) >> time_shift);
        *
        * Where time_offset,time_mult,time_shift and cyc are read in the
        * seqcount loop described above. This delta can then be added to
        * enabled and possible running (if index), improving the scaling:
        *
        *   enabled += delta;
        *   if (index)
        *     running += delta;
        *
        *   quot = count / running;
        *   rem  = count % running;
        *   count = quot * enabled + (rem * enabled) / running;
        */
        __u16	time_shift;
        __u32	time_mult;
        __u64	time_offset;
        /*
        * If cap_usr_time_zero, the hardware clock (e.g. TSC) can be calculated
        * from sample timestamps.
        *
        *   time = timestamp - time_zero;
        *   quot = time / time_mult;
        *   rem  = time % time_mult;
        *   cyc = (quot << time_shift) + (rem << time_shift) / time_mult;
        *
        * And vice versa:
        *
        *   quot = cyc >> time_shift;
        *   rem  = cyc & (((u64)1 << time_shift) - 1);
        *   timestamp = time_zero + quot * time_mult +
        *               ((rem * time_mult) >> time_shift);
        */
        __u64	time_zero;

        __u32	size;			/* Header size up to __reserved[] fields. */
        __u32	__reserved_1;

        /*
        * If cap_usr_time_short, the hardware clock is less than 64bit wide
        * and we must compute the 'cyc' value, as used by cap_usr_time, as:
        *
        *   cyc = time_cycles + ((cyc - time_cycles) & time_mask)
        *
        * NOTE: this form is explicitly chosen such that cap_usr_time_short
        *       is a correction on top of cap_usr_time, and code that doesn't
        *       know about cap_usr_time_short still works under the assumption
        *       the counter doesn't wrap.
        */
        __u64	time_cycles;
        __u64	time_mask;

            /*
            * Hole for extension of the self monitor capabilities
            */

        __u8	__reserved[116*8];	/* align to 1k. */

        /*
        * Control data for the mmap() data buffer.
        *
        * User-space reading the @data_head value should issue an smp_rmb(),
        * after reading this value.
        *
        * When the mapping is PROT_WRITE the @data_tail value should be
        * written by userspace to reflect the last read data, after issueing
        * an smp_mb() to separate the data read from the ->data_tail store.
        * In this case the kernel will not over-write unread data.
        *
        * See perf_output_put_handle() for the data ordering.
        *
        * data_{offset,size} indicate the location and size of the perf record
        * buffer within the mmapped area.
        */
        __u64   data_head;		/* head in the data section */
        __u64	data_tail;		/* user-space written tail */
        __u64	data_offset;		/* where the buffer starts */
        __u64	data_size;		/* data buffer size */

        /*
        * AUX area is defined by aux_{offset,size} fields that should be set
        * by the userspace, so that
        *
        *   aux_offset >= data_offset + data_size
        *
        * prior to mmap()ing it. Size of the mmap()ed area should be aux_size.
        *
        * Ring buffer pointers aux_{head,tail} have the same semantics as
        * data_{head,tail} and same ordering rules apply.
        */
        __u64	aux_head;
        __u64	aux_tail;
        __u64	aux_offset;
        __u64	aux_size;
    };
    ```

This structure contains a lot of additional data, but for consuming the data, we interested in the `perf_event_mmap_page->head`, `perf_event_mmap_page->tail`, `perf_event_mmap_page->data_offset` and `perf_event_mmap_page->data_size` fields.

The `data_offset` and `data_size` indicate the offset from the `addr` pointer where the ring-buffer data starts, and `data_size` where it ends. The region between these two locations will be written to by the kernel. To avoid race conditions, we should avoid reading and writing the same location at the same time. The `head` and `tail` are used to coordinate access. The `head` will be written to by the kernel and should be read by userspace, it moves if new data is available. The `tail` will be written to by userspace to indicate data has been consumed. So if `head` and `tail` are equal, no data is pending. If `head` == `tail-1` then the buffer is full.

To recap, we create a perf event for every logical CPU, and every perf event gets its own ring-buffer.

Once setup, userspace can busy-poll the ring-buffers, but this might cause significant overhead or data loss during spikes. It is recommended to use [`epoll`](https://man7.org/linux/man-pages/man7/epoll.7.html) to setup signalling in combination with the watermark settings for optimal result. Allowing a reader thread to wake-up when there is data, read from all buffers indicated to have pending data by `epoll` and then go back to blocking mode/sleep. This avoids CPU usage when nothing happens and wakes up the thread more frequently when more data is sent. 

Depending on the use case, a author might want to dedicate more threads to reading concurrently, not use watermarks, or still do busy polling, it all depends on optimizing throughput, latency (between eBPF sending and userspace consuming), memory usage, and CPU usage.

Every entry on the ring-buffer will start with the same header:

```c
struct perf_event_header {
        __u32   type;
        __u16   misc;
        __u16   size;
};
```

The event `type` in our particular use case will be `PERF_RECORD_SAMPLE` or `PERF_RECORD_LOST`. `PERF_RECORD_SAMPLE` indicating that there is an actual sample after this header. 

```c
struct perf_event_sample {
	struct perf_event_header header;
	__u64 time;
	__u32 size;
	unsigned char data[];
};
```

And `PERF_RECORD_LOST` indicating that there is a record lost header following the perf event header.

```c
struct {
    struct perf_event_header header;
    u64    id;
    u64    lost;
    struct sample_id sample_id;
};
```

The `lost` field indicates how many samples have been lost. This might happen when the ringbuffer is full or nearly full. The kernel will keep track of the amount of samples it can't write and will populate the ring buffer with record lost messages after there is room on the ringbuffer again.

For further reading, checkout [`tools/perf/design.txt`](https://github.com/torvalds/linux/blob/v6.2/tools/perf/design.txt) which goes into the design and logic behind the above instructions.

### Performance counters

A less common use case is to transfer perf events to eBPF from which programs read. This allows tracing programs to record, for example hardware counters at very specific points such as kprobes or tracepoints to measure performance.

To use the map in this mode, the loader creates a perf event for each CPU, using the [`perf_event_open`](https://man7.org/linux/man-pages/man2/perf_event_open.2.html) syscall with the desired settings, making sure to vary the `cpu` argument.

On the eBPF side, your programs should be able to use [bpf_perf_event_read](../helper-function/bpf_perf_event_read.md) and/or [bpf_perf_event_read_value](../helper-function/bpf_perf_event_read_value.md) to read values from the events given to it by userspace.

## Attributes

Both the [`key_size`](../syscall/BPF_MAP_CREATE.md#key_size) and [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) must be exactly `4`. 

While the [`value_size`](../syscall/BPF_MAP_CREATE.md#value_size) is essentially unrestricted, the  must always be `4` indicating the key is a 32-bit unsigned integer.

The [`max_entries`](../syscall/BPF_MAP_CREATE.md#max_entries) should at least as large as the number of logical CPUs. This number can be discovered in a number of ways, including [`nproc`](https://man7.org/linux/man-pages/man1/nproc.1.html), [`lscpu`](https://linux.die.net/man/1/lscpu), [`/proc/cpuinfo`](https://www.kernel.org/doc/Documentation/admin-guide/cputopology.rst)

## Syscall commands

The following syscall commands work with this map type:

* [`BPF_MAP_LOOKUP_ELEM`](../syscall/BPF_MAP_LOOKUP_ELEM.md)
* [`BPF_MAP_UPDATE_ELEM`](../syscall/BPF_MAP_UPDATE_ELEM.md)
* [`BPF_MAP_GET_NEXT_KEY`](../syscall/BPF_MAP_GET_NEXT_KEY.md)

## Helper functions

<!-- DO NOT EDIT MANUALLY -->
<!-- [MAP_HELPER_FUNC_REF] -->
 * [`bpf_perf_event_output`](../helper-function/bpf_perf_event_output.md)
 * [`bpf_perf_event_read`](../helper-function/bpf_perf_event_read.md)
 * [`bpf_perf_event_read_value`](../helper-function/bpf_perf_event_read_value.md)
 * [`bpf_skb_output`](../helper-function/bpf_skb_output.md)
 * [`bpf_xdp_output`](../helper-function/bpf_xdp_output.md)
<!-- [/MAP_HELPER_FUNC_REF] -->

## Flags

The following flags are supported by this map type.

### `BPF_F_NUMA_NODE`

[:octicons-tag-24: v4.14](https://github.com/torvalds/linux/commit/96eabe7a40aa17e613cf3db2c742ee8b1fc764d0)

When set, the [`numa_node`](../syscall/BPF_MAP_CREATE.md#numa_node) attribute is respected during map creation.

### `BPF_F_PRESERVE_ELEMS`

<!-- [FEATURE_TAG](BPF_F_PRESERVE_ELEMS) -->
[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/792caccc4526bb489e054f9ab61d7c024b15dea2)
<!-- [/FEATURE_TAG] -->

By default, all unread perf events are cleared when the original map file descriptor is closed, even if the map still exists. Setting this flag will make it so any pending elements will stay until explicitly removed or the map is freed. This makes sharing the perf event array between userspace programs easier.

### `BPF_F_RDONLY`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/6e71b04a82248ccf13a94b85cbc674a9fefe53f5)

Setting this flag will make it so the map can only be read via the [syscall](../syscall/index.md) interface, but not written to.

For details please check the [generic description](../syscall/BPF_MAP_CREATE.md#bpf_f_rdonly).

### `BPF_F_WRONLY`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/6e71b04a82248ccf13a94b85cbc674a9fefe53f5)

Setting this flag will make it so the map can only be written to via the [syscall](../syscall/index.md) interface, but not read from.


