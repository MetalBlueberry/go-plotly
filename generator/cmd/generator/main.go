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

// The generate command always executes with the working directory set to the path with the file with the directive

//go:generate go run main.go --config=../../../schemas.yaml

const configPath = "schemas.yaml"

func main() {
	var schemapath string
	flag.StringVar(&schemapath, "config", configPath, "yaml file defining versions to be generated")
	flag.Parse()

	// Define a flag for the version
	schemas := generator.ReadSchemas(schemapath)
	if schemas == nil {
		fmt.Printf("could not find versions\n")
		return
	}
	for _, schema := range schemas {
		generatePackage(schema.Path, schema.Generated, schema.Cdn, schema.Tag)
	}

}

type Creator struct{}

func (c Creator) Create(name string) (io.WriteCloser, error) {
	abs, err := filepath.Abs(name)
	if err != nil {
		return nil, err
	}

	return os.Create(abs)
}

// FindDir goes back in the project directory to find the correct location of the schema and output folders
// this allows the script to run smoothlessly with go generate in the cmd files,
// and also when it is run from the project root as working directory
func FindDir(schema, output string) (string, string) {
	for i := 0; i < 4; i++ {
		if _, err := os.Stat(schema); err == nil {
			break
		}
		schema = filepath.Join("../", schema)
		output = filepath.Join("../", output)
	}
	return schema, output
}

// create the packages and write them into the specified folders
func generatePackage(schema, versionOutput, cdnUrl, tag string) {
	// look for the correct schema and output paths
	schema, versionOutput = FindDir(schema, versionOutput)
	log.Println("schema:", schema, "versionoutput", versionOutput)

	file, err := os.Open(schema)
	if err != nil {
		log.Fatalf("unable to open schema, %s", err)
	}

	graphObjectsOuput := filepath.Join(versionOutput, "graph_objects")
	offlineOuput := filepath.Join(versionOutput, "offline")

	root, err := generator.LoadSchema(file)
	if err != nil {
		log.Fatalf("unable to load schema, %s", err)
	}

	r, err := generator.NewRenderer(Creator{}, root)
	if err != nil {
		log.Fatalf("unable to create a new renderer, %s", err)
	}

	for _, path := range []string{graphObjectsOuput, offlineOuput} {
		err = os.RemoveAll(path)
		if err != nil {
			log.Fatalf("Failed to clean output directory, %s", err)
		}

		if err = os.MkdirAll(path, 0755); err != nil {
			log.Fatalf("Failed to create output dir %s, %s", path, err)
		}
	}

	// plot_gen.go must be separate in an offline package
	err = r.CreatePlotGo(offlineOuput, cdnUrl, tag)
	if err != nil {
		log.Fatalf("unable to write plot.go, %s", err)
	}

	// graphobjects
	err = r.CreatePlotly(graphObjectsOuput)
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

	err = r.CreateUnmarshal(graphObjectsOuput)
	if err != nil {
		log.Fatalf("unable to write unmarshal, %s", err)
	}
}
