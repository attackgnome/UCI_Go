package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	name string
	addr string
}

func main() {

	m := make(map[int]Person)

	var name string
	var address string

	fmt.Println("Enter a Name:")
	fmt.Scan(&name) //input from terminal

	fmt.Println("Enter an Address:")
	fmt.Scan(&address) //input from terminal
	p1 := Person{
		name: name,
		addr: address,
	}
	m[1] = p1

	barr, _ := json.Marshal(m)
	fmt.Println(barr)
	fmt.Println(string(barr))

}
