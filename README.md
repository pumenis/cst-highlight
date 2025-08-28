# cst-highlight go based source highlighter based on tree-sitter cst inside html

currently only following source files are supported:

- go
- html
- js
- css
- sql
- sh - bash

installation:

```
go install github.com/pumenis/cst-highlight@main
```

usage:

```
cst-highlight -preview ./main.go>/tmp/maingo.html
xdg-open /tmp/maingo.html
```
