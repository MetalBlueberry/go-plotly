package offline

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/MetalBlueberry/go-plotly/pkg/types"
	"github.com/pkg/browser"
)

type Options struct {
	Addr string
}

// ToHtml saves the figure as standalone HTML. It still requires internet to load plotly.js from CDN.
func ToHtml(fig types.Fig, path string) {
	buf := figToBuffer(fig)
	os.WriteFile(path, buf.Bytes(), os.ModePerm)
}

// Show displays the figure in your browser.
// Use serve if you want a persistent view
func Show(fig types.Fig) {
	buf := figToBuffer(fig)
	browser.OpenReader(buf)
}

func figToBuffer(fig types.Fig) *bytes.Buffer {
	figBytes, err := json.Marshal(fig)
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New("plotly").Parse(baseHtml)
	if err != nil {
		panic(err)
	}
	buf := &bytes.Buffer{}
	data := struct {
		Version types.Version
		Content string
	}{
		Version: fig.Info(),
		Content: string(figBytes),
	}

	err = tmpl.Execute(buf, data)
	if err != nil {
		panic(err)
	}
	return buf
}

// Serve creates a local web server that displays the image using plotly.js
// Is a good alternative to Show to avoid creating tmp files.
func Serve(fig types.Fig, opt ...Options) {
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
		<script src="{{ .Version.Cdn }}"></script>
	</head>
	<body>
		<div id="plot"></div>
	<script>
		data = JSON.parse('{{ .Content }}')
		Plotly.newPlot('plot', data);
	</script>
	</body>
	`
