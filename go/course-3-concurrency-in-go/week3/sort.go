package main

import (
	"fmt"
	"os"
  "sync"
	"sort"
)

func sortSlice(wg *sync.WaitGroup, slice []int) {
	fmt.Println("Sorting: ", slice)
	sort.Ints(slice)
	wg.Done()
}

func main() {
	var inputNums []int
	var wg sync.WaitGroup

	for {
		fmt.Print("Enter a single int > ")

		var num int
		_, err := fmt.Fscanf(os.Stdin, "%d", &num)

		if err != nil {
			break
		}

		inputNums = append(inputNums, num)
	}

	// nums should now contain the integers
	inputNumsLen := len(inputNums)
	partitionSize := inputNumsLen / 4
	remainder := inputNumsLen % 4

  fmt.Println(inputNums)
  fmt.Printf("inputNumsLen: %d\n", inputNumsLen)
  fmt.Printf("partitionSize: %d\n", partitionSize)
  fmt.Printf("remainder: %d\n", remainder)

  part1 := inputNums[:partitionSize]
  part2 := inputNums[partitionSize : partitionSize*2]
	part3 := inputNums[partitionSize*2 : partitionSize*3]
	part4 := inputNums[partitionSize*3:]

  fmt.Printf("part1: %d",part1)
  fmt.Printf("part2: %d",part2)
  fmt.Printf("part3: %d",part3)
  fmt.Printf("part4: %d",part4)

	// sort each array concurrently
  wg.Add(4)
	go sortSlice(&wg, part1)
	go sortSlice(&wg, part2)
	go sortSlice(&wg, part3)
	go sortSlice(&wg, part4)
  wg.Wait()

  fmt.Println("All Goroutine Done!")
  fmt.Println("Goroutine 1:", part1)
  fmt.Println("Goroutine 2:", part2)
  fmt.Println("Goroutine 3:", part3)
  fmt.Println("Goroutine 4:", part4)

	// Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.
	sort.Ints(inputNums)
	fmt.Println("Merged and sorted: ",inputNums)
}
