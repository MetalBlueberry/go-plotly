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

	scenarios := []TestMarshallScenario[float64]{
		{
			Name: "A single number",
			Input: types.ArrayOK[float64]{
				Value: 12.3,
			},
			Expected: "12.3",
		},
		{
			Name: "A single number in a list",
			Input: types.ArrayOK[float64]{
				Array: []float64{12.3},
			},
			Expected: "[12.3]",
		},
		{
			Name: "Multiple values",
			Input: types.ArrayOK[float64]{
				Array: []float64{1, 2, 3, 4.5},
			},
			Expected: "[1,2,3,4.5]",
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

	scenarios := []TestMarshallScenario[string]{
		{
			Name: "A single string",
			Input: types.ArrayOK[string]{
				Value: "hello",
			},
			Expected: `"hello"`,
		},
		{
			Name: "A single string in a list",
			Input: types.ArrayOK[string]{
				Array: []string{"hello"},
			},
			Expected: `["hello"]`,
		},
		{
			Name: "Multiple strings",
			Input: types.ArrayOK[string]{
				Array: []string{"hello", "world", "foo", "bar"},
			},
			Expected: `["hello","world","foo","bar"]`,
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

type TestUnmarshallScenario[T any] struct {
	Name     string
	Input    string
	Expected types.ArrayOK[T]
}

func TestFloat64Unmarshal(t *testing.T) {
	RegisterTestingT(t)

	scenarios := []TestUnmarshallScenario[float64]{
		{
			Name:  "A single number",
			Input: "12.3",
			Expected: types.ArrayOK[float64]{
				Value: 12.3,
			},
		},
		{
			Name:  "A single number in a list",
			Input: "[12.3]",
			Expected: types.ArrayOK[float64]{
				Array: []float64{12.3},
			},
		},
		{
			Name:  "Multiple values",
			Input: "[1,2,3,4.5]",
			Expected: types.ArrayOK[float64]{
				Array: []float64{1, 2, 3, 4.5},
			},
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.Name, func(t *testing.T) {
			result := types.ArrayOK[float64]{}
			err := json.Unmarshal([]byte(tt.Input), &result)
			Expect(err).To(BeNil())
			Expect(result).To(Equal(tt.Expected))
		})
	}
}

func TestStringUnmarshal(t *testing.T) {
	RegisterTestingT(t)

	scenarios := []TestUnmarshallScenario[string]{
		{
			Name:  "A single string",
			Input: `"hello"`,
			Expected: types.ArrayOK[string]{
				Value: "hello",
			},
		},
		{
			Name:  "A single string in a list",
			Input: `["hello"]`,
			Expected: types.ArrayOK[string]{
				Array: []string{"hello"},
			},
		},
		{
			Name:  "Multiple strings",
			Input: `["hello","world","foo","bar"]`,
			Expected: types.ArrayOK[string]{
				Array: []string{"hello", "world", "foo", "bar"},
			},
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.Name, func(t *testing.T) {
			result := types.ArrayOK[string]{}
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

	scenarios := []TestUnmarshallScenario[Color]{
		{
			Name:  "A single color",
			Input: `{"red": 255, "green": 255, "blue": 255}`,
			Expected: types.ArrayOK[Color]{
				Value: Color{Red: 255, Green: 255, Blue: 255},
			},
		},
		{
			Name:  "A single color in a list",
			Input: `[{"red": 255, "green": 255, "blue": 255}]`,
			Expected: types.ArrayOK[Color]{
				Array: []Color{{Red: 255, Green: 255, Blue: 255}},
			},
		},
		{
			Name:  "Multiple colors",
			Input: `[{"red": 255, "green": 255, "blue": 255}, {"red": 0, "green": 0, "blue": 0}, {"red": 0, "green": 255, "blue": 0}]`,
			Expected: types.ArrayOK[Color]{
				Array: []Color{
					{Red: 255, Green: 255, Blue: 255},
					{Red: 0, Green: 0, Blue: 0},
					{Red: 0, Green: 255, Blue: 0},
				},
			},
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.Name, func(t *testing.T) {
			result := types.ArrayOK[Color]{}
			err := json.Unmarshal([]byte(tt.Input), &result)
			Expect(err).To(BeNil())
			Expect(result).To(Equal(tt.Expected))
		})
	}
}

func TestColorMarshal(t *testing.T) {
	RegisterTestingT(t)

	scenarios := []TestMarshallScenario[Color]{
		{
			Name: "A single color",
			Input: types.ArrayOK[Color]{
				Value: Color{Red: 255, Green: 255, Blue: 255},
			},
			Expected: `{"red":255,"green":255,"blue":255}`,
		},
		{
			Name: "A single color in a list",
			Input: types.ArrayOK[Color]{
				Array: []Color{{Red: 255, Green: 255, Blue: 255}},
			},
			Expected: `[{"red":255,"green":255,"blue":255}]`,
		},
		{
			Name: "Multiple colors",
			Input: types.ArrayOK[Color]{
				Array: []Color{
					{Red: 255, Green: 255, Blue: 255},
					{Red: 0, Green: 0, Blue: 0},
					{Red: 0, Green: 255, Blue: 0},
				},
			},
			Expected: `[{"red":255,"green":255,"blue":255},{"red":0,"green":0,"blue":0},{"red":0,"green":255,"blue":0}]`,
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
