package main

import "testing"

func TestGetSign(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{10, 1},
		{-5, -1},
		{0, 0},
	}

	for _, test := range tests {
		result := getSign(test.input)
		if result != test.expected {
			t.Errorf("getSign(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}
