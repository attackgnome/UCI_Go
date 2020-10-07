package main

import (
	"fmt"
)

//Animal is a structure for hold the characteristics of an animal
type Animal struct {
	food  string
	move  string
	speak string
}

func main() {
	var request string
	var action string
	var animal Animal

	for {
		fmt.Println(">")
		fmt.Scanln(&request, &action) //input from terminal

		if request == "cow" {
			animal.InitAnimal("grass", "walk", "moo")
		} else if request == "bird" {
			animal.InitAnimal("worms", "fly", "peep")
		} else if request == "snake" {
			animal.InitAnimal("mice", "slither", "hsss")
		}

		if action == "eat" {
			animal.Eat()
		} else if action == "move" {
			animal.Move()
		} else if action == "speak" {
			animal.Speak()
		}

	}
}

//InitAnimal is to intialize the instance of an animal
func (animal *Animal) InitAnimal(food, move, speak string) {
	animal.food = food
	animal.move = move
	animal.speak = speak
}

//Eat prints the animals food
func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

//Move prints the animals method of locomotion
func (animal *Animal) Move() {
	fmt.Println(animal.move)
}

//Speak prints the sound the animal makes
func (animal *Animal) Speak() {
	fmt.Println(animal.speak)
}
