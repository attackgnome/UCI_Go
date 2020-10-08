package main

import (
	"fmt"
)

//Animal is a structure for hold the characteristics of an animal
type Animal interface {
	Eat()
	Move()
	Speak()
}

//Cow type
type Cow struct{ name string }

//Eat method for cow type
func (cow Cow) Eat() { fmt.Println("grass") }

//Move method for cow type
func (cow Cow) Move() { fmt.Println("walk") }

//Speak method for cow type
func (cow Cow) Speak() { fmt.Println("moo") }

//Bird type
type Bird struct{ name string }

//Eat method for bird type
func (bird Bird) Eat() { fmt.Println("worms") }

//Move method for bird type
func (bird Bird) Move() { fmt.Println("fly") }

//Speak method for bird type
func (bird Bird) Speak() { fmt.Println("peep") }

//Snake type
type Snake struct{ name string }

//Eat method for snake type
func (snake Snake) Eat() { fmt.Println("mice") }

//Move method for snake type
func (snake Snake) Move() { fmt.Println("slither") }

//Speak method for snake type
func (snake Snake) Speak() { fmt.Println("hsss") }

func main() {
	var command string
	var name string
	var action string
	animals := make(map[string]Animal)

	for {
		fmt.Println(">")
		fmt.Scanln(&command, &name, &action) //input from terminal

		if command == "newanimal" {
			if action == "cow" {
				animal := Cow{name}
				animals[name] = animal
			}
			if action == "bird" {
				animal := Bird{name}
				animals[name] = animal
			}
			if action == "snake" {
				animal := Snake{name}
				animals[name] = animal
			}
			fmt.Println("Created it!")
		}
		if command == "query" {

			if action == "eat" {
				animals[name].Eat()
			}
			if action == "move" {
				animals[name].Move()
			}
			if action == "speak" {
				animals[name].Speak()
			}

		}
	}
}
