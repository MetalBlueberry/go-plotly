package types

import (
	"encoding/json"
)

func ArrayOKValue[T any](value T) ArrayOK[*T] {
	v := &value
	return ArrayOK[*T]{Value: v}
}

func ArrayOKArray[T any](array ...T) ArrayOK[*T] {
	out := make([]*T, len(array))
	for i, v := range array {
		value := v
		out[i] = &value
	}
	return ArrayOK[*T]{
		Array: out,
	}
}

// ArrayOK is a type that allows you to define a single value or an array of values, But not both.
// If Array is defined, Value will be ignored.
type ArrayOK[T any] struct {
	Value T
	Array []T
}

func (arrayOK *ArrayOK[T]) MarshalJSON() ([]byte, error) {
	if arrayOK.Array != nil {
		return json.Marshal(arrayOK.Array)
	}
	return json.Marshal(arrayOK.Value)
}

func (arrayOK *ArrayOK[T]) UnmarshalJSON(data []byte) error {
	arrayOK.Array = nil

	var array []T
	err := json.Unmarshal(data, &array)
	if err == nil {
		arrayOK.Array = array
		return nil
	}

	var value T
	err = json.Unmarshal(data, &value)
	if err != nil {
		return err
	}
	arrayOK.Value = value
	return nil
}
