---
title: KFuncs (Linux)
description: This page lists all KFuncs (Kernel Functions) that are available in the Linux kernel. They are categorized based on their functionality.
hide: toc
---

# KFuncs (Linux)

## cGroup RStat KFuncs

These KFuncs are used to update or flush cgroup rstats efficiently.

- [cgroup_rstat_updated](cgroup_rstat_updated.md)
- [cgroup_rstat_flush](cgroup_rstat_flush.md)

## Key signature verification KFuncs

These KFuncs are used to verify PKCS#7 signed data against keys from a keyring. 

- [bpf_lookup_user_key](bpf_lookup_user_key.md)
- [bpf_lookup_system_key](bpf_lookup_system_key.md)
- [bpf_key_put](bpf_key_put.md)
- [bpf_verify_pkcs7_signature](bpf_verify_pkcs7_signature.md)

## File related kfuncs

- [bpf_get_file_xattr](bpf_get_file_xattr.md)

## CPU mask KFuncs

- [bpf_cpumask_create](bpf_cpumask_create.md)
- [bpf_cpumask_release](bpf_cpumask_release.md)
- [bpf_cpumask_acquire](bpf_cpumask_acquire.md)
- [bpf_cpumask_first](bpf_cpumask_first.md)
- [bpf_cpumask_first_zero](bpf_cpumask_first_zero.md)
- [bpf_cpumask_first_and](bpf_cpumask_first_and.md)
- [bpf_cpumask_set_cpu](bpf_cpumask_set_cpu.md)
- [bpf_cpumask_clear_cpu](bpf_cpumask_clear_cpu.md)
- [bpf_cpumask_test_cpu](bpf_cpumask_test_cpu.md)
- [bpf_cpumask_test_and_set_cpu](bpf_cpumask_test_and_set_cpu.md)
- [bpf_cpumask_test_and_clear_cpu](bpf_cpumask_test_and_clear_cpu.md)
- [bpf_cpumask_setall](bpf_cpumask_setall.md)
- [bpf_cpumask_clear](bpf_cpumask_clear.md)
- [bpf_cpumask_and](bpf_cpumask_and.md)
- [bpf_cpumask_or](bpf_cpumask_or.md)
- [bpf_cpumask_xor](bpf_cpumask_xor.md)
- [bpf_cpumask_equal](bpf_cpumask_equal.md)
- [bpf_cpumask_intersects](bpf_cpumask_intersects.md)
- [bpf_cpumask_subset](bpf_cpumask_subset.md)
- [bpf_cpumask_empty](bpf_cpumask_empty.md)
- [bpf_cpumask_full](bpf_cpumask_full.md)
- [bpf_cpumask_copy](bpf_cpumask_copy.md)
- [bpf_cpumask_any_distribute](bpf_cpumask_any_distribute.md)
- [bpf_cpumask_any_and_distribute](bpf_cpumask_any_and_distribute.md)
- [bpf_cpumask_weight](bpf_cpumask_weight.md)
  
## Generic KFuncs

- [crash_kexec](crash_kexec.md)
- [bpf_throw](bpf_throw.md)

## Object allocation KFuncs

A set of KFuncs to allocate and deallocate custom objects for the purposes of building custom data structures.

- [bpf_obj_new_impl](bpf_obj_new_impl.md)
- [bpf_percpu_obj_new_impl](bpf_percpu_obj_new_impl.md)
- [bpf_obj_drop_impl](bpf_obj_drop_impl.md)
- [bpf_percpu_obj_drop_impl](bpf_percpu_obj_drop_impl.md)
- [bpf_refcount_acquire_impl](bpf_refcount_acquire_impl.md)
- [bpf_list_push_front_impl](bpf_list_push_front_impl.md)
- [bpf_list_push_back_impl](bpf_list_push_back_impl.md)
- [bpf_list_pop_front](bpf_list_pop_front.md)
- [bpf_list_pop_back](bpf_list_pop_back.md)

## BPF task KFuncs

Kfuncs used to aquire and release task reference.

- [bpf_task_acquire](bpf_task_acquire.md)
- [bpf_task_release](bpf_task_release.md)

## BPF cgroup KFuncs

Kfuncs used to create and modify red-black trees.

- [bpf_rbtree_add_impl](bpf_rbtree_add_impl.md)
- [bpf_rbtree_first](bpf_rbtree_first.md)
- [bpf_rbtree_remove](bpf_rbtree_remove.md)

