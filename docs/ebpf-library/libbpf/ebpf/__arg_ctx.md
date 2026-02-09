---
title: "Libbpf eBPF macro '__arg_ctx'"
description: "This page documents the '__arg_ctx' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `__arg_ctx`

[:octicons-tag-24: v1.4.0](https://github.com/libbpf/libbpf/releases/tag/v1.4.0)

The `__arg_ctx` macros is used to tag a function argument to tell the verifier that it is a program context.

## Definition

`#!c #define __arg_ctx __attribute__((btf_decl_tag("arg:ctx")))`

## Usage

This macro can be used to tag a function argument of a [global function](../../../linux/concepts/functions.md#function-by-function-verification) if you want to write that function in such a way that it can be reused between different program types. Global functions are verified function-by-function, so the function can be verifier before any of its callers. The verifier therefore has to use type info to determine possible values. The verifier will already implicitly associate types such as `struct __sk_buff*` and `struct xdp_md*` with the program context and assert only the actual, valid context is passed. However, a function that can work with multiple program contexts needs to use `void *` to be able to compile, which means the verifier is missing type info. When this becomes an issue you can add the `__arg_ctx` macro to the function argument to tell the verifier that the argument is a program context. The verifier will treat the argument as a program context for all intents and purposes and it will enforce a valid context is passed on the call site.

### Example

```c hl_lines="7"
struct {
    __uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
    __uint(key_size, sizeof(__u32));
    __uint(value_size, sizeof(int));
} events SEC(".maps");

int send_msg(void *ctx __arg_ctx)
{
    char msg[] = "Some common message";
	return bpf_perf_event_output(ctx, &events, 0, msg, sizeof(msg));
}

SEC("kprobe/eth_type_trans")
int kprobe__sys_open(struct pt_regs *ctx)
{
    send_msg(ctx);
    return 0;
}

SEC("fentry/eth_type_trans")
int BPF_PROG(fentry_eth_type_trans, struct sk_buff *skb,
	     struct net_device *dev, unsigned short protocol)
{
    // Note: The BPF_PROG does some magic to give us typed arguments, but `ctx` is still preserved as the
    // context passed into the program.
	send_msg(ctx);
    return 0;
}
```
