package main

import (
  "fmt"
)

const LENGTH = 10

func Swap(elems []int, i int) {
  elems[i], elems[i+1] = elems[i+1], elems[i]
}

func BubbleSort(elems []int) {
  swapped := true
  for swapped {
    swapped = false
    for i := 0; i < len(elems) -1 ; i++ {
      if elems[i] > elems[i+1] {
        Swap(elems, i)
        swapped = true
      }
    }
  }
}

func readNumber() int {
  var i int
  for {
    fmt.Printf("Please enter an integer: ")
    _, err := fmt.Scanf("%d", &i)
    if err == nil {
      return i
    }
  }
}

func main(){
  elems := make([]int, LENGTH)
  for i := 0; i < LENGTH; i++ {
    elems[i] = readNumber()
  }
  fmt.Printf("You entered: %v\n", elems)
  BubbleSort(elems)
  fmt.Printf("Sorted: %v\n", elems)
}
