package offlinePlotly

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Figure map[string]interface{}

func NewFigure() Figure {
	return Figure{}
}

func (fig Figure) getData() []interface{} {
	if _, ok := fig["data"]; !ok {
		fig["data"] = make([]interface{}, 0)
	}
	return fig["data"].([]interface{})
}

func (fig Figure) AddTrace(trace interface{}) error {
	data := fig.getData()
	switch v := trace.(type) {
	case Trace:
		fig["data"] = append(data, v)
		return nil
	default:
		return errors.New("Trace type not valid")
	}
}

func (fig Figure) NewLayout() Layout {
	layout := NewLayout()
	fig["layout"] = layout
	return layout
}

// GetJSON is same as json.Marshall
func (fig Figure) GetJSON() ([]byte, error) {
	return json.Marshal(fig)
}

func (fig Figure) GetPlotlySnippet(divId string) ([]byte, error) {
	data, err := fig.GetJSON()
	if err != nil {
		return nil, err
	}
	snippet := fmt.Sprintf("Plotly.newPlot(%s, %s);", divId, data)
	return []byte(snippet), nil
}
