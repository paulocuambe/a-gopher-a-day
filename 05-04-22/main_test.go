package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type test struct {
		input    [2]int
		expected int
	}
	inp := []test{
		{
			input:    [2]int{1, 2},
			expected: 3,
		},
		{
			input:    [2]int{4, 5},
			expected: 9,
		},
	}

	for _, v := range inp {
		if v.input[0]+v.input[1] != v.expected {
			t.Errorf("%d + %d != %d", v.input[0], v.input[1], v.expected)
		}
	}
}
