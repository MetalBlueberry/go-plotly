package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	grob "github.com/MetalBlueberry/go-plotly/graph_objects"
)

func main() {
	/*
	   fig = dict({
	       "data": [{"type": "bar",
	                 "x": [1, 2, 3],
	                 "y": [1, 3, 2]}],
	       "layout": {"title": {"text": "A Figure Specified By Python Dictionary"}}
	   })
	*/
	fig := &grob.Fig{
		Data: grob.Traces{
			&grob.Bar{
				Type: grob.TraceTypeBar,
				X:    []float64{1, 2, 3},
				Y:    []float64{1, 2, 3},
			},
		},
		Layout: &grob.Layout{
			Title: &grob.LayoutTitle{
				Text: "A Figure Specified By Go Struct",
			},
		},
	}

	err := savePNG(fig, "out.png")
	if err != nil {
		panic(err)
	}
}

func savePNG(fig *grob.Fig, path string) error {
	args := "run --rm -i quay.io/plotly/orca graph --format png"
	cmd := exec.Command("docker", strings.Split(args, " ")...)

	in, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
		return fmt.Errorf("Failed to open StdIn, %w", err)
	}
	go func() {
		err := json.NewEncoder(in).Encode(fig)
		if err != nil {
			panic(err)
		}
		err = in.Close()
	}()

	go func() {
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}()

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Cannot create output file")
	}
	defer f.Close()

	out, err := cmd.StdoutPipe()
	_, err = io.Copy(f, out)
	if err != nil {
		return fmt.Errorf("Failed to open StdOut, %w", err)
	}
	return nil
}
