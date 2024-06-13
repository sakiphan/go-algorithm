# Bracket Validator in Go

This Go program checks if a given string of brackets is valid according to the rules of properly closed and nested brackets.

## Features

- Validates strings containing the characters `(`, `)`, `{`, `}`, `[`, `]`.
- Checks if every opening bracket has a corresponding closing bracket.
- Ensures that the brackets are closed in the correct order.

## Requirements

To run this program, you will need:
- Go (version 1.14 or higher recommended)

## Installation

Clone the repository or download the source code to your local machine.

## Usage

To run the program, navigate to the directory containing the source file and use the following command:

```bash
go run main.go
```
Once the program is running, it will prompt you to enter a string of brackets. Type the string and press Enter. The program will then output whether the string is valid or not.

## Examples

Input and Output examples:

| Input    | Output |
|----------|--------|
| `"()"`       | true   |
| `"()[]{}"`   | true   |
| `"(]"`       | false  |
| `"([)]"`     | false  |
| `"{[]}"`     | true   |
| `""`         | true   |
These examples demonstrate various scenarios, including nested and sequential brackets, as well as an empty string.

