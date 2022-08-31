package main

import (
	"fmt"
	"os"
	"sort"
	"sync"
)

const partsCount = 4

func sortSlice(sliceToSort []int, wg *sync.WaitGroup, sliceName string) {
	defer wg.Done()

	sort.Ints(sliceToSort)
	fmt.Printf("%s sorted: %v\n", sliceName, sliceToSort)
}

func main() {
	var countIntsToSort int
	fmt.Println("Please specify number of integers you want to sort (min 8):")
	fmt.Scan(&countIntsToSort)
	if countIntsToSort < 8 {
		fmt.Println("There is no sense to split less than 8 elements to 4 parts and sort each part")
		os.Exit(1)
	}

	var intFromUser int
	sliceToSort := make([]int, countIntsToSort)
	for i := 0; i < countIntsToSort; i++ {
		fmt.Printf("Please enter integer number %d:\n", i+1)
		fmt.Scan(&intFromUser)
		sliceToSort[i] = intFromUser
	}
	fmt.Printf("Initial array: %v\n", sliceToSort)

	var partSize int
	partSize = countIntsToSort / partsCount
	indexFrom := 0
	indexTo := partSize
	var wg sync.WaitGroup
	for i := 0; i < partsCount; i++ {
		wg.Add(1)
		partSlice := make([]int, indexTo-indexFrom)
		copy(partSlice, sliceToSort[indexFrom:indexTo])
		go sortSlice(partSlice, &wg, fmt.Sprintf("Part %d", i+1))
		indexFrom = indexTo
		if i == partsCount-2 {
			indexTo = len(sliceToSort)
		} else {
			indexTo += partSize
		}
	}
	wg.Wait()

	wg.Add(1)
	go sortSlice(sliceToSort, &wg, "Full array")
	wg.Wait()
}
