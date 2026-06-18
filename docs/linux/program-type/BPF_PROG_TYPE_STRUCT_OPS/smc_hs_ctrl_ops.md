---
title: "Struct ops 'smc_hs_ctrl_ops'"
description: "This page documents the 'smc_hs_ctrl_ops' struct ops, its semantics, capabilities, and limitations."
---
# Struct ops `smc_hs_ctrl_ops`

[:octicons-tag-24: v6.19](https://github.com/torvalds/linux/commit/15f295f55656658e65bdbc9b901d6b2e49d68d72)

SMC handshake control ops is a type of struct_ops which allows BPF programs to control whether SMC (Shared Memory Communications) is negotiated for individual TCP connections.

## Usage

[SMC](https://datatracker.ietf.org/doc/html/rfc7609) is a high-performance socket protocol that transparently replaces TCP. By default, SMC negotiation is initiated for all eligible connections via TCP option exchange during the three-way handshake. The `smc_hs_ctrl_ops` struct ops provides hooks that are called before SMC options are set in SYN and SYN-ACK packets, allowing a BPF program to selectively enable or disable SMC negotiation per connection based on runtime information such as local/remote IP address or port numbers.

A registered `smc_hs_ctrl` instance is associated with a network namespace. It can optionally be inherited by child network namespaces when the `SMC_HS_CTRL_FLAG_INHERITABLE` flag is set.

The active controller for a network namespace can be configured via the sysctl `net.smc.hs_ctrl`, which accepts the `name` of a registered `smc_hs_ctrl` instance.

## Fields and ops

```c
struct smc_hs_ctrl {
	char name[SMC_HS_CTRL_NAME_MAX];
	int  flags;
	int (*syn_option)(struct tcp_sock *tp);
	int (*synack_option)(const struct tcp_sock *tp,
	                     struct inet_request_sock *ireq);
};
```

### `name`

[:octicons-tag-24: v6.19](https://github.com/torvalds/linux/commit/15f295f55656658e65bdbc9b901d6b2e49d68d72)

`#!c char name[SMC_HS_CTRL_NAME_MAX]`

The unique name of this `smc_hs_ctrl` instance. `SMC_HS_CTRL_NAME_MAX` is 16 bytes. This name is used to refer to the instance via the `net.smc.hs_ctrl` sysctl.

### `flags`

[:octicons-tag-24: v6.19](https://github.com/torvalds/linux/commit/15f295f55656658e65bdbc9b901d6b2e49d68d72)

`#!c int flags`

Flags controlling the behavior of this `smc_hs_ctrl` instance.

Bitmask values:

- `SMC_HS_CTRL_FLAG_INHERITABLE` (`0x1`) - If set, child network namespaces created from a namespace using this ctrl will inherit it.

### `syn_option`

[:octicons-tag-24: v6.19](https://github.com/torvalds/linux/commit/15f295f55656658e65bdbc9b901d6b2e49d68d72)

`#!c int (*syn_option)(struct tcp_sock *tp)`

This function/program is invoked before SMC options are written into a SYN packet. It controls whether the connecting side advertises SMC capability in its SYN.

**Parameters**

`tp`: The [`struct tcp_sock`](https://elixir.bootlin.com/linux/v6.19-rc1/source/include/linux/tcp.h) for the outgoing SYN.

**Returns**

`0` to disable SMC negotiation for this connection (the SMC option will not be added to the SYN packet). Any other value to allow SMC negotiation to proceed as normal.

### `synack_option`

[:octicons-tag-24: v6.19](https://github.com/torvalds/linux/commit/15f295f55656658e65bdbc9b901d6b2e49d68d72)

`#!c int (*synack_option)(const struct tcp_sock *tp, struct inet_request_sock *ireq)`

This function/program is invoked before SMC options are written into a SYN-ACK packet. It controls whether the server side responds with SMC capability in its SYN-ACK.

**Parameters**

`tp`: The [`struct tcp_sock`](https://elixir.bootlin.com/linux/v6.19-rc1/source/include/linux/tcp.h) associated with the listener.

`ireq`: The [`struct inet_request_sock`](https://elixir.bootlin.com/linux/v6.19-rc1/source/include/net/inet_sock.h) representing the connection request, containing information about the remote peer.

**Returns**

`0` to disable SMC negotiation for this connection (the SMC option will not be added to the SYN-ACK packet). Any other value to allow SMC negotiation to proceed as normal.

## Example

```c
// SPDX-License-Identifier: GPL-2.0
#include <vmlinux.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>

/* Disable SMC for connections from/to port 80 */
SEC("struct_ops/syn_option")
int BPF_PROG(my_syn_option, struct tcp_sock *tp)
{
    struct sock *sk = (struct sock *)tp;

    /* Return 0 to disable SMC, non-zero to enable */
    if (sk->sk_num == 80 || sk->sk_dport == bpf_htons(80))
        return 0;
    return 1;
}

SEC("struct_ops/synack_option")
int BPF_PROG(my_synack_option, const struct tcp_sock *tp,
             struct inet_request_sock *ireq)
{
    /* Return 0 to disable SMC, non-zero to enable */
    if (ireq->ir_rmt_port == bpf_htons(80))
        return 0;
    return 1;
}

SEC(".struct_ops")
struct smc_hs_ctrl my_smc_hs_ctrl = {
    .name           = "my_ctrl",
    .flags          = 0,
    .syn_option     = (void *)my_syn_option,
    .synack_option  = (void *)my_synack_option,
};

char _license[] SEC("license") = "GPL";
```
