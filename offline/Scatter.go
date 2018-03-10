package offlinePlotly

import (
	"fmt"
	"strings"
)

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

func NewScatter() Trace {
	return Trace{
		"type": SCATTER,
	}
}

func (trace Trace) SetMode(mode ScatterMode) {
	trace["mode"] = mode
}
