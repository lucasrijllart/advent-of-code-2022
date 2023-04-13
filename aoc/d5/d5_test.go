package main

import (
	"os"
	"reflect"
	"testing"
)

// ensure that the function catches assignments where the first set is encompassed by the second
// and the second encompasses the first. Also check for single-value sets.
func TestParseInput(t *testing.T) {
	file, _ := os.ReadFile("test.txt")
	containers, moves := parseInput(string(file))
	expectedContainers := [][]string{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}
	expectedMoves := [][]int{
		{1, 2, 1},
		{3, 1, 3},
		{2, 2, 1},
		{1, 1, 2},
	}

	if reflect.DeepEqual(containers, expectedContainers) != true {
		t.Errorf("Containers are not the same: %v, expected %v", containers, expectedContainers)
	}

	if reflect.DeepEqual(moves, expectedMoves) != true {
		t.Errorf("Moves are not the same: %v, expected %v", moves, expectedMoves)
	}
}

// ensure that the function runs the moves on the stacks of containers correctly
func TestMoveContainers(t *testing.T) {
	file, _ := os.ReadFile("test.txt")
	containers, moves := parseInput(string(file))
	movedContainers := MoveContainers(containers, moves)
	expectedContainers := [][]string{
		{"C"},
		{"M"},
		{"P", "D", "N", "Z"},
	}

	if reflect.DeepEqual(movedContainers, expectedContainers) != true {
		t.Errorf("Containers are not the same: %v, expected %v", movedContainers, expectedContainers)
	}
}
