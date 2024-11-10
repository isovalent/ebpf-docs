# Libbpf eBPF side

Libbpf contains a number of C header files containing mostly pre-processor macros, forward declarations and type definitions that make it easier to write eBPF programs. This is an index into these useful definitions.

## [`bpf_helper_defs.h`](https://github.com/libbpf/libbpf/blob/master/src/bpf_helper_defs.h)

The `bpf_helper_defs.h` file is automatically generated from the kernel sources. It contains forward declarations for every type that is used by [eBPF helper functions](../../../linux/helper-function/index.md) and somewhat special forward declarations for the helper functions themselves.

For example, the `bpf_map_lookup_elem` function is declared as: 

`#!c static void *(* const bpf_map_lookup_elem)(void *map, const void *key) = (void *) 1;`

The normal forward declaration of this function would be 

`#!c void *bpf_map_lookup_elem(void *map, const void *key);`.

But what the special declaration does is it casts a pointer of value `1` to a const static function pointer. This causes the compiler to emit a `call 1` instruction which the kernel recognizes as a call to the `bpf_map_lookup_elem` function.

It is entirely possible to copy parts of this file if you are only interested in specific helper functions and their types and even modify their definitions to suit your needs. Though for most people it will be best to include the whole file.

## [`bpf_helpers.h`](https://github.com/libbpf/libbpf/blob/master/src/bpf_helpers.h)

The `bpf_helpers.h` file is the single most useful file in the eBPF side of the libbpf library. It contains a lot of
generic and basic definitions you will use in almost any eBPF program. It also includes the `bpf_helper_defs.h` file,
so you don't need to include it separately.

The file contains definitions for the following:

* BTF map macros / types
    * [`__uint`](__uint.md)
    * [`__uint`](__uint.md)
    * [`__array`](__array.md)
    * [`__ulong`](__ulong.md)
    * [`enum libbpf_pin_type`](enum-libbpf_pin_type.md)
* `SEC`
* Attributes
    * [`__always_inline`](__always_inline.md)
    * [`__noinline`](__noinline.md)
    * [`__weak`](__weak.md)
    * [`__hidden`](__hidden.md)
    * [`__kconfig`](__kconfig.md)
    * [`__ksym`](__ksym.md)
    * [`__kptr_untrusted`](__kptr_untrusted.md)
    * [`__kptr`](__kptr.md)
    * [`__percpu_kptr`](__percpu_kptr.md)
* `NULL`
* `KERNEL_VERSION`
* `offsetof`
* `container_of`
* `barrier`
* `barrier_var`
* `__bpf_unreachable`
* `bpf_tail_call_static`
* `enum libbpf_tristate`
* `bpf_ksym_exists`
* Global function attributes
    * `__arg_ctx`
    * `__arg_nonnull`
    * `__arg_nullable`
    * `__arg_trusted`
    * `__arg_arena`
* <nospell>Printf macros</nospell>
    * `BPF_SEQ_PRINTF`
    * `BPF_SNPRINTF`
    * `bpf_printk`
* Open coded iterator loop macros
    * `bpf_for_each`
    * `bpf_for`
    * `bpf_repeat`

## [`bpf_endian.h`](https://github.com/libbpf/libbpf/blob/master/src/bpf_endian.h)

The `bpf_endian.h` file contains macros for endianess conversion. It is useful when you need to convert data between host and network byte order.

The file contains definitions for the following:

* `bpf_htons`
* `bpf_ntohs`
* `bpf_htonl`
* `bpf_ntohl`
* `bpf_cpu_to_be64`
* `bpf_be64_to_cpu`

## [`bpf_tracing.h`](https://github.com/libbpf/libbpf/blob/master/src/bpf_tracing.h)

The `bpf_tracing.h` file contains macros which are mostly meant for tracing-like program types such as [`BPF_PROG_TYPE_KPROBE`](../../../linux/program-type/BPF_PROG_TYPE_KPROBE.md) and [`BPF_PROG_TYPE_TRACING`](../../../linux/program-type/BPF_PROG_TYPE_TRACING.md). Most provided functionality is related to the given context to parameters and/or return value.

The file contains definitions for the following:

* `PT_REGS_PARM`
* `PT_REGS_RET`
* `PT_REGS_FP`
* `PT_REGS_RC`
* `PT_REGS_SP`
* `PT_REGS_IP`
* `PT_REGS_SYSCALL_REGS`
* `BPF_PROG`
* `BPF_PROG2`
* `BPF_KPROBE`/`BPF_UPROBE`
* `BPF_KRETPROBE`/`BPF_URETPROBE`
* `BPF_KSYSCALL`/`BPF_KPROBE_SYSCALL`

## [`bpf_core_read.h`](https://github.com/libbpf/libbpf/blob/master/src/bpf_core_read.h)

The `bpf_core_read.h` file contains macros for CO-RE(Compile Once Run Everywhere) operations.

The file contains definitions for the following:

* `BPF_CORE_READ_BITFIELD_PROBED`
* `BPF_CORE_READ_BITFIELD`
* `BPF_CORE_WRITE_BITFIELD`
* `bpf_core_field_exists`
* `bpf_core_field_size`
* `bpf_core_field_offset`
* `bpf_core_type_id_local`
* `bpf_core_type_id_kernel`
* `bpf_core_type_exists`
* `bpf_core_type_matches`
* `bpf_core_type_size`
* `bpf_core_enum_value_exists`
* `bpf_core_enum_value`
* `bpf_core_read`
* `bpf_core_read_user`
* `bpf_core_read_str`
* `bpf_core_read_user_str`
* `bpf_core_cast`
* `BPF_CORE_READ_INTO`
* `BPF_CORE_READ_USER_INTO`
* `BPF_CORE_READ_STR_INTO`
* `BPF_CORE_READ_USER_STR_INTO`
* `BPF_CORE_READ_BITFIELD_INTO`
* `BPF_CORE_READ`
* `BPF_CORE_READ_USER`
* `BPF_PROBE_READ_INTO`
* `BPF_PROBE_READ_USER_INTO`
* `BPF_PROBE_READ_STR_INTO`
* `BPF_PROBE_READ_USER_STR_INTO`
* `BPF_PROBE_READ`
* `BPF_PROBE_READ_USER`
