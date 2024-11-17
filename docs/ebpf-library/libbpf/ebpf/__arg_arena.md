---
title: "Libbpf eBPF macro '__arg_arena'"
description: "This page documents the '__arg_arena' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `__arg_arena`

[:octicons-tag-24: v1.4.0](https://github.com/libbpf/libbpf/releases/tag/v1.4.0)

The `__arg_arena` macros is used to tag a function argument to tell the verifier that a value lives on an arena.

## Definition

`#!c #define __arg_arena __attribute((btf_decl_tag("arg:arena")))`

## Usage

This macro can be used to tag a function argument of a [global function](../../../linux/concepts/functions.md#function-by-function-verification) to tell the verifier that its value lives on an arena. Since global functions are verifier out of order, the verifier has implicit way to track this property over global function boundaries. The verifier will enforce at the call site that an actual arena value is passed to the function.

### Example

```c hl_lines="56 69"
/* SPDX-License-Identifier: (GPL-2.0-only OR BSD-2-Clause) */
/* Copyright (c) 2024 Meta Platforms, Inc. and affiliates. */
#pragma once
#include <errno.h>
#include "bpf_arena_alloc.h"
#include "bpf_arena_list.h"

struct htab_bucket {
	struct arena_list_head head;
};
typedef struct htab_bucket __arena htab_bucket_t;

struct htab {
	htab_bucket_t *buckets;
	int n_buckets;
};
typedef struct htab __arena htab_t;

static inline htab_bucket_t *__select_bucket(htab_t *htab, __u32 hash)
{
	htab_bucket_t *b = htab->buckets;

	cast_kern(b);
	return &b[hash & (htab->n_buckets - 1)];
}

static inline arena_list_head_t *select_bucket(htab_t *htab, __u32 hash)
{
	return &__select_bucket(htab, hash)->head;
}

struct hashtab_elem {
	int hash;
	int key;
	int value;
	struct arena_list_node hash_node;
};
typedef struct hashtab_elem __arena hashtab_elem_t;

static hashtab_elem_t *lookup_elem_raw(arena_list_head_t *head, __u32 hash, int key)
{
	hashtab_elem_t *l;

	list_for_each_entry(l, head, hash_node)
		if (l->hash == hash && l->key == key)
			return l;

	return NULL;
}

static int htab_hash(int key)
{
	return key;
}

__weak int htab_lookup_elem(htab_t *htab __arg_arena, int key)
{
	hashtab_elem_t *l_old;
	arena_list_head_t *head;

	cast_kern(htab);
	head = select_bucket(htab, key);
	l_old = lookup_elem_raw(head, htab_hash(key), key);
	if (l_old)
		return l_old->value;
	return 0;
}

__weak int htab_update_elem(htab_t *htab __arg_arena, int key, int value)
{
	hashtab_elem_t *l_new = NULL, *l_old;
	arena_list_head_t *head;

	cast_kern(htab);
	head = select_bucket(htab, key);
	l_old = lookup_elem_raw(head, htab_hash(key), key);

	l_new = bpf_alloc(sizeof(*l_new));
	if (!l_new)
		return -ENOMEM;
	l_new->key = key;
	l_new->hash = htab_hash(key);
	l_new->value = value;

	list_add_head(&l_new->hash_node, head);
	if (l_old) {
		list_del(&l_old->hash_node);
		bpf_free(l_old);
	}
	return 0;
}

void htab_init(htab_t *htab)
{
	void __arena *buckets = bpf_arena_alloc_pages(&arena, NULL, 2, NUMA_NO_NODE, 0);

	cast_user(buckets);
	htab->buckets = buckets;
	htab->n_buckets = 2 * PAGE_SIZE / sizeof(struct htab_bucket);
}
```
