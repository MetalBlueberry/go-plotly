package types_test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/MetalBlueberry/go-plotly/pkg/types"
	. "github.com/onsi/gomega"
)

var testSomeEncodedDataString = "SomeEncodedData"
var testSomeEncodedDataBytes = []byte("SomeEncodedData")
var testSomeEncodedData = base64.StdEncoding.EncodeToString([]byte(testSomeEncodedDataString))

func TestDataArrayType_MarshalJSON(t *testing.T) {

	tests := []struct {
		name     string
		data     *types.DataArrayType
		expected string
	}{
		{
			name:     "Marshal with values",
			data:     types.DataArray([]int{1, 2, 3}),
			expected: `[1,2,3]`,
		},
		{
			name:     "Marshal with bdata",
			data:     types.BDataArray(types.DTypeFloat32, testSomeEncodedDataBytes, []int{2, 2}),
			expected: fmt.Sprintf(`{"dtype":"float32","bdata":"%s","shape":"2,2"}`, testSomeEncodedData),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := tt.data.MarshalJSON()
			if err != nil {
				t.Fatalf("MarshalJSON() error = %v", err)
			}
			if string(b) != tt.expected {
				t.Errorf("MarshalJSON() = %v, want %v", string(b), tt.expected)
			}
		})
	}
}

func TestDataArrayType_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		expected *types.DataArrayType
	}{
		{
			name:     "Unmarshal with values",
			data:     `[1,2,3]`,
			expected: types.DataArray([]float64{1, 2, 3}),
		},
		{
			name:     "Unmarshal with bdata",
			data:     fmt.Sprintf(`{"dtype":"float32","bdata":"%s","shape":"2,2"}`, testSomeEncodedData),
			expected: types.BDataArray(types.DTypeFloat32, testSomeEncodedDataBytes, []int{2, 2}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			got := &types.DataArrayType{}
			if err := json.Unmarshal([]byte(tt.data), got); err != nil {
				t.Fatalf("UnmarshalJSON() error = %v", err)
			}
			g.Expect(tt.expected).To(Equal(got))
		})
	}
}

func TestDataArrayShape_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		data     types.DataArrayShape
		expected string
	}{
		{
			name:     "Marshal shape",
			data:     types.DataArrayShape{2, 3, 4},
			expected: `"2,3,4"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := tt.data.MarshalJSON()
			if err != nil {
				t.Fatalf("MarshalJSON() error = %v", err)
			}
			if string(b) != tt.expected {
				t.Errorf("MarshalJSON() = %v, want %v", string(b), tt.expected)
			}
		})
	}
}

func TestDataArrayShape_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		data     string
		expected types.DataArrayShape
	}{
		{
			name:     "Unmarshal shape",
			data:     `"2,3,4"`,
			expected: types.DataArrayShape{2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got types.DataArrayShape
			if err := json.Unmarshal([]byte(tt.data), &got); err != nil {
				t.Fatalf("UnmarshalJSON() error = %v", err)
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("UnmarshalJSON() = %+v, want %+v", got, tt.expected)
			}
		})
	}
}
