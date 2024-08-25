package main

import (
	"encoding/json"
	"math"

	grob "github.com/MetalBlueberry/go-plotly/generated/v2.31.1/graph_objects"
	"github.com/MetalBlueberry/go-plotly/pkg/offline"
	"github.com/MetalBlueberry/go-plotly/pkg/types"
)

func main() {
	t := linspace(0, 10, 100)
	y := sin(t)

	fig := &grob.Fig{
		Data: []types.Trace{
			&grob.Scatter{
				X:    types.DataArray(t),
				Y:    types.DataArray(y),
				Mode: grob.ScatterModeMarkers,
			},
		},
	}

	bytes, err := json.Marshal(fig)
	if err != nil {
		panic(err)
	}

	// Demonstrates that a json representation can be loaded
	recoveredFig := &grob.Fig{}
	json.Unmarshal(bytes, recoveredFig)
	offline.Show(recoveredFig)
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
