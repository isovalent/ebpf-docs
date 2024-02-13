# Pinning

[:octicons-tag-24: v4.4](https://github.com/torvalds/linux/commit/b2197755b2633e164a439682fb05a9b5ea48f706)

Pinning is a technique whereby we can make a pseudo-file in the BPF file system hold a reference to a BPF object. BPF objects are reference counted, meaning that if all references to a BPF object are gone, the kernel will unload/kill/free that BPF object.

A pin can be created by any process that has a file descriptor to a BPF object, but passing it into the [`BPF_OBJ_PIN`](../syscall/BPF_OBJ_PIN.md) syscall command alongside a valid path inside the BPF file system which is typically mounted at `/sys/fs/bpf`.

If your linux distribution does not automatically mount the BPF file system you can do so manually by executing `#!bash mount -t bpf bpffs /sys/fs/bpf` as root or making it part of a setup/initialization script.

A process can get a file descriptor to a BPF object by calling the [`BPF_OBJ_GET`](../syscall/BPF_OBJ_GET.md) syscall command, passing it a valid path to a pin.

Pins are usually used as an easy method of sharing or transferring a BPF object between processes or applications. Command line tools which have short running processes before existing can for example use them to perform actions on object over multiple invocation. Long running daemons can use pins to ensure resources do not go away while restarting. And tools like iproute2/tc can load a program on behalf of a user and then another program can modify the maps afterwards.

Pins can be removed by using the `rm` cli tool or `unlink` syscall. Pins are ephemeral and do not persist over restarts of the system.

Most loader libraries will offer a API for pinning and opening resources from pins. This is usually an action that needs to be explicitly taken. For BTF style maps, however, there is a property called `pinning` which is set to the macro value `LIBBPF_PIN_BY_NAME`/`1` then most loader libraries will attempt to pin the map by default (sometimes given a path to a directory). If a pin already exists, the library will open the pin instead of creating a new map. If no pin exists, the library creates a new map and pins it, using the name of the map as filename.

Example of a map definition with `pinning`:
:``c
struct {
	__uint(type, BPF_MAP_TYPE_ARRAY);
	__uint(max_entries, 1);
	__type(key, __u32);
	__type(value, __u64);
	__uint(pinning, LIBBPF_PIN_BY_NAME);
} pinmap SEC(".maps");
```
