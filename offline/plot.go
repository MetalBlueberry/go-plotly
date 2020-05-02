package offline

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"time"

	"github.com/MetalBlueberry/go-plotly/plotly"
)

type ShowOptions struct {
	Addr string
}

// Show displays the figure in your browser.
// It works by creating a local server and closing it after serving the data.
func Show(fig *plotly.Fig, opt ...ShowOptions) {
	html := `
	<head>
		<script src="https://cdn.plot.ly/plotly-latest.min.js"></script>
	</head>
	</body>
		<div id="tester"></div>
	<script>
		console.log("start")
		fetch("/data").then((response)=>{
			console.log(response)
			return response.json()
		}).then((data)=>{
			console.log(data)
			TESTER = document.getElementById('tester');
			Plotly.newPlot(TESTER, data);
		})
	</script>
	<body>
	`
	mux := &http.ServeMux{}
	opts := ShowOptions{
		Addr: "localhost:8080",
	}
	if len(opt) > 0 {
		opts = opt[0]
	}
	srv := &http.Server{
		Handler: mux,
		Addr:    opts.Addr,
	}
	done := make(chan bool)
	mux.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Served data")
		encoder := json.NewEncoder(w)
		err := encoder.Encode(fig)
		if err != nil {
			log.Print("Error! %s", err)
		}
		log.Print("Closing")
		go func() {
			ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
			srv.Shutdown(ctx)
			close(done)
		}()
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Served page")
		fmt.Fprint(w, html)
	})

	log.Print("Starting server")
	openbrowser("http://" + opts.Addr)
	err := srv.ListenAndServe()
	log.Print(err)
	<-done
	log.Print("Plot served")
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
