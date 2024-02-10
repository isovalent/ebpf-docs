# eBPF Docs

This project aims to provide documentation for eBPF, with a specific focus on technical details for developers of projects that use eBPF.

## Getting started - Serving docs

In order to see the docs in browser you need to serve them locally. This can be done with docker via the makefile.
```
$ make serve
```

Or you can serve from you own host by following these steps:
```
$ python3 -m venv .venv
$ source .venv/bin/activate
$ pip install -r requirements.txt
$ mkdocs serve -a 127.0.0.1:8000 --watch /docs
```

## Docs about docs

Use use MkDocs to render the markdown in this repository to HTML (and to serve it in development). In addition we use the Material for MkDocs theme and some additional plugins to add features. The following links contains documents for varies components including guides on the markdown syntax and references for our config files:

* https://www.mkdocs.org/
* https://squidfunk.github.io/mkdocs-material/reference/
* https://facelessuser.github.io/pymdown-extensions/extensions/arithmatex/

## Contributing

This project is meant to provide a common knowledge base of the whole eBPF community, everyone is free to submit changes via Github Pull Requests, please read our [Contributions Guide](./contributions-guide.md) for details and guidelines.
