package offline

type Marker map[string]interface{}

func NewMarker() Marker {
	return Marker{}
}

func (m Marker) SetColor(color string) {
	m["color"] = color
}

func (m Marker) SetSize(size float32) {
	m["size"] = size
}
