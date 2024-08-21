package arrayok

import (
	"encoding/json"
)

func Value[T any](value T) Type[*T] {
	v := &value
	return Type[*T]{Value: v}
}

func Array[T any](array ...T) Type[*T] {
	out := make([]*T, len(array))
	for i, v := range array {
		value := v
		out[i] = &value
	}
	return Type[*T]{
		Array: out,
	}
}

// Type is a type that allows you to define a single value or an array of values, But not both.
// If Array is defined, Value will be ignored.
type Type[T any] struct {
	Value T
	Array []T
}

func (arrayOK *Type[T]) MarshalJSON() ([]byte, error) {
	if arrayOK.Array != nil {
		return json.Marshal(arrayOK.Array)
	}
	return json.Marshal(arrayOK.Value)
}

func (arrayOK *Type[T]) UnmarshalJSON(data []byte) error {
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
