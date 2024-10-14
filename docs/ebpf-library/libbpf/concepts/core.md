# BPF CO-RE

BPF CO-RE stand for Compile Once - Run Everywhere
It's a **concept** to build cross-version kernel eBPF application by building in a single binary by bringing together the [BTF] type information, [libbpf] and the compiler. 

## Problem of portability
eBPF programs use the memory and data structures from the kernel. Between different kernel versions, some structures can be modified, altering the memory layout. 
Another problem can be the renaming of a field in the structure. 
If a BPF application uses one of these modified or renamed fields, the program will no longer be compatible.

## Export kernel information

Libbpf relies on the [BTF] information from the actual running Kernel who expose itself BTF information at `/sys/kernel/btf/vmlinux`.
It include all kernel types and structures layout.
A header file `vmlinux.h`, can be generate using [bpftool] :
```sh
bpftool btf dump file /sys/kernel/btf/vmlinux format c > vmlinux.h
```

This eliminate the dependency of the kernel headers

## Emit BTF relocation's with Clang

Clang was extended to emit BTF relocation's. These relocation's capture high-level descriptions of what information the BPF program intends to access.
The compiled BPF program is stored in an ELF (Executable and Link-able Format) object file. This file contains BTF type information and Clang-generated relocation's. 
The ELF format allows libbpf to process and adjust the BPF program for the target kernel dynamically.

## Use Libbpf as CO-RE library and loader

When you run your loader program with libbpf, it serves as the BPF program loader. It takes the compiled BPF ELF object file and post-processing it as necessary. It sets up various kernel objects (maps, programs, etc.) and triggers BPF program loading and verification. 
Libbpf uses the BTF information to match the types and fields in the BPF program with those in the running kernel, adjusting offsets and other relocatable data to ensure the program functions correctly on the specific kernel.

## Examples

Lists of examples programs using libbpf can be found on GitHub [libbpf-bootstrap]

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome


[BTF]: ../../../concepts/btf.md
[bpftool]: https://github.com/libbpf/bpftool
[libbpf]: ../../libbpf/index.md
[libbpf-bootstrap]: https://github.com/libbpf/libbpf-bootstrap
