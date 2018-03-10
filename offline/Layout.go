package offlinePlotly

type Layout map[string]interface{}

func NewLayout() Layout {
	return Layout{}
}

func (layout Layout) SetTitle(title string) {
	layout["titile"] = title
}
