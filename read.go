package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	var filename string
	var names []name

	fmt.Println("Enter a filename:")
	fmt.Scan(&filename) //input from terminal

	// Open the file.
	f, _ := os.Open(filename)
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file, read the runes, and append to name structs
	for scanner.Scan() {
		line := scanner.Text()

		newName := name{
			fname: strings.Split(line, " ")[0],
			lname: strings.Split(line, " ")[1],
		}

		names = append(names, newName)

	}

	//print out files read in.
	for _, n := range names {
		println(n.fname + " " + n.lname)
	}

}
