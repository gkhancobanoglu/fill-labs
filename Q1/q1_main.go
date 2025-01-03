package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Function to count the number of 'a' characters in a word
func countA(word string) int {
	count := 0
	for _, c := range word {
		if c == 'a' {
			count++
		}
	}
	return count
}

// Function to sort words first by the number of 'a's (in decreasing order),
// and then by length if the 'a' counts are the same
func sortWordsByA(words []string) []string {
	sort.SliceStable(words, func(i, j int) bool {
		countA_i := countA(words[i])
		countA_j := countA(words[j])
		// Sort by 'a' count, if equal, by length
		if countA_i == countA_j {
			return len(words[i]) > len(words[j])
		}
		return countA_i > countA_j
	})
	return words
}

func main() {
	// Default words list
	words := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}

	// Create a scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Display sorted words
		sortedWords := sortWordsByA(words)
		fmt.Println("Sorted Words:", sortedWords)

		// Ask for user input
		fmt.Print("Enter a new word (or type 'exit' to quit): ")
		scanner.Scan()
		input := scanner.Text()

		// Exit the loop if 'exit' is entered
		if input == "exit" {
			break
		}

		// Add the new word to the list
		words = append(words, input)
	}

	// Notify user about exiting
	fmt.Println("Exiting program...")
}
