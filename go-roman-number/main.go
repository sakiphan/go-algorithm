package main

import (
	"fmt"
	"strconv"
)

// toRoman converts an integer to Roman numerals
func toRoman(decimalNum int) string {
	// If the input number is not between 1 and 3999, it's marked as invalid input
	if decimalNum < 1 || decimalNum > 3999 {
		return "Invalid Input"
	}

	// Slice to hold Roman numerals and their corresponding values
	romanNums := []struct {
		value  int    // Value of the Roman numeral
		symbol string // Symbol of the Roman numeral
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	// Create a string for the result
	var result string

	// Loop to create symbols of Roman numerals
	for _, rn := range romanNums {
		// As long as the given number is greater than or equal to the value of the Roman numeral
		for decimalNum >= rn.value {
			// Add the symbol of the Roman numeral to the result
			result += rn.symbol
			// Subtract the value of the Roman numeral from the number
			decimalNum -= rn.value
		}
	}

	// Return the generated Roman numeral symbols
	return result
}

func main() {
	// Start an infinite loop, it continues until the user exits
	for {
		// Show information text to the user
		fmt.Print("Please enter a number between 1 and 3999 (Type 'exit' to quit): ")
		var input string
		// Read user input
		fmt.Scanln(&input)
		// If the user enters "exit", exit the program
		if input == "exit" {
			fmt.Println("\nExiting the program... Goodbye!")
			break
		}
		// Convert user input to an integer
		number, err := strconv.Atoi(input)
		// If there's an error or the converted number is not between 1 and 3999
		if err != nil {
			fmt.Println("Invalid input!")
			continue
		}
		// Show the user the Roman numerals equivalent of the converted number
		fmt.Printf("\nThe Roman numeral equivalent of %d: %s\n", number, toRoman(number))
	}
}
