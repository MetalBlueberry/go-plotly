package main

import (
	"github.com/MetalBlueberry/go-plotly/offline"
	"github.com/MetalBlueberry/go-plotly/plotly"
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
	fig := &plotly.Fig{
		Data: plotly.Traces{
			&plotly.Bar{
				Type: plotly.TraceTypeBar,
				X:    []float64{1, 2, 3},
				Y:    []float64{1, 2, 3},
			},
		},
		Layout: plotly.Layout{
			Title: &plotly.LayoutTitle{
				Text: "A Figure Specified By Go Struct",
			},
		},
	}

	offline.Show(fig)
}
