package types_test

import (
	"encoding/json"
	"testing"

	"github.com/MetalBlueberry/go-plotly/types"
	. "github.com/onsi/gomega"
)

type TestMarshallScenario[T any] struct {
	Name     string
	Input    types.ArrayOK[T]
	Expected string
}

func TestFloat64Marshal(t *testing.T) {
	RegisterTestingT(t)

	scenarios := []TestMarshallScenario[*float64]{
		{
			Name:     "A single number",
			Input:    types.ArrayOKValue(12.3),
			Expected: "12.3",
		},
		{
			Name:     "A single number in a list",
			Input:    types.ArrayOKArray(12.3),
			Expected: "[12.3]",
		},
		{
			Name:     "Multiple values",
			Input:    types.ArrayOKArray(1, 2, 3, 4.5),
			Expected: "[1,2,3,4.5]",
		},
		{
			Name:     "Nil Value",
			Input:    types.ArrayOK[*float64]{},
			Expected: "null",
		},
		{
			Name: "Nil Array",
			Input: func() types.ArrayOK[*float64] {
				original := types.ArrayOKArray(1.2, 0, 3.4)
				original.Array[1] = nil
				return original
			}(),
			Expected: "[1.2,null,3.4]",
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := json.Marshal(&tt.Input)
			Expect(err).To(BeNil())
			Expect(string(result)).To(Equal(tt.Expected))
		})
	}
}

func TestStringMarshal(t *testing.T) {
	RegisterTestingT(t)

	scenarios := []TestMarshallScenario[*string]{
		{
			Name:     "A single string",
			Input:    types.ArrayOKValue("hello"),
			Expected: `"hello"`,
		},
		{
			Name:     "A single string in a list",
			Input:    types.ArrayOKArray("hello"),
			Expected: `["hello"]`,
		},
		{
			Name:     "Multiple strings",
			Input:    types.ArrayOKArray("hello", "world", "foo", "bar"),
			Expected: `["hello","world","foo","bar"]`,
		},
		{
			Name:     "Nil Value",
			Input:    types.ArrayOK[*string]{},
			Expected: "null",
		},
		{
			Name: "Nil Array",
			Input: func() types.ArrayOK[*string] {
				original := types.ArrayOKArray("hello", "", "world")
				original.Array[1] = nil
				return original
			}(),
			Expected: `["hello",null,"world"]`,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := json.Marshal(&tt.Input)
			Expect(err).To(BeNil())
			Expect(string(result)).To(Equal(tt.Expected))
		})
	}
}

type TestUnmarshalScenario[T any] struct {
	Name     string
	Input    string
	Expected types.ArrayOK[T]
}

func TestFloat64Unmarshal(t *testing.T) {
	RegisterTestingT(t)

	scenarios := []TestUnmarshalScenario[*float64]{
		{
			Name:     "A single number",
			Input:    "12.3",
			Expected: types.ArrayOKValue(12.3),
		},
		{
			Name:     "A single number in a list",
			Input:    "[12.3]",
			Expected: types.ArrayOKArray(12.3),
		},
		{
			Name:     "Multiple values",
			Input:    "[1,2,3,4.5]",
			Expected: types.ArrayOKArray(1, 2, 3, 4.5),
		},
		{
			Name:     "Nil Value",
			Input:    "null",
			Expected: types.ArrayOK[*float64]{},
		},
		{
			Name:  "Nil Array",
			Input: "[1.2,null,3.4]",
			Expected: func() types.ArrayOK[*float64] {
				original := types.ArrayOKArray(1.2, 0, 3.4)
				original.Array[1] = nil
				return original
			}(),
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.Name, func(t *testing.T) {
			result := types.ArrayOK[*float64]{}
			err := json.Unmarshal([]byte(tt.Input), &result)
			Expect(err).To(BeNil())
			Expect(result).To(Equal(tt.Expected))
		})
	}
}

