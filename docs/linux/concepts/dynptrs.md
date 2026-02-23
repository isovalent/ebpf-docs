# Dynptrs (Dynamic Pointers)

A "dynptr" or "dynamic pointer" is a concept in the Linux eBPF verifier. It is a pointer with additional metadata so that certain safety check can be performed at runtime. This is useful for situations where it might be difficult to statically prove the safety of certain actions.

For example, consider the situation where you have a map with a setting that instructs a program to read or write an arbitrary spot in a packet. Since the map value can be any value and the size of the packet is variable as well, it is challenging to statically prove all cases (though likely not impossible). By using dynptrs, we shift the burden of proof to runtime. If the program tries to access memory outside of the packet, the helper function or kfunc will return an error instead.

To a eBPF program, a dynptr is just an opaque pointer. The verifier will not allow the program to dereference it directly. Instead, the program must use helper functions or kfuncs to access the memory it points to. These functions will perform the necessary safety checks.

## Helper functions and kfuncs

The following functions create or manipulate dynptrs:

* [`bpf_dynptr_from_mem`](../helper-function/bpf_dynptr_from_mem.md) - Creates a dynptr for a map value or global variable.
* [`bpf_dynptr_read`](../helper-function/bpf_dynptr_read.md) - Attempts to read from a dynptr.
* [`bpf_dynptr_write`](../helper-function/bpf_dynptr_write.md) - Attempts to write to a dynptr.
* [`bpf_dynptr_data`](../helper-function/bpf_dynptr_data.md) - Returns a pointer to the underlying data of a dynptr of a given length at a given offset. The verifier knows about the length and offset and will enforce bounds checks statically. 
* [`bpf_ringbuf_reserve_dynptr`](../helper-function/bpf_ringbuf_reserve_dynptr.md) - Reserves a sample in a ring buffer as dynptr, allowing for runtime determined size samples.
* [`bpf_dynptr_adjust`](../kfuncs/bpf_dynptr_adjust.md) - Adjusts the dynptr to a new offset and length within the existing dynptr. This proves more slice-like functionality.
* [`bpf_dynptr_size`](../kfuncs/bpf_dynptr_size.md) - Returns the number of usable bytes inside the dynptr.
* [`bpf_dynptr_from_skb`](../kfuncs/bpf_dynptr_from_skb.md) - Creates a dynptr for packet data inside of a socket buffer. Depending on the program type, the dynptr might be read-only or NULL (if ctx->data is NULL). This dynptr gives access to both the linear and non-linear parts of the packet.
* [`bpf_dynptr_from_xdp`](../kfuncs/bpf_dynptr_from_xdp.md) - Creates a dynptr for packet data inside of XDP metadata. This dynptr gives access to the linear part and fragment data of the packet.
* [`bpf_dynptr_is_null`](../kfuncs/bpf_dynptr_is_null.md) - Returns `true` if the dynptr is NULL.
* [`bpf_dynptr_is_rdonly`](../kfuncs/bpf_dynptr_is_rdonly.md) - Returns `true` if the dynptr is read-only.
* [`bpf_dynptr_slice`](../kfuncs/bpf_dynptr_slice.md) - Similar to `bpf_dynptr_data`, but works with XDP and SKB derived dynptrs where `bpf_dynptr_data` isn't supported. It returns a read-only dynptr to a slice of the original dynptr.
* [`bpf_dynptr_slice_rdwr`](../kfuncs/bpf_dynptr_slice_rdwr.md) - Similar to `bpf_dynptr_slice`, but returns a read-write dynptr.
* [`bpf_dynptr_clone`](../kfuncs/bpf_dynptr_clone.md) - Clones a dynptr. The new dynptr points to the same underlying data and has the same metadata as the original dynptr.
* [`bpf_dynptr_memset`](../kfuncs/bpf_dynptr_memset.md) - Fill dynptr memory with a constant byte.
* [`bpf_probe_read_user_dynptr`](../kfuncs/bpf_probe_read_user_dynptr.md) - Probes user-space data into a dynptr
* [`bpf_probe_read_kernel_dynptr`](../kfuncs/bpf_probe_read_kernel_dynptr.md) - Probes kernel-space data into a dynptr
* [`bpf_probe_read_user_str_dynptr`](../kfuncs/bpf_probe_read_user_str_dynptr.md) - Probes user-space string into a dynptr
* [`bpf_probe_read_kernel_str_dynptr`](../kfuncs/bpf_probe_read_kernel_str_dynptr.md) - probes kernel-space string into a dynptr
* [`bpf_copy_from_user_dynptr`](../kfuncs/bpf_copy_from_user_dynptr.md) - Sleepable, copies user-space data into a dynptr for the current task
* [`bpf_copy_from_user_str_dynptr`](../kfuncs/bpf_copy_from_user_str_dynptr.md) - Sleepable, copies user-space string into a dynptr for the current task
* [`bpf_copy_from_user_task_dynptr`](../kfuncs/bpf_copy_from_user_task_dynptr.md) - Sleepable, copies user-space data of the task into a dynptr
* [`bpf_copy_from_user_task_str_dynptr`](../kfuncs/bpf_copy_from_user_task_str_dynptr.md) - Sleepable, copies user-space string of the task into a dynptr
* [`bpf_dynptr_from_file`](../kfuncs/bpf_dynptr_from_file.md) - Creates a dynptr for a file.
* [`bpf_dynptr_file_discard`](../kfuncs/bpf_dynptr_file_discard.md) - Releases a file dynptr.

