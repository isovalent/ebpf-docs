# Docs about the docs (meta docs)

This page documents how the docs work. For anyone interested in contributing to the docs, or would like to know what goes on under the hood.

## <nospell>MkDocs</nospell>

All docs sources are written in Markdown and rendered to static HTML using [<nospell>MkDocs</nospell>](https://www.mkdocs.org/). <nospell>MkDocs</nospell> uses the [Python-Markdown](https://python-markdown.github.io/) library to actually do the parsing and rendering. 

The rendered HTML itself is very plain, a theme is used to provide styling and additional features. We use the [Material for <nospell>MkDocs</nospell>](https://squidfunk.github.io/mkdocs-material/reference/) theme. The [`docs/stylesheets/extra.css`](https://github.com/isovalent/ebpf-docs/blob/master/docs/stylesheets/extra.css) file contains our own custom CSS to make any fine adjustments.

The [`mkdocs.yml`](https://github.com/isovalent/ebpf-docs/blob/master/mkdocs.yml) file is used to configure most of the parsing and rendering process.

Both [<nospell>MkDocs</nospell>](https://www.mkdocs.org/) and [Python-Markdown](https://python-markdown.github.io/) are extensible. The 

The `plugins` configured in `mkdocs.yml` are MkDocs plugins provided via [pip](https://pypi.org/project/pip/). The `hooks` configured in `mkdocs.yml` are essentially plugins, but defined inside the project. And the `markdown_extensions` in the `mkdocs.yml` are the Python-markdown extensions, again provided via [pip](https://pypi.org/project/pip/).

All pip dependencies are defined in [`requirements.txt`](https://github.com/isovalent/ebpf-docs/blob/master/requirements.txt).

## Markdown dialect

In the basis the [Markdown syntax by <nospell>John Gruber</nospell>](https://daringfireball.net/projects/markdown/syntax) apply, with a few [differences](https://python-markdown.github.io/#differences).

Layered on top is functionality added by <nospell>add-ons</nospell>, plugins, and the theme.

### Admonitions

This is a feature that allows us to make nice looking inline notes like:

```
!!! note
    This is a note
```

!!! note
    This is a note


Much more is possible, see the [theme docs](https://squidfunk.github.io/mkdocs-material/reference/admonitions/) for details.

### Code blocks

A multi-line code block can be defined with three back-ticks. Like so

\`\`\`
Code example
\`\`\`

```
Code example
```

Code is highlighted by [`Pygments`](https://pygments.org/). This works best when instructed on the programming language used:

````
```c
struct abc {
    uint_t some_field;
};
```
````

```c
struct abc {
    uint_t some_field;
};
```

See the [theme docs](https://squidfunk.github.io/mkdocs-material/reference/code-blocks/) for all options.

#### Links in code blocks

A [custom hook](https://github.com/isovalent/ebpf-docs/blob/master/hooks/links_in_code.py) gives us the ability to use Markdown style links inside of code blocks like so:

````
```c
struct abc {
    uint_t [some_field](#codeblocks);
};
```
````

```c
struct abc {
    uint_t [some_field](#codeblocks);
};
```

This allows us to for example add a link to a function definition.

### Emojis and symbols

Emojis and symbols can be inserted with the `:<name of emoji>:` syntax. These docs in particular make heavy use of the `:octicons-tag-24:` :octicons-tag-24: emoji to signal a version tag, most for Linux kernel versions.

The [theme docs](https://squidfunk.github.io/mkdocs-material/reference/icons-emojis/) contains a search function for all available emojis.

### Footnotes

This is a feature that allows for the creation of footnotes [^1]. See [theme docs](https://squidfunk.github.io/mkdocs-material/reference/footnotes/) for details.

[^1]: Example of a footnote

## Markdown plugins

### Search

The [search plugin](https://squidfunk.github.io/mkdocs-material/plugins/search/) creates a search index from our pages at render time. This index is then used by JavaScript in the front-end to provide a search bar. This is all automatic.

### Git revision date

The [`git-revision-date-localized`](https://github.com/timvink/mkdocs-git-revision-date-localized-plugin) plugin uses git to see when the current page was created and updated last and includes this information in the rendered HTML.

### Git committers

The [`git-committers`](https://github.com/ojacques/mkdocs-git-committers-plugin-2) plugin uses the git history to see who contributed to a given page and adds links to their GitHub profiles to the bottom of the page. Our way of providing credit to contributors.

### Social

The [`social`](https://squidfunk.github.io/mkdocs-material/plugins/social/) plugin creates images and links them in the page metadata such that these images are used on social media platforms as previews when links to the site are shared.

## Tools

This project uses a bunch of custom tools to generate specific sections of the docs. This makes writing and maintaining these sections much lower effort then without tools.

### Spellcheck

The [`spellcheck`](https://github.com/isovalent/ebpf-docs/tree/master/tools/spellcheck) tool is a spellchecker that is tailored to the needs of this project. It operates on the rendered output of the site (its HTML) since this is easier to parse then our custom markdown dialect.

The tool parses the HTML and filters out a number of tags which we want to ignore for the purposes of spellchecking, such as `<script>`, `<style>` and `<nav>`, but also `<code>` and `<nospell>`.

The `<code>` block is emitted for any multi-line (\`\`\`...\`\`\`) or single line (\`...\`) code blocks, since these are used for actual code, we exclude them.

The `<nospell>` is a custom HTML tag, which has no meaning to browsers and exists for the purposes of ignoring spellchecking. This is particularly useful when you want to include non-standard words, but you are not planning on using them regularly enough to warrant adding it to the dictionary, like names of people, projects, protocols.

The spellchecker also ignores text between certain markers such as `<!-- [HELPER_FUNC_DEF] -->`/`<!-- [/HELPER_FUNC_DEF] -->` and `<!-- [MTU_TABLE] -->`/`<!-- [/MTU_TABLE] -->`, since these are generated and modifying them would be undone by other tools.

All of the text we do consider is passed to [`aspell`](http://aspell.net/) which does the actual spellchecking against its builtin English dictionary and our supplementary dictionary in [`.aspell.en.pws`](https://github.com/isovalent/ebpf-docs/tree/master/.aspell.en.pws).

If `aspell` marks a word as misspelled, then this wrapper will try to work backwards from the HTML to find which markdown file that must have come from. It then will try and find occurrences of that same word and tell you the possible locations at fault in the markdown.

Its advised to run this tool locally before submitting pull requests with `make spellcheck`.

When a misspelled word is found consider the following solutions:

* Check if you actually misspelled the word, that is what the tool is for after all, and use one of the suggestions given.
* If the word is code or a code reference (names of variables and functions are often abbreviated or truncated), then consider putting them in a code block (multi-line or inline).
* If a word is an abbreviation, but not a code reference, consider spelling it out instead of adding an exception, since it likely will improve readability for someone unfamiliar with the jargon. (file descriptor instead of <nospell>fd</nospell>)
* If a word is very uncommon, like a name (<nospell>Lempel–Ziv–Welch</nospell>). Consider surrounding it with `<nospell>` and `</nospell>`.
* If a word is commonly used (in this project), and not one of the above, add it to the [`.aspell.en.pws`](https://github.com/isovalent/ebpf-docs/tree/master/.aspell.en.pws).

### Version finder

The [version-finder](https://github.com/isovalent/ebpf-docs/tree/master/tools/version-finder) finds the tag and commit in which certain symbols were added to a git repository. It outputs the result in the [`data/feature-versions.yaml`](https://github.com/isovalent/ebpf-docs/blob/master/data/feature-versions.yaml) file to be used by other tools.

This tool isn't called in CI or as part of any make command since it is very slow to run and needs a Linux repository with full history which is very large. Manual effort is needed to make it work.

### Feature tag generator

The [feature-gen](https://github.com/isovalent/ebpf-docs/tree/master/tools/feature-gen) scans all files in a list of hard-coded directories for the `<!-- [FEATURE_TAG](...) -->` and `<!-- [/FEATURE_TAG] -->` HTML comment. When it finds such a marker it looks up the name provided in the parenthesis in the [`data/feature-versions.yaml`](https://github.com/isovalent/ebpf-docs/blob/master/data/feature-versions.yaml) file. It then replaced the text between the start and end marker with the name of the Linux version and a link to the exact commit that symbol was added.

When invoked, the tool also generates the full [timeline](https://github.com/isovalent/ebpf-docs/blob/master/docs/linux/timeline/index.md) page.

### Helper definition scraper

The [helper-def-scraper](https://github.com/isovalent/ebpf-docs/tree/master/tools/helper-def-scraper) downloads the [`bpf_helper_defs.h`](https://github.com/libbpf/libbpf/blob/master/src/bpf_helper_defs.h) from the libbpf project and extracts the descriptions of all helper functions.

It matches up the helper functions with the corresponding helper page, converts the comment into our flavor of markdown, and then adds that markdown between the `<!-- [HELPER_FUNC_DEF] -->` and `<!-- [/HELPER_FUNC_DEF] -->` marker. This ensures our helpers keep an up-to-date definition unless we decide to take ownership.

### Helper reference generator

The [helper-def-gen](https://github.com/isovalent/ebpf-docs/tree/master/tools/helper-def-gen) generates the lists of references between program types and helper functions.

In the kernel, access to helper functions is gated behind a function set on the `get_func_proto` field of the `struct bpf_verifier_ops` of a program type when its registered. The function takes the helper function number / id and converts it into a `struct bpf_func_proto *`, or `NULL` if the helper is not allowed or doesn't exist.

These root functions assigned to `get_func_proto` often call out to other functions which are reused by multiple program types. This makes it difficult to manually resolve all allowed helpers for all program types. 

Our approach is to model these functions (we call them groups) in a <nospell>YAML</nospell> data file (`data/helpers-functions.yaml`). The `helper-def-gen` takes this data file, and resolves the group references so we end up with a flat list of program type <-> helper functions. 

The tool generates lists of markdown for each program type and helper function. For helper functions it searches for the `<!-- [HELPER_FUNC_PROG_REF] -->` and `<!-- [/HELPER_FUNC_PROG_REF] -->` markers on every page and adds the markdown in between. For program types it searches for the `<!-- [PROG_HELPER_FUNC_REF] -->` and `<!-- [/PROG_HELPER_FUNC_REF] -->` markers to place the generated markdown between.

### Kfunc generator

The [kfunc-gen](https://github.com/isovalent/ebpf-docs/tree/master/tools/kfunc-gen) generates both the kfunc definitions (their signatures at least) and it generates the program-type <-> kfunc references.

This tool has two inputs, the `data/kfuncs.yaml` data file, and a vmlinux BTF blob. Kfuncs are defined in sets, every set can be registered to one or more program types. The data file contains these mappings and other interesting information such as the `KF_*` flags for each set. 

The vmlinux contains the "types" of the kfuncs from which the tool generates the C function signatures that users need to forward declare. The vmlinux also allows us to enumerate all known kfuncs, which serves as a check to ensure we are not missing any.

The tool will render markdown with the function signature and notes / warnings for `KF_*` flags. This markdown is inserted between `<!-- [KFUNC_DEF] -->` and `<!-- [/KFUNC_DEF] -->` markers on the kfunc pages.

A reference list of kfunc -> program types is rendered and inserted between the `<!-- [KFUNC_PROG_REF] -->` and `<!-- [/KFUNC_PROG_REF] -->` marker on the kfunc pages.

A reference list of program types -> kfuncs is rendered and inserted between the `<!-- [PROG_KFUNC_REF] -->` and `<!-- [/PROG_KFUNC_REF] -->` marker on the program type pages.

Running this tool can be done with the `vmlinux` file that is embedded in the repo. This file has to be updated manually when newer kernel versions are available. Instructions can be found at [`tools/kfunc-gen/vmlinux-update.md`](https://github.com/isovalent/ebpf-docs/blob/master/tools/kfunc-gen/vmlinux-update.md)

### MTU calculator

The [`mtu-calc`](https://github.com/isovalent/ebpf-docs/tree/master/tools/mtu-calc/) tool calculates the maximum MTU that can be used with a given network device/driver when XDP is also enabled.

The reason this tool exists is because every network driver has its own logic for determining the this maximum MTU number it can handle. They differ since different drivers and NICs need different amounts of extra overhead for metadata or hardware specific reasons. In addition a lot of macros are used which might change this calculation based on factors like the CPU architecture. Most drivers do not advertise this number, even when the limit is hit. And manually calculating every permutation is cumbersome.

The tool works by essentially modeling the same calculations found in the C code of the driver in Go, where we can vary the inputs. The tool then creates a table of the most common architectures plus factors such as XDP fragments. The rendered table is inserted between the `<!-- [MTU_TABLE] -->` and `<!-- [/MTU_TABLE] -->` marker on the [`BPF_PROG_TYPE_XDP`](linux/program-type/BPF_PROG_TYPE_XDP.md) page.

### Libbpf tag generator

The [`libbpf-tag-gen`](https://github.com/isovalent/ebpf-docs/tree/master/tools/libbpf-tag-gen/) tool creates version tags for [Libbpf userspace](ebpf-library/libbpf/userspace/) functions. It parses the [`libbpf.map`](https://github.com/libbpf/libbpf/blob/master/src/libbpf.map) file to see when certain functions were added to the library and creates a version number tag and link which is inserted between the `<!-- [LIBBPF_TAG] -->` and `<!-- [/LIBBPF_TAG] -->` markers on the Libbpf userspace function pages.
