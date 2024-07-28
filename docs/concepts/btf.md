# BTF

BTF (BPF Type Format) is a metadata format designed for encoding debug information related to BPF programs and maps. It focus on describing data types, function information for defined subroutines.
This debug information serves various purposes including map visualization, function signature enhancement for BPF programs, and aiding in the generation of annotated source code, JIT-ed code, and verifier logs.

The BTF specification is divided into two main parts:

* BTF Kernel API: This defines the interface between user space and the kernel. Before usage, the kernel validates the BTF information provided.
* BTF ELF File Format: This establishes the contract between the ELF file and the libbpf loader in user space.

It was created as an alternative to DWARF debug information. The BTF is more space-efficient due to his [deduplication algorithm] while remaining expressive enough to have all the type information of a C programs. 

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

[deduplication algorithm]: https://nakryiko.com/posts/btf-dedup/
