package main

import (
	"log"
	"net/http"
	"sort"

	grob "github.com/MetalBlueberry/go-plotly/generated/v2.19.0/graph_objects"
	"github.com/MetalBlueberry/go-plotly/pkg/offline"
	"github.com/MetalBlueberry/go-plotly/pkg/types"
	"github.com/go-gota/gota/dataframe"
	"golang.org/x/exp/constraints"
)

// https://plotly.com/javascript/gapminder-example/

func readCSVData() dataframe.DataFrame {
	// country,year,pop,continent,lifeExp,gdpPercap
	response, err := http.Get("https://raw.githubusercontent.com/plotly/datasets/master/gapminderDataFiveYear.csv")
	if err != nil {
		log.Fatalf("Unable to fetch csv data, %s", err)
	}
	defer response.Body.Close()

	df := dataframe.ReadCSV(response.Body)

	if err != nil {
		log.Fatalf("Unable to import CSV data, %s", err)
	}
	return df
}

func main() {
	df := readCSVData()

	country := df.Col("country")
	year := df.Col("year")
	population := df.Col("pop")
	continent := df.Col("continent")
	lifeExp := df.Col("lifeExp")
	gdpPercap := df.Col("gdpPercap")

	continentClassifications, continentKey := split(continent.Records(), [][]string{
		year.Records(),
		country.Records(),
		population.Records(),
		lifeExp.Records(),
		gdpPercap.Records(),
	})

	indexYearContinent := make(map[string]map[string][][]string)
	for _, continent := range continentKey {
		continentClassification := continentClassifications[continent]
		year := continentClassification[0]
		yearClassification, yearKeys := split(year, continentClassification[1:])
		for _, year := range yearKeys {
			if indexYearContinent[year] == nil {
				indexYearContinent[year] = make(map[string][][]string)
			}
			indexYearContinent[year][continent] = yearClassification[year]
		}
	}

	frames := []grob.Frame{}
	sliderSteps := []grob.LayoutSliderStep{}

	years := SortedKeys(indexYearContinent)

	for _, year := range years {
		indexContinent := indexYearContinent[year]
		data := []types.Trace{}

		continents := SortedKeys(indexContinent)
		for _, continent := range continents {
			records := indexContinent[continent]
			data = append(data, &grob.Scatter{
				Name: types.S(continent),
				X:    types.DataArray(records[2]),                 // life expectancy
				Y:    types.DataArray(records[3]),                 // gdp per capita
				Ids:  types.DataArray(records[0]),                 // country
				Text: types.ArrayOKArray(types.SA(records[0])...), // country
				Marker: &grob.ScatterMarker{
					// Sizemode: grob.ScatterMarkerSizemodeArea,
					Size: types.ArrayOKArray(types.NSA(records[1])...), // population
					// Sizeref:  types.N(200000),
				},
			})
		}

		frameName := types.S(year)
		frames = append(frames, grob.Frame{
			Name: frameName,
			Data: data,
		})

		sliderSteps = append(sliderSteps, grob.LayoutSliderStep{
			Method: grob.LayoutSliderStepMethodAnimate,
			Label:  frameName,
			Args: []interface{}{
				[]interface{}{frameName},
				&ButtonArgs{
					Mode:       "immediate",
					Transition: map[string]interface{}{"duration": 300},
					Frame:      map[string]interface{}{"duration": 300, "redraw": false},
				},
			},
		})
	}

	indexContinent := indexYearContinent[years[0]]
	data := []types.Trace{}

	continents := SortedKeys(indexContinent)
	for _, continent := range continents {
		records := indexContinent[continent]
		data = append(data, &grob.Scatter{
			Name: types.S(continent),
			X:    types.DataArray(records[2]),                 // life expectancy
			Y:    types.DataArray(records[3]),                 // gdp per capita
			Ids:  types.DataArray(records[0]),                 // country
			Text: types.ArrayOKArray(types.SA(records[0])...), // country
			Mode: grob.ScatterModeMarkers,
			Marker: &grob.ScatterMarker{
				Sizemode: grob.ScatterMarkerSizemodeArea,
				Size:     types.ArrayOKArray(types.NSA(records[1])...), // population
				Sizeref:  types.N(200000),
			},
		})
	}

	fig := &grob.Fig{
		Data: data,
		Layout: &grob.Layout{
			Xaxis: &grob.LayoutXaxis{
				Title: &grob.LayoutXaxisTitle{
					Text: "Life Expectancy",
				},
				Range: []int{30, 85},
			},
			Yaxis: &grob.LayoutYaxis{
				Title: &grob.LayoutYaxisTitle{
					Text: "GDP per Capita",
				},
				Type: grob.LayoutYaxisTypeLog,
			},
			Hovermode: grob.LayoutHovermodeClosest,
			Updatemenus: []grob.LayoutUpdatemenu{
				{
					X:          types.N(0),
					Y:          types.N(0),
					Xanchor:    grob.LayoutUpdatemenuXanchorLeft,
					Yanchor:    grob.LayoutUpdatemenuYanchorTop,
					Showactive: types.False,
					Direction:  grob.LayoutUpdatemenuDirectionLeft,
					Type:       grob.LayoutUpdatemenuTypeButtons,
					Pad: &grob.LayoutUpdatemenuPad{
						T: types.N(87),
						R: types.N(10),
					},
					Buttons: []grob.LayoutUpdatemenuButton{
						{
							Label:  types.S("Play"),
							Method: grob.LayoutUpdatemenuButtonMethodAnimate,
							Args: []*ButtonArgs{
								nil,
								{
									Mode:        "immediate",
									FromCurrent: true,
									Transition:  map[string]interface{}{"duration": 300},
									Frame:       map[string]interface{}{"duration": 500, "redraw": false},
								},
							},
						},
						{
							Label:  types.S("Pause"),
							Method: grob.LayoutUpdatemenuButtonMethodAnimate,
							Args: []interface{}{
								[]interface{}{nil},
								&ButtonArgs{
									Mode:        "immediate",
									FromCurrent: true,
									Transition:  map[string]interface{}{"duration": 0},
									Frame:       map[string]interface{}{"duration": 0, "redraw": false},
								},
							},
						},
					},
				},
			},
			Sliders: []grob.LayoutSlider{
				{
					Pad: &grob.LayoutSliderPad{
						L: types.N(130),
						T: types.N(55),
					},
					Currentvalue: &grob.LayoutSliderCurrentvalue{
						Visible: types.True,
						Prefix:  types.S("Year:"),
						Xanchor: grob.LayoutSliderCurrentvalueXanchorRight,
						Font: &grob.LayoutSliderCurrentvalueFont{
							Size: types.N(20),
						},
					},
					Steps: sliderSteps,
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

	offline.Serve(fig)
}

type ButtonArgs struct {
	Frame       map[string]interface{} `json:"frame,omitempty"`
	Transition  map[string]interface{} `json:"transition,omitempty"`
	FromCurrent bool                   `json:"fromcurrent,omitempty"`
	Mode        string                 `json:"mode,omitempty"`
}

// given a reference slice, it will split the other slices in the same way
// so if reface is ["a","b","a","b"] and slices is [[1,2,3,4],["s1","s2","s3","s4"]]
// it will return
// {
// "a":[[1,3],["s1","s3"]],
// "b":[[2,4],["s2","s4"]]
// }
func split[T constraints.Ordered, Y any](reference []T, slices [][]Y) (map[T][][]Y, []T) {
	indices, keys := findIndices(reference)

	result := map[T][][]Y{}
	for i, slice := range slices {
		sections := splitByIndices(slice, indices)
		for j, key := range keys {
			if result[key] == nil {
				result[key] = make([][]Y, len(slices))
			}
			result[key][i] = sections[j]
		}
	}
	return result, keys
}

// given an slice, it will return the indices and the keys you can use to classify it by its types
// so ["a","b","a","b"] will return [[0,2],[1,3]] and ["a","b"]
func findIndices[T constraints.Ordered](input []T) ([][]int, []T) {
	indexMap := make(map[T][]int)
	var keys []T

	// Populate the map with indices grouped by the value
	for i, val := range input {
		if _, found := indexMap[val]; !found {
			keys = append(keys, val)
		}
		indexMap[val] = append(indexMap[val], i)
	}

	// Collect the grouped indices into a result slice in the order of first appearance
	var result [][]int
	for _, key := range keys {
		result = append(result, indexMap[key])
	}

	return result, keys
}

// given a slice, it will classify it by the given indices.
// so ["a","b","c","d"] with [[0,2],[1,3]] will return [["a","c"],["b","d"]]
func splitByIndices[T any](orginal []T, indices [][]int) [][]T {
	result := [][]T{}
	for i, section := range indices {
		result = append(result, []T{})
		for _, value := range section {
			result[i] = append(result[i], orginal[value])
		}
	}

	return result
}

func SortedKeys[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
