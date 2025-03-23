---
title: "Libbpf userspace function 'libbpf_set_strict_mode'"
description: "This page documents the 'libbpf_set_strict_mode' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `libbpf_set_strict_mode`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.5.0](https://github.com/libbpf/libbpf/releases/tag/v0.5.0)
<!-- [/LIBBPF_TAG] -->

Set the "strict mode" of libbpf.

## Definition

`#!c int libbpf_set_strict_mode(enum libbpf_strict_mode mode);`

**Parameters**

- `mode`: strict mode to set

**Return**

`0`, on success; negative error code, otherwise

### `enum libbpf_strict_mode`

```c
/* As of libbpf 1.0 libbpf_set_strict_mode() and enum libbpf_struct_mode have
 * no effect. But they are left in libbpf_legacy.h so that applications that
 * prepared for libbpf 1.0 before final release by using
 * libbpf_set_strict_mode() still work with libbpf 1.0+ without any changes.
 */
enum libbpf_strict_mode {
	/* Turn on all supported strict features of libbpf to simulate libbpf
	 * v1.0 behavior.
	 * This will be the default behavior in libbpf v1.0.
	 */
	LIBBPF_STRICT_ALL = 0xffffffff,

	/*
	 * Disable any libbpf 1.0 behaviors. This is the default before libbpf
	 * v1.0. It won't be supported anymore in v1.0, please update your
	 * code so that it handles LIBBPF_STRICT_ALL mode before libbpf v1.0.
	 */
	LIBBPF_STRICT_NONE = 0x00,
	/*
	 * Return NULL pointers on error, not ERR_PTR(err).
	 * Additionally, libbpf also always sets errno to corresponding Exx
	 * (positive) error code.
	 */
	LIBBPF_STRICT_CLEAN_PTRS = 0x01,
	/*
	 * Return actual error codes from low-level APIs directly, not just -1.
	 * Additionally, libbpf also always sets errno to corresponding Exx
	 * (positive) error code.
	 */
	LIBBPF_STRICT_DIRECT_ERRS = 0x02,
	/*
	 * Enforce strict BPF program section (SEC()) names.
	 * E.g., while prefiously SEC("xdp_whatever") or SEC("perf_event_blah") were
	 * allowed, with LIBBPF_STRICT_SEC_PREFIX this will become
	 * unrecognized by libbpf and would have to be just SEC("xdp") and
	 * SEC("xdp") and SEC("perf_event").
	 *
	 * Note, in this mode the program pin path will be based on the
	 * function name instead of section name.
	 *
	 * Additionally, routines in the .text section are always considered
	 * sub-programs. Legacy behavior allows for a single routine in .text
	 * to be a program.
	 */
	LIBBPF_STRICT_SEC_NAME = 0x04,
	/*
	 * Disable the global 'bpf_objects_list'. Maintaining this list adds
	 * a race condition to bpf_object__open() and bpf_object__close().
	 * Clients can maintain it on their own if it is valuable for them.
	 */
	LIBBPF_STRICT_NO_OBJECT_LIST = 0x08,
	/*
	 * Automatically bump RLIMIT_MEMLOCK using setrlimit() before the
	 * first BPF program or map creation operation. This is done only if
	 * kernel is too old to support memcg-based memory accounting for BPF
	 * subsystem. By default, RLIMIT_MEMLOCK limit is set to RLIM_INFINITY,
	 * but it can be overridden with libbpf_set_memlock_rlim() API.
	 * Note that libbpf_set_memlock_rlim() needs to be called before
	 * the very first bpf_prog_load(), bpf_map_create() or bpf_object__load()
	 * operation.
	 */
	LIBBPF_STRICT_AUTO_RLIMIT_MEMLOCK = 0x10,
	/*
	 * Error out on any SEC("maps") map definition, which are deprecated
	 * in favor of BTF-defined map definitions in SEC(".maps").
	 */
	LIBBPF_STRICT_MAP_DEFINITIONS = 0x20,

	__LIBBPF_STRICT_LAST,
};
```

## Usage

!!! warning
    as of v1.0 `libbpf_set_strict_mode` is a no-op

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
