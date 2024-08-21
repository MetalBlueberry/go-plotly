package main

import (
	grob "github.com/MetalBlueberry/go-plotly/generated/v2.31.1/graph_objects"
	"github.com/MetalBlueberry/go-plotly/generated/v2.31.1/offline"
	"github.com/MetalBlueberry/go-plotly/types/arrayok"
)

func main() {
	/*
		https://plotly.com/javascript/bar-charts/#waterfall-bar-chart

		// Base

		var xData = ['Product<br>Revenue', 'Services<br>Revenue',
		  'Total<br>Revenue', 'Fixed<br>Costs',
		  'Variable<br>Costs', 'Total<br>Costs', 'Total'
		];

		var yData = [400, 660, 660, 590, 400, 400, 340];

		var textList = ['$430K', '$260K', '$690K', '$-120K', '$-200K', '$-320K', '$370K'];

		//Base

		var trace1 = {
		  x: xData,
		  y: [0, 430, 0, 570, 370, 370, 0],
		  marker: {
		    color: 'rgba(1,1,1,0.0)'
		  },
		  type: 'bar'
		};

		//Revenue

		var trace2 = {
		  x: xData,
		  y: [430, 260, 690, 0, 0, 0, 0],
		  type: 'bar',
		  marker: {
		    color: 'rgba(55,128,191,0.7)',
		    line: {
		      color: 'rgba(55,128,191,1.0)',
		      width: 2
		    }
		  }
		};

		//Cost

		var trace3 = {
		  x: xData,
		  y: [0, 0, 0, 120, 200, 320, 0],
		  type: 'bar',
		  marker: {
		    color: 'rgba(219, 64, 82, 0.7)',
		    line: {
		      color: 'rgba(219, 64, 82, 1.0)',
		      width: 2
		    }
		  }
		};

		//Profit

		var trace4 = {
		  x: xData,
		  y: [0, 0, 0, 0, 0, 0, 370],
		  type: 'bar',
		  marker: {
		    color: 'rgba(50,171, 96, 0.7)',
		    line: {
		      color: 'rgba(50,171,96,1.0)',
		      width: 2
		    }
		  }
		};

		var data = [trace1, trace2, trace3, trace4];

		var layout = {
		  title: 'Annual Profit 2015',
		  barmode: 'stack',
		  paper_bgcolor: 'rgba(245,246,249,1)',
		  plot_bgcolor: 'rgba(245,246,249,1)',
		  width: 600,
		  height: 600,
		  showlegend: false,
		  annotations: []
		};

		for ( var i = 0 ; i < 7 ; i++ ) {
		  var result = {
		    x: xData[i],
		    y: yData[i],
		    text: textList[i],
		    font: {
		      family: 'Arial',
		      size: 14,
		      color: 'rgba(245,246,249,1)'
		    },
		    showarrow: false
		  };
		  layout.annotations.push(result);
		};

		Plotly.newPlot('myDiv', data, layout);
	*/
	xData := []string{
		"Product<br>Revenue",
		"Services<br>Revenue",
		"Total<br>Revenue",
		"Fixed<br>Costs",
		"Variable<br>Costs",
		"Total<br>Costs",
		"Total",
	}
	yData := []int{400, 660, 660, 590, 400, 400, 340}
	textList := []string{"$430K", "$260K", "$690K", "$-120K", "$-200K", "$-320K", "$370K"}

	// Base

	trace1 := &grob.Bar{
		X: xData,
		Y: []int{0, 430, 0, 570, 370, 370, 0},
		Marker: &grob.BarMarker{
			Color: arrayok.Value(grob.UseColor("rgba(1,1,1,0.0)")),
		},
	}

	// Revenue

	trace2 := &grob.Bar{
		X: xData,
		Y: []int{430, 260, 690, 0, 0, 0, 0},
		Marker: &grob.BarMarker{
			Color: arrayok.Value(grob.UseColor("rgba(55,128,191,0.7)")),
			Line: &grob.BarMarkerLine{
				Color: arrayok.Value(grob.UseColor("rgba(55,128,191,1.0)")),
				Width: arrayok.Value(2.0),
			},
		},
	}

	// Cost

	trace3 := &grob.Bar{
		X: xData,
		Y: []int{0, 0, 0, 120, 200, 320, 0},
		Marker: &grob.BarMarker{
			Color: arrayok.Value(grob.UseColor("rgba(219, 64, 82, 0.7)")),
			Line: &grob.BarMarkerLine{
				Color: arrayok.Value(grob.UseColor("rgba(219, 64, 82, 1.0)")),
				Width: arrayok.Value(2.0),
			},
		},
	}

	// Profit

	trace4 := &grob.Bar{
		X: xData,
		Y: []int{0, 0, 0, 0, 0, 0, 370},
		Marker: &grob.BarMarker{
			Color: arrayok.Value(grob.UseColor("rgba(50,171, 96, 0.7)")),
			Line: &grob.BarMarkerLine{
				Color: arrayok.Value(grob.UseColor("rgba(50,171,96,1.0)")),
				Width: arrayok.Value(2.0),
			},
		},
	}

	data := grob.Traces{trace1, trace2, trace3, trace4}

	annotations := make([]Annotation, 7)
	for i := 0; i < len(annotations); i++ {
		annotations[i] = Annotation{
			X:    xData[i],
			Y:    yData[i],
			Text: textList[i],
			Font: AnnotationFont{
				Family: "Arial",
				Size:   14,
				Color:  "rgba(245,246,249,1)",
			},
			Showarrow: false,
		}
	}

	layout := &grob.Layout{
		Title: &grob.LayoutTitle{
			Text: "Annual Profit 2015",
		},
		Barmode:      grob.BarBarmodeStack,
		PaperBgcolor: "rgba(245,246,249,1)",
		PlotBgcolor:  "rgba(245,246,249,1)",
		Width:        600,
		Height:       600,
		Showlegend:   grob.False,
		Annotations:  annotations, // Not implemented yet
	}

	fig := &grob.Fig{
		Data:   data,
		Layout: layout,
	}

	offline.ToHtml(fig, "waterfall.html")
	offline.Show(fig)
}

type Annotation struct {
	X         string         `json:"x"`
	Y         int            `json:"y"`
	Text      string         `json:"text"`
	Font      AnnotationFont `json:"font"`
	Showarrow bool           `json:"showarrow"`
}

type AnnotationFont struct {
	Family string `json:"family"`
	Size   int64  `json:"size"`
	Color  string `json:"color"`
}
