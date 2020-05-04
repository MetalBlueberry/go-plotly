package main

import (
	grob "github.com/MetalBlueberry/go-plotly/graph_objects"
	"github.com/MetalBlueberry/go-plotly/offline"
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
		Layout: &grob.Layout{
			Title: &grob.LayoutTitle{
				Text: "A Figure Specified By Go Struct",
			},
			Shapes: []LineShape{
				{
					Type: "line",
					X0:   1,
					Y0:   0,
					X1:   1,
					Y1:   2,
					Line: grob.ScatterLine{
						Color: "RoyalBlue",
						Width: 3,
					},
				},
				{
					Type: "line",
					X0:   2,
					Y0:   2,
					X1:   5,
					Y1:   2,
					Line: grob.ScatterLine{
						Color: "LightSeaGreen",
						Width: 4,
						Dash:  grob.Scatter3dLineDash_dashdot,
					},
				},
				{
					Type: "line",
					X0:   4,
					Y0:   0,
					X1:   6,
					Y1:   2,
					Line: grob.ScatterLine{
						Color: "MediumPurple",
						Width: 4,
						Dash:  grob.Scatter3dLineDash_dot,
					},
				},
			},
		},
	}

	offline.ToHtml(fig, "bar.html")
	offline.Show(fig)
}

type LineShape struct {
	Type string           `json:"type,omitempty"`
	X0   float64          `json:"x0,omitempty"`
	Y0   float64          `json:"y0,omitempty"`
	X1   float64          `json:"x1,omitempty"`
	Y1   float64          `json:"y1,omitempty"`
	Line grob.ScatterLine `json:"line,omitempty"`
}
