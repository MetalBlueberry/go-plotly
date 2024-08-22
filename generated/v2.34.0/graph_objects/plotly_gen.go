package grob

import (
	"encoding/json"

	"github.com/MetalBlueberry/go-plotly/pkg/types"
)

// Fig is the base type for figures.
type Fig struct {
	// Data The data to be plotted is described in an array usually called data, whose elements are trace objects of various types (e.g. scatter, bar etc) as documented in the Full Reference.
	// https://plotly.com/javascript/reference
	Data []types.Trace `json:"data,omitempty"`

	// Layout The layout of the plot – non-data-related visual attributes such as the title, annotations etc – is described in an object usually called layout, as documented in/ the Full Reference.
	// https://plotly.com/javascript/reference/layout
	Layout *Layout `json:"layout,omitempty"`

	// Config High-level configuration options for the plot, such as the scroll/zoom/hover behaviour, is described in an object usually called config, as documented here. The difference between config and layout is that layout relates to the content of the plot, whereas config relates to the context in which the plot is being shown.
	// https://plotly.com/javascript/configuration-options
	Config *Config `json:"config,omitempty"`

	// Animation is not yet implemented, feel free to insert custom a struct
	Animation *Animation `json:"animation,omitempty"`

	// Frames are the animation frames
	Frames []Frame `json:"frames,omitempty"`
}

// AddTraces Is a shorthand  to add figures to a given figure. It handles the case where the Traces value is nil.
func (fig *Fig) AddTraces(traces ...types.Trace) {
	if fig.Data == nil {
		fig.Data = make([]types.Trace, 0)
	}
	fig.Data = append(fig.Data, traces...)
}

func (fig *Fig) Info() types.Version {
	return types.Version{
		//
		Name: "Plotly 2.34.0",
		//
		Tag: "v2.34.0",
		//
		URL: "https://raw.githubusercontent.com/plotly/plotly.js/v2.34.0/test/plot-schema.json",
		//
		Path: "schemas/v2.34.0/plot-schema.json",
		//
		Generated: "generated/v2.34.0",
		//
		Cdn: "https://cdn.plot.ly/plotly-2.34.0.min.js",
	}
}

// UnmarshalJSON is a custom unmarshal function to properly handle special cases.
func (fig *Fig) UnmarshalJSON(data []byte) error {
	var err error
	tmp := unmarshalFig{}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	fig.Layout = tmp.Layout
	fig.Config = tmp.Config

	for i := range tmp.Data {
		trace, err := UnmarshalTrace(tmp.Data[i])
		if err != nil {
			return err
		}
		fig.AddTraces(trace)
	}
	return nil
}

type unmarshalFig struct {
	Data   []json.RawMessage `json:"data,omitempty"`
	Layout *Layout           `json:"layout,omitempty"`
	Config *Config           `json:"config,omitempty"`
}
