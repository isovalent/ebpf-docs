---
title: "Libbpf eBPF macro 'bpf_for'"
description: "This page documents the 'bpf_for' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `bpf_for`

[:octicons-tag-24: v1.2.0](https://github.com/libbpf/libbpf/releases/tag/v1.2.0)

The `bpf_for` macro is used to make writing a for loop with an open coded iterator easier.

## Definition

```c
#define bpf_for(i, start, end) for (								\
	/* initialize and define destructor */							\
	struct bpf_iter_num ___it __attribute__((aligned(8), /* enforce, just in case */	\
						 cleanup(bpf_iter_num_destroy))),		\
	/* ___p pointer is necessary to call bpf_iter_num_new() *once* to init ___it */		\
			    *___p __attribute__((unused)) = (					\
				bpf_iter_num_new(&___it, (start), (end)),			\
	/* this is a workaround for Clang bug: it currently doesn't emit BTF */			\
	/* for bpf_iter_num_destroy() when used from cleanup() attribute */			\
				(void)bpf_iter_num_destroy, (void *)0);				\
	({											\
		/* iteration step */								\
		int *___t = bpf_iter_num_next(&___it);						\
		/* termination and bounds check */						\
		(___t && ((i) = *___t, (i) >= (start) && (i) < (end)));				\
	});											\
)
```

## Usage

This macro makes writing a for loop with an open coded iterator easier. `bpf_for` specifically uses the numeric open coded iterator, which instead of looping over kernel structures loops over a range of numbers.
The `bpf_for` macro is a neat shorthand for manually writing the iteration logic.

The reason you might want to use this instead of a traditional `for` loop is to get a large loop count without making the verifier upset. Using the open code numeric iterator allows for a large number of iterations without inflating complexity and while keeping the current scope. For details see the [loops concept](../../../linux/concepts/loops.md) page.

### Example

```c
SEC("raw_tp/sys_enter")
int iter_next_rcu(const void *ctx)
{
	int v;

    // Will print 2, 3, 4, 5
    bpf_for(v, 2, 5) {
        bpf_printk("X = %d", v);
    }

	return 0;
}
```
