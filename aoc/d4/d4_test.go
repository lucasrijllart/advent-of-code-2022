package main

import (
	"fmt"
	"testing"
)

// ensure that the function catches assignments where the first set is encompassed by the second
// and the second encompasses the first. Also check for single-value sets.
func TestCheckOverlap(t *testing.T) {
	tests := []struct {
		left     string
		right    string
		expected bool
	}{
		{"2-4", "6-8", false},
		{"2-8", "3-7", true},
		{"3-7", "2-8", true},
		{"6-6", "4-6", true},
		{"4-6", "6-6", true},
		{"7-7", "12-96", false},
	}

	for _, test := range tests {
		t.Run(test.left+","+test.right, func(t *testing.T) {
			result := CheckOverlap(test.left, test.right)
			if result != test.expected {
				t.Errorf("Got %v, expected %v", result, test.expected)
			}
		})
	}
}

// ensure that the function counts the number of lines that have ranges that fully overlap
func TestCountFullyContainedSections(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			input: []string{
				"2-4,6-8",
				"2-3,4-5",
				"5-7,7-9",
				"2-8,3-7",
				"6-6,4-6",
				"2-6,4-8",
			},
			expected: 2,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint(test.input), func(t *testing.T) {
			result := CountFullyContainedSections(test.input)
			if result != test.expected {
				t.Errorf("Got %v, expected %v", result, test.expected)
			}
		})
	}
}
