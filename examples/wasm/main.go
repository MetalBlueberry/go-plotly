package main

import (
	"encoding/json"
	"syscall/js"

	grob "github.com/MetalBlueberry/go-plotly/graph_objects"
)

func getPlot(this js.Value, inputs []js.Value) interface{} {
	fig := &grob.Fig{
		Data: grob.Traces{
			&grob.Choropleth{
				Type:           grob.TraceTypeChoropleth,
				Autocolorscale: grob.True,
				Locationmode:   grob.ChoroplethLocationmode_USA_states,
			},
		},
		Layout: &grob.Layout{
			Title: &grob.LayoutTitle{
				Text: "Demo",
			},
			Geo: &grob.LayoutGeo{
				Scope: grob.LayoutGeoScope_usa,
			},
		},
	}
	b, err := json.Marshal(fig)
	if err != nil {
		return nil
	}
	return string(b)
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
	js.Global().Set("getPlot", js.FuncOf(getPlot))
	js.Global().Set("stop", js.FuncOf(stop))
	<-c
	println("We are out of here")
}
