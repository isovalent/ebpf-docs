# Libbpf userspace

Definitions for the libbpf userspace library are split across a few different header files.

## High level APIs

In the [`libbpf.h`](https://github.com/libbpf/libbpf/blob/master/src/libbpf.h) header file you will find the high level APIs which do a lot of work for you under the hood. These are the most commonly used APIs.

* BPF Object functions
    * [`bpf_object__open`](bpf_object__open.md)
    * [`bpf_object__open_file`](bpf_object__open_file.md)
    * [`bpf_object__open_mem`](bpf_object__open_mem.md)
    * [`bpf_object__load`](bpf_object__load.md)
    * [`bpf_object__close`](bpf_object__close.md)
    * [`bpf_object__pin_maps`](bpf_object__pin_maps.md)
    * [`bpf_object__unpin_maps`](bpf_object__unpin_maps.md)
    * [`bpf_object__pin_programs`](bpf_object__pin_programs.md)
    * [`bpf_object__unpin_programs`](bpf_object__unpin_programs.md)
    * [`bpf_object__pin`](bpf_object__pin.md)
    * [`bpf_object__unpin`](bpf_object__unpin.md)
    * [`bpf_object__name`](bpf_object__name.md)
    * [`bpf_object__kversion`](bpf_object__kversion.md)
    * [`bpf_object__set_kversion`](bpf_object__set_kversion.md)
    * [`bpf_object__token_fd`](bpf_object__token_fd.md)
    * [`bpf_object__btf`](bpf_object__btf.md)
    * [`bpf_object__btf_fd`](bpf_object__btf_fd.md)
    * [`bpf_object__find_program_by_name`](bpf_object__find_program_by_name.md)
    * BPF Skeleton functions
        * [`bpf_object__open_skeleton`](bpf_object__open_skeleton.md)
        * [`bpf_object__load_skeleton`](bpf_object__load_skeleton.md)
        * [`bpf_object__attach_skeleton`](bpf_object__attach_skeleton.md)
        * [`bpf_object__detach_skeleton`](bpf_object__detach_skeleton.md)
        * [`bpf_object__destroy_skeleton`](bpf_object__destroy_skeleton.md)
        * [`bpf_object__open_subskeleton`](bpf_object__open_subskeleton.md)
        * [`bpf_object__destroy_subskeleton`](bpf_object__destroy_subskeleton.md)
        * [`bpf_object__gen_loader`](bpf_object__gen_loader.md)
    * [`bpf_object__next_program`](bpf_object__next_program.md)
    * [`bpf_object__prev_program`](bpf_object__prev_program.md)
    * [`bpf_object__find_map_by_name`](bpf_object__find_map_by_name.md)
    * [`bpf_object__find_map_fd_by_name`](bpf_object__find_map_fd_by_name.md)
    * [`bpf_object__next_map`](bpf_object__next_map.md)
    * [`bpf_object__prev_map`](bpf_object__prev_map.md)
* BPF Program functions
    * [`bpf_program__set_ifindex`](bpf_program__set_ifindex.md)
    * [`bpf_program__name`](bpf_program__name.md)
    * [`bpf_program__section_name`](bpf_program__section_name.md)
    * [`bpf_program__autoload`](bpf_program__autoload.md)
    * [`bpf_program__set_autoload`](bpf_program__set_autoload.md)
    * [`bpf_program__autoattach`](bpf_program__autoattach.md)
    * [`bpf_program__set_autoattach`](bpf_program__set_autoattach.md)
    * [`bpf_program__insns`](bpf_program__insns.md)
    * [`bpf_program__set_insns`](bpf_program__set_insns.md)
    * [`bpf_program__insn_cnt`](bpf_program__insn_cnt.md)
    * [`bpf_program__fd`](bpf_program__fd.md)
    * [`bpf_program__pin`](bpf_program__pin.md)
    * [`bpf_program__unpin`](bpf_program__unpin.md)
    * [`bpf_program__unload`](bpf_program__unload.md)
    * Program attach functions
        * [`bpf_program__attach`](bpf_program__attach.md)
        * [`bpf_program__attach_perf_event`](bpf_program__attach_perf_event.md)
        * [`bpf_program__attach_perf_event_opts`](bpf_program__attach_perf_event_opts.md)
        * [`bpf_program__attach_kprobe`](bpf_program__attach_kprobe.md)
        * [`bpf_program__attach_kprobe_opts`](bpf_program__attach_kprobe_opts.md)
        * [`bpf_program__attach_kprobe_multi_opts`](bpf_program__attach_kprobe_multi_opts.md)
        * [`bpf_program__attach_uprobe_multi`](bpf_program__attach_uprobe_multi.md)
        * [`bpf_program__attach_ksyscall`](bpf_program__attach_ksyscall.md)
        * [`bpf_program__attach_uprobe`](bpf_program__attach_uprobe.md)
        * [`bpf_program__attach_uprobe_opts`](bpf_program__attach_uprobe_opts.md)
        * [`bpf_program__attach_usdt`](bpf_program__attach_usdt.md)
        * [`bpf_program__attach_tracepoint`](bpf_program__attach_tracepoint.md)
        * [`bpf_program__attach_tracepoint_opts`](bpf_program__attach_tracepoint_opts.md)
        * [`bpf_program__attach_raw_tracepoint`](bpf_program__attach_raw_tracepoint.md)
        * [`bpf_program__attach_raw_tracepoint_opts`](bpf_program__attach_raw_tracepoint_opts.md)
        * [`bpf_program__attach_trace`](bpf_program__attach_trace.md)
        * [`bpf_program__attach_trace_opts`](bpf_program__attach_trace_opts.md)
        * [`bpf_program__attach_lsm`](bpf_program__attach_lsm.md)
        * [`bpf_program__attach_cgroup`](bpf_program__attach_cgroup.md)
        * [`bpf_program__attach_netns`](bpf_program__attach_netns.md)
        * [`bpf_program__attach_sockmap`](bpf_program__attach_sockmap.md)
        * [`bpf_program__attach_xdp`](bpf_program__attach_xdp.md)
        * [`bpf_program__attach_freplace`](bpf_program__attach_freplace.md)
        * [`bpf_program__attach_netfilter`](bpf_program__attach_netfilter.md)
        * [`bpf_program__attach_tcx`](bpf_program__attach_tcx.md)
        * [`bpf_program__attach_netkit`](bpf_program__attach_netkit.md)
        * [`bpf_program__attach_iter`](bpf_program__attach_iter.md)
    * [`bpf_program__type`](bpf_program__type.md)
    * [`bpf_program__set_type`](bpf_program__set_type.md)
    * [`bpf_program__set_expected_attach_type`](bpf_program__set_expected_attach_type.md)
    * [`bpf_program__flags`](bpf_program__flags.md)
    * [`bpf_program__set_flags`](bpf_program__set_flags.md)
    * [`bpf_program__log_level`](bpf_program__log_level.md)
    * [`bpf_program__set_log_level`](bpf_program__set_log_level.md)
    * [`bpf_program__log_buf`](bpf_program__log_buf.md)
    * [`bpf_program__set_log_buf`](bpf_program__set_log_buf.md)
    * [`bpf_program__set_attach_target`](bpf_program__set_attach_target.md)
    * [`bpf_program__expected_attach_type`](bpf_program__expected_attach_type.md)
* Link functions
    * `bpf_link__open`
    * `bpf_link__fd`
    * `bpf_link__pin_path`
    * `bpf_link__pin`
    * `bpf_link__unpin`
    * `bpf_link__update_program`
    * `bpf_link__disconnect`
    * `bpf_link__detach`
    * `bpf_link__destroy`
    * `bpf_link__update_map`
* Map functions
    * `bpf_map__attach_struct_ops`
    * `bpf_map__set_autocreate`
    * `bpf_map__autocreate`
    * `bpf_map__set_autoattach`
    * `bpf_map__autoattach`
    * `bpf_map__fd`
    * `bpf_map__reuse_fd`
    * `bpf_map__name`
    * `bpf_map__type`
    * `bpf_map__set_type`
    * `bpf_map__max_entries`
    * `bpf_map__set_max_entries`
    * `bpf_map__map_flags`
    * `bpf_map__set_map_flags`
    * `bpf_map__numa_node`
    * `bpf_map__set_numa_node`
    * `bpf_map__key_size`
    * `bpf_map__set_key_size`
    * `bpf_map__value_size`
    * `bpf_map__set_value_size`
    * `bpf_map__btf_key_type_id`
    * `bpf_map__btf_value_type_id`
    * `bpf_map__ifindex`
    * `bpf_map__set_ifindex`
    * `bpf_map__map_extra`
    * `bpf_map__set_map_extra`
    * `bpf_map__set_initial_value`
    * `bpf_map__initial_value`
    * `bpf_map__is_internal`
    * `bpf_map__set_pin_path`
    * `bpf_map__pin_path`
    * `bpf_map__is_pinned`
    * `bpf_map__pin`
    * `bpf_map__unpin`
    * `bpf_map__set_inner_map_fd`
    * `bpf_map__inner_map`
    * `bpf_map__lookup_elem`
    * `bpf_map__update_elem`
    * `bpf_map__delete_elem`
    * `bpf_map__lookup_and_delete_elem`
    * `bpf_map__get_next_key`
* XDP functions
    * `bpf_xdp_attach`
    * `bpf_xdp_detach`
    * `bpf_xdp_query`
    * `bpf_xdp_query_id`
* TC functions
    * `bpf_tc_hook_create`
    * `bpf_tc_hook_destroy`
    * `bpf_tc_attach`
    * `bpf_tc_detach`
    * `bpf_tc_query`
* Ring buffer functions
    * `ring_buffer__new`
    * `ring_buffer__free`
    * `ring_buffer__add`
    * `ring_buffer__poll`
    * `ring_buffer__consume`
    * `ring_buffer__consume_n`
    * `ring_buffer__epoll_fd`
    * `ring_buffer__ring`
    * Ring functions
        * `ring__consumer_pos`
        * `ring__producer_pos`
        * `ring__avail_data_size`
        * `ring__size`
        * `ring__map_fd`
        * `ring__consume`
        * `ring__consume_n`
* User ring buffer
    * `user_ring_buffer__new`
    * `user_ring_buffer__reserve`
    * `user_ring_buffer__reserve_blocking`
    * `user_ring_buffer__submit`
    * `user_ring_buffer__discard`
    * `user_ring_buffer__free`
* Perf buffer functions
    * `perf_buffer__new`
    * `perf_buffer__new_raw`
    * `perf_buffer__free`
    * `perf_buffer__epoll_fd`
    * `perf_buffer__poll`
    * `perf_buffer__consume`
    * `perf_buffer__consume_buffer`
    * `perf_buffer__buffer_cnt`
    * `perf_buffer__buffer_fd`
    * `perf_buffer__buffer`
* Program line info functions
    * `bpf_prog_linfo__free`
    * `bpf_prog_linfo__new`
    * `bpf_prog_linfo__lfind_addr_func`
    * `bpf_prog_linfo__lfind`
* Linker functions
    * `bpf_linker__new`
    * `bpf_linker__add_file`
    * `bpf_linker__finalize`
    * `bpf_linker__free`
* Misc libbpf functions
    * [`libbpf_major_version`](libbpf_major_version.md)
    * [`libbpf_minor_version`](libbpf_minor_version.md)
    * `libbpf_version_string`
    * [`libbpf_strerror`](libbpf_strerror.md)
    * `libbpf_bpf_attach_type_str`
    * `libbpf_bpf_link_type_str`
    * `libbpf_bpf_map_type_str`
    * `libbpf_bpf_prog_type_str`
    * `libbpf_set_print`
    * `libbpf_prog_type_by_name`
    * `libbpf_attach_type_by_name`
    * `libbpf_find_vmlinux_btf_id`
    * `libbpf_probe_bpf_prog_type`
    * `libbpf_probe_bpf_map_type`
    * `libbpf_probe_bpf_helper`
    * `libbpf_num_possible_cpus`
    * `libbpf_register_prog_handler`
    * `libbpf_unregister_prog_handler`

## BTF APIs

In the `btf.h` header file you will find the BTF APIs, to do more advanced BTF operations other than just loading BTF
from an object file.

* `btf__free`
* `btf__new`
* `btf__new_split`
* `btf__new_empty`
* `btf__new_empty_split`
* `btf__distill_base`
* `btf__parse`
* `btf__parse_split`
* `btf__parse_elf`
* `btf__parse_elf_split`
* `btf__parse_raw`
* `btf__parse_raw_split`
* `btf__load_vmlinux_btf`
* `btf__load_module_btf`
* `btf__load_from_kernel_by_id`
* `btf__load_from_kernel_by_id_split`
* `btf__load_into_kernel`
* `btf__find_by_name`
* `btf__find_by_name_kind`
* `btf__type_cnt`
* `btf__base_btf`
* `btf__type_by_id`
* `btf__pointer_size`
* `btf__set_pointer_size`
* `btf__endianness`
* `btf__set_endianness`
* `btf__resolve_size`
* `btf__resolve_type`
* `btf__align_of`
* `btf__fd`
* `btf__set_fd`
* `btf__raw_data`
* `btf__name_by_offset`
* `btf__str_by_offset`
* `btf_ext__new`
* `btf_ext__free`
* `btf_ext__raw_data`
* `btf_ext__endianness`
* `btf_ext__set_endianness`
* `btf__find_str`
* `btf__add_str`
* `btf__add_type`
* `btf__add_btf`
* `btf__add_int`
* `btf__add_float`
* `btf__add_ptr`
* `btf__add_array`
* `btf__add_struct`
* `btf__add_union`
* `btf__add_field`
* `btf__add_enum`
* `btf__add_enum_value`
* `btf__add_enum64`
* `btf__add_enum64_value`
* `btf__add_fwd`
* `btf__add_typedef`
* `btf__add_volatile`
* `btf__add_const`
* `btf__add_restrict`
* `btf__add_type_tag`
* `btf__add_func`
* `btf__add_func_proto`
* `btf__add_func_param`
* `btf__add_var`
* `btf__add_datasec`
* `btf__add_datasec_var_info`
* `btf__add_decl_tag`
* `btf__dedup`
* `btf__relocate`
* `btf_dump__new`
* `btf_dump__free`
* `btf_dump__dump_type`
* `btf_dump__emit_type_decl`
* `btf_dump__dump_type_data`

## Low level APIs

In the `bpf.h` header file you will find the low level APIs which are used to interact with the kernel. These are basically just wrappers around the `bpf()` syscall. You should only use these if you know what you are doing and can't do something with the high level APIs.

* `libbpf_set_memlock_rlim`
* `bpf_map_create`
* `bpf_prog_load`
* `bpf_btf_load`
* `bpf_map_update_elem`
* `bpf_map_lookup_elem`
* `bpf_map_lookup_elem_flags`
* `bpf_map_lookup_and_delete_elem`
* `bpf_map_lookup_and_delete_elem_flags`
* `bpf_map_delete_elem`
* `bpf_map_delete_elem_flags`
* `bpf_map_get_next_key`
* `bpf_map_freeze`
* `bpf_map_delete_batch`
* `bpf_map_lookup_batch`
* `bpf_map_lookup_and_delete_batch`
* `bpf_map_update_batch`
* `bpf_obj_pin`
* `bpf_obj_pin_opts`
* `bpf_obj_get`
* `bpf_obj_get_opts`
* `bpf_prog_attach`
* `bpf_prog_detach`
* `bpf_prog_detach2`
* `bpf_prog_attach_opts`
* `bpf_prog_detach_opts`
* `bpf_link_create`
* `bpf_link_detach`
* `bpf_link_update`
* `bpf_iter_create`
* `bpf_prog_get_next_id`
* `bpf_map_get_next_id`
* `bpf_btf_get_next_id`
* `bpf_link_get_next_id`
* `bpf_prog_get_fd_by_id`
* `bpf_prog_get_fd_by_id_opts`
* `bpf_map_get_fd_by_id`
* `bpf_map_get_fd_by_id_opts`
* `bpf_btf_get_fd_by_id`
* `bpf_btf_get_fd_by_id_opts`
* `bpf_link_get_fd_by_id`
* `bpf_link_get_fd_by_id_opts`
* `bpf_obj_get_info_by_fd`
* `bpf_prog_get_info_by_fd`
* `bpf_map_get_info_by_fd`
* `bpf_btf_get_info_by_fd`
* `bpf_link_get_info_by_fd`
* `bpf_prog_query_opts`
* `bpf_prog_query`
* `bpf_raw_tracepoint_open_opts`
* `bpf_raw_tracepoint_open`
* `bpf_task_fd_query`
* `bpf_enable_stats`
* `bpf_prog_bind_map`
* `bpf_prog_test_run_opts`
* `bpf_token_create`

## Deprecated APIs

In the [`libbpf_legacy.h`](https://github.com/libbpf/libbpf/blob/master/src/libbpf_legacy.h) header file you will find the deprecated APIs. These are APIs that are no longer recommended to use and might be removed in the future.

* `libbpf_set_strict_mode`
* `libbpf_get_error`
* `libbpf_find_kernel_btf`
* [`bpf_program__get_type`](bpf_program__get_type.md)
* [`bpf_program__get_expected_attach_type`](bpf_program__get_expected_attach_type.md)
* `bpf_map__get_pin_path`
* `btf__get_raw_data`
* `btf_ext__get_raw_data`
