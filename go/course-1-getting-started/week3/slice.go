package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
  "strings"
  "sort"
)

const SLICE_SIZE = 3

func main() {
  sortedSlice := make([]int, SLICE_SIZE)

  var userInput string
  var inputVal int
  reader := bufio.NewReader(os.Stdin)

  for idx := 1;; {
    fmt.Printf("Please enter an integer or 'X' to quit: ")
    in, err := reader.ReadString('\n')

    userInput = strings.TrimSuffix(in, "\n")
    if userInput == "X" {
      break
    }

    inputVal, err = strconv.Atoi(userInput)
    if err != nil {
      fmt.Print("Not an integer or 'X'\n")
      continue
    }

    if idx > SLICE_SIZE {
      sortedSlice = append(sortedSlice, inputVal)
    } else {
      sortedSlice[len(sortedSlice) - idx] = inputVal
      idx++
    }

    sort.Sort(sort.IntSlice(sortedSlice))
    fmt.Println("Slice Contents: ", sortedSlice)
  }
}
