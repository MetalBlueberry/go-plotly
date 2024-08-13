package generator

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
	Name      string `yaml:"Name"`      // name of the version
	Tag       string `yaml:"Tag"`       // git tag of the plotly version
	URL       string `yaml:"URL"`       // url under which the plotly schema json file can be downloaded directly
	Path      string `yaml:"Path"`      // path under which the schema file will be saved locally for future use
	Generated string `yaml:"Generated"` // path for the generated package
	Cdn       string `yaml:"CDN"`       // url for the cdn which should be included in the head of the generated html
}

// cleanJson makes sure that whitespaces and new line characters are removed from the downloaded json
func cleanJson(stream io.ReadCloser) (bytes.Buffer, error) {
	var err error
	var buf bytes.Buffer

	// Read and decode JSON content
	var jsonData interface{}
	err = json.NewDecoder(stream).Decode(&jsonData)
	if err != nil {
		return buf, fmt.Errorf("could not decode json content of response")
	}

	// Re-encode JSON content with minimized whitespace and no new lines
	encoder := json.NewEncoder(&buf)
	encoder.SetIndent("", "")
	err = encoder.Encode(jsonData)
	if err != nil {
		return buf, fmt.Errorf("could not remove whitespaces and new lines from json by encoding content")
	}
	// Remove newline character from the end of the encoded JSON
	buf.Truncate(buf.Len() - 1)
	return buf, nil
}

// DownloadSchema gets the schema file for a given version, and either saves the raw content,
// or prepares the file for the generator if prep is set to true
func DownloadSchema(version string, schemaURL string, outputpath string) (string, error) {
	// Make GET request to download schema
	response, err := http.Get(schemaURL)
	if err != nil {
		return "", err
	}
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("could not schema file for version: %s\n[Status Code: %d (%s)]\nused url: %s\nlist the available tags first and recheck", version, response.StatusCode, response.Status, schemaURL)
	}
	defer response.Body.Close()

	// clean the response
	buf, err := cleanJson(response.Body)
	if err != nil {
		return "", fmt.Errorf("%s: from: %s", err, schemaURL)
	}

	// Create file to save the schema
	err = os.MkdirAll(filepath.Dir(outputpath), 0755)
	if err != nil {
		return "", fmt.Errorf("could not create directory for plotly schema: %s", filepath.Dir(outputpath))
	}

	_, err = os.Stat(outputpath)
	if err == nil {
		fmt.Printf("WARN: overwriting: %s\n", outputpath)
	}
	file, err := os.Create(outputpath)
	if err != nil {
		return "", fmt.Errorf("could not create file for saving: %s", outputpath)
	}
	defer file.Close()

	// Copy response body to file
	fmt.Printf("Downloading %s -> %s\n", schemaURL, outputpath)

	_, err = file.Write([]byte(`{"sha1":"11b662302a42aa0698df091a9974ac8f6e1a2292","modified":true,"schema":`))
	if err != nil {
		return "", fmt.Errorf("could not write json schema prepared for code generator: %s", outputpath)
	}

	// Write JSON content to file
	_, err = io.Copy(file, &buf)
	if err != nil {
		return "", err
	}

	_, err = file.Write([]byte(`}`))
	if err != nil {
		return "", err
	}

	fmt.Printf("Schema downloaded and saved as: %s\n", outputpath)
	return outputpath, nil
}

func ReadSchemas(schemapath string) []Version {
	var err error
	f, err := os.Open(schemapath)
	if err != nil {
		fmt.Printf("error opening schema yaml file: %s\n", err)
		return nil
	}

	var schemas Schemas
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&schemas)
	if err != nil {
		fmt.Printf("error decoding schema yaml file: %s\n", err)
		return nil
	}

	return schemas.Versions
}
