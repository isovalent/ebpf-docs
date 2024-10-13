---
title: "Libbpf"
description: "This is the index page for the libbpf docs. Libbpf is the reference library for eBPF development. It is developed and maintained as part of the kernel tree often in lockstep with the kernel itself. From this page you can explore the different parts of libbpf."
hide: toc
---
# Libbpf

Libbpf is the reference library for eBPF development. It is developed and maintained as [part of the kernel tree](https://github.com/torvalds/linux/tree/master/tools/lib/bpf) often in lockstep with the kernel itself. Since including the whole kernel tree in a project is not practical, a mirror of just the libbpf library is maintained at [https://github.com/libbpf/libbpf](https://github.com/libbpf/libbpf).

Libbpf has both userspace components and eBPF components. The eBPF components are mostly pre-processor statements, forward declarations and type definitions that make it easier to write eBPF programs. The userspace components is a library that for loading eBPF programs and interacting with the loaded resources.

<div class="grid cards" markdown>

-   __Userspace__

    ---

    The userspace loader library for eBPF programs

    [:octicons-arrow-right-24: Userspace](./userspace/index.md)

-   __eBPF side__

    ---

    The eBPF side library for ease of writing eBPF programs

    [:octicons-arrow-right-24: eBPF side](./ebpf/index.md)

-  __Concepts__

    ---

    Libbpf concepts that involve both userspace and eBPF side

    [:octicons-arrow-right-24: Concepts](./concepts/index.md)

</div>
