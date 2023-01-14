package main

import "testing"

func TestSplitString(t *testing.T) {
	tests := []struct {
		input   string
		output1 string
		output2 string
	}{
		{
			input:   "vJrwpWtwJgWrhcsFMMfFFhFp",
			output1: "vJrwpWtwJgWr",
			output2: "hcsFMMfFFhFp",
		},
		{
			input:   "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			output1: "jqHRNqRjqzjGDLGL",
			output2: "rsFMfFZSrLrFZsSL",
		},
	}

	for _, test := range tests {
		result := SplitString(test.input)
		if result[0] != test.output1 {
			t.Errorf("Got %s, expected %s", result, test.output1)
		}
		if result[1] != test.output2 {
			t.Errorf("Got %s, expected %s", result, test.output2)
		}
	}
}

func TestFindCommonItem(t *testing.T) {
	tests := []struct {
		input1   string
		input2   string
		expected string
	}{
		{
			input1:   "vJrwpWtwJgWr",
			input2:   "hcsFMMfFFhFp",
			expected: "p",
		},
		{
			input1:   "jqHRNqRjqzjGDLGL",
			input2:   "rsFMfFZSrLrFZsSL",
			expected: "L",
		},
	}

	for _, test := range tests {
		result := FindCommonItem(test.input1, test.input2)
		if result != test.expected {
			t.Errorf("Got %s, expected %s", result, test.expected)
		}
	}
}

func TestGetItemPriority(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{input: "p", expected: 16},
		{input: "L", expected: 38},
		{input: "P", expected: 42},
		{input: "v", expected: 22},
		{input: "t", expected: 20},
		{input: "Z", expected: 52},
	}

	for _, test := range tests {
		result := GetItemPriority(test.input)
		if result != test.expected {
			t.Errorf("Got %d, expected %d", result, test.expected)
		}
	}
}

func TestGetTotalItemValue(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			input: []string{
				"vJrwpWtwJgWrhcsFMMfFFhFp",
				"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
				"PmmdzqPrVvPwwTWBwg",
				"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
				"ttgJtRGJQctTZtZT",
				"CrZsJsPPZsGzwwsLwLmpwMDw",
			},
			expected: 157,
		},
	}

	for _, test := range tests {
		result := GetTotalItemValue(test.input)
		if result != test.expected {
			t.Errorf("Got %d, expected %d", result, test.expected)
		}
	}
}
