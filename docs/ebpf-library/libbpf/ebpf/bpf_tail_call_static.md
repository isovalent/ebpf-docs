---
title: "Libbpf eBPF function 'bpf_tail_call_static'"
description: "This page documents the 'bpf_tail_call_static' libbpf eBPF function, including its definition, usage, and examples."
---
# Libbpf eBPF function `bpf_tail_call_static`

[:octicons-tag-24: v0.2](https://github.com/libbpf/libbpf/releases/tag/v0.2)

The `bpf_tail_call_static` function is used to make an optimized tail call to another eBPF program, but only if the index into the [program array map](../../../linux/map-type/BPF_MAP_TYPE_PROG_ARRAY.md) is a constant.

## Definition

```c
static __always_inline void
bpf_tail_call_static(void *ctx, const void *map, const __u32 slot)
{
	if (!__builtin_constant_p(slot))
		__bpf_unreachable();

	asm volatile("r1 = %[ctx]\n\t"
		     "r2 = %[map]\n\t"
		     "r3 = %[slot]\n\t"
		     "call 12"
		     :: [ctx]"r"(ctx), [map]"r"(map), [slot]"i"(slot)
		     : "r0", "r1", "r2", "r3", "r4", "r5");
}
```

## Usage

This function emits BPF instructions in a very particular way. It emits a call to the [`bpf_tail_call`](../../../linux/helper-function/bpf_tail_call.md) helper. In addition it makes sure that the parameter registers R1, R2, and R3 are set in the instructions right before the call and that R3 is a constant value. A normal call to `bpf_tail_call` gives the compiler the liberty to move where the parameters are set, which can lead to the verifier not being able to detect that the index is a constant.

The reason why we want the verifier to detect the slot is constant is because the verifier can optimize the  the tail call can be a simple jump instead of a retpoline in that specific case.

Retpolines are a security feature that mitigates speculative execution attacks, but also make execution slow. So we gain a performance benefit by using this function.

### Example

```c
#define TAIL_CALL_ZERO 0
#define TAIL_CALL_ONE 1

struct {
	__uint(type, BPF_MAP_TYPE_PROG_ARRAY);
	__uint(max_entries, 2);
	__uint(key_size, sizeof(__u32));
	__array(values, int (void *));
} prog_array_map SEC(".maps") = {
	.values = {
        [TAIL_CALL_ZERO] = (void *)&tailcall_0,
		[TAIL_CALL_ONE] = (void *)&tailcall_1,
	},
};

SEC("xdp")
int tailcall_0(struct xdp_md *ctx)
{
    return XDP_PASS;
}

SEC("xdp")
int tailcall_1(struct xdp_md *ctx)
{
    return XDP_DROP;
}

SEC("xdp")
int xdp_prog(struct xdp_md *ctx)
{
    void *data_end = (void *)(long)ctx->data_end;
    void *data = (void *)(long)ctx->data;

    if (data + sizeof(struct ethhdr) > data_end)
        return bpf_tail_call_static(skb, &prog_array_map, TAIL_CALL_ONE);

    return bpf_tail_call_static(skb, &prog_array_map, TAIL_CALL_ZERO);
}
```
