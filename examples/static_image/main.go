package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	grob "github.com/MetalBlueberry/go-plotly/generated/v2.31.1/graph_objects"
	"github.com/MetalBlueberry/go-plotly/pkg/types"
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
		Data: []types.Trace{
			&grob.Bar{
				X: []float64{1, 2, 3},
				Y: []float64{1, 2, 3},
			},
		},
		Layout: &grob.Layout{
			Title: &grob.LayoutTitle{
				Text: "A Figure Specified By Go Struct",
			},
		},
	}

	fmt.Fprint(os.Stderr, "saving to PNG...\n")
	err := savePNG(fig, "out.png")
	if err != nil {
		panic(err)
	}
	fmt.Fprint(os.Stderr, "Done!\n")
}

func savePNG(fig *grob.Fig, path string) error {
	args := "run --rm -i quay.io/plotly/orca graph --format png"
	cmd := exec.Command("docker", strings.Split(args, " ")...)

	in, err := cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("Failed to open StdIn, %w", err)
	}
	go func() {
		err := json.NewEncoder(in).Encode(fig)
		if err != nil {
			panic(err)
		}
		err = in.Close()
	}()

	out, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("Failed to open StdOut, %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Cannot create output file")
	}
	go func() {
		defer f.Close()
		_, err = io.Copy(f, out)
		if err != nil {
			panic(err)
		}
	}()

	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("Failed to run command, %w", err)
	}

	return nil
}
