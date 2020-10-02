package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var intEntry string
	var out []string

	fmt.Println("Enter a sequence of up to 10 integers seperated by spaces: ")
	fmt.Scan(&intEntry) //input from terminal

	out = strings.Split(intEntry, " ")

	ints := make([]int, len(out))
	fmt.Println(out)

	for i, s := range out {
		ints[i], _ = strconv.Atoi(s)
	}

	fmt.Println(ints)

}

//BubbleSort() takes a slice of integers as an argument and
//modifies the slice so that the elements are in sorted order.
func Swap(sli []int, i int) {
	var hold int
	hold = sli[i+1]
	sli[i+1] = sli[i]
	sli[i] = hold
}

//
func BubbleSort(sli []int) {

	for i, num := range sli {
		if num > sli[i+1] {
			Swap(sli, i)
		}
	}
}
