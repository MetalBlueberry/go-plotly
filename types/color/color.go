package color

import (
	"encoding/json"
	"fmt"
)

// String A string describing color. Supported formats: - hex (e.g. '#d3d3d3') - rgb (e.g. 'rgb(255, 0, 0)') - rgba (e.g. 'rgb(255, 0, 0, 0.5)') - hsl (e.g. 'hsl(0, 100%, 50%)') - hsv (e.g. 'hsv(0, 100%, 100%)') - named colors (full list: http://www.w3.org/TR/css3-color/#svg-color)",
type String string

func UseColorScaleValues(in []float64) []ColorWithColorScale {
	out := make([]ColorWithColorScale, len(in), len(in))
	for i := 0; i < len(in); i++ {
		out[i] = ColorWithColorScale{
			Value: in[i],
		}
	}
	return out
}

func UseColors(in []String) []ColorWithColorScale {
	out := make([]ColorWithColorScale, len(in), len(in))
	for i := 0; i < len(in); i++ {
		out[i] = ColorWithColorScale{
			Color: &in[i],
		}
	}
	return out
}

func UseColor(in String) ColorWithColorScale {
	return ColorWithColorScale{
		Color: &in,
	}
}

type ColorWithColorScale struct {
	Color *String
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

	var color String
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

// List A list of colors. Must be an {array} containing valid colors.
type List []String

// Scale A Plotly colorscale either picked by a name: (any of Greys, YlGnBu, Greens, YlOrRd, Bluered, RdBu, Reds, Blues, Picnic, Rainbow, Portland, Jet, Hot, Blackbody, Earth, Electric, Viridis, Cividis ) customized as an {array} of 2-element {arrays} where the first element is the normalized color level value (starting at *0* and ending at *1*), and the second item is a valid color string.
type Scale struct {
	Name   string
	Values []ScaleReference
}

type ScaleReference struct {
	NormalizedValue float64
	Color           String
}

func (cs *ScaleReference) MarshalJSON() ([]byte, error) {
	data := []interface{}{cs.NormalizedValue, cs.Color}
	return json.Marshal(data)
}

func (cs *Scale) MarshalJSON() ([]byte, error) {
	if cs.Values != nil {
		return json.Marshal(cs.Values)
	}
	return json.Marshal(cs.Name)
}

func (cs *Scale) UnmarshalJSON(data []byte) error {
	var values []ScaleReference

	err := json.Unmarshal(data, &values)
	if err == nil {
		cs.Values = values
		return nil
	}
	var name string
	err = json.Unmarshal(data, &name)
	if err == nil {
		cs.Name = name
		return nil
	}
	return fmt.Errorf("Unable to Unmarshal ColorScale, %w", err)
}
