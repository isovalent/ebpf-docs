# KFunc `bpf_iter_num_new`

<!-- [FEATURE_TAG](bpf_iter_num_new) -->
[:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/06accc8779c1d558a5b5a21f2ac82b0c95827ddd)
<!-- [/FEATURE_TAG] -->

Initialize a new number iterator.

## Definition

This kfunc initializes the iterator `it`, priming it to do a numeric iteration from `start` to `end` where `start <= i < end`.

If any of the input arguments are invalid, constructor should make sure to still initialize it such that subsequent bpf_iter_num_next() calls will return NULL. I.e., on error, return error and construct empty iterator.

<!-- [KFUNC_DEF] -->
`#!c int bpf_iter_num_new(struct bpf_iter_num *it, int start, int end)`
<!-- [/KFUNC_DEF] -->

## Usage

These open coded iterators allow you to do iteration over kernel objects like [iterator programs](../../program-type/BPF_PROG_TYPE_TRACING/#iterator). Except we can just use a classic iterator pattern.

This specific iterator is a numeric iterator which allows us to iterate over a range of numbers. Effectively this gives us yet another way to do a for loop in BPF. The advantage of this method over bounded loops is that the verifier only has to check two states, the [`bpf_iter_num_next`](bpf_iter_num_next.md) can return NULL or a integer between `start` and `end`. So no matter the size of the range, the verifier only has to check these two states, it understands the contract that the iterator will always produce a NULL at some point and terminate.

The advantages over the [`bpf_loop`](../helper-function/bpf_loop.md) is that we keep the same scope, and we don't have to move variables back and forth between the main program and the callback via the context. 

### Program types

The following program types can make use of this kfunc:

<!-- [KFUNC_PROG_REF] -->
- [BPF_PROG_TYPE_CGROUP_SKB](../../program-types/BPF_PROG_TYPE_CGROUP_SKB.md)
- [BPF_PROG_TYPE_CGROUP_SOCK_ADDR](../../program-types/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md)
- [BPF_PROG_TYPE_LSM](../../program-types/BPF_PROG_TYPE_LSM.md)
- [BPF_PROG_TYPE_LWT_IN](../../program-types/BPF_PROG_TYPE_LWT_IN.md)
- [BPF_PROG_TYPE_LWT_OUT](../../program-types/BPF_PROG_TYPE_LWT_OUT.md)
- [BPF_PROG_TYPE_LWT_SEG6LOCAL](../../program-types/BPF_PROG_TYPE_LWT_SEG6LOCAL.md)
- [BPF_PROG_TYPE_LWT_XMIT](../../program-types/BPF_PROG_TYPE_LWT_XMIT.md)
- [BPF_PROG_TYPE_NETFILTER](../../program-types/BPF_PROG_TYPE_NETFILTER.md)
- [BPF_PROG_TYPE_SCHED_ACT](../../program-types/BPF_PROG_TYPE_SCHED_ACT.md)
- [BPF_PROG_TYPE_SCHED_CLS](../../program-types/BPF_PROG_TYPE_SCHED_CLS.md)
- [BPF_PROG_TYPE_SK_SKB](../../program-types/BPF_PROG_TYPE_SK_SKB.md)
- [BPF_PROG_TYPE_SOCKET_FILTER](../../program-types/BPF_PROG_TYPE_SOCKET_FILTER.md)
- [BPF_PROG_TYPE_STRUCT_OPS](../../program-types/BPF_PROG_TYPE_STRUCT_OPS.md)
- [BPF_PROG_TYPE_SYSCALL](../../program-types/BPF_PROG_TYPE_SYSCALL.md)
- [BPF_PROG_TYPE_TRACING](../../program-types/BPF_PROG_TYPE_TRACING.md)
- [BPF_PROG_TYPE_XDP](../../program-types/BPF_PROG_TYPE_XDP.md)
<!-- [/KFUNC_PROG_REF] -->

### Example

```c
struct bpf_iter_num it;
int *v;

bpf_iter_num_new(&it, 2, 5);
while ((v = bpf_iter_num_next(&it))) {
    bpf_printk("X = %d", *v);
}
bpf_iter_num_destroy(&it);
```

Above snippet should output "X = 2", "X = 3", "X = 4". Note that 5 is
exclusive and is not returned. This matches similar APIs (e.g., slices
in Go or Rust) that implement a range of elements, where end index is
non-inclusive.

Libbpf also provides macros to provide a more natural feeling way to write the above:
```c
int *v;

bpf_for(v, start, end) {
    bpf_printk("X = %d", *v);
}
```

There is also a repeat macros:
```c
int i = 0;
bpf_repeat(5) {
    bpf_printk("X = %d", i);
    i++;
}
```
