package grob

import (
	"encoding/json"

	"github.com/MetalBlueberry/go-plotly/pkg/types"
)

// Types defined here are just meant to allow compilation for the template package
// This simplifies the process of writing the static templates

type Layout struct{}
type Config struct{}

func UnmarshalTrace(json.RawMessage) (types.Trace, error) { return &Bar{}, nil }

type Bar struct {
	Type types.TraceType `json:"type"`
}

func (b *Bar) GetType() types.TraceType { return "bar" }
