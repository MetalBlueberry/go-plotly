package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/MetalBlueberry/go-plotly/generator"
)

type Creator struct{}

func (c Creator) Create(name string) (io.WriteCloser, error) {
	return os.Create(name)
}

func main() {
	schema := flag.String("schema", "schema.json", "plotly schema")
	outputDirectory := flag.String("output-directory", "gen/", "output directory, must exist before generation")
	*schema = "plotly-schema_v2.29.0.json"
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
		panic(err)
	}

	output := *outputDirectory

	if err = os.MkdirAll(output, 0755); err != nil {
		fmt.Println("Error creating directory:", err)
		log.Fatalf("unable to write traces, %w", err)
	}

	err = r.CreateTraces(output)
	if err != nil {
		log.Fatalf("unable to write traces, %w", err)
	}

	err = r.CreateLayout(output)
	if err != nil {
		log.Fatalf("unable to write layout, %w", err)
	}

	err = r.CreateConfig(output)
	if err != nil {
		log.Fatalf("unable to write config, %w", err)
	}

	err = r.CreateUnmarshal(output)
	if err != nil {
		log.Fatalf("unable to write unmarshal, %w", err)
	}
}
