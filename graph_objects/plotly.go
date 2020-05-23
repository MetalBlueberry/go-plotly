package graph_objects

import (
	"encoding/json"
)

//go:generate sh -c "cat schema.json | plate --template main.tmpl; go fmt"

// Traces is a slice of Traces
type Traces []Trace

// Fig is the base type for figures.
type Fig struct {
	Data   Traces  `json:"data,omitempty"`
	Layout *Layout `json:"layout,omitempty"`
	Config *Config `json:"config,omitempty"`
}

// AddTraces Is a shorthand  to add figures to a given figure. It handles the case where the Traces value is nil.
func (fig *Fig) AddTraces(traces ...Trace) {
	if fig.Data == nil {
		fig.Data = make(Traces, 0)
	}
	fig.Data = append(fig.Data, traces...)
}

type unmarshalFig struct {
	Data   []json.RawMessage `json:"data,omitempty"`
	Layout *Layout           `json:"layout,omitempty"`
	Config *Config           `json:"config,omitempty"`
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
		trace, err := UnmarshallTrace(tmp.Data[i])
		if err != nil {
			return err
		}
		fig.AddTraces(trace)
	}
	return nil
}

// This section is to workaround the omitempty for json serialization.

// Bool represents a *bool value. Needed to tell the differenc between false and nil.
type Bool *bool

var (
	trueValue  bool = true
	falseValue bool = false

	// True is a *bool with true value
	True Bool = &trueValue
	// False is a *bool with false value
	False Bool = &falseValue
)

type String interface{}
