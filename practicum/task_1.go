package main

import (
	"fmt"
	"unicode/utf8"
)

var str string

func main() {
	str = "Hello, world!"

	fmt.Println(utf8.RuneCountInString(str))
	fmt.Println(string(str[0]))

	slice := []rune(str)
	slice[5] = 101
	for _, value := range slice {
		fmt.Print(string(value))
	}
}
