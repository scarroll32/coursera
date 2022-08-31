package main

import (
	"fmt"
	"math"
)

func genDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return (a*math.Pow(t, 2))/2 + v0*t + s0
	}
}

func main() {
	var a, v0, s0, t float64

	fmt.Println("This program calculates displacement.\nEnter acceleration:")
	fmt.Scan(&a)
	fmt.Println("Enter initial velocity:")
	fmt.Scan(&v0)
	fmt.Println("Enter initial displacement:")
	fmt.Scan(&s0)

	displaceFn := genDisplaceFn(a, v0, s0)

	fmt.Println("Enter time:")
	fmt.Scan(&t)

	fmt.Printf("Displacement: %f\n", displaceFn(t))
}

