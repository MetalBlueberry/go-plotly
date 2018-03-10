package offline

type Bar map[string]interface{}

func NewBar() Bar {
	return Bar{
		"type": "bar",
	}
}
