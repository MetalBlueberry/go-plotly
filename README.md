# go-plotly

Inspired by [Python Plotly](https://plotly.com/python/creating-and-updating-figures/)

The goal of the go-plotly package is to provide a pleasant Go interface for creating figure specifications which are displayed by the plotly.js JavaScript graphing library.

In the context of the plotly.js library, a figure is specified by a declarative JSON data structure.

Therefore, you should always keep in mind as you are creating and updating figures using the go-plotly package that its ultimate goal is to help users produce Go structures that can be automatically serialized into the JSON data structure that the plotly.js graphing library understands.

> Yes, that text is a copy paste from Python description.

The good thing about this package is that it's **automatically generated** based on the schema described [here](https://plotly.com/chart-studio-help/json-chart-schema/) so It will be easy to keep it up to date or to generate different versions of it.

## Example

```go
package main

import (
    grob "github.com/MetalBlueberry/go-plotly/graph_objects"
    "github.com/MetalBlueberry/go-plotly/plotly"
)

func main() {
    fig := &grob.Fig{
        Data: grob.Traces{
            &grob.Bar{
                Type: grob.TraceTypeBar,
                X:    []float64{1, 2, 3},
                Y:    []float64{1, 2, 3},
            },
        },
        Layout: grob.Layout{
            Title: &grob.LayoutTitle{
                Text: "A Figure Specified By Go Struct",
            },
        },
    }

    offline.Show(fig)
}
```

This will open your browser and display the following plot

![Bar](img/Bar.png)

And that's it.

See the examples dir for more examples.

## Structure

All the traces are defined in [traces.go](./graph_objects/traces.go) and identified by the interface "Trace". Feel free to browse that file, but I prefer to use [Plotly's documentation](https://plotly.com/python/).

The values that can hold single values or arrays are defined as `interfaces{}`. Most common case are X and Y values. You can pass any number slice and it will work (`[]float64`,`[]int`,`[]int64`...). In case of Hovertext, you can provide a `[]string` to display a text for each point, a `string` to display the same for all or `[]int` to display a number.

Enumerated values doesn't have a special type and you can assign whatever you want to them. But for type safety and autocompletion, you can look for the const value associated. For example, in case of "visible" you have to look for the const {{Type}}Visible{{Value}}. to hide a Scatter trace. `grob.ScatterVisibleFalse`. To known if a field is Enumerated, check the field description. the second word should be **enumerated**

Nested Properties, are defined as new types. This is great for auto completion using vscode because you can write all the boilerplate with ctrl+space. For example, the field `Title.Text` is accessed by the property `Title` of type {{Type}}Title that contains the property `Text`. The Type is always the struct that contains the field. For Layout It is `LayoutTitle`.

Flaglist Properties are defined like Enumerated values as constant fields. The difference is that they can be composed. Right now there is no special tool to deal with this and you can use the const to have autocompletion but you must be careful when joining them. For example, Scatter with mode="markers+lines" is `Mode: grob.ScatterModeMarkers + "+" + grob.ScatterModeLines,`

## Tested

Examples, Code generation and Offline package are based on version 1.54.0, but It should work with other versions as this library just generates standard JSON.

## Progress

The main focus is to have the structures to hold the data and provide auto competition. Once we get there, we will be on v1.0.0.

Currently, most common use cases should be already covered. Feel free to create an issue if you find something that it's not working.

## FAQ

### What's the meaning of "grob"?

In python, the component is called graph_objects, like in this package, but that's too long to write it over and over again. In Python, usually it's called "go" from Graph_Objects but in Go... that's not possible for a conflict with a keyword. as an alternative, I've decided to use "grob" from GRaph_OBjects.

### How are the graph_object files generated?

The tmpl files are parsed mainly with go templates with just a few extra functions that I've publish as a CLI tool called "[plate](https://github.com/MetalBlueberry/plate)". You can go there to read more about It. The good thing is that you don't need to use plate unless you want to contribute to this repo.

### What are the usecases?

1. Send plotly figures to the frontend ready to be drawn, avoiding possible mistakes in JS thanks to types!

2. Generate an awesome dynamic plot in a local file with offline package.

3. Tell me what you can do with this.