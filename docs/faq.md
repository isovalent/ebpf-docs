# Frequently Asked Questions

There are some frequently asked questions by people seeing these docs for the first time. Hopefully this Q/A section can answer most of them. If still left wondering about something please reach out in the `#ebpf` channel on the [Cilium & eBPF slack](https://cilium.herokuapp.com/) or open a GH issue.

## What is the goal of this project?

The goal of this project is to provide a single, searchable, and interlinked resource for eBPF related information.

## Who are these docs for? What is the target audience?

These docs are (ideally) for everyone in the eBPF community, specifically

* eBPF program authors (from beginner to power-user)
* eBPF library maintainers
* eBPF researchers/intellectuals (those who need to or want to know about eBPF without working on or with eBPF directly for a number of reasons)

## What is the scope of the project? Will you add `XYZ`?

The Linux kernel, libraries maintained alongside the Linux kernel as "reference implementation" (specifically `iproute2`, `libbpf`, `libxdp`), eBPF on Windows (when it matures). 

## Why did this project get started?

At the time of project creation there were a couple of different resources available for eBPF which all told a piece of the puzzle. A list of helper functions can be found at `A`. A list of program types could be found at `B` but not description of how to use them. And at `C` a tutorial can be found on how to use a specific map type, but not how to get started with eBPF. 

Most resources do not include information about when features were added which turns out to be very important when running and distributing eBPF based software. And while not a replacement for feature probing, knowing roughly which kernel version compatibility upfront can be a huge lifesaver during development.

Lastly, it turns out a lot of features have limitations or interactions which are rarely documented. A reference of which helper functions work in which programs for example did not exist because no resource ever documented both a the same time.

## Why a separate project? Don't these docs belong in the kernel tree?

The usefulness of documentation grows exponentially the closer it gets to covering "all there is to know" about a subject. The more useful a project is the more people can be found to contribute to it. And the more people to contribute to docs the more it can cover. So there is a positive feedback loop here and we have a lot to cover, it is therefore important that the friction of adding coverage be as low as possible, at least at the start.

While the kernel is technically the correct place for docs about about a kernel feature, getting changes into the kernel is a long process with a lot of friction. This is a large part of the reason eBPF exists, to allow for rapid iteration without the friction of the kernel contribution process. If we were to start this effort in the kernel, we would not be able to maintain the same velocity and positive feedback as we are able to outside of the tree. And without that velocity the project will likely die once the energy of the initial contributors runs out.

Once this project has proven its worth and has the support it needs, we can start feeding docs into the kernel tree once they have had enough time to mature and be properly revised. What this process will look like and how long that is going to take is up for debate. Until then, this project will do its best to provide great eBPF docs.
