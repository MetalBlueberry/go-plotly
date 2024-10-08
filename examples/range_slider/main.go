package main

import (
	"net/http"
	"strings"

	"github.com/go-gota/gota/dataframe"

	grob "github.com/MetalBlueberry/go-plotly/generated/v2.31.1/graph_objects"
	"github.com/MetalBlueberry/go-plotly/pkg/offline"
	"github.com/MetalBlueberry/go-plotly/pkg/types"
)

func main() {
	/*
		https://plotly.com/python/range-slider/
		import plotly.graph_objects as go

		import pandas as pd

		# Load data
		df = pd.read_csv(
		    "https://raw.githubusercontent.com/plotly/datasets/master/finance-charts-apple.csv")
		df.columns = [col.replace("AAPL.", "") for col in df.columns]

		# Create figure
		fig = go.Figure()

		fig.add_trace(
		    go.Scatter(x=list(df.Date), y=list(df.High)))

		# Set title
		fig.update_layout(
		    title_text="Time series with range slider and selectors"
		)

		# Add range slider
		fig.update_layout(
		    xaxis=dict(
		        rangeselector=dict(
		            buttons=list([
		                dict(count=1,
		                     label="1m",
		                     step="month",
		                     stepmode="backward"),
		                dict(count=6,
		                     label="6m",
		                     step="month",
		                     stepmode="backward"),
		                dict(count=1,
		                     label="YTD",
		                     step="year",
		                     stepmode="todate"),
		                dict(count=1,
		                     label="1y",
		                     step="year",
		                     stepmode="backward"),
		                dict(step="all")
		            ])
		        ),
		        rangeslider=dict(
		            visible=True
		        ),
		        type="date"
		    )
		)

		fig.show()
	*/

	request, err := http.Get("https://raw.githubusercontent.com/plotly/datasets/master/finance-charts-apple.csv")
	if err != nil {
		panic(err)
	}

	df := dataframe.ReadCSV(request.Body)
	for _, col := range df.Names() {
		df = df.Rename(strings.TrimPrefix(col, "AAPL."), col)
	}

	fig := &grob.Fig{}
	fig.Data = append(fig.Data, &grob.Scatter{
		X: types.DataArray(df.Col("Date").Records()),
		Y: types.DataArray(df.Col("High").Float()),
	})

	fig.Layout = &grob.Layout{
		Title: &grob.LayoutTitle{
			Text: types.S("Time series with range slider and selectors"),
		},
		Xaxis: &grob.LayoutXaxis{
			Rangeselector: &grob.LayoutXaxisRangeselector{
				Buttons: []grob.LayoutXaxisRangeselectorButton{
					{
						Count:    types.N(1),
						Label:    types.S("1m"),
						Step:     "month",
						Stepmode: "backward",
					},
					{
						Count:    types.N(6),
						Label:    types.S("6m"),
						Step:     "month",
						Stepmode: "backward",
					},
					{
						Count:    types.N(1),
						Label:    types.S("YTD"),
						Step:     "year",
						Stepmode: "todate",
					},
					{
						Count:    types.N(1),
						Label:    types.S("1y"),
						Step:     "year",
						Stepmode: "backward",
					},
					{
						Step: "all",
					},
				},
			},
			Rangeslider: &grob.LayoutXaxisRangeslider{
				Visible: types.True,
			},
			Type: grob.LayoutXaxisTypeDate,
		},
	}

	offline.ToHtml(fig, "range_slider.html")
	offline.Show(fig)
}
