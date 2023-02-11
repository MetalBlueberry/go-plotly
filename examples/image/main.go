package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"

	grob "github.com/MetalBlueberry/go-plotly/graph_objects"
	"github.com/MetalBlueberry/go-plotly/offline"
)

// bytes to base64 image
func image64FromBytes(b []byte) string {
	var base64Encoding string
	mimeType := http.DetectContentType(b)

	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += base64.StdEncoding.EncodeToString(b)
	return base64Encoding
}

// load image from path
func image64(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return image64FromBytes(bytes)
}

func main() {
	path := "go-logo.png"
	img := image64(path)

	fig := &grob.Fig{
		Data: grob.Traces{
			&grob.Image{
				Type:   grob.TraceTypeImage,
				Source: img,
				Name:   "Image",
			},
		},
		Layout: &grob.Layout{
			Title: &grob.LayoutTitle{
				Text: path,
			},
			Xaxis: &grob.LayoutXaxis{
				Tickformat: "d",
			},
		},
		Config: &grob.Config{
			Displaymodebar: grob.ConfigDisplaymodebarTrue,
			Doubleclick:    grob.ConfigDoubleclickResetPlusautosize,
			Scrollzoom:     grob.ConfigScrollzoomTrue,
		},
	}

	offline.ToHtml(fig, "image.html")
	offline.Show(fig)
}
