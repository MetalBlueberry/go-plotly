package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MetalBlueberry/go-plotly/generator"
)

// schemaPath location of the schemas file, holding all version to be created
const configPath = "schemas.yaml"

func main() {
	var schemapath string
	flag.StringVar(&schemapath, "config", configPath, "yaml file defining versions to be generated")
	flag.Parse()

	// Define a flag for the version
	versions := generator.ReadSchemas(schemapath)
	if versions == nil {
		fmt.Printf("could not find versions\n")
		return
	}
	var err error
	for _, version := range versions {

		// check if version already downloaded:
		_, err = os.Stat(version.Path)
		if err == nil {
			fmt.Printf("already downloaded version: %s\n", version.Tag)
			continue
		}

		// download schema
		_, err = generator.DownloadSchema(version.Tag, version.URL, version.Path)
		if err != nil {
			fmt.Printf("failed to download version: %s\n", version.Tag)
			return
		}
	}
}
