package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	chunkNo := 4 //split into four groups

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter the value: ")
	a, _ := reader.ReadString('\n')

	s := strings.Split(strings.TrimSpace(a), " ")
	arr := make([]int, 0, len(s))

	for _, str := range s {
		n, _ := strconv.Atoi(str)
		arr = append(arr, n)
	}
	fmt.Println(arr)

	/*Partition the array/slice into 4 broken parts*/
	sliceSize := len(arr) / chunkNo
	slice1 := arr[:sliceSize]
	slice2 := arr[sliceSize : 2*(sliceSize)]
	slice3 := arr[2*(sliceSize) : 3*(sliceSize)]
	slice4 := arr[3*(sliceSize):]

	fmt.Println("Here are your arrays when partitioned", slice1, slice2, slice3, slice4)

	wg.Add(chunkNo) //Create a waitGroup that executes 4 concurrent goroutines
	go sortList(slice1)
	wg.Done()
	go sortList(slice2)
	wg.Done()
	go sortList(slice3)
	wg.Done()
	go sortList(slice4)
	wg.Done()
	wg.Wait()

	/*Merge the Slices into one slice and sort them in main goroutine*/
	newSlice := mergeAndSort(slice1, slice2, slice3, slice4)

	//Print the new slice in the goroutine
	fmt.Println("Your Integers merged and sorted is follows:", newSlice)
}

// sortList takes a slice of integers and returns a sorted slice
func sortList(unsortedSlice []int) []int {
	sort.Ints(unsortedSlice)
	return unsortedSlice
}

//mergeAndSort merges the 4 intermediary slices together
func mergeAndSort(list1 []int, list2 []int, list3 []int, list4 []int) []int {
	newSlice := []int{}
	newSlice = append(list1, list2...)
	newSlice = append(newSlice, list3...)
	newSlice = append(newSlice, list4...)
	sort.Ints(newSlice)
	return newSlice
}