func TestStringUnmarshal(t *testing.T) {
	RegisterTestingT(t)

	scenarios := []TestUnmarshalScenario[*string]{
		{
			Name:     "A single string",
			Input:    `"hello"`,
			Expected: types.ArrayOKValue("hello"),
		},
		{
			Name:     "A single string in a list",
			Input:    `["hello"]`,
			Expected: types.ArrayOKArray("hello"),
		},
		{
			Name:     "Multiple strings",
			Input:    `["hello","world","foo","bar"]`,
			Expected: types.ArrayOKArray("hello", "world", "foo", "bar"),
		},
		{
			Name:     "Nil Value",
			Input:    "null",
			Expected: types.ArrayOK[*string]{},
		},
		{
			Name:  "Nil Array",
			Input: `["hello",null,"world"]`,
			Expected: func() types.ArrayOK[*string] {
				original := types.ArrayOKArray("hello", "", "world")
				original.Array[1] = nil
				return original
			}(),
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.Name, func(t *testing.T) {
			result := types.ArrayOK[*string]{}
			err := json.Unmarshal([]byte(tt.Input), &result)
			Expect(err).To(BeNil())
			Expect(result).To(Equal(tt.Expected))
		})
	}
}

type Color struct {
	Red   int `json:"red"`
	Green int `json:"green"`
	Blue  int `json:"blue"`
}

func TestColorUnmarshal(t *testing.T) {
	RegisterTestingT(t)

	scenarios := []TestUnmarshalScenario[*Color]{
		{
			Name:     "A single color",
			Input:    `{"red": 255, "green": 255, "blue": 255}`,
			Expected: types.ArrayOKValue(Color{Red: 255, Green: 255, Blue: 255}),
		},
		{
			Name:  "A single color in a list",
			Input: `[{"red": 255, "green": 255, "blue": 255}]`,
			Expected: types.ArrayOK[*Color]{
				Array: []*Color{{Red: 255, Green: 255, Blue: 255}},
			},
		},
		{
			Name:  "Multiple colors",
			Input: `[{"red": 255, "green": 255, "blue": 255}, {"red": 0, "green": 0, "blue": 0}, {"red": 0, "green": 255, "blue": 0}]`,
			Expected: types.ArrayOK[*Color]{
				Array: []*Color{
					{Red: 255, Green: 255, Blue: 255},
					{Red: 0, Green: 0, Blue: 0},
					{Red: 0, Green: 255, Blue: 0},
				},
			},
		},
		{
			Name:     "Nil Value",
			Input:    "null",
			Expected: types.ArrayOK[*Color]{},
		},
		{
			Name:  "Nil Array",
			Input: `[{"red": 255, "green": 255, "blue": 255}, null, {"red": 0, "green": 0, "blue": 0}]`,
			Expected: types.ArrayOK[*Color]{
				Array: []*Color{
					{Red: 255, Green: 255, Blue: 255},
					nil,
					{Red: 0, Green: 0, Blue: 0},
				},
			},
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.Name, func(t *testing.T) {
			result := types.ArrayOK[*Color]{}
			err := json.Unmarshal([]byte(tt.Input), &result)
			Expect(err).To(BeNil())
			Expect(result).To(Equal(tt.Expected))
		})
	}
}

