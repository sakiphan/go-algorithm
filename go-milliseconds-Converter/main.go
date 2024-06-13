package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Function to convert milliseconds to hours, minutes, and seconds.
func convertMillis(millis int) string {
	if millis < 1000 {
		return fmt.Sprintf("just %d millisecond/s", millis)
	}

	seconds := millis / 1000
	minutes := seconds / 60
	hours := minutes / 60

	seconds %= 60
	minutes %= 60

	var result strings.Builder

	if hours > 0 {
		result.WriteString(fmt.Sprintf("%d hour/s ", hours))
	}
	if minutes > 0 {
		result.WriteString(fmt.Sprintf("%d minute/s ", minutes))
	}
	if seconds > 0 {
		result.WriteString(fmt.Sprintf("%d second/s", seconds))
	}

	return strings.TrimSpace(result.String())
}

func main() {
	fmt.Println("### This program converts milliseconds into hours, minutes, and seconds ###\n(To exit the program, please type \"exit\")")

	for {
		fmt.Print("Please enter the milliseconds (should be greater than zero): ")
		var input string
		fmt.Scanln(&input)

		input = strings.TrimSpace(input)
		if strings.EqualFold(input, "exit") {
			fmt.Println("Exiting the program... Good Bye")
			break
		}

		millis, err := strconv.Atoi(input)
		if err != nil || millis < 1 {
			fmt.Println("Not Valid Input !!!")
			continue
		}

		fmt.Println(convertMillis(millis))
	}
}
