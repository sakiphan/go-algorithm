# Customer Validation with Security Questions

This is a simple Go program that validates a customer by asking security questions. It checks the customer's name, surname, and birthday against a predefined database.

## Learning Outcomes

By completing this coding challenge, you will:

- Analyze a problem, identify, and apply programming knowledge for an appropriate solution.
- Implement conditional statements effectively to solve a problem.
- Demonstrate knowledge of algorithmic design principles by solving the problem effectively.

## Problem Statement

Companies use security questions to verify if a customer is valid and exists in their database before executing any operations requested by the customer (e.g., telecommunication companies, banks, etc.).

In this coding challenge, you will create a security questions program to assist the customer representative with questions which should be answered by the customer.

The customer should be checked for **customer's name**, **customer's surname**, and **customer's birthday**.

The database is given in Python lists. The matching name, surname, and birthday combinations are in the same indexes of the lists. So if the customer's name is in the first index of the name list, his surname and birthday are also going to be in the first indexes of the surname and birthday lists.

You should start by asking for the customer's name, then continue with the customer's surname and birthday respectively.

Each time you get the information from the customer, you should check if that information is valid or matches with existing records.

If the input is not valid or the combination that is provided does not match, the program should give an appropriate message and terminate without asking the rest of the questions.

If all the information is valid and matches the pattern, the program should give an appropriate validation message.

## Usage

### Prerequisites

- Go programming environment

### Running the Program

1. Clone the repository or download the program file.
2. Open a terminal and navigate to the directory containing the program file.
3. Run the following command to execute the program:

   ```bash
   go run main.go
   ```

4. Follow the prompts to enter the customer's name, surname, and birthday.

### Sample Database

The sample database contains the following entries:

| Name   | Surname | Birthday    |
|--------|---------|-------------|
| Ahmet  | Yılmaz  | 01-01-1990  |
| Ayşe   | Kaya    | 02-02-1985  |
| Mehmet | Demir   | 03-03-1992  |

### Example Usage

```sh
Enter customer's name: Ahmet
Enter customer's surname: Yılmaz
Enter customer's birthday (DD-MM-YYYY): 01-01-1990
Customer validated successfully!
```

### Invalid Example Usage

```sh
Enter customer's name: Ahmet
Enter customer's surname: Kaya
Invalid surname.
```

## Code Explanation

The program consists of the following parts:

- **Sample Database**: Contains hardcoded customer data (names, surnames, and birthdays).
- **Input Handling**: Uses `bufio.NewReader` and `ReadString` to read user input.
- **Trimming Input**: Uses `strings.TrimSpace` to remove any leading or trailing whitespace.
- **Validation**:
  - The program checks if the entered name exists in the `names` slice using the `findIndex` function.
  - If the name is valid, it then checks if the surname at the same index matches the entered surname.
  - If the surname is valid, it checks if the birthday at the same index matches the entered birthday.
  - If all checks are passed, the program prints "Customer validated successfully!". Otherwise, it prints an appropriate error message and terminates.


