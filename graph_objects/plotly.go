package grob

import (
	"encoding/json"
)

//go:generate go run ../generator/cmd/generator/main.go --schema ../generator/schema.json --output-directory .

type TraceType string

type Trace interface {
	GetType() TraceType
}

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

// String is a string value, numeric values are converted to string by plotly
type String interface{}

// Color A string describing color. Supported formats: - hex (e.g. '#d3d3d3') - rgb (e.g. 'rgb(255, 0, 0)') - rgba (e.g. 'rgb(255, 0, 0, 0.5)') - hsl (e.g. 'hsl(0, 100%, 50%)') - hsv (e.g. 'hsv(0, 100%, 100%)') - named colors (full list: http://www.w3.org/TR/css3-color/#svg-color)",
type Color interface{}

// ColorList A list of colors. Must be an {array} containing valid colors.
type ColorList []Color

// ColorScale A Plotly colorscale either picked by a name: (any of Greys, YlGnBu, Greens, YlOrRd, Bluered, RdBu, Reds, Blues, Picnic, Rainbow, Portland, Jet, Hot, Blackbody, Earth, Electric, Viridis, Cividis ) customized as an {array} of 2-element {arrays} where the first element is the normalized color level value (starting at *0* and ending at *1*), and the second item is a valid color string.
type ColorScale interface{}
