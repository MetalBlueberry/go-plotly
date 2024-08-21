package grob

import "encoding/json"

// Types defined here are just meant to allow compilation for the template package
// This simplifies the process of writing the static templates

type Layout struct{}
type Config struct{}

func UnmarshalTrace(json.RawMessage) (Trace, error) { return &Bar{}, nil }

type Bar struct {
	Type TraceType `json:"type"`
}

func (b *Bar) GetType() TraceType { return "bar" }
