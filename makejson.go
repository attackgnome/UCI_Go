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

	m := make(map[string]int)

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
	m[p1.name] = 1
	m[p1.addr] = 1

	barr, _ := json.Marshal(m)
	fmt.Println(string(barr))

}
