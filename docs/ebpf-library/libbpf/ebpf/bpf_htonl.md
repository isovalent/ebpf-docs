---
title: "Libbpf eBPF macro 'bpf_htonl'"
description: "This page documents the 'bpf_htonl' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `bpf_htonl`

[:octicons-tag-24: v0.0.6](https://github.com/libbpf/libbpf/releases/tag/v0.0.6)

The `bpf_htonl` macro is used to convert a 32-bit number from host byte order to network byte order.

## Definition

```c
#define ___bpf_mvb(x, b, n, m) ((__u##b)(x) << (b-(n+1)*8) >> (b-8) << (m*8))

#define ___bpf_swab32(x) ((__u32)(			\
			  ___bpf_mvb(x, 32, 0, 3) |	\
			  ___bpf_mvb(x, 32, 1, 2) |	\
			  ___bpf_mvb(x, 32, 2, 1) |	\
			  ___bpf_mvb(x, 32, 3, 0)))

#if __BYTE_ORDER__ == __ORDER_LITTLE_ENDIAN__
# define __bpf_htonl(x)			__builtin_bswap32(x)
# define __bpf_constant_htonl(x)	___bpf_swab32(x)
#elif __BYTE_ORDER__ == __ORDER_BIG_ENDIAN__
# define __bpf_htonl(x)			(x)
# define __bpf_constant_htonl(x)	(x)
#else
# error "Fix your compiler's __BYTE_ORDER__?!"
#endif

#define bpf_htonl(x)				\
	(__builtin_constant_p(x) ?		\
	 __bpf_constant_htonl(x) : __bpf_htonl(x))
```

## Usage

This macro implements the analog of the `htonl` function from the standard C library. `htonl` being short for "host to network long", converts a 32-bit number (a `long`) from host byte order to network byte order.

The implementation checks the endianness of the host system and if the number is a compile time constant or not. If the endianness of the system we are compiling on is already in network order, the macro simply returns the number as is. Otherwise if conversion is needed, and the number is a compile time constant, the conversion is done at compile time. If the number is not a compile time constant, a compiler builtin is used to emit byte swap instructions.

### Example

Only allow a socket to bind to `192.168.1.254` and port `4040`.

```c hl_lines="16"
#define SERV4_IP		0xc0a801feU /* 192.168.1.254 */
#define SERV4_PORT		4040

SEC("cgroup/bind4")
int bind_v4_prog(struct bpf_sock_addr *ctx)
{
	struct bpf_sock *sk;

	sk = ctx->sk;
	if (!sk)
		return 0;

	if (sk->family != AF_INET)
		return 0;

	if (ctx->user_ip4 != bpf_htonl(SERV4_IP) ||
	    ctx->user_port != bpf_htons(SERV4_PORT))
		return 0;
    
    return 1;
}
```
