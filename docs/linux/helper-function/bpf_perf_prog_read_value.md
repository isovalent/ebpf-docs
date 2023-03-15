# Helper function `bpf_perf_prog_read_value`

<!-- [FEATURE_TAG](bpf_perf_prog_read_value) -->
[:octicons-tag-24: v4.15](https://github.com/torvalds/linux/commit/4bebdc7a85aa400c0222b5329861e4ad9252f1e5)
<!-- [/FEATURE_TAG] -->

This helper retrieves the value of an event counter.

## Definition

**Returns**
0 on success, or a negative error in case of failure.

`#!c static long (*bpf_perf_prog_read_value)(struct bpf_perf_event_data *ctx, struct bpf_perf_event_value *buf, __u32 buf_size) = (void *) 56;`

## Usage

For an eBPF program attached to a perf event, retrieve the value of the event counter associated to `ctx` and store it in the structure pointed by `buf` and of size `buf_size`. Enabled and running times are also stored in the structure (see description of helper [`bpf_perf_event_read_value`](bpf_perf_event_read_value.md) for more details).

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->

<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_PERF_EVENT](../program-type/BPF_PROG_TYPE_PERF_EVENT.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

<!-- TODO add C / Rust example -->
