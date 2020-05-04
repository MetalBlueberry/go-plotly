package main

import (
	"fmt"
	"math"
	"net/http"

	grob "github.com/MetalBlueberry/go-plotly/graph_objects"
	"github.com/MetalBlueberry/go-plotly/offline"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {

	request, err := http.Get("https://covid19.isciii.es/resources/serie_historica_acumulados.csv")
	if err != nil {
		panic(err)
	}

	df := dataframe.ReadCSV(request.Body)
	df = df.Filter(
		dataframe.F{"CASOS", series.Greater, 1000},
		dataframe.F{"PCR+", series.Greater, 00},
		dataframe.F{"TestAc+", series.Greater, 00},
	)
	fmt.Print(df.Names())
	clean := df.Select(
		[]string{"CASOS", "PCR+", "TestAc+"},
	).Rapply(func(in series.Series) series.Series {
		data := in.Float()
		if !math.IsNaN(data[0]) && (math.IsNaN(data[1]) || math.IsNaN(data[2])) {
			return series.Floats(data[0])
		}
		if math.IsNaN(data[0]) && !math.IsNaN(data[1]) && !math.IsNaN(data[2]) {
			return series.Floats(data[1] + data[2])
		}
		if math.IsNaN(data[0]) && !math.IsNaN(data[1]) {
			return series.Floats(data[1])
		}
		if math.IsNaN(data[0]) {
			return series.Floats(0)
		}
		return series.Floats(data[0])
	}).Rename("CASOS_FIX", "X0").Col("CASOS_FIX")
	df = df.Mutate(clean)
	fmt.Println(df)

	fig := &grob.Fig{}

	ccaas := unique(df.Col("CCAA").Records())
	fmt.Print(ccaas)
	for _, ccaa := range ccaas {
		fmt.Println(ccaa)
		dfl := df.Filter(
			dataframe.F{"CCAA", series.Eq, ccaa},
		).Arrange(
			dataframe.Sort("CASOS_FIX"),
		)
		fmt.Println(dfl)
		fmt.Println(dfl.Col("CASOS_FIX"))
		fmt.Println(dfl.Col("FECHA").Records())
		fig.Data = append(fig.Data, &grob.Bar{
			Type: grob.TraceTypeBar,
			Y:    dfl.Col("CASOS_FIX").Float(),
			X:    dfl.Col("FECHA").Records(),
			Name: ccaa,
		})
	}

	fig.Layout = &grob.Layout{
		Barmode: grob.LayoutBarmode_stack,
	}

	offline.Show(fig)

}

func unique(in []string) []string {
	u := make(map[string]bool)
	list := make([]string, 0)
	for i := 0; i < len(in); i++ {
		_, exist := u[in[i]]
		if !exist {
			u[in[i]] = true
			list = append(list, in[i])
		}
	}
	return list
}
