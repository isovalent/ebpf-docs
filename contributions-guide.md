# Contributions guide

Thanks current/future contributor for taking the time to improve the eBPF docs. Before you submit your changes, please read this contributions guide to avoid any surprises during the PR feedback process.

In order to ensure the quality and consistency of the documentation, we have a few rules which should be followed:

## Submit English docs

While not everyone's native language, it is the language of the internet and the Linux kernel. Supporting multiple languages requires contributors whom speak that language to fact-check it. This is currently not a burden we are willing to undertake.

## Keep docs vendor neutral

These docs are meant for the community as a whole and we do not want to promote certain projects over others. A fine line has to be drawn between mentioning projects to inform readers of their existence and promoting one over the other.

The primary goal of the project is to document eBPF to such an extent that developers do not have to go the eBPF kernel sources to find out how to use eBPF. This will inevitably include mentioning the APIs of loader projects and eBPF kernel side libraries, and showing examples. Having that said, we do not want to document these projects.

The exception to this rule are tools and libraries which originate in the Linux kernel or are maintained by kernel developers alongside the kernel such as: `libbpf`, `libxdp`, `bpftool` and `iproute2`.

## Include copyrights, and link to original sources when copying

In general we can copy and modify documentation and code from other sources, as long as they have licenses that permit us to do so. When copying large sections such as code examples or whole paragraphs, we should preserve any copyright lines present and link back to the sources.
