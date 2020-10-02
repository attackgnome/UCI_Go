package main

import (
	"fmt"
	"io/ioutil"
)
package main

import (
	"bufio"
	"fmt"
	"os"
)

type name struct {
	fname string
	lname string
}

func main() {
	var filename string
	var err error
	var names name[]

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
		  fname: line[:19],
		  lname: line[20:],
	  }

		  names = append(names, newName)

	}
}

//Your program will define a name struct
//hich has two fields, fname for the first name, and
//lname for the last name. Each field will be a string of size 20 (characters).
