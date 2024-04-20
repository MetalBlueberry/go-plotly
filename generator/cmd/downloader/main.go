package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Tag represents a git tag content
type Tag struct {
	Name       string `json:"name"`
	ZipballURL string `json:"zipball_url"`
	TarballURL string `json:"tarball_url"`
	Commit     struct {
		SHA string `json:"sha"`
		URL string `json:"url"`
	} `json:"commit"`
	NodeID string `json:"node_id"`
}

// getSchemaUrl returns the url to 'plot-schema.json' for the given version saved in the plotly.js repo
func getSchemaUrl(version string) string {
	return fmt.Sprintf("https://raw.githubusercontent.com/plotly/plotly.js/%s/test/plot-schema.json", version)
}

// downloadSchema gets the schema file for a given version, and either saves the raw content,
// or prepares the file for the generator if prep is set to true
func downloadSchema(version string, prep bool) error {
	schemaURL := getSchemaUrl(version)

	// Make GET request to download schema
	response, err := http.Get(schemaURL)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("could not schema file for version: %s\n[Status Code: %d (%s)]\nused url: %s\nlist the available tags first and recheck", version, response.StatusCode, response.Status, schemaURL)
	}
	defer response.Body.Close()

	// Read and decode JSON content
	var jsonData interface{}
	err = json.NewDecoder(response.Body).Decode(&jsonData)
	if err != nil {
		return err
	}

	// Re-encode JSON content with minimized whitespace and no new lines
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetIndent("", "")
	err = encoder.Encode(jsonData)
	if err != nil {
		return err
	}

	// Remove newline character from the end of the encoded JSON
	buf.Truncate(buf.Len() - 1)

	// Create file to save the schema
	var filename string
	if prep {
		filename = fmt.Sprintf("plotly-schema-%s_ready_for_gen.json", version)
	} else {
		filename = fmt.Sprintf("plotly-schema-%s_raw.json", version)
	}

	_, err = os.Stat(filename)
	if err == nil {
		fmt.Printf("WARN: overwriting: %s\n", filename)
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Copy response body to file
	fmt.Printf("Downloading %s -> %s [adjusted for generator: %t]\n", schemaURL, filename, prep)
	if prep {
		_, err = file.Write([]byte(`{"sha1":"11b662302a42aa0698df091a9974ac8f6e1a2292","modified":true,"schema":`))
		if err != nil {
			return err
		}
	}
	// Write JSON content to file
	_, err = io.Copy(file, &buf)
	if err != nil {
		return err
	}
	if prep {
		_, err = file.Write([]byte(`}`))
		if err != nil {
			return err
		}
	}

	fmt.Println("Schema downloaded and saved as", filename)
	return nil
}

// getTags lists the git tags of the plotly repo, for which a schema file can be downloaded and go files can be generated
func getTags() ([]Tag, error) {
	const url = "https://api.github.com/repos/plotly/plotly.js/tags"
	// Make GET request
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making GET request. %w", err)
	}
	defer response.Body.Close()

	// Decode JSON response
	var tags []Tag
	err = json.NewDecoder(response.Body).Decode(&tags)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON. %w", err)
	}
	return tags, nil
}

// printTags prints the name of given tags to std out.
func printTags(tags []Tag) {
	// Print names of each entry
	fmt.Println("Tags:")
	var tagNames []string
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}
	fmt.Println(strings.Join(tagNames, `,`))
}

func printExamples() {
	fmt.Println(`
Run with current working directory set to the project root.

Download 'v2.31.1' and Prepare it for use by the generator script:
go run generator/cmd/downloader/main.go -version=v2.31.1 -p

Check the url used to download 'v2.31.1':
go run generator/cmd/downloader/main.go -version=v2.31.1 -s

List available versions from the git tags:
go run generator/cmd/downloader/main.go -l`,
	)
}

func main() {
	// Define a flag for the version
	flag.Usage()
	versionPtr := flag.String("version", "v2.29.0", "version of Plotly.js")
	listPtr := flag.Bool("l", false, "list available versions")
	showPtr := flag.Bool("s", false, "only show schema url based on version - no download")
	prepPtr := flag.Bool("p", false, "prepare 'schema.json' used by generator to great *_gen.go files.")
	examplePtr := flag.Bool("u", false, "display usage examples.")
	// Parse the command-line flags
	flag.Parse()

	if *examplePtr {
		printExamples()
		return
	}

	// if list flag is passed, only parse tags
	if *listPtr {
		tags, err := getTags()
		if err != nil {
			fmt.Println("error fetching versions:", err)
			return
		}
		printTags(tags)
		return
	}

	// Check if version flag is provided
	if *versionPtr == "" {
		fmt.Println("Please provide the version using -version flag.")
		return
	}

	// If show is true, only the url to the schema has to be displayed, otherwise the content is downloaded
	if *showPtr {
		fmt.Println(getSchemaUrl(*versionPtr))
	} else {
		err := downloadSchema(*versionPtr, *prepPtr)
		if err != nil {
			fmt.Printf("error downloading the schema. %s", err)
			return
		}
	}

}
