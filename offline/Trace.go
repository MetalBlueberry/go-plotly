package offlinePlotly

import (
	"fmt"
)

type Trace map[string]interface{}

type TraceType int

const (
	SCATTER TraceType = 0 + iota
	LINE
)

var traceTypes = [...]string{
	"scatter",
	"line",
}

func (tr TraceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", traceTypes[tr])), nil
}

func (trace Trace) addNumbericValue(name string, x interface{}) error {
	switch xtype := x.(type) {
	case []int:
		if _, ok := trace[name]; !ok {
			trace[name] = make([]int, 0)
		}
		xvalues, ok := trace[name].([]int)
		if !ok {
			return fmt.Errorf("previous data in X doesn't match the type, %T vs %T", xtype, trace[name])
		}
		trace[name] = append(xvalues, xtype...)
	case []float32:
		if _, ok := trace[name]; !ok {
			trace[name] = make([]float32, 0)
		}
		xvalues, ok := trace[name].([]float32)
		if !ok {
			return fmt.Errorf("previous data in X doesn't match the type, %T vs %T", xtype, trace[name])
		}
		trace[name] = append(xvalues, xtype...)
	case []float64:
		if _, ok := trace[name]; !ok {
			trace[name] = make([]float64, 0)
		}
		xvalues, ok := trace[name].([]float64)
		if !ok {
			return fmt.Errorf("previous data in X doesn't match the type, %T vs %T", xtype, trace[name])
		}
		trace[name] = append(xvalues, xtype...)
	case int:
		return trace.addNumbericValue(name, []int{xtype})
	case float32:
		return trace.addNumbericValue(name, []float32{xtype})
	case float64:
		return trace.addNumbericValue(name, []float64{xtype})
	default:
		return fmt.Errorf("Data type is not valid, use int, float32 or float64")
	}
	return nil
}

func (trace Trace) AddPoints(x, y interface{}) error {
	if err := trace.addNumbericValue("x", x); err != nil {
		return err
	}

	if err := trace.addNumbericValue("y", y); err != nil {
		return err
	}

	return nil
}
