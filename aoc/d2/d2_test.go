package main

import "testing"

func TestD2(t *testing.T) {
	tests := []struct {
		inputs   string
		expected int
	}{
		{
			inputs:   "A Y\nB X\nC Z",
			expected: 15,
		},
	}

	for _, test := range tests {
		result := RPSWin(test.inputs)
		if result != test.expected {
			t.Errorf("Got %d, expected %d", result, test.expected)
		}
	}
}

func TestScoreFromOneGame(t *testing.T) {
	tests := []struct {
		move     string
		response string
		expected int
	}{
		{move: "A", response: "Y", expected: 8},
		{move: "B", response: "X", expected: 1},
		{move: "C", response: "Z", expected: 6},
	}

	for _, test := range tests {
		result := ScoreFromOneGame(test.move, test.response)
		if result != test.expected {
			t.Errorf("Got %d, expected %d", result, test.expected)
		}
	}
}
