# KFunc `bpf_get_fsverity_digest`

<!-- [FEATURE_TAG](bpf_get_fsverity_digest) -->
[:octicons-tag-24: v6.8](https://github.com/torvalds/linux/commit/67814c00de3161181cddd06c77aeaf86ac4cc584)
<!-- [/FEATURE_TAG] -->

Get the fs-verity digest of a file.

## Definition

<!-- [KFUNC_DEF] -->
`#!c int bpf_get_fsverity_digest(struct file *file, struct bpf_dynptr_kern *digest_ptr)`
<!-- [/KFUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_LSM](../../program-types/BPF_PROG_TYPE_LSM.md)
- [BPF_PROG_TYPE_TRACING](../../program-types/BPF_PROG_TYPE_TRACING.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

