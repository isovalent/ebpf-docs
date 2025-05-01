# Concepts

## [BPF CO-RE] (Compile Once - Run Everywhere)
It enables the development of portable BPF programs by integrating kernel BTF type information, Clang compiler support for relocation, and libbpf loader adjustments, ensuring compatibility across different kernel versions without runtime compilation overhead.

[BPF CO-RE]: core.md

## [BTF] (BPF Type Format)
It is compact, efficient format for describing C program type information. It enables runtime accessibility of kernel types crucial for BPF program development and verification.

[BTF]: btf.md

## [ELF] (Executable and Linkable Format)
It is a standard and versatile file format for executable files, object code, shared libraries, and core dumps on Unix-alike systems. The ELF format is extensible and cross-platform by design thus supporting different CPUs or ISA with different endiannesses and address sizes.

[ELF]: elf.md
