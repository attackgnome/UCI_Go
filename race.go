package main

import "fmt"

func main() {
	counter := 0
	go Add("First", counter)
	go Add("Second", counter)

	//Program exists after main go routine so nothing is printed.

	fmt.Println("Program finished")
}

//printAdd prints a string with a running counter
func Add(some string, counter int) int {
	for i := 1; i < 10; i++ {
		fmt.Println(some, i)
		counter++
	}
	return counter
}
