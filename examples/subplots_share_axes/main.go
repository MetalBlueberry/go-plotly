package main

import (
	grob "github.com/MetalBlueberry/go-plotly/generated/v2.31.1/graph_objects"
	"github.com/MetalBlueberry/go-plotly/generated/v2.31.1/offline"
)

func main() {
	/*
		https://plotly.com/javascript/subplots/#subplots-with-shared-axes

		var trace1 = {
		  x: [1, 2, 3],
		  y: [2, 3, 4],
		  type: 'scatter'
		};

		var trace2 = {
		  x: [20, 30, 40],
		  y: [5, 5, 5],
		  xaxis: 'x2',
		  yaxis: 'y',
		  type: 'scatter'
		};

		var trace3 = {
		  x: [2, 3, 4],
		  y: [600, 700, 800],
		  xaxis: 'x',
		  yaxis: 'y3',
		  type: 'scatter'
		};

		var trace4 = {
		  x: [4000, 5000, 6000],
		  y: [7000, 8000, 9000],
		  xaxis: 'x4',
		  yaxis: 'y4',
		  type: 'scatter'
		};

		var data = [trace1, trace2, trace3, trace4];

		var layout = {
		  grid: {
		    rows: 2,
		    columns: 2,
		    subplots:[['xy','x2y'], ['xy3','x4y4']],
		    roworder:'bottom to top'
		  }
		};

		Plotly.newPlot('myDiv', data, layout);
	*/
	fig := &grob.Fig{
		Data: grob.Traces{
			&grob.Scatter{
				Type: grob.TraceTypeScatter,
				X:    []float64{1, 2, 3},
				Y:    []float64{2, 3, 4},
			},
			&grob.Scatter{
				Type:  grob.TraceTypeScatter,
				X:     []float64{20, 30, 40},
				Y:     []float64{5, 5, 5},
				Xaxis: "x2",
				Yaxis: "y",
			},
			&grob.Scatter{
				Type:  grob.TraceTypeScatter,
				X:     []float64{2, 3, 4},
				Y:     []float64{600, 700, 800},
				Xaxis: "x",
				Yaxis: "y3",
			},
			&grob.Scatter{
				Type:  grob.TraceTypeScatter,
				X:     []float64{4000, 5000, 6000},
				Y:     []float64{7000, 8000, 9000},
				Xaxis: "x4",
				Yaxis: "y4",
			},
		},
		Layout: &grob.Layout{
			Grid: &grob.LayoutGrid{
				Rows:    2,
				Columns: 2,
				Subplots: [][]string{
					{"xy", "x2y"},
					{"xy3", "x4y4"},
				},
				Roworder: grob.LayoutGridRoworderBottomToTop,
			},
		},
	}

	offline.ToHtml(fig, "subplots_share_axes.html")
	offline.Show(fig)
}
