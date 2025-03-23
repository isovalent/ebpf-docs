---
title: "Libbpf userspace function 'btf_dump__emit_type_decl'"
description: "This page documents the 'btf_dump__emit_type_decl' libbpf userspace function, including its definition, usage, and examples."
---
# Libbpf userspace function `btf_dump__emit_type_decl`

<!-- [LIBBPF_TAG] -->
[:octicons-tag-24: 0.0.7](https://github.com/libbpf/libbpf/releases/tag/v0.0.7)
<!-- [/LIBBPF_TAG] -->

Emit type declaration (e.g., field type declaration in a struct or argument declaration in function prototype) in correct C syntax.

## Definition

`#!c int btf_dump__emit_type_decl(struct btf_dump *d, __u32 id, const struct btf_dump_emit_type_decl_opts *opts);`

**Parameters**

- `d`: pointer to a `struct btf_dump` object
- `id`: BTF type ID to emit declaration for
- `opts`: pointer to a `struct btf_dump_emit_type_decl_opts` object

**Return**

`0`, on success; `-errno`, on error.

### `struct btf_dump_emit_type_decl_opts`

```c
struct btf_dump_emit_type_decl_opts {
	/* size of this struct, for forward/backward compatibility */
	size_t sz;
	const char *field_name;
	int indent_level;
	bool strip_mods;
	size_t :0;
};
```

#### `field_name`

[:octicons-tag-24: 0.0.7](https://github.com/libbpf/libbpf/commit/5ec0ba653073ae308c7acf438971044389afddf3)

optional field name for type declaration, e.g.:

 - `struct my_struct <FNAME>`
 - `void (*<FNAME>)(int)`
 - `char (*<FNAME>)[123]`

#### `indent_level`

[:octicons-tag-24: 0.0.7](https://github.com/libbpf/libbpf/commit/5ec0ba653073ae308c7acf438971044389afddf3)

extra indentation level (in number of tabs) to emit for multi-line type declarations (e.g., anonymous struct); applies for lines starting from the second one (first line is assumed to have necessary indentation already

#### `strip_mods`

[:octicons-tag-24: 0.1.0](https://github.com/libbpf/libbpf/commit/de60a31eba192e42f25ef5c00d88c7018304bfe1)

strip all the `const`/`volatile`/`restrict` modifiers from the type declaration

## Usage

For most types it's trivial, but there are few quirky type declaration cases worth mentioning:

 - function prototypes (especially nesting of function prototypes);
 - arrays;
 - `const`/`volatile`/`restrict` for pointers vs other types.

For a good discussion of *PARSING* C syntax (as a human), see <nospell>Peter van der Linden's</nospell> "Expert C Programming: Deep C Secrets", Ch.3 "Unscrambling Declarations in C".
 
It won't help with BTF to C conversion much, though, as it's an opposite  problem. So we came up with this algorithm in reverse to <nospell>van der Linden's</nospell> parsing algorithm. It goes from structured BTF representation of type  declaration to a valid compilable C syntax.

For instance, consider this C type definition:
  `typedef const int * const * arr[10] arr_t;`

It will be represented in BTF with this chain of BTF types:
  `[typedef] -> [array] -> [ptr] -> [const] -> [ptr] -> [const] -> [int]`

Notice how `const` modifier always goes before type it modifies in BTF type graph, but in C syntax, const/volatile/restrict modifiers are written to the right of pointers, but to the left of other types. There are also other quirks, like function pointers, arrays of them, functions returning other functions, etc.

We handle that by pushing all the types to a stack, until we hit "terminal" type (int/enum/struct/union/fwd). Then depending on the kind of a type on top of a stack, modifiers are handled differently. Array/function pointers have also wildly different syntax and how nesting of them are done. See code for authoritative definition.

To avoid allocating new stack for each independent chain of BTF types, we share one bigger stack, with each chain working only on its own local view of a stack frame. Some care is required to "pop" stack frames after processing type declaration chain.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
