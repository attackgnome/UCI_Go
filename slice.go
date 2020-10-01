package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	var intSlice = make([]int, 3)
	var input string
	var inputInteger int
	var e error

	i := 0

	for {
		fmt.Println("Enter an integer:")
		fmt.Scan(&input) //input from terminal
		if input == "X" {
			break
		}
		inputInteger, e = strconv.Atoi(input) //cast to lower case
		if e == nil {
			intSlice = append(intSlice, inputInteger)
		}
		sort.Ints(intSlice)
		fmt.Println(intSlice)
		i = i + 1

	}

}
