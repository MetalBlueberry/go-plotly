package main

import (
	"math/rand"

	grob "github.com/MetalBlueberry/go-plotly/generated/v2.19.0/graph_objects"
	"github.com/MetalBlueberry/go-plotly/pkg/offline"
	"github.com/MetalBlueberry/go-plotly/pkg/types"
)

// Plotly.newPlot('myDiv', [{
//   x: [1, 2, 3],
//   y: [0, 0.5, 1],
//   line: {simplify: false},
// }]);

//	function randomize() {
//	  Plotly.animate('myDiv', {
//	    data: [{y: [Math.random(), Math.random(), Math.random()]}],
//	    traces: [0],
//	    layout: {}
//	  }, {
//	    transition: {
//	      duration: 500,
//	      easing: 'cubic-in-out'
//	    },
//	    frame: {
//	      duration: 500
//	    }
//	  })
//	}
func main() {

	frames := []grob.Frame{}
	// Because this example runs entirely in go, build a few random frames
	for i := 0; i < 20; i++ {
		frames = append(frames, grob.Frame{
			Data: []types.Trace{
				&grob.Scatter{
					Y: types.DataArray([]float64{rand.Float64(), rand.Float64(), rand.Float64()}),
				},
			},
		})
	}

	fig := &grob.Fig{
		Data: []types.Trace{
			&grob.Scatter{
				X: types.DataArray([]float64{1, 2, 3}),
				Y: types.DataArray([]float64{0, 0.5, 1}),
				Line: &grob.ScatterLine{
					Simplify: types.False,
				},
			},
		},
		Layout: &grob.Layout{
			Yaxis: &grob.LayoutYaxis{
				Range: []int{0, 1},
			},
			Updatemenus: []grob.LayoutUpdatemenu{
				{
					Type:       grob.LayoutUpdatemenuTypeButtons,
					Showactive: types.False,
					Buttons: []grob.LayoutUpdatemenuButton{
						{
							Label:  types.S("Play"),
							Method: grob.LayoutUpdatemenuButtonMethodAnimate,
							Args: []*ButtonArgs{
								nil,
								{
									Mode:        "immediate",
									FromCurrent: false,
								},
							},
						},
					},
				},
			},
		},
		Frames: frames,
		Animation: &grob.Animation{
			Transition: &grob.AnimationTransition{
				Duration: types.N(500),
				Easing:   grob.AnimationTransitionEasingCubicInOut,
			},
			Frame: &grob.AnimationFrame{
				Duration: types.N(500),
				Redraw:   types.True,
			},
		},
	}

	offline.Show(fig)
}

type ButtonArgs struct {
	Frame       map[string]interface{} `json:"frame,omitempty"`
	Transition  map[string]interface{} `json:"transition,omitempty"`
	FromCurrent bool                   `json:"fromcurrent,omitempty"`
	Mode        string                 `json:"mode,omitempty"`
}
