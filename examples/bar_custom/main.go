package main

import (
	"image/color"
	"strconv"

	"github.com/lucasb-eyer/go-colorful"

	grob "github.com/MetalBlueberry/go-plotly/generated/v2.31.1/graph_objects"
	"github.com/MetalBlueberry/go-plotly/pkg/offline"
	"github.com/MetalBlueberry/go-plotly/pkg/types"
)

func main() {
	/*
		https://plotly.com/javascript/bar-charts/
		var xValue = ['Product A', 'Product B', 'Product C'];

		var yValue = [20, 14, 23];

		var trace1 = {
		x: xValue,
		y: yValue,
		type: 'bar',
		text: yValue.map(String),
		textposition: 'auto',
		hoverinfo: 'none',
		marker: {
			color: 'rgb(158,202,225)',
			opacity: 0.6,
			line: {
			color: 'rgb(8,48,107)',
			width: 1.5
			}
		}
		};

		var data = [trace1];

		var layout = {
		title: 'January 2013 Sales Report',
		barmode: 'stack'
		};

	*/
	xValue := []string{"Product A", "Product B", "Product C"}
	yValue := []int{20, 14, 23}

	markerColor, ok := colorful.MakeColor(color.NRGBA{158, 202, 225, 1})
	if !ok {
		panic("fail to build a color")
	}

	trace1 := &grob.Bar{
		X:            types.DataArray(xValue),
		Y:            types.DataArray(yValue),
		Text:         types.ArrayOKArray(toString(yValue)...),
		Textposition: types.ArrayOKValue(grob.BarTextpositionAuto),
		Hoverinfo:    types.ArrayOKValue(grob.BarHoverinfoNone),
		Marker: &grob.BarMarker{
			Color: types.ArrayOKValue(types.UseColor(types.Color(
				markerColor.Hex(), // Use colorfull
			))),
			Opacity: types.ArrayOKValue(0.6),
			Line: &grob.BarMarkerLine{
				Color: types.ArrayOKValue(types.UseColor("rgb(8,48,107)")), // Or just write the string
				Width: types.ArrayOKValue(1.5),
			},
		},
	}

	layout := &grob.Layout{
		Title: &grob.LayoutTitle{
			Text: "A Figure Specified By Go Struct",
		},
	}

	fig := &grob.Fig{
		Data:   []types.Trace{trace1},
		Layout: layout,
	}

	offline.ToHtml(fig, "bar_custom.html")
	offline.Show(fig)
}

func toString(in []int) []types.String {
	out := make([]types.String, len(in))
	for i := range in {
		out[i] = strconv.Itoa(in[i])
	}
	return out
}
