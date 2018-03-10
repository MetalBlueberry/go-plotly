package offline

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Figure map[string]interface{}

func NewFigure() Figure {
	return Figure{}
}

//GetTraceList return a list with all the traces inside the figure
func (fig Figure) GetTraceList() []interface{} {
	if _, ok := fig["data"]; !ok {
		fig["data"] = make([]interface{}, 0)
	}
	return fig["data"].([]interface{})
}

// AddTrace  Adds the trace to the figure, is easier to use function to directly create charts, like NewScatter, NewGantt....
func (fig Figure) AddTrace(trace interface{}) error {
	data := fig.GetTraceList()
	switch v := trace.(type) {
	case Scatter:
		fig["data"] = append(data, v)
		return nil
	default:
		return errors.New("Trace type not valid")
	}
}

func (fig Figure) NewScatter() (Scatter, error) {
	s := NewScatter()
	return s, fig.AddTrace(s)
}

// NewGantt will create a new Gantt Figure and update the Conf and Layout of the figure to properly show Gantt charts.
func (fig Figure) NewGantt() (Gantt, error) {
	g, l, c := NewGantt()
	err := fig.AddTrace(g)
	fig.UpdateConfig(c)
	fig.UpdateLayout(l)
	return g, err
}
func (fig Figure) NewLayout() Layout {
	layout := NewLayout()
	fig["layout"] = layout
	return layout
}
func (fig Figure) NewConfig() Config {
	config := NewConfig()
	fig["config"] = config
	return config
}

func (fig Figure) SetLayout(l Layout) {
	fig["layout"] = l
}

func (fig Figure) SetConfig(c Config) {
	fig["Config"] = c
}

func (fig Figure) UpdateLayout(l Layout) {
	if _, ok := fig["layout"]; !ok {
		fig.SetLayout(l)
	} else {
		for k, v := range l {
			fig["layout"].(map[string]interface{})[k] = v
		}
	}
}

func (fig Figure) UpdateConfig(c Config) {
	if _, ok := fig["config"]; !ok {
		fig.SetConfig(c)
	} else {
		for k, v := range c {
			fig["config"].(map[string]interface{})[k] = v
		}
	}
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