## Kfuncs for aquiring and releasing cgroup references

These kfuncs allow you to take a reference to a cgroup and store them as kptrs in maps.

- [bpf_cgroup_acquire](bpf_cgroup_acquire.md)
- [bpf_cgroup_release](bpf_cgroup_release.md)
- [bpf_cgroup_ancestor](bpf_cgroup_ancestor.md)
- [bpf_cgroup_from_id](bpf_cgroup_from_id.md)

## Kfuncs for querying tasks

- [bpf_task_under_cgroup](bpf_task_under_cgroup.md)
- [bpf_task_get_cgroup1](bpf_task_get_cgroup1.md)
- [bpf_task_from_pid](bpf_task_from_pid.md)

## Kfuncs for casting pointers

- [bpf_cast_to_kern_ctx](bpf_cast_to_kern_ctx.md)
- [bpf_rdonly_cast](bpf_rdonly_cast.md)

## Kfuncs for taking and releasing RCU read locks

- [bpf_rcu_read_lock](bpf_rcu_read_lock.md)
- [bpf_rcu_read_unlock](bpf_rcu_read_unlock.md)

## Kfuncs for dynamic pointer slices

- [bpf_dynptr_slice](bpf_dynptr_slice.md)
- [bpf_dynptr_slice_rdwr](bpf_dynptr_slice_rdwr.md)

## Kfuncs for open coded numeric iterators

- [bpf_iter_num_new](bpf_iter_num_new.md)
- [bpf_iter_num_next](bpf_iter_num_next.md)
- [bpf_iter_num_destroy](bpf_iter_num_destroy.md)

## Kfuncs for open coded VMA iterators

- [bpf_iter_task_vma_new](bpf_iter_task_vma_new.md)
- [bpf_iter_task_vma_next](bpf_iter_task_vma_next.md)
- [bpf_iter_task_vma_destroy](bpf_iter_task_vma_destroy.md)

## Kfuncs for open coded task cGroup iterators

- [bpf_iter_css_task_new](bpf_iter_css_task_new.md)
- [bpf_iter_css_task_next](bpf_iter_css_task_next.md)
- [bpf_iter_css_task_destroy](bpf_iter_css_task_destroy.md)

## Kfuncs for open coded cGroup iterators

- [bpf_iter_css_new](bpf_iter_css_new.md)
- [bpf_iter_css_next](bpf_iter_css_next.md)
- [bpf_iter_css_destroy](bpf_iter_css_destroy.md)

## Kfuncs for open coded task iterators

- [bpf_iter_task_new](bpf_iter_task_new.md)
- [bpf_iter_task_next](bpf_iter_task_next.md)
- [bpf_iter_task_destroy](bpf_iter_task_destroy.md)

## Kfuncs for dynamic pointers 

- [bpf_dynptr_adjust](bpf_dynptr_adjust.md)
- [bpf_dynptr_is_null](bpf_dynptr_is_null.md)
- [bpf_dynptr_is_rdonly](bpf_dynptr_is_rdonly.md)
- [bpf_dynptr_size](bpf_dynptr_size.md)
- [bpf_dynptr_clone](bpf_dynptr_clone.md)

## Preemption KFuncs

- [bpf_preempt_disable](bpf_preempt_disable.md)
- [bpf_preempt_enable](bpf_preempt_enable.md)

## Workqueue KFuncs

- [bpf_wq_init](bpf_wq_init.md)
- [bpf_wq_set_callback_impl](bpf_wq_set_callback_impl.md)
- [bpf_wq_start](bpf_wq_start.md)

## Misc KFuncs

- [bpf_map_sum_elem_count](bpf_map_sum_elem_count.md)
- [bpf_get_fsverity_digest](bpf_get_fsverity_digest.md)

## XDP metadata kfuncs

- [bpf_xdp_metadata_rx_timestamp](bpf_xdp_metadata_rx_timestamp.md)
- [bpf_xdp_metadata_rx_hash](bpf_xdp_metadata_rx_hash.md)
- [bpf_xdp_metadata_rx_vlan_tag](bpf_xdp_metadata_rx_vlan_tag.md)


## XDP/SKB dynamic pointer kfuncs

- [bpf_dynptr_from_skb](bpf_dynptr_from_skb.md)
- [bpf_dynptr_from_xdp](bpf_dynptr_from_xdp.md)

## Socket related kfuncs

