package main

import (
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/MetalBlueberry/go-plotly/generator"
)

type Creator struct{}

func (c Creator) Create(name string) (io.WriteCloser, error) {
	abs, err := filepath.Abs(name)
	if err != nil {
		return nil, err
	}

	return os.Create(abs)
}

//go:generate go run main.go --clean --schema ../../schema.json --output-directory ../../../graph_objects

func main() {
	clean := flag.Bool("clean", false, "clean the output directory first. Mandatory on CI")
	schema := flag.String("schema", "schema.json", "plotly schema")
	outputDirectory := flag.String("output-directory", "gen/", "output directory, must exist before generation")

	flag.Parse()

	file, err := os.Open(*schema)
	if err != nil {
		log.Fatalf("unable to open schema, %s", err)
	}

	root, err := generator.LoadSchema(file)
	if err != nil {
		log.Fatalf("unable to load schema, %s", err)
	}

	r, err := generator.NewRenderer(Creator{}, root)
	if err != nil {
		log.Fatalf("unable to create a new renderer, %s", err)
	}

	output := *outputDirectory

	if *clean {
		err = os.RemoveAll(output)
		if err != nil {
			log.Fatalf("Failed to clean output directory, %s", err)
		}
	}

	err = os.MkdirAll(output, 0755)
	if err != nil {
		log.Fatalf("Failed to create output dir %s, %s", *outputDirectory, err)

	}

	err = r.CreatePlotly(output)
	if err != nil {
		log.Fatalf("unable to write plotly, %s", err)
	}

	err = r.CreateTraces(output)
	if err != nil {
		log.Fatalf("unable to write traces, %s", err)
	}

	err = r.CreateLayout(output)
	if err != nil {
		log.Fatalf("unable to write layout, %s", err)
	}

	err = r.CreateConfig(output)
	if err != nil {
		log.Fatalf("unable to write config, %s", err)
	}

	err = r.CreateUnmarshal(output)
	if err != nil {
		log.Fatalf("unable to write unmarshal, %s", err)
	}
}
