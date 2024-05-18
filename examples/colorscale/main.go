package main

import (
	"encoding/json"
	"math"

	grob "github.com/MetalBlueberry/go-plotly/generated/v2.31.1/graph_objects"
	"github.com/MetalBlueberry/go-plotly/generated/v2.31.1/offline"
)

type ColorScale struct {
	Position   float64
	ColorScale string
}

func (cs ColorScale) MarshalJSON() ([]byte, error) {
	data := []interface{}{cs.Position, cs.ColorScale}
	return json.Marshal(data)
}

func main() {
	colorScale := []ColorScale{
		{0.0, "#6e40aa"},
		{0.1, "#963db3"},
		{0.2, "#bf3caf"},
		{0.3, "#e4419d"},
		{0.4, "#fe4b83"},
		{0.5, "#ff5e63"},
		{0.6, "#ff7847"},
		{0.7, "#fb9633"},
		{0.8, "#e2b72f"},
		{0.9, "#c6d63c"},
		{1.0, "#aff05b"},
	}

	t := linspace(0, 10, 50)
	x := cos(t)
	y := sin(t)
	z := t
	fig := &grob.Fig{
		Data: grob.Traces{
			&grob.Scatter3d{
				Type: grob.TraceTypeScatter3d,
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
					Color:          z,
					Colorscale:     colorScale,
					Showscale:      grob.True,
					Size:           4,
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
