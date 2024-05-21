package types

import (
    "encoding/json"
    "reflect"
)

type ArrayOK[T any] struct {
    Value T
    Array []T
}

func (arrayOK *ArrayOK[T]) MarshalJSON() ([]byte, error) {
    if arrayOK.Array!= nil {
        return json.Marshal(arrayOK.Array)
    }
    return json.Marshal(arrayOK.Value)
}

func (arrayOK *ArrayOK[T]) UnmarshalJSON(data []byte) error {
    arrayOK.Array = nil
    arrayOK.Value = reflect.Zero(reflect.TypeOf(arrayOK.Value)).Interface().(T)

    var array []T
    err := json.Unmarshal(data, &array)
    if err == nil {
        arrayOK.Array = array
        return nil
    }

    var value T
    err = json.Unmarshal(data, &value)
    if err!= nil {
        return err
    }
    arrayOK.Value = value
    return nil
}