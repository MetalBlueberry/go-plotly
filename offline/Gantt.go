package offline

type Gantt map[string]interface{}

func NewGantt() (Gantt, Layout, Config) {
	return Gantt{}, NewGanttLayout(), NewGanttConfig()
}

func NewGanttConfig() Config {
	return Config{
		"barmode":   "stack",
		"hovermode": "closest",
		"yaxis": map[string]interface{}{
			"fixedrange": true,
		},
	}
}

func NewGanttLayout() Layout {
	return Layout{
		"modeBarButtonsToRemove": []string{"hoverCompareCartesian"},
	}
}

/*
TESTER = document.getElementById("tester");

Plotly.plot(
	TESTER,
	[
		{
			type: "bar",
			orientation: "h",
			x: [1, 2, 3, 3, 5],
			base: [0, 2, 0, 0, 0],
			y: [1, 1, 4, 1, 16],
			text: ["ts1", "ts3", "ts2", "tst", "tst"],
			marker: {
				color: ["red", "red", "red", "blue", "blue", "blue"],
				opacity: 1,
				line: {
					color: "black",
					width: 4,
					opacity: 1
				}
			}
		}
	],
	{
		margin: { t: 0 },
		barmode: "stack",
		hovermode: "closest",
		xaxis: {
			range: [0, 1] // to set the xaxis range to 0 to 1
		},
		yaxis: {
			fixedrange: true
		}
	},
	{
		//displayModeBar: false
		modeBarButtonsToRemove: ["hoverCompareCartesian"]
	}
);

console.log(Plotly.BUILD);
*/
