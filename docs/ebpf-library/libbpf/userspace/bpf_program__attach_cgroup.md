---
title: "Libbpf userspace function 'bpf_program__attach_cgroup'"
description: "This page documents the 'bpf_program__attach_cgroup' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `bpf_program__attach_cgroup`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.8](https://github.com/libbpf/libbpf/releases/tag/v0.0.8)
<!-- [/LIBBPF_TAG] -->

Attach a [`BPF_PROG_TYPE_CGROUP_*`](../../../linux/program-type/index.md#cgroup-program-types) program type or an [`BPF_PROG_TYPE_LSM`](../../../linux/program-type/BPF_PROG_TYPE_LSM.md) program type using the [`BPF_LSM_CGROUP`](../../../linux/syscall/BPF_LINK_CREATE.md#bpf_lsm_cgroup) attachment type.

## Definition

`#!c struct bpf_link * bpf_program__attach_cgroup(const struct bpf_program *prog, int cgroup_fd);`

**Parameters**

- `prog`: BPF program to attach
- `cgroup_fd`: file descriptor of the cgroup to attach the program to

**Return**

Reference to the newly created BPF link; or `NULL` is returned on error, error code is stored in [`errno`](https://man7.org/linux/man-pages/man3/errno.3.html)

## Usage

[`bpf_program__attach_cgroup`](bpf_program__attach_cgroup.md) attaches a BPF program to a given cGroup to enforce fine-grained, per-cGroup policies.

### Example


```c
int main(int argc, char **argv)
{
	if (argc != 2) {
		fprintf(stderr, "Usage: %s <cgroup_path> (e.g '/sys/fs/cgroup/system.slice/'\n", argv[0]);
		return 1;
	}

	// Obtain cgroup fd
	const char *cgroup_path = argv[1];
	int cgroup_fd = open(cgroup_path, O_RDONLY);
	if (cgroup_fd < 0) {
		perror("open");
		return 1;
	}

	LIBBPF_OPTS(bpf_object_open_opts, opts);

	struct perf_buffer *pb = NULL;
	struct file_lsm_bpf *obj;
	int err;

	obj = file_lsm_bpf__open_opts(&opts);
	err = bpf_object__load(obj->obj);
	struct bpf_link *link = bpf_program__attach_cgroup(obj->progs.lsm_file, cgroup_fd);
	if (!link) {
		fprintf(stderr, "failed to attach BPF program to cgroup\n");
		goto cleanup;
	}
}
```
