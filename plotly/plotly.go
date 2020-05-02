package plotly

//go:generate sh -c "cat schema.json | plate --template main.tmpl"

// Plotly Schema SHA ccc513331344efa777630dc699af5ecc8a8c208a

type Traces []Trace


type Fig struct {
	Data   Traces `json:"data,omitempty"`
	Layout Layout `json:"layout,omitempty"`
}