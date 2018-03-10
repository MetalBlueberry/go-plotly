package offline

import (
	"fmt"
	"strings"
)

type Scatter map[string]interface{}

// ScatterMode : Type to set ScatterMode
type ScatterMode int

// Flags to set ScatterMode
const (
	MODELINES   ScatterMode = 1 << iota
	MODEMARKERS ScatterMode = 1 << iota
	MODETEXT    ScatterMode = 1 << iota
)

func (md ScatterMode) MarshalJSON() ([]byte, error) {
	modes := make([]string, 0)
	if md|MODELINES == md {
		modes = append(modes, "lines")
	}
	if md|MODEMARKERS == md {
		modes = append(modes, "markers")
	}
	if md|MODETEXT == md {
		modes = append(modes, "text")
	}
	if len(modes) == 0 {
		return []byte("\"\""), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", strings.Join(modes, "+"))), nil

}

func NewScatter() Scatter {
	return Scatter{
		"type": "scatter",
	}
}

func (s Scatter) SetMode(mode ScatterMode) {
	s["mode"] = mode
}

func (s Scatter) SetName(name string) {
	s["name"] = (name)
}

func (s Scatter) AddPoints(x, y interface{}) error {
	if err := addNumbericValue(s, "x", x); err != nil {
		return err
	}
	return addNumbericValue(s, "y", y)

}

func (s Scatter) NewMarker() Marker {
	s["marker"] = NewMarker()
	return s["marker"].(Marker)
}
