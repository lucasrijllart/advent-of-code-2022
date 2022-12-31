package main

import "testing"

func TestD2P1(t *testing.T) {
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
		result := GoThroughFile(test.inputs, ScoreFromOneGame)
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

func TestD2P2(t *testing.T) {
	tests := []struct {
		inputs   string
		expected int
	}{
		{
			inputs:   "A Y\nB X\nC Z",
			expected: 12,
		},
	}

	for _, test := range tests {
		result := GoThroughFile(test.inputs, ScoreFromRoundEnd)
		if result != test.expected {
			t.Errorf("Got %d, expected %d", result, test.expected)
		}
	}
}

func TestRoundEnd(t *testing.T) {
	tests := []struct {
		move     string
		response string
		expected int
	}{
		{move: "A", response: "Y", expected: 4},
		{move: "B", response: "X", expected: 1},
		{move: "C", response: "Z", expected: 7},
	}

	for _, test := range tests {
		result := ScoreFromRoundEnd(test.move, test.response)
		if result != test.expected {
			t.Errorf("Got %d, expected %d", result, test.expected)
		}
	}
}
