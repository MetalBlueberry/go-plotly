//go:build ignore

package main

import (
	"encoding/json"
	"syscall/js"

	grob "github.com/MetalBlueberry/go-plotly/generated/v2.31.1/graph_objects"
)

// to run this, wasm need to be set: GOOS=js GOARCH=wasm go build -o main.wasm

func plot(this js.Value, inputs []js.Value) interface{} {
	fig := &grob.Fig{
		Data: grob.Traces{
			&grob.Choropleth{
				Autocolorscale: grob.True,
				Locationmode:   grob.ChoroplethLocationmodeUSAStates,
			},
		},
		Layout: &grob.Layout{
			Title: &grob.LayoutTitle{
				Text: "Demo",
			},
			Geo: &grob.LayoutGeo{
				Scope: grob.LayoutGeoScopeUsa,
			},
		},
	}
	b, err := json.Marshal(fig)
	if err != nil {
		return nil
	}

	plot := js.Global().Get("JSON").Call("parse", string(b))
	plotly := js.Global().Get("Plotly")
	plotly.Call("plot", "plot", plot)
	return nil
}

var c chan bool

func init() {
	c = make(chan bool)
}

func stop(this js.Value, inputs []js.Value) interface{} {
	c <- true
	return nil
}

func main() {
	js.Global().Set("plot", js.FuncOf(plot))
	js.Global().Set("stop", js.FuncOf(stop))
	<-c
	println("We are out of here")
}
