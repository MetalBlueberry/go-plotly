package main

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	// Test case 1: Basic functionality
	reference := []string{"a", "b", "a", "b"}
	slices := [][]interface{}{
		{1, 2, 3, 4},
		{"s1", "s2", "s3", "s4"},
	}
	expected := map[string][][]interface{}{
		"a": {
			{1, 3},
			{"s1", "s3"},
		},
		"b": {
			{2, 4},
			{"s2", "s4"},
		},
	}

	result, _ := split(reference, slices)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 1 failed. Expected %v, got %v", expected, result)
	}

	// Test case 2: Single element in reference
	reference = []string{"a"}
	slices = [][]interface{}{
		{5},
		{"single"},
	}
	expected = map[string][][]interface{}{
		"a": {
			{5},
			{"single"},
		},
	}

	result, _ = split(reference, slices)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 2 failed. Expected %v, got %v", expected, result)
	}

	// Test case 3: Empty slices
	reference = []string{}
	slices = [][]interface{}{}
	expected = map[string][][]interface{}{}

	result, _ = split(reference, slices)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 3 failed. Expected %v, got %v", expected, result)
	}

	// Test case 4: Multiple same keys
	reference = []string{"x", "x", "y", "y"}
	slices = [][]interface{}{
		{10, 20, 30, 40},
		{"a", "b", "c", "d"},
	}
	expected = map[string][][]interface{}{
		"x": {
			{10, 20},
			{"a", "b"},
		},
		"y": {
			{30, 40},
			{"c", "d"},
		},
	}

	result, _ = split(reference, slices)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test case 4 failed. Expected %v, got %v", expected, result)
	}
}
