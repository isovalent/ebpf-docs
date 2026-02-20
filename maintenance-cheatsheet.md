# Maintenance cheatsheet

This is a handy cheatsheet of commands which can be used on a checked out kernel git repo to see which changes have been made since last kernel release. Making the docs maintenance a bit easier.

## Keeping track of BPF changes

* `git log -S "__bpf_kfunc"` - Changes to kfuncs
* `git log -S "ndo_xdp_xmit"` - XDP driver support
* `git log -S "xdp_frame_has_frags"` - XDP frags support
* `git log {prev tag}..{curr tag} -S "bpf" --pretty=oneline kernel/bpf kernel/trace kernel/events kernel/sched net/ security/ fs/`
  To get a verbose list of commits that did something BPF related.
