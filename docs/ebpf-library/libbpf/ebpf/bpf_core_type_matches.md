---
title: "Libbpf eBPF macro 'bpf_core_type_matches'"
description: "This page documents the 'bpf_core_type_matches' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `bpf_core_type_matches`

[:octicons-tag-24: v1.0.0](https://github.com/libbpf/libbpf/releases/tag/v1.0.0)

The `bpf_core_type_matches` macro is used to check that provided named type (struct/union/enum/typedef) "matches" that in a target kernel.

## Definition

```c
#define bpf_core_type_matches(type)					    \
	__builtin_preserve_type_info(*___bpf_typeof(type), BPF_TYPE_MATCHES)
```

## Usage

The `bpf_core_type_matches` macro is used to check that provided named type (struct/union/enum/typedef) "matches" that in a target kernel.

The matching relation is defined as follows:

* modifiers and typedefs are stripped (and, hence, effectively ignored)
* generally speaking types need to be of same kind (struct vs. struct, union vs. union, etc.)
    * exceptions are struct/union behind a pointer which could also match a forward declaration of a struct or union, respectively, and enum vs. enum64 (see below)

Then, depending on type:

* integers:
    * match if size and signedness match
* arrays & pointers:
    * target types are recursively matched
* structs & unions:
    * local members need to exist in target with the same name
    * for each member we recursively check match unless it is already behind a pointer, in which case we only check matching names and compatible kind
* enums:
    * local variants have to have a match in target by symbolic name (but not numeric value)
    * size has to match (but enum may match enum64 and vice versa)
* function pointers:
    * number and position of arguments in local type has to match target
    * for each argument and the return value we recursively check match

Returns:

 * 1, if the type matches in the target kernel's BTF
 * 0, if the type does not match any in the target kernel

This result is determined by the loader library such as libbpf, and set at load time. If a branch is never taken based on the result, it will not be evaluated by the verifier.

### Example

```c hl_lines="18 21"
struct rw_semaphore___old {
	struct task_struct *owner;
} __attribute__((preserve_access_index));

struct rw_semaphore___new {
	atomic_long_t owner;
} __attribute__((preserve_access_index));

static inline struct task_struct *get_lock_owner(__u64 lock, __u32 flags)
{
	struct task_struct *task;
	__u64 owner = 0;

	if (flags & LCB_F_MUTEX) {
		struct mutex *mutex = (void *)lock;
		owner = BPF_CORE_READ(mutex, owner.counter);
	} else if (flags == LCB_F_READ || flags == LCB_F_WRITE) {
        if (bpf_core_type_matches(struct rw_semaphore___old)) {
            struct rw_semaphore___old *rwsem = (void *)lock;
            owner = (unsigned long)BPF_CORE_READ(rwsem, owner);
        } else if (bpf_core_type_matches(struct rw_semaphore___new)) {
            struct rw_semaphore___new *rwsem = (void *)lock;
            owner = BPF_CORE_READ(rwsem, owner.counter);
        }
	}

	if (!owner)
		return NULL;

	task = (void *)(owner & ~7UL);
	return task;
}

```
