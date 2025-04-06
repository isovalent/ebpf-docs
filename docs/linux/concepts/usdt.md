# USDT (User Statically-Defined Tracing)

USDT is a technique for defining tracepoints in userspace programs which system level tooling can attach to trace the execution of these programs.

The principles of operation are fairly straightforward. 

1. Tracepoint definitions are placed by a userspace program developer at locations of interest. 
2. During compilation, these tracepoints are turned into a series of CPU instructions that prepare any arguments passed, and a <nospell>NOP</nospell> (No operation) instruction (more on this later). The resulting ELF file will contain "notes" that contain information about where in the process memory the tracepoint will be located once the program is executing, its name, and some other bits of information.
3. A tracing tool (traditionally <nospell>[SystemTap](https://sourceware.org/systemtap/)</nospell> or <nospell>[DTrace](https://dtrace.org/)</nospell>) will read the "notes" from a program executable and attach to the tracepoints to collect information.

!!! note
    There are also mechanisms for defining these tracepoints dynamically (at runtime). This is necessary for programs written in languages that are not pre-compiled. This is a bit more involved and discussed in its [own section](#dynamic-tracepoints).

## Defining tracepoints

Defining a tracepoints requires emitting a <nospell>NOP</nospell> instruction in the produced machine code and adding a "note" in the ELF section. Users will likely want to use a library for this purpose. For C and C++ the de facto library `<sys/sdt.h>` from [SystemTap](https://sourceware.org/git/?p=systemtap.git;a=blob;f=includes/sys/sdt.h;h=e743f1f29bdabb6d7f5b4bbb7ab0a448767658a0;hb=HEAD). But any library can be used as long as the expected result is emitted into the compiled executables.

An example of defining a tracepoint might look like this:

```c
#include <sys/sdt.h>

// Some function called at some point in our program
int somefunction(int8_t a, uint32_t b) {
    DTRACE_PROBE2("my_provider", "somefunction-enter", a, b);
    // [...]
}
```

The first string is the "provider", which allows a tracer to see who defined a tracepoint, since libraries may include their own tracepoints as well as the main program. The second string is the name of the tracepoint. And after that we pass two arguments. The `DTRACE_PROBE` macro passes no arguments, the `DTRACE_PROBE1` macro passes 1 arguments, and so on up to `DTRACE_PROBE12`.

The notes are added to the `.note.stapsdt` ELF section and follow the format described in [https://sourceware.org/systemtap/wiki/UserSpaceProbeImplementation](https://sourceware.org/systemtap/wiki/UserSpaceProbeImplementation)

## Attaching with eBPF

In other to attach an eBPF program to a USDT tracepoint we have to know where to attach to. So we need a loader program which can parse the USDT notes, and do some math to find out where in process memory the <nospell>NOP</nospell> instruction is located.

This location is then used to attach a [uprobe](../program-type/BPF_PROG_TYPE_KPROBE.md#usage). The <nospell>NOP</nospell> instruction is replaced with a INT3 instruction on x86 (other CPU interrupt instructions on other architectures). When the userspace program executes this instruction, the BPF program is called.

The slightly tricky bit is handling arguments passed to a tracepoint. When the USDT note is create, it records where that argument is located. But unlike with function calls, there is no ABI here, no rule for which arguments go where. The note passes the location of the argument as [GAS(GNU assembler) operand](https://en.wikibooks.org/wiki/X86_Assembly/GNU_assembly_syntax). It is up the the loader and eBPF program to figure out how to turn this into logic to actually extract these arguments from the process and use them.

Fortunately, the heavy lifting is often taken care of by libraries such as libbpf, which provides the loader logic via [`bpf_program__attach_usdt`](../../ebpf-library/libbpf/userspace/bpf_program__attach_usdt.md) (implementation in [`usdt.c`](https://github.com/libbpf/libbpf/blob/master/src/usdt.c)).

Libbpf expects that any USDT probes are written using the helpers from [`usdt.bpf.h`](../../ebpf-library/libbpf/ebpf/index.md#usdtbpfh) which includes the BPF logic to read arguments from the process according to a spec provided by the loader via maps.

Example of defining a USDT eBPF program is:

```c hl_lines="4"
// SPDX-License-Identifier: GPL-2.0
/* Copyright (c) 2022 Meta Platforms, Inc. and affiliates. */
SEC("usdt/./urandom_read:urand:read_without_sema")
int BPF_USDT(urand_read_without_sema, int iter_num, int iter_cnt, int buf_sz)
{
	if (urand_pid != (bpf_get_current_pid_tgid() >> 32))
		return 0;

	__sync_fetch_and_add(&urand_read_without_sema_call_cnt, 1);
	__sync_fetch_and_add(&urand_read_without_sema_buf_sz_sum, buf_sz);

	return 0;
}
```

The ELF section specified in [`SEC`](../../ebpf-library/libbpf/ebpf/SEC.md) can be used to tell libbpf where this program should be [auto-attached](../../ebpf-library/libbpf/userspace/bpf_program__attach.md). Starting with `usdt` for the program type, `./urandom_read` for the path to the binary, can be relative or absolute. `urand` for the provider, and `read_without_sema` for the tracepoint name.

## Semaphores

Semaphores are an optional USDT feature. A semaphore is a number which is incremented when a probe is attached and decremented when detached. This allows a program to see if it being traced. An example use could could be that you would like to expose internal state to a tracepoint, but accessing that state is costly. You can first see if any probes are attached, and only collect the arguments if at least one is.

The location of the semaphore is included in the <nospell>SDT</nospell> note.

When using the [`bpf_program__attach_usdt`](../../ebpf-library/libbpf/userspace/bpf_program__attach_usdt.md) the semaphore location is parsed from the note and set internally.

When attaching an eBPF program the kernel will increment the semaphore. Its location has to be passed to [perf_event_open](https://man7.org/linux/man-pages/man2/perf_event_open.2.html) when manually attaching a USDT via [`bpf_program__attach_uprobe_opts`](../../ebpf-library/libbpf/userspace/bpf_program__attach_uprobe_opts.md#ref_ctr_offset) in the `ref_ctr_offset` option.

When using [`bpf_program__attach_uprobe_multi`](../../ebpf-library/libbpf/userspace/bpf_program__attach_uprobe_multi.md#ref_ctr_offsets) via the `ref_ctr_offsets` option.

## Dynamic tracepoints

As the name User **Statically**-Defined Tracing implies, the tracepoints were originally intended to be statically defined at compile time. This works for programs that are written in statically compiled languages, but for programs that are more dynamic such as those using an interpreter or runtime it does not work.

For these use cases dynamic tracepoints were created. The logic is provided by libraries such as [libstapsdt](https://github.com/linux-usdt/libstapsdt). It works by creating a dynamic library (.so file) on the fly (while running), containing the tracepoints and the <nospell>SDT</nospell> notes. This dynamic library is then loaded and the exported function which simply wrap the tracepoints can be called.

Tracers are expected to scan all libraries which are dynamically linked into a process to discover these tracepoints.
