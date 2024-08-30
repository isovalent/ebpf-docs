---
title: "Loops"
description: "This page explains the limits of loops in eBPF. It explains different methods of looping and their pros, cons, and when you can use them."
---
# Loops in BPF

Loops in programming is a common concept, however, in BPF they can be a bit more complicated than in most environments. This is due to the verifier and the guaranteed "safe" nature of BPF programs. 

## Unrolling

Before [:octicons-tag-24: v5.3](https://github.com/torvalds/linux/commit/2589726d12a1b12eaaa93c7f1ea64287e383c7a5) loops in BPF bytecode were not allowed because the verifier wasn't smart enough to determine if a loop would always terminate. The workaround for a long time was to unroll loops in the compiler. Unrolling loops increases the size of a program and can only be done if the amount of iterations is known at compile time. To unroll a loop you can use the `#pragma unroll` pragma as such:

```c
#pragma unroll
for (int i = 0; i < 10; i++) {
    // do something
}
```

## Bounded loops

Since [:octicons-tag-24: v5.3](https://github.com/torvalds/linux/commit/2589726d12a1b12eaaa93c7f1ea64287e383c7a5) the verifier is smart enough to determine if a loop will stop or not. These are referred to as "bounded loops". Users of this feature still have to be careful though, because its easy to write a loop which makes your program too complex for the verifier to handle. The verifier will check every possible permutation of a loop, so if you have a loop that goes up to 100 times with a body of 20 instructions and a few branches, then that loop counts for a few thousand instructions towards the complexity limit.

A common mistake is to use variables with a huge range as the bounds for a loop. For example:

```c
void *data = ctx->data;
void *data_end = ctx->data_end;
struct iphdr *ip = data + sizeof(struct ethhdr);
if (ip + sizeof(struct iphdr) > data_end)
    return XDP_DROP;

if (ip + sizeof(struct iphdr) > data_end)
    return XDP_DROP;

for (int i = 0; i < ip->tot_len; i++) {
    // scan IP body for something
}
```

Since `ip->tot_len` is a 16 bit integer, the verifier will check the body for every possible value of `i` up to 65535. Depending on the instructions and branches in the body, you will run out of complexity very quickly. Most of the time scanning the first X bytes of a body is enough, so you can limit the loop to that:

```c
void *data = ctx->data;
void *data_end = ctx->data_end;
struct iphdr *ip = data + sizeof(struct ethhdr);
if (ip + sizeof(struct iphdr) > data_end)
    return XDP_DROP;

if (ip + sizeof(struct iphdr) > data_end)
    return XDP_DROP;

int max = ip->tot_len;
if (max > 100)
    max = 100;

for (int i = 0; i < max; i++) {
    // scan IP body for something
}
```

## Map iteration helper

Since [:octicons-tag-24: v5.13](https://github.com/torvalds/linux/commit/69c087ba6225b574afb6e505b72cb75242a3d844) it is possible to use the [bpf_for_each_map_elem](../helper-function/bpf_for_each_map_elem.md) helper to iterate over maps so you don't have to use loops for that (if the map type supports it). The helper is called with a pointer to a map and a callback function. The callback function is called for each element in the map. The callback function is passed the map, the key, the value and a context pointer. The context pointer can be used to pass information from the main program to the callback function and back. The return value of the callback can be used to break out of the loop early.

## Loop helper

Sometimes you really need to iterate over a huge range. For cases where any of the above solutions result in complexity issues the [bpf_loop](../helper-function/bpf_loop.md) helper functions was introduced in [:octicons-tag-24: v5.17](https://github.com/torvalds/linux/commit/e6f2dd0f80674e9d5960337b3e9c2a242441b326). The helper allows for loops up to 1 << 23 (~8 million) iterations. The helper guarantees that the loop will terminate without the verifier having to check each iteration. The body is a callback function with a `index` and `ctx` argument. The context can be any type, passed in from the main program and shared between iterations which can be used for both the input and output of the loop. The return value of the callback function can be used to continue or break out of the loop early.

## Numeric open coded iterators

In [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/06accc8779c1d558a5b5a21f2ac82b0c95827ddd) open-coded iterators were introduced. Which allow programs to iterate over kernel objects. The numeric iterator allows us to iterate over a range of numbers, allowing us to make a for loop.

The advantage of this method is that the verifier only has to check two states as opposed to the amount of iterations like with a bounded loop and we don't require a callback function like with the loop helper.

Every iterator type has a `bpf_iter_<type>_new` function to initialize the iterator, a `bpf_iter_<type>_next` function to get the next element, and a `bpf_iter_<type>_destroy` function to clean up the iterator. In the case of the numeric iterator, the `bpf_iter_num_new`, `bpf_iter_num_next` and `bpf_iter_num_destroy` functions are used.

The most basic example of a numeric iterator is:

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
int v;

bpf_for(v, start, end) {
    bpf_printk("X = %d", v);
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

At a 10,000-foot point of view this works because, `next` methods are the points of forking a verification state, which are conceptually similar to what verifier is doing when validating conditional jump. We branch out at a `call bpf_iter_<type>_next` instruction and simulate two outcomes: NULL (iteration is done) and non-NULL (new element is returned). NULL is simulated first and is supposed to reach exit without looping. After that non-NULL case is validated and it either reaches exit (for trivial examples with no real loop), or reaches another `call bpf_iter_<type>_next` instruction with the state equivalent to already (partially) validated one. State equivalency at that point means we technically are going to be looping forever without "breaking out" out of established "state envelope" (i.e., subsequent iterations don't add any new knowledge or constraints to the verifier state, so running 1, 2, 10, or a million of them doesn't matter). But taking into account the contract stating that iterator next method *has to* return NULL eventually, we can conclude that loop body is safe and will eventually terminate. Given we validated logic outside of the loop (NULL case), and concluded that loop body is safe (though potentially looping many times), verifier can claim safety of the overall program logic.
