# eBPF Docs

## Serving docs

In order to see the docs in browser you need to serve them locally. This can be done with docker via the makefile.
```
$ make serve
```

Or you can serve from you own host by following these steps:
```
$ pip install -r requirements.txt
$ mkdocs serve -a 127.0.0.1:8000 --watch /docs
```
