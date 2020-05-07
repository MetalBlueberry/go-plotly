package graph_objects

//go:generate sh -c "cat schema.json | plate --template main.tmpl"

// Plotly Schema SHA ccc513331344efa777630dc699af5ecc8a8c208a

type Traces []Trace

type Fig struct {
	Data   Traces  `json:"data,omitempty"`
	Layout *Layout `json:"layout,omitempty"`
	Config *Config `json:"config,omitempty"`
}

func (fig *Fig) AddTraces(traces ...Trace) {
	fig.Data = append(fig.Data, traces...)
}

// This section is to workaround the omitempty for json serialization.

type Bool *bool

var (
	trueValue  bool = true
	falseValue bool = false

	True  Bool = &trueValue
	False Bool = &falseValue
)

type String interface{}
