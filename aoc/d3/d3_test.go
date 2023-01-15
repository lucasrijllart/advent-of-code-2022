package main

import (
	"fmt"
	"reflect"
	"testing"
)

// ensure the strings given get split in two and returned as
// the first and second items in a slice
func TestSplitString(t *testing.T) {
	tests := []struct {
		input  string
		output []string
	}{
		{
			input:  "abcd",
			output: []string{"ab", "cd"},
		},
		{
			input:  "vJrwpWtwJgWrhcsFMMfFFhFp",
			output: []string{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
		},
		{
			input:  "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			output: []string{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint(test.input), func(t *testing.T) {
			result := SplitString(test.input)
			if reflect.DeepEqual(result, test.output) == false {
				t.Errorf("Got %s, expected %s", result, test.output)
			}
		})
	}
}

// ensure the function returns the character that appears in both strings
func TestFindCommonItem(t *testing.T) {
	tests := []struct {
		input1, input2, expected string
	}{
		{
			input1:   "abc",
			input2:   "cde",
			expected: "c",
		},
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
		t.Run(fmt.Sprint(test.input1, test.input2), func(t *testing.T) {
			if result := FindCommonItem(test.input1, test.input2); result != test.expected {
				t.Errorf("Got %s, expected %s", result, test.expected)
			}
		})
	}
}

// ensure every letter in the alphabet is given the correct value
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
		t.Run(fmt.Sprint(test.input), func(t *testing.T) {
			result := GetItemPriority(test.input)
			if result != test.expected {
				t.Errorf("Got %d, expected %d", result, test.expected)
			}
		})
	}
}

// ensure the example rucksacks given equals the expected value
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
