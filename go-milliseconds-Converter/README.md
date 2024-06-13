# Milliseconds Converter

### Installation
This Go program is designed to convert milliseconds into a human-readable format, expressing the time in hours, minutes, and seconds. It's a simple yet practical tool for those who need to make quick conversions from milliseconds to a more understandable time format. The program handles various edge cases and provides user-friendly messages for invalid inputs.
### Running the Program
To run the program, navigate to the directory containing your .go file in the terminal and use the following command:


```sh
go run main.go
```
### Usage
Upon running the program, you will be greeted with a message explaining what the program does and prompting you to enter a number of milliseconds. Here's what to expect:

```sh
(To exit the program, please type "exit")
Please enter the milliseconds (should be greater than zero):
```
Simply enter the number of milliseconds you wish to convert and press Enter. The program will then display the converted time in a human-readable format. If you enter an invalid input (e.g., a non-numeric value or a negative number), the program will notify you and prompt you to try again.
To exit the program, type exit when prompted for input.

### Examples

Here are some examples of user inputs and their corresponding outputs:

```plaintext
Input: 555
Output: just 555 millisecond/s

Input: 2001
Output: 2 second/s

Input: 3661011
Output: 1 hour/s 1 minute/s 1 second/s

Input: -8
Output: Not Valid Input !!!

Input: Exit
Output: Exiting the program... Good Bye
```


