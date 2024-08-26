package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"

	"golang.org/x/exp/constraints"
)

type DataArrayTypeValues interface {
	constraints.Ordered
}

// DataArrayType An {array} of data.
// The value must represent an {array} or it will be ignored, but this array can be provided in several forms:
// (1) a regular {array} object
// (2) a typed array (e.g. Float32Array)
// (3) an object with keys dtype, bdata, and optionally shape.
// In this 3rd form, dtype is one of *f8*, *f4*. *i4*, *u4*, *i2*, *u2*, *i1*, *u1* or *u1c* for Uint8ClampedArray.
// In addition to shorthand `dtype` above one could also use the following forms: *float64*, *float32*, *int32*, *uint32*, *int16*, *uint16*, *int8*, *uint8* or *uint8c*  for Uint8ClampedArray.
// `bdata` is either a base64-encoded string or the ArrayBuffer of an integer or float typed array.
// For either multi-dimensional arrays you must also provide its dimensions separated by comma via `shape`.
// For example using `dtype`: *f4* and `shape`: *5,100* you can declare a 2-D array that has 5 rows and 100 columns containing
// float32 values i.e. 4 bits per value. `shape` is optional for one dimensional arrays.",
type DataArrayType struct {
	// Option 1
	values interface{}

	// Option 2 // doesn't apply in go as option one already has types

	// Option 3
	dtype DType
	// Base64 encoded data
	bdata []byte
	shape DataArrayShape
}

type BDataArrayType struct {
	Dtype DType          `json:"dtype,omitempty"`
	Bdata []byte         `json:"bdata,omitempty"`
	Shape DataArrayShape `json:"shape,omitempty"`
}

// Value returns either the underlying interface values OR the BDataArrayType, depending on how the DataArray was created.
func (da *DataArrayType) Value() interface{} {
	if da.values != nil {
		return da.values
	}
	if da.bdata != nil {
		return BDataArrayType{
			Dtype: da.dtype,
			Bdata: da.bdata,
			Shape: da.shape,
		}
	}
	return nil
}

type DataArrayShape []int

type DType string

const (
	DTypeFloat64 DType = "float64"
	DTypeFloat32 DType = "float32"
	DTypeInt32   DType = "int32"
	DTypeUInt32  DType = "uint32"
	DTypeInt16   DType = "int16"
	DTypeUInt16  DType = "uint16"
	DTypeInt8    DType = "int8"
	DTypeUInt8   DType = "uint8"
	DTypeUInt8c  DType = "uint8c"
)

func DataArray[T any](data []T) *DataArrayType {
	return &DataArrayType{
		values: data,
	}
}

func BDataArray(dtype DType, bdata []byte, shape []int) *DataArrayType {
	return &DataArrayType{
		dtype: dtype,
		bdata: bdata,
		shape: shape,
	}
}

func (d *DataArrayType) MarshalJSON() ([]byte, error) {
	if d.values != nil {
		// If Values is not nil, marshal it directly
		return json.Marshal(d.values)
	}

	return json.Marshal(BDataArrayType{
		Dtype: d.dtype,
		Bdata: d.bdata,
		Shape: d.shape,
	})
}

func (d *DataArrayType) UnmarshalJSON(data []byte) error {

	bData := &BDataArrayType{}
	err := json.Unmarshal(data, bData)
	if err == nil {
		d.bdata = bData.Bdata
		d.dtype = bData.Dtype
		d.shape = bData.Shape
		return nil
	}

	values, err := unmarshalToGenericSlice(data)
	if err != nil {
		return err
	}
	d.values = values
	return nil
}

// unmarshalToGenericSlice unmarshals JSON data into a generic slice.
// It returns the data as a slice of interface{} and its detected type.
func unmarshalToGenericSlice(data []byte) (interface{}, error) {
	var genericSlice []interface{}
	if err := json.Unmarshal(data, &genericSlice); err != nil {
		return nil, err
	}

	if len(genericSlice) == 0 {
		return genericSlice, nil
	}

	// Determine the type of elements in the slice
	elemType := reflect.TypeOf(genericSlice[0])
	return convertToType(genericSlice, elemType.Kind())
}

// Convert a slice of interface{} to a slice of the detected type
func convertToType(slice []interface{}, elemType reflect.Kind) (interface{}, error) {
	switch elemType {
	// I believe json package will never use this type, but will use float64
	case reflect.Int:
		intSlice := make([]int, len(slice))
		for i, v := range slice {
			if v, ok := v.(float64); ok { // JSON numbers are float64 by default
				intSlice[i] = int(v)
			} else {
				return nil, fmt.Errorf("unexpected type %T for int", v)
			}
		}
		return intSlice, nil
	case reflect.Float64:
		floatSlice := make([]float64, len(slice))
		for i, v := range slice {
			if v, ok := v.(float64); ok {
				floatSlice[i] = v
			} else {
				return nil, fmt.Errorf("unexpected type %T for float64", v)
			}
		}
		return floatSlice, nil
	case reflect.String:
		stringSlice := make([]string, len(slice))
		for i, v := range slice {
			if v, ok := v.(string); ok {
				stringSlice[i] = v
			} else {
				return nil, fmt.Errorf("unexpected type %T for string", v)
			}
		}
		return stringSlice, nil
	default:
		return nil, fmt.Errorf("unsupported type %s", elemType)
	}
}

func (d DataArrayShape) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal([]int(d))
	if err != nil {
		return b, err
	}
	// make data compatible with what plotty expects
	b = bytes.Replace(b, []byte("["), []byte("\""), -1)
	b = bytes.Replace(b, []byte("]"), []byte("\""), -1)

	return b, nil
}

func (d *DataArrayShape) UnmarshalJSON(data []byte) error {
	data = bytes.Replace(data, []byte("\""), []byte("["), 1)
	data = bytes.Replace(data, []byte("\""), []byte("]"), 1)

	intSlice := []int{}
	err := json.Unmarshal(data, &intSlice)
	if err != nil {
		return err
	}
	(*d) = intSlice
	return nil
}
