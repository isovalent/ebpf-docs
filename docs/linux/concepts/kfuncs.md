---
title: KFuncs
description: This page explains the concept of kfunc in eBPF. It explains what kfunc is, how to use it, and when to use it.
---
# KFuncs

[:octicons-tag-24: v5.13](https://github.com/torvalds/linux/commit/e6ac2450d6dee3121cd8bbf2907b78a68a8a353d)

KFunc also known as a kernel function is a function within the kernel that has been annotated and specifically designated as being callable from eBPF programs. KFuncs are an alternative to [helper functions](../helper-function/index.md), a new method to provide similar functionality. 

Officially KFuncs are unstable, unlike helper functions, kfuncs have no UAPI guarantees. In practice this might mean that kfuncs can change or be removed between kernel versions. Though as with all features, the kernel community will try to avoid breaking changes, and will provide deprecation warnings when possible. Users of kfuncs might need to be more vigilant about changes in the kernel, and be prepared to update their programs more frequently or write more complex code to handle different kernel versions.

## Usage

Using a KFunc is fairly straightforward. The first step is to copy the function signature(return type, name, and parameters) of the kfunc we would like to call. These function signatures are usually found in the kernel source code or in [KFunc pages](../kfuncs/index.md).
Second step is to add the `extern` keyword, this tells the compiler that the function isn't defined in our compilation unit. Lastly we add the `__ksym` attribute which tells the loader that references to the function should be resolved with the address of a kernel symbol (kernel function).

After we have done this we can call the kfunc as if it was a normal function.

```c
#include <vmlinux.h>
#include <bpf/bpf_helpers.h>

extern struct task_struct *bpf_task_acquire(struct task_struct *p) __ksym;

extern void bpf_task_release(struct task_struct *p) __ksym;

SEC("tp_btf/task_newtask")
int BPF_PROG(task_acquire_release_example, struct task_struct *task, u64 clone_flags)
{
    struct task_struct *acquired;

    acquired = bpf_task_acquire(task);
    if (acquired)
            /*
                * In a typical program you'd do something like store
                * the task in a map, and the map will automatically
                * release it later. Here, we release it manually.
                */
            bpf_task_release(acquired);
    return 0;
}

char _license[] SEC("license") = "GPL";
```

!!! note
    The definition of `__ksym` is `#define __ksym __attribute__((section(".ksyms")))`

### Kernel modules

The [KFunc index](../kfuncs/index.md) includes all KFuncs defined in the Linux kernel sources. Depending on the KConfig used to compile the kernel not all of these might be available or might be available via a kernel module.

KFuncs can be dynamically added to the kernel via kernel modules, so both builtin and third party modules can add KFuncs. The usage mechanism is the same, but you might have to handle situations where a module isn't loaded.

## Parameter annotations

The parameters/arguments of a KFunc can be annotated with a number of suffixes. These indicate how they should be used. The [verifier](verifier.md) is aware of these annotations and will enforce the rules they imply.

### `__sz` annotation

A parameter with the `__sz` suffix is used to indeicate the size of the pointer.

Take the following example:

```c
void bpf_memzero(void *mem, int mem__sz)
```

In this case `mem__sz` indicates the size of the memory pointed to by `mem`. The verifier will enforce that `mem` is a valid pointer and that the size of the `mem__sz` will not result in an out of bounds access.

### `__szk` annotation

A parameter with the `__szk` suffix is similar to `__sz` but its value has to be a constant at compile time. This is typically used when the parameter before is a pointer to a kernel structure which might change size between kernel versions. In which case the `sizeof()` that struct is to be used.

### `__k` annotation

A parameter with the `__k` suffix is used to indicate that its value must be a scalar value (just a number) and a well-known constant. This is typically used for BTF IDs.

For example:

```c
void *bpf_obj_new_impl(u64 local_type_id__k, void *meta__ign)
```

The `bpf_obj_new_impl` KFunc creates a new object with the shape of the given BTF type, which can be a struct for example. Since the size of the object that is returned depends on `local_type_id__k`, the verifier will enforce that `local_type_id__k` is a valid BTF ID and that it knows the size of the object that will be returned.

For other functions a different type of constant might be used.

### `__ign` annotation

A parameter with the `__ign` suffix is used to indicate that this parameter is ignored during typechecks. So any type can be passed into it, no restrictions.

### `__uninit` annotation

A parameter with the `__uninit` suffix is used to indicate that the parameter will be treated as if its uninitialized. Normally, without this annotation, the verifier will enforce that all parameters are initialized before they are used.

So its typically used in situations where a KFunc initializes a object for you.

### `__alloc` annotation

A parameter with the `__alloc` suffix is used to indicate that the parameter is a pointer to a memory region that has been allocated by a KFunc at some time. 

This is typically used on KFuncs such as [`bpf_obj_drop_impl`](../kfuncs/bpf_obj_drop_impl.md) which frees the memory allocated by [`bpf_obj_new_impl`](../kfuncs/bpf_obj_new_impl.md). Here we want to prevent a pointer to stack or map value to be passed in.

### `__opt` annotation

A parameter with the `__opt` suffix is used to indicate that the parameter associated with an `__sz` or `__szk` it is optional. This means that the parameter can be `NULL`.

```c
void *bpf_dynptr_slice(..., void *buffer__opt, u32 buffer__szk)
```

### `__refcounted_kptr` annotation

A parameter with the `__refcounted_kptr` suffix is used to indicate that values passed to this parameter must be a reference counted kernel pointer.

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### `__nullable` annotation

A parameter with the `__nullable` suffix is used to indicate that the parameter can be `NULL`. Normally a typed pointer must be non-NULL, but with this annotation the verifier will allow `NULL` to be passed in.

For example:

```c
int bpf_iter_task_new(struct bpf_iter_task *it, struct task_struct *task__nullable, unsigned int flags)
```

### `__str` annotation

A parameter with the `__str` suffix is used to indicate that the parameter is a constant string.

For example

```c
bpf_get_file_xattr(..., const char *name__str, ...)
```

Can be called

```c
bpf_get_file_xattr(..., "xattr_name", ...);
```

or 

```c
const char name[] = "xattr_name";  /* This need to be global */
int BPF_PROG(...)
{
        ...
        bpf_get_file_xattr(..., name, ...);
        ...
}
```

But not with a string that isn't known at compile time.

## KFunc flags

KFuncs can have flags associated with them. These aren't visible in the function signature, but are used to indicate certain properties of the function. Whenever a flag has significant impact on the behavior of the function, it will be documented in the KFunc page.

A functions can have multiple flags at the same time, they are not mutually exclusive in most cases.

This section will document the flags that are available for the sake of completeness.

### `KF_ACQUIRE`

The `KF_ACQUIRE` flag is used to indicate that the KFunc returns a reference to a kernel object. This means that the caller is responsible for releasing the reference when it is done with it.

Typically a `KF_ACQUIRE` KFunc will have a corresponding `KF_RELEASE` KFunc, such pairs are easy to spot.

### `KF_RELEASE`

The `KF_RELEASE` flag is used to indicate that the KFunc takes a reference to a kernel object and releases it. 

Typically a `KF_ACQUIRE` KFunc will have a corresponding `KF_RELEASE` KFunc, such pairs are easy to spot.

### `KF_RET_NULL`

The `KF_RET_NULL` flag is used to indicate that the KFunc can return `NULL`. The verifier will enforce that the return value is checked for `NULL` before it is passed to another KFunc that doesn't accept nullable values or is dereferenced.

### `KF_TRUSTED_ARGS`

The `KF_TRUSTED_ARGS` flag is used to indicate that pointers to kernel objects passed to this KFunc must be "valid". And that all pointers to BTF objects must be in their unmodified form.

Being a "valid" kernel pointer means one of the following:

* Pointers which are passed as tracepoint or struct_ops callback arguments.
* Pointers which were returned from a `KF_ACQUIRE` kfunc.

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### `KF_SLEEPABLE`

The `KF_SLEEPABLE` flag is used to indicate that the KFunc can sleep. This means that this KFunc can only be called from programs that have been loaded as sleepable ([`BPF_F_SLEEPABLE`](../syscall/BPF_PROG_LOAD.md#bpf_f_sleepable) flag set during loading).

### `KF_DESTRUCTIVE`

The `KF_DESTRUCTIVE` flag is used to indicate that the KFunc is destructive to the system. A call can for example cause the kernel to panic or reboot. Due to the risk, only programs loaded by a user with `CAP_SYS_BOOT` can call such KFuncs.

### `KF_RCU`

The `KF_RCU` flag is used to indicate that the KFunc that arguments to this KFunc must be RCU protected.

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### `KF_ITER_NEW`

The `KF_ITER_NEW` flag is used to indicate that the KFunc is used to initialize an iterator. This means that the KFunc will return a pointer to an iterator that can be used to iterate over a set of objects.

The verifier will guarantee that an iterator is destroyed by a function with the `KF_ITER_DESTROY` flag. Typically a `KF_ITER_NEW` KFunc will have a corresponding `KF_ITER_DESTROY` KFunc, such pairs are easy to spot.

### `KF_ITER_NEXT`

The `KF_ITER_NEXT` flag is used to indicate that the KFunc is used to advance an iterator. This means that the KFunc will take a pointer to an iterator and advance it to the next object.

The verifier will enforce that a `KF_ITER_NEXT` KFunc is only called with an interator created a `KF_ITER_NEW` KFunc. Typically a `KF_ITER_NEW` KFunc will have a corresponding `KF_ITER_NEXT` KFunc, such pairs are easy to spot.

### `KF_ITER_DESTROY`

The `KF_ITER_DESTROY` flag is used to indicate that the KFunc is used to destroy an iterator. This means that the KFunc will take a pointer to an iterator and destroy it.

The verifier will guarantee that an iterator is destroyed by a function with the `KF_ITER_DESTROY` flag. Typically a `KF_ITER_NEW` KFunc will have a corresponding `KF_ITER_DESTROY` KFunc, such pairs are easy to spot.

### `KF_RCU_PROTECTED`

The `KF_RCU_PROTECTED` flag is used to indicate that the KFunc can only be used within a RCU critical section. This means that sleepable programs must explicitly use [`bpf_rcu_read_lock`](../kfuncs/bpf_rcu_read_lock.md) and [`bpf_rcu_read_unlock`](../kfuncs/bpf_rcu_read_unlock.md) to protect the calls to such KFuncs. Programs that run in the context of a RCU critical section can call these KFuncs without any additional protection.
