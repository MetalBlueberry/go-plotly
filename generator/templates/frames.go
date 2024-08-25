package grob

// {{ . }}

import (
	"github.com/MetalBlueberry/go-plotly/pkg/types"
)

// Frame
type Frame struct {

	// Baseframe
	// The name of the frame into which this frame's properties are merged before applying. This is used to unify properties and avoid needing to specify the same values for the same properties in multiple frames.
	Baseframe types.StringType `json:"baseframe,omitempty"`

	// Data
	// A list of traces this frame modifies. The format is identical to the normal trace definition.
	Data []types.Trace `json:"data,omitempty"`

	// Group
	// An identifier that specifies the group to which the frame belongs, used by animate to select a subset of frames.
	Group types.StringType `json:"group,omitempty"`

	// Layout
	// Layout properties which this frame modifies. The format is identical to the normal layout definition.
	Layout *Layout `json:"layout,omitempty"`

	// Name
	// A label by which to identify the frame
	Name types.StringType `json:"name,omitempty"`

	// Traces
	// A list of trace indices that identify the respective traces in the data attribute
	Traces []int `json:"traces,omitempty"`
}
