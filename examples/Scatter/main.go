package main

import (
	"fmt"
	"math/rand"
	"time"

	plotly "github.com/metalblueberry/go-plotly/offline"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	// Create base figure
	fig := plotly.NewFigure()

	// Create layout configuration for the figure
	layout := fig.NewLayout()
	layout.SetTitle("Test Plot")

	// Create a new Scatter
	trace1, _ := fig.NewScatter()
	// set Mode
	trace1.SetMode(plotly.MODEMARKERS)
	trace1.SetName("10 Random points")

	// Add 10 random points 1 by 1
	for x := 0.0; x < 10.0; x++ {
		if err := trace1.AddPoints(rand.Float64()+x, rand.Float64()+x); err != nil {
			panic(err)
		}
	}

	// Create a new Scatter
	trace2, _ := fig.NewScatter()
	trace2.SetName("Array")
	// Add points as array
	if err := trace2.AddPoints([]int{1, 5, 10}, []int{1, 5, 10}); err != nil {
		panic(err)
	}
	// set Mode to use Markers and Lines
	trace2.SetMode(plotly.MODEMARKERS + plotly.MODELINES)
	marker2 := trace2.NewMarker()
	marker2.SetColor("blue")
	marker2.SetSize(50)

	// Generate a json that can be passed directly to plotly.js
	// json, err := fig.GetJSON()

	// Generate a PlotlyScriptSnippet
	snippet, err := fig.GetPlotlySnippet("plot")
	if err != nil {
		panic(err)
	}

	// Add the snippet to a simple web template
	const tpl = `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<script src="https://cdn.plot.ly/plotly-latest.min.js"></script>
			<title>%s</title>
		</head>
		<body>
		<div id="plot"></div>
		<script>
		  <!-- JAVASCRIPT CODE GOES HERE -->
			%s
		</script>
		</body>
	</html>
	`
	// Output to stdout
	fmt.Printf(tpl, "Sample Offline Plot", snippet)

}
