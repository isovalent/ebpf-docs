---
title: "Timers"
description: "This page explains the concept of timers in eBPF. It explains what timers are, how to use them, and when to use them."
---
# eBPF timers

[:octicons-tag-24: v5.15](https://github.com/torvalds/linux/commit/b00628b1c7d595ae5b544e059c27b1f5828314b4)

Timers allow eBPF programs to schedule the execution of an eBPF function at a later time. Use cases for this feature include garbage collection of map values or performing periodic checks. For example, we might want to prune DNS records from an LRU map if their TTL has expired to pro-actively make room instead of risking entries with valid TTLs being pruned due to inactivity.

Timers are stored in map values as `struct bpf_timer` fields. For example:

```c
struct map_elem {
    int counter;
    struct bpf_timer timer;
};

struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, 1000);
    __type(key, int);
    __type(value, struct map_elem);
} hmap SEC(".maps");
```

The definition of such a timer is: `#!c struct bpf_timer { __u64 :64; __u64 :64; };`.

!!! note
    Only programs with CAP_BPF are allowed to use bpf_timer.

The timers are attached to the life cycle of the map, if the map is freed/deleted, the all pending timers in that map will be canceled.

Pending timers will keep a reference to the program containing the callback, so even if no other references exist, programs will stay loaded until all timers have fired or are canceled.

A timer has to be initialized with the [`bpf_timer_init`](../helper-function/bpf_timer_init.md) helper function. After initialization a callback can be assigned to the timer with the [`bpf_timer_set_callback`](../helper-function/bpf_timer_set_callback.md) helper function. Lastly the timer is started with the [`bpf_timer_start`](../helper-function/bpf_timer_start.md) helper function. A pending timer can also be canceled with the [`bpf_timer_cancel`](../helper-function/bpf_timer_cancel.md) helper function.

These three helper calls do not necessarily have to happen in the same program at the same time. The following use case is valid:

* map1 is shared by `prog1`, `prog2`, `prog3`.
* `prog1` calls `bpf_timer_init` for some `map1` elements
* `prog2` calls `bpf_timer_set_callback` for some `map1` elements.
  * Those that were not `bpf_timer_init`-ed will return `-EINVAL`.
* `prog3` calls `bpf_timer_start` for some `map1` elements.
  * Those that were not both `bpf_timer_init`-ed and `bpf_timer_set_callback`-ed will return `-EINVAL`.


[`bpf_timer_init`](../helper-function/bpf_timer_init.md) and [`bpf_timer_set_callback`](../helper-function/bpf_timer_set_callback.md) will return `-EPERM` if map doesn't have user references (is not held by open file descriptor from user space and not pinned in bpffs).

The callback passed to the timer has the following signature `#!c static int callback_fn(void *map, {map key type} *key, {map value type} *value)`.
The callback is invoked with a pointer to the map, map key, and map value associated with the timer. It has no context unlike normal eBPF program execution, and thus is unable to perform work that requires operating on a context or helper side-effects. Its only input and output are maps.

!!! note
    The callback function *must* always return `0`, otherwise the verifier will reject the program.

A callback can chose to re-schedule its own timer by calling [`bpf_timer_start`](../helper-function/bpf_timer_start.md) on `value->timer`. Thus making it possible to not just have a one shot delay from a given eBPF program run, but to have a periodic function running endlessly after just a single trigger event.


