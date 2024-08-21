package grob

import (
	"encoding/json"
)

// TraceType is the type for the TraceType field on every trace
type TraceType string

// Trace Every trace implements this interface
// It is useful for autocompletion, it is a better idea to use
// type assertions/switches to identify trace types
type Trace interface {
	GetType() TraceType
}

// Traces is a slice of Traces
type Traces []Trace

// Fig is the base type for figures.
type Fig struct {
	// Data The data to be plotted is described in an array usually called data, whose elements are trace objects of various types (e.g. scatter, bar etc) as documented in the Full Reference.
	// https://plotly.com/javascript/reference
	Data Traces `json:"data,omitempty"`

	// Layout The layout of the plot – non-data-related visual attributes such as the title, annotations etc – is described in an object usually called layout, as documented in/ the Full Reference.
	// https://plotly.com/javascript/reference/layout
	Layout *Layout `json:"layout,omitempty"`

	// Config High-level configuration options for the plot, such as the scroll/zoom/hover behaviour, is described in an object usually called config, as documented here. The difference between config and layout is that layout relates to the content of the plot, whereas config relates to the context in which the plot is being shown.
	// https://plotly.com/javascript/configuration-options
	Config *Config `json:"config,omitempty"`

	// Animation is not yet implemented, feel free to insert custom a struct
	Animation interface{} `json:"animation,omitempty"`
}

// AddTraces Is a shorthand  to add figures to a given figure. It handles the case where the Traces value is nil.
func (fig *Fig) AddTraces(traces ...Trace) {
	if fig.Data == nil {
		fig.Data = make(Traces, 0)
	}
	fig.Data = append(fig.Data, traces...)
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

// Bool represents a *bool value. Needed to tell the different between false and nil.
type Bool *bool

var (
	trueValue  bool = true
	falseValue bool = false

	// True is a *bool with true value
	True Bool = &trueValue
	// False is a *bool with false value
	False Bool = &falseValue
)

// String is a string value, can be a []string if arrayOK is true.
// numeric values are converted to string by plotly, so []<number> can work
type String interface{}
