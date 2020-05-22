# graph objects

Almost all content from this package is automatically generated from the schema.json. Files can be identified by `auto_` prefix. files with `.tmpl` are the template files that can be parsed with `plate` tool. The command is in `plotly.go` with go generate.

````go
//go:generate sh -c "cat schema.json | plate --template main.tmpl"
```

For development is a good idea to set an autoreload environment to automatically regenerate the files when something changes. I do it with reflex

```bash
reflex -g "*.tmpl" -- sh -c "cat schema.json | plate --template main.tmpl"
````
