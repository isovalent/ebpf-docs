# Trampolines

Trampolines in computing has a number of meanings. In the context of Linux this refers to locations in memory containing addresses of logic to jump to. Trampolines are also referred to as "indirect jump vectors". It is a mechanism that has a number of use cases such as interrupt service routines or I/O routines. In these classic use cases the hardware hard-codes memory locations to which execution will jump when certain events such as interrupt happen. A trampoline typically jumps immediately to some other function where the actual handler lives, hence the term trampoline.

## ftrace

[:octicons-tag-24: v2.6.27](https://github.com/torvalds/linux/commit/16444a8a40d4c7b4f6de34af0cae1f76a4f6c901)

Ftrace (function trace) is a mechanism in the kernel for observing function execution, traditionally for debugging purposes. Ftrace also uses trampolines, but in a slightly different way. It is enabled when compiled with the `CONFIG_FUNCTION_TRACER` kernel configuration. When enabled, most files in the kernel are compiled with the `-pg` and `-mnop-mcount=...` flags. The `-pg` flag enables profiling, which adds some additional CPU instructions to the start of most global functions. The `-mnop-mcount=...` flag makes it so these CPU instructions are [`NOP` (No Operation)](https://en.wikipedia.org/wiki/NOP_(code)) instructions as apposed to the normal profiling logic. 

So by default, when a function is called, the CPU encounters a few `NOP` instructions which it will ignore and then continue the actual function. However, at runtime the kernel can replace these `NOP` instructions with a trampoline to in the case of ftrace some tracing logic.

Without `CONFIG_FUNCTION_TRACER`, the the assembly code for a function might look like this:

```
some_kernel_func:
  PUSH RBP
  MOV RSP, RBP
  ...
  RET
```

With `CONFIG_FUNCTION_TRACER` but without any tracing enabled, the assembly code for a function might look like this:
```
# Symbol where instructions like: CALL some_kernel_func will jump to.
some_kernel_func:
  NOP
  NOP
# actual start of the function
  PUSH RBP
  MOV RSP, RBP
  ...
  RET
```

The `NOP` instructions are ignored by the CPU as if they were not there.
 
Then at runtime, the kernel can replace the `NOP` instructions with code that adds the tracing. For example, the assembly code for a function with tracing enabled might look like this:

```
# Symbol where instructions like: CALL some_kernel_func will jump to.
some_kernel_func:
  CALL trace_call
  NOP
# actual start of the function
  PUSH RBP
  MOV RSP, RBP
  [...]
  RET
```

In this case the `trace_call` function would implement the actual tracing logic. The second `NOP` is here to illustrate that code of different sizes can be patched into the room left by the `NOP` instructions.

It should be noted that certain directories, files or functions in the kernel are excluded. For example the trace subsystem (kernel/trace) is excluded and any functions with the `notrace` attribute. This is done to protect the user from infinite recursion or breaking assumptions in critical code sections of the kernel.

## BPF trampolines

[:octicons-tag-24: v5.5](https://github.com/torvalds/linux/commit/fec56f5890d93fc2ed74166c397dc186b1c25951)

BPF programs can be attached to any function in the kernel that has enough `NOP` instructions. When this is done, a "BPF trampoline" is created. Somewhat confusing since normally the section of `NOP` instructions is called a trampoline, but in this case when we talk about a "BPF trampoline" we are talking about the code we jump to from a trampoline.

When a BPF program is attached, the kernel patches the `NOP` instruction on the original function like so:

```
# Symbol where instructions like: CALL some_kernel_func will jump to.
some_kernel_func:
  CALL generated_bpf_trampoline
  RET
# actual start of the function
  PUSH RBP
  MOV RSP, RBP
  [...]
  RET
```

This does not just add some additional logic like when tracing, it effectively redirects the execution of the original function to the generated BPF trampoline. The trampoline can now choose to never call the original (the `freplace` use case), call a BPF program before the original (the `fentry` use case), after the original (the `fexit` use case) or modify the return value of the original (the `fmodify_return` use case). It is also this generated trampoline that allows multiple programs to co-exist on the same function.

The generated BPF trampoline is architecture specific dynamically generated machine code that boils down to:

* Allocate room on the stack for all function arguments + return value
* Copy all arguments to the stack, from the registers specified by the calling convention
* For each `fentry` program attached
    * Call the `fentry` program with a pointer to the stack as context (if any is attached)
        * Disable CPU migration (preemption on older kernels) before calling the program, re-enable after
        * If stats tracking is enabled, start timer before execution and add run time to stats after execution
* For each `fmodify_return` program attached
    * Call a `fmodify_return` program with a pointer to the stack as context (if any is attached)
        * Disable CPU migration (preemption on older kernels) before calling the program, re-enable after
        * If stats tracking is enabled, start timer before execution and add run time to stats after execution
        * If the return value is non zero, return it instead of continuing
* Call the original function (unless a `fmodify_return` program returned a non `0` value)
* Copy return value from register onto the stack
* For each `fexit` program attached
    * Call a `fexit` program with a pointer to the stack as context (if any is attached)
        * Disable CPU migration (preemption on older kernels) before calling the program, re-enable after
        * If stats tracking is enabled, start timer before execution and add run time to stats after execution
* Return the return value from the stack

### Architecture support

Since BPF trampolines are architecture specific, support for a given architecture may be added later than the initial support for BPF trampolines. Here is a table of when support was added for each architecture:

| Architecture | Support added |
|--------------|---------------|
| X86-64       | [:octicons-tag-24: v5.5](https://github.com/torvalds/linux/commit/fec56f5890d93fc2ed74166c397dc186b1c25951) |
| ARM64        | [:octicons-tag-24: v6.0](https://github.com/torvalds/linux/commit/efc9909fdce00a827a37609628223cd45bf95d0b) |
| RISC-V       | [:octicons-tag-24: v6.3](https://github.com/torvalds/linux/commit/49b5e77ae3e214acff4728595b4ac7bf776693ca) |
| S390         | [:octicons-tag-24: v6.3](https://github.com/torvalds/linux/commit/528eb2cb87bc1353235a6384696b4849bde8b0ba) |

Other architectures currently lack support for BPF trampolines.

### fentry/fexit/fmodify_return

[:octicons-tag-24: v5.5](https://github.com/torvalds/linux/commit/f1b9509c2fb0ef4db8d22dac9aef8e856a5d81f6)

The [`BPF_PROG_TYPE_TRACING`](../program-type/BPF_PROG_TYPE_TRACING.md) fentry/fexit/fmodify_return programs make use of these trampolines to attach and execute just before entering a function or right after it exits. Since the trampoline is essentially just a function call, the overhead is very minimal. So a way faster alternative to a [kprobe](../program-type/BPF_PROG_TYPE_KPROBE.md) which uses an interrupt, and thus a context switch which is way more expensive.

Not only native kernel functions can have these blank spots for trampolines. When BPF programs are JIT-ed, the kernel also gives them these spots which can be instrumented in the same manner. This allows fentry/fexit/fmodify_return programs to attach to other programs, for the purposes of observability.

### Program replacement

[:octicons-tag-24: v5.6](https://github.com/torvalds/linux/commit/be8704ff07d2374bcc5c675526f95e70c6459683)

Trampolines are also used to implement [freplace programs](../program-type/BPF_PROG_TYPE_EXT.md), where one program replaces another. When a freplace program is attached, it installs a trampoline before the original program, jumps to the extension program and executes it instead, then returns without calling the original.

### LSM

[:octicons-tag-24: v5.7](https://github.com/torvalds/linux/commit/9e4e01dfd3254c7f04f24b7c6b29596bc12332f3)

[LSM](../program-type/BPF_PROG_TYPE_LSM.md) programs also attach via trampolines very similar to `fexit`/`fmodify_return` programs. The kernel defines placeholder functions for every hook, always starting with the prefix `bpf_lsm_`. These placeholders simply return the default return value. When attached, the LSM program acts as `fexit` or `fmodify_return` probe for the purposes of the BPF trampoline on these placeholder hooks.
