package main

import (
	"fmt"
	"strconv"
)

func main() {
	var intEntry string
	var entryNum string
	var ints []int

	//Prompt the user to input the number of integers to be sorted
	fmt.Println("Enter a the number of Integers in the sequence")
	fmt.Scan(&entryNum)
	intNum, _ := strconv.Atoi(entryNum)

	//Loop to prompt for integer entry and store the inputs to a slice
	for i := 1; i <= intNum; i++ {
		fmt.Println("Enter an Integer: ")
		fmt.Scan(&intEntry)
		intNum, _ := strconv.Atoi(intEntry)
		ints = append(ints, intNum)
	}
	//Call Bubblesort and print the output
	BubbleSort(ints)
	fmt.Print(ints)

}

//Swap is a function to swap the contents of the slice in
//position i with the contents in position i+1.
func Swap(sli []int, i int) {
	var hold int
	hold = sli[i+1]
	sli[i+1] = sli[i]
	sli[i] = hold
}

//BubbleSort takes a slice of integers as an argument and
//modifies the slice so that the elements are in sorted order.
func BubbleSort(sli []int) {
	swapped := true //tracker for if elements are being swapped
	for swapped {
		// set swapped to false
		swapped = false
		for i := 0; i < len(sli)-1; i++ {
			if sli[i] > sli[i+1] {
				Swap(sli, i)
				swapped = true
			}
		}
	}
}
