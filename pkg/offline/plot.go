package offline

import (
	"bytes"
	"encoding/base64"
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
	tmpl, err := template.New("plotly").Parse(singleFileHTML)
	if err != nil {
		panic(err)
	}
	buf := &bytes.Buffer{}
	data := struct {
		Version    types.Version
		B64Content string
	}{
		Version: fig.Info(),
		// Encode to avoid problems with special characters
		B64Content: base64.StdEncoding.EncodeToString(figBytes),
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

	mux.HandleFunc("/content", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(fig)
		if err != nil {
			log.Printf("Error rendering template, %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/", webPage(fig, "/content"))

	log.Printf("Starting server at %s", srv.Addr)
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

var singleFileHTML = `
	<head>
		<script src="{{ .Version.Cdn }}"></script>
	</head>
	<body>
		<div id="plot"></div>
	<script>
		data = JSON.parse(atob('{{ .B64Content }}'))
		Plotly.newPlot('plot', data);
	</script>
	</body>
	`

type serverHTMLdata struct {
	Version    types.Version
	ContentURL string
}

func webPage(fig types.Fig, contentURL string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("server").Parse(serverHTML))
		data := serverHTMLdata{
			Version:    fig.Info(),
			ContentURL: contentURL,
		}
		err := tmpl.Execute(w, data)
		if err != nil {
			log.Printf("Error rendering template, %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

var serverHTML = `
	<head>
		<script src="{{ .Version.Cdn }}"></script>
	</head>
	<body>
		<div id="plot"></div>
	<script>
		const url = '{{ .ContentURL }}'
		fetch(url)
        	.then(response => response.json())  // Parse the JSON data from the response
        	.then(data => {
            	// Use the fetched data to create the Plotly plot
            	Plotly.newPlot('plot', data);
        	})
        	.catch(error => console.error('Error fetching data:', error));
	</script>
	</body>
	`
