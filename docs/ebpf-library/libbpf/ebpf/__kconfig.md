---
title: "Libbpf eBPF macro '__kconfig'"
description: "This page documents the '__kconfig' libbpf eBPF macro, including its definition, usage, and examples."
---
# Libbpf eBPF macro `__kconfig`

[:octicons-tag-24: v0.0.7](https://github.com/libbpf/libbpf/releases/tag/v0.0.7)

The `__kconfig` macros is used to instruct the loader to provide the value of a kernel configuration of the system the program is being loaded on to the program.

## Definition

`#!c #define __kconfig __attribute__((section(".kconfig")))`

## Usage

This macro places a global variable in the `.kconfig` section of the eBPF object file. This will signal the loader (library) to initialize the value of the variable to the value of the kernel configuration option with the same name as the variable. Since kernel configuration values can be `y`, `m`, or `n`, the variable will be initialized to enum values instead.

```c
enum libbpf_tristate {
	TRI_NO = 0,
	TRI_YES = 1,
	TRI_MODULE = 2,
};
```

This enum type is also provided by the `bpf_helpers.h` file. 

This capability ties into CO-RE (Compile Once - Run Everywhere). Since the value is provided at load time it allows the user to write multiple variations / code paths that are enabled or disabled based on the value of, for example, in this case a kernel configuration option. Therefore allowing the same pre-compiled eBPF program to be loaded on different systems with different kernel configurations.

If a value for a kernel configuration option is not found, the loader (library) will error out, unless the [`__weak`](__weak.md) attribute is also used.

### Example

```c hl_lines="1"
extern unsigned CONFIG_HZ __kconfig;

#define USER_HZ		100
#define NSEC_PER_SEC	1000000000ULL
static clock_t jiffies_to_clock_t(unsigned long x)
{
	/* The implementation here tailored to a particular
	 * setting of USER_HZ.
	 */
	u64 tick_nsec = (NSEC_PER_SEC + CONFIG_HZ/2) / CONFIG_HZ;
	u64 user_hz_nsec = NSEC_PER_SEC / USER_HZ;

	if ((tick_nsec % user_hz_nsec) == 0) {
		if (CONFIG_HZ < USER_HZ)
			return x * (USER_HZ / CONFIG_HZ);
		else
			return x / (CONFIG_HZ / USER_HZ);
	}
	return x * tick_nsec/user_hz_nsec;
}
```
