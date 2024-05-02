---
title: "Verifier"
description: "This page explains the eBPF verifier. It explains what the verifier is, why it exists, and what features it has."
---
# Verifier

The verifier is a core component of the BPF subsystem. Its main responsibility is to ensure that the BPF program is "safe" to execute. It does this by checking the program against a set of rules. The verifier also performs some additional tasks, mainly optimizations for which it uses the information gathered during the verification process.

The verifier exists because BPF programs are translated into native machine code and executed in kernel mode. This means BPF programs can do really bad things to the system if they are not properly checked such as corrupting memory, leaking sensitive information, causing the kernel to crash or causing the kernel to hang/deadlock.

This model is a tradeoff between ease of use and performance. Once you are able to pass the verifier, there are no expensive runtime checks, so BPF programs can run at native speed. An alternative model with a virtual machine or interpreter would have been much slower.

## Basics

So what is this "safe" concept we have been talking about? The general idea is that BPF programs are not allowed to break the kernel in any way and it should not violate the security model of the system. This results in a long list of don'ts. Here is a non-exhaustive list of things that are not allowed to illustrate the point:

* Programs must always terminate (within a reasonable amount of time) - So no infinite loops or infinite recursion.
* Programs are not allowed to read arbitrary memory - Being able to read any memory would allow a program to leak sensitive information. There are exceptions, tracing programs have access to helpers that allow them to read memory in a controlled way. But these program types require root privileges and thus are not a security risk.
* Network programs are not allowed to access memory outside of packet bounds because adjacent memory could contain sensitive information. See the point above.
* Programs are not allowed to deadlock, so any held spinlocks must be released and only one lock can be held at a time to avoid deadlocks over multiple programs.
* Programs are not allowed to read uninitialized memory - This could leak sensitive information.

The list goes on. A lot of rules are conditional, there are additional rules per program type. Not all program types can use the same helper functions or access the same context fields. These restrictions are discussed in more detail in the pages about the different program types and helper functions.

### Analysis

The basic premise is that the verifier checks every possible permutation of a program mathematically. It starts by walking the code and constructing a graph based on branching instructions. It will reject any statically-dead-code unreachable code might be a link in an exploit chain.

Next the verifier starts at the top, setting the initial registers. R1 for example is almost always a pointer to the context. It walks over each instruction and updates the state of the registers and stack. This state contains information like smax32 (what is the largest 32 bit signed integer that could be in this register). It has many such variables which it can use to evaluate if a branch such as "if R1 > 123" is always taken, sometimes taken or never taken.

Every time the verifier encounters a branching instruction, it will fork the current state, queue one of the branches+state for later investigation and update states. For example, if I have a register R3 with a value between 10 and 30 and I then encounter a "if R3 > 20" instruction, one fork will have a R3 of 10-20 and the other 21-30. This is a very simple example, but it illustrates the point.

It also keeps track of linked registers. If I go R2 = R3, then do the above example, the verifier knows that R2 also has the same range as R3. This is commonly used for packet bounds checks.

The verifier also keeps track of data types, before I mentioned the pointer to a context. It also knows when we are dealing with normal numbers or pointers to map values for example. Every time an offset from the context is dereferenced for example it will check that access is allowed for the current program type and that the offset is within bounds of the context. It can also keep track of possible null values, such as those returned from map lookups. And uses that information to enforce that null checks are done before dereferencing pointers.

It uses this same type info tracking to assert that the correct parameters are passed to helper functions or function calls. The verifier can also use BTF to enforce that a map value contains a timer field for example or a spinlock. BTF is also used to enforce that the correct parameters are passed to fkuncs, that BTF func definitions match the actual BPF functions and that these BTF func definitions match callbacks.

The verifier will attempt to asses all queued states and branches. But to protect itself it has limits. It tracks the amount of instructions inspected, this is for any permutation, so the complexity of a program not only depends on the amount of instructions, but also on the amount of branches. The verifier only has a limited amount of storage for states, so infinite recursion doesn't consume to much memory. 

