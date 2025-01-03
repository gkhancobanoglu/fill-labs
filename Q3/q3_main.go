package main

import (
	"bufio"
	"fmt"
	"os"
)

// Function to find the most repeated element in the array
func findMostRepeated(data []string) string {
	// Create a map to store the frequency of each element
	frequency := make(map[string]int)

	// Iterate through the array and update the frequency of each element
	for _, item := range data {
		frequency[item]++
	}

	// Initialize variables to track the most repeated element
	var mostRepeated string
	maxCount := 0

	// Iterate through the map to find the element with the highest frequency
	for item, count := range frequency {
		if count > maxCount {
			mostRepeated = item
			maxCount = count
		}
	}

	// Return the most repeated element
	return mostRepeated
}

func main() {
	// Default data
	data := []string{"apple", "pie", "apple", "red", "red", "red"}

	// Create a scanner to get input from the user
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Find and print the most repeated element
		result := findMostRepeated(data)
		fmt.Printf("Most repeated element: %s\n", result)

		// Prompt user for input
		fmt.Print("Enter a new word (or type 'exit' to quit): ")
		scanner.Scan()
		input := scanner.Text()

		// Exit condition
		if input == "exit" {
			break
		}

		// Add the new word to the data
		data = append(data, input)
	}

	fmt.Println("Exiting program...")
}
