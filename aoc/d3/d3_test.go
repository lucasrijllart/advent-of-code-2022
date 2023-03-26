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
			if result, _ := FindCommonItem(test.input1, test.input2); result != test.expected {
				t.Errorf("Got %s, expected %s", result, test.expected)
			}
		})
	}
}

// test exception creation by the function
func TestFindCommonItemError(t *testing.T) {
	tests := []struct {
		left, right string
	}{
		{
			"abc",
			"cde",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprint(test.left, test.right), func(t *testing.T) {
			if _, error := FindCommonItem(test.left, test.right); error != nil {
				t.Errorf("Got %v, expected %v", error, nil)
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
func TestGetRucksackDuplicateValue(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", 16},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", 38},
		{"PmmdzqPrVvPwwTWBwg", 42},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", 22},
		{"ttgJtRGJQctTZtZT", 20},
		{"CrZsJsPPZsGzwwsLwLmpwMDw", 19},
	}

	for _, test := range tests {
		result := GetRucksackDuplicateValue(test.input)
		if result != test.expected {
			t.Errorf("Got %d, expected %d", result, test.expected)
		}
	}
}

func TestGetCommonBadge(t *testing.T) {
	tests := []struct {
		first    string
		second   string
		third    string
		expected string
	}{
		{
			"vJrwpWtwJgWrhcsFMMfFFhFp",
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			"PmmdzqPrVvPwwTWBwg",
			"r",
		},
		{
			"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			"ttgJtRGJQctTZtZT",
			"CrZsJsPPZsGzwwsLwLmpwMDw",
			"Z",
		},
	}

	for _, test := range tests {
		result, _ := GetCommonBadge(test.first, test.second, test.third)
		if result != test.expected {
			t.Errorf("Got %v, expected %v", result, test.expected)
		}
	}
}
