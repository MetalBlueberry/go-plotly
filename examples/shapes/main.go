package main

import (
	grob "github.com/MetalBlueberry/go-plotly/generated/v2.19.0/graph_objects"
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
		Layout: &grob.Layout{
			Title: &grob.LayoutTitle{
				Text: types.S("A Figure Specified By Go Struct"),
			},
			Shapes: []grob.LayoutShape{
				{
					Type: "line",
					X0:   1,
					Y0:   0,
					X1:   1,
					Y1:   2,
					Line: &grob.LayoutShapeLine{
						Color: "RoyalBlue",
						Width: types.N(3),
					},
				},
				{
					Type: "line",
					X0:   2,
					Y0:   2,
					X1:   5,
					Y1:   2,
					Line: &grob.LayoutShapeLine{
						Color: "LightSeaGreen",
						Width: types.N(4),
						Dash:  types.S(string(grob.Scatter3dLineDashDashdot)),
					},
				},
				{
					Type: "line",
					X0:   4,
					Y0:   0,
					X1:   6,
					Y1:   2,
					Line: &grob.LayoutShapeLine{
						Color: "MediumPurple",
						Width: types.N(4),
						Dash:  types.S(string(grob.Scatter3dLineDashDot)),
					},
				},
			},
		},
	}

	offline.ToHtml(fig, "bar.html")
	offline.Show(fig)
}
