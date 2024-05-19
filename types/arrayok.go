package types

import "encoding/json"

type ArrayOKfloat64 struct {
	Value float64
	Array []float64
}

func (arrayOK *ArrayOKfloat64) MarshalJSON() ([]byte, error) {
	if arrayOK.Array != nil {
		return json.Marshal(arrayOK.Array)
	}
	return json.Marshal(arrayOK.Value)
}

func (arrayOK *ArrayOKfloat64) UnmarshalJSON(data []byte) error {
	arrayOK.Array = nil
	arrayOK.Value = 0

	var array []float64
	err := json.Unmarshal(data, &array)
	if err == nil {
		arrayOK.Array = array
		return nil
	}

	var value float64
	err = json.Unmarshal(data, &value)
	if err != nil {
		return err
	}
	arrayOK.Value = value
	return nil
	return nil

}
