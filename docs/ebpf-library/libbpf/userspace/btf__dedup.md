---
title: "Libbpf userspace function 'btf__dedup'"
description: "This page documents the 'btf__dedup' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf__dedup`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/releases/tag/v0.6.0)
<!-- [/LIBBPF_TAG] -->

Deduplicate BTF types and strings.

## Definition

`#!c int btf__dedup(struct btf *btf, const struct btf_dedup_opts *opts);`

**Parameters**

- `btf`: pointer to a `struct btf` object
- `opts`: pointer to a `struct btf_dedup_opts` object

**Return**

`0`, on success; `<0`, on error.

### `struct btf_dedup_opts`

```c
struct btf_dedup_opts {
	size_t sz;
	struct btf_ext *btf_ext;
	bool force_collisions;
	size_t :0;
};
```

#### `btf_ext`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/releases/tag/v0.6.0)

optional `.BTF.ext` info to dedup along the main BTF info

#### `force_collisions`

[:octicons-tag-24: 0.6.0](https://github.com/libbpf/libbpf/releases/tag/v0.6.0)

force hash collisions (used for testing)

## Usage

BTF deduplication algorithm takes as an input `struct btf` representing `.BTF` ELF section with all BTF type descriptors and string data. It overwrites that memory in-place with deduplicated types and strings without any loss of information. If optional `struct btf_ext` representing '.BTF.ext' ELF section is provided, all the strings referenced from .BTF.ext section are honored and updated to point to the right offsets after deduplication.

If function returns with error, type/string data might be garbled and should be discarded.

More verbose and detailed description of both problem `btf_dedup` is solving, as well as solution could be found at: [https://facebookmicrosites.github.io/bpf/blog/2018/11/14/btf-enhancement.html](https://facebookmicrosites.github.io/bpf/blog/2018/11/14/btf-enhancement.html)

##  Problem description and justification

BTF type information is typically emitted either as a result of conversion from DWARF to BTF or directly by compiler. In both cases, each compilation unit contains information about a subset of all the types that are used in an application. These subsets are frequently overlapping and contain a lot of duplicated information when later concatenated together into a single binary. This algorithm ensures that each unique type is represented by single BTF type descriptor, greatly reducing resulting size of BTF data.

Compilation unit isolation and subsequent duplication of data is not the only problem. The same type hierarchy (e.g., struct and all the type that struct references) in different compilation units can be represented in BTF to various degrees of completeness (or, rather, incompleteness) due to struct/union forward declarations.

Let's take a look at an example, that we'll use to better understand the problem (and solution). Suppose we have two compilation units, each using same `struct S`, but each of them having incomplete type information about struct's fields:

```c
// CU #1:
struct S;
struct A {
	int a;
	struct A* self;
	struct S* parent;
};
struct B;
struct S {
	struct A* a_ptr;
	struct B* b_ptr;
};

// CU #2:
struct S;
struct A;
struct B {
	int b;
	struct B* self;
	struct S* parent;
};
struct S {
	struct A* a_ptr;
	struct B* b_ptr;
};
```

In case of CU #1, BTF data will know only that `struct B` exist (but no more), but will know the complete type information about `struct A`. While for CU #2, it will know full type information about `struct B`, but will only know about forward declaration of `struct A` (in BTF terms, it will have `BTF_KIND_FWD` type descriptor with name `B`).

This compilation unit isolation means that it's possible that there is no single CU with complete type information describing structs `S`, `A`, and `B`. Also, we might get tons of duplicated and redundant type information.

Additional complication we need to keep in mind comes from the fact that types, in general, can form graphs containing cycles, not just directed acyclic graph.

While algorithm does deduplication, it also merges and resolves type information (unless disabled thought `struct btf_opts`), whenever possible. E.g., in the example above with two compilation units having partial type information for structs `A` and `B`, the output of algorithm will emit a single copy of each BTF type that describes structs `A`, `B`, and `S` (as well as type information for `int` and pointers), as if they were defined in a single compilation unit as:

```c
struct A {
	int a;
	struct A* self;
	struct S* parent;
};
struct B {
	int b;
	struct B* self;
	struct S* parent;
};
struct S {
	struct A* a_ptr;
	struct B* b_ptr;
};
```

## Algorithm summary

Algorithm completes its work in 7 separate passes:

1. Strings deduplication.
2. Primitive types deduplication (int, enum, fwd).
3. Struct/union types deduplication.
4. Resolve unambiguous forward declarations.
5. Reference types deduplication <nospell>(pointers, typedefs, arrays, funcs, func protos, and const/volatile/restrict modifiers)</nospell>.
6. Types compaction.
7. Types remapping.

Algorithm determines canonical type descriptor, which is a single representative type for each truly unique type. This canonical type is the one that will go into final deduplicated BTF type information. For struct/unions, it is also the type that algorithm will merge additional type information into (while resolving forward declarations), as it discovers it from data in other compilation units. Each input BTF type eventually gets either mapped to itself, if that type is canonical, or to some other type, if that type is equivalent and was chosen as canonical representative. This mapping is stored in `btf_dedup->map` array. This map is also used to record STRUCT/UNION that FWD type got resolved to.

To facilitate fast discovery of canonical types, we also maintain canonical  index (`btf_dedup->dedup_table`), which maps type descriptor's signature hash  (i.e., hashed kind, name, size, fields, etc) into a list of canonical types  that match that signature. With sufficiently good choice of type signature  hashing function, we can limit number of canonical types for each unique type  signature to a very small number, allowing to find canonical type for any  duplicated type very quickly.

Struct/union deduplication is the most critical part and algorithm for deduplicating structs/unions is described in greater details in comments for `btf_dedup_is_equiv` function.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
