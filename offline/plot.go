package offline

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"

	grob "github.com/MetalBlueberry/go-plotly/graph_objects"
	"github.com/pkg/browser"
)

type Options struct {
	Addr string
}

// ToHtml saves the figure as standalone HTML. It still requires internet to load plotly.js from CDN.
func ToHtml(fig *grob.Fig, path string) {
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
	ioutil.WriteFile(path, buf.Bytes(), os.ModePerm)
}

// Show displays the figure in your browser.
// Use serve if you want a persistent view
func Show(fig *grob.Fig) {
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
func Serve(fig *grob.Fig, opt ...Options) {
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

var baseHtml = `
	<head>
		<script src="https://cdn.plot.ly/plotly-1.54.0.min.js"></script>
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
