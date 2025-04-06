# Libbpf eBPF side

Libbpf contains a number of C header files containing mostly pre-processor macros, forward declarations and type definitions that make it easier to write eBPF programs. This is an index into these useful definitions.

## [`bpf_helper_defs.h`](https://github.com/libbpf/libbpf/blob/master/src/bpf_helper_defs.h)

The `bpf_helper_defs.h` file is automatically generated from the kernel sources. It contains forward declarations for every type that is used by [eBPF helper functions](../../../linux/helper-function/index.md) and somewhat special forward declarations for the helper functions themselves.

For example, the [`bpf_map_lookup_elem`](../../../linux/helper-function/bpf_map_lookup_elem.md) function is declared as: 

`#!c static void *(* const bpf_map_lookup_elem)(void *map, const void *key) = (void *) 1;`

The normal forward declaration of this function would be 

`#!c void *bpf_map_lookup_elem(void *map, const void *key);`.

But what the special declaration does is it casts a pointer of value `1` to a const static function pointer. This causes the compiler to emit a `call 1` instruction which the kernel recognizes as a call to the [`bpf_map_lookup_elem`](../../../linux/helper-function/bpf_map_lookup_elem.md) function.

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
* Global function attributes
    * [`__arg_ctx`](__arg_ctx.md)
    * [`__arg_nonnull`](__arg_nonnull.md)
    * [`__arg_nullable`](__arg_nullable.md)
    * [`__arg_trusted`](__arg_trusted.md)
    * [`__arg_arena`](__arg_arena.md)
* [`SEC`](SEC.md)
* [`KERNEL_VERSION`](KERNEL_VERSION.md)
* [`offsetof`](offsetof.md)
* [`container_of`](container_of.md)
* [`barrier`](barrier.md)
* [`barrier_var`](barrier_var.md)
* [`__bpf_unreachable`](__bpf_unreachable.md)
* [`bpf_tail_call_static`](bpf_tail_call_static.md)
* [`bpf_ksym_exists`](bpf_ksym_exists.md)
* <nospell>Printf macros</nospell>
    * [`BPF_SEQ_PRINTF`](bpf_seq_printf.md)
    * [`BPF_SNPRINTF`](bpf_snprintf.md)
    * [`bpf_printk`](bpf_printk.md)
* Open coded iterator loop macros
    * [`bpf_for_each`](bpf_for_each.md)
    * [`bpf_for`](bpf_for.md)
    * [`bpf_repeat`](bpf_repeat.md)
* Utility macros
    * [`___bpf_fill`](___bpf_fill.md)

## [`bpf_endian.h`](https://github.com/libbpf/libbpf/blob/master/src/bpf_endian.h)

The `bpf_endian.h` file contains macros for endianess conversion. It is useful when you need to convert data between host and network byte order.

The file contains definitions for the following:

* [`bpf_htons`](bpf_htons.md)
* [`bpf_ntohs`](bpf_ntohs.md)
* [`bpf_htonl`](bpf_htonl.md)
* [`bpf_ntohl`](bpf_ntohl.md)
* [`bpf_cpu_to_be64`](bpf_cpu_to_be64.md)
* [`bpf_be64_to_cpu`](bpf_be64_to_cpu.md)

## [`bpf_tracing.h`](https://github.com/libbpf/libbpf/blob/master/src/bpf_tracing.h)

The `bpf_tracing.h` file contains macros which are mostly meant for tracing-like program types such as [`BPF_PROG_TYPE_KPROBE`](../../../linux/program-type/BPF_PROG_TYPE_KPROBE.md) and [`BPF_PROG_TYPE_TRACING`](../../../linux/program-type/BPF_PROG_TYPE_TRACING.md). Most provided functionality is related to the given context to parameters and/or return value.

The file contains definitions for the following:

