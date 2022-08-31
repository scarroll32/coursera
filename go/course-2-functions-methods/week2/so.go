package main

import "fmt"

func readFloat(title string) float64 {
  var userInput float64

  for {
    fmt.Println("Please enter a float: ")
    _, err := fmt.Scanf("%f", &userInput)

    if err != nil {
      fmt.Printf("Wooops! That's not a float\n")
    } else {
      return userInput
    }
  }
}

func main() {
  var f float64

  f = readFloat("acceleration")
	fmt.Printf("You entered: %.04f\n", f)
}
