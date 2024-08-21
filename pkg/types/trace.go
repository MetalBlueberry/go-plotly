package types

// TraceType is the type for the TraceType field on every trace
type TraceType string

// Trace Every trace implements this interface
// It is useful for autocompletion, it is a better idea to use
// type assertions/switches to identify trace types
type Trace interface {
	GetType() TraceType
}
