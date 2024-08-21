package types

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
