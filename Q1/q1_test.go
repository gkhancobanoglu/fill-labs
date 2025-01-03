package main

import (
	"testing"
)

func TestSortWordsByA(t *testing.T) {
	words := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
	expected := []string{"aaaasd", "aaabcd", "aab", "a", "lklklklklklklklkl", "cssssssd", "fdz", "ef", "kf", "zc", "l"}
	sortedWords := sortWordsByA(words)

	// Check if the sorted words match the expected result
	for i, word := range sortedWords {
		if word != expected[i] {
			t.Errorf("Expected %v, but got %v", expected[i], word)
		}
	}
}
