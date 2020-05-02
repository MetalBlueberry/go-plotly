package offline

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/MetalBlueberry/go-plotly/plotly"
	"github.com/pkg/browser"
)

type Options struct {
	Addr string
}

// Show displays the figure in your browser.
// Use serve if you want a persistent view
func Show(fig *plotly.Fig) {
	figBytes, err := json.Marshal(fig)
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("plotly").Parse(baseHtml)
	if err != nil {
		panic(err)
	}
	buf := &bytes.Buffer{}
	tmpl.Execute(buf, string(figBytes))
	browser.OpenReader(buf)
}

// Serve creates a local web server that displays the image using plotly.js
func Serve(fig *plotly.Fig, opt ...Options) {
	opts := computeOptions(Options{
		Addr: "localhost:8080",
	}, opt...)

	mux := &http.ServeMux{}
	srv := &http.Server{
		Handler: mux,
		Addr:    opts.Addr,
	}
	mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Served data")
		encoder := json.NewEncoder(w)
		err := encoder.Encode(fig)
		if err != nil {
			log.Printf("Error! %s", err)
		}
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Served page")
		fmt.Fprint(w, baseHtml)
	})

	log.Print("Starting server")
	err := srv.ListenAndServe()
	log.Print(err)
	log.Print("Plot served")
}

func computeOptions(def Options, opt ...Options) Options {
	if len(opt) == 1 {
		opts := opt[0]
		if opts.Addr != "" {
			def.Addr = opts.Addr
		}
	}
	return def
}

// Serve allows you to see the figure in the browser at http://localhost:8080 by default

var baseHtml = `
	<head>
		<script src="https://cdn.plot.ly/plotly-1.33.0.min.js"></script>
	</head>
	</body>
		<div id="tester"></div>
	<script>
		console.log("start")
		data = JSON.parse('{{ . }}')

		TESTER = document.getElementById('tester');
		Plotly.newPlot(TESTER, data);
	</script>
	<body>
	`
