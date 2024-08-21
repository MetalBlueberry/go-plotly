package main

import (
	"math"

	grob "github.com/MetalBlueberry/go-plotly/generated/v2.31.1/graph_objects"
	"github.com/MetalBlueberry/go-plotly/generated/v2.31.1/offline"
	"github.com/MetalBlueberry/go-plotly/types/arrayok"
)

func main() {
	t := linspace(0, 10, 50)
	x := cos(t)
	y := sin(t)
	z := t
	fig := &grob.Fig{
		Data: grob.Traces{
			&grob.Scatter3d{
				X:    x,
				Y:    y,
				Z:    z,
				Mode: grob.Scatter3dModeMarkers,
				Marker: &grob.Scatter3dMarker{
					Autocolorscale: grob.False,
					Cauto:          grob.False,
					Cmin:           0,
					Cmid:           5,
					Cmax:           10,
					Color:          arrayok.Array(grob.UseColorScaleValues(z)...),
					Colorscale: grob.ColorScale{
						Values: []grob.ColorScaleReference{
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
					Showscale: grob.True,
					Size:      arrayok.Value(4.0),
				},
			},
		},
		Layout: &grob.Layout{
			Height: 700,
			Width:  1200,
			Title: &grob.LayoutTitle{
				Text: "3D Spiral",
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
