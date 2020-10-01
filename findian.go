package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Println("Enter a string:")
	fmt.Scan(&input)               //input from terminal
	input = strings.ToLower(input) //cast to lower case

	if strings.HasPrefix(input, "i") && strings.Contains(input, "a") && strings.HasSuffix(input, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
