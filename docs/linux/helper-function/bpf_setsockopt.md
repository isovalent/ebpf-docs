# Helper function `bpf_setsockopt`

<!-- [FEATURE_TAG](bpf_setsockopt) -->
[:octicons-tag-24: v4.13](https://github.com/torvalds/linux/commit/8c4b4c7e9ff0447995750d9329949fa082520269)
<!-- [/FEATURE_TAG] -->

## Definition

<!-- [HELPER_FUNC_DEF] -->
Emulate a call to **setsockopt()** on the socket associated to _bpf_socket_, which must be a full socket. The _level_ at which the option resides and the name _optname_ of the option must be specified, see **setsockopt(2)** for more information. The option value of length _optlen_ is pointed by _optval_.

_bpf_socket_ should be one of the following:

* **struct bpf_sock_ops** for **BPF_PROG_TYPE_SOCK_OPS**.
* **struct bpf_sock_addr** for **BPF_CGROUP_INET4_CONNECT**
  and **BPF_CGROUP_INET6_CONNECT**.

This helper actually implements a subset of **setsockopt()**. It supports the following _level_s:

* **SOL_SOCKET**, which supports the following _optname_s:
  **SO_RCVBUF**, **SO_SNDBUF**, **SO_MAX_PACING_RATE**,   **SO_PRIORITY**, **SO_RCVLOWAT**, **SO_MARK**,   **SO_BINDTODEVICE**, **SO_KEEPALIVE**, **SO_REUSEADDR**,   **SO_REUSEPORT**, **SO_BINDTOIFINDEX**, **SO_TXREHASH**. * **IPPROTO_TCP**, which supports the following _optname_s:
  **TCP_CONGESTION**, **TCP_BPF_IW**,   **TCP_BPF_SNDCWND_CLAMP**, **TCP_SAVE_SYN**,   **TCP_KEEPIDLE**, **TCP_KEEPINTVL**, **TCP_KEEPCNT**,   **TCP_SYNCNT**, **TCP_USER_TIMEOUT**, **TCP_NOTSENT_LOWAT**,   **TCP_NODELAY**, **TCP_MAXSEG**, **TCP_WINDOW_CLAMP**,   **TCP_THIN_LINEAR_TIMEOUTS**, **TCP_BPF_DELACK_MAX**,   **TCP_BPF_RTO_MIN**. * **IPPROTO_IP**, which supports _optname_ **IP_TOS**.
* **IPPROTO_IPV6**, which supports the following _optname_s:
  **IPV6_TCLASS**, **IPV6_AUTOFLOWLABEL**.

### Returns

0 on success, or a negative error in case of failure.

`#!c static long (*bpf_setsockopt)(void *bpf_socket, int level, int optname, void *optval, int optlen) = (void *) 49;`
<!-- [/HELPER_FUNC_DEF] -->

## Usage

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome

### Program types

This helper call can be used in the following program types:

<!-- DO NOT EDIT MANUALLY -->
<!-- [HELPER_FUNC_PROG_REF] -->
 * [BPF_PROG_TYPE_CGROUP_SOCKOPT](../program-type/BPF_PROG_TYPE_CGROUP_SOCKOPT.md) [:octicons-tag-24: v5.15](2c531639deb5e3ddfd6e8123b82052b2d9fbc6e5)
 * [BPF_PROG_TYPE_CGROUP_SOCK_ADDR](../program-type/BPF_PROG_TYPE_CGROUP_SOCK_ADDR.md) [:octicons-tag-24: v5.8](beecf11bc2188067824591612151c4dc6ec383c7)
 * [BPF_PROG_TYPE_LSM](../program-type/BPF_PROG_TYPE_LSM.md) [:octicons-tag-24: v6.0](9113d7e48e9128522b9f5a54dfd30dff10509a92)
 * [BPF_PROG_TYPE_SOCK_OPS](../program-type/BPF_PROG_TYPE_SOCK_OPS.md)
<!-- [/HELPER_FUNC_PROG_REF] -->

### Example

!!! example "Docs could be improved"
    This part of the docs is incomplete, contributions are very welcome