func TestColorMarshal(t *testing.T) {
	RegisterTestingT(t)

	scenarios := []TestMarshallScenario[*Color]{
		{
			Name:     "A single color",
			Input:    types.ArrayOKValue(Color{Red: 255, Green: 255, Blue: 255}),
			Expected: `{"red":255,"green":255,"blue":255}`,
		},
		{
			Name: "A single color in a list",
			Input: types.ArrayOK[*Color]{
				Array: []*Color{{Red: 255, Green: 255, Blue: 255}},
			},
			Expected: `[{"red":255,"green":255,"blue":255}]`,
		},
		{
			Name: "Multiple colors",
			Input: types.ArrayOK[*Color]{
				Array: []*Color{
					{Red: 255, Green: 255, Blue: 255},
					{Red: 0, Green: 0, Blue: 0},
					{Red: 0, Green: 255, Blue: 0},
				},
			},
			Expected: `[{"red":255,"green":255,"blue":255},{"red":0,"green":0,"blue":0},{"red":0,"green":255,"blue":0}]`,
		},
		{
			Name:     "Nil Value",
			Input:    types.ArrayOK[*Color]{},
			Expected: "null",
		},
		{
			Name: "Nil Array",
			Input: types.ArrayOK[*Color]{
				Array: []*Color{
					{Red: 255, Green: 255, Blue: 255},
					nil,
					{Red: 0, Green: 0, Blue: 0},
				},
			},
			Expected: `[{"red":255,"green":255,"blue":255},null,{"red":0,"green":0,"blue":0}]`,
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := json.Marshal(&tt.Input)
			Expect(err).To(BeNil())
			Expect(string(result)).To(Equal(tt.Expected))
		})
	}
}

func TestNestedArrayOKUnmarshal(t *testing.T) {
	type Color struct {
		Red   int `json:"red"`
		Green int `json:"green"`
		Blue  int `json:"blue"`
	}

	tests := []struct {
		Name     string
		Input    interface{}
		Expected string
	}{
		{
			Name: "Nested ArrayOK within another struct",
			Input: struct {
				Name   string                `json:"name"`
				Colors *types.ArrayOK[Color] `json:"colors"`
			}{
				Name: "Test",
				Colors: &types.ArrayOK[Color]{
					Array: []Color{
						{Red: 255, Green: 255, Blue: 255},
						{Red: 0, Green: 0, Blue: 0},
						{Red: 0, Green: 255, Blue: 0},
					},
				},
			},
			Expected: `{"name":"Test","colors":[{"red":255,"green":255,"blue":255},{"red":0,"green":0,"blue":0},{"red":0,"green":255,"blue":0}]}`,
		},
		{
			Name: "Nested ArrayOK within another struct with nil Colors",
			Input: struct {
				Name   string                `json:"name"`
				Colors *types.ArrayOK[Color] `json:"colors"`
			}{
				Name:   "Test",
				Colors: nil,
			},
			Expected: `{"name":"Test","colors":null}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			bytes, err := json.Marshal(tt.Input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if string(bytes) != tt.Expected {
				t.Errorf("expected %s, but got %s", tt.Expected, string(bytes))
			}
		})
	}
}
func TestNestedArrayOKMarshal(t *testing.T) {
	tests := []struct {
		Name     string
		Input    interface{}
		Expected string
	}{
		{
			Name: "Marshal ArrayOK nested within another struct with actual value",
			Input: struct {
				Name   string                `json:"name"`
				Colors *types.ArrayOK[Color] `json:"colors"`
			}{
				Name: "Test",
				Colors: &types.ArrayOK[Color]{
					Array: []Color{
						{Red: 255, Green: 255, Blue: 255},
						{Red: 0, Green: 0, Blue: 0},
						{Red: 0, Green: 255, Blue: 0},
					},
				},
			},
			Expected: `{"name":"Test","colors":[{"red":255,"green":255,"blue":255},{"red":0,"green":0,"blue":0},{"red":0,"green":255,"blue":0}]}`,
		},
		{
			Name: "Marshal ArrayOK nested within another struct with nil value",
			Input: struct {
				Name   string                `json:"name"`
				Colors *types.ArrayOK[Color] `json:"colors"`
			}{
				Name:   "Test",
				Colors: nil,
			},
			Expected: `{"name":"Test","colors":null}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			bytes, err := json.Marshal(tt.Input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if string(bytes) != tt.Expected {
				t.Errorf("expected %s, but got %s", tt.Expected, string(bytes))
			}
		})
	}
}
