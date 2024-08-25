package main

import (
	"encoding/json"
	"os/exec"

	grob "github.com/MetalBlueberry/go-plotly/generated/v2.31.1/graph_objects"
	"github.com/MetalBlueberry/go-plotly/pkg/offline"
	"github.com/MetalBlueberry/go-plotly/pkg/types"
)

func main() {
	// https://plotly.com/javascript/heatmaps/
	// var data = [
	//   {
	//     z: [[1, null, 30, 50, 1], [20, 1, 60, 80, 30], [30, 60, 1, -10, 20]],
	//     x: ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday'],
	//     y: ['Morning', 'Afternoon', 'Evening'],
	//     type: 'heatmap',
	//     hoverongaps: false
	//   }
	// ];
	zJSON, err := exec.Command("make").Output()
	if err != nil {
		panic(err)
	}
	type bdata struct {
		DType types.DType
		Shape []int
		BData []byte
	}

	z := &bdata{}
	err = json.Unmarshal(zJSON, z)
	if err != nil {
		panic(err)
	}

	fig := &grob.Fig{
		Data: []types.Trace{
			&grob.Heatmap{
				X:           types.DataArray([]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}),
				Y:           types.DataArray([]string{"Morning", "Afternoon", "Evening"}),
				Z:           types.BDataArray(z.DType, z.BData, z.Shape),
				Hoverongaps: types.False,
			},
		},
	}

	offline.Show(fig)

}
