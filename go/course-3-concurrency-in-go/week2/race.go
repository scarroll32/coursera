/* Both add and subtract are accessing the variable x, the `go` function means the outcome in unpredicitible. */

package main

import (
  "fmt"
)

func add(x *int) {
  (*x)++
  fmt.Println(*x)
}

func subtract(x *int) {
  (*x)--
  fmt.Println(*x)
}
func main() {
  y := 0
  go add(&y)
  go subtract(&y)

  y++
  fmt.Println()
}
