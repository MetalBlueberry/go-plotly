package types

import (
	"strconv"
)

// BoolType represents a *bool value. Needed to tell the different between false and nil.
type BoolType *bool

func B(v bool) BoolType {
	return BoolType(&v)
}

// BA converts a list of bool into BooleanType
func BA(v []bool) []BoolType {
	result := make([]BoolType, len(v))
	for i := range v {
		result[i] = B(v[i])
	}
	return result
}

var (
	trueValue  bool = true
	falseValue bool = false

	// True is a *bool with true value
	True BoolType = &trueValue
	// False is a *bool with false value
	False BoolType = &falseValue
)

// StringType as defined by plotly schema
// This is not *string because I do not see any case where an empty string is a desired input for plotly and having this as string makes it much easier to work with the package.
type StringType string

// S converts a string to a StringType
// It is not needed, but if you use this method, it will make it easier to migrate to different implementations of String type in case we actually need *string instead of string
func S(v string) StringType {
	return StringType(v)
}

// NumberType as defined by plotly schema
type NumberType *float64

func N(n float64) NumberType {
	return NumberType(&n)
}

// NA converts a list of float64 to NumberType
func NA(n []float64) []NumberType {
	result := make([]NumberType, len(n))
	for i := range n {
		result[i] = N(n[i])
	}
	return result
}

// NS Given a string, parses it as a float64 number
// Panics if the string is not a float number
func NS(n string) NumberType {
	v, err := strconv.ParseFloat(n, 64)
	if err != nil {
		panic(err)
	}
	return NumberType(&v)
}

// IntegerType as defined by plotly schema
type IntegerType *int

func I(n int) IntegerType {
	return IntegerType(&n)
}

// IA converts a list of int to IntegerType
func IA(n []int) []IntegerType {
	result := make([]IntegerType, len(n))
	for i := range n {
		result[i] = I(n[i])
	}
	return result
}

// IS Given a string, parses it as an integer number
// Panics if the string is not an integer number
func IS(n string) IntegerType {
	v, err := strconv.Atoi(n)
	if err != nil {
		panic(err)
	}
	return IntegerType(&v)
}
