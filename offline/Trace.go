package offline

// import (
// 	"fmt"
// )

// type Trace map[string]interface{}

// type TraceType int

// const (
// 	SCATTER TraceType = 0 + iota
// 	LINE
// )

// var traceTypes = [...]string{
// 	"scatter",
// 	"line",
// 	"bar",
// }

// func (tr TraceType) MarshalJSON() ([]byte, error) {
// 	return []byte(fmt.Sprintf("\"%s\"", traceTypes[tr])), nil
// }

// func (trace Trace) AddPoints(x, y interface{}) error {
// 	if err := addNumbericValue(trace, "x", x); err != nil {
// 		return err
// 	}

// 	if err := addNumbericValue(trace, "y", y); err != nil {
// 		return err
// 	}

// 	return nil
// }
// func (trace Trace) SetName(name string) {
// 	trace["name"] = name
// }

// func (trace Trace) NewMarker() Marker {
// 	trace["marker"] = NewMarker()
// 	return trace["marker"].(Marker)
// }
// func (trace Trace) SetMarker(marker Marker) {
// 	trace["marker"] = marker
// }
