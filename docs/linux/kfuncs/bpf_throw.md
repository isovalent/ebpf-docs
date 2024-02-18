# KFunc `bpf_throw`

<!-- [FEATURE_TAG](bpf_throw) -->
[:octicons-tag-24: v6.7](https://github.com/torvalds/linux/commit/fd5d27b70188379bb441d404c29a0afb111e1753)
<!-- [/FEATURE_TAG] -->

Throw a BPF exception from the program

## Definition

Throw a BPF exception from the program, immediately terminating its execution and unwinding the stack. The supplied `cookie` parameter will be the return value of the program when an exception is thrown, and the default exception callback is used. Otherwise, if an exception callback is set using the `__exception_cb(callback)` declaration tag on the main program, the `cookie` parameter will be the callback's only input argument.

Thus, in case of default exception callback, `cookie` is subjected to constraints on the program's return value (as with R0 on exit). Otherwise, the return value of the marked exception callback will be subjected to the same checks.

!!! note
    throwing an exception with lingering resources (locks, references, etc.) will lead to a verification error.

!!! note
    callbacks **cannot** call this helper.

**Return**

Never.

**Throws**

An exception with the specified `cookie` value.

<!-- [KFUNC_DEF] -->
`#!c void bpf_throw(u64 cookie)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md)
- [BPF_PROG_TYPE_SCHED_CLS](../program-type/BPF_PROG_TYPE_SCHED_CLS.md)
- [BPF_PROG_TYPE_STRUCT_OPS](../program-type/BPF_PROG_TYPE_STRUCT_OPS.md)
- [BPF_PROG_TYPE_TRACING](../program-type/BPF_PROG_TYPE_TRACING.md)
- [BPF_PROG_TYPE_XDP](../program-type/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

