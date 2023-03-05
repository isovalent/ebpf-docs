# Concurrency

## Spin locks
eBPF spin locks lock a map element while an eBPF program accesses it.

In user space a flag `BPF_F_LOCK` can be used with `bpf_map_update_elem` and `bpf_map_lookup_elem_flags` helper functions. 

In eBPF kernel space a semaphore can be added to map elements. 

```c
struct concurrent_element {
	struct bpf_spin_lock semaphore;
	int count;
}

...

struct bpf_map_def SEC("maps") concurrent_map = {
	.type = BPF_MAP_TYPE_HASH,
	.key_size = sizeof(int),
	.value_size = sizeof(struct concurrent_element),
	.max_entries = 100,
};
// this is legacy, this whole map definition should be BTF style and this macro
// won't be necessary.
BPF_ANNOTATE_KV_PAIR(concurrent_map, int, struct concurrent_element);

...

int bpf_program(struct pt_regs *ctx) {
	int key = 0;
	struct concurrent_element init_value = {};
	struct concurrent_element *read_value;

	bpf_map_create_elem(&concurrent_map, &key, &init_value, BPF_NOEXIST);

	read_value = bpf_map_lookup_elem(&concurrent_map, &key);
	bpf_spin_lock(&read_value->semaphore);
	read_value->count += 100;
	bpf_spin_unlock(&read_value->semaphore);
}
```

<!-- TODO: -->
<!-- Per CPU maps -->
<!-- Atomics -->
<!-- Helper calls -->
