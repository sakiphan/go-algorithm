package main

import "fmt"

func main() {
	// Declare variables to store the numbers and the maximum found so far.
	var num, max int

	// Prompt user for input
	fmt.Println("Enter 5 numbers:")

	// Initialize max to the lowest possible integer for safety
	max = -1 << 31 // Assigns the minimum value for an int32

	// Read and compare each number
	for i := 0; i < 5; i++ {
		fmt.Scanln(&num)
		if num > max {
			max = num
		}
	}

	// Print the largest number
	fmt.Println("The largest number is:", max)
}
