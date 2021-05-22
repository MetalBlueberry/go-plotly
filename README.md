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
    grob "github.com/MetalBlueberry/go-plotly/graph_objects"
    "github.com/MetalBlueberry/go-plotly/offline"
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

Each trace type has its own file on **graph_objecs (grob)** package. The file contains the main structure and all the needed nested objects. Files ending with **_gen** are automatically generated files running `go generate`. This is executing the code in **generator** package to generate the structures from the plotly schema. The types are documented, but you can find more examples and extended documentation at [Plotly's documentation](https://plotly.com/python/).

The values that can hold single values or arrays are defined as custom types that are a type definition of `interfaces{}`. Most common case are X and Y values. You can pass any number slice and it will work (`[]float64`,`[]int`,`[]int64`...). In case of Hovertext, you can provide a `[]string` to display a text for each point, a `string` to display the same for all or `[]int` to display a number.

Nested Properties, are defined as new types. This is great for auto completion using vscode because you can write all the boilerplate with ctrl+space. For example, the field `Title.Text` is accessed by the property `Title` of type {{Type}}Title that contains the property `Text`. The Type is always the struct that contains the field. For Layout It is `LayoutTitle`.

Flaglist and Enumerated fields are defined with a type and its constant values. This provides useful autocompletion. Keep in mind that you can compose multiple values for a Flaglist like `Mode: grob.ScatterModeMarkers + "+" + grob.ScatterModeLines,`. You can read de inline documentation to see the default value.

## Tested

Examples, Code generation and Offline package are based on version 1.58.4, but It should work with other versions as this library just generates standard JSON.

## Testing

The package lacks of unit testing basically because it's just building JSON to be consumed by plotly.js. This means that I do not see a clear way of building valuable unit testing that doesn't involve the usage of plotly.js. If the package compiles, It means that types are generated correctly.

Said that, I'm thinking about adding some integration testing with Docker, but for now, I've enough if the examples are working.

If you have any good idea of how to test this code, I will be happy to hear it.

## Progress

The main focus is to have the structures to hold the data and provide auto competition. Once we get there, we will be on v1.0.0.

Currently, most common use cases should be already covered. Feel free to create an issue if you find something that it's not working.

## FAQ

### What's the meaning of "grob"?

In python, the component is called graph_objects, like in this package, but that's too long to write it over and over again. In Python, usually it's called "go" from Graph_Objects but in Go... that's not possible for a conflict with a keyword. as an alternative, I've decided to use "grob" from GRaph_OBjects.

### How are the graph_object files generated?

I was using "[plate](https://github.com/MetalBlueberry/plate)", but it was a bad idea. Now it is just plan go code inside the **generator package**. This should be much easier to understand and to contribute. Let me know if you want to contribute!!

### What are the usecases?

1. Send plotly figures to the frontend ready to be drawn, avoiding possible mistakes in JS thanks to types!

2. Generate an awesome dynamic plot in a local file with offline package.

3. I don't know, just do something awesome and let me know it.

### Why are the String and Bool types defined?

> I'm thinking about defining everything as a pointer, it should be better in the long run

This is to handle the omitempty flag in json serialization. Turns out that if the flag is set, you cannot create a json object with a flag set to `false`. For example, turn off visibility of the legend will be impossible without this.

For bool, the solution is as simple as defining the types again inside graph_objects and it feels like using normal bool values.

For strings... This is a little bit more complicated, In AWS package they are using `aws.String` which maps to `*string` to workaround this issue, but I find that really annoying because you have to wrap every single string with `aws.String("whatever")`. For now I've decided to define the type String but leave it as `interface{}` instead of `*string` to allow you to use raw strings. The draw back is that you can pass any value of your choice... Hopefully you can live with this :).

For numbers... It's similar to strings, Right now you cannot create plots with integer/float numbers with 0 value. I've only encounter problems when trying to remove the margin and can be workaround with an small value like `0.001`. I would like to avoid using interface{} or defining types again to keep the package interface as simple as possible.
