package offline

import (
	"bytes"
	"encoding/json"
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
	buf := figToBuffer(fig)
	ioutil.WriteFile(path, buf.Bytes(), os.ModePerm)
}

// Show displays the figure in your browser.
// Use serve if you want a persistent view
func Show(fig *grob.Fig) {
	buf := figToBuffer(fig)
	browser.OpenReader(buf)
}

func figToBuffer(fig *grob.Fig) *bytes.Buffer {
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
	return buf
}

// Serve creates a local web server that displays the image using plotly.js
// Is a good alternative to Show to avoid creating tmp files.
func Serve(fig *grob.Fig, opt ...Options) {
	opts := computeOptions(Options{
		Addr: "localhost:8080",
	}, opt...)

	mux := &http.ServeMux{}
	srv := &http.Server{
		Handler: mux,
		Addr:    opts.Addr,
	}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		buf := figToBuffer(fig)
		buf.WriteTo(w)
	})

	log.Print("Starting server")
	if err := srv.ListenAndServe(); err != nil {
		log.Print(err)
	}
	log.Print("Stop server")
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
		<script src="https://cdn.plot.ly/plotly-1.58.4.min.js"></script>
	</head>
	</body>
		<div id="plot"></div>
	<script>
		data = JSON.parse('{{ . }}')
		Plotly.newPlot('plot', data);
	</script>
	<body>
	`
