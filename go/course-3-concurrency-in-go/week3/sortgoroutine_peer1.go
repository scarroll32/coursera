package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const divnum = 4

var finalpieces []int

func main() {
	var c = make(chan []int, 2)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter integers separated by space, when finished press [enter]:")
	snums, _ := reader.ReadString('\n')
	snums = strings.TrimSuffix(snums, "\n")
	nums := strings.Fields(snums)

	theOneArray := make([]int, len(nums))
	copy(theOneArray, makeIntArray(nums))

	n := len(theOneArray)
	splitsize := int(math.Ceil(float64(n) / float64(divnum)))

	for i := 0; i < divnum-1; i++ {
		carr := make([]int, splitsize)
		copy(carr, theOneArray[i*splitsize:i*splitsize+splitsize])
		go sendSubArray(carr, c)
	}
	carr := make([]int, n-(splitsize*(divnum-1)))
	copy(carr, theOneArray[(divnum-1)*splitsize:])
	go sendSubArray(carr, c)

	for i := 0; i < divnum; i++ {
		one := <-c
		finalpieces = append(finalpieces, one...)
	}
	fmt.Println("final ordering", orderSlice(finalpieces))
}

func makeIntArray(arr []string) []int {
	var res []int
	for _, s := range arr {
		tx, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Failed at %v", err)
			os.Exit(1)
		}
		res = append(res, tx)
	}
	return res
}

func sendSubArray(arr []int, c chan []int) {
	fmt.Println("Im goroutine and will sort:", arr)
	res := orderSlice(arr)
	c <- res
}

func orderSlice(arr []int) []int {
	n := len(arr)
	for n > 1 {
		newn := 0
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				swap(arr, i-1, i)
				newn = i
			}
		}
		n = newn
	}
	return arr
}

func swap(s []int, ia int, ib int) {
	var temp int
	temp = s[ia]
	s[ia] = s[ib]
	s[ib] = temp
}
