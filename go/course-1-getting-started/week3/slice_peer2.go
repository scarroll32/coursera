package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var listint = make([]int64, 3)
	var scanner = bufio.NewScanner(os.Stdin)
	var counter int

	fmt.Print("Enter an interger number: ")
	for scanner.Scan() {
		if scanner.Text() == "X" { // Input of capital X
			fmt.Println("Game Over")
			break
		}

		value, err := strconv.ParseInt(scanner.Text(), 10, 64) // convert input to integer

		if err != nil { // catch conversion error, when user input is no integer
			fmt.Print("Invalid input. Enter an interger number: ")
		} else {
			if counter != len(listint) { // for the first 3 numbers the slice has suffcient lenght
				for i := 0; i < 3; i++ { // since array is sorted below, replace element with value 0 only
					if listint[i] == 0 { // when element value is 0, replace with user input
						listint[i] = value
						break // exit inner for loop
					}
				}
			} else {
				listint = append(listint, value) // start appending element to slice after 3 entries
			}
			sort.SliceStable(listint, func(i, j int) bool { return listint[i] < listint[j] })
			counter++
			fmt.Println("The numbers you entered are:", listint)
			fmt.Print("Enter another integer number or a capital 'X' to end: ")
		}
	}

	if err := scanner.Err(); err != nil { // Output error in case of failure reading from stdin
		fmt.Fprintln(os.Stderr, "Reading from standard input:", err)
	}
}

