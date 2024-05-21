package types_test

import (
	"encoding/json"
	"testing"

	"github.com/MetalBlueberry/go-plotly/types"
	. "github.com/onsi/gomega"
)

type TestMarshallScenario struct {
	Name     string
	Input    types.ArrayOK[float64]
	Expected string
}

func TestFloat64Marshal(t *testing.T) {
	RegisterTestingT(t)

	table := []TestMarshallScenario{
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
	for _, tt := range table {
		t.Run(tt.Name, func(t *testing.T) {
			result, err := json.Marshal(&tt.Input)
			Expect(err).To(BeNil())
			Expect(string(result)).To(Equal(tt.Expected))

		})

	}

}

type TestUnmarshallScenario struct {
	Name     string
	Input    string
	Expected types.ArrayOK[float64]
}

func TestFloat64Unmarshal(t *testing.T) {
	RegisterTestingT(t)

	table := []TestUnmarshallScenario{
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
	for _, tt := range table {
		t.Run(tt.Name, func(t *testing.T) {
			result := types.ArrayOK[float64]{}
			err := json.Unmarshal([]byte(tt.Input), &result)
			Expect(err).To(BeNil())
			Expect(result).To(Equal(tt.Expected))

		})

	}

}