The following functions are not dynptr centric, but do require dynptrs in their arguments:

* [`bpf_ringbuf_submit_dynptr`](../helper-function/bpf_ringbuf_submit_dynptr.md) - Submits a dynptr to a ring buffer.
* [`bpf_ringbuf_discard_dynptr`](../helper-function/bpf_ringbuf_discard_dynptr.md) - Discards a dynptr from a ring buffer.
* [`bpf_user_ringbuf_drain`](../helper-function/bpf_user_ringbuf_drain.md) - Drains samples from a user ring buffer and invokes a callback for each sample, with a dynptr containing the sample data.
* [`bpf_crypto_decrypt`](../kfuncs/bpf_crypto_decrypt.md) - Decrypts a buffer using a crypto context and IV data. The source, destination, and IV buffers are dynptrs.
* [`bpf_crypto_encrypt`](../kfuncs/bpf_crypto_encrypt.md) - Encrypts a buffer using a crypto context and IV data. The source, destination, and IV buffers are dynptrs.
* [`bpf_get_file_xattr`](../kfuncs/bpf_get_file_xattr.md) - Retrieves an extended attribute from a file descriptor. A dynptr must be provided to which the attribute value will be written.
* [`bpf_get_fsverity_digest`](../kfuncs/bpf_get_fsverity_digest.md) - Retrieves the `fs-verity` digest of a file. A dynptr must be provided to which the digest will be written.
* [`bpf_verify_pkcs7_signature`](../kfuncs/bpf_verify_pkcs7_signature.md) - Verifies a <nospell>PKCS7</nospell> signature. The signature and data buffers are dynptrs.

## Examples

