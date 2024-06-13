package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isVowel(c byte) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, rune(c))
}

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
