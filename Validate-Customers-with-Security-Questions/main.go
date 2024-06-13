package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Sample database
	names := []string{"Ahmet", "Ayşe", "Mehmet"}
	surnames := []string{"Yılmaz", "Kaya", "Demir"}
	birthdays := []string{"01-01-1990", "02-02-1985", "03-03-1992"}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter customer's name: ")
	customerName, _ := reader.ReadString('\n')
	customerName = strings.TrimSpace(customerName)

	index := findIndex(names, customerName)
	if index == -1 {
		fmt.Println("Invalid name.")
		return
	}

	fmt.Print("Enter customer's surname: ")
	customerSurname, _ := reader.ReadString('\n')
	customerSurname = strings.TrimSpace(customerSurname)

	if customerSurname != surnames[index] {
		fmt.Println("Invalid surname.")
		return
	}

	fmt.Print("Enter customer's birthday (DD-MM-YYYY): ")
	customerBirthday, _ := reader.ReadString('\n')
	customerBirthday = strings.TrimSpace(customerBirthday)

	if customerBirthday != birthdays[index] {
		fmt.Println("Invalid birthday.")
		return
	}

	fmt.Println("Customer validated successfully!")
}

func findIndex(slice []string, value string) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}
