package grob

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestFig_Marshal_Unmarshal_does_not_change_content(t *testing.T) {
	tests := []struct {
		name   string
		figure Fig
	}{
		{
			name: "Simple scatter plot",
			figure: Fig{
				Data: Traces{
					&Scatter{
						Type: TraceTypeScatter,
						Name: "Simple Scatter",
						X:    []float64{1.1, 2.2, 3.3},
						Y:    []float64{3.3, 2.2, 1.1},
					},
				},
			},
		},
		{
			name: "Simple box plot",
			figure: Fig{
				Data: Traces{
					&Box{
						Type: TraceTypeBox,
						Name: "Simple Scatter",
						X:    []float64{1.1, 2.2, 3.3},
						Y:    []float64{3.3, 2.2, 1.1},
					},
				},
			},
		},
		{
			name: "with Bool values",
			figure: Fig{
				Data: Traces{
					&Scatter{
						Type:       TraceTypeScatter,
						Name:       "Simple Scatter",
						X:          []float64{1.1, 2.2, 3.3},
						Y:          []float64{3.3, 2.2, 1.1},
						Visible:    True,
						Cliponaxis: False,
					},
				},
			},
		},
		{
			name: "with multiple traces",
			figure: Fig{
				Data: Traces{
					&Scatter{
						Type: TraceTypeScatter,
						Name: "Simple Scatter",
						X:    []float64{1.1, 2.2, 3.3},
						Y:    []float64{3.3, 2.2, 1.1},
					},
					&Box{
						Type: TraceTypeBox,
						Name: "Simple Box",
						X:    []float64{1.1, 2.2, 3.3},
						Y:    []float64{3.3, 2.2, 1.1},
					},
				},
			},
		},
		{
			name: "with integer values",
			figure: Fig{
				Data: Traces{
					&Scatter{
						Type: TraceTypeScatter,
						Name: "Simple Scatter",
						X:    []int{1, 2, 3},
						Y:    []int{3, 2, 1},
					},
				},
			},
		},
		{
			name: "with Config And Layout",
			figure: Fig{
				Data: Traces{
					&Scatter{
						Type: TraceTypeScatter,
						Name: "Simple Scatter",
						X:    []float64{1.1, 2.2, 3.3},
						Y:    []float64{3.3, 2.2, 1.1},
					},
				},
				Config: &Config{
					Editable: True,
				},
				Layout: &Layout{
					Datarevision: "Date Revision 1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.figure)
			if err != nil {
				t.Errorf("Marshal should not fail, %s", err)
				return
			}
			result := Fig{}
			err = json.Unmarshal(data, &result)
			if err != nil {
				t.Errorf("Unmarshal should not fail, %s", err)
				return
			}
			marshalledResult, err := json.Marshal(result)
			if err != nil {
				t.Errorf("Marshal result should not fail, %s", err)
				return
			}
			if !bytes.Equal(data, marshalledResult) {
				t.Errorf("Marshal/Unmarshal figure modifies the content.\nbefore: %s\nafter:  %s\n", string(data), string(marshalledResult))
				return

			}
		})
	}
}
