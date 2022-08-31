package main
import (
	"fmt"
)

func main() {
  var userInput float64
  fmt.Println("Please enter a float: ")
	_, err := fmt.Scanf("%f", &userInput)

	if err != nil {
	  fmt.Printf("Sorry that is not a valid floating point number\n")
	} else {
    fmt.Printf("You entered: %f \n", userInput)
		fmt.Printf("Truncated: %d \n", int64(userInput))
	}

}