- [bpf_sock_addr_set_sun_path](bpf_sock_addr_set_sun_path.md)
- [bpf_sock_destroy](bpf_sock_destroy.md)

## Network crypto kfuncs

- [bpf_crypto_ctx_create](bpf_crypto_ctx_create.md)
- [bpf_crypto_ctx_acquire](bpf_crypto_ctx_acquire.md)
- [bpf_crypto_ctx_release](bpf_crypto_ctx_release.md)
- [bpf_crypto_decrypt](bpf_crypto_decrypt.md)
- [bpf_crypto_encrypt](bpf_crypto_encrypt.md)

## BBR congestion control kfuncs

- [bbr_init](bbr_init.md)
- [bbr_main](bbr_main.md)
- [bbr_sndbuf_expand](bbr_sndbuf_expand.md)
- [bbr_undo_cwnd](bbr_undo_cwnd.md)
- [bbr_cwnd_event](bbr_cwnd_event.md)
- [bbr_ssthresh](bbr_ssthresh.md)
- [bbr_min_tso_segs](bbr_min_tso_segs.md)
- [bbr_set_state](bbr_set_state.md)

## Cubic TCP congestion control kfuncs

- [cubictcp_init](cubictcp_init.md)
- [cubictcp_recalc_ssthresh](cubictcp_recalc_ssthresh.md)
- [cubictcp_cong_avoid](cubictcp_cong_avoid.md)
- [cubictcp_state](cubictcp_state.md)
- [cubictcp_cwnd_event](cubictcp_cwnd_event.md)
- [cubictcp_acked](cubictcp_acked.md)

## DC TCP congestion control kfuncs

- [dctcp_init](dctcp_init.md)
- [dctcp_update_alpha](dctcp_update_alpha.md)
- [dctcp_cwnd_event](dctcp_cwnd_event.md)
- [dctcp_ssthresh](dctcp_ssthresh.md)
- [dctcp_cwnd_undo](dctcp_cwnd_undo.md)
- [dctcp_state](dctcp_state.md)

## TCP Reno congestion control kfuncs

- [tcp_reno_ssthresh](tcp_reno_ssthresh.md)
- [tcp_reno_cong_avoid](tcp_reno_cong_avoid.md)
- [tcp_reno_undo_cwnd](tcp_reno_undo_cwnd.md)
- [tcp_slow_start](tcp_slow_start.md)
- [tcp_cong_avoid_ai](tcp_cong_avoid_ai.md)

## Foo over UDP KFuncs

- [bpf_skb_set_fou_encap](bpf_skb_set_fou_encap.md)
- [bpf_skb_get_fou_encap](bpf_skb_get_fou_encap.md)

## Contrack KFuncs

- [bpf_ct_set_nat_info](bpf_ct_set_nat_info.md)
- [bpf_xdp_ct_alloc](bpf_xdp_ct_alloc.md)
- [bpf_xdp_ct_lookup](bpf_xdp_ct_lookup.md)
- [bpf_skb_ct_alloc](bpf_skb_ct_alloc.md)
- [bpf_skb_ct_lookup](bpf_skb_ct_lookup.md)
- [bpf_ct_insert_entry](bpf_ct_insert_entry.md)
- [bpf_ct_release](bpf_ct_release.md)
- [bpf_ct_set_timeout](bpf_ct_set_timeout.md)
- [bpf_ct_change_timeout](bpf_ct_change_timeout.md)
- [bpf_ct_set_status](bpf_ct_set_status.md)
- [bpf_ct_change_status](bpf_ct_change_status.md)

## XFRM KFuncs

- [bpf_skb_get_xfrm_info](bpf_skb_get_xfrm_info.md)
- [bpf_skb_set_xfrm_info](bpf_skb_set_xfrm_info.md)
- [bpf_xdp_get_xfrm_state](bpf_xdp_get_xfrm_state.md)
- [bpf_xdp_xfrm_state_release](bpf_xdp_xfrm_state_release.md)

## HID Kfuncs

- [hid_bpf_get_data](hid_bpf_get_data.md)
- [hid_bpf_attach_prog](hid_bpf_attach_prog.md)
- [hid_bpf_allocate_context](hid_bpf_allocate_context.md)
- [hid_bpf_release_context](hid_bpf_release_context.md)
- [hid_bpf_hw_request](hid_bpf_hw_request.md)

## KProbe session Kfuncs
  - [bpf_session_cookie](bpf_session_cookie.md)
  - [bpf_session_is_return](bpf_session_is_return.md)
