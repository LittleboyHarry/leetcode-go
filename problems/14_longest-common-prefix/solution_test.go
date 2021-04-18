package main

import (
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{""}, ""},
		{[]string{}, ""},
	}
	for id, test := range tests {
		t.Run("case "+strconv.Itoa(id+1), func(t *testing.T) {
			actual := longestCommonPrefix(test.input)
			if actual != test.expected {
				t.Errorf("input %v wanted %v but actual %v", test.input, test.expected, actual)
			}
		})
	}
}
