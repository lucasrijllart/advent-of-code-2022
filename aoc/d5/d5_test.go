package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

// ensure that the function catches assignments where the first set is encompassed by the second
// and the second encompasses the first. Also check for single-value sets.
func TestParseInput(t *testing.T) {
	file, _ := ioutil.ReadFile("test.txt")
	containers, moves := parseInput(string(file))
	expected_containers := [][]string{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}
	expected_moves := [][]int{
		{1, 2, 1},
		{3, 1, 3},
		{2, 2, 1},
		{1, 1, 2},
	}

	if reflect.DeepEqual(containers, expected_containers) != true {
		t.Errorf("Containers are not the same: %v, expected %v", containers, expected_containers)
	}

	if reflect.DeepEqual(moves, expected_moves) != true {
		t.Errorf("Containers are not the same: %v, expected %v", moves, expected_moves)
	}
}
