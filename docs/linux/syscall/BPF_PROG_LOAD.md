---
title: "Syscall command 'BPF_PROG_LOAD'"
description: "This page documents the 'BPF_PROG_LOAD' eBPF syscall command, including its definition, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_PROG_LOAD` command

<!-- [FEATURE_TAG](BPF_PROG_LOAD) -->
[:octicons-tag-24: v3.18](https://github.com/torvalds/linux/commit/09756af46893c18839062976c3252e93a1beeba7)
<!-- [/FEATURE_TAG] -->

The `BPF_PROG_LOAD` command loads a program into the kernel.

## Return type

This command will return the file descriptor of the program (positive integer) or an error number (negative integer) if the program wasn't loaded for whatever reason.

## Attributes

```c
union bpf_attr {
    struct {
		__u32           [prog_type](#prog_type);
		__u32           [insn_cnt](#insn_cnt);
		__aligned_u64   [insns](#insns);
		__aligned_u64   [license](#license);
		__u32           [log_level](#log_level);
		__u32           [log_size](#log_size);
		__aligned_u64   [log_buf](#log_buf);
		__u32           [kern_version](#kern_version);
		__u32           [prog_flags](#prog_flags);
		char            [prog_name](#prog_name)[BPF_OBJ_NAME_LEN];
		__u32           [prog_ifindex](#prog_ifindex);
		__u32           [expected_attach_type](#expected_attach_type);
		__u32           [prog_btf_fd](#prog_btf_fd);
		__u32           [func_info_rec_size](#func_info_rec_size);
		__aligned_u64   [func_info](#func_info);
		__u32           [func_info_cnt](#func_info_cnt);
		__u32           [line_info_rec_size](#line_info_rec_size);
		__aligned_u64   [line_info](#line_info);
		__u32           [line_info_cnt](#line_info_cnt);
		__u32           [attach_btf_id](#attach_btf_id);
		union {
			__u32       [attach_prog_fd](#attach_prog_fd);
			__u32       [attach_btf_obj_fd](#attach_btf_obj_fd);
		};
		__u32           [core_relo_cnt](#core_relo_cnt);
		__aligned_u64   [fd_array](#fd_array);
		__aligned_u64   [core_relos](#core_relos);
		__u32           [core_relo_rec_size](#core_relo_rec_size);
		__u32           [log_true_size](#log_true_size);
		__s32           [prog_token_fd](#prog_token_fd);
		__u32           [fd_array_cnt](#fd_array_cnt);
	};
}
```

### `prog_type`

[:octicons-tag-24: v3.18](https://github.com/torvalds/linux/commit/09756af46893c18839062976c3252e93a1beeba7)

This attribute specifies the type of the program to be loaded and must be one of the types defined in [program types](../program-type/index.md).

### `insn_cnt`

[:octicons-tag-24: v3.18](https://github.com/torvalds/linux/commit/09756af46893c18839062976c3252e93a1beeba7)

This attribute specifies the number of eBPF instructions which are passed to `insns`. This is used to know how much memory to read so it must be correctly sized. If only the amount of bytes is known, one can simply divided by `8` since every eBPF instruction is 8 bytes wide.

### `insns`

[:octicons-tag-24: v3.18](https://github.com/torvalds/linux/commit/09756af46893c18839062976c3252e93a1beeba7)

This attributes specifies the actual eBPF instructions of the program to be loaded. It should be a pointer to memory containing the instructions. The size of this blob is indicated by `insn_cnt`.

### `license`

[:octicons-tag-24: v3.18](https://github.com/torvalds/linux/commit/09756af46893c18839062976c3252e93a1beeba7)

This attribute specifies the software license of the eBPF program to be loaded. It should be a pointer to a null-terminated string containing the human readable license. For example `GPL`, `MIT` or `Proprietary`.

A number of helper functions in the kernel are GPL-licensed and may only be called from "GPL compatible" programs. The following license strings are recognized as "GPL compatible":

* GPL
* GPL v2
* GPL and additional rights
* Dual BSD/GPL
* Dual MIT/GPL
* Dual MPL/GPL

### `log_level`

[:octicons-tag-24: v3.18](https://github.com/torvalds/linux/commit/cbd357008604925355ae7b54a09137dabb81b580)

The lower 2 bits of this value are the log level:

* `0` = no log
* `1` = basic logging   (`BPF_LOG_LEVEL1`)
* `2` = verbose logging (`BPF_LOG_LEVEL2`)

The remaining bits are flags:

* `1 << 3` (`BPF_LOG_STATS`) If set the kernel will output statistics to the log. Flags can be used since [:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/06ee7115b0d1742de745ad143fb5e06d77d27fba)

* `1 << 4` (`BPF_LOG_FIXED`) since [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/1216640938035e63bdbd32438e91c9bcc1fd8ee1), the verifier log rotates instead of truncating. When `log_size` is exceeded. Setting this flag preserves the old behavior of truncating the log to `log_size` bytes.

### `log_size`

[:octicons-tag-24: v3.18](https://github.com/torvalds/linux/commit/cbd357008604925355ae7b54a09137dabb81b580)

This attributes indicates the size of the memory region in bytes indicated by `log_buf` which can safely be written to by the kernel.

### `log_buf`

[:octicons-tag-24: v3.18](https://github.com/torvalds/linux/commit/cbd357008604925355ae7b54a09137dabb81b580)

This attributes can be set to a pointer to a memory region allocated/reserved by the loader process where the verifier log will be written to. The detail of the log is set by `log_level`. The verifier log is often the only indication in addition to the error code of why the syscall command failed to load the program.

The log is also written to on success. If the kernel runs out of space in the buffer while loading, on older kernels the loading process will fail and the command will return with an error code of `-ENOSPC`. Callers had to take care to allocate enough buffer room, or load the program repeatedly and increasing the buffer when `-ENOSPC` was returned.

Since [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/1216640938035e63bdbd32438e91c9bcc1fd8ee1) the kernel rotates the verifier log by default. The bottom of the log is always included, which is most of the time the most interesting part. Old behavior can be requested by setting the `BPF_LOG_FIXED` flag in `log_level`.

### `kern_version`

[:octicons-tag-24: v4.1](https://github.com/torvalds/linux/commit/2541517c32be2531e0da59dfd7efc1ce844644f5)

!!! warning
    This attribute is no longer used as of [:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/6c4fc209fcf9d27efbaa48368773e4d2bfbd59aa). The field is still present for compatibly reasons but does not do anything.

Before its deprecation, this field was supposed to point to a string containing the current kernel version. This string was checked against the actual kernel version for programs of type [`BPF_PROG_TYPE_KPROBE`](../program-type/BPF_PROG_TYPE_KPROBE.md).The rational behind the field was that kprobe are fundamentally unstable and thus had to be recompiled for every kernel version (this was before [CO-RE](../../concepts/core.md) was introduced), having to set this field would make this apparent to users.

The field was retired due to the invention of [CO-RE](../../concepts/core.md) and the tendency of users/libraries to automate setting this field anyway based on `uname` without actually re-compiling.

### `prog_flags`

[:octicons-tag-24: v4.12](https://github.com/torvalds/linux/commit/e07b98d9bffe410019dfcf62c3428d4a96c56a2c)

This attribute specifies flags for all sorts of purposes, please see the [`Flags`](#flags) section for details.

### `prog_name`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/067cae47771c864604969fd902efe10916e0d79c)

This attribute specifies the name of the program. It is a 16 byte array which should be filled with a null-terminated
string thus leaving 15 characters for the name which must be one of (A-Z, a-z, 0-9, `-`, `_`).

This name is reported back to the user in the output of [`BPF_OBJ_GET`](./BPF_OBJ_GET.md). No uniqueness guarantees are made by the kernel, so it is only useful for human feedback.

### `prog_ifindex`

[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/1f6f4cb7ba219b00a3fa9afe8049fa16444d8b52)

This attribute specifies the network interface index the user intends to attach this program to after loading. If the user intends to offload a given program to a network device, they must set this field so the drivers of that network device can validate the program in addition to the kernel verifier to gauge if the selected network device can offload the given program.

### `expected_attach_type`

[:octicons-tag-24: v4.17](https://github.com/torvalds/linux/commit/5e43f899b03a3492ce5fc44e8900becb04dae9c0)

This attribute specifies the attach type the user expects to use when attaching the program. For certain program types, the attach type may changes aspects like the context type that will be given, the meaning of return values, and which helper function are or are not available. Therefor the verifier must know the attach type during loading time to enforce correct behavior of the program to be loaded.

The expected attach type is known to be important in the following cases:

* For `BPF_PROG_TYPE_LSM` programs only programs attached with type `BPF_LSM_CGROUP` are allowed to use certain helper functions.
* For `BPF_PROG_TYPE_TRACING` programs the attach type determine access to helper calls
* For `BPF_PROG_TYPE_CGROUP_SOCK_ADDR` programs the verifier restricts valid return values depending on attach type
* For `BPF_PROG_TYPE_CGROUP_SKB` programs the verifier restricts valid return values depending on attach type
* For `BPF_PROG_TYPE_CGROUP_SOCKOPT` programs the attach type determines accessability for certain context fields and helper functions.
* Only `BPF_PROG_TYPE_XDP` programs with `BPF_XDP_CPUMAP` attach type can be added to the values of `BPF_MAP_TYPE_CPUMAP` maps
* Only `BPF_PROG_TYPE_XDP` programs with `BPF_XDP_DEVMAP` attach type can be added to the values of `BPF_MAP_TYPE_DEVMAP` maps

!!! note
    For `BPF_PROG_TYPE_STRUCT_OPS` program types the `expected_attach_type` doesn't contain a constant or enum value but rater the member index of the [BTF](../../concepts/btf.md) struct specified by `attach_btf_id` which is to be replaced by this eBPF program

### `prog_btf_fd`

[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/838e96904ff3fc6c30e5ebbc611474669856e3c0)

This attribute specifies the file descriptor of the [BTF](../../concepts/btf.md) object which contains type information associated with the program we are loaded.

Loading BTF for your program is optional, but highly recommended since a ever growing number of features require BTF to properly function.

### `func_info_rec_size`

[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/838e96904ff3fc6c30e5ebbc611474669856e3c0)

This attribute specifies the size of the records in `func_info`, this allows for compatibility between newer and older loaders and kernel versions if the size of the function info records ever changes.

### `func_info`

[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/838e96904ff3fc6c30e5ebbc611474669856e3c0)

This attribute specifies the memory region where extended BTF function info is located. This should be a pointer to an array of function info records with the size of `func_info_rec_size`. The array should contain `func_info_cnt` of these records.

<!-- TODO link to the structure of these records in the BTF page -->

This function info contains the signatures of functions within the program and is used to validate these signatures match expected signatures when used as callbacks for certain helper functions like [`bpf_loop`](../helper-function/bpf_loop.md) and[`bpf_timer_set_callback`](../helper-function/bpf_timer_set_callback.md).

### `func_info_cnt`

[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/838e96904ff3fc6c30e5ebbc611474669856e3c0)

This attribute specifies the amount of function records that are present in `func_info`.

### `line_info_rec_size`

[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/c454a46b5efd8eff8880e88ece2976e60a26bf35)

This attribute specifies the size of the records in `line_info`, this allows for compatibility between newer and older loaders and kernel versions if the size of the line info records ever changes.

### `line_info`

[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/c454a46b5efd8eff8880e88ece2976e60a26bf35)

This attribute specifies the memory region where extended BTF line info is located. This should be a pointer to an array of line info records with the size of `line_info_rec_size`. The array should contain `line_info_cnt` of these records.

<!-- TODO link to the structure of these records in the BTF page -->

This line information associates information like the filename+path, line number, column number and an snippet of source code which produced a given piece of the eBPF code. This information is available in the verifier log to make understanding the output easier as well as in output of [`BPF_OBJ_GET`](./BPF_OBJ_GET.md).

!!! note
    The verifier also enforces that for every function info record, there also exists a line info record on the same instruction.

### `line_info_cnt`

[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/c454a46b5efd8eff8880e88ece2976e60a26bf35)

This attribute specifies the amount of function records that are present in `line_info`.

### `attach_btf_id`

[:octicons-tag-24: v5.5](https://github.com/torvalds/linux/commit/ccfe29eb29c2edcea6552072ef00ff4117f53e83)

This attribute specifies the [BTF](../../concepts/btf.md) type ID of kernel types the current program wishes to attach to. This ID refers the ID within the `vmlinux` object, not the BTF object specified by `prog_btf_fd`. This attribute can have different meaning depending on the program type.

* For `BPF_PROG_TYPE_STRUCT_OPS` this attribute is the ID of the ops struct of which the user wants to replace a function pointer with an eBPF program.
* For `BPF_PROG_TYPE_LSM` this attribute specifies the LSM hook point where we intend to attach it to.
<!-- TODO It is used in a number of other locations, need more research -->


### `attach_prog_fd`

[:octicons-tag-24: v5.15](https://github.com/torvalds/linux/commit/5b92a28aae4dd0f88778d540ecfdcdaec5a41723)

This attribute specifies the file descriptor of an already loaded eBPF program. It is used in [`BPF_PROG_TYPE_EXT`](../program-type/BPF_PROG_TYPE_EXT.md) program types to select which existing program should be extended.

### `attach_btf_obj_fd`

[:octicons-tag-24: v5.11](https://github.com/torvalds/linux/commit/290248a5b7d829871b3ea3c62578613a580a1744)

This attribute specifies the file descriptor of a BTF object which the kernel should use instead of its internal vmlinux object. This is mainly used to hook BTF-dependant program types such as raw tracepoints, fentry/fexit, and LSM into kernel modules.

### `core_relo_cnt`

[:octicons-tag-24: v5.17](https://github.com/torvalds/linux/commit/fbd94c7afcf99c9f3b1ba1168657ecc428eb2c8d)

This attribute specifies the size of the records in `core_relos`, this allows for compatibility between newer and older loaders and kernel versions if the size of the CO-RE relocation records ever changes.

### `fd_array`

[:octicons-tag-24: v5.14](https://github.com/torvalds/linux/commit/387544bfa291a22383d60b40f887360e2b931ec6)

This attribute specifies an array of file descriptors. This value should be a pointer to an array of 32 bit values containing file descriptors. When using this feature, loaders don't have to rewrite the eBPF program so the blob in the ELF can be signed. The instructions will instead contain index into this array and the actual file descriptors which may be different between program runs are thus not included in any signable blob.

If [`fd_array_cnt`](#fd_array_cnt) is `0`, or in kernel versions without the attribute. The array can be sparse (no every element had to contain a valid file descriptor, just the indices referenced by instructions) and only map file descriptors are allowed.

If [`fd_array_cnt`](#fd_array_cnt) is set, the array must be contiguous of [`fd_array_cnt`](#fd_array_cnt) elements and all file descriptors must be valid. The array may contain file descriptors of maps and BTF objects. The program will take a refcount on all file descriptors in the array, not just the ones referenced by instructions.

### `core_relos`

[:octicons-tag-24: v5.17](https://github.com/torvalds/linux/commit/fbd94c7afcf99c9f3b1ba1168657ecc428eb2c8d)

This attribute specifies the memory region where CO-RE relocations is located. This should be a pointer to an array of CO-Re relocation records with the size of `core_relo_rec_size`. The array should contain `core_relo_cnt` of these records.

Before the addition of this field, CO-RE relocations had to be performed by the loader in userspace. This process modifies the eBPF program passed in via `insns` which makes signing of the program difficult. This field passes the CO-RE relocations to the kernel so that these can happen after verifying a potential program signature.

### `core_relo_rec_size`

[:octicons-tag-24: v5.17](https://github.com/torvalds/linux/commit/fbd94c7afcf99c9f3b1ba1168657ecc428eb2c8d)

This attribute specifies the amount of function records that are present in `core_relos`.

### `log_true_size`

[:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/47a71c1f9af0a334c9dfa97633c41de4feda4287)

Since :octicons-tag-24: v6.4, the verifier log is not truncated to `log_size` anymore, instead it is rotated such that the bottom of the log is always included, which is most of the time the most interesting part. 

This field is an output, it will be set to the actual size of the log. If a user wants to get the full verifier log, this allows them to allocate a buffer of the correct size before calling `BPF_PROG_LOAD` again with the same inputs (other than the resized buffer).

### `prog_token_fd`

[:octicons-tag-24: v6.9](https://github.com/torvalds/linux/commit/caf8f28e036c4ba1e823355da6c0c01c39e70ab9)

The file descriptor of a [BPF token](../../linux/concepts/token.md) can be passed to this attribute. If the BPF token grants permission to create a program of the type specified in [`prog_type`](#prog_type), the kernel will allow the program to be loaded for a user without `CAP_BPF`.

### `fd_array_cnt`

[:octicons-tag-24: v6.14](https://github.com/torvalds/linux/commit/4d3ae294f900fb7232fb6c890dbd3176b8a5f121)

If set, the [`fd_array`](#fd_array) must be contiguous of `fd_array_cnt` elements and all file descriptors must be valid. The array may contain file descriptors of maps and BTF objects. The program will take a refcount on all file descriptors in the array, not just the ones referenced by instructions.

## Flags

### `BPF_F_STRICT_ALIGNMENT`

<!-- [FEATURE_TAG](BPF_F_STRICT_ALIGNMENT) -->
[:octicons-tag-24: v4.12](https://github.com/torvalds/linux/commit/e07b98d9bffe410019dfcf62c3428d4a96c56a2c)
<!-- [/FEATURE_TAG] -->

If `BPF_F_STRICT_ALIGNMENT` is used in `BPF_PROG_LOAD` command, the verifier will perform strict alignment checking as if the kernel has been built with `CONFIG_EFFICIENT_UNALIGNED_ACCESS` not set, and `NET_IP_ALIGN` defined to 2.

### `BPF_F_ANY_ALIGNMENT`

<!-- [FEATURE_TAG](BPF_F_ANY_ALIGNMENT) -->
[:octicons-tag-24: v5.0](https://github.com/torvalds/linux/commit/e9ee9efc0d176512cdce9d27ff8549d7ffa2bfcd)
<!-- [/FEATURE_TAG] -->

If `BPF_F_ANY_ALIGNMENT` is used in `BPF_PROF_LOAD` command, the verifier will allow any alignment whatsoever.  On platforms with strict alignment requirements for loads ands stores (such as sparc and mips) the verifier validates that all loads and stores provably follow this requirement.  This flag turns that checking and enforcement off.

It is mostly used for testing when we want to validate the context and memory access aspects of the verifier, but because of an unaligned access the alignment check would trigger before the one we are interested in.

### `BPF_F_TEST_RND_HI32`

<!-- [FEATURE_TAG](BPF_F_TEST_RND_HI32) -->
[:octicons-tag-24: v5.3](https://github.com/torvalds/linux/commit/c240eff63a1cf1c4edc768e0cfc374811c02f069)
<!-- [/FEATURE_TAG] -->

!!! warning
    `BPF_F_TEST_RND_HI32` is used for testing purpose, not meant for production usage.

Verifier does sub-register def/use analysis and identifies instructions whose def only matters for low 32-bit, high 32-bit is never referenced later through implicit zero extension. Therefore verifier notifies JIT back-ends that it is safe to ignore clearing high 32-bit for these instructions. This saves some back-ends a lot of code-gen. However such optimization is not necessary on some arches, for example x86_64, arm64 etc, whose JIT back-ends hence hasn't used verifier's analysis result. But, we really want to have a way to be able to verify the correctness of the described optimization on x86_64 on which test suites are frequently exercised.

So, this flag is introduced. Once it is set, verifier will randomize high 32-bit for those instructions who has been identified as safe to ignore them. Then, if verifier is not doing correct analysis, such randomization will regress tests to expose bugs.

### `BPF_F_TEST_STATE_FREQ`

<!-- [FEATURE_TAG](BPF_F_TEST_STATE_FREQ) -->
[:octicons-tag-24: v5.4](https://github.com/torvalds/linux/commit/10d274e880eb208ec6a76261a9f8f8155020f771)
<!-- [/FEATURE_TAG] -->

The verifier internal test flag used for stress testing state pruning. 

!!! warning
    Behavior is undefined 

### `BPF_F_SLEEPABLE`

<!-- [FEATURE_TAG](BPF_F_SLEEPABLE) -->
[:octicons-tag-24: v5.10](https://github.com/torvalds/linux/commit/1e6c62a8821557720a9b2ea9617359b264f2f67c)
<!-- [/FEATURE_TAG] -->

If `BPF_F_SLEEPABLE` is used in `BPF_PROG_LOAD` command, the verifier will restrict map and helper usage for such programs. Sleepable BPF programs can only be attached to hooks where kernel execution context allows sleeping. Such programs are allowed to use helpers that may sleep like [`bpf_copy_from_user`](../helper-function/bpf_copy_from_user.md).

### `BPF_F_XDP_HAS_FRAGS`

<!-- [FEATURE_TAG](BPF_F_XDP_HAS_FRAGS) -->
[:octicons-tag-24: v5.18](https://github.com/torvalds/linux/commit/c2f2cdbeffda7b153c19e0f3d73149c41026c0db)
<!-- [/FEATURE_TAG] -->

This flag notifies the kernel that the XDP program supports XDP fragments. If set, the XDP program may be called with a context that doesn't include the full packet in a single linear piece of memory, which breaks assumptions most XDP programs have, hence the flag.

For more details, check out the [XDP program type page](../program-type/BPF_PROG_TYPE_XDP.md#xdp-fragments)

### `BPF_F_XDP_DEV_BOUND_ONLY`

[:octicons-tag-24: v6.3](https://github.com/torvalds/linux/commit/2b3486bc2d237ec345b3942b7be5deabf8c8fed1)

If `BPF_F_XDP_DEV_BOUND_ONLY` is used in `BPF_PROG_LOAD` command, the loaded program becomes device-bound but can access XDP metadata.

### `BPF_F_TEST_REG_INVARIANTS`

[:octicons-tag-24: v6.8](https://github.com/torvalds/linux/commit/ff8867af01daa7ea770bebf5f91199b7434b74e5)

The verifier internal test flag. Behavior is undefined 

### `BPF_F_TOKEN_FD`

<!-- [FEATURE_TAG](BPF_F_TOKEN_FD) -->
[:octicons-tag-24: v6.9](https://github.com/torvalds/linux/commit/a177fc2bf6fd83704854feaf7aae926b1df4f0b9)
<!-- [/FEATURE_TAG] -->

When set, the kernel will use the BPF token in [`prog_token_fd`](#prog_token_fd) to authorize the creation of the map instead of checking the capabilities of the current user.
