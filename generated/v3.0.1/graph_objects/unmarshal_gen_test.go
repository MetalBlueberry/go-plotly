package grob

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestUnmarshalTrace(t *testing.T) {
	data := `{"type":"bar"}`
	trace, err := UnmarshalTrace([]byte(data))
	if err != nil {
		t.Logf("Error during Unmarshal, %s", err)
		t.FailNow()
	}
	if traceType := trace.GetType(); traceType != "bar" {
		t.Logf("Error recovering type, Expected \"bar\", got %s", traceType)
		t.FailNow()
	}
}

func TestMarshalTrace(t *testing.T) {
	bar := &Bar{}
	v, err := json.Marshal(bar)
	if err != nil {
		t.Logf("Error during Marshal, %s", err)
		t.FailNow()
	}
	if !strings.Contains(string(v), "type") {
		t.Logf("Output doesn't contain the type field, Expected to contain \"type\", got: %s", v)
		t.FailNow()
	}
	if strings.Contains(string(v), "null") {
		t.Logf("Output contains null fields, Expected to not contain any, got: %s", v)
		t.FailNow()
	}
}