* [`PT_REGS_PARM`](PT_REGS_PARM.md)
* [`PT_REGS_RET`](PT_REGS_RET.md)
* [`PT_REGS_FP`](PT_REGS_FP.md)
* [`PT_REGS_RC`](PT_REGS_RC.md)
* [`PT_REGS_SP`](PT_REGS_SP.md) 
* [`PT_REGS_IP`](PT_REGS_IP.md)
* [`PT_REGS_SYSCALL_REGS`](PT_REGS_SYSCALL_REGS.md)
* [`BPF_PROG`](BPF_PROG.md)
* [`BPF_PROG2`](BPF_PROG2.md)
* [`BPF_KPROBE`](BPF_KPROBE.md)/[`BPF_UPROBE`](BPF_UPROBE.md)
* [`BPF_KRETPROBE`](BPF_KRETPROBE.md)/[`BPF_URETPROBE`](BPF_URETPROBE.md)
* [`BPF_KSYSCALL`](BPF_KSYSCALL.md)/[`BPF_KPROBE_SYSCALL`](BPF_KPROBE_SYSCALL.md)

## [`bpf_core_read.h`](https://github.com/libbpf/libbpf/blob/master/src/bpf_core_read.h)

The `bpf_core_read.h` file contains macros for CO-RE(Compile Once Run Everywhere) operations.

The file contains definitions for the following:

* CO-RE memory access
    * [`BPF_CORE_READ`](BPF_CORE_READ.md)
    * [`BPF_CORE_READ_INTO`](BPF_CORE_READ_INTO.md)
    * [`bpf_core_read`](bpf_core_read.md)
    * [`BPF_CORE_READ_STR_INTO`](BPF_CORE_READ_STR_INTO.md)
    * [`bpf_core_read_str`](bpf_core_read_str.md)
    * [`BPF_CORE_READ_USER`](BPF_CORE_READ_USER.md)
    * [`BPF_CORE_READ_USER_INTO`](BPF_CORE_READ_USER_INTO.md)
    * [`bpf_core_read_user`](bpf_core_read_user.md)
    * [`BPF_CORE_READ_USER_STR_INTO`](BPF_CORE_READ_USER_STR_INTO.md)
    * [`bpf_core_read_user_str`](bpf_core_read_user_str.md)
    * [`BPF_CORE_READ_BITFIELD`](BPF_CORE_READ_BITFIELD.md)
    * [`BPF_CORE_READ_BITFIELD_PROBED`](BPF_CORE_READ_BITFIELD_PROBED.md)
    * [`BPF_CORE_WRITE_BITFIELD`](BPF_CORE_WRITE_BITFIELD.md)
* CO-RE queries
    * [`bpf_core_field_exists`](bpf_core_field_exists.md)
    * [`bpf_core_field_size`](bpf_core_field_size.md)
    * [`bpf_core_field_offset`](bpf_core_field_offset.md)
    * [`bpf_core_type_id_local`](bpf_core_type_id_local.md)
    * [`bpf_core_type_id_kernel`](bpf_core_type_id_kernel.md)
    * [`bpf_core_type_exists`](bpf_core_type_exists.md)
    * [`bpf_core_type_matches`](bpf_core_type_matches.md)
    * [`bpf_core_type_size`](bpf_core_type_size.md)
    * [`bpf_core_enum_value_exists`](bpf_core_enum_value_exists.md)
    * [`bpf_core_enum_value`](bpf_core_enum_value.md)
* [`bpf_core_cast`](bpf_core_cast.md)
* Non CO-RE macros
    * [`BPF_PROBE_READ`](BPF_PROBE_READ.md)
    * [`BPF_PROBE_READ_INTO`](BPF_PROBE_READ_INTO.md)
    * [`BPF_PROBE_READ_USER_INTO`](BPF_PROBE_READ_USER_INTO.md)
    * [`BPF_PROBE_READ_STR_INTO`](BPF_PROBE_READ_STR_INTO.md)
    * [`BPF_PROBE_READ_USER_STR_INTO`](BPF_PROBE_READ_USER_STR_INTO.md)
    * [`BPF_PROBE_READ_USER`](BPF_PROBE_READ_USER.md)

## [`usdt.bpf.h`](https://github.com/libbpf/libbpf/blob/master/src/usdt.bpf.h)

* [`BPF_USDT`](BPF_USDT.md)
* [`bpf_usdt_arg_cnt`](bpf_usdt_arg_cnt.md)
* [`bpf_usdt_arg_size`](bpf_usdt_arg_size.md)
* [`bpf_usdt_arg`](bpf_usdt_arg.md)
* [`bpf_usdt_cookie`](bpf_usdt_cookie.md)
