package main

import (
	"fmt"
	"math"
)

// Recursive function to generate the sequence
func reverseFnuc(input int) {
	var array []int

	// Inner recursive function
	var findRecursive func(input int)
	findRecursive = func(input int) {
		if input > 1 {
			array = append(array, input)
		}
		number := int(math.Floor(float64(input / 2)))
		if number > 1 {
			findRecursive(number)
		}
	}

	// Start recursion
	findRecursive(input)

	// Reverse the array and print values
	for i := len(array) - 1; i >= 0; i-- {
		fmt.Println(array[i])
	}
}

func main() {
	// Call the function with 9
	reverseFnuc(9)
}
