# Password Generator

This tool generates a random 10-digit password each time it is executed. It is written in Go and can be easily used on Mac OS via the zsh terminal..

## Installation
Follow these steps to install the `password-generator` tool on your system:
### Compile the Go Program
```sh
go build password-generator.go
```
## Make the Executable Globally Accessible
Move the password-generator executable to the `/usr/local/bin` directory to make it globally accessible on your system:
```sh
sudo mv password-generator /usr/local/bin/password-generator
```
This command may require administrator privileges.

## Usage
After installing the password-generator tool, you can generate a random password by executing the following command in any terminal window:
```sh
password-generator
```
This command will display a random 10-digit password in the terminal.

## Development

If you wish to make changes to the password-generator, edit the Go file and repeat the installation steps above to apply your changes.


