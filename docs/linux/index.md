---
title: eBPF on Linux
---
# eBPF on Linux

<!-- TODO(dylandreimerink): linux is the first platform but not the only one -->

## Programs

The most central part of eBPF are the programs. eBPF programs can be attached at different points in the kernel and will be called like a function. Programs can have lots of different purposes, they can for example record information, modify information, make decisions and cause side effects. Where a program can attach and what it is allowed to do depends on its [program type](./program-type/index.md).

Programs get called with a context which is a struct with information the kernel is making easily available to the program. Typical examples are a [socket buffer](./program-context/__sk_buff.md) or CPU registers. What context is passed to a program depends on its type.

Program like functions also have return values, the meaning of which is again determined by the program type. A return value can for example indicate the amount of bytes of a packet to keep or be an enum of actions which could be taken like discarding a packet, accepting it, or redirecting it.

eBPF program are typically written in C and compiled with LLVM, but this isn't necessarily the only way to do it. Any program which can generate byte-code (following the eBPF instruction set) can author eBPF programs. eBPF programs are typically serialized into a relocatable ELF file.

Ultimately eBPF programs are loaded into the kernel using the [BPF syscall](./syscall/index.md), the userspace program that does this is refereed to as a loader. In practice loaders range from applications that just load the eBPF program to complex systems that constantly interacts with multiple programs and maps to provide advanced features. Loaders often use [loader libraries](./../ebpf-library/index.md) to provide higher-level APIs than the syscall to ease development.

When the loader loads a program the kernel will verify that the program is "safe". This job is done by a component of the kernel called the verifier. "safe" in this context means that programs are not allowed to crash the kernel or break critical components. eBPF programs have to pass quite a number of stringent requirements before being allowed anywhere near kernel memory. For more details checkout the verifier page.

## Helper functions

Programs on their own are quite limited, they can read from and write to a local stack, perform maths on registers, call internal functions and do conditional jumps. All of this is within its own little bubble. The final thing programs can do is call so called "helper functions". These are actually regular C functions defined by the kernel. These functions form a sort of internal API/ABI between the eBPF programs and the kernel. These helpers can allow eBPF programs to perform tasks they otherwise couldn't be cause it wouldn't get past the verifier.

These helper functions, take up to 5 arguments and return a single return value. Not every program type can execute every helper call to enforce the same restrictions the verifier does.

Helper functions have a large variety of purposes ranging from simply getting some additional information like what CPU core we are executing on to invoking major side effects like redirecting packets. For a complete overview checkout the [helper functions](./helper-function/index.md) page.

## Maps

eBPF maps are datastructures that live in the kernel. Both eBPF programs and userspace programs can access these maps and thus they are the communication layer between eBPF programs and userspace as well as a place to persist data between program calls. Maps like all other BPF objects are shared over the whole host, and multiple programs can access the same maps at the same time. Thus maps can also be used to transfer information between programs of different types at different attached points.

Examples of these maps are [`BPF_MAP_TYPE_ARRAY`](./map-type/BPF_MAP_TYPE_ARRAY.md) which is an array of arbitrary values or a [`BPF_MAP_TYPE_HASH`](./map-type/BPF_MAP_TYPE_HASH.md) which is a hash map with arbitrary key and value types. For more details check out the [map type overview](./map-type/index.md).

## Objects

eBPF programs and maps are BPF objects along with some others we didn't mention yet. All of these objects managed in roughly the same way. Such BPF objects are created by a loader which gets a file descriptor to that object. The file descriptor is used to interact with the object further, but it also is a reference that keeps the object "alive". Objects are free-ed as soon as there exist no more references to that object.

Applications can transfer copies of these file descriptors to other processes via interprocess communication techniques like unix sockets which is quite universal. A more eBPF special technique is called pinning, which allows the loader to reference a BPF object with a special file called a pin. These pins can only be made in the special BPF File System which needs to be mounted somewhere (typically at /sys/bpf but this can change between distros). As long as a pin exists it will keep the object it refers to alive. These pins can be read by any program that has permissions to access the pin file and get a reference of the object that way. Thus, multiple program can share the same objects at the same time.

## Capabilities

<!-- TODO explain CAP_BPF and others -->


