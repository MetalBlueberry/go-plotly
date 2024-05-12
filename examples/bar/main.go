package main

import (
	"fmt"
	"path/filepath"

	grob "github.com/MetalBlueberry/go-plotly/generated/v2.29.1/graph_objects"
)

func main() {
	/*
	   fig = dict({
	       "data": [{"type": "bar",
	                 "x": [1, 2, 3],
	                 "y": [1, 3, 2]}],
	       "layout": {"title": {"text": "A Figure Specified By Python Dictionary"}}
	   })
	*/
	fig := &grob.Fig{
		Data: grob.Traces{
			&grob.Bar{
				Type: grob.TraceTypeBar,
				X:    []float64{1, 2, 3},
				Y:    []float64{1, 2, 3},
			},
		},
		Layout: &grob.Layout{
			Title: &grob.LayoutTitle{
				Text: "A Figure Specified By Go Struct",
			},
		},
	}

	// by default, using the cdn reference, downloading the plotly js on demand
	grob.ToHtml(fig, "bar.html")
	grob.Show(fig)

	// example for using static assets, which can be embedded into the golang application
	abs, err := filepath.Abs("asset/plotly-2.29.1.min.js")
	if err != nil {
		return
	}
	grob.ToHtml(fig, "bar.html", grob.FigOptions{HeadContent: fmt.Sprintf(`<title>Offline Bars</title><script src="%s"></script>`, abs)})
	grob.Show(fig, grob.FigOptions{HeadContent: fmt.Sprintf(`<title>Offline Bars</title><script src="%s"></script>`, abs)})
}
