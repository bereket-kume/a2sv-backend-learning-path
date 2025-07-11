package main

import "testing"

func TestCalculateAverage(t *testing.T) {
	tests := []struct {
		input    map[string]int
		expected float32
	}{
		{map[string]int{"Math": 90, "Sc": 80, "Eng": 70}, 80.0},
		{map[string]int{"His": 100, "Geo": 100}, 100.0},
		{map[string]int{"Phy": 50}, 50.0},
		{map[string]int{}, 0.0},
	}

	for _, test := range tests {
		result, _ := calculateAverage(test.input, len(test.input))

		if result != test.expected {
			t.Errorf("calculateAverage(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