!!! note
    Until [:octicons-tag-24: v5.2](https://github.com/torvalds/linux/commit/c04c0d2b968ac45d6ef020316808ef6c82325a82) there was a hard 4k instruction limit and a 128k complexity limit. Afterwards both are 1 million.

## Features

### Tail calls

Tail calls allow a BPF program to call another BPF program, basically a GOTO to another program and not a function call. These programs are loaded and verified separately and thus do not count towards the complexity limit of the verifier. Therefore tail calls are a popular method to work around the verifier complexity limit by splitting to logic of a program into multiple programs.

For details check out the [Tail calls](tail-calls.md) page.

### Dead code elimination

The first iteration of dead code elimination was added in [:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/c131187db2d3fa2f8bf32fdf4e9a4ef805168467). From then on any dynamically dead code (reachable via conditional statement, but the condition is always true or always false) is replaced by NOP instructions. This doesn't yet eliminate dead code but renders it harmless (we don't want to JIT code that isn't checked, even if we never jump to it ourselves).

In :octicons-tag-24: v5.1 dead code elimination was added. The [first step](https://github.com/torvalds/linux/commit/e2ae4ca266a1c9a0163738129506dbc63d5cca80) was to convert the conditional branching instructions into unconditional jump instructions to avoid misprediction penalties. 

The [second step](https://github.com/torvalds/linux/commit/52875a04f4b26e7ef30a288ea096f7cfec0e93cd) was to actually remove the dead code. This requires recalculation of relative jumps and the adjustment of BTF line info.

The [third step](https://github.com/torvalds/linux/commit/a1b14abc009d9c13be355dbd4a4c4d47816ad3db) is to remove conditional jumps with empty bodies since they don't do anything.

The dead code elimination in v5.1 only happens for privileged programs. Since this is an optimization step, it is not strictly required and any bug might cause security issues, so by not performing it on unprivileged programs we can avoid potential privilege escalation.

### Bounded loops

[:octicons-tag-24: v5.3](https://github.com/torvalds/linux/commit/2589726d12a1b12eaaa93c7f1ea64287e383c7a5)

Before bounded loops were introduced, the verifier would reject any program that contained a loop. The workaround for a long time was to unroll loops in the compiler. This is not a great solution because it increases the size of the program and it is not always possible to unroll loops.

Bounded loops allow the verifier to check that a loop will always terminate. The downside is that to do so the verifier will check every permutation of the loop. So if you have a loop that goes up to 100 times with a body of 20 instructions and a few branches, then that loop counts for a few thousand instructions towards the complexity limit.

See the [Loops](loops.md) page for more details on doing loops in BPF.

### Function by function verification

[:octicons-tag-24: v5.6](https://github.com/torvalds/linux/commit/51c39bb1d5d105a02e29aa7960f0a395086e6342)

Before this feature every BPF-to-BPF function had to be `static`. Static functions are verified from the perspective of the caller. Every time a program invokes a function, verification continues in that function with the state of the arguments to prove invocation from every call site is safe. This means that the verifier might need to verify certain functions multiple times which is slow and drives up complexity.

This feature allows you to use global functions (functions without the `static` keyword). These have slightly different constraints. The verifier will assume no information about the arguments and will verify the function in isolation. This means that the verifier only needs to verify the function once, no matter how many times it is called. This is much faster and reduces complexity.

Additionally, global functions can be replaced by [freplace](../program-type/BPF_PROG_TYPE_EXT.md) programs because there are assumptions about these functions outside of their signature.

### Callbacks

[:octicons-tag-24: v5.13](https://github.com/torvalds/linux/commit/69c087ba6225b574afb6e505b72cb75242a3d844)

The [`bpf_for_each_map_elem`](../helper-function/bpf_for_each_map_elem.md) helper also introduced the concept of callbacks. This allows users to declare a static function that is not directly called by the BPF program but is passed as function pointer to a helper to be called.

In later versions this mechanism is also used for [timers](timers.md), bpf_find_vma, and [loops](loops.md).
