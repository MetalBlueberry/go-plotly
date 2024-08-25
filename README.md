[![Go Report Card](https://goreportcard.com/badge/github.com/MetalBlueberry/go-plotly)](https://goreportcard.com/report/github.com/MetalBlueberry/go-plotly)
[![godoc](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/mod/github.com/MetalBlueberry/go-plotly?tab=overview)

# go-plotly

Inspired by [Python Plotly](https://plotly.com/python/creating-and-updating-figures/)

The goal of the go-plotly package is to provide a pleasant Go interface for creating figure specifications which are displayed by the plotly.js JavaScript graphing library.

In the context of the plotly.js library, a figure is specified by a declarative JSON data structure.

Therefore, you should always keep in mind as you are creating and updating figures using the go-plotly package that its ultimate goal is to help users produce Go structures that can be automatically serialized into the JSON data structure that the plotly.js graphing library understands.

> Yes, that text is a copy paste from Python description.

The good thing about this package is that it's **automatically generated** based on the schema described [here](https://plotly.com/chart-studio-help/json-chart-schema/) so It will be easy to keep it up to date or to generate different versions of it.

> Just a reminder from semver: **Major version zero (0.y.z) is for initial development. Anything MAY change at any time. The public API SHOULD NOT be considered stable.**

## Example

```go
package main

import (
	grob "github.com/MetalBlueberry/go-plotly/generated/v2.31.1/graph_objects"
	"github.com/MetalBlueberry/go-plotly/pkg/offline"
	"github.com/MetalBlueberry/go-plotly/pkg/types"
)

func main() {
	/*
	   fig = dict({
	       "data": [{"type": "bar",
	                 "x": [1, 2, 3],
	                 "y": [1, 3, 2]}],
	       "layout": {"title": {"text": "A Figure Specified By Python Dictionary"}}
	   })
	*/
	fig := &grob.Fig{
		Data: []types.Trace{
			&grob.Bar{
				X: types.DataArray([]float64{1, 2, 3}),
				Y: types.DataArray([]float64{1, 2, 3}),
			},
		},
		Layout: &grob.Layout{
			Title: &grob.LayoutTitle{
				Text: "A Figure Specified By Go Struct",
			},
		},
	}

	offline.Show(fig)
}
```

This will open your browser and display the following plot

![Bar](./examples/bar/Bar.png)

And that's it.

See the examples dir for more examples.

## Structure

To keep the plotly.js independent of the version of this package, the generated directory contains a directory per plotly version supported. The plan is to support all minor releases and update as patches are released. But because I can not do it myself, I will accept PRs if anyone wants any specific version.

Updates to the package will affect all schemas. This will be done as we improve the generator.

Each trace type has its own file on **graph_objecs (grob)** package. The file contains the main structure and all the needed nested objects. Files ending with **_gen** are automatically generated files running `go generate`. This is executing the code in **generator** package to generate the structures from the plotly schema. The types are documented, but you can find more examples and extended documentation at [Plotly's documentation](https://plotly.com/python/).

The values that can hold single values or arrays are defined as custom types that are a type definition of `interfaces{}`. Most common case are X and Y values. You can pass any number slice and it will work (`[]float64`,`[]int`,`[]int64`...). In case of Hovertext, you can provide a `[]string` to display a text for each point, a `string` to display the same for all or `[]int` to display a number.

Nested Properties, are defined as new types. This is great for auto completion using vscode because you can write all the boilerplate with ctrl+space. For example, the field `Title.Text` is accessed by the property `Title` of type {{Type}}Title that contains the property `Text`. The Type is always the struct that contains the field. For Layout It is `LayoutTitle`.

Flaglist and Enumerated fields are defined with a type and its constant values. This provides useful autocompletion. Keep in mind that you can compose multiple values for a Flaglist like `Mode: grob.ScatterModeMarkers + "+" + grob.ScatterModeLines,`. You can read de inline documentation to see the default value.

## Progress

The main focus is to have the structures to hold the data and provide auto competition. Once we get there, we will be on v1.0.0.

Currently, most common use cases should be already covered. Feel free to create an issue if you find something that it's not working.

- [x] Traces
- [x] Layout
- [x] Config
- [x] Animation
- [x] Frames
  - [ ] defs
    - [ ] defs valobjects
      - [ ] angle
      - [x] any
      - [x] boolean
      - [x] color
      - [x] colorlist
      - [x] colorscale
      - [x] data_array
      - [x] enumerated
      - [x] flaglist
      - [ ] info_array
      - [ ] integer **Needs support for nil values**
      - [ ] number **Needs support for nil values**
      - [x] string
      - [ ] subplotid
    - [ ] defs_modifier
      - [ ] arrayOK
      - [ ] min/max validations
      - [x] dflt  **This is not needed in the output, as plotly will do it. But it would be nice to have a method to fetch it**
      - [ ] noBlank validation
      - [ ] strict validation
      - [ ] values list validation

## FAQ

### What's the meaning of "grob"?

In python, the component is called graph_objects, like in this package, but that's too long to write it over and over again. In Python, usually it's called "go" from Graph_Objects but in Go... that's not possible for a conflict with a keyword. as an alternative, I've decided to use "grob" from GRaph_OBjects.

### What are the usecases?

1. Quickly visualize data using only Go

2. Send plotly figures to the frontend ready to be drawn, avoiding possible mistakes in JS thanks to types!

3. Generate an awesome dynamic plot in a local file with offline package.

4. I don't know, just do something awesome and let me know it.

### Go Plotly Update to any json schema version

[Example PR](https://github.com/MetalBlueberry/go-plotly/pull/29)

#### Update the config to add a new version

1. To add a new version, add a new entry in: [schemas.yaml](schemas.yaml)
    > The documentation for each field can be found in [schema.go](generator%2Fschema.go)

    Example entry:

    ```yaml
      -  Name: Plotly 2.31.1
         Tag: v2.31.1
         URL: https://raw.githubusercontent.com/plotly/plotly.js/v2.31.1/test/plot-schema.json
         # relative to the project root.
         Path: schemas/v2.31.1/plot-schema.json
         Generated: generated/v2.31.1
         CDN: https://cdn.plot.ly/plotly-2.31.1.min.js
    ```

#### Run download and regeneration

1. Download all the schemas:

   ```shell
   go run generator/cmd/downloader/main.go --config="schemas.yaml"
   ```

2. Then run the generator:

   ```shell
   go generate ./...
   ```

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=Metalblueberry/go-plotly&type=Date)](https://star-history.com/#Metalblueberry/go-plotly&Date)

## Other tools and information links

### Plotly online editor sandbox

http://plotly-json-editor.getforge.io/

### Official Plotly Release Notes

For detailed changes please follow the release notes of the original JS repo: https://github.com/plotly/plotly.js/releases