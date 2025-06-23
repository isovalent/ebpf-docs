---
title: "Libbpf eBPF macro 'bpf_repeat'"
description: "This page documents the 'bpf_repeat' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `bpf_repeat`

[:octicons-tag-24: v1.2.0](https://github.com/libbpf/libbpf/releases/tag/v1.2.0)

The `bpf_repeat` macro makes it easier to write a loop that repeats X times using open coded iterators.

## Definition

```c
#define bpf_repeat(N) for (									\
	/* initialize and define destructor */							\
	struct bpf_iter_num ___it __attribute__((aligned(8), /* enforce, just in case */	\
						 cleanup(bpf_iter_num_destroy))),		\
	/* ___p pointer is necessary to call bpf_iter_num_new() *once* to init ___it */		\
			    *___p __attribute__((unused)) = (					\
				bpf_iter_num_new(&___it, 0, (N)),				\
	/* this is a workaround for Clang bug: it currently doesn't emit BTF */			\
	/* for bpf_iter_num_destroy() when used from cleanup() attribute */			\
				(void)bpf_iter_num_destroy, (void *)0);				\
	bpf_iter_num_next(&___it);								\
	/* nothing here  */									\
)
```

## Usage

This macro makes writing a loop that repeats X times using open coded iterators easier. `bpf_repeat` is a simplified version of the [`bpf_for`](bpf_for.md) macro, which simply executes the loop body `N` times and does not make the iteration count available.

### Example

```c
SEC("raw_tp/sys_enter")
int iter_next_rcu(const void *ctx)
{
    bpf_repeat(64) {
        if (try_to_do_something())
            break;
    }
}
```
