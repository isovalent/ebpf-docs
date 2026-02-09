---
title: "Syscall command 'BPF_ENABLE_STATS'"
description: "This page documents the 'BPF_ENABLE_STATS' eBPF syscall command, including its definition, usage, program types that can use it, and examples."
---
# BPF Syscall `BPF_ENABLE_STATS` command

<!-- [FEATURE_TAG](BPF_ENABLE_STATS) -->
[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/d46edd671a147032e22cfeb271a5734703093649)
<!-- [/FEATURE_TAG] -->

This syscall command is used to temporarily enable statistics tracking, globally for all programs. This allows for the benchmarking or monitoring of programs <nospell>in situ</nospell>.

## Return value

On success this syscall command returns a file descriptor. Unlike other BPF related file descriptor this file descriptor can not be pinned. Statistics tracking is guaranteed to be working until this file descriptor is released/closed.

## Usage

When enabled, the kernel will start to update the `run_time_ns` and `run_cnt` fields of the `struct bpf_prog_info` associated with each loaded program. This information can be queries with the [`BPF_OBJ_GET_INFO_BY_FD`](BPF_OBJ_GET_INFO_BY_FD.md) syscall command when used on the file descriptor of a program.

The `run_time_ns` value holds the accumulated amount of nano seconds the BPF program has ran and the `run_cnt` the amount of times it ran. Users might be interested in these values for the sake of monitoring the amount of CPU time BPF programs take up or to benchmark programs under real world conditions.

These statistics are not enabled by default since BPF programs might be called quite frequently and recording this information increases the overhead of each BPF program run. It is therefore recommended to not permanently enable this feature in production environments or to do so when CPU usage can be spared.

The typical usage pattern would be:

1. Enable statistics tracking with this command. 
2. Take the base-line measurements of the programs you are interested in.
3. Wait for a certain about of time.
4. Take new measurements.
5. Calculate the difference to see the by how much they incremented.

!!! note
    This syscall command requires the `CAP_SYS_ADMIN` to use.

## Attributes

### `type`

[:octicons-tag-24: v5.8](https://github.com/torvalds/linux/commit/d46edd671a147032e22cfeb271a5734703093649)

This field specifies the type of statistic you would like to enable. Currently the only valid value is `BPF_STATS_RUN_TIME`
