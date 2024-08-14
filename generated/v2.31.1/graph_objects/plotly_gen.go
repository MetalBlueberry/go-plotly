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

// Color A string describing color. Supported formats: - hex (e.g. '#d3d3d3') - rgb (e.g. 'rgb(255, 0, 0)') - rgba (e.g. 'rgb(255, 0, 0, 0.5)') - hsl (e.g. 'hsl(0, 100%, 50%)') - hsv (e.g. 'hsv(0, 100%, 100%)') - named colors (full list: http://www.w3.org/TR/css3-color/#svg-color)",
type Color string

func UseColorScaleValues(in []float64) []ColorWithColorScale {
	out := make([]ColorWithColorScale, len(in), len(in))
	for i := 0; i < len(in); i++ {
		out[i] = ColorWithColorScale{
			Value: in[i],
		}
	}
	return out
}

func UseColors(in []Color) []ColorWithColorScale {
	out := make([]ColorWithColorScale, len(in), len(in))
	for i := 0; i < len(in); i++ {
		out[i] = ColorWithColorScale{
			Color: &in[i],
		}
	}
	return out
}

func UseColor(in Color) ColorWithColorScale {
	return ColorWithColorScale{
		Color: &in,
	}
}

type ColorWithColorScale struct {
	Color *Color
	Value float64
}

func (c *ColorWithColorScale) MarshalJSON() ([]byte, error) {
	if c.Color != nil {
		return json.Marshal(c.Color)
	}
	return json.Marshal(c.Value)
}

func (c *ColorWithColorScale) UnmarshalJSON(data []byte) error {
	c.Color = nil

	var color Color
	err := json.Unmarshal(data, &color)
	if err == nil {
		c.Color = &color
		return nil
	}

	var value float64
	err = json.Unmarshal(data, &value)
	if err != nil {
		return err
	}
	c.Value = value
	return nil
}

// ColorList A list of colors. Must be an {array} containing valid colors.
type ColorList []Color

// ColorScale A Plotly colorscale either picked by a name: (any of Greys, YlGnBu, Greens, YlOrRd, Bluered, RdBu, Reds, Blues, Picnic, Rainbow, Portland, Jet, Hot, Blackbody, Earth, Electric, Viridis, Cividis ) customized as an {array} of 2-element {arrays} where the first element is the normalized color level value (starting at *0* and ending at *1*), and the second item is a valid color string.
type ColorScale interface{}

func ArrayOKValue[T any](value T) ArrayOK[*T] {
	v := &value
	return ArrayOK[*T]{Value: v}
}

func ArrayOKArray[T any](array ...T) ArrayOK[*T] {
	out := make([]*T, len(array))
	for i, v := range array {
		value := v
		out[i] = &value
	}
	return ArrayOK[*T]{
		Array: out,
	}
}

// ArrayOK is a type that allows you to define a single value or an array of values, But not both.
// If Array is defined, Value will be ignored.
type ArrayOK[T any] struct {
	Value T
	Array []T
}

func (arrayOK *ArrayOK[T]) MarshalJSON() ([]byte, error) {
	if arrayOK.Array != nil {
		return json.Marshal(arrayOK.Array)
	}
	return json.Marshal(arrayOK.Value)
}

func (arrayOK *ArrayOK[T]) UnmarshalJSON(data []byte) error {
	arrayOK.Array = nil

	var array []T
	err := json.Unmarshal(data, &array)
	if err == nil {
		arrayOK.Array = array
		return nil
	}

	var value T
	err = json.Unmarshal(data, &value)
	if err != nil {
		return err
	}
	arrayOK.Value = value
	return nil
}