Original source [link](https://elixir.bootlin.com/linux/v6.10.7/source/tools/testing/selftests/bpf/progs/dynptr_success.c)
```c
// SPDX-License-Identifier: GPL-2.0
/* Copyright (c) 2022 Facebook */

#include <string.h>
#include <stdbool.h>
#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>
#include "bpf_misc.h"
#include "bpf_kfuncs.h"
#include "errno.h"

char _license[] SEC("license") = "GPL";

int pid, err, val;

struct sample {
	int pid;
	int seq;
	long value;
	char comm[16];
};

struct {
	__uint(type, BPF_MAP_TYPE_RINGBUF);
	__uint(max_entries, 4096);
} ringbuf SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_ARRAY);
	__uint(max_entries, 1);
	__type(key, __u32);
	__type(value, __u32);
} array_map SEC(".maps");

SEC("?tp/syscalls/sys_enter_nanosleep")
int test_read_write(void *ctx)
{
	char write_data[64] = "hello there, world!!";
	char read_data[64] = {};
	struct bpf_dynptr ptr;
	int i;

	if (bpf_get_current_pid_tgid() >> 32 != pid)
		return 0;

	bpf_ringbuf_reserve_dynptr(&ringbuf, sizeof(write_data), 0, &ptr);

	/* Write data into the dynptr */
	err = bpf_dynptr_write(&ptr, 0, write_data, sizeof(write_data), 0);

	/* Read the data that was written into the dynptr */
	err = err ?: bpf_dynptr_read(read_data, sizeof(read_data), &ptr, 0, 0);

	/* Ensure the data we read matches the data we wrote */
	for (i = 0; i < sizeof(read_data); i++) {
		if (read_data[i] != write_data[i]) {
			err = 1;
			break;
		}
	}

	bpf_ringbuf_discard_dynptr(&ptr, 0);
	return 0;
}

SEC("?tp/syscalls/sys_enter_nanosleep")
int test_dynptr_data(void *ctx)
{
	__u32 key = 0, val = 235, *map_val;
	struct bpf_dynptr ptr;
	__u32 map_val_size;
	void *data;

	map_val_size = sizeof(*map_val);

	if (bpf_get_current_pid_tgid() >> 32 != pid)
		return 0;

	bpf_map_update_elem(&array_map, &key, &val, 0);

	map_val = bpf_map_lookup_elem(&array_map, &key);
	if (!map_val) {
		err = 1;
		return 0;
	}

	bpf_dynptr_from_mem(map_val, map_val_size, 0, &ptr);

	/* Try getting a data slice that is out of range */
	data = bpf_dynptr_data(&ptr, map_val_size + 1, 1);
	if (data) {
		err = 2;
		return 0;
	}

	/* Try getting more bytes than available */
	data = bpf_dynptr_data(&ptr, 0, map_val_size + 1);
	if (data) {
		err = 3;
		return 0;
	}

	data = bpf_dynptr_data(&ptr, 0, sizeof(__u32));
	if (!data) {
		err = 4;
		return 0;
	}

	*(__u32 *)data = 999;

	err = bpf_probe_read_kernel(&val, sizeof(val), data);
	if (err)
		return 0;

	if (val != *(int *)data)
		err = 5;

	return 0;
}

static int ringbuf_callback(__u32 index, void *data)
{
	struct sample *sample;

	struct bpf_dynptr *ptr = (struct bpf_dynptr *)data;

	sample = bpf_dynptr_data(ptr, 0, sizeof(*sample));
	if (!sample)
		err = 2;
	else
		sample->pid += index;

	return 0;
}

SEC("?tp/syscalls/sys_enter_nanosleep")
int test_ringbuf(void *ctx)
{
	struct bpf_dynptr ptr;
	struct sample *sample;

	if (bpf_get_current_pid_tgid() >> 32 != pid)
		return 0;

	val = 100;

	/* check that you can reserve a dynamic size reservation */
	err = bpf_ringbuf_reserve_dynptr(&ringbuf, val, 0, &ptr);

	sample = err ? NULL : bpf_dynptr_data(&ptr, 0, sizeof(*sample));
	if (!sample) {
		err = 1;
		goto done;
	}

	sample->pid = 10;

	/* Can pass dynptr to callback functions */
	bpf_loop(10, ringbuf_callback, &ptr, 0);

	if (sample->pid != 55)
		err = 2;

done:
	bpf_ringbuf_discard_dynptr(&ptr, 0);
	return 0;
}

SEC("?cgroup_skb/egress")
int test_skb_readonly(struct __sk_buff *skb)
{
	__u8 write_data[2] = {1, 2};
	struct bpf_dynptr ptr;
	int ret;

	if (bpf_dynptr_from_skb(skb, 0, &ptr)) {
		err = 1;
		return 1;
	}

	/* since cgroup skbs are read only, writes should fail */
	ret = bpf_dynptr_write(&ptr, 0, write_data, sizeof(write_data), 0);
	if (ret != -EINVAL) {
		err = 2;
		return 1;
	}

	return 1;
}

SEC("?cgroup_skb/egress")
int test_dynptr_skb_data(struct __sk_buff *skb)
{
	struct bpf_dynptr ptr;
	__u64 *data;

	if (bpf_dynptr_from_skb(skb, 0, &ptr)) {
		err = 1;
		return 1;
	}

	/* This should return NULL. Must use bpf_dynptr_slice API */
	data = bpf_dynptr_data(&ptr, 0, 1);
	if (data) {
		err = 2;
		return 1;
	}

	return 1;
}

SEC("tp/syscalls/sys_enter_nanosleep")
int test_adjust(void *ctx)
{
	struct bpf_dynptr ptr;
	__u32 bytes = 64;
	__u32 off = 10;
	__u32 trim = 15;

	if (bpf_get_current_pid_tgid() >> 32 != pid)
		return 0;

	err = bpf_ringbuf_reserve_dynptr(&ringbuf, bytes, 0, &ptr);
	if (err) {
		err = 1;
		goto done;
	}

	if (bpf_dynptr_size(&ptr) != bytes) {
		err = 2;
		goto done;
	}

	/* Advance the dynptr by off */
	err = bpf_dynptr_adjust(&ptr, off, bpf_dynptr_size(&ptr));
	if (err) {
		err = 3;
		goto done;
	}

	if (bpf_dynptr_size(&ptr) != bytes - off) {
		err = 4;
		goto done;
	}

	/* Trim the dynptr */
	err = bpf_dynptr_adjust(&ptr, off, 15);
	if (err) {
		err = 5;
		goto done;
	}

	/* Check that the size was adjusted correctly */
	if (bpf_dynptr_size(&ptr) != trim - off) {
		err = 6;
		goto done;
	}

done:
	bpf_ringbuf_discard_dynptr(&ptr, 0);
	return 0;
}

SEC("tp/syscalls/sys_enter_nanosleep")
int test_adjust_err(void *ctx)
{
	char write_data[45] = "hello there, world!!";
	struct bpf_dynptr ptr;
	__u32 size = 64;
	__u32 off = 20;

	if (bpf_get_current_pid_tgid() >> 32 != pid)
		return 0;

	if (bpf_ringbuf_reserve_dynptr(&ringbuf, size, 0, &ptr)) {
		err = 1;
		goto done;
	}

	/* Check that start can't be greater than end */
	if (bpf_dynptr_adjust(&ptr, 5, 1) != -EINVAL) {
		err = 2;
		goto done;
	}

	/* Check that start can't be greater than size */
	if (bpf_dynptr_adjust(&ptr, size + 1, size + 1) != -ERANGE) {
		err = 3;
		goto done;
	}

	/* Check that end can't be greater than size */
	if (bpf_dynptr_adjust(&ptr, 0, size + 1) != -ERANGE) {
		err = 4;
		goto done;
	}

	if (bpf_dynptr_adjust(&ptr, off, size)) {
		err = 5;
		goto done;
	}

	/* Check that you can't write more bytes than available into the dynptr
	 * after you've adjusted it
	 */
	if (bpf_dynptr_write(&ptr, 0, &write_data, sizeof(write_data), 0) != -E2BIG) {
		err = 6;
		goto done;
	}

	/* Check that even after adjusting, submitting/discarding
	 * a ringbuf dynptr works
	 */
	bpf_ringbuf_submit_dynptr(&ptr, 0);
	return 0;

done:
	bpf_ringbuf_discard_dynptr(&ptr, 0);
	return 0;
}

SEC("tp/syscalls/sys_enter_nanosleep")
int test_zero_size_dynptr(void *ctx)
{
	char write_data = 'x', read_data;
	struct bpf_dynptr ptr;
	__u32 size = 64;

	if (bpf_get_current_pid_tgid() >> 32 != pid)
		return 0;

	if (bpf_ringbuf_reserve_dynptr(&ringbuf, size, 0, &ptr)) {
		err = 1;
		goto done;
	}

	/* After this, the dynptr has a size of 0 */
	if (bpf_dynptr_adjust(&ptr, size, size)) {
		err = 2;
		goto done;
	}

	/* Test that reading + writing non-zero bytes is not ok */
	if (bpf_dynptr_read(&read_data, sizeof(read_data), &ptr, 0, 0) != -E2BIG) {
		err = 3;
		goto done;
	}

	if (bpf_dynptr_write(&ptr, 0, &write_data, sizeof(write_data), 0) != -E2BIG) {
		err = 4;
		goto done;
	}

	/* Test that reading + writing 0 bytes from a 0-size dynptr is ok */
	if (bpf_dynptr_read(&read_data, 0, &ptr, 0, 0)) {
		err = 5;
		goto done;
	}

	if (bpf_dynptr_write(&ptr, 0, &write_data, 0, 0)) {
		err = 6;
		goto done;
	}

	err = 0;

done:
	bpf_ringbuf_discard_dynptr(&ptr, 0);
	return 0;
}

SEC("tp/syscalls/sys_enter_nanosleep")
int test_dynptr_is_null(void *ctx)
{
	struct bpf_dynptr ptr1;
	struct bpf_dynptr ptr2;
	__u64 size = 4;

	if (bpf_get_current_pid_tgid() >> 32 != pid)
		return 0;

	/* Pass in invalid flags, get back an invalid dynptr */
	if (bpf_ringbuf_reserve_dynptr(&ringbuf, size, 123, &ptr1) != -EINVAL) {
		err = 1;
		goto exit_early;
	}

	/* Test that the invalid dynptr is null */
	if (!bpf_dynptr_is_null(&ptr1)) {
		err = 2;
		goto exit_early;
	}

	/* Get a valid dynptr */
	if (bpf_ringbuf_reserve_dynptr(&ringbuf, size, 0, &ptr2)) {
		err = 3;
		goto exit;
	}

	/* Test that the valid dynptr is not null */
	if (bpf_dynptr_is_null(&ptr2)) {
		err = 4;
		goto exit;
	}

exit:
	bpf_ringbuf_discard_dynptr(&ptr2, 0);
exit_early:
	bpf_ringbuf_discard_dynptr(&ptr1, 0);
	return 0;
}

SEC("cgroup_skb/egress")
int test_dynptr_is_rdonly(struct __sk_buff *skb)
{
	struct bpf_dynptr ptr1;
	struct bpf_dynptr ptr2;
	struct bpf_dynptr ptr3;

	/* Pass in invalid flags, get back an invalid dynptr */
	if (bpf_dynptr_from_skb(skb, 123, &ptr1) != -EINVAL) {
		err = 1;
		return 0;
	}

	/* Test that an invalid dynptr is_rdonly returns false */
	if (bpf_dynptr_is_rdonly(&ptr1)) {
		err = 2;
		return 0;
	}

	/* Get a read-only dynptr */
	if (bpf_dynptr_from_skb(skb, 0, &ptr2)) {
		err = 3;
		return 0;
	}

	/* Test that the dynptr is read-only */
	if (!bpf_dynptr_is_rdonly(&ptr2)) {
		err = 4;
		return 0;
	}

	/* Get a read-writeable dynptr */
	if (bpf_ringbuf_reserve_dynptr(&ringbuf, 64, 0, &ptr3)) {
		err = 5;
		goto done;
	}

	/* Test that the dynptr is read-only */
	if (bpf_dynptr_is_rdonly(&ptr3)) {
		err = 6;
		goto done;
	}

done:
	bpf_ringbuf_discard_dynptr(&ptr3, 0);
	return 0;
}

SEC("cgroup_skb/egress")
int test_dynptr_clone(struct __sk_buff *skb)
{
	struct bpf_dynptr ptr1;
	struct bpf_dynptr ptr2;
	__u32 off = 2, size;

	/* Get a dynptr */
	if (bpf_dynptr_from_skb(skb, 0, &ptr1)) {
		err = 1;
		return 0;
	}

	if (bpf_dynptr_adjust(&ptr1, off, bpf_dynptr_size(&ptr1))) {
		err = 2;
		return 0;
	}

	/* Clone the dynptr */
	if (bpf_dynptr_clone(&ptr1, &ptr2)) {
		err = 3;
		return 0;
	}

	size = bpf_dynptr_size(&ptr1);

	/* Check that the clone has the same size and rd-only */
	if (bpf_dynptr_size(&ptr2) != size) {
		err = 4;
		return 0;
	}

	if (bpf_dynptr_is_rdonly(&ptr2) != bpf_dynptr_is_rdonly(&ptr1)) {
		err = 5;
		return 0;
	}

	/* Advance and trim the original dynptr */
	bpf_dynptr_adjust(&ptr1, 5, 5);

	/* Check that only original dynptr was affected, and the clone wasn't */
	if (bpf_dynptr_size(&ptr2) != size) {
		err = 6;
		return 0;
	}

	return 0;
}

SEC("?cgroup_skb/egress")
int test_dynptr_skb_no_buff(struct __sk_buff *skb)
{
	struct bpf_dynptr ptr;
	__u64 *data;

	if (bpf_dynptr_from_skb(skb, 0, &ptr)) {
		err = 1;
		return 1;
	}

	/* This may return NULL. SKB may require a buffer */
	data = bpf_dynptr_slice(&ptr, 0, NULL, 1);

	return !!data;
}

SEC("?cgroup_skb/egress")
int test_dynptr_skb_strcmp(struct __sk_buff *skb)
{
	struct bpf_dynptr ptr;
	char *data;

	if (bpf_dynptr_from_skb(skb, 0, &ptr)) {
		err = 1;
		return 1;
	}

	/* This may return NULL. SKB may require a buffer */
	data = bpf_dynptr_slice(&ptr, 0, NULL, 10);
	if (data) {
		bpf_strncmp(data, 10, "foo");
		return 1;
	}

	return 1;
}
```
