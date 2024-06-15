package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Reader function to read user input and return it as an integer
func Reader() (int, error) {
	// Initiate user input reader
	reader := bufio.NewReader(os.Stdin)

	// Print the instruction to the reader in the console
	fmt.Println("Please enter your age:")

	// Call the reader to read user's input
	ageStr, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	// Trim the newline character and convert the string to an integer
	ageStr = strings.TrimSpace(ageStr)
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return 0, err
	}

	return age, nil
}

func main() {

	age, err := Reader()
	if err != nil {
		panic(err)
	}

	switch {
	case age >= 1946 && age <= 1964:
		fmt.Println("Hello, Boomer!")
	case age >= 1964 && age <= 1980:
		fmt.Println("Hello, X!")
	case age >= 1981 && age <= 1996:
		fmt.Println("Hello, Millenial!")
	case age >= 1997 && age <= 2012:
		fmt.Println("Hello, Zoomer!")
	case age >= 2013 && age <= 2024:
		fmt.Println("Hello, Alfa!")
	default:
		fmt.Println("You should be dead")
	}

	date := time.Date(2000, 12, 1, 0, 0, 0, 0, time.UTC)

	fmt.Println(date.Year())
}
