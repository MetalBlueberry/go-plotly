package main

import (
	"math"

	grob "github.com/MetalBlueberry/go-plotly/generated/v2.31.1/graph_objects"
	"github.com/MetalBlueberry/go-plotly/generated/v2.31.1/offline"
)

func main() {
	/*
		import plotly.graph_objects as go
		import numpy as np

		N = 1000
		t = np.linspace(0, 10, 100)
		y = np.sin(t)

		fig = go.Figure(data=go.Scatter(x=t, y=y, mode='markers'))

		fig.show()
	*/
	t := linspace(0, 10, 100)
	y := sin(t)

	fig := &grob.Fig{
		Data: grob.Traces{
			&grob.Scatter{
				Type: grob.TraceTypeScatter,
				X:    t,
				Y:    y,
				Mode: grob.ScatterModeMarkers,
			},
		},
	}

	offline.ToHtml(fig, "scatter.html")
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
	for i := 1; i < len(x); i++ {
		y[i] = math.Sin(x[i])
	}
	return y
}
