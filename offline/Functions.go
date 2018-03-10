package offline

import "fmt"

func addNumbericValue(trace map[string]interface{}, name string, x interface{}) error {
	switch xtype := x.(type) {
	case []int:
		if _, ok := trace[name]; !ok {
			trace[name] = make([]int, 0)
		}
		xvalues, ok := trace[name].([]int)
		if !ok {
			return fmt.Errorf("previous data in X doesn't match the type, %T vs %T", xtype, trace[name])
		}
		trace[name] = append(xvalues, xtype...)
	case []float32:
		if _, ok := trace[name]; !ok {
			trace[name] = make([]float32, 0)
		}
		xvalues, ok := trace[name].([]float32)
		if !ok {
			return fmt.Errorf("previous data in X doesn't match the type, %T vs %T", xtype, trace[name])
		}
		trace[name] = append(xvalues, xtype...)
	case []float64:
		if _, ok := trace[name]; !ok {
			trace[name] = make([]float64, 0)
		}
		xvalues, ok := trace[name].([]float64)
		if !ok {
			return fmt.Errorf("previous data in X doesn't match the type, %T vs %T", xtype, trace[name])
		}
		trace[name] = append(xvalues, xtype...)
	case int:
		return addNumbericValue(trace, name, []int{xtype})
	case float32:
		return addNumbericValue(trace, name, []float32{xtype})
	case float64:
		return addNumbericValue(trace, name, []float64{xtype})
	default:
		return fmt.Errorf("Data type is not valid, use int, float32 or float64")
	}
	return nil
}
