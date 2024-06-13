# Check Consecutive Vowels

This program checks if a given string contains consecutive vowels. It takes a string as input from the user and outputs "Positive" if the string contains consecutive vowels and "Negative" otherwise.

## Learning Outcomes

By completing this coding challenge, you will be able to:

- Analyze a problem, identify, and apply programming knowledge for an appropriate solution.
- Implement conditional statements effectively to solve a problem.
- Implement loops to solve a problem.
- Execute operations on strings.
- Demonstrate knowledge of algorithmic design principles by solving the problem effectively.

## Problem Statement

In this coding challenge, you will write a program that takes a string and checks if it contains consecutive vowels or not. It should output `Positive` if it contains consecutive vowels and `Negative` otherwise. 

### Examples

1. For the string `saetqi`, the output should be `Positive` because it contains `a` adjacent to `e`.
2. For the string `statoqag`, the output should be `Negative` because it does not contain any consecutive vowels.

### Expected Output

```text
Please enter a string: gasdph
Negative

Please enter a string: aiou
Positive

Please enter a string: taoum
Positive

Please enter a string: a
Negative
```

## Implementation

Here is the complete implementation in Go:

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to check if a character is a vowel
func isVowel(c byte) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, rune(c))
}

// Function to check for consecutive vowels
func hasConsecutiveVowels(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if isVowel(s[i]) && isVowel(s[i+1]) {
			return true
		}
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter a string: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if hasConsecutiveVowels(input) {
		fmt.Println("Positive")
	} else {
		fmt.Println("Negative")
	}
}
```

### Explanation

1. **isVowel Function**: This function checks if a given character is a vowel. It uses `strings.ContainsRune` to determine if the character is in the string of vowels (`"aeiouAEIOU"`).
2. **hasConsecutiveVowels Function**: This function iterates through the input string and checks for any consecutive vowels using the `isVowel` function.
3. **main Function**:
   - Reads a string from the user using `bufio.NewReader`.
   - Trims any leading or trailing whitespace or newline characters from the input string.
   - Calls the `hasConsecutiveVowels` function to check for consecutive vowels.
   - Prints "Positive" if consecutive vowels are found, otherwise prints "Negative".

### How to Run the Program

1. Ensure you have Go installed on your system.
2. Save the code to a file named `main.go`.
3. Open a terminal and navigate to the directory containing `main.go`.
4. Run the program using the command: `go run main.go`.
5. Enter a string when prompted, and the program will output whether the string contains consecutive vowels or not.