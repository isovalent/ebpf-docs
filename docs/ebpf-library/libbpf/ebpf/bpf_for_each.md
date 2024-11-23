---
title: "Libbpf eBPF macro 'bpf_for_each'"
description: "This page documents the 'bpf_for_each' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `bpf_for_each`

[:octicons-tag-24: v1.2.0](https://github.com/libbpf/libbpf/releases/tag/v1.2.0)

The `bpf_for_each` macro is used to make looping over open-coded iterators easier.

## Definition

```c
#define bpf_for_each(type, cur, args...) for (							\
	/* initialize and define destructor */							\
	struct bpf_iter_##type ___it __attribute__((aligned(8), /* enforce, just in case */,	\
						    cleanup(bpf_iter_##type##_destroy))),	\
	/* ___p pointer is just to call bpf_iter_##type##_new() *once* to init ___it */		\
			       *___p __attribute__((unused)) = (				\
					bpf_iter_##type##_new(&___it, ##args),			\
	/* this is a workaround for Clang bug: it currently doesn't emit BTF */			\
	/* for bpf_iter_##type##_destroy() when used from cleanup() attribute */		\
					(void)bpf_iter_##type##_destroy, (void *)0);		\
	/* iteration and termination check */							\
	(((cur) = bpf_iter_##type##_next(&___it)));						\
)
```

## Usage

This macro makes looping over open coded iterators easier. Open coded iterators were introduced in kernel [:octicons-tag-24: v6.4](https://github.com/torvalds/linux/commit/06accc8779c1d558a5b5a21f2ac82b0c95827ddd) and allow for iterating/looping over specific kernel data structures.

Normally, you would have to write iteration logic manually like:
```c
SEC("raw_tp/sys_enter")
int my_example(const void *ctx)
{
	struct bpf_iter_task task_it;
	struct task_struct *task_ptr;

    // Initialize the iterator, request to iterate over all processes on the system
	bpf_iter_task_new(&task_it, NULL, BPF_TASK_ITER_ALL_PROCS);

    // Loop until `bpf_iter_task_next` returns NULL
    while((task_ptr = bpf_iter_task_next(&task_it))) {
	    // Do something with the task pointer

        if (/* We found the task we are looking for*/)
            break;
    }

	bpf_iter_task_destroy(&task_it);
	return 0;
}
```

There are a few different iterator types, all follow the same naming convention: `bpf_iter_<type>_{new,next,destroy}`. The `bpf_for_each` macro makes use of this to simplify the iteration logic the user has to write.

The first argument is the `<type>` part of the iterator functions. The second argument is the variable name that will be used to store the current iterator value. And the rest of the arguments are passed to the `bpf_iter_<type>_new` function, the exact meaning depends on the iterator type.

### Example

```c
SEC("raw_tp/sys_enter")
int my_example(const void *ctx)
{
	struct task_struct *task_ptr;
    bpf_for_each(task, task_ptr, NULL, BPF_TASK_ITER_ALL_PROCS) {
        // Do something with the task pointer

        if (/* We found the task we are looking for*/)
            break;
    }

	return 0;
}
```

