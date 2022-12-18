package main

import "testing"

func TestD1(t *testing.T) {
	tests := []struct {
		inputs   string
		expected int
	}{
		{
			inputs:   "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000",
			expected: 24000,
		},
		{
			inputs:   "1000\n\n2000\n\n3000\n\n4000",
			expected: 4000,
		},
	}

	for _, test := range tests {
		result := find_max_cal(test.inputs)
		if result != test.expected {
			t.Errorf("Got %d, expected %d", result, test.expected)
		}
	}
}

func TestD2(t *testing.T) {
	tests := []struct {
		inputs   string
		expected int
	}{
		{
			inputs:   "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000",
			expected: 45000,
		},
		{
			inputs:   "1000\n\n2000\n\n3000\n\n4000",
			expected: 9000,
		},
	}

	for _, test := range tests {
		result := find_top_three(test.inputs)
		if result != test.expected {
			t.Errorf("Got %d, expected %d", result, test.expected)
		}
	}
}
