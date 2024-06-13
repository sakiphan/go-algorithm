package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	phonebook := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the phonebook application")
	for {
		fmt.Println("1. Find phone number")
		fmt.Println("2. Insert a phone number")
		fmt.Println("3. Delete a person from the phonebook")
		fmt.Println("4. Terminate")
		fmt.Print("Select operation on Phonebook App (1/2/3/4) : ")

		if !scanner.Scan() {
			continue
		}
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Find the phone number of : ")
			scanner.Scan()
			name := scanner.Text()
			if number, found := phonebook[name]; found {
				fmt.Println(number)
			} else {
				fmt.Printf("Couldn't find phone number of %s\n", name)
			}
		case "2":
			fmt.Print("Insert name of the person : ")
			scanner.Scan()
			name := scanner.Text()
			fmt.Print("Insert phone number of the person: ")
			scanner.Scan()
			number := scanner.Text()
			if _, found := phonebook[name]; !found {
				phonebook[name] = number
				fmt.Printf("Phone number of %s is inserted into the phonebook\n", name)
			} else {
				fmt.Println("That person already exists in the phonebook.")
			}
		case "3":
			fmt.Print("Whom to delete from phonebook : ")
			scanner.Scan()
			name := scanner.Text()
			if _, found := phonebook[name]; found {
				delete(phonebook, name)
				fmt.Printf("%s is deleted from the phonebook\n", name)
			} else {
				fmt.Printf("Couldn't find %s in phonebook\n", name)
			}
		case "4":
			fmt.Println("Exiting Phonebook")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
