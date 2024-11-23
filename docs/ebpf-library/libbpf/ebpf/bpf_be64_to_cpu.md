---
title: "Libbpf eBPF macro 'bpf_be64_to_cpu'"
description: "This page documents the 'bpf_be64_to_cpu' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `bpf_be64_to_cpu`

[:octicons-tag-24: v0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)

The `bpf_be64_to_cpu` macro is used to convert a 64-bit number from network byte order to host byte order.

## Definition

```c
#define ___bpf_mvb(x, b, n, m) ((__u##b)(x) << (b-(n+1)*8) >> (b-8) << (m*8))

#define ___bpf_swab64(x) ((__u64)(			\
			  ___bpf_mvb(x, 64, 0, 7) |	\
			  ___bpf_mvb(x, 64, 1, 6) |	\
			  ___bpf_mvb(x, 64, 2, 5) |	\
			  ___bpf_mvb(x, 64, 3, 4) |	\
			  ___bpf_mvb(x, 64, 4, 3) |	\
			  ___bpf_mvb(x, 64, 5, 2) |	\
			  ___bpf_mvb(x, 64, 6, 1) |	\
			  ___bpf_mvb(x, 64, 7, 0)))

#if __BYTE_ORDER__ == __ORDER_LITTLE_ENDIAN__
# define __bpf_be64_to_cpu(x)		__builtin_bswap64(x)
# define __bpf_constant_be64_to_cpu(x)	___bpf_swab64(x)
#elif __BYTE_ORDER__ == __ORDER_BIG_ENDIAN__
# define __bpf_be64_to_cpu(x)		(x)
# define __bpf_constant_be64_to_cpu(x)  (x)
#else
# error "Fix your compiler's __BYTE_ORDER__?!"
#endif

#define bpf_be64_to_cpu(x)			\
	(__builtin_constant_p(x) ?		\
	 __bpf_constant_be64_to_cpu(x) : __bpf_be64_to_cpu(x))
```

## Usage

Converts a 64-bit number (a `long long`) from host byte order to network byte order.

The implementation checks the endianness of the host system and if the number is a compile time constant or not. If the endianness of the system we are compiling on is already in network order, the macro simply returns the number as is. Otherwise if conversion is needed, and the number is a compile time constant, the conversion is done at compile time. If the number is not a compile time constant, a compiler builtin is used to emit byte swap instructions.

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
