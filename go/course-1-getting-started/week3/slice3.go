package main

import (
	"sort"
	"fmt"
)

func main() {
	i := 0

	arr := make([]int, 10, 10)
	fmt.Println("Input the numbers (X to cancel): ")
	for {
		fmt.Scan(&arr[i])
		if arr[i] == 'X' {
			break
		}
		i++
		sort.Ints(arr)
		fmt.Print(arr)
	}
}

