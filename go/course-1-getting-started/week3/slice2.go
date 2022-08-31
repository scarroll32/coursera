package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const InitialSliceSize = 3

func main() {

	slice := make([]int, InitialSliceSize)
	i := 1

	for {
		var input string
		var intNumber int

		fmt.Printf("Enter a integer to add to slice, or 'X' for exit: ")

		reader := bufio.NewReader(os.Stdin)
		in, err := reader.ReadString('\n')

		if err != nil {

			log.Print("Unknown input. Type integer or 'X' for exit")
			continue

		} else {

			input = strings.ToLower(strings.TrimSuffix(in, "\n"))
			if input == "x" {
				break
			}

			intNumber, err = strconv.Atoi(input)
			if err != nil {
				continue
			}

			if i <= InitialSliceSize {
				slice[len(slice)-i] = intNumber
				i++
			} else {
				slice = append(slice, intNumber)
			}
			sort.Ints(slice)
			fmt.Println(slice)

		}
	}
}
