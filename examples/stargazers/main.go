package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/pkg/browser"

	grob "github.com/MetalBlueberry/go-plotly/generated/v2.31.1/graph_objects"
	"github.com/MetalBlueberry/go-plotly/pkg/types/arrayok"
)

type User struct {
	Login     string `json:"login,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

type GitHubStargazers struct {
	StarredAt time.Time `json:"starred_at,omitempty"`
	User      User      `json:"user,omitempty"`
}

func main() {

	request, err := http.NewRequest("GET", "https://api.github.com/repos/MetalBlueberry/go-plotly/stargazers", nil)
	if err != nil {
		log.Fatalf("Failed to build request, %s", err)
	}
	request.Header.Set("Accept", "application/vnd.github.star+json")

	q := request.URL.Query()
	q.Set("per_page", "100")
	request.URL.RawQuery = q.Encode()

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalf("Failed to make request to github, %s", err)
	}

	defer response.Body.Close()

	buf := &bytes.Buffer{}

	starredAt := []GitHubStargazers{}
	err = json.NewDecoder(io.TeeReader(response.Body, buf)).Decode(&starredAt)
	if err != nil {
		log.Fatalf("Failed to decode response, %s\n%s", err, buf.String())
	}

	x := []string{}
	y := []int{}
	text := []string{}
	link := []interface{}{}

	for i := 0; i < len(starredAt); i++ {
		x = append(x, starredAt[i].StarredAt.Format(time.RFC3339))
		y = append(y, i+1)
		text = append(text, starredAt[i].User.Login)
		link = append(link, starredAt[i].User.AvatarURL)

	}

	fig := &grob.Fig{
		Data: grob.Traces{
			&grob.Scatter{
				X:    x,
				Y:    y,
				Text: arrayok.Array(text...),
				Meta: arrayok.Array(link...),
				Mode: grob.ScatterModeLines + "+" + grob.ScatterModeMarkers,
				Name: "Stars",
				Line: &grob.ScatterLine{
					Color: "#f0ed46",
				},
			},
		},

		Layout: &grob.Layout{
			Title: &grob.LayoutTitle{
				Text: "Metalblueberry/go-plotly Stargazers",
			},
			Legend: &grob.LayoutLegend{},
		},
	}

	db, _ := json.MarshalIndent(fig, "", " ")
	fmt.Println(string(db))

	buf = figToBuffer(fig)
	browser.OpenReader(buf)
}

var page = `
	<head>
	    <script src="https://cdn.plot.ly/plotly-2.29.0.min.js"></script>
	</head>
	</body>
		<div id="plot"></div>
		<img id="avatar" />
	<script>
		data = JSON.parse('{{ . }}')
		Plotly.newPlot('plot', data);

		myPlot = document.getElementById('plot')
		avatar = document.getElementById('avatar')
		myPlot.on('plotly_hover', function(data){
			console.log(data)
			var point = data.points[0];
		console.log(point)
		avatar.src = point.data.meta[point.pointIndex]
		});
	</script>
	<body>

`

func figToBuffer(fig *grob.Fig) *bytes.Buffer {
	figBytes, err := json.Marshal(fig)
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("plotly").Parse(page)
	if err != nil {
		panic(err)
	}
	buf := &bytes.Buffer{}
	tmpl.Execute(buf, string(figBytes))
	return buf
}
