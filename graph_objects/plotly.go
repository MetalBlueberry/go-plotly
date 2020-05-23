package graph_objects

import (
	"encoding/json"
)

//go:generate sh -c "cat schema.json | plate --template main.tmpl; go fmt"

type Traces []Trace

type Fig struct {
	Data   Traces  `json:"data,omitempty"`
	Layout *Layout `json:"layout,omitempty"`
	Config *Config `json:"config,omitempty"`
}

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

type Bool *bool

var (
	trueValue  bool = true
	falseValue bool = false

	True  Bool = &trueValue
	False Bool = &falseValue
)

type String interface{}
