* [Concept](libxdp.md)
* Manage programs
  * Load
    * [`xdp_program__from_bpf_obj`](./functions/xdp_program__from_bpf_obj.md)
    * [`xdp_program__find_file`](./functions/xdp_program__find_file.md)
    * [`xdp_program__open_file`](./functions/xdp_program__open_file.md)
    * [`xdp_program__from_fd`](./functions/xdp_program__from_fd.md)
    * [`xdp_program__from_id`](./functions/xdp_program__from_id.md)
    * [`xdp_program__from_pin`](./functions/xdp_program__from_pin.md)
  * Metadata
    * [`xdp_program__run_prio`](./functions/xdp_program__run_prio.md)
    * [`xdp_program__set_run_prio`](./functions/xdp_program__set_run_prio.md)
    * [`xdp_program__chain_call_enabled`](./functions/xdp_program__chain_call_enabled.md)
    * [`xdp_program__set_chain_call_enabled`](./functions/xdp_program__set_chain_call_enabled.md)
    * [`xdp_program__print_chain_call_actions`](./functions/xdp_program__print_chain_call_actions.md)
  * Dispatcher
    * [`xdp_multiprog__get_from_ifindex`](./functions/xdp_multiprog__get_from_ifindex.md)
    * [`xdp_multiprog__next_prog`](./functions/xdp_multiprog__next_prog.md)
    * [`xdp_multiprog__close`](./functions/xdp_multiprog__close.md)
    * [`xdp_multiprog__detach`](./functions/xdp_multiprog__detach.md)
    * [`xdp_multiprog__attach_mode`](./functions/xdp_multiprog__attach_mode.md)
    * [`xdp_multiprog__main_prog`](./functions/xdp_multiprog__main_prog.md)
    * [`xdp_multiprog__hw_prog`](./functions/xdp_multiprog__hw_prog.md)
    * [`xdp_multiprog__is_legacy`](./functions/xdp_multiprog__is_legacy.md)
* AF_XDP sockets
  * Control path
    * Umem Area
      * [`xsk_umem__create`](./functions/xsk_umem__create.md)
      * [`xsk_umem__create_with_fd`](./functions/xsk_umem__create_with_fd.md)
      * [`xsk_umem__delete`](./functions/xsk_umem__delete.md)
      * [`xsk_umem__fd`](./functions/xsk_umem__fd.md)
      * [`xsk_umem__get_data`](./functions/xsk_umem__get_data.md)
      * [`xsk_umem__extract_addr`](./functions/xsk_umem__extract_addr.md)
      * [`xsk_umem__extract_offset`](./functions/xsk_umem__extract_offset.md)
      * [`xsk_umem__add_offset_to_addr`](./functions/xsk_umem__add_offset_to_addr.md)
    * Sockets
      * [`xsk_socket__create`](./functions/xsk_socket__create.md)
      * [`xsk_socket__create_shared`](./functions/xsk_socket__create_shared.md)
      * [`xsk_socket__delete`](./functions/xsk_socket__delete.md)
      * [`xsk_socket__fd`](./functions/xsk_socket__fd.md)
      * [`xsk_setup_xdp_prog`](./functions/xsk_setup_xdp_prog.md)
      * [`xsk_socket__update_xskmap`](./functions/xsk_socket__update_xskmap.md)
  * Data path
    * Producer rings
      * [`xsk_ring_prod__reserve`](./functions/xsk_ring_prod__reserve.md)
      * [`xsk_ring_prod__submit`](./functions/xsk_ring_prod__submit.md)
      * [`xsk_ring_prod__fill_addr`](./functions/xsk_ring_prod__fill_addr.md)
      * [`xsk_ring_prod__tx_desc`](./functions/xsk_ring_prod__tx_desc.md)
      * [`xsk_ring_prod__needs_wakeup`](./functions/xsk_ring_prod__needs_wakeup.md)
    * Consumer rings
      * [`xsk_ring_cons__peek`](./functions/xsk_ring_cons__peek.md)
      * [`xsk_ring_cons__cancel`](./functions/xsk_ring_cons__cancel.md)
      * [`xsk_ring_cons__release`](./functions/xsk_ring_cons__release.md)
      * [`xsk_ring_cons__comp_addr`](./functions/xsk_ring_cons__comp_addr.md)
      * [`xsk_ring_cons__rx_desc`](./functions/xsk_ring_cons__rx_desc.md)
