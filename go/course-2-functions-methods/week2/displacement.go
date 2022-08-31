package main

import (
	"fmt"
	"math"
)

func readFloat(title string) float64 {
  var userInput float64

  for {
    fmt.Printf("Please enter the %s: ", title)
    _, err := fmt.Scanf("%f", &userInput)

    if err != nil {
      fmt.Printf("Wooops! That's not a float\n")
    } else {
      return userInput
    }
  }
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
  return func(t float64) float64 {
    return (1/2)*a*math.Pow(t, 2) + (v0 * t) + s0
  }
}

func main() {
  var a, v0, s0, t float64

  a = readFloat("acceleration")
  v0 = readFloat("initial velocity")
  s0 = readFloat("initial displacement")
  t = readFloat("time")
  fmt.Println("Initial values")
  fmt.Printf("a: %.4f, v0: %.4f, s0: %.4f, t: %.4f\n", a, v0, s0, t)

  displacementFunction := GenDisplaceFn(a, v0, s0)
  fmt.Printf("Displacement is: %.2f\n", displacementFunction(t))
}