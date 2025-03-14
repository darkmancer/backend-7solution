package main

import (
	"reflect"
	"testing"
)

func TestCountMeatsTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:  "Original Sample",
			input: "Fatback t-bone T-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone.",
			expected: map[string]int{
				"fatback":  1,
				"t-bone":   4,
				"pastrami": 1,
				"pork":     1,
				"meatloaf": 1,
				"jowl":     1,
				"enim":     1,
				"bresaola": 1,
			},
		},
		{
			name:     "Empty String",
			input:    "",
			expected: map[string]int{},
		},
		{
			name:     "Single Word",
			input:    "Tenderloin",
			expected: map[string]int{"tenderloin": 1},
		},
		{
			name:     "Newlines and Tabs",
			input:    "tenderloin\ntenderloin, TENDERLOIN",
			expected: map[string]int{"tenderloin": 3},
		},
		{
			name:     "Numbers",
			input:    "123 123, 123",
			expected: map[string]int{"123": 3},
		},
		{
			name:     "Only Punctuation",
			input:    ".,;!?",
			expected: map[string]int{},
		},
		{
			name: "Long Input",
			input: "Steak ribeye t-bone T-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  " +
				"Bresaola t-bone. Tenderloin, steak, ribeye, TENDERLOIN, t-bone, ribeye.",
			expected: map[string]int{
				"steak":      2,
				"ribeye":     3,
				"t-bone":     5,
				"pastrami":   1,
				"pork":       1,
				"meatloaf":   1,
				"jowl":       1,
				"enim":       1,
				"bresaola":   1,
				"tenderloin": 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countMeats(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("For input %q, expected %v but got %v", tt.input, tt.expected, result)
			}
		})
	}
}
