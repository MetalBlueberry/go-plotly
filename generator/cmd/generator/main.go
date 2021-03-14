package main

import (
	"os"

	"github.com/MetalBlueberry/go-plotly/generator"
)

func main() {
	file, err := os.Open("schema.json")
	if err != nil {
		panic(err)
	}
	root, err := generator.LoadSchema(file)
	if err != nil {
		panic(err)
	}
	r, err := generator.NewRenderer("templates/*.tmpl", root)
	if err != nil {
		panic(err)
	}

	r.WriteTraces(os.Stdout)
}
