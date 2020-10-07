package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var v_o float64
	var s_o float64
	var t float64

	fmt.Println("Enter a value for acceleration: ")
	fmt.Scan(&a) //input from terminal

	fmt.Println("Enter a value for velocity: ")
	fmt.Scan(&v_o) //input from terminal

	fmt.Println("Enter a value for displacement: ")
	fmt.Scan(&s_o) //input from terminal

	fn := GenDisplaceFn(a, v_o, s_o)

	fmt.Println("Enter a value for time: ")
	fmt.Scan(&t) //input from terminal

	fmt.Println(fn(t))

}

// GenDisplaceFn computes displace as a function of time
func GenDisplaceFn(a, v_o, s_o float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v_o*t + s_o
	}
	return fn
}
