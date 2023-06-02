# Helper function `bpf_sock_ops_cb_flags_set`

<!-- [FEATURE_TAG](bpf_sock_ops_cb_flags_set) -->
[:octicons-tag-24: v4.16](https://github.com/torvalds/linux/commit/b13d880721729384757f235166068c315326f4a1)
<!-- [/FEATURE_TAG] -->

## Definition

<!-- [HELPER_FUNC_DEF] -->
Attempt to set the value of the **bpf_sock_ops_cb_flags** field for the full TCP socket associated to _bpf_sock_ops_ to _argval_.

The primary use of this field is to determine if there should be calls to eBPF programs of type **BPF_PROG_TYPE_SOCK_OPS** at various points in the TCP code. A program of the same type can change its value, per connection and as necessary, when the connection is established. This field is directly accessible for reading, but this helper must be used for updates in order to return an error if an eBPF program tries to set a callback that is not supported in the current kernel.

_argval_ is a flag array which can combine these flags:

* **BPF_SOCK_OPS_RTO_CB_FLAG** (retransmission time out)
* **BPF_SOCK_OPS_RETRANS_CB_FLAG** (retransmission)
* **BPF_SOCK_OPS_STATE_CB_FLAG** (TCP state change)
* **BPF_SOCK_OPS_RTT_CB_FLAG** (every RTT)


Therefore, this function can be used to clear a callback flag by setting the appropriate bit to zero. e.g. to disable the RTO callback:

**bpf_sock_ops_cb_flags_set(bpf_sock,**

&nbsp;&nbsp;&nbsp;&nbsp;**bpf_sock->bpf_sock_ops_cb_flags & ~BPF_SOCK_OPS_RTO_CB_FLAG)**

Here are some examples of where one could call such eBPF program:

* When RTO fires.
* When a packet is retransmitted.
* When the connection terminates.
* When a packet is sent.
* When a packet is received.


### Returns

Code **-EINVAL** if the socket is not a full TCP socket; otherwise, a positive number containing the bits that could not be set is returned (which comes down to 0 if all bits were set as required).

`#!c static long (*bpf_sock_ops_cb_flags_set)(struct bpf_sock_ops *bpf_sock, int argval) = (void *) 59;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_SOCK_OPS](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
