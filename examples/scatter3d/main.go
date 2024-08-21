package main

import (
	"math"

	grob "github.com/MetalBlueberry/go-plotly/generated/v2.31.1/graph_objects"
	"github.com/MetalBlueberry/go-plotly/pkg/offline"
)

func main() {
	/*
		https://plotly.com/python/3d-scatter-plots/
		import plotly.graph_objects as go
		import numpy as np

		# Helix equation
		t = np.linspace(0, 10, 50)
		x, y, z = np.cos(t), np.sin(t), t

		fig = go.Figure(data=[go.Scatter3d(x=x, y=y, z=z,
										mode='markers')])
		fig.show()
	*/
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
			},
		},
	}

	offline.ToHtml(fig, "scatter3d.html")
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
