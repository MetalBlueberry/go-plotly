package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Schemas struct {
	Versions []Version `yaml:"versions"`
}

type Version struct {
	Name      string `yaml:"Name"`
	Tag       string `yaml:"Tag"`
	URL       string `yaml:"URL"`
	Path      string `yaml:"Path"`
	Generated string `yaml:"Generated"`
}

// downloadSchema gets the schema file for a given version, and either saves the raw content,
// or prepares the file for the generator if prep is set to true
func downloadSchema(version string, schemaURL string, prep bool) (string, error) {
	//schemaURL := getSchemaUrl(version)

	// Make GET request to download schema
	response, err := http.Get(schemaURL)
	if err != nil {
		return "", err
	}
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("could not schema file for version: %s\n[Status Code: %d (%s)]\nused url: %s\nlist the available tags first and recheck", version, response.StatusCode, response.Status, schemaURL)
	}
	defer response.Body.Close()

	// Read and decode JSON content
	var jsonData interface{}
	err = json.NewDecoder(response.Body).Decode(&jsonData)
	if err != nil {
		return "", fmt.Errorf("could not decode json content of response from: %s", schemaURL)
	}

	// Re-encode JSON content with minimized whitespace and no new lines
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetIndent("", "")
	err = encoder.Encode(jsonData)
	if err != nil {
		return "", fmt.Errorf("could not remove whitespaces and new lines from json by encoding content of response from: %s", schemaURL)
	}

	// Remove newline character from the end of the encoded JSON
	buf.Truncate(buf.Len() - 1)

	// Create file to save the schema
	var fullSavepath string

	if prep {
		fullSavepath = filepath.Join("schemas", version, "plotly-schema_for_gen.json")
	} else {
		fullSavepath = filepath.Join("schemas", version, "plotly-schema.json")
	}

	err = os.MkdirAll(filepath.Dir(fullSavepath), 0755)
	if err != nil {
		return "", fmt.Errorf("could not create directory for plotly schema: %s", filepath.Dir(fullSavepath))
	}

	_, err = os.Stat(fullSavepath)
	if err == nil {
		fmt.Printf("WARN: overwriting: %s\n", fullSavepath)
	}
	file, err := os.Create(fullSavepath)
	if err != nil {
		return "", fmt.Errorf("could not create file for saving: %s", fullSavepath)
	}
	defer file.Close()

	// Copy response body to file
	fmt.Printf("Downloading %s -> %s [adjusted for generator: %t]\n", schemaURL, fullSavepath, prep)
	if prep {
		_, err = file.Write([]byte(`{"sha1":"11b662302a42aa0698df091a9974ac8f6e1a2292","modified":true,"schema":`))
		if err != nil {
			return "", fmt.Errorf("could not write json schema prepared for code generator: %s", fullSavepath)
		}
	}
	// Write JSON content to file
	_, err = io.Copy(file, &buf)
	if err != nil {
		return "", err
	}
	if prep {
		_, err = file.Write([]byte(`}`))
		if err != nil {
			return "", err
		}
	}

	fmt.Println("Schema downloaded and saved as", fullSavepath)
	return fullSavepath, nil
}

func readSchemas(schemapath string) []Version {
	var err error
	f, err := os.Open(schemapath)
	if err != nil {
		fmt.Printf("error opening schema yaml file: %s", err)
		return nil
	}

	var schemas Schemas
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&schemas)
	if err != nil {
		fmt.Printf("error decoding schema yaml file: %s", err)
		return nil
	}

	return schemas.Versions
}

// schemaPath location of the schemas file, holding all version to be created
const schemaPath = "schemas.yaml"

func main() {
	// Define a flag for the version
	versions := readSchemas(schemaPath)
	if versions == nil {
		fmt.Printf("could not find versions")
		return
	}
	for _, version := range versions {
		_, err := downloadSchema(version.Tag, version.URL, true)
		if err != nil {
			fmt.Printf("failed to download version: %s", version.Tag)
			return
		}
	}
}
