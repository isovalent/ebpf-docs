---
title: "Program Type 'BPF_PROG_TYPE_EXT'"
description: "This page documents the 'BPF_PROG_TYPE_EXT' eBPF program type, including its defintion, usage, program types that can use it, and examples."
---
# Program type `BPF_PROG_TYPE_EXT`

<!-- [FEATURE_TAG](BPF_PROG_TYPE_EXT) -->
[:octicons-tag-24: v5.6](https://github.com/torvalds/linux/commit/be8704ff07d2374bcc5c675526f95e70c6459683)
<!-- [/FEATURE_TAG] -->

Extension programs can be used to dynamically extend another BPF program.

## Usage

These programs can be used to replace global functions in already loaded BPF programs. Global functions are verified individually by the verifier based on their types only.
Hence the global function in the new program which types match older function can
safely replace that corresponding function.

Programs of this type are typically placed in an ELF section prefixed with `freplace/`

The main use case for extensions is to provide generic mechanism to plug external programs into policy program or function call chaining. The [libxdp](https://github.com/xdp-project/xdp-tools/tree/master/lib/libxdp) project uses this functionality to implement XDP program chaining from a dispatcher program.

This new function/program is called 'an extension' of old program. At load time
the verifier uses (attach_prog_fd, attach_btf_id) pair to identify the function
to be replaced. The BPF program type is derived from the target program into
extension program. 

!!! note
    The verifier allows only one level of replacement. Meaning that the extension program cannot recursively extend an extension.

!!! note
    The extension program has its own stack + depth limit. So the combined limit increases to 1024 stack and 16 calls deep.

## Context

The context of the extension program depends on the function to be replaced. Function by function verification of global function supports scalars and pointer to context only. Hence program extensions are supported for such class of global functions only. In the future the verifier will be extended with support to pointers to structures, arrays with sizes, etc.

## Attachment

Program extensions are attached via a [BPF link](../syscall/BPF_LINK_CREATE.md). The `prog_fd` is set to the file descriptor of the extension program, the `target_fd` set to the file descriptor of the program to be extended, `target_btf_id` set to the BTF ID of the global function to replace and the `attach_type` set to `0`.

## Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
