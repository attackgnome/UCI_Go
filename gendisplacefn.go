package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var vO float64
	var sO float64
	var t float64

	fmt.Println("Enter a value for acceleration: ")
	fmt.Scan(&a) //input from terminal

	fmt.Println("Enter a value for velocity: ")
	fmt.Scan(&vO) //input from terminal

	fmt.Println("Enter a value for displacement: ")
	fmt.Scan(&sO) //input from terminal

	fn := GenDisplaceFn(a, vO, sO)

	fmt.Println("Enter a value for time: ")
	fmt.Scan(&t) //input from terminal

	fmt.Println(fn(t))

}

// GenDisplaceFn computes displace as a function of time
func GenDisplaceFn(a, vO, sO float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + vO*t + sO
	}
	return fn
}
