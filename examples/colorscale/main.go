package main

import (
	"math"

	grob "github.com/MetalBlueberry/go-plotly/generated/v2.31.1/graph_objects"
	"github.com/MetalBlueberry/go-plotly/pkg/offline"
	"github.com/MetalBlueberry/go-plotly/pkg/types"
)

func main() {
	t := linspace(0, 10, 50)
	x := cos(t)
	y := sin(t)
	z := t
	fig := &grob.Fig{
		Data: []types.Trace{
			&grob.Scatter3d{
				X:    types.DataArray(x),
				Y:    types.DataArray(y),
				Z:    types.DataArray(z),
				Mode: grob.Scatter3dModeMarkers,
				Marker: &grob.Scatter3dMarker{
					Autocolorscale: types.False,
					Cauto:          types.False,
					Cmin:           types.N(0),
					Cmid:           types.N(5),
					Cmax:           types.N(10),
					Color:          types.ArrayOKArray(types.UseColorScaleValues(z)...),
					Colorscale: &types.ColorScale{
						Values: []types.ColorScaleReference{
							{NormalizedValue: 0.0, Color: "#6e40aa"},
							{NormalizedValue: 0.1, Color: `#963db3`},
							{NormalizedValue: 0.2, Color: "#bf3caf"},
							{NormalizedValue: 0.3, Color: "#e4419d"},
							{NormalizedValue: 0.4, Color: "#fe4b83"},
							{NormalizedValue: 0.5, Color: "#ff5e63"},
							{NormalizedValue: 0.6, Color: "#ff7847"},
							{NormalizedValue: 0.7, Color: "#fb9633"},
							{NormalizedValue: 0.8, Color: "#e2b72f"},
							{NormalizedValue: 0.9, Color: "#c6d63c"},
							{NormalizedValue: 1.0, Color: "#aff05b"},
						},
					},
					Showscale: types.True,
					Size:      types.ArrayOKValue(types.N(4.0)),
				},
			},
		},
		Layout: &grob.Layout{
			Height: types.N(700),
			Width:  types.N(1200),
			Title: &grob.LayoutTitle{
				Text: types.S("3D Spiral"),
			},
		},
	}
	offline.Show(fig)
}

func linspace(start, stop float64, points int) []float64 {
	step := (stop - start) / float64(points)
	out := make([]float64, points)

	for i := 0; i < points; i++ {
		out[i] = start + step*float64(i)
	}
	return out
}

func sin(x []float64) []float64 {
	y := make([]float64, len(x))
	for i := 0; i < len(x); i++ {
		y[i] = math.Sin(x[i])
	}
	return y
}

func cos(x []float64) []float64 {
	z := make([]float64, len(x))
	for i := 0; i < len(x); i++ {
		z[i] = math.Cos(x[i])
	}
	return z

}
