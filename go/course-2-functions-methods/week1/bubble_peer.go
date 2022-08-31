package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const totalItems = 10

func main() {
	items, err := getItems()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Original slice: %v \n", items)

	bubblesort(items)
	fmt.Printf("Sorted slice: %v \n", items)
}

func bubblesort(items []int) {
	var (
		n      = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				swap(items, i)
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}

func swap(items []int, position int) {
	items[position+1], items[position] = items[position], items[position+1]
}

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func getItems() ([]int, error) {
	fmt.Print("Press Enter to generate random items or type integers separated by space: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Trim(scanner.Text(), " \r\n\t")
	if input == "" {
		return generateSlice(totalItems), nil
	}
	items, parseError := parseInput(input)
	if parseError != nil {
		return nil, parseError
	}
	return items, nil
}

func parseInput(input string) ([]int, error) {
	stringItems := strings.Split(input, " ")
	intItems := make([]int, 0, len(stringItems))
	for _, value := range stringItems {
		if value == "" {
			continue
		}
		iValue, convError := strconv.Atoi(value)
		if convError != nil {
			return nil, errors.New("Error: " + value + " is not integer")
		}
		intItems = append(intItems, iValue)
	}
	return intItems, nil
}

