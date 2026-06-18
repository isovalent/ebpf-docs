---
title: "Struct ops 'io_uring_bpf_ops'"
description: "This page documents the 'io_uring_bpf_ops' struct ops, its semantics, capabilities, and limitations."
---
# Struct ops `io_uring_bpf_ops`

[:octicons-tag-24: v7.1](https://github.com/torvalds/linux/commit/d0e437b76bd3c979ddaa6205f5e9ad3e0f95faef)

The `io_uring` BPF ops allows users to write BPF programs that can interact with a `io_uring`, replacing what normally would be userspace logic with a BPF program.

## Usage

The intended use cases for this are:

* Syscall avoidance. Instead of returning to the userspace for Completion Queue Entry processing, a part of the logic can be moved into BPF to avoid excessive number of syscalls.
* Access to in-kernel `io_uring` resources. For example, there are registered buffers that can't be directly accessed by the userspace, however we can give BPF the ability to peek at them. It can be used to take a look at in-buffer app level headers to decide what to do with data next and issuing IO using it.
* Smarter request ordering and linking. Request links are pretty limited and inflexible as they can't pass information from one request to another. With BPF we can peek at Completion Queue Entries and memory and compile a subsequent request.
* Feature semi-deprecation. It can be used to simplify handling of deprecated features by moving it into the callback out core `io_uring`. For example, it should be trivial to simulate `IOSQE_IO_DRAIN`. Another target could be request linking logic.
* It can serve as a base for custom algorithms and fine tuning. Often, it'd be impractical to introduce a generic feature because it's either niche or requires a lot of configuration. For example, there is support min-wait, however BPF can help to further fine tune it by doing it in multiple steps with different number of Completion Queue Entries / timeouts. Another feature people were asking about is allowing to over queue Submission Queue Entries but make the kernel to maintain a given Queue Depth.
* Smarter polling. NAPI polling is performed only once per syscall and then it switches to waiting. We can do smarter and intermix polling with waiting using the hook.

## Fields and ops

```c
struct io_uring_bpf_ops {
	int (*[loop_step](#loop_step))(struct iou_ctx *, struct iou_loop_params *lp);

	__u32 ring_fd;
	void *priv;
};;
```

### `loop_step`

[:octicons-tag-24: v7.1](https://github.com/torvalds/linux/commit/d0e437b76bd3c979ddaa6205f5e9ad3e0f95faef)

`#!c int (*loop_step)(struct iou_ctx *, struct iou_loop_params *lp)`

`loop_step` replaces the job of the [`io_uring_enter`](https://man7.org/linux/man-pages/man2/io_uring_enter.2.html) syscall.

It is called in a loop, it should return `IOU_LOOP_CONTINUE` to continue execution or `IOU_LOOP_STOP` to return to the user space. 

!!! note 
    The kernel may decide to prematurely terminate it as well, for example in case the process was signalled or killed.

The hook takes a structure with parameters (`lp`). It can be used to ask the kernel to wait for CQEs by setting 
`lp->cq_wait_idx` to the CQE index it wants to wait for. Spurious wake ups are possible and even likely, the callback
is expected to handle it. There will be more parameters in the future like timeout.

The [`bpf_io_uring_get_region`](../../kfuncs/bpf_io_uring_get_region.md) kfunc can be used to gain access to the shared memory region to read from or write to the rings. The [`bpf_io_uring_submit_sqes`](../../kfuncs/bpf_io_uring_submit_sqes.md) kfunc allows you to submit new submission queue entries after they have been written to the ring.

**Parameters**

`ctx`: An opaque context pointer which can be passed to kfuncs.
`lp`: Parameters shared by multiple calls of the loop, can contain both inputs and outputs.

## Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
