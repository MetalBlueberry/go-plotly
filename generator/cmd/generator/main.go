package main

import (
	"flag"
	"fmt"
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

// The generate command always executes with the working directory set to the path with the file with the directive
//go:generate go run main.go --config=../../../schemas.yaml -clean

const configPath = "schemas.yaml"

func main() {
	var schemapath string
	var clean bool

	flag.StringVar(&schemapath, "config", configPath, "yaml file defining versions to be generated")
	flag.BoolVar(&clean, "clean", false, "clean the output directory first. Mandatory on CI")

	flag.Parse()

	// Define a flag for the version
	schemas := generator.ReadSchemas(schemapath)
	if schemas == nil {
		fmt.Printf("could not find versions\n")
		return
	}

	root := filepath.Dir(schemapath)

	for _, schema := range schemas {
		generatePackage(root, schema, clean)
	}

}

func rootDirectories(root, schema, output string) (string, string) {
	schema = filepath.Join(root, schema)
	output = filepath.Join(root, output)

	return schema, output
}

// create the packages and write them into the specified folders

func generatePackage(projectRoot string, version generator.Version, clean bool) {
	// look for the correct schema and output paths
	schema, relativeVersionOutput := rootDirectories(projectRoot, version.Path, version.Generated)
	log.Println("schema:", schema, "versionoutput", relativeVersionOutput)

	file, err := os.Open(schema)
	if err != nil {
		log.Fatalf("unable to open schema, %s", err)
	}

	graphObjectsOuput := filepath.Join(relativeVersionOutput, "graph_objects")

	schemaRoot, err := generator.LoadSchema(file)
	if err != nil {
		log.Fatalf("unable to load schema, %s", err)
	}

	r, err := generator.NewRenderer(Creator{}, schemaRoot)
	if err != nil {
		log.Fatalf("unable to create a new renderer, %s", err)
	}

	if clean {
		for _, path := range []string{graphObjectsOuput} {
			err = os.RemoveAll(path)
			if err != nil {
				log.Fatalf("Failed to clean output directory, %s", err)
			}

			if err = os.MkdirAll(path, 0755); err != nil {
				log.Fatalf("Failed to create output dir %s, %s", path, err)
			}
		}
	}

	err = r.CreatePlotly(graphObjectsOuput, version)
	if err != nil {
		log.Fatalf("unable to write plotly, %s", err)
	}

	err = r.CreateTraces(graphObjectsOuput)
	if err != nil {
		log.Fatalf("unable to write traces, %s", err)
	}

	err = r.CreateLayout(graphObjectsOuput)
	if err != nil {
		log.Fatalf("unable to write layout, %s", err)
	}

	err = r.CreateConfig(graphObjectsOuput)
	if err != nil {
		log.Fatalf("unable to write config, %s", err)
	}

	err = r.CreateAnimation(graphObjectsOuput)
	if err != nil {
		log.Fatalf("unable to write layout, %s", err)
	}

	err = r.CreateFrames(graphObjectsOuput)
	if err != nil {
		log.Fatalf("unable to write layout, %s", err)
	}

	err = r.CreateUnmarshal(graphObjectsOuput)
	if err != nil {
		log.Fatalf("unable to write unmarshal, %s", err)
	}

	err = r.CreateUnmarshalTests(graphObjectsOuput)
	if err != nil {
		log.Fatalf("unable to write unmarshal, %s", err)
	}

}
