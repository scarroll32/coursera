package main

import (
	"fmt"
	"sort"
)

func main() {
	i := 0

	a := make([]int, 10, 10)
	fmt.Println("Input the numbers")
	for {
		fmt.Scan(&a[i])
		if a[i] == 'x' || a[i] == 'X' {
			break
		}
		i++
		sort.Ints(a)
		fmt.Print(a)
	}
}

