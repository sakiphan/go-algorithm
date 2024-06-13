package main

import "fmt"

func isValid(s string) bool {
	matching := map[rune]rune{')': '(', '}': '{', ']': '['}
	var stack []rune

	for _, char := range s {
		if closing, ok := matching[char]; ok {
			var top rune
			if len(stack) > 0 {
				top = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			if top != closing {
				return false
			}
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}

func main() {
	var input string
	fmt.Println("Enter a string of brackets to validate:")
	fmt.Scanln(&input) // Taking input from the user

	result := isValid(input)
	fmt.Printf("Input: %s\tOutput: %v\n", input, result)
}
