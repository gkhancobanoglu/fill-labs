package main

import (
	"testing"
)

// Test function for findMostRepeated
func TestFindMostRepeated(t *testing.T) {
	// Sample test cases
	tests := []struct {
		data     []string
		expected string
	}{
		{[]string{"apple", "pie", "apple", "red", "red", "red"}, "red"},
		{[]string{"apple", "apple", "apple"}, "apple"},
		{[]string{"red", "blue", "green"}, "red"},
	}

	// Loop through test cases
	for _, test := range tests {
		t.Run("Testing Most Repeated", func(t *testing.T) {
			result := findMostRepeated(test.data)
			if result != test.expected {
				t.Errorf("Expected %s but got %s", test.expected, result)
			}
		})
	}
}
