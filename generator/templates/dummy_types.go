package grob

import "encoding/json"

// Types defined here are just meant to allow compilation for the template package

type Layout struct{}
type Config struct{}

func UnmarshalTrace(json.RawMessage) (Trace, error) { return nil, nil }
